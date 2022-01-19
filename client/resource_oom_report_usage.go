package client

// Type for resource Oom_report.Usage
type ResourceOomReportUsage struct {
	// Pointer to client
	Client *Client

	// Action Oom_report.Usage#Index
	Index *ActionOomReportUsageIndex
	// Action Oom_report.Usage#Index
	List *ActionOomReportUsageIndex
	// Action Oom_report.Usage#Show
	Show *ActionOomReportUsageShow
	// Action Oom_report.Usage#Show
	Find *ActionOomReportUsageShow
}

func NewResourceOomReportUsage(client *Client) *ResourceOomReportUsage {
	actionIndex := NewActionOomReportUsageIndex(client)
	actionShow := NewActionOomReportUsageShow(client)

	return &ResourceOomReportUsage{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
