package client

// Type for resource Dns_record
type ResourceDnsRecord struct {
	// Pointer to client
	Client *Client

	// Action Dns_record#Create
	Create *ActionDnsRecordCreate
	// Action Dns_record#Create
	New *ActionDnsRecordCreate
	// Action Dns_record#Delete
	Delete *ActionDnsRecordDelete
	// Action Dns_record#Delete
	Destroy *ActionDnsRecordDelete
	// Action Dns_record#Dynamic_update
	DynamicUpdate *ActionDnsRecordDynamicUpdate
	// Action Dns_record#Index
	Index *ActionDnsRecordIndex
	// Action Dns_record#Index
	List *ActionDnsRecordIndex
	// Action Dns_record#Show
	Show *ActionDnsRecordShow
	// Action Dns_record#Show
	Find *ActionDnsRecordShow
	// Action Dns_record#Update
	Update *ActionDnsRecordUpdate
}

func NewResourceDnsRecord(client *Client) *ResourceDnsRecord {
	actionCreate := NewActionDnsRecordCreate(client)
	actionDelete := NewActionDnsRecordDelete(client)
	actionDynamicUpdate := NewActionDnsRecordDynamicUpdate(client)
	actionIndex := NewActionDnsRecordIndex(client)
	actionShow := NewActionDnsRecordShow(client)
	actionUpdate := NewActionDnsRecordUpdate(client)

	return &ResourceDnsRecord{
		Client:        client,
		Create:        actionCreate,
		New:           actionCreate,
		Delete:        actionDelete,
		Destroy:       actionDelete,
		DynamicUpdate: actionDynamicUpdate,
		Index:         actionIndex,
		List:          actionIndex,
		Show:          actionShow,
		Find:          actionShow,
		Update:        actionUpdate,
	}
}
