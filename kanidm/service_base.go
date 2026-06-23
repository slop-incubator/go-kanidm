package kanidm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// serviceBase holds the shared HTTP client and credentials used by all
// service wrappers. It is intentionally unexported — callers use the typed
// service structs (AccountsService, GroupsService, …) instead.
type serviceBase struct {
	httpClient *http.Client
	baseURL    string
	token      string // bearer token; never logged
}

// get executes a GET request and decodes the JSON response into dst.
func (s *serviceBase) get(ctx context.Context, path string, dst any) error {
	req, err := s.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}
	return s.do(req, dst)
}

// post executes a POST request with a JSON body and decodes the response.
func (s *serviceBase) post(ctx context.Context, path string, body, dst any) error {
	req, err := s.newRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return err
	}
	return s.do(req, dst)
}

// patch executes a PATCH request with a JSON body and decodes the response.
func (s *serviceBase) patch(ctx context.Context, path string, body, dst any) error {
	req, err := s.newRequest(ctx, http.MethodPatch, path, body)
	if err != nil {
		return err
	}
	return s.do(req, dst)
}

// delete executes a DELETE request (no response body expected).
func (s *serviceBase) delete(ctx context.Context, path string) error {
	req, err := s.newRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return s.do(req, nil)
}

// newRequest builds an authenticated HTTP request. The Authorization header is
// set here and nowhere else, ensuring it is not accidentally duplicated or logged.
func (s *serviceBase) newRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("kanidm: encoding request body: %w", err)
		}
		bodyReader = bytes.NewReader(encoded)
	}

	req, err := http.NewRequestWithContext(ctx, method, s.baseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("kanidm: building request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// do executes the request and handles non-2xx responses. If dst is non-nil,
// the response body is decoded into it.
func (s *serviceBase) do(req *http.Request, dst any) error {
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("kanidm: executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return wrapHTTPError(resp.StatusCode, string(bodyBytes))
	}

	if dst != nil && resp.StatusCode != http.StatusNoContent {
		if err := decodeJSON(resp.Body, dst); err != nil {
			return fmt.Errorf("kanidm: decoding response: %w", err)
		}
	}
	return nil
}

// decodeJSON is a small helper that returns a clear error on invalid JSON.
func decodeJSON(r io.Reader, dst any) error {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		return fmt.Errorf("JSON decode: %w", err)
	}
	return nil
}
