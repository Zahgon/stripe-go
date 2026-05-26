//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ApplePayDomainService is used to invoke /v1/apple_pay/domains APIs.
type v1ApplePayDomainService struct {
	B   Backend
	Key string
}

// Create an apple pay domain.
func (c v1ApplePayDomainService) Create(ctx context.Context, params *ApplePayDomainCreateParams) (*ApplePayDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve an apple pay domain.
func (c v1ApplePayDomainService) Retrieve(ctx context.Context, id string, params *ApplePayDomainRetrieveParams) (*ApplePayDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete an apple pay domain.
func (c v1ApplePayDomainService) Delete(ctx context.Context, id string, params *ApplePayDomainDeleteParams) (*ApplePayDomain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List apple pay domains.
func (c v1ApplePayDomainService) List(ctx context.Context, listParams *ApplePayDomainListParams) *V1List[*ApplePayDomain] {
	_ = "STUB: not implemented"
	return nil
}
