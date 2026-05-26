//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1SourceTransactionService is used to invoke sourcetransaction related APIs.
type v1SourceTransactionService struct {
	B   Backend
	Key string
}

// List source transactions for a given source.
func (c v1SourceTransactionService) List(ctx context.Context, listParams *SourceTransactionListParams) *V1List[*SourceTransaction] {
	_ = "STUB: not implemented"
	return nil
}
