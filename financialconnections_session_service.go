//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1FinancialConnectionsSessionService is used to invoke /v1/financial_connections/sessions APIs.
type v1FinancialConnectionsSessionService struct {
	B   Backend
	Key string
}

// To launch the Financial Connections authorization flow, create a Session. The session's client_secret can be used to launch the flow using Stripe.js.
func (c v1FinancialConnectionsSessionService) Create(ctx context.Context, params *FinancialConnectionsSessionCreateParams) (*FinancialConnectionsSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a Financial Connections Session
func (c v1FinancialConnectionsSessionService) Retrieve(ctx context.Context, id string, params *FinancialConnectionsSessionRetrieveParams) (*FinancialConnectionsSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
