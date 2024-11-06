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
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	DstNode int64  `json:"dst_node"`
	FromId  int64  `json:"from_id"`
	Limit   int64  `json:"limit"`
	SrcNode int64  `json:"src_node"`
	State   string `json:"state"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDstNode sets parameter DstNode to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetDstNode(value int64) *ActionMigrationPlanVpsMigrationIndexInput {
	in.DstNode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDstNodeNil(false)
	in._selectedParameters["DstNode"] = nil
	return in
}

// SetDstNodeNil sets parameter DstNode to nil and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetDstNodeNil(set bool) *ActionMigrationPlanVpsMigrationIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["DstNode"] = nil
		in.SelectParameters("DstNode")
	} else {
		delete(in._nilParameters, "DstNode")
	}
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetFromId(value int64) *ActionMigrationPlanVpsMigrationIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
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

// SetSrcNode sets parameter SrcNode to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetSrcNode(value int64) *ActionMigrationPlanVpsMigrationIndexInput {
	in.SrcNode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSrcNodeNil(false)
	in._selectedParameters["SrcNode"] = nil
	return in
}

// SetSrcNodeNil sets parameter SrcNode to nil and selects it for sending
func (in *ActionMigrationPlanVpsMigrationIndexInput) SetSrcNodeNil(set bool) *ActionMigrationPlanVpsMigrationIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["SrcNode"] = nil
		in.SelectParameters("SrcNode")
	} else {
		delete(in._nilParameters, "SrcNode")
	}
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

// UnselectParameters unsets parameters from ActionMigrationPlanVpsMigrationIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMigrationPlanVpsMigrationIndexInput) UnselectParameters(params ...string) *ActionMigrationPlanVpsMigrationIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	CleanupData       bool                              `json:"cleanup_data"`
	CreatedAt         string                            `json:"created_at"`
	DstNode           *ActionNodeShowOutput             `json:"dst_node"`
	FinishedAt        string                            `json:"finished_at"`
	Id                int64                             `json:"id"`
	MaintenanceWindow bool                              `json:"maintenance_window"`
	SrcNode           *ActionNodeShowOutput             `json:"src_node"`
	StartedAt         string                            `json:"started_at"`
	State             string                            `json:"state"`
	TransactionChain  *ActionTransactionChainShowOutput `json:"transaction_chain"`
	Vps               *ActionVpsShowOutput              `json:"vps"`
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
		Path:   "/v7.0/migration_plans/{migration_plan_id}/vps_migrations",
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
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) NewInput() *ActionMigrationPlanVpsMigrationIndexInput {
	inv.Input = &ActionMigrationPlanVpsMigrationIndexInput{}
	return inv.Input
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) NewMetaInput() *ActionMigrationPlanVpsMigrationIndexMetaGlobalInput {
	inv.MetaInput = &ActionMigrationPlanVpsMigrationIndexMetaGlobalInput{}
	return inv.MetaInput
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		if inv.IsParameterSelected("DstNode") {
			ret["vps_migration[dst_node]"] = convertInt64ToString(inv.Input.DstNode)
		}
		if inv.IsParameterSelected("FromId") {
			ret["vps_migration[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["vps_migration[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("SrcNode") {
			ret["vps_migration[src_node]"] = convertInt64ToString(inv.Input.SrcNode)
		}
		if inv.IsParameterSelected("State") {
			ret["vps_migration[state]"] = inv.Input.State
		}
	}
}

func (inv *ActionMigrationPlanVpsMigrationIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
