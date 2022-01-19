package client

import (
	"strings"
)

// ActionVpsSwapWith is a type for action Vps#Swap_with
type ActionVpsSwapWith struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsSwapWith(client *Client) *ActionVpsSwapWith {
	return &ActionVpsSwapWith{
		Client: client,
	}
}

// ActionVpsSwapWithMetaGlobalInput is a type for action global meta input parameters
type ActionVpsSwapWithMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsSwapWithMetaGlobalInput) SetIncludes(value string) *ActionVpsSwapWithMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsSwapWithMetaGlobalInput) SetNo(value bool) *ActionVpsSwapWithMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsSwapWithMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsSwapWithMetaGlobalInput) SelectParameters(params ...string) *ActionVpsSwapWithMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsSwapWithMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsSwapWithInput is a type for action input parameters
type ActionVpsSwapWithInput struct {
	Configs bool `json:"configs"`
	Expirations bool `json:"expirations"`
	Hostname bool `json:"hostname"`
	Resources bool `json:"resources"`
	Vps int64 `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetConfigs sets parameter Configs to value and selects it for sending
func (in *ActionVpsSwapWithInput) SetConfigs(value bool) *ActionVpsSwapWithInput {
	in.Configs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Configs"] = nil
	return in
}
// SetExpirations sets parameter Expirations to value and selects it for sending
func (in *ActionVpsSwapWithInput) SetExpirations(value bool) *ActionVpsSwapWithInput {
	in.Expirations = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Expirations"] = nil
	return in
}
// SetHostname sets parameter Hostname to value and selects it for sending
func (in *ActionVpsSwapWithInput) SetHostname(value bool) *ActionVpsSwapWithInput {
	in.Hostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Hostname"] = nil
	return in
}
// SetResources sets parameter Resources to value and selects it for sending
func (in *ActionVpsSwapWithInput) SetResources(value bool) *ActionVpsSwapWithInput {
	in.Resources = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Resources"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionVpsSwapWithInput) SetVps(value int64) *ActionVpsSwapWithInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsSwapWithInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsSwapWithInput) SelectParameters(params ...string) *ActionVpsSwapWithInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsSwapWithInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsSwapWithRequest is a type for the entire action request
type ActionVpsSwapWithRequest struct {
	Vps map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}


// ActionVpsSwapWithMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsSwapWithMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsSwapWithResponse struct {
	Action *ActionVpsSwapWith `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsSwapWithMetaGlobalOutput `json:"_meta"`
	}
}


// Prepare the action for invocation
func (action *ActionVpsSwapWith) Prepare() *ActionVpsSwapWithInvocation {
	return &ActionVpsSwapWithInvocation{
		Action: action,
		Path: "/v6.0/vpses/{vps_id}/swap_with",
	}
}

// ActionVpsSwapWithInvocation is used to configure action for invocation
type ActionVpsSwapWithInvocation struct {
	// Pointer to the action
	Action *ActionVpsSwapWith

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsSwapWithInput
	// Global meta input parameters
	MetaInput *ActionVpsSwapWithMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsSwapWithInvocation) SetPathParamInt(param string, value int64) *ActionVpsSwapWithInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsSwapWithInvocation) SetPathParamString(param string, value string) *ActionVpsSwapWithInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsSwapWithInvocation) NewInput() *ActionVpsSwapWithInput {
	inv.Input = &ActionVpsSwapWithInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsSwapWithInvocation) SetInput(input *ActionVpsSwapWithInput) *ActionVpsSwapWithInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsSwapWithInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsSwapWithInvocation) NewMetaInput() *ActionVpsSwapWithMetaGlobalInput {
	inv.MetaInput = &ActionVpsSwapWithMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsSwapWithInvocation) SetMetaInput(input *ActionVpsSwapWithMetaGlobalInput) *ActionVpsSwapWithInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsSwapWithInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsSwapWithInvocation) Call() (*ActionVpsSwapWithResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsSwapWithInvocation) callAsBody() (*ActionVpsSwapWithResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsSwapWithResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsSwapWithResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsSwapWithResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsSwapWithResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsSwapWithResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsSwapWithResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsSwapWithInvocation) makeAllInputParams() *ActionVpsSwapWithRequest {
	return &ActionVpsSwapWithRequest{
		Vps: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsSwapWithInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Configs") {
			ret["configs"] = inv.Input.Configs
		}
		if inv.IsParameterSelected("Expirations") {
			ret["expirations"] = inv.Input.Expirations
		}
		if inv.IsParameterSelected("Hostname") {
			ret["hostname"] = inv.Input.Hostname
		}
		if inv.IsParameterSelected("Resources") {
			ret["resources"] = inv.Input.Resources
		}
		if inv.IsParameterSelected("Vps") {
			ret["vps"] = inv.Input.Vps
		}
	}

	return ret
}

func (inv *ActionVpsSwapWithInvocation) makeMetaInputParams() map[string]interface{} {
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
