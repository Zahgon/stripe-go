//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1IssuingAuthorizationService is used to invoke /v1/issuing/authorizations APIs.
type v1IssuingAuthorizationService struct {
	B   Backend
	Key string
}

// Retrieves an Issuing Authorization object.
func (c v1IssuingAuthorizationService) Retrieve(ctx context.Context, id string, params *IssuingAuthorizationRetrieveParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Authorization object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1IssuingAuthorizationService) Update(ctx context.Context, id string, params *IssuingAuthorizationUpdateParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: [Deprecated] Approves a pending Issuing Authorization object. This request should be made within the timeout window of the [real-time authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations) flow.
// This method is deprecated. Instead, [respond directly to the webhook request to approve an authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations#authorization-handling).
func (c v1IssuingAuthorizationService) Approve(ctx context.Context, id string, params *IssuingAuthorizationApproveParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: [Deprecated] Declines a pending Issuing Authorization object. This request should be made within the timeout window of the [real time authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations) flow.
// This method is deprecated. Instead, [respond directly to the webhook request to decline an authorization](https://docs.stripe.com/docs/issuing/controls/real-time-authorizations#authorization-handling).
func (c v1IssuingAuthorizationService) Decline(ctx context.Context, id string, params *IssuingAuthorizationDeclineParams) (*IssuingAuthorization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Issuing Authorization objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingAuthorizationService) List(ctx context.Context, listParams *IssuingAuthorizationListParams) *V1List[*IssuingAuthorization] {
	_ = "STUB: not implemented"
	return nil
}
