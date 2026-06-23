package kanidm

import (
	"errors"
	"fmt"
	"net/http"
)

// Sentinel errors that callers can match with errors.Is.
var (
	ErrMissingBaseURL        = errors.New("kanidm: BaseURL is required")
	ErrMissingToken          = errors.New("kanidm: Token is required")
	ErrUnauthorized          = errors.New("kanidm: authentication failed — check token validity")
	ErrForbidden             = errors.New("kanidm: permission denied")
	ErrNotFound              = errors.New("kanidm: resource not found")
	ErrConflict              = errors.New("kanidm: resource already exists")
	ErrSchemaVersionMismatch = errors.New("kanidm: client schema version does not match server")
)

// APIError wraps a non-2xx HTTP response. Callers can use errors.As to
// inspect the status code while still using errors.Is against sentinels.
type APIError struct {
	StatusCode int
	Body       string
	Err        error // one of the sentinel errors above, or nil
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("kanidm API error %d: %s", e.StatusCode, e.Err)
	}
	return fmt.Sprintf("kanidm API error %d: %s", e.StatusCode, e.Body)
}

func (e *APIError) Unwrap() error { return e.Err }

// wrapHTTPError maps an HTTP status code to a typed APIError.
// body is used verbatim for error messages but never logged.
func wrapHTTPError(statusCode int, body string) error {
	base := &APIError{StatusCode: statusCode, Body: body}
	switch statusCode {
	case http.StatusUnauthorized:
		base.Err = ErrUnauthorized
	case http.StatusForbidden:
		base.Err = ErrForbidden
	case http.StatusNotFound:
		base.Err = ErrNotFound
	case http.StatusConflict:
		base.Err = ErrConflict
	}
	return base
}
