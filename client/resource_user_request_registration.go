package client

// Type for resource User_request.Registration
type ResourceUserRequestRegistration struct {
	// Pointer to client
	Client *Client

	// Action User_request.Registration#Create
	Create *ActionUserRequestRegistrationCreate
	// Action User_request.Registration#Create
	New *ActionUserRequestRegistrationCreate
	// Action User_request.Registration#Index
	Index *ActionUserRequestRegistrationIndex
	// Action User_request.Registration#Index
	List *ActionUserRequestRegistrationIndex
	// Action User_request.Registration#Preview
	Preview *ActionUserRequestRegistrationPreview
	// Action User_request.Registration#Resolve
	Resolve *ActionUserRequestRegistrationResolve
	// Action User_request.Registration#Show
	Show *ActionUserRequestRegistrationShow
	// Action User_request.Registration#Show
	Find *ActionUserRequestRegistrationShow
	// Action User_request.Registration#Update
	Update *ActionUserRequestRegistrationUpdate
}

func NewResourceUserRequestRegistration(client *Client) *ResourceUserRequestRegistration {
	actionCreate := NewActionUserRequestRegistrationCreate(client)
	actionIndex := NewActionUserRequestRegistrationIndex(client)
	actionPreview := NewActionUserRequestRegistrationPreview(client)
	actionResolve := NewActionUserRequestRegistrationResolve(client)
	actionShow := NewActionUserRequestRegistrationShow(client)
	actionUpdate := NewActionUserRequestRegistrationUpdate(client)

	return &ResourceUserRequestRegistration{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Index:   actionIndex,
		List:    actionIndex,
		Preview: actionPreview,
		Resolve: actionResolve,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
