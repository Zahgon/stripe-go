//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1CustomerService is used to invoke /v1/customers APIs.
type v1CustomerService struct {
	B   Backend
	Key string
}

// Creates a new customer object.
func (c v1CustomerService) Create(ctx context.Context, params *CustomerCreateParams) (*Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Customer object.
func (c v1CustomerService) Retrieve(ctx context.Context, id string, params *CustomerRetrieveParams) (*Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified customer by setting the values of the parameters passed. Any parameters not provided are left unchanged. For example, if you pass the source parameter, that becomes the customer's active source (such as a card) to be used for all charges in the future. When you update a customer to a new valid card source by passing the source parameter: for each of the customer's current subscriptions, if the subscription bills automatically and is in the past_due state, then the latest open invoice for the subscription with automatic collection enabled is retried. This retry doesn't count as an automatic retry, and doesn't affect the next regularly scheduled payment for the invoice. Changing the default_source for a customer doesn't trigger this behavior.
//
// This request accepts mostly the same arguments as the customer creation call.
func (c v1CustomerService) Update(ctx context.Context, id string, params *CustomerUpdateParams) (*Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Permanently deletes a customer. It cannot be undone. Also immediately cancels any active subscriptions on the customer.
func (c v1CustomerService) Delete(ctx context.Context, id string, params *CustomerDeleteParams) (*Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve funding instructions for a customer cash balance. If funding instructions do not yet exist for the customer, new
// funding instructions will be created. If funding instructions have already been created for a given customer, the same
// funding instructions will be retrieved. In other words, we will return the same funding instructions each time.
func (c v1CustomerService) CreateFundingInstructions(ctx context.Context, id string, params *CustomerCreateFundingInstructionsParams) (*FundingInstructions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Removes the currently applied discount on a customer.
func (c v1CustomerService) DeleteDiscount(ctx context.Context, id string, params *CustomerDeleteDiscountParams) (*Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a PaymentMethod object for a given Customer.
func (c v1CustomerService) RetrievePaymentMethod(ctx context.Context, id string, params *CustomerRetrievePaymentMethodParams) (*PaymentMethod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your customers. The customers are returned sorted by creation date, with the most recent customers appearing first.
func (c v1CustomerService) List(ctx context.Context, listParams *CustomerListParams) *V1List[*Customer] {
	_ = "STUB: not implemented"
	return nil
}

// Returns a list of PaymentMethods for a given Customer
func (c v1CustomerService) ListPaymentMethods(ctx context.Context, listParams *CustomerListPaymentMethodsParams) *V1List[*PaymentMethod] {
	_ = "STUB: not implemented"
	return nil
}

// Search for customers you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func (c v1CustomerService) Search(ctx context.Context, params *CustomerSearchParams) *V1SearchList[*Customer] {
	_ = "STUB: not implemented"
	return nil
}
