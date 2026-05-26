//
//
// File generated from our OpenAPI spec
//
//

// Package invoice provides the /v1/invoices APIs
package invoice

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/invoices APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// This endpoint creates a draft invoice for a given customer. The invoice remains a draft until you [finalize the invoice, which allows you to [pay](/api/invoices/pay) or <a href="/api/invoices/send">send](https://docs.stripe.com/api/invoices/finalize) the invoice to your customers.
func New(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil,

		// This endpoint creates a draft invoice for a given customer. The invoice remains a draft until you [finalize the invoice, which allows you to [pay](/api/invoices/pay) or <a href="/api/invoices/send">send](https://docs.stripe.com/api/invoices/finalize) the invoice to your customers.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the invoice with the given ID.
func Get(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the invoice with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Draft invoices are fully editable. Once an invoice is [finalized](https://docs.stripe.com/docs/billing/invoices/workflow#finalized),
// monetary values, as well as collection_method, become uneditable.
//
// If you would like to stop the Stripe Billing engine from automatically finalizing, reattempting payments on,
// sending reminders for, or [automatically reconciling](https://docs.stripe.com/docs/billing/invoices/reconciliation) invoices, pass
// auto_advance=false.
func Update(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Draft invoices are fully editable. Once an invoice is [finalized](https://docs.stripe.com/docs/billing/invoices/workflow#finalized),
// monetary values, as well as collection_method, become uneditable.
//
// If you would like to stop the Stripe Billing engine from automatically finalizing, reattempting payments on,
// sending reminders for, or [automatically reconciling](https://docs.stripe.com/docs/billing/invoices/reconciliation) invoices, pass
// auto_advance=false.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Permanently deletes a one-off invoice draft. This cannot be undone. Attempts to delete invoices that are no longer in a draft state will fail; once an invoice has been finalized or if an invoice is for a subscription, it must be [voided](https://docs.stripe.com/api/invoices/void).
func Del(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Permanently deletes a one-off invoice draft. This cannot be undone. Attempts to delete invoices that are no longer in a draft state will fail; once an invoice has been finalized or if an invoice is for a subscription, it must be [voided](https://docs.stripe.com/api/invoices/void).
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.InvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Adds multiple line items to an invoice. This is only possible when an invoice is still a draft.
func AddLines(id string, params *stripe.InvoiceAddLinesParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Adds multiple line items to an invoice. This is only possible when an invoice is still a draft.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) AddLines(id string, params *stripe.InvoiceAddLinesParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attaches a PaymentIntent or an Out of Band Payment to the invoice, adding it to the list of payments.
//
// For the PaymentIntent, when the PaymentIntent's status changes to succeeded, the payment is credited
// to the invoice, increasing its amount_paid. When the invoice is fully paid, the
// invoice's status becomes paid.
//
// If the PaymentIntent's status is already succeeded when it's attached, it's
// credited to the invoice immediately.
//
// See: [Partial payments](https://docs.stripe.com/docs/invoicing/partial-payments) to learn more.
func AttachPayment(id string, params *stripe.InvoiceAttachPaymentParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attaches a PaymentIntent or an Out of Band Payment to the invoice, adding it to the list of payments.
//
// For the PaymentIntent, when the PaymentIntent's status changes to succeeded, the payment is credited
// to the invoice, increasing its amount_paid. When the invoice is fully paid, the
// invoice's status becomes paid.
//
// If the PaymentIntent's status is already succeeded when it's attached, it's
// credited to the invoice immediately.
//
// See: [Partial payments](https://docs.stripe.com/docs/invoicing/partial-payments) to learn more.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) AttachPayment(id string, params *stripe.InvoiceAttachPaymentParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// At any time, you can preview the upcoming invoice for a subscription or subscription schedule. This will show you all the charges that are pending, including subscription renewal charges, invoice item charges, etc. It will also show you any discounts that are applicable to the invoice.
//
// You can also preview the effects of creating or updating a subscription or subscription schedule, including a preview of any prorations that will take place. To ensure that the actual proration is calculated exactly the same as the previewed proration, you should pass the subscription_details.proration_date parameter when doing the actual subscription update.
//
// The recommended way to get only the prorations being previewed on the invoice is to consider line items where parent.subscription_item_details.proration is true.
//
// Note that when you are viewing an upcoming invoice, you are simply viewing a preview – the invoice has not yet been created. As such, the upcoming invoice will not show up in invoice listing calls, and you cannot use the API to pay or edit the invoice. If you want to change the amount that your customer will be billed, you can add, remove, or update pending invoice items, or update the customer's discount.
//
// Note: Currency conversion calculations use the latest exchange rates. Exchange rates may vary between the time of the preview and the time of the actual invoice creation. [Learn more](https://docs.stripe.com/currencies/conversions)
func CreatePreview(params *stripe.InvoiceCreatePreviewParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// At any time, you can preview the upcoming invoice for a subscription or subscription schedule. This will show you all the charges that are pending, including subscription renewal charges, invoice item charges, etc. It will also show you any discounts that are applicable to the invoice.
//
// You can also preview the effects of creating or updating a subscription or subscription schedule, including a preview of any prorations that will take place. To ensure that the actual proration is calculated exactly the same as the previewed proration, you should pass the subscription_details.proration_date parameter when doing the actual subscription update.
//
// The recommended way to get only the prorations being previewed on the invoice is to consider line items where parent.subscription_item_details.proration is true.
//
// Note that when you are viewing an upcoming invoice, you are simply viewing a preview – the invoice has not yet been created. As such, the upcoming invoice will not show up in invoice listing calls, and you cannot use the API to pay or edit the invoice. If you want to change the amount that your customer will be billed, you can add, remove, or update pending invoice items, or update the customer's discount.
//
// Note: Currency conversion calculations use the latest exchange rates. Exchange rates may vary between the time of the preview and the time of the actual invoice creation. [Learn more](https://docs.stripe.com/currencies/conversions)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreatePreview(params *stripe.InvoiceCreatePreviewParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stripe automatically finalizes drafts before sending and attempting payment on invoices. However, if you'd like to finalize a draft invoice manually, you can do so using this method.
func FinalizeInvoice(id string, params *stripe.InvoiceFinalizeInvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stripe automatically finalizes drafts before sending and attempting payment on invoices. However, if you'd like to finalize a draft invoice manually, you can do so using this method.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) FinalizeInvoice(id string, params *stripe.InvoiceFinalizeInvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marking an invoice as uncollectible is useful for keeping track of bad debts that can be written off for accounting purposes.
func MarkUncollectible(id string, params *stripe.InvoiceMarkUncollectibleParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marking an invoice as uncollectible is useful for keeping track of bad debts that can be written off for accounting purposes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) MarkUncollectible(id string, params *stripe.InvoiceMarkUncollectibleParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.
func Pay(id string, params *stripe.InvoicePayParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Pay(id string, params *stripe.InvoicePayParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Removes multiple line items from an invoice. This is only possible when an invoice is still a draft.
func RemoveLines(id string, params *stripe.InvoiceRemoveLinesParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Removes multiple line items from an invoice. This is only possible when an invoice is still a draft.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RemoveLines(id string, params *stripe.InvoiceRemoveLinesParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stripe will automatically send invoices to customers according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to manually send an invoice to your customer out of the normal schedule, you can do so. When sending invoices that have already been paid, there will be no reference to the payment in the email.
//
// Requests made in test-mode result in no emails being sent, despite sending an invoice.sent event.
func SendInvoice(id string, params *stripe.InvoiceSendInvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stripe will automatically send invoices to customers according to your [subscriptions settings](https://dashboard.stripe.com/account/billing/automatic). However, if you'd like to manually send an invoice to your customer out of the normal schedule, you can do so. When sending invoices that have already been paid, there will be no reference to the payment in the email.
//
// Requests made in test-mode result in no emails being sent, despite sending an invoice.sent event.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SendInvoice(id string, params *stripe.InvoiceSendInvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates multiple line items on an invoice. This is only possible when an invoice is still a draft.
func UpdateLines(id string, params *stripe.InvoiceUpdateLinesParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates multiple line items on an invoice. This is only possible when an invoice is still a draft.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) UpdateLines(id string, params *stripe.InvoiceUpdateLinesParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mark a finalized invoice as void. This cannot be undone. Voiding an invoice is similar to [deletion](https://docs.stripe.com/api/invoices/delete), however it only applies to finalized invoices and maintains a papertrail where the invoice can still be found.
//
// Consult with local regulations to determine whether and how an invoice might be amended, canceled, or voided in the jurisdiction you're doing business in. You might need to [issue another invoice or <a href="/api/credit_notes/create">credit note](https://docs.stripe.com/api/invoices/create) instead. Stripe recommends that you consult with your legal counsel for advice specific to your business.
func VoidInvoice(id string, params *stripe.InvoiceVoidInvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mark a finalized invoice as void. This cannot be undone. Voiding an invoice is similar to [deletion](https://docs.stripe.com/api/invoices/delete), however it only applies to finalized invoices and maintains a papertrail where the invoice can still be found.
//
// Consult with local regulations to determine whether and how an invoice might be amended, canceled, or voided in the jurisdiction you're doing business in. You might need to [issue another invoice or <a href="/api/credit_notes/create">credit note](https://docs.stripe.com/api/invoices/create) instead. Stripe recommends that you consult with your legal counsel for advice specific to your business.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) VoidInvoice(id string, params *stripe.InvoiceVoidInvoiceParams) (*stripe.Invoice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can list all invoices, or list the invoices for a specific customer. The invoices are returned sorted by creation date, with the most recently created invoices appearing first.
func List(params *stripe.InvoiceListParams) *Iter { _ = "STUB: not implemented"; return nil }

// You can list all invoices, or list the invoices for a specific customer. The invoices are returned sorted by creation date, with the most recently created invoices appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.InvoiceListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for invoices.
type Iter struct {
	*stripe.Iter
}

// Invoice returns the invoice which the iterator is currently pointing to.
func (i *Iter) Invoice() *stripe.Invoice { _ = "STUB: not implemented"; return nil }

// InvoiceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) InvoiceList() *stripe.InvoiceList { _ = "STUB: not implemented"; return nil }

// When retrieving an invoice, you'll get a lines property containing the total count of line items and the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func ListLines(params *stripe.InvoiceListLinesParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving an invoice, you'll get a lines property containing the total count of line items and the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLines(listParams *stripe.InvoiceListLinesParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// LineItemIter is an iterator for invoice line items.
type LineItemIter struct {
	*stripe.Iter
}

// InvoiceLineItem returns the invoice line item which the iterator is currently pointing to.
func (i *LineItemIter) InvoiceLineItem() *stripe.InvoiceLineItem {
	_ = "STUB: not implemented"
	return nil
}

// InvoiceLineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) InvoiceLineItemList() *stripe.InvoiceLineItemList {
	_ = "STUB: not implemented"
	return nil
}

// Search for invoices you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func Search(params *stripe.InvoiceSearchParams) *SearchIter { _ = "STUB: not implemented"; return nil }

// Search for invoices you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Search(params *stripe.InvoiceSearchParams) *SearchIter {
	_ = "STUB: not implemented"
	return nil
}

// SearchIter is an iterator for invoices.
type SearchIter struct {
	*stripe.SearchIter
}

// Invoice returns the invoice which the iterator is currently pointing to.
func (i *SearchIter) Invoice() *stripe.Invoice { _ = "STUB: not implemented"; return nil }

// InvoiceSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) InvoiceSearchResult() *stripe.InvoiceSearchResult {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
