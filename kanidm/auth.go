package kanidm

// This file intentionally does not implement password-based authentication.
//
// Kanidm service accounts authenticate via bearer tokens generated out-of-band
// using the Kanidm CLI:
//
//	kanidm service-account api-token generate <account> <label>
//
// Embedding passwords in application code violates the principle of least
// privilege and makes secret rotation harder. Distribute tokens via environment
// variables or a secrets manager (Vault, AWS Secrets Manager, etc.).
//
// For token rotation in long-running services, create a new token before
// expiry using the Kanidm CLI or management API, update your secrets store,
// and construct a new Client with New(). The old client can be abandoned once
// in-flight requests complete.
//
// See: https://kanidm.github.io/kanidm/stable/service_accounts.html
