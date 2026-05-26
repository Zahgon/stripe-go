//
//
// File generated from our OpenAPI spec
//
//

// Package card provides the /v1/issuing/cards APIs
package card

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/cards APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Updates the shipping status of the specified Issuing Card object to delivered.
func DeliverCard(id string, params *stripe.TestHelpersIssuingCardDeliverCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to delivered.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) DeliverCard(id string, params *stripe.TestHelpersIssuingCardDeliverCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to failure.
func FailCard(id string, params *stripe.TestHelpersIssuingCardFailCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to failure.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) FailCard(id string, params *stripe.TestHelpersIssuingCardFailCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to returned.
func ReturnCard(id string, params *stripe.TestHelpersIssuingCardReturnCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to returned.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReturnCard(id string, params *stripe.TestHelpersIssuingCardReturnCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to shipped.
func ShipCard(id string, params *stripe.TestHelpersIssuingCardShipCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to shipped.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ShipCard(id string, params *stripe.TestHelpersIssuingCardShipCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to submitted. This method requires Stripe Version ‘2024-09-30.acacia' or later.
func SubmitCard(id string, params *stripe.TestHelpersIssuingCardSubmitCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the shipping status of the specified Issuing Card object to submitted. This method requires Stripe Version ‘2024-09-30.acacia' or later.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) SubmitCard(id string, params *stripe.TestHelpersIssuingCardSubmitCardParams) (*stripe.IssuingCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
