package client

// Type for resource User.Mail_template_recipient
type ResourceUserMailTemplateRecipient struct {
	// Pointer to client
	Client *Client

	// Action User.Mail_template_recipient#Index
	Index *ActionUserMailTemplateRecipientIndex
	// Action User.Mail_template_recipient#Index
	List *ActionUserMailTemplateRecipientIndex
	// Action User.Mail_template_recipient#Show
	Show *ActionUserMailTemplateRecipientShow
	// Action User.Mail_template_recipient#Show
	Find *ActionUserMailTemplateRecipientShow
	// Action User.Mail_template_recipient#Update
	Update *ActionUserMailTemplateRecipientUpdate
}

func NewResourceUserMailTemplateRecipient(client *Client) *ResourceUserMailTemplateRecipient {
	actionIndex := NewActionUserMailTemplateRecipientIndex(client)
	actionShow := NewActionUserMailTemplateRecipientShow(client)
	actionUpdate := NewActionUserMailTemplateRecipientUpdate(client)

	return &ResourceUserMailTemplateRecipient{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
		Update: actionUpdate,
	}
}
