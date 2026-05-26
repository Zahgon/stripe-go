//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TransferReversalService is used to invoke /v1/transfers/{id}/reversals APIs.
type v1TransferReversalService struct {
	B   Backend
	Key string
}

// When you create a new reversal, you must specify a transfer to create it on.
//
// When reversing transfers, you can optionally reverse part of the transfer. You can do so as many times as you wish until the entire transfer has been reversed.
//
// Once entirely reversed, a transfer can't be reversed again. This method will return an error when called on an already-reversed transfer, or when trying to reverse more money than is left on a transfer.
func (c v1TransferReversalService) Create(ctx context.Context, params *TransferReversalCreateParams) (*TransferReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// By default, you can see the 10 most recent reversals stored directly on the transfer object, but you can also retrieve details about a specific reversal stored on the transfer.
func (c v1TransferReversalService) Retrieve(ctx context.Context, id string, params *TransferReversalRetrieveParams) (*TransferReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified reversal by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// This request only accepts metadata and description as arguments.
func (c v1TransferReversalService) Update(ctx context.Context, id string, params *TransferReversalUpdateParams) (*TransferReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can see a list of the reversals belonging to a specific transfer. Note that the 10 most recent reversals are always available by default on the transfer object. If you need more than those 10, you can use this API method and the limit and starting_after parameters to page through additional reversals.
func (c v1TransferReversalService) List(ctx context.Context, listParams *TransferReversalListParams) *V1List[*TransferReversal] {
	_ = "STUB: not implemented"
	return nil
}
