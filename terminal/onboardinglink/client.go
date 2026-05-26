//
//
// File generated from our OpenAPI spec
//
//

// Package onboardinglink provides the /v1/terminal/onboarding_links APIs
package onboardinglink

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/terminal/onboarding_links APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new OnboardingLink object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.
func New(params *stripe.TerminalOnboardingLinkParams) (*stripe.TerminalOnboardingLink, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new OnboardingLink object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TerminalOnboardingLinkParams) (*stripe.TerminalOnboardingLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
