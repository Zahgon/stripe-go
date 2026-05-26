// Package oauth provides the OAuth APIs
package oauth

import (
	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /oauth and related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// AuthorizeURL builds an OAuth authorize URL.
func AuthorizeURL(params *stripe.AuthorizeURLParams) string { _ = "STUB: not implemented"; return "" }

// AuthorizeURL builds an OAuth authorize URL.
func (c Client) AuthorizeURL(params *stripe.AuthorizeURLParams) string {
	_ = "STUB: not implemented"
	return ""
}

// New creates an OAuth token using a code after successful redirection back.
func New(params *stripe.OAuthTokenParams) (*stripe.OAuthToken, error) {
	_ = "STUB: not implemented"
	return nil,

		// New creates an OAuth token using a code after successful redirection back.
		nil
}

func (c Client) New(params *stripe.OAuthTokenParams) (*stripe.OAuthToken, error) {
	_ = "STUB: not implemented"
	// client_secret is sent in the post body for this endpoint.
	return nil, nil
}

// Del deauthorizes a connected account.
func Del(params *stripe.DeauthorizeParams) (*stripe.Deauthorize, error) {
	_ = "STUB: not implemented"
	return nil,

		// Del deauthorizes a connected account.
		nil
}

func (c Client) Del(params *stripe.DeauthorizeParams) (*stripe.Deauthorize, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getC() Client { _ = "STUB: not implemented"; return *new(Client) }
