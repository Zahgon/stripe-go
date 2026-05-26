package stripe

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stripe/stripe-go/v85/form"
)

//
// Public constants
//

// Contains constants for the names of parameters used for pagination in list APIs.
const (
	EndingBefore  = "ending_before"
	StartingAfter = "starting_after"
)

//
// Public types
//

// ExtraValues are extra parameters that are attached to an API request.
// They're implemented as a custom type so that they can have their own
// AppendTo implementation.
type ExtraValues struct {
	url.Values `form:"-" json:"-"` // See custom AppendTo implementation
}

// AppendTo implements custom form encoding for extra parameter values.
func (v ExtraValues) AppendTo(body *form.Values, keyParts []string) {
	_ = "STUB: not implemented"
	return
}

// Filters is a structure that contains a collection of filters for list-related APIs.
type Filters struct {
	f []*filter `form:"-" json:"-"` // See custom AppendTo implementation
}

// AddFilter adds a new filter with a given key, op and value.
func (f *Filters) AddFilter(key, op, value string) { _ = "STUB: not implemented"; return }

// AppendTo implements custom form encoding for filters.
func (f Filters) AppendTo(body *form.Values, keyParts []string) { _ = "STUB: not implemented"; return }

// ListContainer is a general interface for which all list object structs
// should comply. They achieve this by embedding a ListMeta struct and
// inheriting its implementation of this interface.
type ListContainer interface {
	GetListMeta() *ListMeta
}

// ListMeta is the structure that contains the common properties
// of List iterators. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type ListMeta struct {
	HasMore bool   `json:"has_more"`
	URL     string `json:"url"`

	// TotalCount is the total number of objects in the collection (beyond just
	// on the current page). This is not returned in most list calls.
	//
	// Deprecated: TotalCount is only included in some legacy situations and
	// not generally available anymore.
	TotalCount uint32 `json:"total_count"`
}

// GetListMeta returns a ListMeta struct (itself). It exists because any
// structs that embed ListMeta will inherit it, and thus implement the
// ListContainer interface.
func (l *ListMeta) GetListMeta() *ListMeta { _ = "STUB: not implemented"; return nil }

type V2ListMeta struct {
	NextPageURL     string `json:"next_page_url"`
	PreviousPageURL string `json:"previous_page_url"`
}

// ListParams is the structure that contains the common properties
// of any *ListParams structure.
type ListParams struct {
	// Context used for request. It may carry deadlines, cancellation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on Stripe's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `form:"-"`

	EndingBefore *string `form:"ending_before"`
	// Deprecated: Please use Expand in the surrounding struct instead.
	Expand  []*string `form:"expand"`
	Filters Filters   `form:"*"`
	Limit   *int64    `form:"limit"`

	// Single specifies whether this is a single page iterator. By default,
	// listing through an iterator will automatically grab additional pages as
	// the query progresses. To change this behavior and just load a single
	// page, set this to true.
	Single bool `form:"-"` // Not an API parameter

	StartingAfter *string `form:"starting_after"`

	// StripeAccount may contain the ID of a connected account. By including
	// this field, the request is made as if it originated from the connected
	// account instead of under the account of the owner of the configured
	// Stripe key.
	StripeAccount *string `form:"-"` // Passed as header

	// StripeContext is used to set the Stripe-Context header on a request.
	// The Stripe-Context header can be used to set the account with which
	// the request is made. It is preferred to using StripeAccount in new code.
	StripeContext *string `form:"-" json:"-"` // Passed as header
}

// AddExpand on the embedded ListParams struct is deprecated.
// Deprecated: please use AddExpand on the surrounding struct instead.
func (p *ListParams) AddExpand(f string) { _ = "STUB: not implemented"; return }

// GetListParams returns a ListParams struct (itself). It exists because any
// structs that embed ListParams will inherit it, and thus implement the
// ListParamsContainer interface.
func (p *ListParams) GetListParams() *ListParams {
	_ = "STUB: not implemented"

	// GetParams returns ListParams as a Params struct. It exists because any
	// structs that embed Params will inherit it, and thus implement the
	// ParamsContainer interface.
	return nil
}

func (p *ListParams) GetParams() *Params { _ = "STUB: not implemented"; return nil }

// SetStripeAccount sets a value for the Stripe-Account header.
func (p *ListParams) SetStripeAccount(val string) { _ = "STUB: not implemented"; return }

// SetStripeContext sets a value for the Stripe-Context header.
func (p *ListParams) SetStripeContext(val string) { _ = "STUB: not implemented"; return }

// SetStripeContextFrom sets a value for the Stripe-Context header using a stripe.Context object.
func (p *ListParams) SetStripeContextFrom(val *Context) { _ = "STUB: not implemented"; return }

// ToParams converts a ListParams to a Params by moving over any fields that
// have valid targets in the new type. This is useful because fields in
// Params can be injected directly into an http.Request while generally
// ListParams is only used to build a set of parameters.
func (p *ListParams) ToParams() *Params { _ = "STUB: not implemented"; return nil }

