package client

// Type for resource Webauthn
type ResourceWebauthn struct {
	// Pointer to client
	Client *Client

	// Resource Webauthn.Authentication
	Authentication *ResourceWebauthnAuthentication
	// Resource Webauthn.Registration
	Registration *ResourceWebauthnRegistration
}

func NewResourceWebauthn(client *Client) *ResourceWebauthn {

	return &ResourceWebauthn{
		Client:         client,
		Authentication: NewResourceWebauthnAuthentication(client),
		Registration:   NewResourceWebauthnRegistration(client),
	}
}
