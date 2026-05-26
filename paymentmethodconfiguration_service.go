//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1PaymentMethodConfigurationService is used to invoke /v1/payment_method_configurations APIs.
type v1PaymentMethodConfigurationService struct {
	B   Backend
	Key string
}

// Creates a payment method configuration
func (c v1PaymentMethodConfigurationService) Create(ctx context.Context, params *PaymentMethodConfigurationCreateParams) (*PaymentMethodConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve payment method configuration
func (c v1PaymentMethodConfigurationService) Retrieve(ctx context.Context, id string, params *PaymentMethodConfigurationRetrieveParams) (*PaymentMethodConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update payment method configuration
func (c v1PaymentMethodConfigurationService) Update(ctx context.Context, id string, params *PaymentMethodConfigurationUpdateParams) (*PaymentMethodConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List payment method configurations
func (c v1PaymentMethodConfigurationService) List(ctx context.Context, listParams *PaymentMethodConfigurationListParams) *V1List[*PaymentMethodConfiguration] {
	_ = "STUB: not implemented"
	return nil
}
