//
//
// File generated from our OpenAPI spec
//
//

// Package earlyfraudwarning provides the /v1/radar/early_fraud_warnings APIs
package earlyfraudwarning

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/radar/early_fraud_warnings APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an early fraud warning that has previously been created.
//
// Please refer to the [early fraud warning](https://docs.stripe.com/api#early_fraud_warning_object) object reference for more details.
func Get(id string, params *stripe.RadarEarlyFraudWarningParams) (*stripe.RadarEarlyFraudWarning, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of an early fraud warning that has previously been created.
	//
	// Please refer to the [early fraud warning](https://docs.stripe.com/api#early_fraud_warning_object) object reference for more details.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.RadarEarlyFraudWarningParams) (*stripe.RadarEarlyFraudWarning, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of early fraud warnings.
func List(params *stripe.RadarEarlyFraudWarningListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of early fraud warnings.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.RadarEarlyFraudWarningListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for radar early fraud warnings.
type Iter struct {
	*stripe.Iter
}

// RadarEarlyFraudWarning returns the radar early fraud warning which the iterator is currently pointing to.
func (i *Iter) RadarEarlyFraudWarning() *stripe.RadarEarlyFraudWarning {
	_ = "STUB: not implemented"
	return nil
}

// RadarEarlyFraudWarningList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) RadarEarlyFraudWarningList() *stripe.RadarEarlyFraudWarningList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
