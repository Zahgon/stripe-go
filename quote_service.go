//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1QuoteService is used to invoke /v1/quotes APIs.
type v1QuoteService struct {
	B        Backend
	BUploads Backend
	Key      string
}

// A quote models prices and services for a customer. Default options for header, description, footer, and expires_at can be set in the dashboard via the [quote template](https://dashboard.stripe.com/settings/billing/quote).
func (c v1QuoteService) Create(ctx context.Context, params *QuoteCreateParams) (*Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the quote with the given ID.
func (c v1QuoteService) Retrieve(ctx context.Context, id string, params *QuoteRetrieveParams) (*Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A quote models prices and services for a customer.
func (c v1QuoteService) Update(ctx context.Context, id string, params *QuoteUpdateParams) (*Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Accepts the specified quote.
func (c v1QuoteService) Accept(ctx context.Context, id string, params *QuoteAcceptParams) (*Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels the quote.
func (c v1QuoteService) Cancel(ctx context.Context, id string, params *QuoteCancelParams) (*Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finalizes the quote.
func (c v1QuoteService) FinalizeQuote(ctx context.Context, id string, params *QuoteFinalizeQuoteParams) (*Quote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Download the PDF for a finalized quote. Explanation for special handling can be found [here](https://docs.stripe.com/quotes/overview#quote_pdf)
func (c v1QuoteService) PDF(ctx context.Context, id string, params *QuotePDFParams) (*APIStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your quotes.
func (c v1QuoteService) List(ctx context.Context, listParams *QuoteListParams) *V1List[*Quote] {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving a quote, there is an includable [computed.upfront.line_items](https://stripe.com/docs/api/quotes/object#quote_object-computed-upfront-line_items) property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of upfront line items.
func (c v1QuoteService) ListComputedUpfrontLineItems(ctx context.Context, listParams *QuoteListComputedUpfrontLineItemsParams) *V1List[*LineItem] {
	_ = "STUB: not implemented"
	return nil
}

// When retrieving a quote, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
func (c v1QuoteService) ListLineItems(ctx context.Context, listParams *QuoteListLineItemsParams) *V1List[*LineItem] {
	_ = "STUB: not implemented"
	return nil
}
