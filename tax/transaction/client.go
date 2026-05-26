//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the /v1/tax/transactions APIs
package transaction

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/tax/transactions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a Tax Transaction object.
func Get(id string, params *stripe.TaxTransactionParams) (*stripe.TaxTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Tax Transaction object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TaxTransactionParams) (*stripe.TaxTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Creates a Tax Transaction from a calculation, if that calculation hasn't expired. Calculations expire after 90 days.
func CreateFromCalculation(params *stripe.TaxTransactionCreateFromCalculationParams) (*stripe.TaxTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Creates a Tax Transaction from a calculation, if that calculation hasn't expired. Calculations expire after 90 days.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreateFromCalculation(params *stripe.TaxTransactionCreateFromCalculationParams) (*stripe.TaxTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Partially or fully reverses a previously created Transaction.
func CreateReversal(params *stripe.TaxTransactionCreateReversalParams) (*stripe.TaxTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Partially or fully reverses a previously created Transaction.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreateReversal(params *stripe.TaxTransactionCreateReversalParams) (*stripe.TaxTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the line items of a committed standalone transaction as a collection.
func ListLineItems(params *stripe.TaxTransactionListLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// Retrieves the line items of a committed standalone transaction as a collection.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLineItems(listParams *stripe.TaxTransactionListLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// LineItemIter is an iterator for tax transaction line items.
type LineItemIter struct {
	*stripe.Iter
}

// TaxTransactionLineItem returns the tax transaction line item which the iterator is currently pointing to.
func (i *LineItemIter) TaxTransactionLineItem() *stripe.TaxTransactionLineItem {
	_ = "STUB: not implemented"
	return nil
}

// TaxTransactionLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) TaxTransactionLineItemList() *stripe.TaxTransactionLineItemList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
