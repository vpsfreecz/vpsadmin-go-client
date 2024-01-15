package client

import (
	"strings"
)

// ActionDatasetUpdate is a type for action Dataset#Update
type ActionDatasetUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetUpdate(client *Client) *ActionDatasetUpdate {
	return &ActionDatasetUpdate{
		Client: client,
	}
}

// ActionDatasetUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetUpdateMetaGlobalInput) SetIncludes(value string) *ActionDatasetUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetUpdateMetaGlobalInput) SetNo(value bool) *ActionDatasetUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetUpdateInput is a type for action input parameters
type ActionDatasetUpdateInput struct {
	AdminLockType string `json:"admin_lock_type"`
	AdminOverride bool   `json:"admin_override"`
	Atime         bool   `json:"atime"`
	Compression   bool   `json:"compression"`
	Quota         int64  `json:"quota"`
	Recordsize    int64  `json:"recordsize"`
	Refquota      int64  `json:"refquota"`
	Relatime      bool   `json:"relatime"`
	Sharenfs      string `json:"sharenfs"`
	Sync          string `json:"sync"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAdminLockType sets parameter AdminLockType to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetAdminLockType(value string) *ActionDatasetUpdateInput {
	in.AdminLockType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AdminLockType"] = nil
	return in
}

// SetAdminOverride sets parameter AdminOverride to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetAdminOverride(value bool) *ActionDatasetUpdateInput {
	in.AdminOverride = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AdminOverride"] = nil
	return in
}

// SetAtime sets parameter Atime to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetAtime(value bool) *ActionDatasetUpdateInput {
	in.Atime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Atime"] = nil
	return in
}

// SetCompression sets parameter Compression to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetCompression(value bool) *ActionDatasetUpdateInput {
	in.Compression = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Compression"] = nil
	return in
}

// SetQuota sets parameter Quota to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetQuota(value int64) *ActionDatasetUpdateInput {
	in.Quota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Quota"] = nil
	return in
}

// SetRecordsize sets parameter Recordsize to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetRecordsize(value int64) *ActionDatasetUpdateInput {
	in.Recordsize = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Recordsize"] = nil
	return in
}

// SetRefquota sets parameter Refquota to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetRefquota(value int64) *ActionDatasetUpdateInput {
	in.Refquota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Refquota"] = nil
	return in
}

// SetRelatime sets parameter Relatime to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetRelatime(value bool) *ActionDatasetUpdateInput {
	in.Relatime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Relatime"] = nil
	return in
}

// SetSharenfs sets parameter Sharenfs to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetSharenfs(value string) *ActionDatasetUpdateInput {
	in.Sharenfs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sharenfs"] = nil
	return in
}

// SetSync sets parameter Sync to value and selects it for sending
func (in *ActionDatasetUpdateInput) SetSync(value string) *ActionDatasetUpdateInput {
	in.Sync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sync"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetUpdateInput) SelectParameters(params ...string) *ActionDatasetUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetUpdateInput) UnselectParameters(params ...string) *ActionDatasetUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetUpdateRequest is a type for the entire action request
type ActionDatasetUpdateRequest struct {
	Dataset map[string]interface{} `json:"dataset"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionDatasetUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDatasetUpdateResponse struct {
	Action *ActionDatasetUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDatasetUpdateMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionDatasetUpdate) Prepare() *ActionDatasetUpdateInvocation {
	return &ActionDatasetUpdateInvocation{
		Action: action,
		Path:   "/v6.0/datasets/{dataset_id}",
	}
}

// ActionDatasetUpdateInvocation is used to configure action for invocation
type ActionDatasetUpdateInvocation struct {
	// Pointer to the action
	Action *ActionDatasetUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetUpdateInput
	// Global meta input parameters
	MetaInput *ActionDatasetUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetUpdateInvocation) SetPathParamInt(param string, value int64) *ActionDatasetUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetUpdateInvocation) SetPathParamString(param string, value string) *ActionDatasetUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetUpdateInvocation) NewInput() *ActionDatasetUpdateInput {
	inv.Input = &ActionDatasetUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetUpdateInvocation) SetInput(input *ActionDatasetUpdateInput) *ActionDatasetUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetUpdateInvocation) NewMetaInput() *ActionDatasetUpdateMetaGlobalInput {
	inv.MetaInput = &ActionDatasetUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetUpdateInvocation) SetMetaInput(input *ActionDatasetUpdateMetaGlobalInput) *ActionDatasetUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetUpdateInvocation) Call() (*ActionDatasetUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetUpdateInvocation) callAsBody() (*ActionDatasetUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
			Timeout:  timeout,
			UpdateIn: updateIn,
			Status:   pollResp.Output.Status,
			Current:  pollResp.Output.Current,
			Total:    pollResp.Output.Total,
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
func (resp *ActionDatasetUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetUpdateInvocation) makeAllInputParams() *ActionDatasetUpdateRequest {
	return &ActionDatasetUpdateRequest{
		Dataset: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AdminLockType") {
			ret["admin_lock_type"] = inv.Input.AdminLockType
		}
		if inv.IsParameterSelected("AdminOverride") {
			ret["admin_override"] = inv.Input.AdminOverride
		}
		if inv.IsParameterSelected("Atime") {
			ret["atime"] = inv.Input.Atime
		}
		if inv.IsParameterSelected("Compression") {
			ret["compression"] = inv.Input.Compression
		}
		if inv.IsParameterSelected("Quota") {
			ret["quota"] = inv.Input.Quota
		}
		if inv.IsParameterSelected("Recordsize") {
			ret["recordsize"] = inv.Input.Recordsize
		}
		if inv.IsParameterSelected("Refquota") {
			ret["refquota"] = inv.Input.Refquota
		}
		if inv.IsParameterSelected("Relatime") {
			ret["relatime"] = inv.Input.Relatime
		}
		if inv.IsParameterSelected("Sharenfs") {
			ret["sharenfs"] = inv.Input.Sharenfs
		}
		if inv.IsParameterSelected("Sync") {
			ret["sync"] = inv.Input.Sync
		}
	}

	return ret
}

func (inv *ActionDatasetUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
