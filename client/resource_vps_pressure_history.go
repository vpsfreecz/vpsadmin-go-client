package client

// Type for resource Vps_pressure_history
type ResourceVpsPressureHistory struct {
	// Pointer to client
	Client *Client

	// Action Vps_pressure_history#Index
	Index *ActionVpsPressureHistoryIndex
}

func NewResourceVpsPressureHistory(client *Client) *ResourceVpsPressureHistory {
	actionIndex := NewActionVpsPressureHistoryIndex(client)

	return &ResourceVpsPressureHistory{
		Client: client,
		Index:  actionIndex,
	}
}
