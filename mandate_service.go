//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1MandateService is used to invoke /v1/mandates APIs.
type v1MandateService struct {
	B   Backend
	Key string
}

// Retrieves a Mandate object.
func (c v1MandateService) Retrieve(ctx context.Context, id string, params *MandateRetrieveParams) (*Mandate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
