//
//
// File generated from our OpenAPI spec
//
//

// Package applicationfee provides the /v1/application_fees APIs
package applicationfee

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/application_fees APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.
func Get(id string, params *stripe.ApplicationFeeParams) (*stripe.ApplicationFee, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ApplicationFeeParams) (*stripe.ApplicationFee, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of application fees you've previously collected. The application fees are returned in sorted order, with the most recent fees appearing first.
func List(params *stripe.ApplicationFeeListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of application fees you've previously collected. The application fees are returned in sorted order, with the most recent fees appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ApplicationFeeListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for application fees.
type Iter struct {
	*stripe.Iter
}

// ApplicationFee returns the application fee which the iterator is currently pointing to.
func (i *Iter) ApplicationFee() *stripe.ApplicationFee { _ = "STUB: not implemented"; return nil }

// ApplicationFeeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ApplicationFeeList() *stripe.ApplicationFeeList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
