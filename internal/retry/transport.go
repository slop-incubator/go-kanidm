// Package retry provides an http.RoundTripper that retries transient HTTP
// errors with exponential backoff.
package retry

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

// Transport wraps an http.RoundTripper with configurable retry logic.
//
// Retried status codes: 429 (Too Many Requests), 502, 503, 504.
// Never retried: any 4xx other than 429 (auth errors, bad requests, etc.).
// Never retried: non-idempotent methods with a body (POST, PATCH) — this
// avoids replaying writes that may have partially succeeded. Retries are safe
// for GET, HEAD, DELETE, and idempotent PUT.
type Transport struct {
	// Base is the underlying RoundTripper. Required.
	Base http.RoundTripper

	// MaxRetries is the maximum number of additional attempts after the first.
	// Defaults to 3 if zero.
	MaxRetries int

	// BaseDelay is the initial backoff duration. Each subsequent attempt
	// doubles the delay (capped at 30s). Defaults to 500ms if zero.
	BaseDelay time.Duration
}

// retryableError records that the last attempt returned a retryable status.
type retryableError struct {
	StatusCode int
}

func (e *retryableError) Error() string {
	return fmt.Sprintf("retry: server returned %d", e.StatusCode)
}

func (t *Transport) maxRetries() int {
	if t.MaxRetries > 0 {
		return t.MaxRetries
	}
	return 3
}

func (t *Transport) baseDelay() time.Duration {
	if t.BaseDelay > 0 {
		return t.BaseDelay
	}
	return 500 * time.Millisecond
}

const maxDelay = 30 * time.Second

// RoundTrip implements http.RoundTripper.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Only retry safe/idempotent methods to avoid replaying writes.
	if !isSafeMethod(req.Method) {
		return t.Base.RoundTrip(req)
	}

	var lastErr error
	max := t.maxRetries()

	for attempt := 0; attempt <= max; attempt++ {
		if attempt > 0 {
			delay := t.delayFor(attempt)
			select {
			case <-time.After(delay):
			case <-req.Context().Done():
				return nil, req.Context().Err()
			}
		}

		resp, err := t.Base.RoundTrip(req)
		if err != nil {
			// Network-level error — retry if context is still alive.
			lastErr = err
			continue
		}

		if !shouldRetry(resp.StatusCode) {
			return resp, nil
		}

		// Drain and close the body before the next attempt so the
		// connection can be reused.
		resp.Body.Close()
		lastErr = &retryableError{StatusCode: resp.StatusCode}
	}

	return nil, lastErr
}

func (t *Transport) delayFor(attempt int) time.Duration {
	delay := time.Duration(math.Pow(2, float64(attempt-1))) * t.baseDelay()
	if delay > maxDelay {
		delay = maxDelay
	}
	return delay
}

// shouldRetry reports whether a status code warrants a retry.
func shouldRetry(code int) bool {
	switch code {
	case http.StatusTooManyRequests,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	}
	return false
}

// isSafeMethod reports whether the HTTP method is safe to retry.
// POST and PATCH are excluded because they may have side effects.
func isSafeMethod(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodDelete, http.MethodPut, http.MethodOptions:
		return true
	}
	return false
}
