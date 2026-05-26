//
//
// File generated from our OpenAPI spec
//
//

// Package verificationsession provides the /v1/identity/verification_sessions APIs
package verificationsession

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/identity/verification_sessions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a VerificationSession object.
//
// After the VerificationSession is created, display a verification modal using the session client_secret or send your users to the session's url.
//
// If your API key is in test mode, verification checks won't actually process, though everything else will occur as if in live mode.
//
// Related guide: [Verify your users' identity documents](https://docs.stripe.com/docs/identity/verify-identity-documents)
func New(params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a VerificationSession object.
		//
		// After the VerificationSession is created, display a verification modal using the session client_secret or send your users to the session's url.
		//
		// If your API key is in test mode, verification checks won't actually process, though everything else will occur as if in live mode.
		//
		// Related guide: [Verify your users' identity documents](https://docs.stripe.com/docs/identity/verify-identity-documents)
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a VerificationSession that was previously created.
//
// When the session status is requires_input, you can use this method to retrieve a valid
// client_secret or url to allow re-submission.
func Get(id string, params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of a VerificationSession that was previously created.
	//
	// When the session status is requires_input, you can use this method to retrieve a valid
	// client_secret or url to allow re-submission.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a VerificationSession object.
//
// When the session status is requires_input, you can use this method to update the
// verification check and options.
func Update(id string, params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a VerificationSession object.
//
// When the session status is requires_input, you can use this method to update the
// verification check and options.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A VerificationSession object can be canceled when it is in requires_input [status](https://docs.stripe.com/docs/identity/how-sessions-work).
//
// Once canceled, future submission attempts are disabled. This cannot be undone. [Learn more](https://docs.stripe.com/docs/identity/verification-sessions#cancel).
func Cancel(id string, params *stripe.IdentityVerificationSessionCancelParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A VerificationSession object can be canceled when it is in requires_input [status](https://docs.stripe.com/docs/identity/how-sessions-work).
//
// Once canceled, future submission attempts are disabled. This cannot be undone. [Learn more](https://docs.stripe.com/docs/identity/verification-sessions#cancel).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.IdentityVerificationSessionCancelParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Redact a VerificationSession to remove all collected information from Stripe. This will redact
// the VerificationSession and all objects related to it, including VerificationReports, Events,
// request logs, etc.
//
// A VerificationSession object can be redacted when it is in requires_input or verified
// [status](https://docs.stripe.com/docs/identity/how-sessions-work). Redacting a VerificationSession in requires_action
// state will automatically cancel it.
//
// The redaction process may take up to four days. When the redaction process is in progress, the
// VerificationSession's redaction.status field will be set to processing; when the process is
// finished, it will change to redacted and an identity.verification_session.redacted event
// will be emitted.
//
// Redaction is irreversible. Redacted objects are still accessible in the Stripe API, but all the
// fields that contain personal data will be replaced by the string [redacted] or a similar
// placeholder. The metadata field will also be erased. Redacted objects cannot be updated or
// used for any purpose.
//
// [Learn more](https://docs.stripe.com/docs/identity/verification-sessions#redact).
func Redact(id string, params *stripe.IdentityVerificationSessionRedactParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Redact a VerificationSession to remove all collected information from Stripe. This will redact
// the VerificationSession and all objects related to it, including VerificationReports, Events,
// request logs, etc.
//
// A VerificationSession object can be redacted when it is in requires_input or verified
// [status](https://docs.stripe.com/docs/identity/how-sessions-work). Redacting a VerificationSession in requires_action
// state will automatically cancel it.
//
// The redaction process may take up to four days. When the redaction process is in progress, the
// VerificationSession's redaction.status field will be set to processing; when the process is
// finished, it will change to redacted and an identity.verification_session.redacted event
// will be emitted.
//
// Redaction is irreversible. Redacted objects are still accessible in the Stripe API, but all the
// fields that contain personal data will be replaced by the string [redacted] or a similar
// placeholder. The metadata field will also be erased. Redacted objects cannot be updated or
// used for any purpose.
//
// [Learn more](https://docs.stripe.com/docs/identity/verification-sessions#redact).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Redact(id string, params *stripe.IdentityVerificationSessionRedactParams) (*stripe.IdentityVerificationSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of VerificationSessions
func List(params *stripe.IdentityVerificationSessionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of VerificationSessions
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.IdentityVerificationSessionListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for identity verification sessions.
type Iter struct {
	*stripe.Iter
}

// IdentityVerificationSession returns the identity verification session which the iterator is currently pointing to.
func (i *Iter) IdentityVerificationSession() *stripe.IdentityVerificationSession {
	_ = "STUB: not implemented"
	return nil
}

// IdentityVerificationSessionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IdentityVerificationSessionList() *stripe.IdentityVerificationSessionList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
