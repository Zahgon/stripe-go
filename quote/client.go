//
//
// File generated from our OpenAPI spec
//
//

// Package quote provides the /v1/quotes APIs
package quote

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/quotes APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B        stripe.Backend
	BUploads stripe.Backend
	Key      string
}

// A quote models prices and services for a customer. Default options for header, description, footer, and expires_at can be set in the dashboard via the [quote template](https://dashboard.stripe.com/settings/billing/quote).
func New(params *stripe.QuoteParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil,

		// A quote models prices and services for a customer. Default options for header, description, footer, and expires_at can be set in the dashboard via the [quote template](https://dashboard.stripe.com/settings/billing/quote).
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.QuoteParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the quote with the given ID.
func Get(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the quote with the given ID.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A quote models prices and services for a customer.
func Update(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A quote models prices and services for a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.QuoteParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Accepts the specified quote.
func Accept(id string, params *stripe.QuoteAcceptParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Accepts the specified quote.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Accept(id string, params *stripe.QuoteAcceptParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels the quote.
func Cancel(id string, params *stripe.QuoteCancelParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels the quote.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.QuoteCancelParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finalizes the quote.
func FinalizeQuote(id string, params *stripe.QuoteFinalizeQuoteParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finalizes the quote.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) FinalizeQuote(id string, params *stripe.QuoteFinalizeQuoteParams) (*stripe.Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Download the PDF for a finalized quote. Explanation for special handling can be found [here](https://docs.stripe.com/quotes/overview#quote_pdf)
func PDF(id string, params *stripe.QuotePDFParams) (*stripe.APIStream, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Download the PDF for a finalized quote. Explanation for special handling can be found [here](https://docs.stripe.com/quotes/overview#quote_pdf)
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) PDF(id string, params *stripe.QuotePDFParams) (*stripe.APIStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your quotes.
func List(params *stripe.QuoteListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of your quotes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.QuoteListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for quotes.
type Iter struct {
	*stripe.Iter
}

// Quote returns the quote which the iterator is currently pointing to.
func (i *Iter) Quote() *stripe.Quote { _ = "STUB: not implemented"; return nil }

// QuoteList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) QuoteList() *stripe.QuoteList { _ = "STUB: not implemented"; return nil }

// When retrieving a quote, there is an includable [computed.upfront.line_items](https://stripe.com/docs/api/quotes/object#quote_object-computed-upfront-line_items) property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of upfront line items.
func ListComputedUpfrontLineItems(params *stripe.QuoteListComputedUpfrontLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving a quote, there is an includable [computed.upfront.line_items](https://stripe.com/docs/api/quotes/object#quote_object-computed-upfront-line_items) property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of upfront line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListComputedUpfrontLineItems(listParams *stripe.QuoteListComputedUpfrontLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// LineItemIter is an iterator for line items.
type LineItemIter struct {
	*stripe.Iter
}

// LineItem returns the line item which the iterator is currently pointing to.
func (i *LineItemIter) LineItem() *stripe.LineItem { _ = "STUB: not implemented"; return nil }

// LineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *LineItemIter) LineItemList() *stripe.LineItemList { _ = "STUB: not implemented"; return nil }

// When retrieving a quote, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func ListLineItems(params *stripe.QuoteListLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving a quote, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListLineItems(listParams *stripe.QuoteListLineItemsParams) *LineItemIter {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
