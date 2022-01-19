package client

// Type for resource Api_server
type ResourceApiServer struct {
	// Pointer to client
	Client *Client

	// Action Api_server#Unlock_transaction_signing_key
	UnlockTransactionSigningKey *ActionApiServerUnlockTransactionSigningKey
}

func NewResourceApiServer(client *Client) *ResourceApiServer {
	actionUnlockTransactionSigningKey := NewActionApiServerUnlockTransactionSigningKey(client)

	return &ResourceApiServer{
		Client:                      client,
		UnlockTransactionSigningKey: actionUnlockTransactionSigningKey,
	}
}
