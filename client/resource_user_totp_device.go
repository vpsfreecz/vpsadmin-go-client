package client

// Type for resource User.Totp_device
type ResourceUserTotpDevice struct {
	// Pointer to client
	Client *Client

	// Action User.Totp_device#Index
	Index *ActionUserTotpDeviceIndex
	// Action User.Totp_device#Index
	List *ActionUserTotpDeviceIndex
	// Action User.Totp_device#Show
	Show *ActionUserTotpDeviceShow
	// Action User.Totp_device#Show
	Find *ActionUserTotpDeviceShow
	// Action User.Totp_device#Create
	Create *ActionUserTotpDeviceCreate
	// Action User.Totp_device#Create
	New *ActionUserTotpDeviceCreate
	// Action User.Totp_device#Confirm
	Confirm *ActionUserTotpDeviceConfirm
	// Action User.Totp_device#Update
	Update *ActionUserTotpDeviceUpdate
	// Action User.Totp_device#Delete
	Delete *ActionUserTotpDeviceDelete
	// Action User.Totp_device#Delete
	Destroy *ActionUserTotpDeviceDelete
}

func NewResourceUserTotpDevice(client *Client) *ResourceUserTotpDevice {
	actionIndex := NewActionUserTotpDeviceIndex(client)
	actionShow := NewActionUserTotpDeviceShow(client)
	actionCreate := NewActionUserTotpDeviceCreate(client)
	actionConfirm := NewActionUserTotpDeviceConfirm(client)
	actionUpdate := NewActionUserTotpDeviceUpdate(client)
	actionDelete := NewActionUserTotpDeviceDelete(client)

	return &ResourceUserTotpDevice{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Create: actionCreate,
		New: actionCreate,
		Confirm: actionConfirm,
		Update: actionUpdate,
		Delete: actionDelete,
		Destroy: actionDelete,
	}
}
