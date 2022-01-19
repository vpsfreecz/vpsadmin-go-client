package client

import (
	"strings"
)

// ActionMigrationPlanVpsMigrationCreate is a type for action Migration_plan.Vps_migration#Create
type ActionMigrationPlanVpsMigrationCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionMigrationPlanVpsMigrationCreate(client *Client) *ActionMigrationPlanVpsMigrationCreate {
	return &ActionMigrationPlanVpsMigrationCreate{
		Client: client,
	}
}

// ActionMigrationPlanVpsMigrationCreateMetaGlobalInput is a type for action global meta input parameters
type ActionMigrationPlanVpsMigrationCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput) SetIncludes(value string) *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanVpsMigrationCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput) SelectParameters(params ...string) *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMigrationPlanVpsMigrationCreateInput is a type for action input parameters
type ActionMigrationPlanVpsMigrationCreateInput struct {
	CleanupData bool `json:"cleanup_data"`
	DstNode int64 `json:"dst_node"`
	MaintenanceWindow bool `json:"maintenance_window"`
	Vps int64 `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCleanupData sets parameter CleanupData to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationCreateInput) SetCleanupData(value bool) *ActionMigrationPlanVpsMigrationCreateInput {
	in.CleanupData = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CleanupData"] = nil
	return in
}
// SetDstNode sets parameter DstNode to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationCreateInput) SetDstNode(value int64) *ActionMigrationPlanVpsMigrationCreateInput {
	in.DstNode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DstNode"] = nil
	return in
}
// SetMaintenanceWindow sets parameter MaintenanceWindow to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationCreateInput) SetMaintenanceWindow(value bool) *ActionMigrationPlanVpsMigrationCreateInput {
	in.MaintenanceWindow = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaintenanceWindow"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionMigrationPlanVpsMigrationCreateInput) SetVps(value int64) *ActionMigrationPlanVpsMigrationCreateInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanVpsMigrationCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanVpsMigrationCreateInput) SelectParameters(params ...string) *ActionMigrationPlanVpsMigrationCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanVpsMigrationCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMigrationPlanVpsMigrationCreateRequest is a type for the entire action request
type ActionMigrationPlanVpsMigrationCreateRequest struct {
	VpsMigration map[string]interface{} `json:"vps_migration"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMigrationPlanVpsMigrationCreateOutput is a type for action output parameters
type ActionMigrationPlanVpsMigrationCreateOutput struct {
	CleanupData bool `json:"cleanup_data"`
	CreatedAt string `json:"created_at"`
	DstNode *ActionNodeShowOutput `json:"dst_node"`
	FinishedAt string `json:"finished_at"`
	Id int64 `json:"id"`
	MaintenanceWindow bool `json:"maintenance_window"`
	SrcNode *ActionNodeShowOutput `json:"src_node"`
	StartedAt string `json:"started_at"`
	State string `json:"state"`
	TransactionChain *ActionTransactionChainShowOutput `json:"transaction_chain"`
	Vps *ActionVpsShowOutput `json:"vps"`
}


// Type for action response, including envelope
type ActionMigrationPlanVpsMigrationCreateResponse struct {
	Action *ActionMigrationPlanVpsMigrationCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsMigration *ActionMigrationPlanVpsMigrationCreateOutput `json:"vps_migration"`
	}

	// Action output without the namespace
	Output *ActionMigrationPlanVpsMigrationCreateOutput
}


// Prepare the action for invocation
func (action *ActionMigrationPlanVpsMigrationCreate) Prepare() *ActionMigrationPlanVpsMigrationCreateInvocation {
	return &ActionMigrationPlanVpsMigrationCreateInvocation{
		Action: action,
		Path: "/v6.0/migration_plans/{migration_plan_id}/vps_migrations",
	}
}

// ActionMigrationPlanVpsMigrationCreateInvocation is used to configure action for invocation
type ActionMigrationPlanVpsMigrationCreateInvocation struct {
	// Pointer to the action
	Action *ActionMigrationPlanVpsMigrationCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMigrationPlanVpsMigrationCreateInput
	// Global meta input parameters
	MetaInput *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) SetPathParamInt(param string, value int64) *ActionMigrationPlanVpsMigrationCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) SetPathParamString(param string, value string) *ActionMigrationPlanVpsMigrationCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) NewInput() *ActionMigrationPlanVpsMigrationCreateInput {
	inv.Input = &ActionMigrationPlanVpsMigrationCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) SetInput(input *ActionMigrationPlanVpsMigrationCreateInput) *ActionMigrationPlanVpsMigrationCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) NewMetaInput() *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput {
	inv.MetaInput = &ActionMigrationPlanVpsMigrationCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) SetMetaInput(input *ActionMigrationPlanVpsMigrationCreateMetaGlobalInput) *ActionMigrationPlanVpsMigrationCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) Call() (*ActionMigrationPlanVpsMigrationCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) callAsBody() (*ActionMigrationPlanVpsMigrationCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMigrationPlanVpsMigrationCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsMigration
	}
	return resp, err
}




func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) makeAllInputParams() *ActionMigrationPlanVpsMigrationCreateRequest {
	return &ActionMigrationPlanVpsMigrationCreateRequest{
		VpsMigration: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("CleanupData") {
			ret["cleanup_data"] = inv.Input.CleanupData
		}
		if inv.IsParameterSelected("DstNode") {
			ret["dst_node"] = inv.Input.DstNode
		}
		if inv.IsParameterSelected("MaintenanceWindow") {
			ret["maintenance_window"] = inv.Input.MaintenanceWindow
		}
		if inv.IsParameterSelected("Vps") {
			ret["vps"] = inv.Input.Vps
		}
	}

	return ret
}

func (inv *ActionMigrationPlanVpsMigrationCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
