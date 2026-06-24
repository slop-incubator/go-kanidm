package retry_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/slop-incubator/go-kanidm/internal/retry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransport_SuccessOnFirstAttempt(t *testing.T) {
	calls := atomic.Int32{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	tr := &retry.Transport{Base: http.DefaultTransport, MaxRetries: 3, BaseDelay: time.Millisecond}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(srv.URL)
	require.NoError(t, err)
	resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(1), calls.Load(), "should not retry on success")
}

func TestTransport_RetriesOn503(t *testing.T) {
	const wantAttempts = 4 // 1 initial + 3 retries
	calls := atomic.Int32{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	tr := &retry.Transport{Base: http.DefaultTransport, MaxRetries: 3, BaseDelay: time.Millisecond}
	client := &http.Client{Transport: tr}

	_, err := client.Get(srv.URL)
	// All attempts exhausted — error expected.
	require.Error(t, err)
	assert.Equal(t, int32(wantAttempts), calls.Load())
}

func TestTransport_DoesNotRetryOn401(t *testing.T) {
	calls := atomic.Int32{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer srv.Close()

	tr := &retry.Transport{Base: http.DefaultTransport, MaxRetries: 3, BaseDelay: time.Millisecond}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(srv.URL)
	require.NoError(t, err)
	resp.Body.Close()
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, int32(1), calls.Load(), "must not retry auth errors")
}

func TestTransport_DoesNotRetryPOST(t *testing.T) {
	calls := atomic.Int32{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	tr := &retry.Transport{Base: http.DefaultTransport, MaxRetries: 3, BaseDelay: time.Millisecond}
	client := &http.Client{Transport: tr}

	resp, err := client.Post(srv.URL, "application/json", http.NoBody)
	require.NoError(t, err)
	resp.Body.Close()
	// POST is not retried even on 503.
	assert.Equal(t, int32(1), calls.Load(), "must not retry non-safe methods")
}

func TestTransport_ContextCancellation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	tr := &retry.Transport{Base: http.DefaultTransport, MaxRetries: 10, BaseDelay: 50 * time.Millisecond}
	client := &http.Client{Transport: tr}

	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL, nil)
	_, err := client.Do(req)
	require.Error(t, err)
	// Should fail with a context error, not exhaust all retries.
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestTransport_SucceedsAfterTransientFailure(t *testing.T) {
	calls := atomic.Int32{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := calls.Add(1)
		if n < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	tr := &retry.Transport{Base: http.DefaultTransport, MaxRetries: 3, BaseDelay: time.Millisecond}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(srv.URL)
	require.NoError(t, err)
	resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(3), calls.Load())
}
