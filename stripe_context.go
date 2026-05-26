package stripe

// Context provides a container and convenience methods for interacting with the `Stripe-Context` header. All methods return a new instance of Context.
// You can use its `StringPtr` method whenever you're initializing a `StripeClient` or sending `StripeContext` with a request. It's also found in the `EventNotification.Context` property.
type Context struct {
	Segments []string
}

// NewStripeContext creates a new stripe.Context with the given segments.
// If segments is nil or empty, creates an empty context.
func NewStripeContext(segments []string) *Context { _ = "STUB: not implemented"; return nil }

// Create a copy to ensure immutability

// ParseStripeContext parses a context string into a stripe.Context instance.
// If contextStr is empty, returns nil.
func ParseStripeContext(contextStr string) *Context { _ = "STUB: not implemented"; return nil }

// Push creates a new StripeContext with an additional segment appended.
// Returns an error if the segment is empty.
func (c *Context) Push(segment string) (*Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pop creates a new StripeContext with the last segment removed.
// If there are no segments, returns an error.
func (c *Context) Pop() (*Context, error) { _ = "STUB: not implemented"; return nil, nil }

// StringPtr returns the string representation of the stripe.Context.
// Segments are joined with "/" as the separator.
func (c *Context) StringPtr() *string { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface for stripe.Context.
func (c *Context) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
