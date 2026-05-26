//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersTreasuryReceivedDebitService is used to invoke /v1/treasury/received_debits APIs.
type v1TestHelpersTreasuryReceivedDebitService struct {
	B   Backend
	Key string
}

// Use this endpoint to simulate a test mode ReceivedDebit initiated by a third party. In live mode, you can't directly create ReceivedDebits initiated by third parties.
func (c v1TestHelpersTreasuryReceivedDebitService) Create(ctx context.Context, params *TestHelpersTreasuryReceivedDebitCreateParams) (*TreasuryReceivedDebit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
