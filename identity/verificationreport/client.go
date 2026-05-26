//
//
// File generated from our OpenAPI spec
//
//

// Package verificationreport provides the /v1/identity/verification_reports APIs
package verificationreport

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/identity/verification_reports APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an existing VerificationReport
func Get(id string, params *stripe.IdentityVerificationReportParams) (*stripe.IdentityVerificationReport, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves an existing VerificationReport
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.IdentityVerificationReportParams) (*stripe.IdentityVerificationReport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List all verification reports.
func List(params *stripe.IdentityVerificationReportListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// List all verification reports.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.IdentityVerificationReportListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for identity verification reports.
type Iter struct {
	*stripe.Iter
}

// IdentityVerificationReport returns the identity verification report which the iterator is currently pointing to.
func (i *Iter) IdentityVerificationReport() *stripe.IdentityVerificationReport {
	_ = "STUB: not implemented"
	return nil
}

// IdentityVerificationReportList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IdentityVerificationReportList() *stripe.IdentityVerificationReportList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
