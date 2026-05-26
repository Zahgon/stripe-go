//
//
// File generated from our OpenAPI spec
//
//

// Package testclock provides the /v1/test_helpers/test_clocks APIs
package testclock

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/test_helpers/test_clocks APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new test clock that can be attached to new customers and quotes.
func New(params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil,

		// Creates a new test clock that can be attached to new customers and quotes.
		//
		// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
		//
		// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
		nil
}

func (c Client) New(params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a test clock.
func Get(id string, params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Retrieves a test clock.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Get(id string, params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletes a test clock.
func Del(id string, params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Deletes a test clock.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) Del(id string, params *stripe.TestHelpersTestClockParams) (*stripe.TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Starts advancing a test clock to a specified time in the future. Advancement is done when status changes to Ready.
func Advance(id string, params *stripe.TestHelpersTestClockAdvanceParams) (*stripe.TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Starts advancing a test clock to a specified time in the future. Advancement is done when status changes to Ready.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Advance(id string, params *stripe.TestHelpersTestClockAdvanceParams) (*stripe.TestHelpersTestClock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of your test clocks.
func List(params *stripe.TestHelpersTestClockListParams) *Iter {
	_ = "STUB: not implemented"
	return nil

	// Returns a list of your test clocks.
	//
	// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
	//
	// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
}

func (c Client) List(listParams *stripe.TestHelpersTestClockListParams) *Iter {
	_ = "STUB: not implemented"
	return nil
}

// Iter is an iterator for test helpers test clocks.
type Iter struct {
	*stripe.Iter
}

// TestHelpersTestClock returns the test helpers test clock which the iterator is currently pointing to.
func (i *Iter) TestHelpersTestClock() *stripe.TestHelpersTestClock {
	_ = "STUB: not implemented"
	return nil
}

// TestHelpersTestClockList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TestHelpersTestClockList() *stripe.TestHelpersTestClockList {
	_ = "STUB: not implemented"
	return nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
