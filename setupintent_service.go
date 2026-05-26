//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1SetupIntentService is used to invoke /v1/setup_intents APIs.
type v1SetupIntentService struct {
	B   Backend
	Key string
}

// Creates a SetupIntent object.
//
// After you create the SetupIntent, attach a payment method and [confirm](https://docs.stripe.com/docs/api/setup_intents/confirm)
// it to collect any required permissions to charge the payment method later.
func (c v1SetupIntentService) Create(ctx context.Context, params *SetupIntentCreateParams) (*SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a SetupIntent that has previously been created.
//
// Client-side retrieval using a publishable key is allowed when the client_secret is provided in the query string.
//
// When retrieved with a publishable key, only a subset of properties will be returned. Please refer to the [SetupIntent](https://docs.stripe.com/api#setup_intent_object) object reference for more details.
func (c v1SetupIntentService) Retrieve(ctx context.Context, id string, params *SetupIntentRetrieveParams) (*SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates a SetupIntent object.
func (c v1SetupIntentService) Update(ctx context.Context, id string, params *SetupIntentUpdateParams) (*SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// You can cancel a SetupIntent object when it's in one of these statuses: requires_payment_method, requires_confirmation, or requires_action.
//
// After you cancel it, setup is abandoned and any operations on the SetupIntent fail with an error. You can't cancel the SetupIntent for a Checkout Session. [Expire the Checkout Session](https://docs.stripe.com/docs/api/checkout/sessions/expire) instead.
func (c v1SetupIntentService) Cancel(ctx context.Context, id string, params *SetupIntentCancelParams) (*SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Confirm that your customer intends to set up the current or
// provided payment method. For example, you would confirm a SetupIntent
// when a customer hits the “Save” button on a payment method management
// page on your website.
//
// If the selected payment method does not require any additional
// steps from the customer, the SetupIntent will transition to the
// succeeded status.
//
// Otherwise, it will transition to the requires_action status and
// suggest additional actions via next_action. If setup fails,
// the SetupIntent will transition to the
// requires_payment_method status or the canceled status if the
// confirmation limit is reached.
func (c v1SetupIntentService) Confirm(ctx context.Context, id string, params *SetupIntentConfirmParams) (*SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verifies microdeposits on a SetupIntent object.
func (c v1SetupIntentService) VerifyMicrodeposits(ctx context.Context, id string, params *SetupIntentVerifyMicrodepositsParams) (*SetupIntent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of SetupIntents.
func (c v1SetupIntentService) List(ctx context.Context, listParams *SetupIntentListParams) *V1List[*SetupIntent] {
	_ = "STUB: not implemented"
	return nil
}
