//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1IssuingTransactionService is used to invoke /v1/issuing/transactions APIs.
type v1IssuingTransactionService struct {
	B   Backend
	Key string
}

// Retrieves an Issuing Transaction object.
func (c v1IssuingTransactionService) Retrieve(ctx context.Context, id string, params *IssuingTransactionRetrieveParams) (*IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Transaction object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1IssuingTransactionService) Update(ctx context.Context, id string, params *IssuingTransactionUpdateParams) (*IssuingTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Issuing Transaction objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingTransactionService) List(ctx context.Context, listParams *IssuingTransactionListParams) *V1List[*IssuingTransaction] {
	_ = "STUB: not implemented"
	return nil
}
