package client

// Type for resource Dns_record_log
type ResourceDnsRecordLog struct {
	// Pointer to client
	Client *Client

	// Action Dns_record_log#Index
	Index *ActionDnsRecordLogIndex
	// Action Dns_record_log#Index
	List *ActionDnsRecordLogIndex
	// Action Dns_record_log#Show
	Show *ActionDnsRecordLogShow
	// Action Dns_record_log#Show
	Find *ActionDnsRecordLogShow
}

func NewResourceDnsRecordLog(client *Client) *ResourceDnsRecordLog {
	actionIndex := NewActionDnsRecordLogIndex(client)
	actionShow := NewActionDnsRecordLogShow(client)

	return &ResourceDnsRecordLog{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
