package client

// Type for resource Token
type AuthTokenResourceToken struct {
	// Pointer to client
	Client *Client

	// Action Token#Request
	Request *AuthTokenActionTokenRequest
	// Action Token#Revoke
	Revoke *AuthTokenActionTokenRevoke
	// Action Token#Renew
	Renew *AuthTokenActionTokenRenew
}

func NewAuthTokenResourceToken(client *Client) *AuthTokenResourceToken {
	actionRequest := NewAuthTokenActionTokenRequest(client)
	actionRevoke := NewAuthTokenActionTokenRevoke(client)
	actionRenew := NewAuthTokenActionTokenRenew(client)

	return &AuthTokenResourceToken{
		Client: client,
		Request: actionRequest,
		Revoke: actionRevoke,
		Renew: actionRenew,
	}
}
