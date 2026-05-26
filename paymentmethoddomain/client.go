//
//
// File generated from our OpenAPI spec
//
//

// Package paymentmethoddomain provides the /v1/payment_method_domains APIs
package paymentmethoddomain

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/payment_method_domains APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a payment method domain.
func New(params *stripe.PaymentMethodDomainParams) (*stripe.PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a payment method domain.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.PaymentMethodDomainParams) (*stripe.PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing payment method domain.
func Get(id string, params *stripe.PaymentMethodDomainParams) (*stripe.PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an existing payment method domain.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.PaymentMethodDomainParams) (*stripe.PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing payment method domain.
func Update(id string, params *stripe.PaymentMethodDomainParams) (*stripe.PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing payment method domain.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentMethodDomainParams) (*stripe.PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Some payment methods might require additional steps to register a domain. If the requirements weren't satisfied when the domain was created, the payment method will be inactive on the domain.
// The payment method doesn't appear in Elements or Embedded Checkout for this domain until it is active.
//
// To activate a payment method on an existing payment method domain, complete the required registration steps specific to the payment method, and then validate the payment method domain with this endpoint.
//
// Related guides: [Payment method domains](https://docs.stripe.com/docs/payments/payment-methods/pmd-registration).
func Validate(id string, params *stripe.PaymentMethodDomainValidateParams) (*stripe.PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Some payment methods might require additional steps to register a domain. If the requirements weren't satisfied when the domain was created, the payment method will be inactive on the domain.
// The payment method doesn't appear in Elements or Embedded Checkout for this domain until it is active.
//
// To activate a payment method on an existing payment method domain, complete the required registration steps specific to the payment method, and then validate the payment method domain with this endpoint.
//
// Related guides: [Payment method domains](https://docs.stripe.com/docs/payments/payment-methods/pmd-registration).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Validate(id string, params *stripe.PaymentMethodDomainValidateParams) (*stripe.PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists the details of existing payment method domains.
func List(params *stripe.PaymentMethodDomainListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Lists the details of existing payment method domains.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.PaymentMethodDomainListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for payment method domains.
type Iter struct {
	*stripe.Iter
}

// PaymentMethodDomain returns the payment method domain which the iterator is currently pointing to.
func (i *Iter) PaymentMethodDomain() *stripe.PaymentMethodDomain {
	_ = "STUB: not implemented"
	return nil
}

// PaymentMethodDomainList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentMethodDomainList() *stripe.PaymentMethodDomainList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
