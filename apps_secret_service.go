//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1AppsSecretService is used to invoke /v1/apps/secrets APIs.
type v1AppsSecretService struct {
	B   Backend
	Key string
}

// Create or replace a secret in the secret store.
func (c v1AppsSecretService) Create(ctx context.Context, params *AppsSecretCreateParams) (*AppsSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a secret from the secret store by name and scope.
func (c v1AppsSecretService) DeleteWhere(ctx context.Context, params *AppsSecretDeleteWhereParams) (*AppsSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finds a secret in the secret store by name and scope.
func (c v1AppsSecretService) Find(ctx context.Context, params *AppsSecretFindParams) (*AppsSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List all secrets stored on the given scope.
func (c v1AppsSecretService) List(ctx context.Context, listParams *AppsSecretListParams) *V1List[*AppsSecret] {
	_ = "STUB: not implemented"
	return nil
}
