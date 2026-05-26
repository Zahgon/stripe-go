//
//
// File generated from our OpenAPI spec
//
//

// Package personalizationdesign provides the /v1/issuing/personalization_designs APIs
package personalizationdesign

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/issuing/personalization_designs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Updates the status of the specified testmode personalization design object to active.
func Activate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignActivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the status of the specified testmode personalization design object to active.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Activate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignActivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the status of the specified testmode personalization design object to inactive.
func Deactivate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignDeactivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the status of the specified testmode personalization design object to inactive.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Deactivate(id string, params *stripe.TestHelpersIssuingPersonalizationDesignDeactivateParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the status of the specified testmode personalization design object to rejected.
func Reject(id string, params *stripe.TestHelpersIssuingPersonalizationDesignRejectParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the status of the specified testmode personalization design object to rejected.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Reject(id string, params *stripe.TestHelpersIssuingPersonalizationDesignRejectParams) (*stripe.IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
