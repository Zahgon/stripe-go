//
//
// File generated from our OpenAPI spec
//
//

// Package calculation provides the /v1/tax/calculations APIs
package calculation

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/tax/calculations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Calculates tax based on the input and returns a Tax Calculation object.
func New(params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	_ = "STUB: not implemented"
	return nil,

		// Calculates tax based on the input and returns a Tax Calculation object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Tax Calculation object, if the calculation hasn't expired.
func Get(id string, params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Tax Calculation object, if the calculation hasn't expired.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TaxCalculationParams) (*stripe.TaxCalculation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the line items of a tax calculation as a collection, if the calculation hasn't expired.
func ListLineItems(params *stripe.TaxCalculationListLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// Retrieves the line items of a tax calculation as a collection, if the calculation hasn't expired.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLineItems(listParams *stripe.TaxCalculationListLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// LineItemIter is an iterator for tax calculation line items.
type LineItemIter struct {
	*stripe.Iter
}

// TaxCalculationLineItem returns the tax calculation line item which the iterator is currently pointing to.
func (i *LineItemIter) TaxCalculationLineItem() *stripe.TaxCalculationLineItem {
	_ = "STUB: not implemented"
	return nil
}

// TaxCalculationLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) TaxCalculationLineItemList() *stripe.TaxCalculationLineItemList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
