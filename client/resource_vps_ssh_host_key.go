package client

// Type for resource Vps.Ssh_host_key
type ResourceVpsSshHostKey struct {
	// Pointer to client
	Client *Client

	// Action Vps.Ssh_host_key#Index
	Index *ActionVpsSshHostKeyIndex
	// Action Vps.Ssh_host_key#Index
	List *ActionVpsSshHostKeyIndex
	// Action Vps.Ssh_host_key#Show
	Show *ActionVpsSshHostKeyShow
	// Action Vps.Ssh_host_key#Show
	Find *ActionVpsSshHostKeyShow
}

func NewResourceVpsSshHostKey(client *Client) *ResourceVpsSshHostKey {
	actionIndex := NewActionVpsSshHostKeyIndex(client)
	actionShow := NewActionVpsSshHostKeyShow(client)

	return &ResourceVpsSshHostKey{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
