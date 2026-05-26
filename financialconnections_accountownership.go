//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Describes a snapshot of the owners of an account at a particular point in time.
type FinancialConnectionsAccountOwnership struct {
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// A paginated list of owners for this account.
	Owners *FinancialConnectionsAccountOwnerList `json:"owners"`
}

// UnmarshalJSON handles deserialization of a FinancialConnectionsAccountOwnership.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *FinancialConnectionsAccountOwnership) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
