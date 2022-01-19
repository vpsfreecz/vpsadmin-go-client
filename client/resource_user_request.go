package client

// Type for resource User_request
type ResourceUserRequest struct {
	// Pointer to client
	Client *Client

	// Resource User_request.Change
	Change *ResourceUserRequestChange
	// Resource User_request.Registration
	Registration *ResourceUserRequestRegistration
}

func NewResourceUserRequest(client *Client) *ResourceUserRequest {

	return &ResourceUserRequest{
		Client:       client,
		Change:       NewResourceUserRequestChange(client),
		Registration: NewResourceUserRequestRegistration(client),
	}
}
