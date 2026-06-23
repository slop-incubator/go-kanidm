package kanidm

import (
	"net/http"

	"github.com/google/uuid"
)

const requestIDHeader = "X-Kanidm-Request-ID"

// requestIDTransport injects a unique request ID header into every outgoing
// request. This enables server-side tracing without leaking user identity.
// It never logs the Authorization header or any credential.
type requestIDTransport struct {
	base      http.RoundTripper
	idFunc    func() string
}

func newRequestIDTransport(base http.RoundTripper, idFunc func() string) http.RoundTripper {
	if idFunc == nil {
		idFunc = func() string { return uuid.NewString() }
	}
	return &requestIDTransport{base: base, idFunc: idFunc}
}

func (t *requestIDTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request so we don't mutate the caller's copy.
	r := req.Clone(req.Context())
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(requestIDHeader, t.idFunc())
	return t.base.RoundTrip(r)
}
