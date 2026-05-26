//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TaxTransactionService is used to invoke /v1/tax/transactions APIs.
type v1TaxTransactionService struct {
	B   Backend
	Key string
}

// Retrieves a Tax Transaction object.
func (c v1TaxTransactionService) Retrieve(ctx context.Context, id string, params *TaxTransactionRetrieveParams) (*TaxTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Creates a Tax Transaction from a calculation, if that calculation hasn't expired. Calculations expire after 90 days.
func (c v1TaxTransactionService) CreateFromCalculation(ctx context.Context, params *TaxTransactionCreateFromCalculationParams) (*TaxTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Partially or fully reverses a previously created Transaction.
func (c v1TaxTransactionService) CreateReversal(ctx context.Context, params *TaxTransactionCreateReversalParams) (*TaxTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the line items of a committed standalone transaction as a collection.
func (c v1TaxTransactionService) ListLineItems(ctx context.Context, listParams *TaxTransactionListLineItemsParams) *V1List[*TaxTransactionLineItem] {
	_ = "STUB: not implemented"
	return nil
}
