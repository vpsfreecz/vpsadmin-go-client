package client

// Type for resource Token
type AuthTokenResourceToken struct {
	// Pointer to client
	Client *Client

	// Action Token#Renew
	Renew *AuthTokenActionTokenRenew
	// Action Token#Request
	Request *AuthTokenActionTokenRequest
	// Action Token#Revoke
	Revoke *AuthTokenActionTokenRevoke
	// Action Token#Totp
	Totp *AuthTokenActionTokenTotp
}

func NewAuthTokenResourceToken(client *Client) *AuthTokenResourceToken {
	actionRenew := NewAuthTokenActionTokenRenew(client)
	actionRequest := NewAuthTokenActionTokenRequest(client)
	actionRevoke := NewAuthTokenActionTokenRevoke(client)
	actionTotp := NewAuthTokenActionTokenTotp(client)

	return &AuthTokenResourceToken{
		Client: client,
		Renew: actionRenew,
		Request: actionRequest,
		Revoke: actionRevoke,
		Totp: actionTotp,
	}
}
