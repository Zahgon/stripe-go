//
//
// File generated from our OpenAPI spec
//
//

// Package secret provides the /v1/apps/secrets APIs
package secret

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/apps/secrets APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create or replace a secret in the secret store.
func New(params *stripe.AppsSecretParams) (*stripe.AppsSecret, error) {
	_ = "STUB: not implemented"
	return nil,

		// Create or replace a secret in the secret store.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.AppsSecretParams) (*stripe.AppsSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a secret from the secret store by name and scope.
func DeleteWhere(params *stripe.AppsSecretDeleteWhereParams) (*stripe.AppsSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a secret from the secret store by name and scope.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) DeleteWhere(params *stripe.AppsSecretDeleteWhereParams) (*stripe.AppsSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finds a secret in the secret store by name and scope.
func Find(params *stripe.AppsSecretFindParams) (*stripe.AppsSecret, error) {
	_ = "STUB: not implemented"
	return nil,

		// Finds a secret in the secret store by name and scope.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) Find(params *stripe.AppsSecretFindParams) (*stripe.AppsSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List all secrets stored on the given scope.
func List(params *stripe.AppsSecretListParams) *Iter { _ = "STUB: not implemented"; return nil }

// List all secrets stored on the given scope.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.AppsSecretListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for apps secrets.
type Iter struct {
	*stripe.Iter
}

// AppsSecret returns the apps secret which the iterator is currently pointing to.
func (i *Iter) AppsSecret() *stripe.AppsSecret { _ = "STUB: not implemented"; return nil }

// AppsSecretList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) AppsSecretList() *stripe.AppsSecretList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
