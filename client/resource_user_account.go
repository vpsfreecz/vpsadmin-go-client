package client

// Type for resource User_account
type ResourceUserAccount struct {
	// Pointer to client
	Client *Client

	// Action User_account#Index
	Index *ActionUserAccountIndex
	// Action User_account#Index
	List *ActionUserAccountIndex
	// Action User_account#Show
	Show *ActionUserAccountShow
	// Action User_account#Show
	Find *ActionUserAccountShow
	// Action User_account#Update
	Update *ActionUserAccountUpdate
}

func NewResourceUserAccount(client *Client) *ResourceUserAccount {
	actionIndex := NewActionUserAccountIndex(client)
	actionShow := NewActionUserAccountShow(client)
	actionUpdate := NewActionUserAccountUpdate(client)

	return &ResourceUserAccount{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
