//
//
// File generated from our OpenAPI spec
//
//

// Package customer provides the /v1/customers APIs
package customer

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/customers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new customer object.
func New(params *stripe.CustomerParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new customer object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.CustomerParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Customer object.
func Get(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Customer object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified customer by setting the values of the parameters passed. Any parameters not provided are left unchanged. For example, if you pass the source parameter, that becomes the customer's active source (such as a card) to be used for all charges in the future. When you update a customer to a new valid card source by passing the source parameter: for each of the customer's current subscriptions, if the subscription bills automatically and is in the past_due state, then the latest open invoice for the subscription with automatic collection enabled is retried. This retry doesn't count as an automatic retry, and doesn't affect the next regularly scheduled payment for the invoice. Changing the default_source for a customer doesn't trigger this behavior.
//
// This request accepts mostly the same arguments as the customer creation call.
func Update(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified customer by setting the values of the parameters passed. Any parameters not provided are left unchanged. For example, if you pass the source parameter, that becomes the customer's active source (such as a card) to be used for all charges in the future. When you update a customer to a new valid card source by passing the source parameter: for each of the customer's current subscriptions, if the subscription bills automatically and is in the past_due state, then the latest open invoice for the subscription with automatic collection enabled is retried. This retry doesn't count as an automatic retry, and doesn't affect the next regularly scheduled payment for the invoice. Changing the default_source for a customer doesn't trigger this behavior.
//
// This request accepts mostly the same arguments as the customer creation call.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Permanently deletes a customer. It cannot be undone. Also immediately cancels any active subscriptions on the customer.
func Del(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Permanently deletes a customer. It cannot be undone. Also immediately cancels any active subscriptions on the customer.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.CustomerParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve funding instructions for a customer cash balance. If funding instructions do not yet exist for the customer, new
// funding instructions will be created. If funding instructions have already been created for a given customer, the same
// funding instructions will be retrieved. In other words, we will return the same funding instructions each time.
func CreateFundingInstructions(id string, params *stripe.CustomerCreateFundingInstructionsParams) (*stripe.FundingInstructions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieve funding instructions for a customer cash balance. If funding instructions do not yet exist for the customer, new
// funding instructions will be created. If funding instructions have already been created for a given customer, the same
// funding instructions will be retrieved. In other words, we will return the same funding instructions each time.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreateFundingInstructions(id string, params *stripe.CustomerCreateFundingInstructionsParams) (*stripe.FundingInstructions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Removes the currently applied discount on a customer.
func DeleteDiscount(id string, params *stripe.CustomerDeleteDiscountParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Removes the currently applied discount on a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) DeleteDiscount(id string, params *stripe.CustomerDeleteDiscountParams) (*stripe.Customer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a PaymentMethod object for a given Customer.
func RetrievePaymentMethod(id string, params *stripe.CustomerRetrievePaymentMethodParams) (*stripe.PaymentMethod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a PaymentMethod object for a given Customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RetrievePaymentMethod(id string, params *stripe.CustomerRetrievePaymentMethodParams) (*stripe.PaymentMethod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your customers. The customers are returned sorted by creation date, with the most recent customers appearing first.
func List(params *stripe.CustomerListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your customers. The customers are returned sorted by creation date, with the most recent customers appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CustomerListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for customers.
type Iter struct {
	*stripe.Iter
}

// Customer returns the customer which the iterator is currently pointing to.
func (i *Iter) Customer() *stripe.Customer { _ = "STUB: not implemented"; return nil }

// CustomerList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CustomerList() *stripe.CustomerList { _ = "STUB: not implemented"; return nil }

// Returns a list of PaymentMethods for a given Customer
func ListPaymentMethods(params *stripe.CustomerListPaymentMethodsParams) *PaymentMethodIter {
	_ = "STUB: not implemented"
	return nil
}

// Returns a list of PaymentMethods for a given Customer
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListPaymentMethods(listParams *stripe.CustomerListPaymentMethodsParams) *PaymentMethodIter {
	_ = "STUB: not implemented"
	return nil
}

// PaymentMethodIter is an iterator for payment methods.
type PaymentMethodIter struct {
	*stripe.Iter
}

// PaymentMethod returns the payment method which the iterator is currently pointing to.
func (i *PaymentMethodIter) PaymentMethod() *stripe.PaymentMethod {
	_ = "STUB: not implemented"
	return nil
}

// PaymentMethodList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *PaymentMethodIter) PaymentMethodList() *stripe.PaymentMethodList {
	_ = "STUB: not implemented"
	return nil
}

// Search for customers you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func Search(params *stripe.CustomerSearchParams) *SearchIter { _ = "STUB: not implemented"; return nil }

// Search for customers you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.CustomerSearchParams) *SearchIter {
	_ = "STUB: not implemented"
	return nil
}

// SearchIter is an iterator for customers.
type SearchIter struct {
	*stripe.SearchIter
}

// Customer returns the customer which the iterator is currently pointing to.
func (i *SearchIter) Customer() *stripe.Customer { _ = "STUB: not implemented"; return nil }

// CustomerSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) CustomerSearchResult() *stripe.CustomerSearchResult {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
