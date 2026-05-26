//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TerminalReaderService is used to invoke /v1/terminal/readers APIs.
type v1TerminalReaderService struct {
	B   Backend
	Key string
}

// Creates a new Reader object.
func (c v1TerminalReaderService) Create(ctx context.Context, params *TerminalReaderCreateParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Reader object.
func (c v1TerminalReaderService) Retrieve(ctx context.Context, id string, params *TerminalReaderRetrieveParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Reader object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1TerminalReaderService) Update(ctx context.Context, id string, params *TerminalReaderUpdateParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a Reader object.
func (c v1TerminalReaderService) Delete(ctx context.Context, id string, params *TerminalReaderDeleteParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels the current reader action. See [Programmatic Cancellation](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven#programmatic-cancellation) for more details.
func (c v1TerminalReaderService) CancelAction(ctx context.Context, id string, params *TerminalReaderCancelActionParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates an [input collection flow](https://docs.stripe.com/docs/terminal/features/collect-inputs) on a Reader to display input forms and collect information from your customers.
func (c v1TerminalReaderService) CollectInputs(ctx context.Context, id string, params *TerminalReaderCollectInputsParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation. See [Collecting a Payment method](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#collect-a-paymentmethod) for more details.
func (c v1TerminalReaderService) CollectPaymentMethod(ctx context.Context, id string, params *TerminalReaderCollectPaymentMethodParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finalizes a payment on a Reader. See [Confirming a Payment](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#confirm-the-paymentintent) for more details.
func (c v1TerminalReaderService) ConfirmPaymentIntent(ctx context.Context, id string, params *TerminalReaderConfirmPaymentIntentParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates a payment flow on a Reader. See [process the payment](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=immediately#process-payment) for more details.
func (c v1TerminalReaderService) ProcessPaymentIntent(ctx context.Context, id string, params *TerminalReaderProcessPaymentIntentParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates a SetupIntent flow on a Reader. See [Save directly without charging](https://docs.stripe.com/docs/terminal/features/saving-payment-details/save-directly) for more details.
func (c v1TerminalReaderService) ProcessSetupIntent(ctx context.Context, id string, params *TerminalReaderProcessSetupIntentParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates an in-person refund on a Reader. See [Refund an Interac Payment](https://docs.stripe.com/docs/terminal/payments/regional?integration-country=CA#refund-an-interac-payment) for more details.
func (c v1TerminalReaderService) RefundPayment(ctx context.Context, id string, params *TerminalReaderRefundPaymentParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sets the reader display to show [cart details](https://docs.stripe.com/docs/terminal/features/display).
func (c v1TerminalReaderService) SetReaderDisplay(ctx context.Context, id string, params *TerminalReaderSetReaderDisplayParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Reader objects.
func (c v1TerminalReaderService) List(ctx context.Context, listParams *TerminalReaderListParams) *V1List[*TerminalReader] {
	_ = "STUB: not implemented"
	return nil
}
