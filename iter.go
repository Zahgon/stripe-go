package stripe

import (
	"context"

	"github.com/stripe/stripe-go/v85/form"
)

// Iter provides a convenient interface
// for iterating over the elements
// returned from paginated list API calls.
// Successive calls to the Next method
// will step through each item in the list,
// fetching pages of items as needed.
// Iterators are not thread-safe, so they should not be consumed
// across multiple goroutines.
type Iter struct {
	cur        interface{}
	err        error
	formValues *form.Values
	list       ListContainer
	listParams ListParams
	meta       *ListMeta
	query      Query
	values     []interface{}
}

// Current returns the most recent item
// visited by a call to Next.
func (it *Iter) Current() interface{} {
	_ = "STUB: not implemented"

	// Err returns the error, if any,
	// that caused the Iter to stop.
	// It must be inspected
	// after Next returns false.
	return nil
}

func (it *Iter) Err() error {
	_ = "STUB: not implemented"

	// List returns the current list object which the iterator is currently using.
	// List objects will change as new API calls are made to continue pagination.
	return nil
}

func (it *Iter) List() ListContainer {
	_ = "STUB: not implemented"

	// Meta returns the list metadata.
	return *new(ListContainer)
}

func (it *Iter) Meta() *ListMeta {
	_ = "STUB: not implemented"

	// Next advances the Iter to the next item in the list,
	// which will then be available
	// through the Current method.
	// It returns false when the iterator stops
	// at the end of the list.
	return nil
}

func (it *Iter) Next() bool { _ = "STUB: not implemented"; return false }

// determine if we're moving forward or backwards in paging

func (it *Iter) getPage() { _ = "STUB: not implemented"; return }

// We are moving backward,
// but items arrive in forward order.

// Query is the function used to get a page listing.
type Query func(*Params, *form.Values) ([]interface{}, ListContainer, error)

// GetIter returns a new Iter for a given query and its options.
func GetIter(container ListParamsContainer, query Query) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// See the comment on Call in stripe.go.

// V1List provides a convenient interface for iterating over the elements
// returned from paginated list API calls. It is meant to be an improvement
// over the Iter type, which was written before Go introduced generics and iter.Seq2.
// Calling the `All` allows you to iterate over all items in the list,
// with automatic pagination.
type V1List[T any] struct {
	err        error
	formValues *form.Values
	listParams ListParams
	query      v1Query[T]
	backward   bool
	v1Page     *v1Page[T]
}

// v1Page represents a single page returned from a V1 List API call.
// The internal state will be updated by the parent V1List when the
// Page method is called.
type v1Page[T any] struct {
	APIResource
	ListMeta
	Data []T `json:"data"`
}

// All returns a Seq2 that will be evaluated on each item in a V1List.
// The All function will continue to fetch pages of items as needed.
func (l *V1List[T]) All(ctx context.Context) Seq2[T, error] { _ = "STUB: not implemented"; return nil }

// Data returns the data for the current page.
func (l *V1List[T]) Data() []T { _ = "STUB: not implemented"; return nil }

// Err returns the error for the current page.
func (l *V1List[T]) Err() error {
	_ = "STUB: not implemented"

	// Meta returns the metadata for the current page.
	return nil
}

func (l *V1List[T]) Meta() ListMeta {
	_ = "STUB: not implemented"
	return *

	// LastResponse returns the last response for the current page.
	new(ListMeta)
}

func (l *V1List[T]) LastResponse() *APIResponse { _ = "STUB: not implemented"; return nil }

// page updates the V1List's state by fetching the next page of items.
func (l *V1List[T]) page(ctx context.Context) { _ = "STUB: not implemented"; return }

// We are moving backward,
// but items arrive in forward order.

// hasMore returns true if there is another page of items to fetch.
func (l *V1List[T]) hasMore() bool { _ = "STUB: not implemented"; return false }

// maybeAddLastResponseV1 adds the LastResponse to the items in the page.
// It parses the page's JSON and adds each `data` item's JSON to the
// LastResponse of the corresponding resource. Note that not
// every resource implements the LastResponseSetter interface.
func maybeAddLastResponseV1[T any](page *v1Page[T]) error { _ = "STUB: not implemented"; return nil }

// Note that not every resource implements the LastResponseSetter interface
// (e.g. CreditNoteLineItem).

