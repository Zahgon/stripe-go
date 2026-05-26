//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersIssuingCardService is used to invoke /v1/issuing/cards APIs.
type v1TestHelpersIssuingCardService struct {
	B   Backend
	Key string
}

// Updates the shipping status of the specified Issuing Card object to delivered.
func (c v1TestHelpersIssuingCardService) DeliverCard(ctx context.Context, id string, params *TestHelpersIssuingCardDeliverCardParams) (*IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to failure.
func (c v1TestHelpersIssuingCardService) FailCard(ctx context.Context, id string, params *TestHelpersIssuingCardFailCardParams) (*IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to returned.
func (c v1TestHelpersIssuingCardService) ReturnCard(ctx context.Context, id string, params *TestHelpersIssuingCardReturnCardParams) (*IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to shipped.
func (c v1TestHelpersIssuingCardService) ShipCard(ctx context.Context, id string, params *TestHelpersIssuingCardShipCardParams) (*IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to submitted. This method requires Stripe Version ‘2024-09-30.acacia' or later.
func (c v1TestHelpersIssuingCardService) SubmitCard(ctx context.Context, id string, params *TestHelpersIssuingCardSubmitCardParams) (*IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
