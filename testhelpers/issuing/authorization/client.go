//
//
// File generated from our OpenAPI spec
//
//

// Package authorization provides the /v1/issuing/authorizations APIs
package authorization

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/authorizations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a test-mode authorization.
func New(params *stripe.TestHelpersIssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil,

		// Create a test-mode authorization.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TestHelpersIssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture a test-mode authorization.
func Capture(id string, params *stripe.TestHelpersIssuingAuthorizationCaptureParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture a test-mode authorization.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Capture(id string, params *stripe.TestHelpersIssuingAuthorizationCaptureParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expire a test-mode Authorization.
func Expire(id string, params *stripe.TestHelpersIssuingAuthorizationExpireParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expire a test-mode Authorization.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Expire(id string, params *stripe.TestHelpersIssuingAuthorizationExpireParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finalize the amount on an Authorization prior to capture, when the initial authorization was for an estimated amount.
func FinalizeAmount(id string, params *stripe.TestHelpersIssuingAuthorizationFinalizeAmountParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finalize the amount on an Authorization prior to capture, when the initial authorization was for an estimated amount.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) FinalizeAmount(id string, params *stripe.TestHelpersIssuingAuthorizationFinalizeAmountParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Increment a test-mode Authorization.
func Increment(id string, params *stripe.TestHelpersIssuingAuthorizationIncrementParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Increment a test-mode Authorization.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Increment(id string, params *stripe.TestHelpersIssuingAuthorizationIncrementParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Respond to a fraud challenge on a testmode Issuing authorization, simulating either a confirmation of fraud or a correction of legitimacy.
func Respond(id string, params *stripe.TestHelpersIssuingAuthorizationRespondParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Respond to a fraud challenge on a testmode Issuing authorization, simulating either a confirmation of fraud or a correction of legitimacy.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Respond(id string, params *stripe.TestHelpersIssuingAuthorizationRespondParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reverse a test-mode Authorization.
func Reverse(id string, params *stripe.TestHelpersIssuingAuthorizationReverseParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reverse a test-mode Authorization.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Reverse(id string, params *stripe.TestHelpersIssuingAuthorizationReverseParams) (*stripe.IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
