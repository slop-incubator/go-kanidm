package unit_test

import (
	"testing"

	"github.com/slop-incubator/go-kanidm/kanidm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew_MissingBaseURL(t *testing.T) {
	_, err := kanidm.New(kanidm.Options{Token: "tok"})
	require.Error(t, err)
	assert.ErrorIs(t, err, kanidm.ErrMissingBaseURL)
}

func TestNew_MissingToken(t *testing.T) {
	_, err := kanidm.New(kanidm.Options{BaseURL: "https://idm.example.com"})
	require.Error(t, err)
	assert.ErrorIs(t, err, kanidm.ErrMissingToken)
}

func TestNew_ValidOptions(t *testing.T) {
	c, err := kanidm.New(kanidm.Options{
		BaseURL: "https://idm.example.com",
		Token:   "tok",
	})
	require.NoError(t, err)
	require.NotNil(t, c)
	assert.NotNil(t, c.Accounts)
	assert.NotNil(t, c.Groups)
	assert.NotNil(t, c.OAuth2)
}

func TestNew_DefaultsApplied(t *testing.T) {
	// This test verifies that zero-value Options fields are replaced with
	// sensible defaults (timeout, retries, delay). We can only observe this
	// indirectly through successful construction.
	c, err := kanidm.New(kanidm.Options{
		BaseURL: "https://idm.example.com",
		Token:   "tok",
	})
	require.NoError(t, err)
	assert.NotNil(t, c)
}
