package client

// Type for resource Webauthn.Registration
type ResourceWebauthnRegistration struct {
	// Pointer to client
	Client *Client

	// Action Webauthn.Registration#Begin
	Begin *ActionWebauthnRegistrationBegin
	// Action Webauthn.Registration#Finish
	Finish *ActionWebauthnRegistrationFinish
}

func NewResourceWebauthnRegistration(client *Client) *ResourceWebauthnRegistration {
	actionBegin := NewActionWebauthnRegistrationBegin(client)
	actionFinish := NewActionWebauthnRegistrationFinish(client)

	return &ResourceWebauthnRegistration{
		Client: client,
		Begin:  actionBegin,
		Finish: actionFinish,
	}
}
