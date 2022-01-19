package client

// Type for resource User.Totp_device
type ResourceUserTotpDevice struct {
	// Pointer to client
	Client *Client

	// Action User.Totp_device#Confirm
	Confirm *ActionUserTotpDeviceConfirm
	// Action User.Totp_device#Create
	Create *ActionUserTotpDeviceCreate
	// Action User.Totp_device#Create
	New *ActionUserTotpDeviceCreate
	// Action User.Totp_device#Delete
	Delete *ActionUserTotpDeviceDelete
	// Action User.Totp_device#Delete
	Destroy *ActionUserTotpDeviceDelete
	// Action User.Totp_device#Index
	Index *ActionUserTotpDeviceIndex
	// Action User.Totp_device#Index
	List *ActionUserTotpDeviceIndex
	// Action User.Totp_device#Show
	Show *ActionUserTotpDeviceShow
	// Action User.Totp_device#Show
	Find *ActionUserTotpDeviceShow
	// Action User.Totp_device#Update
	Update *ActionUserTotpDeviceUpdate
}

func NewResourceUserTotpDevice(client *Client) *ResourceUserTotpDevice {
	actionConfirm := NewActionUserTotpDeviceConfirm(client)
	actionCreate := NewActionUserTotpDeviceCreate(client)
	actionDelete := NewActionUserTotpDeviceDelete(client)
	actionIndex := NewActionUserTotpDeviceIndex(client)
	actionShow := NewActionUserTotpDeviceShow(client)
	actionUpdate := NewActionUserTotpDeviceUpdate(client)

	return &ResourceUserTotpDevice{
		Client: client,
		Confirm: actionConfirm,
		Create: actionCreate,
		New: actionCreate,
		Delete: actionDelete,
		Destroy: actionDelete,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
		Update: actionUpdate,
	}
}
