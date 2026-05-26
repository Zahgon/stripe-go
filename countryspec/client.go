//
//
// File generated from our OpenAPI spec
//
//

// Package countryspec provides the /v1/country_specs APIs
package countryspec

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/country_specs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Returns a Country Spec for a given Country code.
func Get(id string, params *stripe.CountrySpecParams) (*stripe.CountrySpec, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Returns a Country Spec for a given Country code.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.CountrySpecParams) (*stripe.CountrySpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists all Country Spec objects available in the API.
func List(params *stripe.CountrySpecListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Lists all Country Spec objects available in the API.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CountrySpecListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for country specs.
type Iter struct {
	*stripe.Iter
}

// CountrySpec returns the country spec which the iterator is currently pointing to.
func (i *Iter) CountrySpec() *stripe.CountrySpec { _ = "STUB: not implemented"; return nil }

// CountrySpecList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CountrySpecList() *stripe.CountrySpecList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
