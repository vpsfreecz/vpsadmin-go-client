package client

// Type for resource Incoming_payment
type ResourceIncomingPayment struct {
	// Pointer to client
	Client *Client

	// Action Incoming_payment#Index
	Index *ActionIncomingPaymentIndex
	// Action Incoming_payment#Index
	List *ActionIncomingPaymentIndex
	// Action Incoming_payment#Show
	Show *ActionIncomingPaymentShow
	// Action Incoming_payment#Show
	Find *ActionIncomingPaymentShow
	// Action Incoming_payment#Update
	Update *ActionIncomingPaymentUpdate
}

func NewResourceIncomingPayment(client *Client) *ResourceIncomingPayment {
	actionIndex := NewActionIncomingPaymentIndex(client)
	actionShow := NewActionIncomingPaymentShow(client)
	actionUpdate := NewActionIncomingPaymentUpdate(client)

	return &ResourceIncomingPayment{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
		Update: actionUpdate,
	}
}