// Create a copy of the original response with individual item's raw JSON

// v1Query is the function used to get a page listing.
type v1Query[T any] func(context.Context, *Params, *form.Values) (*v1Page[T], error)

// newV1List returns a new v1List for a given query and its options, and initializes
// it by fetching the first page of items.
func newV1List[T any](ctx context.Context, container ListParamsContainer, query v1Query[T]) *V1List[T] {
	_ = "STUB: not implemented"
	return nil
}

// See the comment on Call in stripe.go.

func listItemID[T any](x T) string { _ = "STUB: not implemented"; return "" }

func reverse[T any](a []T) { _ = "STUB: not implemented"; return }

// Seq2 is the same as the iter.Seq2 type in Go 1.23+. It is used as the return type
// of List methods. If you are using Go 1.23+, you can just range over the an List
// method directly, e.g.,
//
//	for event, err := range sc.V2CoreEvents.List(...) {
//		// check err and do something with event
//	}
//
// For older versions of Go, the yield function should return false
// to stop iteration or true to continue.
type Seq2[K, V any] func(yield func(K, V) bool)

// V2List contains a page of data received from a List API call,
// and the means to paginate to the next page of data via the fetch function.
type V2List[T any] struct {
	fetch       v2Query[T]
	params      ParamsContainer
	initialized bool
	err         error
	// Page contains the items returned from the last API call.
	v2Page *V2Page[T]
}

// V2Page is represents a single page returned from a V2 List API call.
type V2Page[T any] struct {
	APIResource
	V2ListMeta
	Data []T `json:"data"`
}

// Data returns the data for the current page.
func (l *V2List[T]) Data() []T { _ = "STUB: not implemented"; return nil }

// Err returns the error for the current page.
func (l *V2List[T]) Err() error {
	_ = "STUB: not implemented"

	// Meta returns the metadata for the current page.
	return nil
}

func (l *V2List[T]) Meta() V2ListMeta { _ = "STUB: not implemented"; return *new(V2ListMeta) }

// LastResponse returns the last response for the current page.
func (l *V2List[T]) LastResponse() *APIResponse { _ = "STUB: not implemented"; return nil }

// All returns a Seq2 that will be evaluated on each item in a V2List.
// The All function will continue to fetch pages of items as needed.
func (l *V2List[T]) All(ctx context.Context) Seq2[T, error] { _ = "STUB: not implemented"; return nil }

// page fetches the next page of items and updates the V2List's state.
// It returns an error if the fetch fails.
func (l *V2List[T]) page(ctx context.Context) {
	_ = "STUB: not implemented"
	// if we've already fetched a page, the next page URL
	// already contains all of the query parameters
	return
}

// maybeAddLastResponseV2 adds the LastResponse to the items in the page.
// It parses the page's JSON and adds each `data` item's JSON to the
// LastResponse of the corresponding resource. Note that not
// every resource implements the LastResponseSetter interface.
func maybeAddLastResponseV2[T any](page *V2Page[T]) error { _ = "STUB: not implemented"; return nil }

// Note that not every resource implements the LastResponseSetter interface
// (e.g. CreditNoteLineItem).

// Create a copy of the original response with individual item's raw JSON

// hasMore returns true if there is another page of items to fetch.
func (l *V2List[T]) hasMore() bool { _ = "STUB: not implemented"; return false }

// newV2List creates a new V2List with the given path and fetch function.
func newV2List[T any](ctx context.Context, path string, p ParamsContainer, fetch v2Query[T]) *V2List[T] {
	_ = "STUB: not implemented"
	return nil
}

// v2Query is a function that fetches a page of items.
type v2Query[T any] func(ctx context.Context, path string, p ParamsContainer) (*V2Page[T], error)

// Fetch is a function that fetches a page of items.
// Deprecated: This type is intended for internal use only, and will be removed in a future version.
type Fetch[T any] func(path string, p ParamsContainer) (*V2Page[T], error)

// NewV2List creates a new V2List with the given path and fetch function.
// Deprecated: This function is intended for internal use only, and will be removed in a future version.
func NewV2List[T any](path string, p ParamsContainer, fetch Fetch[T]) *V2List[T] {
	_ = "STUB: not implemented"
	return nil
}
