package stripe

import (
	"errors"
	"time"
)

//
// Public constants
//

const (
	// DefaultTolerance indicates that signatures older than this will be rejected by ConstructEvent.
	WebhookDefaultTolerance time.Duration = 300 * time.Second
	// signingVersion represents the version of the signature we currently use.
	signingVersion string = "v1"
)

//
// Public variables
//

// This block represents the list of errors that could be raised when using the webhook package.
var (
	ErrWebhookInvalidHeader    = errors.New("webhook has invalid Stripe-Signature header")
	ErrWebhookNoValidSignature = errors.New("webhook had no valid signature")
	ErrWebhookNotSigned        = errors.New("webhook has no Stripe-Signature header")
	ErrWebhookTooOld           = errors.New("timestamp wasn't within tolerance")
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
// APIVersion constant.
func ConstructEvent(payload []byte, header string, secret string, opts ...WebhookOption) (Event, error) {
	_ = "STUB: not implemented"
	return *new(Event), nil
}

// ValidatePayload validates the payload against the Stripe-Signature header
// using the specified signing secret. Returns an error if the body or
// Stripe-Signature header provided are unreadable, if the signature doesn't
// match, or if the timestamp for the signature is older than DefaultTolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ValidatePayload(payload []byte, header string, secret string, opts ...WebhookOption) error {
	_ = "STUB: not implemented"
	return nil
}

type WebhookOption func(*webhookConfig)

// WithTolerance validates event timestamps using a custom Tolerance window. If this is
// not set and `IgnoreTolerance` is false, will default to
// `WebhookDefaultTolerance`.
func WithTolerance(tolerance time.Duration) WebhookOption {
	_ = "STUB: not implemented"
	return *new(WebhookOption)
}

// WithIgnoreTolerance will ignore the the event signature's timestamp.
func WithIgnoreTolerance() WebhookOption { _ = "STUB: not implemented"; return *new(WebhookOption) }

// WithIgnoreAPIVersionMismatch will ignore validating whether an event's API version
// matches the stripe-go API version. This is currently only used for ConstructEvent.
func WithIgnoreAPIVersionMismatch() WebhookOption {
	_ = "STUB: not implemented"
	return *new(WebhookOption)
}

type webhookConfig struct {
	Tolerance                time.Duration
	IgnoreTolerance          bool
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

func isCompatibleAPIVersion(sdkAPIVersion, eventAPIVersion string) bool {
	_ = "STUB: not implemented"
	// If the event api version is from before we started adding
	// a release train, there's no way its compatible with this
	// version
	return false
}

// if the SDK is pinned to a preview version, the event's API version must match exactly

// versions are yyyy-MM-dd.train

func constructEvent(payload []byte, sigHeader string, secret string, cfg webhookConfig) (Event, error) {
	_ = "STUB: not implemented"
	return *new(Event), nil
}

func checkEventNotification(payload []byte) error { _ = "STUB: not implemented"; return nil }

func parseSignatureHeader(header string) (*signedHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Signed header looks like "t=1495999758,v1=ABC,v1=DEF,v0=GHI"

// Ignore invalid signatures

// Ignore unknown parts of the header

func validatePayload(payload []byte, sigHeader string, secret string, cfg webhookConfig) error {
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
