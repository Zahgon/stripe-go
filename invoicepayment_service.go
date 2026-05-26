//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1InvoicePaymentService is used to invoke /v1/invoice_payments APIs.
type v1InvoicePaymentService struct {
	B   Backend
	Key string
}

// Retrieves the invoice payment with the given ID.
func (c v1InvoicePaymentService) Retrieve(ctx context.Context, id string, params *InvoicePaymentRetrieveParams) (*InvoicePayment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When retrieving an invoice, there is an includable payments property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of payments.
func (c v1InvoicePaymentService) List(ctx context.Context, listParams *InvoicePaymentListParams) *V1List[*InvoicePayment] {
	_ = "STUB: not implemented"
	return nil
}
