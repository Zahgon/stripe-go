//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TerminalOnboardingLinkService is used to invoke /v1/terminal/onboarding_links APIs.
type v1TerminalOnboardingLinkService struct {
	B   Backend
	Key string
}

// Creates a new OnboardingLink object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.
func (c v1TerminalOnboardingLinkService) Create(ctx context.Context, params *TerminalOnboardingLinkCreateParams) (*TerminalOnboardingLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
