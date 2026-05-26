package stripe

import (
	"context"
)

// oauthService is used to invoke /oauth and related APIs.
type oauthService struct {
	B   Backend
	Key string
}

// AuthorizeURL builds an OAuth authorize URL.
func (c oauthService) AuthorizeURL(params *AuthorizeURLParams) string {
	_ = "STUB: not implemented"
	return ""
}

// Create creates an OAuth token using a code after successful redirection back.
func (c oauthService) Create(ctx context.Context, params *OAuthTokenParams) (*OAuthToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete deauthorizes a connected account.
func (c oauthService) Delete(ctx context.Context, params *DeauthorizeParams) (*Deauthorize, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
