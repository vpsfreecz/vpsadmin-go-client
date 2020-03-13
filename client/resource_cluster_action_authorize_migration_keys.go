package client

import (
)

// ActionClusterAuthorizeMigrationKeys is a type for action Cluster#Authorize_migration_keys
type ActionClusterAuthorizeMigrationKeys struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterAuthorizeMigrationKeys(client *Client) *ActionClusterAuthorizeMigrationKeys {
	return &ActionClusterAuthorizeMigrationKeys{
		Client: client,
	}
}

// ActionClusterAuthorizeMigrationKeysMetaGlobalInput is a type for action global meta input parameters
type ActionClusterAuthorizeMigrationKeysMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterAuthorizeMigrationKeysMetaGlobalInput) SetNo(value bool) *ActionClusterAuthorizeMigrationKeysMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterAuthorizeMigrationKeysMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterAuthorizeMigrationKeysMetaGlobalInput) SelectParameters(params ...string) *ActionClusterAuthorizeMigrationKeysMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterAuthorizeMigrationKeysMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionClusterAuthorizeMigrationKeysRequest is a type for the entire action request
type ActionClusterAuthorizeMigrationKeysRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}


// ActionClusterAuthorizeMigrationKeysMetaGlobalOutput is a type for global output metadata parameters
type ActionClusterAuthorizeMigrationKeysMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionClusterAuthorizeMigrationKeysResponse struct {
	Action *ActionClusterAuthorizeMigrationKeys `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionClusterAuthorizeMigrationKeysMetaGlobalOutput `json:"_meta"`
	}
}

// Call the action directly without any path or input parameters
func (action *ActionClusterAuthorizeMigrationKeys) Call() (*ActionClusterAuthorizeMigrationKeysResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *ActionClusterAuthorizeMigrationKeys) Prepare() *ActionClusterAuthorizeMigrationKeysInvocation {
	return &ActionClusterAuthorizeMigrationKeysInvocation{
		Action: action,
		Path: "/v6.0/cluster/authorize_migration_keys",
	}
}

// ActionClusterAuthorizeMigrationKeysInvocation is used to configure action for invocation
type ActionClusterAuthorizeMigrationKeysInvocation struct {
	// Pointer to the action
	Action *ActionClusterAuthorizeMigrationKeys

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterAuthorizeMigrationKeysMetaGlobalInput
}


// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterAuthorizeMigrationKeysInvocation) NewMetaInput() *ActionClusterAuthorizeMigrationKeysMetaGlobalInput {
	inv.MetaInput = &ActionClusterAuthorizeMigrationKeysMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterAuthorizeMigrationKeysInvocation) SetMetaInput(input *ActionClusterAuthorizeMigrationKeysMetaGlobalInput) *ActionClusterAuthorizeMigrationKeysInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterAuthorizeMigrationKeysInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterAuthorizeMigrationKeysInvocation) Call() (*ActionClusterAuthorizeMigrationKeysResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterAuthorizeMigrationKeysInvocation) callAsBody() (*ActionClusterAuthorizeMigrationKeysResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterAuthorizeMigrationKeysResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionClusterAuthorizeMigrationKeysResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionClusterAuthorizeMigrationKeysResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionClusterAuthorizeMigrationKeysResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionClusterAuthorizeMigrationKeysResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionClusterAuthorizeMigrationKeysResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionClusterAuthorizeMigrationKeysInvocation) makeAllInputParams() *ActionClusterAuthorizeMigrationKeysRequest {
	return &ActionClusterAuthorizeMigrationKeysRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionClusterAuthorizeMigrationKeysInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
