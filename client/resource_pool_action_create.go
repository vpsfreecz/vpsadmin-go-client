package client

import (
)

// ActionPoolCreate is a type for action Pool#Create
type ActionPoolCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionPoolCreate(client *Client) *ActionPoolCreate {
	return &ActionPoolCreate{
		Client: client,
	}
}

// ActionPoolCreateMetaGlobalInput is a type for action global meta input parameters
type ActionPoolCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionPoolCreateMetaGlobalInput) SetNo(value bool) *ActionPoolCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionPoolCreateMetaGlobalInput) SetIncludes(value string) *ActionPoolCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionPoolCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPoolCreateMetaGlobalInput) SelectParameters(params ...string) *ActionPoolCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPoolCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPoolCreateInput is a type for action input parameters
type ActionPoolCreateInput struct {
	Node int64 `json:"node"`
	Label string `json:"label"`
	Filesystem string `json:"filesystem"`
	Role string `json:"role"`
	RefquotaCheck bool `json:"refquota_check"`
	Atime bool `json:"atime"`
	Compression bool `json:"compression"`
	Recordsize int64 `json:"recordsize"`
	Quota int64 `json:"quota"`
	Refquota int64 `json:"refquota"`
	Relatime bool `json:"relatime"`
	Sync string `json:"sync"`
	Sharenfs string `json:"sharenfs"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionPoolCreateInput) SetNode(value int64) *ActionPoolCreateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionPoolCreateInput) SetLabel(value string) *ActionPoolCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetFilesystem sets parameter Filesystem to value and selects it for sending
func (in *ActionPoolCreateInput) SetFilesystem(value string) *ActionPoolCreateInput {
	in.Filesystem = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Filesystem"] = nil
	return in
}
// SetRole sets parameter Role to value and selects it for sending
func (in *ActionPoolCreateInput) SetRole(value string) *ActionPoolCreateInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}
// SetRefquotaCheck sets parameter RefquotaCheck to value and selects it for sending
func (in *ActionPoolCreateInput) SetRefquotaCheck(value bool) *ActionPoolCreateInput {
	in.RefquotaCheck = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RefquotaCheck"] = nil
	return in
}
// SetAtime sets parameter Atime to value and selects it for sending
func (in *ActionPoolCreateInput) SetAtime(value bool) *ActionPoolCreateInput {
	in.Atime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Atime"] = nil
	return in
}
// SetCompression sets parameter Compression to value and selects it for sending
func (in *ActionPoolCreateInput) SetCompression(value bool) *ActionPoolCreateInput {
	in.Compression = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Compression"] = nil
	return in
}
// SetRecordsize sets parameter Recordsize to value and selects it for sending
func (in *ActionPoolCreateInput) SetRecordsize(value int64) *ActionPoolCreateInput {
	in.Recordsize = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Recordsize"] = nil
	return in
}
// SetQuota sets parameter Quota to value and selects it for sending
func (in *ActionPoolCreateInput) SetQuota(value int64) *ActionPoolCreateInput {
	in.Quota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Quota"] = nil
	return in
}
// SetRefquota sets parameter Refquota to value and selects it for sending
func (in *ActionPoolCreateInput) SetRefquota(value int64) *ActionPoolCreateInput {
	in.Refquota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Refquota"] = nil
	return in
}
// SetRelatime sets parameter Relatime to value and selects it for sending
func (in *ActionPoolCreateInput) SetRelatime(value bool) *ActionPoolCreateInput {
	in.Relatime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Relatime"] = nil
	return in
}
// SetSync sets parameter Sync to value and selects it for sending
func (in *ActionPoolCreateInput) SetSync(value string) *ActionPoolCreateInput {
	in.Sync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sync"] = nil
	return in
}
// SetSharenfs sets parameter Sharenfs to value and selects it for sending
func (in *ActionPoolCreateInput) SetSharenfs(value string) *ActionPoolCreateInput {
	in.Sharenfs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sharenfs"] = nil
	return in
}

// SelectParameters sets parameters from ActionPoolCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPoolCreateInput) SelectParameters(params ...string) *ActionPoolCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPoolCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPoolCreateRequest is a type for the entire action request
type ActionPoolCreateRequest struct {
	Pool map[string]interface{} `json:"pool"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionPoolCreateOutput is a type for action output parameters
type ActionPoolCreateOutput struct {
	Id int64 `json:"id"`
	Node *ActionNodeShowOutput `json:"node"`
	Label string `json:"label"`
	Filesystem string `json:"filesystem"`
	Role string `json:"role"`
	RefquotaCheck bool `json:"refquota_check"`
	Atime bool `json:"atime"`
	Compression bool `json:"compression"`
	Recordsize int64 `json:"recordsize"`
	Quota int64 `json:"quota"`
	Refquota int64 `json:"refquota"`
	Relatime bool `json:"relatime"`
	Sync string `json:"sync"`
	Sharenfs string `json:"sharenfs"`
	Used int64 `json:"used"`
	Referenced int64 `json:"referenced"`
	Avail int64 `json:"avail"`
}

// ActionPoolCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionPoolCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionPoolCreateResponse struct {
	Action *ActionPoolCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Pool *ActionPoolCreateOutput `json:"pool"`
		// Global output metadata
		Meta *ActionPoolCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionPoolCreateOutput
}


