//
//
// File generated from our OpenAPI spec
//
//

// Package reader provides the /v1/terminal/readers APIs
package reader

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/terminal/readers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new Reader object.
func New(params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new Reader object.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a Reader object.
func Get(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a Reader object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Reader object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func Update(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a Reader object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a Reader object.
func Del(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes a Reader object.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.TerminalReaderParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels the current reader action. See [Programmatic Cancellation](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven#programmatic-cancellation) for more details.
func CancelAction(id string, params *stripe.TerminalReaderCancelActionParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancels the current reader action. See [Programmatic Cancellation](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven#programmatic-cancellation) for more details.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CancelAction(id string, params *stripe.TerminalReaderCancelActionParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates an [input collection flow](https://docs.stripe.com/docs/terminal/features/collect-inputs) on a Reader to display input forms and collect information from your customers.
func CollectInputs(id string, params *stripe.TerminalReaderCollectInputsParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates an [input collection flow](https://docs.stripe.com/docs/terminal/features/collect-inputs) on a Reader to display input forms and collect information from your customers.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CollectInputs(id string, params *stripe.TerminalReaderCollectInputsParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation. See [Collecting a Payment method](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#collect-a-paymentmethod) for more details.
func CollectPaymentMethod(id string, params *stripe.TerminalReaderCollectPaymentMethodParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation. See [Collecting a Payment method](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#collect-a-paymentmethod) for more details.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CollectPaymentMethod(id string, params *stripe.TerminalReaderCollectPaymentMethodParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finalizes a payment on a Reader. See [Confirming a Payment](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#confirm-the-paymentintent) for more details.
func ConfirmPaymentIntent(id string, params *stripe.TerminalReaderConfirmPaymentIntentParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finalizes a payment on a Reader. See [Confirming a Payment](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#confirm-the-paymentintent) for more details.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ConfirmPaymentIntent(id string, params *stripe.TerminalReaderConfirmPaymentIntentParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates a payment flow on a Reader. See [process the payment](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=immediately#process-payment) for more details.
func ProcessPaymentIntent(id string, params *stripe.TerminalReaderProcessPaymentIntentParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates a payment flow on a Reader. See [process the payment](https://docs.stripe.com/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=immediately#process-payment) for more details.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ProcessPaymentIntent(id string, params *stripe.TerminalReaderProcessPaymentIntentParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates a SetupIntent flow on a Reader. See [Save directly without charging](https://docs.stripe.com/docs/terminal/features/saving-payment-details/save-directly) for more details.
func ProcessSetupIntent(id string, params *stripe.TerminalReaderProcessSetupIntentParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates a SetupIntent flow on a Reader. See [Save directly without charging](https://docs.stripe.com/docs/terminal/features/saving-payment-details/save-directly) for more details.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ProcessSetupIntent(id string, params *stripe.TerminalReaderProcessSetupIntentParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates an in-person refund on a Reader. See [Refund an Interac Payment](https://docs.stripe.com/docs/terminal/payments/regional?integration-country=CA#refund-an-interac-payment) for more details.
func RefundPayment(id string, params *stripe.TerminalReaderRefundPaymentParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiates an in-person refund on a Reader. See [Refund an Interac Payment](https://docs.stripe.com/docs/terminal/payments/regional?integration-country=CA#refund-an-interac-payment) for more details.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RefundPayment(id string, params *stripe.TerminalReaderRefundPaymentParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sets the reader display to show [cart details](https://docs.stripe.com/docs/terminal/features/display).
func SetReaderDisplay(id string, params *stripe.TerminalReaderSetReaderDisplayParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sets the reader display to show [cart details](https://docs.stripe.com/docs/terminal/features/display).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SetReaderDisplay(id string, params *stripe.TerminalReaderSetReaderDisplayParams) (*stripe.TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of Reader objects.
func List(params *stripe.TerminalReaderListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of Reader objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TerminalReaderListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for terminal readers.
type Iter struct {
	*stripe.Iter
}

// TerminalReader returns the terminal reader which the iterator is currently pointing to.
func (i *Iter) TerminalReader() *stripe.TerminalReader { _ = "STUB: not implemented"; return nil }

// TerminalReaderList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TerminalReaderList() *stripe.TerminalReaderList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
