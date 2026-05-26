//
//
// File generated from our OpenAPI spec
//
//

// Package setupintent provides the /v1/setup_intents APIs
package setupintent

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/setup_intents APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a SetupIntent object.
//
// After you create the SetupIntent, attach a payment method and [confirm](https://docs.stripe.com/docs/api/setup_intents/confirm)
// it to collect any required permissions to charge the payment method later.
func New(params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a SetupIntent object.
		//
		// After you create the SetupIntent, attach a payment method and [confirm](https://docs.stripe.com/docs/api/setup_intents/confirm)
		// it to collect any required permissions to charge the payment method later.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a SetupIntent that has previously been created.
//
// Client-side retrieval using a publishable key is allowed when the client_secret is provided in the query string.
//
// When retrieved with a publishable key, only a subset of properties will be returned. Please refer to the [SetupIntent](https://docs.stripe.com/api#setup_intent_object) object reference for more details.
func Get(id string, params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of a SetupIntent that has previously been created.
	//
	// Client-side retrieval using a publishable key is allowed when the client_secret is provided in the query string.
	//
	// When retrieved with a publishable key, only a subset of properties will be returned. Please refer to the [SetupIntent](https://docs.stripe.com/api#setup_intent_object) object reference for more details.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a SetupIntent object.
func Update(id string, params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a SetupIntent object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.SetupIntentParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can cancel a SetupIntent object when it's in one of these statuses: requires_payment_method, requires_confirmation, or requires_action.
//
// After you cancel it, setup is abandoned and any operations on the SetupIntent fail with an error. You can't cancel the SetupIntent for a Checkout Session. [Expire the Checkout Session](https://docs.stripe.com/docs/api/checkout/sessions/expire) instead.
func Cancel(id string, params *stripe.SetupIntentCancelParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can cancel a SetupIntent object when it's in one of these statuses: requires_payment_method, requires_confirmation, or requires_action.
//
// After you cancel it, setup is abandoned and any operations on the SetupIntent fail with an error. You can't cancel the SetupIntent for a Checkout Session. [Expire the Checkout Session](https://docs.stripe.com/docs/api/checkout/sessions/expire) instead.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.SetupIntentCancelParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Confirm that your customer intends to set up the current or
// provided payment method. For example, you would confirm a SetupIntent
// when a customer hits the “Save” button on a payment method management
// page on your website.
//
// If the selected payment method does not require any additional
// steps from the customer, the SetupIntent will transition to the
// succeeded status.
//
// Otherwise, it will transition to the requires_action status and
// suggest additional actions via next_action. If setup fails,
// the SetupIntent will transition to the
// requires_payment_method status or the canceled status if the
// confirmation limit is reached.
func Confirm(id string, params *stripe.SetupIntentConfirmParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Confirm that your customer intends to set up the current or
// provided payment method. For example, you would confirm a SetupIntent
// when a customer hits the “Save” button on a payment method management
// page on your website.
//
// If the selected payment method does not require any additional
// steps from the customer, the SetupIntent will transition to the
// succeeded status.
//
// Otherwise, it will transition to the requires_action status and
// suggest additional actions via next_action. If setup fails,
// the SetupIntent will transition to the
// requires_payment_method status or the canceled status if the
// confirmation limit is reached.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Confirm(id string, params *stripe.SetupIntentConfirmParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verifies microdeposits on a SetupIntent object.
func VerifyMicrodeposits(id string, params *stripe.SetupIntentVerifyMicrodepositsParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verifies microdeposits on a SetupIntent object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) VerifyMicrodeposits(id string, params *stripe.SetupIntentVerifyMicrodepositsParams) (*stripe.SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of SetupIntents.
func List(params *stripe.SetupIntentListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of SetupIntents.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.SetupIntentListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for setup intents.
type Iter struct {
	*stripe.Iter
}

// SetupIntent returns the setup intent which the iterator is currently pointing to.
func (i *Iter) SetupIntent() *stripe.SetupIntent { _ = "STUB: not implemented"; return nil }

// SetupIntentList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SetupIntentList() *stripe.SetupIntentList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
