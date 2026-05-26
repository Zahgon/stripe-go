//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
)

// v1ChargeService is used to invoke /v1/charges APIs.
type v1ChargeService struct {
	B   Backend
	Key string
}

// This method is no longer recommended—use the [Payment Intents API](https://docs.stripe.com/docs/api/payment_intents)
// to initiate a new payment instead. Confirmation of the PaymentIntent creates the Charge
// object used to request payment.
func (c v1ChargeService) Create(ctx context.Context, params *ChargeCreateParams) (*Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the details of a charge that has previously been created. Supply the unique charge ID that was returned from your previous request, and Stripe will return the corresponding charge information. The same information is returned when creating or refunding the charge.
func (c v1ChargeService) Retrieve(ctx context.Context, id string, params *ChargeRetrieveParams) (*Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Updates the specified charge by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1ChargeService) Update(ctx context.Context, id string, params *ChargeUpdateParams) (*Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capture the payment of an existing, uncaptured charge that was created with the capture option set to false.
//
// Uncaptured payments expire a set number of days after they are created ([7 by default](https://docs.stripe.com/docs/charges/placing-a-hold)), after which they are marked as refunded and capture attempts will fail.
//
// Don't use this method to capture a PaymentIntent-initiated charge. Use [Capture a PaymentIntent](https://docs.stripe.com/docs/api/payment_intents/capture).
func (c v1ChargeService) Capture(ctx context.Context, id string, params *ChargeCaptureParams) (*Charge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns a list of charges you've previously created. The charges are returned in sorted order, with the most recent charges appearing first.
func (c v1ChargeService) List(ctx context.Context, listParams *ChargeListParams) *V1List[*Charge] {
	_ = "STUB: not implemented"
	return nil
}

// Search for charges you've previously created using Stripe's [Search Query Language](https://docs.stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
func (c v1ChargeService) Search(ctx context.Context, params *ChargeSearchParams) *V1SearchList[*Charge] {
	_ = "STUB: not implemented"
	return nil
}
