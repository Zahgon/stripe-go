package webhook

import (
	"errors"
	"time"

	"github.com/stripe/stripe-go/v85"
)

//
// Public constants
//

const (
	// DefaultTolerance indicates that signatures older than this will be rejected by ConstructEvent.
	DefaultTolerance time.Duration = 300 * time.Second
	// signingVersion represents the version of the signature we currently use.
	signingVersion string = "v1"
)

//
// Public variables
//

// This block represents the list of errors that could be raised when using the webhook package.
var (
	ErrInvalidHeader    = errors.New("webhook has invalid Stripe-Signature header")
	ErrNoValidSignature = errors.New("webhook had no valid signature")
	ErrNotSigned        = errors.New("webhook has no Stripe-Signature header")
	ErrTooOld           = errors.New("timestamp wasn't within tolerance")
)

//
// Public functions
//

// ComputeSignature computes a webhook signature using Stripe's v1 signing
// method.
//
// See https://stripe.com/docs/webhooks#signatures for more information.
func ComputeSignature(t time.Time, payload []byte, secret string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// ConstructEvent initializes an Event object from a JSON webhook payload, validating
// the Stripe-Signature header using the specified signing secret. Returns an error
// if the body or Stripe-Signature header provided are unreadable, if the
// signature doesn't match, or if the timestamp for the signature is older than
// DefaultTolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
// This will return an error if the event API version does not match the
// stripe.APIVersion constant.
func ConstructEvent(payload []byte, header string, secret string) (stripe.Event, error) {
	_ = "STUB: not implemented"
	return *new(stripe.Event), nil
}

// ConstructEventIgnoringTolerance initializes an Event object from a JSON webhook
// payload, validating the Stripe-Signature header using the specified signing secret.
// Returns an error if the body or Stripe-Signature header provided are unreadable or
// if the signature doesn't match. Does not check the signature's timestamp.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
// This will return an error if the event API version does not match the
// stripe.APIVersion constant.
func ConstructEventIgnoringTolerance(payload []byte, header string, secret string) (stripe.Event, error) {
	_ = "STUB: not implemented"
	return *new(stripe.Event), nil
}

// ConstructEventWithTolerance initializes an Event object from a JSON webhook payload,
// validating the signature in the Stripe-Signature header using the specified signing
// secret and tolerance window. Returns an error if the body or Stripe-Signature header
// provided are unreadable, if the signature doesn't match, or if the timestamp
// for the signature is older than the specified tolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
// This will return an error if the event API version does not match the
// stripe.APIVersion constant.
func ConstructEventWithTolerance(payload []byte, header string, secret string, tolerance time.Duration) (stripe.Event, error) {
	_ = "STUB: not implemented"
	return *new(stripe.Event), nil
}

// ConstructEventWithOptions initializes an Event object from a JSON webhook payload,
// validating the signature in the Stripe-Signature header using the specified signing
// secret and tolerance window provided by the options, if applicable.
//
// See `ConstructEventOptions` for more details on each of the options.
//
// Returns an error if the signature doesn't match, or:
//   - if `IgnoreTolerance` is false and the timestamp embedded in the event
//     header is not within the tolerance window (similar to `ConstructEventWithTolerance`)
//   - if `IgnoreAPIVersionMismatch` is false and the webhook event API version
//     does not match the API version of the stripe-go library, as defined in
//     `stripe.APIVersion`.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ConstructEventWithOptions(payload []byte, header string, secret string, options ConstructEventOptions) (stripe.Event, error) {
	_ = "STUB: not implemented"
	return *new(stripe.Event), nil
}

// ValidatePayload validates the payload against the Stripe-Signature header
// using the specified signing secret. Returns an error if the body or
// Stripe-Signature header provided are unreadable, if the signature doesn't
// match, or if the timestamp for the signature is older than DefaultTolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ValidatePayload(payload []byte, header string, secret string) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidatePayloadIgnoringTolerance validates the payload against the Stripe-Signature header
// using the specified signing secret. Returns an error if the body or
// Stripe-Signature header provided are unreadable or if the signature doesn't match.
// Does not check the signature's timestamp.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ValidatePayloadIgnoringTolerance(payload []byte, header string, secret string) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidatePayloadWithTolerance validates the payload against the Stripe-Signature header
// using the specified signing secret and tolerance window. Returns an error if the body
// or Stripe-Signature header provided are unreadable, if the signature doesn't match, or
// if the timestamp for the signature is older than the specified tolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ValidatePayloadWithTolerance(payload []byte, header string, secret string, tolerance time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

type ConstructEventOptions struct {
	// Validates event timestamps using a custom Tolerance window. If this is
	// not set and `IgnoreTolerance` is false, will default to
	// `DefaultTolerance`.
	Tolerance time.Duration

	// If set to true, will ignore the `tolerance` option entirely and will not
	// check the event signature's timestamp. Defaults to false. When false,
	// constructing an event will fail with an error if the timestamp is not
	// within the `Tolerance` window.
	IgnoreTolerance bool

	// If set to true, will ignore validating whether an event's API version
	// matches the stripe-go API version. Defaults to false, returning an error
	// when there is a mismatch.
	IgnoreAPIVersionMismatch bool
}

//
// Private types
//

type signedHeader struct {
	timestamp  time.Time
	signatures [][]byte
}

//
// Private functions
//

func isCompatibleAPIVersion(sdkApiVersion, eventApiVersion string) bool {
	_ = "STUB: not implemented"
	// If the event api version is from before we started adding
	// a release train, there's no way its compatible with this
	// version
	return false
}

// if the SDK is pinned to a preview version, the event's API version must match exactly

// versions are yyyy-MM-dd.train

func constructEvent(payload []byte, sigHeader string, secret string, options ConstructEventOptions) (stripe.Event, error) {
	_ = "STUB: not implemented"
	return *new(stripe.Event), nil
}

func checkEventNotification(payload []byte) error { _ = "STUB: not implemented"; return nil }

func parseSignatureHeader(header string) (*signedHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Signed header looks like "t=1495999758,v1=ABC,v1=DEF,v0=GHI"

// Ignore invalid signatures

// Ignore unknown parts of the header

func validatePayload(payload []byte, sigHeader string, secret string, tolerance time.Duration, enforceTolerance bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Check all given v1 signatures, multiple signatures will be sent temporarily in the case of a rolled signature secret

// For mocking webhook events
type UnsignedPayload struct {
	Payload   []byte
	Secret    string
	Timestamp time.Time
	Scheme    string
}

type SignedPayload struct {
	UnsignedPayload

	Signature []byte
	Header    string
}

func GenerateTestSignedPayload(options *UnsignedPayload) *SignedPayload {
	_ = "STUB: not implemented"
	return nil
}

func generateHeader(p SignedPayload) string { _ = "STUB: not implemented"; return "" }
