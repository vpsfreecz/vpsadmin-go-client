package client

// Type for resource Oom_report
type ResourceOomReport struct {
	// Pointer to client
	Client *Client

	// Resource Oom_report.Stat
	Stat *ResourceOomReportStat
	// Resource Oom_report.Task
	Task *ResourceOomReportTask
	// Resource Oom_report.Usage
	Usage *ResourceOomReportUsage
	// Action Oom_report#Index
	Index *ActionOomReportIndex
	// Action Oom_report#Index
	List *ActionOomReportIndex
	// Action Oom_report#Show
	Show *ActionOomReportShow
	// Action Oom_report#Show
	Find *ActionOomReportShow
}

func NewResourceOomReport(client *Client) *ResourceOomReport {
	actionIndex := NewActionOomReportIndex(client)
	actionShow := NewActionOomReportShow(client)

	return &ResourceOomReport{
		Client: client,
		Stat: NewResourceOomReportStat(client),
		Task: NewResourceOomReportTask(client),
		Usage: NewResourceOomReportUsage(client),
		Index: actionIndex,
		List: actionIndex,
		Show: actionShow,
		Find: actionShow,
	}
}
