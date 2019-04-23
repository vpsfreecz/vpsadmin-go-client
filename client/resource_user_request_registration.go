package client

// Type for resource User_request.Registration
type ResourceUserRequestRegistration struct {
	// Pointer to client
	Client *Client

	// Action User_request.Registration#Index
	Index *ActionUserRequestRegistrationIndex
	// Action User_request.Registration#Index
	List *ActionUserRequestRegistrationIndex
	// Action User_request.Registration#Show
	Show *ActionUserRequestRegistrationShow
	// Action User_request.Registration#Show
	Find *ActionUserRequestRegistrationShow
	// Action User_request.Registration#Create
	Create *ActionUserRequestRegistrationCreate
	// Action User_request.Registration#Create
	New *ActionUserRequestRegistrationCreate
	// Action User_request.Registration#Resolve
	Resolve *ActionUserRequestRegistrationResolve
	// Action User_request.Registration#Preview
	Preview *ActionUserRequestRegistrationPreview
	// Action User_request.Registration#Update
	Update *ActionUserRequestRegistrationUpdate
}

func NewResourceUserRequestRegistration(client *Client) *ResourceUserRequestRegistration {
	actionIndex := NewActionUserRequestRegistrationIndex(client)
	actionShow := NewActionUserRequestRegistrationShow(client)
	actionCreate := NewActionUserRequestRegistrationCreate(client)
	actionResolve := NewActionUserRequestRegistrationResolve(client)
	actionPreview := NewActionUserRequestRegistrationPreview(client)
	actionUpdate := NewActionUserRequestRegistrationUpdate(client)

	return &ResourceUserRequestRegistration{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Resolve: actionResolve,
		Preview: actionPreview,
		Update: actionUpdate,
	}
}