// Prepare the action for invocation
func (action *ActionPoolCreate) Prepare() *ActionPoolCreateInvocation {
	return &ActionPoolCreateInvocation{
		Action: action,
		Path: "/v6.0/pools",
	}
}

// ActionPoolCreateInvocation is used to configure action for invocation
type ActionPoolCreateInvocation struct {
	// Pointer to the action
	Action *ActionPoolCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionPoolCreateInput
	// Global meta input parameters
	MetaInput *ActionPoolCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionPoolCreateInvocation) NewInput() *ActionPoolCreateInput {
	inv.Input = &ActionPoolCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionPoolCreateInvocation) SetInput(input *ActionPoolCreateInput) *ActionPoolCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionPoolCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionPoolCreateInvocation) NewMetaInput() *ActionPoolCreateMetaGlobalInput {
	inv.MetaInput = &ActionPoolCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionPoolCreateInvocation) SetMetaInput(input *ActionPoolCreateMetaGlobalInput) *ActionPoolCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionPoolCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionPoolCreateInvocation) Call() (*ActionPoolCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionPoolCreateInvocation) callAsBody() (*ActionPoolCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionPoolCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Pool
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionPoolCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionPoolCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionPoolCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionPoolCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionPoolCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionPoolCreateInvocation) makeAllInputParams() *ActionPoolCreateRequest {
	return &ActionPoolCreateRequest{
		Pool: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionPoolCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Node") {
			ret["node"] = inv.Input.Node
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Filesystem") {
			ret["filesystem"] = inv.Input.Filesystem
		}
		if inv.IsParameterSelected("Role") {
			ret["role"] = inv.Input.Role
		}
		if inv.IsParameterSelected("RefquotaCheck") {
			ret["refquota_check"] = inv.Input.RefquotaCheck
		}
		if inv.IsParameterSelected("Atime") {
			ret["atime"] = inv.Input.Atime
		}
		if inv.IsParameterSelected("Compression") {
			ret["compression"] = inv.Input.Compression
		}
		if inv.IsParameterSelected("Recordsize") {
			ret["recordsize"] = inv.Input.Recordsize
		}
		if inv.IsParameterSelected("Quota") {
			ret["quota"] = inv.Input.Quota
		}
		if inv.IsParameterSelected("Refquota") {
			ret["refquota"] = inv.Input.Refquota
		}
		if inv.IsParameterSelected("Relatime") {
			ret["relatime"] = inv.Input.Relatime
		}
		if inv.IsParameterSelected("Sync") {
			ret["sync"] = inv.Input.Sync
		}
		if inv.IsParameterSelected("Sharenfs") {
			ret["sharenfs"] = inv.Input.Sharenfs
		}
	}

	return ret
}

func (inv *ActionPoolCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
