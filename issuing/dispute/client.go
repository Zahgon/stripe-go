//
//
// File generated from our OpenAPI spec
//
//

// Package dispute provides the /v1/issuing/disputes APIs
package dispute

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/disputes APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an Issuing Dispute object. Individual pieces of evidence within the evidence object are optional at this point. Stripe only validates that required evidence is present during submission. Refer to [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence) for more details about evidence requirements.
func New(params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates an Issuing Dispute object. Individual pieces of evidence within the evidence object are optional at this point. Stripe only validates that required evidence is present during submission. Refer to [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence) for more details about evidence requirements.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves an Issuing Dispute object.
func Get(id string, params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves an Issuing Dispute object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Dispute object by setting the values of the parameters passed. Any parameters not provided will be left unchanged. Properties on the evidence object can be unset by passing in an empty string.
func Update(id string, params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified Issuing Dispute object by setting the values of the parameters passed. Any parameters not provided will be left unchanged. Properties on the evidence object can be unset by passing in an empty string.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.IssuingDisputeParams) (*stripe.IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Submits an Issuing Dispute to the card network. Stripe validates that all evidence fields required for the dispute's reason are present. For more details, see [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence).
func Submit(id string, params *stripe.IssuingDisputeSubmitParams) (*stripe.IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Submits an Issuing Dispute to the card network. Stripe validates that all evidence fields required for the dispute's reason are present. For more details, see [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Submit(id string, params *stripe.IssuingDisputeSubmitParams) (*stripe.IssuingDispute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Issuing Dispute objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingDisputeListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of Issuing Dispute objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IssuingDisputeListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for issuing disputes.
type Iter struct {
	*stripe.Iter
}

// IssuingDispute returns the issuing dispute which the iterator is currently pointing to.
func (i *Iter) IssuingDispute() *stripe.IssuingDispute { _ = "STUB: not implemented"; return nil }

// IssuingDisputeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingDisputeList() *stripe.IssuingDisputeList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
