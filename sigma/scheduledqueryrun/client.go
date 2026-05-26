//
//
// File generated from our OpenAPI spec
//
//

// Package scheduledqueryrun provides the /v1/sigma/scheduled_query_runs APIs
// For more details, see: https://stripe.com/docs/api#scheduled_queries
package scheduledqueryrun

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/sigma/scheduled_query_runs APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an scheduled query run.
func Get(id string, params *stripe.SigmaScheduledQueryRunParams) (*stripe.SigmaScheduledQueryRun, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an scheduled query run.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.SigmaScheduledQueryRunParams) (*stripe.SigmaScheduledQueryRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of scheduled query runs.
func List(params *stripe.SigmaScheduledQueryRunListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of scheduled query runs.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.SigmaScheduledQueryRunListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for sigma scheduled query runs.
type Iter struct {
	*stripe.Iter
}

// SigmaScheduledQueryRun returns the sigma scheduled query run which the iterator is currently pointing to.
func (i *Iter) SigmaScheduledQueryRun() *stripe.SigmaScheduledQueryRun {
	_ = "STUB: not implemented"
	return nil
}

// SigmaScheduledQueryRunList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SigmaScheduledQueryRunList() *stripe.SigmaScheduledQueryRunList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
