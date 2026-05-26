//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PaymentMethodDomainService is used to invoke /v1/payment_method_domains APIs.
type v1PaymentMethodDomainService struct {
	B   Backend
	Key string
}

// Creates a payment method domain.
func (c v1PaymentMethodDomainService) Create(ctx context.Context, params *PaymentMethodDomainCreateParams) (*PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing payment method domain.
func (c v1PaymentMethodDomainService) Retrieve(ctx context.Context, id string, params *PaymentMethodDomainRetrieveParams) (*PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates an existing payment method domain.
func (c v1PaymentMethodDomainService) Update(ctx context.Context, id string, params *PaymentMethodDomainUpdateParams) (*PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Some payment methods might require additional steps to register a domain. If the requirements weren't satisfied when the domain was created, the payment method will be inactive on the domain.
// The payment method doesn't appear in Elements or Embedded Checkout for this domain until it is active.
//
// To activate a payment method on an existing payment method domain, complete the required registration steps specific to the payment method, and then validate the payment method domain with this endpoint.
//
// Related guides: [Payment method domains](https://docs.stripe.com/docs/payments/payment-methods/pmd-registration).
func (c v1PaymentMethodDomainService) Validate(ctx context.Context, id string, params *PaymentMethodDomainValidateParams) (*PaymentMethodDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lists the details of existing payment method domains.
func (c v1PaymentMethodDomainService) List(ctx context.Context, listParams *PaymentMethodDomainListParams) *V1List[*PaymentMethodDomain] {
	_ = "STUB: not implemented"
	return nil
}
