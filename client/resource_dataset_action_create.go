package client

import (
)

// ActionDatasetCreate is a type for action Dataset#Create
type ActionDatasetCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetCreate(client *Client) *ActionDatasetCreate {
	return &ActionDatasetCreate{
		Client: client,
	}
}

// ActionDatasetCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetCreateMetaGlobalInput) SetIncludes(value string) *ActionDatasetCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetCreateMetaGlobalInput) SetNo(value bool) *ActionDatasetCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetCreateInput is a type for action input parameters
type ActionDatasetCreateInput struct {
	Atime bool `json:"atime"`
	Automount bool `json:"automount"`
	Compression bool `json:"compression"`
	Dataset int64 `json:"dataset"`
	InheritUserNamespaceMap bool `json:"inherit_user_namespace_map"`
	Name string `json:"name"`
	Quota int64 `json:"quota"`
	Recordsize int64 `json:"recordsize"`
	Refquota int64 `json:"refquota"`
	Relatime bool `json:"relatime"`
	Sharenfs string `json:"sharenfs"`
	Sync string `json:"sync"`
	UserNamespaceMap int64 `json:"user_namespace_map"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetAtime sets parameter Atime to value and selects it for sending
func (in *ActionDatasetCreateInput) SetAtime(value bool) *ActionDatasetCreateInput {
	in.Atime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Atime"] = nil
	return in
}
// SetAutomount sets parameter Automount to value and selects it for sending
func (in *ActionDatasetCreateInput) SetAutomount(value bool) *ActionDatasetCreateInput {
	in.Automount = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Automount"] = nil
	return in
}
// SetCompression sets parameter Compression to value and selects it for sending
func (in *ActionDatasetCreateInput) SetCompression(value bool) *ActionDatasetCreateInput {
	in.Compression = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Compression"] = nil
	return in
}
// SetDataset sets parameter Dataset to value and selects it for sending
func (in *ActionDatasetCreateInput) SetDataset(value int64) *ActionDatasetCreateInput {
	in.Dataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Dataset"] = nil
	return in
}
// SetInheritUserNamespaceMap sets parameter InheritUserNamespaceMap to value and selects it for sending
func (in *ActionDatasetCreateInput) SetInheritUserNamespaceMap(value bool) *ActionDatasetCreateInput {
	in.InheritUserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["InheritUserNamespaceMap"] = nil
	return in
}
// SetName sets parameter Name to value and selects it for sending
func (in *ActionDatasetCreateInput) SetName(value string) *ActionDatasetCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}
// SetQuota sets parameter Quota to value and selects it for sending
func (in *ActionDatasetCreateInput) SetQuota(value int64) *ActionDatasetCreateInput {
	in.Quota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Quota"] = nil
	return in
}
// SetRecordsize sets parameter Recordsize to value and selects it for sending
func (in *ActionDatasetCreateInput) SetRecordsize(value int64) *ActionDatasetCreateInput {
	in.Recordsize = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Recordsize"] = nil
	return in
}
// SetRefquota sets parameter Refquota to value and selects it for sending
func (in *ActionDatasetCreateInput) SetRefquota(value int64) *ActionDatasetCreateInput {
	in.Refquota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Refquota"] = nil
	return in
}
// SetRelatime sets parameter Relatime to value and selects it for sending
func (in *ActionDatasetCreateInput) SetRelatime(value bool) *ActionDatasetCreateInput {
	in.Relatime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Relatime"] = nil
	return in
}
// SetSharenfs sets parameter Sharenfs to value and selects it for sending
func (in *ActionDatasetCreateInput) SetSharenfs(value string) *ActionDatasetCreateInput {
	in.Sharenfs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sharenfs"] = nil
	return in
}
// SetSync sets parameter Sync to value and selects it for sending
func (in *ActionDatasetCreateInput) SetSync(value string) *ActionDatasetCreateInput {
	in.Sync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sync"] = nil
	return in
}
// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionDatasetCreateInput) SetUserNamespaceMap(value int64) *ActionDatasetCreateInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserNamespaceMap"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetCreateInput) SelectParameters(params ...string) *ActionDatasetCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetCreateRequest is a type for the entire action request
type ActionDatasetCreateRequest struct {
	Dataset map[string]interface{} `json:"dataset"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionDatasetCreateOutput is a type for action output parameters
type ActionDatasetCreateOutput struct {
	Atime bool `json:"atime"`
	Avail int64 `json:"avail"`
	Compression bool `json:"compression"`
	CurrentHistoryId int64 `json:"current_history_id"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	Export *ActionExportShowOutput `json:"export"`
	Id int64 `json:"id"`
	Name string `json:"name"`
	Parent *ActionDatasetShowOutput `json:"parent"`
	Quota int64 `json:"quota"`
	Recordsize int64 `json:"recordsize"`
	Referenced int64 `json:"referenced"`
	Refquota int64 `json:"refquota"`
	Relatime bool `json:"relatime"`
	Sharenfs string `json:"sharenfs"`
	Sync string `json:"sync"`
	Used int64 `json:"used"`
	User *ActionUserShowOutput `json:"user"`
	UserNamespaceMap *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
}

// ActionDatasetCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDatasetCreateResponse struct {
	Action *ActionDatasetCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Dataset *ActionDatasetCreateOutput `json:"dataset"`
		// Global output metadata
		Meta *ActionDatasetCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDatasetCreateOutput
}


// Prepare the action for invocation
func (action *ActionDatasetCreate) Prepare() *ActionDatasetCreateInvocation {
	return &ActionDatasetCreateInvocation{
		Action: action,
		Path: "/v6.0/datasets",
	}
}

// ActionDatasetCreateInvocation is used to configure action for invocation
type ActionDatasetCreateInvocation struct {
	// Pointer to the action
	Action *ActionDatasetCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetCreateInput
	// Global meta input parameters
	MetaInput *ActionDatasetCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetCreateInvocation) NewInput() *ActionDatasetCreateInput {
	inv.Input = &ActionDatasetCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetCreateInvocation) SetInput(input *ActionDatasetCreateInput) *ActionDatasetCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetCreateInvocation) NewMetaInput() *ActionDatasetCreateMetaGlobalInput {
	inv.MetaInput = &ActionDatasetCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetCreateInvocation) SetMetaInput(input *ActionDatasetCreateMetaGlobalInput) *ActionDatasetCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetCreateInvocation) Call() (*ActionDatasetCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionDatasetCreateInvocation) callAsBody() (*ActionDatasetCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Dataset
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionDatasetCreateInvocation) makeAllInputParams() *ActionDatasetCreateRequest {
	return &ActionDatasetCreateRequest{
		Dataset: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Atime") {
			ret["atime"] = inv.Input.Atime
		}
		if inv.IsParameterSelected("Automount") {
			ret["automount"] = inv.Input.Automount
		}
		if inv.IsParameterSelected("Compression") {
			ret["compression"] = inv.Input.Compression
		}
		if inv.IsParameterSelected("Dataset") {
			ret["dataset"] = inv.Input.Dataset
		}
		if inv.IsParameterSelected("InheritUserNamespaceMap") {
			ret["inherit_user_namespace_map"] = inv.Input.InheritUserNamespaceMap
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
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
		if inv.IsParameterSelected("UserNamespaceMap") {
			ret["user_namespace_map"] = inv.Input.UserNamespaceMap
		}
	}

	return ret
}

func (inv *ActionDatasetCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
