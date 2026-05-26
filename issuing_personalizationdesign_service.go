//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1IssuingPersonalizationDesignService is used to invoke /v1/issuing/personalization_designs APIs.
type v1IssuingPersonalizationDesignService struct {
	B   Backend
	Key string
}

// Creates a personalization design object.
func (c v1IssuingPersonalizationDesignService) Create(ctx context.Context, params *IssuingPersonalizationDesignCreateParams) (*IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a personalization design object.
func (c v1IssuingPersonalizationDesignService) Retrieve(ctx context.Context, id string, params *IssuingPersonalizationDesignRetrieveParams) (*IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a card personalization object.
func (c v1IssuingPersonalizationDesignService) Update(ctx context.Context, id string, params *IssuingPersonalizationDesignUpdateParams) (*IssuingPersonalizationDesign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of personalization design objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingPersonalizationDesignService) List(ctx context.Context, listParams *IssuingPersonalizationDesignListParams) *V1List[*IssuingPersonalizationDesign] {
	_ = "STUB: not implemented"
	return nil
}
