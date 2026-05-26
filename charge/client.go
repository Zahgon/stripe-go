//
//
// File generated from our OpenAPI spec
//
//

// Package charge provides the /v1/charges APIs
package charge

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/charges APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// This method is no longer recommended—use the [Payment Intents API](https://docs.stripe.com/docs/api/payment_intents)
// to initiate a new payment instead. Confirmation of the PaymentIntent creates the Charge
// object used to request payment.
func New(params *stripe.ChargeParams) (*stripe.Charge, error) {
	_ = "STUB: not implemented"
	return nil,

		// This method is no longer recommended—use the [Payment Intents API](https://docs.stripe.com/docs/api/payment_intents)
		// to initiate a new payment instead. Confirmation of the PaymentIntent creates the Charge
		// object used to request payment.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.ChargeParams) (*stripe.Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a charge that has previously been created. Supply the unique charge ID that was returned from your previous request, and Stripe will return the corresponding charge information. The same information is returned when creating or refunding the charge.
func Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of a charge that has previously been created. Supply the unique charge ID that was returned from your previous request, and Stripe will return the corresponding charge information. The same information is returned when creating or refunding the charge.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified charge by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified charge by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture the payment of an existing, uncaptured charge that was created with the capture option set to false.
//
// Uncaptured payments expire a set number of days after they are created ([7 by default](https://docs.stripe.com/docs/charges/placing-a-hold)), after which they are marked as refunded and capture attempts will fail.
//
// Don't use this method to capture a PaymentIntent-initiated charge. Use [Capture a PaymentIntent](https://docs.stripe.com/docs/api/payment_intents/capture).
func Capture(id string, params *stripe.ChargeCaptureParams) (*stripe.Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture the payment of an existing, uncaptured charge that was created with the capture option set to false.
//
// Uncaptured payments expire a set number of days after they are created ([7 by default](https://docs.stripe.com/docs/charges/placing-a-hold)), after which they are marked as refunded and capture attempts will fail.
//
// Don't use this method to capture a PaymentIntent-initiated charge. Use [Capture a PaymentIntent](https://docs.stripe.com/docs/api/payment_intents/capture).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Capture(id string, params *stripe.ChargeCaptureParams) (*stripe.Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of charges you've previously created. The charges are returned in sorted order, with the most recent charges appearing first.
func List(params *stripe.ChargeListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of charges you've previously created. The charges are returned in sorted order, with the most recent charges appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ChargeListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for charges.
type Iter struct {
	*stripe.Iter
}

// Charge returns the charge which the iterator is currently pointing to.
func (i *Iter) Charge() *stripe.Charge { _ = "STUB: not implemented"; return nil }

// ChargeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ChargeList() *stripe.ChargeList { _ = "STUB: not implemented"; return nil }

// Search for charges you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func Search(params *stripe.ChargeSearchParams) *SearchIter { _ = "STUB: not implemented"; return nil }

// Search for charges you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.ChargeSearchParams) *SearchIter {
	_ = "STUB: not implemented"
	return nil
}

// SearchIter is an iterator for charges.
type SearchIter struct {
	*stripe.SearchIter
}

// Charge returns the charge which the iterator is currently pointing to.
func (i *SearchIter) Charge() *stripe.Charge { _ = "STUB: not implemented"; return nil }

// ChargeSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) ChargeSearchResult() *stripe.ChargeSearchResult {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
