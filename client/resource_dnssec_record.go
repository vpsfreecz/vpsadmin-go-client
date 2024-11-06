package client

// Type for resource Dnssec_record
type ResourceDnssecRecord struct {
	// Pointer to client
	Client *Client

	// Action Dnssec_record#Index
	Index *ActionDnssecRecordIndex
	// Action Dnssec_record#Index
	List *ActionDnssecRecordIndex
	// Action Dnssec_record#Show
	Show *ActionDnssecRecordShow
	// Action Dnssec_record#Show
	Find *ActionDnssecRecordShow
}

func NewResourceDnssecRecord(client *Client) *ResourceDnssecRecord {
	actionIndex := NewActionDnssecRecordIndex(client)
	actionShow := NewActionDnssecRecordShow(client)

	return &ResourceDnssecRecord{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
