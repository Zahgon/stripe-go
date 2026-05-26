//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersIssuingAuthorizationService is used to invoke /v1/issuing/authorizations APIs.
type v1TestHelpersIssuingAuthorizationService struct {
	B   Backend
	Key string
}

// Create a test-mode authorization.
func (c v1TestHelpersIssuingAuthorizationService) Create(ctx context.Context, params *TestHelpersIssuingAuthorizationCreateParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture a test-mode authorization.
func (c v1TestHelpersIssuingAuthorizationService) Capture(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationCaptureParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expire a test-mode Authorization.
func (c v1TestHelpersIssuingAuthorizationService) Expire(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationExpireParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finalize the amount on an Authorization prior to capture, when the initial authorization was for an estimated amount.
func (c v1TestHelpersIssuingAuthorizationService) FinalizeAmount(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationFinalizeAmountParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Increment a test-mode Authorization.
func (c v1TestHelpersIssuingAuthorizationService) Increment(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationIncrementParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Respond to a fraud challenge on a testmode Issuing authorization, simulating either a confirmation of fraud or a correction of legitimacy.
func (c v1TestHelpersIssuingAuthorizationService) Respond(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationRespondParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reverse a test-mode Authorization.
func (c v1TestHelpersIssuingAuthorizationService) Reverse(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationReverseParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
