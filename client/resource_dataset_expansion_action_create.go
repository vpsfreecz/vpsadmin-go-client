package client

import ()

// ActionDatasetExpansionCreate is a type for action Dataset_expansion#Create
type ActionDatasetExpansionCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetExpansionCreate(client *Client) *ActionDatasetExpansionCreate {
	return &ActionDatasetExpansionCreate{
		Client: client,
	}
}

// ActionDatasetExpansionCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetExpansionCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetExpansionCreateMetaGlobalInput) SetIncludes(value string) *ActionDatasetExpansionCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetExpansionCreateMetaGlobalInput) SetNo(value bool) *ActionDatasetExpansionCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetExpansionCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetExpansionCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionCreateInput is a type for action input parameters
type ActionDatasetExpansionCreateInput struct {
	AddedSpace             int64 `json:"added_space"`
	Dataset                int64 `json:"dataset"`
	EnableNotifications    bool  `json:"enable_notifications"`
	EnableShrink           bool  `json:"enable_shrink"`
	MaxOverRefquotaSeconds int64 `json:"max_over_refquota_seconds"`
	StopVps                bool  `json:"stop_vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddedSpace sets parameter AddedSpace to value and selects it for sending
func (in *ActionDatasetExpansionCreateInput) SetAddedSpace(value int64) *ActionDatasetExpansionCreateInput {
	in.AddedSpace = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AddedSpace"] = nil
	return in
}

// SetDataset sets parameter Dataset to value and selects it for sending
func (in *ActionDatasetExpansionCreateInput) SetDataset(value int64) *ActionDatasetExpansionCreateInput {
	in.Dataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDatasetNil(false)
	in._selectedParameters["Dataset"] = nil
	return in
}

// SetDatasetNil sets parameter Dataset to nil and selects it for sending
func (in *ActionDatasetExpansionCreateInput) SetDatasetNil(set bool) *ActionDatasetExpansionCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Dataset"] = nil
		in.SelectParameters("Dataset")
	} else {
		delete(in._nilParameters, "Dataset")
	}
	return in
}

// SetEnableNotifications sets parameter EnableNotifications to value and selects it for sending
func (in *ActionDatasetExpansionCreateInput) SetEnableNotifications(value bool) *ActionDatasetExpansionCreateInput {
	in.EnableNotifications = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableNotifications"] = nil
	return in
}

// SetEnableShrink sets parameter EnableShrink to value and selects it for sending
func (in *ActionDatasetExpansionCreateInput) SetEnableShrink(value bool) *ActionDatasetExpansionCreateInput {
	in.EnableShrink = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableShrink"] = nil
	return in
}

// SetMaxOverRefquotaSeconds sets parameter MaxOverRefquotaSeconds to value and selects it for sending
func (in *ActionDatasetExpansionCreateInput) SetMaxOverRefquotaSeconds(value int64) *ActionDatasetExpansionCreateInput {
	in.MaxOverRefquotaSeconds = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxOverRefquotaSeconds"] = nil
	return in
}

// SetStopVps sets parameter StopVps to value and selects it for sending
func (in *ActionDatasetExpansionCreateInput) SetStopVps(value bool) *ActionDatasetExpansionCreateInput {
	in.StopVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["StopVps"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionCreateInput) SelectParameters(params ...string) *ActionDatasetExpansionCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetExpansionCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetExpansionCreateInput) UnselectParameters(params ...string) *ActionDatasetExpansionCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetExpansionCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionCreateRequest is a type for the entire action request
type ActionDatasetExpansionCreateRequest struct {
	DatasetExpansion map[string]interface{} `json:"dataset_expansion"`
	Meta             map[string]interface{} `json:"_meta"`
}

// ActionDatasetExpansionCreateOutput is a type for action output parameters
type ActionDatasetExpansionCreateOutput struct {
	AddedSpace             int64                    `json:"added_space"`
	CreatedAt              string                   `json:"created_at"`
	Dataset                *ActionDatasetShowOutput `json:"dataset"`
	EnableNotifications    bool                     `json:"enable_notifications"`
	EnableShrink           bool                     `json:"enable_shrink"`
	Id                     int64                    `json:"id"`
	MaxOverRefquotaSeconds int64                    `json:"max_over_refquota_seconds"`
	OriginalRefquota       int64                    `json:"original_refquota"`
	OverRefquotaSeconds    int64                    `json:"over_refquota_seconds"`
	State                  string                   `json:"state"`
	StopVps                bool                     `json:"stop_vps"`
	Vps                    *ActionVpsShowOutput     `json:"vps"`
}

// ActionDatasetExpansionCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionDatasetExpansionCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDatasetExpansionCreateResponse struct {
	Action *ActionDatasetExpansionCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DatasetExpansion *ActionDatasetExpansionCreateOutput `json:"dataset_expansion"`
		// Global output metadata
		Meta *ActionDatasetExpansionCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDatasetExpansionCreateOutput
}

// Prepare the action for invocation
func (action *ActionDatasetExpansionCreate) Prepare() *ActionDatasetExpansionCreateInvocation {
	return &ActionDatasetExpansionCreateInvocation{
		Action: action,
		Path:   "/v6.0/dataset_expansions",
	}
}

// ActionDatasetExpansionCreateInvocation is used to configure action for invocation
type ActionDatasetExpansionCreateInvocation struct {
	// Pointer to the action
	Action *ActionDatasetExpansionCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetExpansionCreateInput
	// Global meta input parameters
	MetaInput *ActionDatasetExpansionCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetExpansionCreateInvocation) NewInput() *ActionDatasetExpansionCreateInput {
	inv.Input = &ActionDatasetExpansionCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetExpansionCreateInvocation) SetInput(input *ActionDatasetExpansionCreateInput) *ActionDatasetExpansionCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetExpansionCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetExpansionCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetExpansionCreateInvocation) NewMetaInput() *ActionDatasetExpansionCreateMetaGlobalInput {
	inv.MetaInput = &ActionDatasetExpansionCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetExpansionCreateInvocation) SetMetaInput(input *ActionDatasetExpansionCreateMetaGlobalInput) *ActionDatasetExpansionCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetExpansionCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetExpansionCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetExpansionCreateInvocation) Call() (*ActionDatasetExpansionCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetExpansionCreateInvocation) callAsBody() (*ActionDatasetExpansionCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetExpansionCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DatasetExpansion
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDatasetExpansionCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDatasetExpansionCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDatasetExpansionCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDatasetExpansionCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDatasetExpansionCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDatasetExpansionCreateInvocation) makeAllInputParams() *ActionDatasetExpansionCreateRequest {
	return &ActionDatasetExpansionCreateRequest{
		DatasetExpansion: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetExpansionCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AddedSpace") {
			ret["added_space"] = inv.Input.AddedSpace
		}
		if inv.IsParameterSelected("Dataset") {
			if inv.IsParameterNil("Dataset") {
				ret["dataset"] = nil
			} else {
				ret["dataset"] = inv.Input.Dataset
			}
		}
		if inv.IsParameterSelected("EnableNotifications") {
			ret["enable_notifications"] = inv.Input.EnableNotifications
		}
		if inv.IsParameterSelected("EnableShrink") {
			ret["enable_shrink"] = inv.Input.EnableShrink
		}
		if inv.IsParameterSelected("MaxOverRefquotaSeconds") {
			ret["max_over_refquota_seconds"] = inv.Input.MaxOverRefquotaSeconds
		}
		if inv.IsParameterSelected("StopVps") {
			ret["stop_vps"] = inv.Input.StopVps
		}
	}

	return ret
}

func (inv *ActionDatasetExpansionCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
