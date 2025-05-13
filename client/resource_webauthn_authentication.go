package client

// Type for resource Webauthn.Authentication
type ResourceWebauthnAuthentication struct {
	// Pointer to client
	Client *Client

	// Action Webauthn.Authentication#Begin
	Begin *ActionWebauthnAuthenticationBegin
	// Action Webauthn.Authentication#Finish
	Finish *ActionWebauthnAuthenticationFinish
}

func NewResourceWebauthnAuthentication(client *Client) *ResourceWebauthnAuthentication {
	actionBegin := NewActionWebauthnAuthenticationBegin(client)
	actionFinish := NewActionWebauthnAuthenticationFinish(client)

	return &ResourceWebauthnAuthentication{
		Client: client,
		Begin:  actionBegin,
		Finish: actionFinish,
	}
}
