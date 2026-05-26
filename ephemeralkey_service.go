//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1EphemeralKeyService is used to invoke /v1/ephemeral_keys APIs.
type v1EphemeralKeyService struct {
	B   Backend
	Key string
}

// Creates a short-lived API key for a given resource.
func (c v1EphemeralKeyService) Create(ctx context.Context, params *EphemeralKeyCreateParams) (*EphemeralKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Invalidates a short-lived API key for a given resource.
func (c v1EphemeralKeyService) Delete(ctx context.Context, id string, params *EphemeralKeyDeleteParams) (*EphemeralKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
