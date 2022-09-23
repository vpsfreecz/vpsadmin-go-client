package client

import ()

// ActionPoolCreate is a type for action Pool#Create
type ActionPoolCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionPoolCreate(client *Client) *ActionPoolCreate {
	return &ActionPoolCreate{
		Client: client,
	}
}

// ActionPoolCreateMetaGlobalInput is a type for action global meta input parameters
type ActionPoolCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionPoolCreateMetaGlobalInput) SetIncludes(value string) *ActionPoolCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionPoolCreateMetaGlobalInput) SetNo(value bool) *ActionPoolCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionPoolCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPoolCreateMetaGlobalInput) SelectParameters(params ...string) *ActionPoolCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPoolCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPoolCreateInput is a type for action input parameters
type ActionPoolCreateInput struct {
	Atime         bool    `json:"atime"`
	CheckedAt     string  `json:"checked_at"`
	Compression   bool    `json:"compression"`
	Filesystem    string  `json:"filesystem"`
	Label         string  `json:"label"`
	MaxDatasets   int64   `json:"max_datasets"`
	Node          int64   `json:"node"`
	Quota         int64   `json:"quota"`
	Recordsize    int64   `json:"recordsize"`
	Refquota      int64   `json:"refquota"`
	RefquotaCheck bool    `json:"refquota_check"`
	Relatime      bool    `json:"relatime"`
	Role          string  `json:"role"`
	Scan          string  `json:"scan"`
	ScanPercent   float64 `json:"scan_percent"`
	Sharenfs      string  `json:"sharenfs"`
	State         string  `json:"state"`
	Sync          string  `json:"sync"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAtime sets parameter Atime to value and selects it for sending
func (in *ActionPoolCreateInput) SetAtime(value bool) *ActionPoolCreateInput {
	in.Atime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Atime"] = nil
	return in
}

// SetCheckedAt sets parameter CheckedAt to value and selects it for sending
func (in *ActionPoolCreateInput) SetCheckedAt(value string) *ActionPoolCreateInput {
	in.CheckedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CheckedAt"] = nil
	return in
}

// SetCompression sets parameter Compression to value and selects it for sending
func (in *ActionPoolCreateInput) SetCompression(value bool) *ActionPoolCreateInput {
	in.Compression = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Compression"] = nil
	return in
}

// SetFilesystem sets parameter Filesystem to value and selects it for sending
func (in *ActionPoolCreateInput) SetFilesystem(value string) *ActionPoolCreateInput {
	in.Filesystem = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Filesystem"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionPoolCreateInput) SetLabel(value string) *ActionPoolCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetMaxDatasets sets parameter MaxDatasets to value and selects it for sending
func (in *ActionPoolCreateInput) SetMaxDatasets(value int64) *ActionPoolCreateInput {
	in.MaxDatasets = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxDatasets"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionPoolCreateInput) SetNode(value int64) *ActionPoolCreateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionPoolCreateInput) SetNodeNil(set bool) *ActionPoolCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Node"] = nil
		in.SelectParameters("Node")
	} else {
		delete(in._nilParameters, "Node")
	}
	return in
}

// SetQuota sets parameter Quota to value and selects it for sending
func (in *ActionPoolCreateInput) SetQuota(value int64) *ActionPoolCreateInput {
	in.Quota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Quota"] = nil
	return in
}

// SetRecordsize sets parameter Recordsize to value and selects it for sending
func (in *ActionPoolCreateInput) SetRecordsize(value int64) *ActionPoolCreateInput {
	in.Recordsize = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Recordsize"] = nil
	return in
}

// SetRefquota sets parameter Refquota to value and selects it for sending
func (in *ActionPoolCreateInput) SetRefquota(value int64) *ActionPoolCreateInput {
	in.Refquota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Refquota"] = nil
	return in
}

// SetRefquotaCheck sets parameter RefquotaCheck to value and selects it for sending
func (in *ActionPoolCreateInput) SetRefquotaCheck(value bool) *ActionPoolCreateInput {
	in.RefquotaCheck = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RefquotaCheck"] = nil
	return in
}

// SetRelatime sets parameter Relatime to value and selects it for sending
func (in *ActionPoolCreateInput) SetRelatime(value bool) *ActionPoolCreateInput {
	in.Relatime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Relatime"] = nil
	return in
}

// SetRole sets parameter Role to value and selects it for sending
func (in *ActionPoolCreateInput) SetRole(value string) *ActionPoolCreateInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}

// SetScan sets parameter Scan to value and selects it for sending
func (in *ActionPoolCreateInput) SetScan(value string) *ActionPoolCreateInput {
	in.Scan = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Scan"] = nil
	return in
}

// SetScanPercent sets parameter ScanPercent to value and selects it for sending
func (in *ActionPoolCreateInput) SetScanPercent(value float64) *ActionPoolCreateInput {
	in.ScanPercent = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ScanPercent"] = nil
	return in
}

// SetSharenfs sets parameter Sharenfs to value and selects it for sending
func (in *ActionPoolCreateInput) SetSharenfs(value string) *ActionPoolCreateInput {
	in.Sharenfs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sharenfs"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionPoolCreateInput) SetState(value string) *ActionPoolCreateInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetSync sets parameter Sync to value and selects it for sending
func (in *ActionPoolCreateInput) SetSync(value string) *ActionPoolCreateInput {
	in.Sync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sync"] = nil
	return in
}

// SelectParameters sets parameters from ActionPoolCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPoolCreateInput) SelectParameters(params ...string) *ActionPoolCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionPoolCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionPoolCreateInput) UnselectParameters(params ...string) *ActionPoolCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionPoolCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPoolCreateRequest is a type for the entire action request
type ActionPoolCreateRequest struct {
	Pool map[string]interface{} `json:"pool"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionPoolCreateOutput is a type for action output parameters
type ActionPoolCreateOutput struct {
	Atime         bool                  `json:"atime"`
	Avail         int64                 `json:"avail"`
	CheckedAt     string                `json:"checked_at"`
	Compression   bool                  `json:"compression"`
	Filesystem    string                `json:"filesystem"`
	Id            int64                 `json:"id"`
	Label         string                `json:"label"`
	MaxDatasets   int64                 `json:"max_datasets"`
	Name          string                `json:"name"`
	Node          *ActionNodeShowOutput `json:"node"`
	Quota         int64                 `json:"quota"`
	Recordsize    int64                 `json:"recordsize"`
	Referenced    int64                 `json:"referenced"`
	Refquota      int64                 `json:"refquota"`
	RefquotaCheck bool                  `json:"refquota_check"`
	Relatime      bool                  `json:"relatime"`
	Role          string                `json:"role"`
	Scan          string                `json:"scan"`
	ScanPercent   float64               `json:"scan_percent"`
	Sharenfs      string                `json:"sharenfs"`
	State         string                `json:"state"`
	Sync          string                `json:"sync"`
	Used          int64                 `json:"used"`
}

// ActionPoolCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionPoolCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionPoolCreateResponse struct {
	Action *ActionPoolCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Pool *ActionPoolCreateOutput `json:"pool"`
		// Global output metadata
		Meta *ActionPoolCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionPoolCreateOutput
}