// ListParamsContainer is a general interface for which all list parameter
// structs should comply. They achieve this by embedding a ListParams struct
// and inheriting its implementation of this interface.
type ListParamsContainer interface {
	GetListParams() *ListParams
}

type APIMode string

var V1APIMode APIMode = "v1"
var V2APIMode APIMode = "v2"

func (m APIMode) contentType() string { _ = "STUB: not implemented"; return "" }

// The only way we can get here is if someone has mutated the APIMode
// variables, which would lead to unexpected behavior.

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	// Context used for request. It may carry deadlines, cancellation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on Stripe's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `form:"-" json:"-"`

	// Deprecated: please use Expand in the surrounding struct instead.
	Expand []*string    `form:"expand" json:"-"`
	Extra  *ExtraValues `form:"*"  json:"-"`

	// Headers may be used to provide extra header lines on the HTTP request.
	Headers http.Header `form:"-" json:"-"`

	IdempotencyKey *string `form:"-" json:"-"` // Passed as header

	// Deprecated: Please use Metadata in the surrounding struct instead.
	Metadata map[string]string `form:"metadata" json:"-"`

	// StripeAccount may contain the ID of a connected account. By including
	// this field, the request is made as if it originated from the connected
	// account instead of under the account of the owner of the configured
	// Stripe key.
	StripeAccount *string `form:"-" json:"-"` // Passed as header

	// StripeContext is used to set the Stripe-Context header on a request.
	// The Stripe-Context header can be used to set the account with which
	// the request is made. It is preferred to using StripeAccount in new code.
	StripeContext *string `form:"-" json:"-"` // Passed as header

	usage []string `form:"-" json:"-"` // Tracked behaviors
}

// AddExpand on the Params embedded struct is deprecated.
// Deprecated: please use Expand in the surrounding struct instead.
func (p *Params) AddExpand(f string) { _ = "STUB: not implemented"; return }

// InternalSetUsage sets the usage field on the Params struct, removing duplicates.
// Unstable: for internal stripe-go usage only.
func (p *Params) InternalSetUsage(usage []string) {
	_ = "STUB: not implemented"
	// Optimization for nil or empty usage
	return
}

// Use a map to track unique usage values

// Reset the slice to avoid retaining old values

// AddExtra adds a new arbitrary key-value pair to the request data
func (p *Params) AddExtra(key, value string) { _ = "STUB: not implemented"; return }

// AddMetadata on the Params embedded struct is deprecated.
// Deprecated: please use .AddMetadata of the surrounding struct.
func (p *Params) AddMetadata(key, value string) { _ = "STUB: not implemented"; return }

// GetParams returns a Params struct (itself). It exists because any structs
// that embed Params will inherit it, and thus implement the ParamsContainer
// interface.
func (p *Params) GetParams() *Params {
	_ = "STUB: not implemented"

	// SetIdempotencyKey sets a value for the Idempotency-Key header.
	return nil
}

func (p *Params) SetIdempotencyKey(val string) { _ = "STUB: not implemented"; return }

// SetStripeAccount sets a value for the Stripe-Account header.
func (p *Params) SetStripeAccount(val string) { _ = "STUB: not implemented"; return }

// SetStripeContext sets a value for the Stripe-Context header.
func (p *Params) SetStripeContext(val string) { _ = "STUB: not implemented"; return }

// SetStripeContextFrom sets a value for the Stripe-Context header using a stripe.Context object.
func (p *Params) SetStripeContextFrom(val *Context) { _ = "STUB: not implemented"; return }

// ParamsContainer is a general interface for which all parameter structs
// should comply. They achieve this by embedding a Params struct and inheriting
// its implementation of this interface.
type ParamsContainer interface {
	GetParams() *Params
}

type RawParams struct {
	Params        `form:"*"`
	StripeContext string `form:"-"`
}

// RangeQueryParams are a set of generic request parameters that are used on
// list endpoints to filter their results by some timestamp.
type RangeQueryParams struct {
	// GreaterThan specifies that values should be a greater than this
	// timestamp.
	GreaterThan int64 `form:"gt"`

	// GreaterThanOrEqual specifies that values should be greater than or equal
	// to this timestamp.
	GreaterThanOrEqual int64 `form:"gte"`

	// LesserThan specifies that values should be lesser than this timestamp.
	LesserThan int64 `form:"lt"`

	// LesserThanOrEqual specifies that values should be lesser than or
	// equal to this timestamp.
	LesserThanOrEqual int64 `form:"lte"`
}

//
// Public functions
//

// NewIdempotencyKey generates a new idempotency key that
// can be used on a request.
func NewIdempotencyKey() string { _ = "STUB: not implemented"; return "" }

//
// Private types
//

// filter is the structure that contains a filter for list-related APIs.
// It ends up passing query string parameters in the format key[op]=value.
type filter struct {
	Key, Op, Val string
}
