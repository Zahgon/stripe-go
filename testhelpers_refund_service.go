//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1TestHelpersRefundService is used to invoke /v1/refunds APIs.
type v1TestHelpersRefundService struct {
	B   Backend
	Key string
}

// Expire a refund with a status of requires_action.
func (c v1TestHelpersRefundService) Expire(ctx context.Context, id string, params *TestHelpersRefundExpireParams) (*Refund, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
