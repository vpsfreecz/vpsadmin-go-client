package client

// Type for resource User.Known_device
type ResourceUserKnownDevice struct {
	// Pointer to client
	Client *Client

	// Action User.Known_device#Delete
	Delete *ActionUserKnownDeviceDelete
	// Action User.Known_device#Delete
	Destroy *ActionUserKnownDeviceDelete
	// Action User.Known_device#Index
	Index *ActionUserKnownDeviceIndex
	// Action User.Known_device#Index
	List *ActionUserKnownDeviceIndex
	// Action User.Known_device#Show
	Show *ActionUserKnownDeviceShow
	// Action User.Known_device#Show
	Find *ActionUserKnownDeviceShow
}

func NewResourceUserKnownDevice(client *Client) *ResourceUserKnownDevice {
	actionDelete := NewActionUserKnownDeviceDelete(client)
	actionIndex := NewActionUserKnownDeviceIndex(client)
	actionShow := NewActionUserKnownDeviceShow(client)

	return &ResourceUserKnownDevice{
		Client:  client,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
	}
}
