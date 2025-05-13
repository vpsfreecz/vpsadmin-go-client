package client

// Type for resource User.Webauthn_credential
type ResourceUserWebauthnCredential struct {
	// Pointer to client
	Client *Client

	// Action User.Webauthn_credential#Delete
	Delete *ActionUserWebauthnCredentialDelete
	// Action User.Webauthn_credential#Delete
	Destroy *ActionUserWebauthnCredentialDelete
	// Action User.Webauthn_credential#Index
	Index *ActionUserWebauthnCredentialIndex
	// Action User.Webauthn_credential#Index
	List *ActionUserWebauthnCredentialIndex
	// Action User.Webauthn_credential#Show
	Show *ActionUserWebauthnCredentialShow
	// Action User.Webauthn_credential#Show
	Find *ActionUserWebauthnCredentialShow
	// Action User.Webauthn_credential#Update
	Update *ActionUserWebauthnCredentialUpdate
}

func NewResourceUserWebauthnCredential(client *Client) *ResourceUserWebauthnCredential {
	actionDelete := NewActionUserWebauthnCredentialDelete(client)
	actionIndex := NewActionUserWebauthnCredentialIndex(client)
	actionShow := NewActionUserWebauthnCredentialShow(client)
	actionUpdate := NewActionUserWebauthnCredentialUpdate(client)

	return &ResourceUserWebauthnCredential{
		Client:  client,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
