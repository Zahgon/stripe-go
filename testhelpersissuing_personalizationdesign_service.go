//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersIssuingPersonalizationDesignService is used to invoke /v1/issuing/personalization_designs APIs.
type v1TestHelpersIssuingPersonalizationDesignService struct {
	B   Backend
	Key string
}

// Updates the status of the specified testmode personalization design object to active.
func (c v1TestHelpersIssuingPersonalizationDesignService) Activate(ctx context.Context, id string, params *TestHelpersIssuingPersonalizationDesignActivateParams) (*IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the status of the specified testmode personalization design object to inactive.
func (c v1TestHelpersIssuingPersonalizationDesignService) Deactivate(ctx context.Context, id string, params *TestHelpersIssuingPersonalizationDesignDeactivateParams) (*IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the status of the specified testmode personalization design object to rejected.
func (c v1TestHelpersIssuingPersonalizationDesignService) Reject(ctx context.Context, id string, params *TestHelpersIssuingPersonalizationDesignRejectParams) (*IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
