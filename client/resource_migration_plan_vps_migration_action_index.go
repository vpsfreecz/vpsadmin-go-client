package client

import (
	"strings"
)

// ActionMigrationPlanVpsMigrationIndex is a type for action Migration_plan.Vps_migration#Index
type ActionMigrationPlanVpsMigrationIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMigrationPlanVpsMigrationIndex(client *Client) *ActionMigrationPlanVpsMigrationIndex {
	return &ActionMigrationPlanVpsMigrationIndex{
		Client: client,
	}
}

// ActionMigrationPlanVpsMigrationIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMigrationPlanVpsMigrationIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput) SetCount(value bool) *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput) SetIncludes(value string) *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanVpsMigrationIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMigrationPlanVpsMigrationIndexInput is a type for action input parameters
type ActionMigrationPlanVpsMigrationIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	State string `json:"state"`
	SrcNode int64 `json:"src_node"`
	DstNode int64 `json:"dst_node"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetOffset(value int64) *ActionMigrationPlanVpsMigrationIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetLimit(value int64) *ActionMigrationPlanVpsMigrationIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetState sets parameter State to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetState(value string) *ActionMigrationPlanVpsMigrationIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}
// SetSrcNode sets parameter SrcNode to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetSrcNode(value int64) *ActionMigrationPlanVpsMigrationIndexInput {
	in.SrcNode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SrcNode"] = nil
	return in
}
// SetDstNode sets parameter DstNode to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetDstNode(value int64) *ActionMigrationPlanVpsMigrationIndexInput {
	in.DstNode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DstNode"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanVpsMigrationIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanVpsMigrationIndexInput) SelectParameters(params ...string) *ActionMigrationPlanVpsMigrationIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanVpsMigrationIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionMigrationPlanVpsMigrationIndexOutput is a type for action output parameters
type ActionMigrationPlanVpsMigrationIndexOutput struct {
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
type ActionMigrationPlanVpsMigrationIndexResponse struct {
	Action *ActionMigrationPlanVpsMigrationIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsMigrations []*ActionMigrationPlanVpsMigrationIndexOutput `json:"vps_migrations"`
	}

	// Action output without the namespace
	Output []*ActionMigrationPlanVpsMigrationIndexOutput
}


// Prepare the action for invocation
func (action *ActionMigrationPlanVpsMigrationIndex) Prepare() *ActionMigrationPlanVpsMigrationIndexInvocation {
	return &ActionMigrationPlanVpsMigrationIndexInvocation{
		Action: action,
		Path: "/v5.0/migration_plans/:migration_plan_id/vps_migrations",
	}
}

// ActionMigrationPlanVpsMigrationIndexInvocation is used to configure action for invocation
type ActionMigrationPlanVpsMigrationIndexInvocation struct {
	// Pointer to the action
	Action *ActionMigrationPlanVpsMigrationIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMigrationPlanVpsMigrationIndexInput
	// Global meta input parameters
	MetaInput *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) SetPathParamInt(param string, value int64) *ActionMigrationPlanVpsMigrationIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) SetPathParamString(param string, value string) *ActionMigrationPlanVpsMigrationIndexInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) SetInput(input *ActionMigrationPlanVpsMigrationIndexInput) *ActionMigrationPlanVpsMigrationIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) SetMetaInput(input *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput) *ActionMigrationPlanVpsMigrationIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) Call() (*ActionMigrationPlanVpsMigrationIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) callAsQuery() (*ActionMigrationPlanVpsMigrationIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMigrationPlanVpsMigrationIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsMigrations
	}
	return resp, err
}



func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["vps_migration[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["vps_migration[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("State") {
			ret["vps_migration[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("SrcNode") {
			ret["vps_migration[src_node]"] = convertInt64ToString(inv.Input.SrcNode)
		}
		if inv.IsParameterSelected("DstNode") {
			ret["vps_migration[dst_node]"] = convertInt64ToString(inv.Input.DstNode)
		}
	}
}

func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

