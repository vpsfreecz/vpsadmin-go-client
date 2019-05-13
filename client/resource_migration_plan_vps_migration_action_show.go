package client

import (
	"strings"
)

// ActionMigrationPlanVpsMigrationShow is a type for action Migration_plan.Vps_migration#Show
type ActionMigrationPlanVpsMigrationShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMigrationPlanVpsMigrationShow(client *Client) *ActionMigrationPlanVpsMigrationShow {
	return &ActionMigrationPlanVpsMigrationShow{
		Client: client,
	}
}

// ActionMigrationPlanVpsMigrationShowMetaGlobalInput is a type for action global meta input parameters
type ActionMigrationPlanVpsMigrationShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationShowMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanVpsMigrationShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationShowMetaGlobalInput) SetIncludes(value string) *ActionMigrationPlanVpsMigrationShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanVpsMigrationShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanVpsMigrationShowMetaGlobalInput) SelectParameters(params ...string) *ActionMigrationPlanVpsMigrationShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanVpsMigrationShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionMigrationPlanVpsMigrationShowOutput is a type for action output parameters
type ActionMigrationPlanVpsMigrationShowOutput struct {
	Id int64 `json:"id"`
	State string `json:"state"`
	TransactionChain *ActionTransactionChainShowOutput `json:"transaction_chain"`
	SrcNode *ActionNodeShowOutput `json:"src_node"`
	Vps *ActionVpsShowOutput `json:"vps"`
	DstNode *ActionNodeShowOutput `json:"dst_node"`
	OutageWindow bool `json:"outage_window"`
	CleanupData bool `json:"cleanup_data"`
	CreatedAt string `json:"created_at"`
	StartedAt string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
}


// Type for action response, including envelope
type ActionMigrationPlanVpsMigrationShowResponse struct {
	Action *ActionMigrationPlanVpsMigrationShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsMigration *ActionMigrationPlanVpsMigrationShowOutput `json:"vps_migration"`
	}

	// Action output without the namespace
	Output *ActionMigrationPlanVpsMigrationShowOutput
}


// Prepare the action for invocation
func (action *ActionMigrationPlanVpsMigrationShow) Prepare() *ActionMigrationPlanVpsMigrationShowInvocation {
	return &ActionMigrationPlanVpsMigrationShowInvocation{
		Action: action,
		Path: "/v5.0/migration_plans/{migration_plan_id}/vps_migrations/{vps_migration_id}",
	}
}

// ActionMigrationPlanVpsMigrationShowInvocation is used to configure action for invocation
type ActionMigrationPlanVpsMigrationShowInvocation struct {
	// Pointer to the action
	Action *ActionMigrationPlanVpsMigrationShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMigrationPlanVpsMigrationShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMigrationPlanVpsMigrationShowInvocation) SetPathParamInt(param string, value int64) *ActionMigrationPlanVpsMigrationShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMigrationPlanVpsMigrationShowInvocation) SetPathParamString(param string, value string) *ActionMigrationPlanVpsMigrationShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMigrationPlanVpsMigrationShowInvocation) NewMetaInput() *ActionMigrationPlanVpsMigrationShowMetaGlobalInput {
	inv.MetaInput = &ActionMigrationPlanVpsMigrationShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMigrationPlanVpsMigrationShowInvocation) SetMetaInput(input *ActionMigrationPlanVpsMigrationShowMetaGlobalInput) *ActionMigrationPlanVpsMigrationShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMigrationPlanVpsMigrationShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMigrationPlanVpsMigrationShowInvocation) Call() (*ActionMigrationPlanVpsMigrationShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMigrationPlanVpsMigrationShowInvocation) callAsQuery() (*ActionMigrationPlanVpsMigrationShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMigrationPlanVpsMigrationShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsMigration
	}
	return resp, err
}




func (inv *ActionMigrationPlanVpsMigrationShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

