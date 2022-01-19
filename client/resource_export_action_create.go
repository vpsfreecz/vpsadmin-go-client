package client

import ()

// ActionExportCreate is a type for action Export#Create
type ActionExportCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionExportCreate(client *Client) *ActionExportCreate {
	return &ActionExportCreate{
		Client: client,
	}
}

// ActionExportCreateMetaGlobalInput is a type for action global meta input parameters
type ActionExportCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportCreateMetaGlobalInput) SetIncludes(value string) *ActionExportCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportCreateMetaGlobalInput) SetNo(value bool) *ActionExportCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportCreateMetaGlobalInput) SelectParameters(params ...string) *ActionExportCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportCreateInput is a type for action input parameters
type ActionExportCreateInput struct {
	AllVps       bool  `json:"all_vps"`
	Dataset      int64 `json:"dataset"`
	Enabled      bool  `json:"enabled"`
	RootSquash   bool  `json:"root_squash"`
	Rw           bool  `json:"rw"`
	Snapshot     int64 `json:"snapshot"`
	SubtreeCheck bool  `json:"subtree_check"`
	Sync         bool  `json:"sync"`
	Threads      int64 `json:"threads"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetAllVps sets parameter AllVps to value and selects it for sending
func (in *ActionExportCreateInput) SetAllVps(value bool) *ActionExportCreateInput {
	in.AllVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AllVps"] = nil
	return in
}

// SetDataset sets parameter Dataset to value and selects it for sending
func (in *ActionExportCreateInput) SetDataset(value int64) *ActionExportCreateInput {
	in.Dataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Dataset"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionExportCreateInput) SetEnabled(value bool) *ActionExportCreateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetRootSquash sets parameter RootSquash to value and selects it for sending
func (in *ActionExportCreateInput) SetRootSquash(value bool) *ActionExportCreateInput {
	in.RootSquash = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RootSquash"] = nil
	return in
}

// SetRw sets parameter Rw to value and selects it for sending
func (in *ActionExportCreateInput) SetRw(value bool) *ActionExportCreateInput {
	in.Rw = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Rw"] = nil
	return in
}

// SetSnapshot sets parameter Snapshot to value and selects it for sending
func (in *ActionExportCreateInput) SetSnapshot(value int64) *ActionExportCreateInput {
	in.Snapshot = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Snapshot"] = nil
	return in
}

// SetSubtreeCheck sets parameter SubtreeCheck to value and selects it for sending
func (in *ActionExportCreateInput) SetSubtreeCheck(value bool) *ActionExportCreateInput {
	in.SubtreeCheck = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SubtreeCheck"] = nil
	return in
}

// SetSync sets parameter Sync to value and selects it for sending
func (in *ActionExportCreateInput) SetSync(value bool) *ActionExportCreateInput {
	in.Sync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sync"] = nil
	return in
}

// SetThreads sets parameter Threads to value and selects it for sending
func (in *ActionExportCreateInput) SetThreads(value int64) *ActionExportCreateInput {
	in.Threads = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Threads"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportCreateInput) SelectParameters(params ...string) *ActionExportCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportCreateRequest is a type for the entire action request
type ActionExportCreateRequest struct {
	Export map[string]interface{} `json:"export"`
	Meta   map[string]interface{} `json:"_meta"`
}

// ActionExportCreateOutput is a type for action output parameters
type ActionExportCreateOutput struct {
	AllVps         bool                             `json:"all_vps"`
	CreatedAt      string                           `json:"created_at"`
	Dataset        *ActionDatasetShowOutput         `json:"dataset"`
	Enabled        bool                             `json:"enabled"`
	ExpirationDate string                           `json:"expiration_date"`
	HostIpAddress  *ActionHostIpAddressShowOutput   `json:"host_ip_address"`
	Id             int64                            `json:"id"`
	IpAddress      *ActionIpAddressShowOutput       `json:"ip_address"`
	Path           string                           `json:"path"`
	RootSquash     bool                             `json:"root_squash"`
	Rw             bool                             `json:"rw"`
	Snapshot       *ActionDatasetSnapshotShowOutput `json:"snapshot"`
	SubtreeCheck   bool                             `json:"subtree_check"`
	Sync           bool                             `json:"sync"`
	Threads        int64                            `json:"threads"`
	UpdatedAt      string                           `json:"updated_at"`
	User           *ActionUserShowOutput            `json:"user"`
}

// ActionExportCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionExportCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionExportCreateResponse struct {
	Action *ActionExportCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Export *ActionExportCreateOutput `json:"export"`
		// Global output metadata
		Meta *ActionExportCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionExportCreateOutput
}

// Prepare the action for invocation
func (action *ActionExportCreate) Prepare() *ActionExportCreateInvocation {
	return &ActionExportCreateInvocation{
		Action: action,
		Path:   "/v6.0/exports",
	}
}

// ActionExportCreateInvocation is used to configure action for invocation
type ActionExportCreateInvocation struct {
	// Pointer to the action
	Action *ActionExportCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionExportCreateInput
	// Global meta input parameters
	MetaInput *ActionExportCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionExportCreateInvocation) NewInput() *ActionExportCreateInput {
	inv.Input = &ActionExportCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionExportCreateInvocation) SetInput(input *ActionExportCreateInput) *ActionExportCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionExportCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportCreateInvocation) NewMetaInput() *ActionExportCreateMetaGlobalInput {
	inv.MetaInput = &ActionExportCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportCreateInvocation) SetMetaInput(input *ActionExportCreateMetaGlobalInput) *ActionExportCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionExportCreateInvocation) Call() (*ActionExportCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionExportCreateInvocation) callAsBody() (*ActionExportCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionExportCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Export
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionExportCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionExportCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionExportCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionExportCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionExportCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionExportCreateInvocation) makeAllInputParams() *ActionExportCreateRequest {
	return &ActionExportCreateRequest{
		Export: inv.makeInputParams(),
		Meta:   inv.makeMetaInputParams(),
	}
}

func (inv *ActionExportCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AllVps") {
			ret["all_vps"] = inv.Input.AllVps
		}
		if inv.IsParameterSelected("Dataset") {
			ret["dataset"] = inv.Input.Dataset
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("RootSquash") {
			ret["root_squash"] = inv.Input.RootSquash
		}
		if inv.IsParameterSelected("Rw") {
			ret["rw"] = inv.Input.Rw
		}
		if inv.IsParameterSelected("Snapshot") {
			ret["snapshot"] = inv.Input.Snapshot
		}
		if inv.IsParameterSelected("SubtreeCheck") {
			ret["subtree_check"] = inv.Input.SubtreeCheck
		}
		if inv.IsParameterSelected("Sync") {
			ret["sync"] = inv.Input.Sync
		}
		if inv.IsParameterSelected("Threads") {
			ret["threads"] = inv.Input.Threads
		}
	}

	return ret
}

func (inv *ActionExportCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
