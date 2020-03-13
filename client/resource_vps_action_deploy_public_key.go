package client

import (
	"strings"
)

// ActionVpsDeployPublicKey is a type for action Vps#Deploy_public_key
type ActionVpsDeployPublicKey struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsDeployPublicKey(client *Client) *ActionVpsDeployPublicKey {
	return &ActionVpsDeployPublicKey{
		Client: client,
	}
}

// ActionVpsDeployPublicKeyMetaGlobalInput is a type for action global meta input parameters
type ActionVpsDeployPublicKeyMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsDeployPublicKeyMetaGlobalInput) SetNo(value bool) *ActionVpsDeployPublicKeyMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsDeployPublicKeyMetaGlobalInput) SetIncludes(value string) *ActionVpsDeployPublicKeyMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsDeployPublicKeyMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsDeployPublicKeyMetaGlobalInput) SelectParameters(params ...string) *ActionVpsDeployPublicKeyMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsDeployPublicKeyMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsDeployPublicKeyInput is a type for action input parameters
type ActionVpsDeployPublicKeyInput struct {
	PublicKey int64 `json:"public_key"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetPublicKey sets parameter PublicKey to value and selects it for sending
func (in *ActionVpsDeployPublicKeyInput) SetPublicKey(value int64) *ActionVpsDeployPublicKeyInput {
	in.PublicKey = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PublicKey"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsDeployPublicKeyInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsDeployPublicKeyInput) SelectParameters(params ...string) *ActionVpsDeployPublicKeyInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsDeployPublicKeyInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsDeployPublicKeyRequest is a type for the entire action request
type ActionVpsDeployPublicKeyRequest struct {
	Vps map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}


// ActionVpsDeployPublicKeyMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsDeployPublicKeyMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsDeployPublicKeyResponse struct {
	Action *ActionVpsDeployPublicKey `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsDeployPublicKeyMetaGlobalOutput `json:"_meta"`
	}
}


// Prepare the action for invocation
func (action *ActionVpsDeployPublicKey) Prepare() *ActionVpsDeployPublicKeyInvocation {
	return &ActionVpsDeployPublicKeyInvocation{
		Action: action,
		Path: "/v6.0/vpses/{vps_id}/deploy_public_key",
	}
}

// ActionVpsDeployPublicKeyInvocation is used to configure action for invocation
type ActionVpsDeployPublicKeyInvocation struct {
	// Pointer to the action
	Action *ActionVpsDeployPublicKey

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsDeployPublicKeyInput
	// Global meta input parameters
	MetaInput *ActionVpsDeployPublicKeyMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsDeployPublicKeyInvocation) SetPathParamInt(param string, value int64) *ActionVpsDeployPublicKeyInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsDeployPublicKeyInvocation) SetPathParamString(param string, value string) *ActionVpsDeployPublicKeyInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsDeployPublicKeyInvocation) NewInput() *ActionVpsDeployPublicKeyInput {
	inv.Input = &ActionVpsDeployPublicKeyInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsDeployPublicKeyInvocation) SetInput(input *ActionVpsDeployPublicKeyInput) *ActionVpsDeployPublicKeyInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsDeployPublicKeyInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsDeployPublicKeyInvocation) NewMetaInput() *ActionVpsDeployPublicKeyMetaGlobalInput {
	inv.MetaInput = &ActionVpsDeployPublicKeyMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsDeployPublicKeyInvocation) SetMetaInput(input *ActionVpsDeployPublicKeyMetaGlobalInput) *ActionVpsDeployPublicKeyInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsDeployPublicKeyInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsDeployPublicKeyInvocation) Call() (*ActionVpsDeployPublicKeyResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsDeployPublicKeyInvocation) callAsBody() (*ActionVpsDeployPublicKeyResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsDeployPublicKeyResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsDeployPublicKeyResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsDeployPublicKeyResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsDeployPublicKeyResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsDeployPublicKeyResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsDeployPublicKeyResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsDeployPublicKeyInvocation) makeAllInputParams() *ActionVpsDeployPublicKeyRequest {
	return &ActionVpsDeployPublicKeyRequest{
		Vps: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsDeployPublicKeyInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("PublicKey") {
			ret["public_key"] = inv.Input.PublicKey
		}
	}

	return ret
}

func (inv *ActionVpsDeployPublicKeyInvocation) makeMetaInputParams() map[string]interface{} {
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
