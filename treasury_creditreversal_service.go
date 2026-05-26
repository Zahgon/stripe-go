//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryCreditReversalService is used to invoke /v1/treasury/credit_reversals APIs.
type v1TreasuryCreditReversalService struct {
	B   Backend
	Key string
}

// Reverses a ReceivedCredit and creates a CreditReversal object.
func (c v1TreasuryCreditReversalService) Create(ctx context.Context, params *TreasuryCreditReversalCreateParams) (*TreasuryCreditReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of an existing CreditReversal by passing the unique CreditReversal ID from either the CreditReversal creation request or CreditReversal list
func (c v1TreasuryCreditReversalService) Retrieve(ctx context.Context, id string, params *TreasuryCreditReversalRetrieveParams) (*TreasuryCreditReversal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of CreditReversals.
func (c v1TreasuryCreditReversalService) List(ctx context.Context, listParams *TreasuryCreditReversalListParams) *V1List[*TreasuryCreditReversal] {
	_ = "STUB: not implemented"
	return nil
}
