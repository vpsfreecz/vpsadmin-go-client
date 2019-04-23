package client

// Type for resource User_request
type ResourceUserRequest struct {
	// Pointer to client
	Client *Client

	// Resource User_request.Registration
	Registration *ResourceUserRequestRegistration
	// Resource User_request.Change
	Change *ResourceUserRequestChange
}

func NewResourceUserRequest(client *Client) *ResourceUserRequest {

	return &ResourceUserRequest{
		Client: client,
		Registration: NewResourceUserRequestRegistration(client),
		Change: NewResourceUserRequestChange(client),
	}
}
