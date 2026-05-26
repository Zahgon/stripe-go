//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1IssuingDisputeService is used to invoke /v1/issuing/disputes APIs.
type v1IssuingDisputeService struct {
	B   Backend
	Key string
}

// Creates an Issuing Dispute object. Individual pieces of evidence within the evidence object are optional at this point. Stripe only validates that required evidence is present during submission. Refer to [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence) for more details about evidence requirements.
func (c v1IssuingDisputeService) Create(ctx context.Context, params *IssuingDisputeCreateParams) (*IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an Issuing Dispute object.
func (c v1IssuingDisputeService) Retrieve(ctx context.Context, id string, params *IssuingDisputeRetrieveParams) (*IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Dispute object by setting the values of the parameters passed. Any parameters not provided will be left unchanged. Properties on the evidence object can be unset by passing in an empty string.
func (c v1IssuingDisputeService) Update(ctx context.Context, id string, params *IssuingDisputeUpdateParams) (*IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Submits an Issuing Dispute to the card network. Stripe validates that all evidence fields required for the dispute's reason are present. For more details, see [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence).
func (c v1IssuingDisputeService) Submit(ctx context.Context, id string, params *IssuingDisputeSubmitParams) (*IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Issuing Dispute objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingDisputeService) List(ctx context.Context, listParams *IssuingDisputeListParams) *V1List[*IssuingDispute] {
	_ = "STUB: not implemented"
	return nil
}
