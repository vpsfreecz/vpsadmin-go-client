package client

import (
)

// ActionClusterGenerateMigrationKeys is a type for action Cluster#Generate_migration_keys
type ActionClusterGenerateMigrationKeys struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterGenerateMigrationKeys(client *Client) *ActionClusterGenerateMigrationKeys {
	return &ActionClusterGenerateMigrationKeys{
		Client: client,
	}
}

// ActionClusterGenerateMigrationKeysMetaGlobalInput is a type for action global meta input parameters
type ActionClusterGenerateMigrationKeysMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterGenerateMigrationKeysMetaGlobalInput) SetNo(value bool) *ActionClusterGenerateMigrationKeysMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterGenerateMigrationKeysMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterGenerateMigrationKeysMetaGlobalInput) SelectParameters(params ...string) *ActionClusterGenerateMigrationKeysMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterGenerateMigrationKeysMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionClusterGenerateMigrationKeysRequest is a type for the entire action request
type ActionClusterGenerateMigrationKeysRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}


// ActionClusterGenerateMigrationKeysMetaGlobalOutput is a type for global output metadata parameters
type ActionClusterGenerateMigrationKeysMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionClusterGenerateMigrationKeysResponse struct {
	Action *ActionClusterGenerateMigrationKeys `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionClusterGenerateMigrationKeysMetaGlobalOutput `json:"_meta"`
	}
}

// Call the action directly without any path or input parameters
func (action *ActionClusterGenerateMigrationKeys) Call() (*ActionClusterGenerateMigrationKeysResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *ActionClusterGenerateMigrationKeys) Prepare() *ActionClusterGenerateMigrationKeysInvocation {
	return &ActionClusterGenerateMigrationKeysInvocation{
		Action: action,
		Path: "/v6.0/cluster/generate_migration_keys",
	}
}

// ActionClusterGenerateMigrationKeysInvocation is used to configure action for invocation
type ActionClusterGenerateMigrationKeysInvocation struct {
	// Pointer to the action
	Action *ActionClusterGenerateMigrationKeys

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterGenerateMigrationKeysMetaGlobalInput
}


// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterGenerateMigrationKeysInvocation) NewMetaInput() *ActionClusterGenerateMigrationKeysMetaGlobalInput {
	inv.MetaInput = &ActionClusterGenerateMigrationKeysMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterGenerateMigrationKeysInvocation) SetMetaInput(input *ActionClusterGenerateMigrationKeysMetaGlobalInput) *ActionClusterGenerateMigrationKeysInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterGenerateMigrationKeysInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterGenerateMigrationKeysInvocation) Call() (*ActionClusterGenerateMigrationKeysResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterGenerateMigrationKeysInvocation) callAsBody() (*ActionClusterGenerateMigrationKeysResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterGenerateMigrationKeysResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionClusterGenerateMigrationKeysResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionClusterGenerateMigrationKeysResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionClusterGenerateMigrationKeysResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionClusterGenerateMigrationKeysResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionClusterGenerateMigrationKeysResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionClusterGenerateMigrationKeysInvocation) makeAllInputParams() *ActionClusterGenerateMigrationKeysRequest {
	return &ActionClusterGenerateMigrationKeysRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionClusterGenerateMigrationKeysInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
