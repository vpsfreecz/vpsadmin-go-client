package client

// Type for resource Oom_report.Task
type ResourceOomReportTask struct {
	// Pointer to client
	Client *Client

	// Action Oom_report.Task#Index
	Index *ActionOomReportTaskIndex
	// Action Oom_report.Task#Index
	List *ActionOomReportTaskIndex
	// Action Oom_report.Task#Show
	Show *ActionOomReportTaskShow
	// Action Oom_report.Task#Show
	Find *ActionOomReportTaskShow
}

func NewResourceOomReportTask(client *Client) *ResourceOomReportTask {
	actionIndex := NewActionOomReportTaskIndex(client)
	actionShow := NewActionOomReportTaskShow(client)

	return &ResourceOomReportTask{
		Client: client,
		Index:  actionIndex,
		List:   actionIndex,
		Show:   actionShow,
		Find:   actionShow,
	}
}
