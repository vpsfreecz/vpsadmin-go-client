package client

import (
)

// ActionVpsConfigCreate is a type for action Vps_config#Create
type ActionVpsConfigCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsConfigCreate(client *Client) *ActionVpsConfigCreate {
	return &ActionVpsConfigCreate{
		Client: client,
	}
}

// ActionVpsConfigCreateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsConfigCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsConfigCreateMetaGlobalInput) SetNo(value bool) *ActionVpsConfigCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsConfigCreateMetaGlobalInput) SetIncludes(value string) *ActionVpsConfigCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConfigCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConfigCreateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsConfigCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConfigCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConfigCreateInput is a type for action input parameters
type ActionVpsConfigCreateInput struct {
	Name string `json:"name"`
	Label string `json:"label"`
	Config string `json:"config"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionVpsConfigCreateInput) SetName(value string) *ActionVpsConfigCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionVpsConfigCreateInput) SetLabel(value string) *ActionVpsConfigCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetConfig sets parameter Config to value and selects it for sending
func (in *ActionVpsConfigCreateInput) SetConfig(value string) *ActionVpsConfigCreateInput {
	in.Config = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Config"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConfigCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConfigCreateInput) SelectParameters(params ...string) *ActionVpsConfigCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConfigCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConfigCreateRequest is a type for the entire action request
type ActionVpsConfigCreateRequest struct {
	VpsConfig map[string]interface{} `json:"vps_config"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsConfigCreateOutput is a type for action output parameters
type ActionVpsConfigCreateOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	Config string `json:"config"`
}

// ActionVpsConfigCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsConfigCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsConfigCreateResponse struct {
	Action *ActionVpsConfigCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsConfig *ActionVpsConfigCreateOutput `json:"vps_config"`
		// Global output metadata
		Meta *ActionVpsConfigCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionVpsConfigCreateOutput
}


// Prepare the action for invocation
func (action *ActionVpsConfigCreate) Prepare() *ActionVpsConfigCreateInvocation {
	return &ActionVpsConfigCreateInvocation{
		Action: action,
		Path: "/v5.0/vps_configs",
	}
}

// ActionVpsConfigCreateInvocation is used to configure action for invocation
type ActionVpsConfigCreateInvocation struct {
	// Pointer to the action
	Action *ActionVpsConfigCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsConfigCreateInput
	// Global meta input parameters
	MetaInput *ActionVpsConfigCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsConfigCreateInvocation) NewInput() *ActionVpsConfigCreateInput {
	inv.Input = &ActionVpsConfigCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsConfigCreateInvocation) SetInput(input *ActionVpsConfigCreateInput) *ActionVpsConfigCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsConfigCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsConfigCreateInvocation) NewMetaInput() *ActionVpsConfigCreateMetaGlobalInput {
	inv.MetaInput = &ActionVpsConfigCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsConfigCreateInvocation) SetMetaInput(input *ActionVpsConfigCreateMetaGlobalInput) *ActionVpsConfigCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsConfigCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsConfigCreateInvocation) Call() (*ActionVpsConfigCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsConfigCreateInvocation) callAsBody() (*ActionVpsConfigCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsConfigCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsConfig
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsConfigCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsConfigCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsConfigCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsConfigCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsConfigCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsConfigCreateInvocation) makeAllInputParams() *ActionVpsConfigCreateRequest {
	return &ActionVpsConfigCreateRequest{
		VpsConfig: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsConfigCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Config") {
			ret["config"] = inv.Input.Config
		}
	}

	return ret
}

func (inv *ActionVpsConfigCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
