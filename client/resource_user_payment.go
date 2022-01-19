package client

// Type for resource User_payment
type ResourceUserPayment struct {
	// Pointer to client
	Client *Client

	// Action User_payment#Create
	Create *ActionUserPaymentCreate
	// Action User_payment#Create
	New *ActionUserPaymentCreate
	// Action User_payment#Index
	Index *ActionUserPaymentIndex
	// Action User_payment#Index
	List *ActionUserPaymentIndex
	// Action User_payment#Show
	Show *ActionUserPaymentShow
	// Action User_payment#Show
	Find *ActionUserPaymentShow
}

func NewResourceUserPayment(client *Client) *ResourceUserPayment {
	actionCreate := NewActionUserPaymentCreate(client)
	actionIndex := NewActionUserPaymentIndex(client)
	actionShow := NewActionUserPaymentShow(client)

	return &ResourceUserPayment{
		Client: client,
		Create: actionCreate,
		New:    actionCreate,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
