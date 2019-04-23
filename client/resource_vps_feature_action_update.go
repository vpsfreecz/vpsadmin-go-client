package client

import (
	"strings"
)

// ActionVpsFeatureUpdate is a type for action Vps.Feature#Update
type ActionVpsFeatureUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsFeatureUpdate(client *Client) *ActionVpsFeatureUpdate {
	return &ActionVpsFeatureUpdate{
		Client: client,
	}
}

// ActionVpsFeatureUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsFeatureUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsFeatureUpdateMetaGlobalInput) SetNo(value bool) *ActionVpsFeatureUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsFeatureUpdateMetaGlobalInput) SetIncludes(value string) *ActionVpsFeatureUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsFeatureUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsFeatureUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsFeatureUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsFeatureUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsFeatureUpdateInput is a type for action input parameters
type ActionVpsFeatureUpdateInput struct {
	Enabled bool `json:"enabled"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionVpsFeatureUpdateInput) SetEnabled(value bool) *ActionVpsFeatureUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsFeatureUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsFeatureUpdateInput) SelectParameters(params ...string) *ActionVpsFeatureUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsFeatureUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsFeatureUpdateRequest is a type for the entire action request
type ActionVpsFeatureUpdateRequest struct {
	Feature map[string]interface{} `json:"feature"`
	Meta map[string]interface{} `json:"_meta"`
}


// ActionVpsFeatureUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsFeatureUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsFeatureUpdateResponse struct {
	Action *ActionVpsFeatureUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsFeatureUpdateMetaGlobalOutput `json:"_meta"`
	}
}


// Prepare the action for invocation
func (action *ActionVpsFeatureUpdate) Prepare() *ActionVpsFeatureUpdateInvocation {
	return &ActionVpsFeatureUpdateInvocation{
		Action: action,
		Path: "/v5.0/vpses/:vps_id/features/:feature_id",
	}
}

// ActionVpsFeatureUpdateInvocation is used to configure action for invocation
type ActionVpsFeatureUpdateInvocation struct {
	// Pointer to the action
	Action *ActionVpsFeatureUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsFeatureUpdateInput
	// Global meta input parameters
	MetaInput *ActionVpsFeatureUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsFeatureUpdateInvocation) SetPathParamInt(param string, value int64) *ActionVpsFeatureUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsFeatureUpdateInvocation) SetPathParamString(param string, value string) *ActionVpsFeatureUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsFeatureUpdateInvocation) SetInput(input *ActionVpsFeatureUpdateInput) *ActionVpsFeatureUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsFeatureUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsFeatureUpdateInvocation) SetMetaInput(input *ActionVpsFeatureUpdateMetaGlobalInput) *ActionVpsFeatureUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsFeatureUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsFeatureUpdateInvocation) Call() (*ActionVpsFeatureUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsFeatureUpdateInvocation) callAsBody() (*ActionVpsFeatureUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsFeatureUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsFeatureUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsFeatureUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsFeatureUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	req.SetInput(&ActionActionStatePollInput{
		Timeout: timeout,
	})
	req.Input.SelectParameters("Timeout")
	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsFeatureUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	req.SetInput(&ActionActionStatePollInput{
		Timeout: timeout,
		UpdateIn: updateIn,
	})
	req.Input.SelectParameters("Timeout", "UpdateIn")
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
func (resp *ActionVpsFeatureUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsFeatureUpdateInvocation) makeAllInputParams() *ActionVpsFeatureUpdateRequest {
	return &ActionVpsFeatureUpdateRequest{
		Feature: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsFeatureUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
	}

	return ret
}

func (inv *ActionVpsFeatureUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
