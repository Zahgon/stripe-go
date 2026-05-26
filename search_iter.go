package stripe

import (
	"context"

	"github.com/stripe/stripe-go/v85/form"
)

//
// Public constants
//

// Contains constants for the names of parameters used for pagination in search APIs.
const (
	Page = "page"
)

//
// Public types
//

// SearchIter provides a convenient interface
// for iterating over the elements
// returned from paginated search API calls.
// Successive calls to the Next method
// will step through each item in the search results,
// fetching pages of items as needed.
// Iterators are not thread-safe, so they should not be consumed
// across multiple goroutines.
type SearchIter struct {
	cur             interface{}
	err             error
	formValues      *form.Values
	searchContainer SearchContainer
	searchParams    SearchParams
	meta            *SearchMeta
	query           SearchQuery
	values          []interface{}
}

// Current returns the most recent item
// visited by a call to Next.
func (it *SearchIter) Current() interface{} {
	_ = "STUB: not implemented"

	// Err returns the error, if any,
	// that caused the SearchIter to stop.
	// It must be inspected
	// after Next returns false.
	return nil
}

func (it *SearchIter) Err() error {
	_ = "STUB: not implemented"

	// SearchResult returns the current search result container which the iterator is currently using.
	// Objects will change as new API calls are made to continue pagination.
	return nil
}

func (it *SearchIter) SearchResult() SearchContainer {
	_ = "STUB: not implemented"
	return *new(SearchContainer)
}

// Meta returns the search metadata.
func (it *SearchIter) Meta() *SearchMeta {
	_ = "STUB: not implemented"

	// Next advances the SearchIter to the next item in the search results,
	// which will then be available
	// through the Current method.
	// It returns false when the iterator stops
	// at the end of the search results.
	return nil
}

func (it *SearchIter) Next() bool { _ = "STUB: not implemented"; return false }

func (it *SearchIter) getPage() { _ = "STUB: not implemented"; return }

// SearchQuery is the function used to get search results.
type SearchQuery func(*Params, *form.Values) ([]interface{}, SearchContainer, error)

//
// Public functions
//

// GetSearchIter returns a new SearchIter for a given query and its options.
func GetSearchIter(container SearchParamsContainer, query SearchQuery) *SearchIter {
	_ = "STUB: not implemented"
	return nil
}

// See the comment on Call in stripe.go.

// V1SearchList provides a convenient interface for iterating over the elements
// returned from paginated list API calls. It is meant to be an improvement
// over the SearchIter type, which was written before Go introduced generics and iter.Seq2.
// Calling the `All` allows you to iterate over all items in the list,
// with automatic pagination.
type V1SearchList[T any] struct {
	err          error
	formValues   *form.Values
	searchParams SearchParams
	query        v1SearchQuery[T]
	v1SearchPage *v1SearchPage[T]
}

// v1SearchPage represents a single page returned from a V1 Search API call.
// The internal state will be updated by the parent V1SearchList when the
// Page method is called.
type v1SearchPage[T any] struct {
	APIResource
	SearchMeta
	Data []T `json:"data"`
}

// Data returns the data for the current page.
func (l *V1SearchList[T]) Data() []T { _ = "STUB: not implemented"; return nil }

// Err returns the error for the current page.
func (l *V1SearchList[T]) Err() error {
	_ = "STUB: not implemented"

	// Meta returns the metadata for the current page.
	return nil
}

func (l *V1SearchList[T]) Meta() SearchMeta { _ = "STUB: not implemented"; return *new(SearchMeta) }

// LastResponse returns the last response for the current page.
func (l *V1SearchList[T]) LastResponse() *APIResponse { _ = "STUB: not implemented"; return nil }

// All returns a Seq2 that will be evaluated on each item in a V1SearchList.
// The All function will continue to fetch pages of items as needed.
func (l *V1SearchList[T]) All(ctx context.Context) Seq2[T, error] {
	_ = "STUB: not implemented"
	return nil
}

// page fetches the next page of items and updates the V1SearchList's state.
func (l *V1SearchList[T]) page(ctx context.Context) { _ = "STUB: not implemented"; return }

// hasMore returns true if there is another page of items to fetch.
func (l *V1SearchList[T]) hasMore() bool { _ = "STUB: not implemented"; return false }

// maybeAddLastResponse adds the LastResponse to the items in the page.
// It parses the page's JSON and adds each `data` item's JSON to the
// LastResponse of the corresponding resource. Note that not
// every resource implements the LastResponseSetter interface.
func maybeAddLastResponseSearch[T any](page *v1SearchPage[T]) error {
	_ = "STUB: not implemented"
	return nil
}

// Note that not every resource implements the LastResponseSetter interface
// (e.g. CreditNoteLineItem).

// Create a copy of the original response with individual item's raw JSON

// v1SearchQuery is the function used to get search results.
type v1SearchQuery[T any] func(context.Context, *Params, *form.Values) (*v1SearchPage[T], error)

//
// Public functions
//

// newV1SearchList returns a new V1SearchList for a given query and its options, and initializes
// it by fetching the first page of items.
func newV1SearchList[T any](ctx context.Context, container SearchParamsContainer, query v1SearchQuery[T]) *V1SearchList[T] {
	_ = "STUB: not implemented"
	return nil
}

// This is a little unfortunate, but Go makes it impossible to compare
// an interface value to nil without the use of the reflect package and
// its true disciples insist that this is a feature and not a bug.
//
// Here we do invoke reflect because (1) we have to reflect anyway to
// use encode with the form package, and (2) the corresponding removal
// of boilerplate that this enables makes the small performance penalty
// worth it.
