//
//
// File generated from our OpenAPI spec
//
//

// Package paymentmethodconfiguration provides the /v1/payment_method_configurations APIs
package paymentmethodconfiguration

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/payment_method_configurations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a payment method configuration
func New(params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a payment method configuration
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve payment method configuration
func Get(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieve payment method configuration
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update payment method configuration
func Update(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update payment method configuration
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List payment method configurations
func List(params *stripe.PaymentMethodConfigurationListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// List payment method configurations
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.PaymentMethodConfigurationListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for payment method configurations.
type Iter struct {
	*stripe.Iter
}

// PaymentMethodConfiguration returns the payment method configuration which the iterator is currently pointing to.
func (i *Iter) PaymentMethodConfiguration() *stripe.PaymentMethodConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// PaymentMethodConfigurationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentMethodConfigurationList() *stripe.PaymentMethodConfigurationList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
