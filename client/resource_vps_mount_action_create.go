package client

import (
	"strings"
)

// ActionVpsMountCreate is a type for action Vps.Mount#Create
type ActionVpsMountCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMountCreate(client *Client) *ActionVpsMountCreate {
	return &ActionVpsMountCreate{
		Client: client,
	}
}

// ActionVpsMountCreateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMountCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMountCreateMetaGlobalInput) SetNo(value bool) *ActionVpsMountCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMountCreateMetaGlobalInput) SetIncludes(value string) *ActionVpsMountCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMountCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountCreateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMountCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMountCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMountCreateInput is a type for action input parameters
type ActionVpsMountCreateInput struct {
	Dataset int64 `json:"dataset"`
	Snapshot int64 `json:"snapshot"`
	UserNamespaceMap int64 `json:"user_namespace_map"`
	Mountpoint string `json:"mountpoint"`
	Mode string `json:"mode"`
	OnStartFail string `json:"on_start_fail"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetDataset sets parameter Dataset to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetDataset(value int64) *ActionVpsMountCreateInput {
	in.Dataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Dataset"] = nil
	return in
}
// SetSnapshot sets parameter Snapshot to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetSnapshot(value int64) *ActionVpsMountCreateInput {
	in.Snapshot = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Snapshot"] = nil
	return in
}
// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetUserNamespaceMap(value int64) *ActionVpsMountCreateInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserNamespaceMap"] = nil
	return in
}
// SetMountpoint sets parameter Mountpoint to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetMountpoint(value string) *ActionVpsMountCreateInput {
	in.Mountpoint = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Mountpoint"] = nil
	return in
}
// SetMode sets parameter Mode to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetMode(value string) *ActionVpsMountCreateInput {
	in.Mode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Mode"] = nil
	return in
}
// SetOnStartFail sets parameter OnStartFail to value and selects it for sending
func (in *ActionVpsMountCreateInput) SetOnStartFail(value string) *ActionVpsMountCreateInput {
	in.OnStartFail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OnStartFail"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMountCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountCreateInput) SelectParameters(params ...string) *ActionVpsMountCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMountCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMountCreateRequest is a type for the entire action request
type ActionVpsMountCreateRequest struct {
	Mount map[string]interface{} `json:"mount"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsMountCreateOutput is a type for action output parameters
type ActionVpsMountCreateOutput struct {
	Id int64 `json:"id"`
	Vps *ActionVpsShowOutput `json:"vps"`
	Dataset *ActionDatasetShowOutput `json:"dataset"`
	Snapshot *ActionDatasetSnapshotShowOutput `json:"snapshot"`
	UserNamespaceMap *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
	Mountpoint string `json:"mountpoint"`
	Mode string `json:"mode"`
	OnStartFail string `json:"on_start_fail"`
	ExpirationDate string `json:"expiration_date"`
	Enabled bool `json:"enabled"`
	MasterEnabled bool `json:"master_enabled"`
	CurrentState string `json:"current_state"`
}

// ActionVpsMountCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsMountCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsMountCreateResponse struct {
	Action *ActionVpsMountCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Mount *ActionVpsMountCreateOutput `json:"mount"`
		// Global output metadata
		Meta *ActionVpsMountCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionVpsMountCreateOutput
}


// Prepare the action for invocation
func (action *ActionVpsMountCreate) Prepare() *ActionVpsMountCreateInvocation {
	return &ActionVpsMountCreateInvocation{
		Action: action,
		Path: "/v5.0/vpses/:vps_id/mounts",
	}
}

// ActionVpsMountCreateInvocation is used to configure action for invocation
type ActionVpsMountCreateInvocation struct {
	// Pointer to the action
	Action *ActionVpsMountCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsMountCreateInput
	// Global meta input parameters
	MetaInput *ActionVpsMountCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMountCreateInvocation) SetPathParamInt(param string, value int64) *ActionVpsMountCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMountCreateInvocation) SetPathParamString(param string, value string) *ActionVpsMountCreateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsMountCreateInvocation) NewInput() *ActionVpsMountCreateInput {
	inv.Input = &ActionVpsMountCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsMountCreateInvocation) SetInput(input *ActionVpsMountCreateInput) *ActionVpsMountCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsMountCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMountCreateInvocation) NewMetaInput() *ActionVpsMountCreateMetaGlobalInput {
	inv.MetaInput = &ActionVpsMountCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMountCreateInvocation) SetMetaInput(input *ActionVpsMountCreateMetaGlobalInput) *ActionVpsMountCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMountCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMountCreateInvocation) Call() (*ActionVpsMountCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsMountCreateInvocation) callAsBody() (*ActionVpsMountCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsMountCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Mount
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsMountCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsMountCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsMountCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsMountCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)
	input.SetUpdateIn(updateIn)

	pollResp, err := req.Call()

	if err != nil {
		return pollResp, err
	} else if pollResp.Output.Finished {
		return pollResp, nil
	}

	if callback(pollResp.Output) == StopWatching {
		return pollResp, nil
	}

	for {
		req = resp.Action.Client.ActionState.Poll.Prepare()
		req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
		req.SetInput(&ActionActionStatePollInput{
			Timeout: timeout,
			UpdateIn: updateIn,
			Status: pollResp.Output.Status,
			Current: pollResp.Output.Current,
			Total: pollResp.Output.Total,
		})
		pollResp, err = req.Call()

		if err != nil {
			return pollResp, err
		} else if pollResp.Output.Finished {
			return pollResp, nil
		}

		if callback(pollResp.Output) == StopWatching {
			return pollResp, nil
		}
	}
}

// CancelOperation cancels the current blocking operation
func (resp *ActionVpsMountCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsMountCreateInvocation) makeAllInputParams() *ActionVpsMountCreateRequest {
	return &ActionVpsMountCreateRequest{
		Mount: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsMountCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Dataset") {
			ret["dataset"] = inv.Input.Dataset
		}
		if inv.IsParameterSelected("Snapshot") {
			ret["snapshot"] = inv.Input.Snapshot
		}
		if inv.IsParameterSelected("UserNamespaceMap") {
			ret["user_namespace_map"] = inv.Input.UserNamespaceMap
		}
		if inv.IsParameterSelected("Mountpoint") {
			ret["mountpoint"] = inv.Input.Mountpoint
		}
		if inv.IsParameterSelected("Mode") {
			ret["mode"] = inv.Input.Mode
		}
		if inv.IsParameterSelected("OnStartFail") {
			ret["on_start_fail"] = inv.Input.OnStartFail
		}
	}

	return ret
}

func (inv *ActionVpsMountCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