// Prepare the action for invocation
func (action *ActionPoolCreate) Prepare() *ActionPoolCreateInvocation {
	return &ActionPoolCreateInvocation{
		Action: action,
		Path:   "/v6.0/pools",
	}
}

// ActionPoolCreateInvocation is used to configure action for invocation
type ActionPoolCreateInvocation struct {
	// Pointer to the action
	Action *ActionPoolCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionPoolCreateInput
	// Global meta input parameters
	MetaInput *ActionPoolCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionPoolCreateInvocation) NewInput() *ActionPoolCreateInput {
	inv.Input = &ActionPoolCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionPoolCreateInvocation) SetInput(input *ActionPoolCreateInput) *ActionPoolCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionPoolCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionPoolCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionPoolCreateInvocation) NewMetaInput() *ActionPoolCreateMetaGlobalInput {
	inv.MetaInput = &ActionPoolCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionPoolCreateInvocation) SetMetaInput(input *ActionPoolCreateMetaGlobalInput) *ActionPoolCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionPoolCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionPoolCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionPoolCreateInvocation) Call() (*ActionPoolCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionPoolCreateInvocation) callAsBody() (*ActionPoolCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionPoolCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Pool
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionPoolCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionPoolCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionPoolCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionPoolCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionPoolCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionPoolCreateInvocation) makeAllInputParams() *ActionPoolCreateRequest {
	return &ActionPoolCreateRequest{
		Pool: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionPoolCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Atime") {
			ret["atime"] = inv.Input.Atime
		}
		if inv.IsParameterSelected("CheckedAt") {
			ret["checked_at"] = inv.Input.CheckedAt
		}
		if inv.IsParameterSelected("Compression") {
			ret["compression"] = inv.Input.Compression
		}
		if inv.IsParameterSelected("Filesystem") {
			ret["filesystem"] = inv.Input.Filesystem
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("MaxDatasets") {
			ret["max_datasets"] = inv.Input.MaxDatasets
		}
		if inv.IsParameterSelected("Node") {
			if inv.IsParameterNil("Node") {
				ret["node"] = nil
			} else {
				ret["node"] = inv.Input.Node
			}
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
		if inv.IsParameterSelected("RefquotaCheck") {
			ret["refquota_check"] = inv.Input.RefquotaCheck
		}
		if inv.IsParameterSelected("Relatime") {
			ret["relatime"] = inv.Input.Relatime
		}
		if inv.IsParameterSelected("Role") {
			ret["role"] = inv.Input.Role
		}
		if inv.IsParameterSelected("Scan") {
			ret["scan"] = inv.Input.Scan
		}
		if inv.IsParameterSelected("ScanPercent") {
			ret["scan_percent"] = inv.Input.ScanPercent
		}
		if inv.IsParameterSelected("Sharenfs") {
			ret["sharenfs"] = inv.Input.Sharenfs
		}
		if inv.IsParameterSelected("State") {
			ret["state"] = inv.Input.State
		}
		if inv.IsParameterSelected("Sync") {
			ret["sync"] = inv.Input.Sync
		}
	}

	return ret
}

func (inv *ActionPoolCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
