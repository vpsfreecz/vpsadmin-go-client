package client

// Type for resource Oom_report_rule
type ResourceOomReportRule struct {
	// Pointer to client
	Client *Client

	// Action Oom_report_rule#Create
	Create *ActionOomReportRuleCreate
	// Action Oom_report_rule#Create
	New *ActionOomReportRuleCreate
	// Action Oom_report_rule#Delete
	Delete *ActionOomReportRuleDelete
	// Action Oom_report_rule#Delete
	Destroy *ActionOomReportRuleDelete
	// Action Oom_report_rule#Index
	Index *ActionOomReportRuleIndex
	// Action Oom_report_rule#Index
	List *ActionOomReportRuleIndex
	// Action Oom_report_rule#Show
	Show *ActionOomReportRuleShow
	// Action Oom_report_rule#Show
	Find *ActionOomReportRuleShow
	// Action Oom_report_rule#Update
	Update *ActionOomReportRuleUpdate
}

func NewResourceOomReportRule(client *Client) *ResourceOomReportRule {
	actionCreate := NewActionOomReportRuleCreate(client)
	actionDelete := NewActionOomReportRuleDelete(client)
	actionIndex := NewActionOomReportRuleIndex(client)
	actionShow := NewActionOomReportRuleShow(client)
	actionUpdate := NewActionOomReportRuleUpdate(client)

	return &ResourceOomReportRule{
		Client:  client,
		Create:  actionCreate,
		New:     actionCreate,
		Delete:  actionDelete,
		Destroy: actionDelete,
		Index:   actionIndex,
		List:    actionIndex,
		Show:    actionShow,
		Find:    actionShow,
		Update:  actionUpdate,
	}
}
