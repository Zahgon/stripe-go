//
//
// File generated from our OpenAPI spec
//
//

// Package personalizationdesign provides the /v1/issuing/personalization_designs APIs
package personalizationdesign

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/personalization_designs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a personalization design object.
func New(params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a personalization design object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a personalization design object.
func Get(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a personalization design object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a card personalization object.
func Update(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a card personalization object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingPersonalizationDesignParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of personalization design objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingPersonalizationDesignListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of personalization design objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.IssuingPersonalizationDesignListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for issuing personalization designs.
type Iter struct {
	*stripe.Iter
}

// IssuingPersonalizationDesign returns the issuing personalization design which the iterator is currently pointing to.
func (i *Iter) IssuingPersonalizationDesign() *stripe.IssuingPersonalizationDesign {
	_ = "STUB: not implemented"
	return nil
}

// IssuingPersonalizationDesignList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingPersonalizationDesignList() *stripe.IssuingPersonalizationDesignList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
