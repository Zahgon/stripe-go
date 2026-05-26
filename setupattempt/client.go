//
//
// File generated from our OpenAPI spec
//
//

// Package setupattempt provides the /v1/setup_attempts APIs
// For more details, see: https://stripe.com/docs/api/?lang=go#setup_attempts
package setupattempt

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/setup_attempts APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Returns a list of SetupAttempts that associate with a provided SetupIntent.
func List(params *stripe.SetupAttemptListParams) *Iter { _ = "STUB: not implemented"; return nil }

// Returns a list of SetupAttempts that associate with a provided SetupIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.SetupAttemptListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for setup attempts.
type Iter struct {
	*stripe.Iter
}

// SetupAttempt returns the setup attempt which the iterator is currently pointing to.
func (i *Iter) SetupAttempt() *stripe.SetupAttempt { _ = "STUB: not implemented"; return nil }

// SetupAttemptList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SetupAttemptList() *stripe.SetupAttemptList { _ = "STUB: not implemented"; return nil }

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
