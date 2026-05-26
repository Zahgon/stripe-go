//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersTerminalReaderService is used to invoke /v1/terminal/readers APIs.
type v1TestHelpersTerminalReaderService struct {
	B   Backend
	Key string
}

// Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.
func (c v1TestHelpersTerminalReaderService) PresentPaymentMethod(ctx context.Context, id string, params *TestHelpersTerminalReaderPresentPaymentMethodParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use this endpoint to trigger a successful input collection on a simulated reader.
func (c v1TestHelpersTerminalReaderService) SucceedInputCollection(ctx context.Context, id string, params *TestHelpersTerminalReaderSucceedInputCollectionParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use this endpoint to complete an input collection with a timeout error on a simulated reader.
func (c v1TestHelpersTerminalReaderService) TimeoutInputCollection(ctx context.Context, id string, params *TestHelpersTerminalReaderTimeoutInputCollectionParams) (*TerminalReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
