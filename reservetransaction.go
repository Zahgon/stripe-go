//
//
// File generated from our OpenAPI spec
//
//

package stripe

type ReserveTransaction struct {
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// UnmarshalJSON handles deserialization of a ReserveTransaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *ReserveTransaction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
