//
//
// File generated from our OpenAPI spec
//
//

// Package reporttype provides the /v1/reporting/report_types APIs
package reporttype

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/reporting/report_types APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Report Type. (Certain report types require a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).)
func Get(id string, params *stripe.ReportingReportTypeParams) (*stripe.ReportingReportType, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves the details of a Report Type. (Certain report types require a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).)
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.ReportingReportTypeParams) (*stripe.ReportingReportType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a full list of Report Types.
func List(params *stripe.ReportingReportTypeListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a full list of Report Types.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.ReportingReportTypeListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for reporting report types.
type Iter struct {
	*stripe.Iter
}

// ReportingReportType returns the reporting report type which the iterator is currently pointing to.
func (i *Iter) ReportingReportType() *stripe.ReportingReportType {
	_ = "STUB: not implemented"
	return nil
}

// ReportingReportTypeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ReportingReportTypeList() *stripe.ReportingReportTypeList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
