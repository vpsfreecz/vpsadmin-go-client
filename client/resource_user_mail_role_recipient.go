package client

// Type for resource User.Mail_role_recipient
type ResourceUserMailRoleRecipient struct {
	// Pointer to client
	Client *Client

	// Action User.Mail_role_recipient#Index
	Index *ActionUserMailRoleRecipientIndex
	// Action User.Mail_role_recipient#Index
	List *ActionUserMailRoleRecipientIndex
	// Action User.Mail_role_recipient#Show
	Show *ActionUserMailRoleRecipientShow
	// Action User.Mail_role_recipient#Show
	Find *ActionUserMailRoleRecipientShow
	// Action User.Mail_role_recipient#Update
	Update *ActionUserMailRoleRecipientUpdate
}

func NewResourceUserMailRoleRecipient(client *Client) *ResourceUserMailRoleRecipient {
	actionIndex := NewActionUserMailRoleRecipientIndex(client)
	actionShow := NewActionUserMailRoleRecipientShow(client)
	actionUpdate := NewActionUserMailRoleRecipientUpdate(client)

	return &ResourceUserMailRoleRecipient{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
