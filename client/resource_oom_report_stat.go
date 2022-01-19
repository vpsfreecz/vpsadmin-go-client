package client

// Type for resource Oom_report.Stat
type ResourceOomReportStat struct {
	// Pointer to client
	Client *Client

	// Action Oom_report.Stat#Index
	Index *ActionOomReportStatIndex
	// Action Oom_report.Stat#Index
	List *ActionOomReportStatIndex
	// Action Oom_report.Stat#Show
	Show *ActionOomReportStatShow
	// Action Oom_report.Stat#Show
	Find *ActionOomReportStatShow
}

func NewResourceOomReportStat(client *Client) *ResourceOomReportStat {
	actionIndex := NewActionOomReportStatIndex(client)
	actionShow := NewActionOomReportStatShow(client)

	return &ResourceOomReportStat{
		Client: client,
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
