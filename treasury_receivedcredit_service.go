//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TreasuryReceivedCreditService is used to invoke /v1/treasury/received_credits APIs.
type v1TreasuryReceivedCreditService struct {
	B   Backend
	Key string
}

// Retrieves the details of an existing ReceivedCredit by passing the unique ReceivedCredit ID from the ReceivedCredit list.
func (c v1TreasuryReceivedCreditService) Retrieve(ctx context.Context, id string, params *TreasuryReceivedCreditRetrieveParams) (*TreasuryReceivedCredit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of ReceivedCredits.
func (c v1TreasuryReceivedCreditService) List(ctx context.Context, listParams *TreasuryReceivedCreditListParams) *V1List[*TreasuryReceivedCredit] {
	_ = "STUB: not implemented"
	return nil
}
