//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TransferService is used to invoke /v1/transfers APIs.
type v1TransferService struct {
	B   Backend
	Key string
}

// To send funds from your Stripe account to a connected account, you create a new transfer object. Your [Stripe balance](https://docs.stripe.com/api#balance) must be able to cover the transfer amount, or you'll receive an “Insufficient Funds” error.
func (c v1TransferService) Create(ctx context.Context, params *TransferCreateParams) (*Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing transfer. Supply the unique transfer ID from either a transfer creation request or the transfer list, and Stripe will return the corresponding transfer information.
func (c v1TransferService) Retrieve(ctx context.Context, id string, params *TransferRetrieveParams) (*Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified transfer by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request accepts only metadata as an argument.
func (c v1TransferService) Update(ctx context.Context, id string, params *TransferUpdateParams) (*Transfer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of existing transfers sent to connected accounts. The transfers are returned in sorted order, with the most recently created transfers appearing first.
func (c v1TransferService) List(ctx context.Context, listParams *TransferListParams) *V1List[*Transfer] {
	_ = "STUB: not implemented"
	return nil
}
