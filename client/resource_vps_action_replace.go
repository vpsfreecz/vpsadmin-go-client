package client

import (
	"strings"
)

// ActionVpsReplace is a type for action Vps#Replace
type ActionVpsReplace struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsReplace(client *Client) *ActionVpsReplace {
	return &ActionVpsReplace{
		Client: client,
	}
}

// ActionVpsReplaceMetaGlobalInput is a type for action global meta input parameters
type ActionVpsReplaceMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsReplaceMetaGlobalInput) SetIncludes(value string) *ActionVpsReplaceMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsReplaceMetaGlobalInput) SetNo(value bool) *ActionVpsReplaceMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsReplaceMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsReplaceMetaGlobalInput) SelectParameters(params ...string) *ActionVpsReplaceMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsReplaceMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsReplaceInput is a type for action input parameters
type ActionVpsReplaceInput struct {
	ExpirationDate string `json:"expiration_date"`
	Node           int64  `json:"node"`
	Start          bool   `json:"start"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetExpirationDate sets parameter ExpirationDate to value and selects it for sending
func (in *ActionVpsReplaceInput) SetExpirationDate(value string) *ActionVpsReplaceInput {
	in.ExpirationDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpirationDate"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsReplaceInput) SetNode(value int64) *ActionVpsReplaceInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionVpsReplaceInput) SetNodeNil(set bool) *ActionVpsReplaceInput {
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

// SetStart sets parameter Start to value and selects it for sending
func (in *ActionVpsReplaceInput) SetStart(value bool) *ActionVpsReplaceInput {
	in.Start = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Start"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsReplaceInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsReplaceInput) SelectParameters(params ...string) *ActionVpsReplaceInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsReplaceInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsReplaceInput) UnselectParameters(params ...string) *ActionVpsReplaceInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsReplaceInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsReplaceRequest is a type for the entire action request
type ActionVpsReplaceRequest struct {
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsReplaceOutput is a type for action output parameters
type ActionVpsReplaceOutput struct {
	Config           string                       `json:"config"`
	Cpu              int64                        `json:"cpu"`
	CpuIdle          float64                      `json:"cpu_idle"`
	CpuIowait        float64                      `json:"cpu_iowait"`
	CpuIrq           float64                      `json:"cpu_irq"`
	CpuLimit         int64                        `json:"cpu_limit"`
	CpuNice          float64                      `json:"cpu_nice"`
	CpuSoftirq       float64                      `json:"cpu_softirq"`
	CpuSystem        float64                      `json:"cpu_system"`
	CpuUser          float64                      `json:"cpu_user"`
	CreatedAt        string                       `json:"created_at"`
	Dataset          *ActionDatasetShowOutput     `json:"dataset"`
	Diskspace        int64                        `json:"diskspace"`
	DnsResolver      *ActionDnsResolverShowOutput `json:"dns_resolver"`
	Hostname         string                       `json:"hostname"`
	Id               int64                        `json:"id"`
	InRescueMode     bool                         `json:"in_rescue_mode"`
	Info             string                       `json:"info"`
	IsRunning        bool                         `json:"is_running"`
	Loadavg          float64                      `json:"loadavg"`
	ManageHostname   bool                         `json:"manage_hostname"`
	Memory           int64                        `json:"memory"`
	Node             *ActionNodeShowOutput        `json:"node"`
	Onboot           bool                         `json:"onboot"`
	Onstartall       bool                         `json:"onstartall"`
	OsTemplate       *ActionOsTemplateShowOutput  `json:"os_template"`
	ProcessCount     int64                        `json:"process_count"`
	StartMenuTimeout int64                        `json:"start_menu_timeout"`
	Swap             int64                        `json:"swap"`
	Uptime           int64                        `json:"uptime"`
	UsedDiskspace    int64                        `json:"used_diskspace"`
	UsedMemory       int64                        `json:"used_memory"`
	UsedSwap         int64                        `json:"used_swap"`
	User             *ActionUserShowOutput        `json:"user"`
}

// ActionVpsReplaceMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsReplaceMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsReplaceResponse struct {
	Action *ActionVpsReplace `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Vps *ActionVpsReplaceOutput `json:"vps"`
		// Global output metadata
		Meta *ActionVpsReplaceMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionVpsReplaceOutput
}

// Prepare the action for invocation
func (action *ActionVpsReplace) Prepare() *ActionVpsReplaceInvocation {
	return &ActionVpsReplaceInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/replace",
	}
}

// ActionVpsReplaceInvocation is used to configure action for invocation
type ActionVpsReplaceInvocation struct {
	// Pointer to the action
	Action *ActionVpsReplace

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsReplaceInput
	// Global meta input parameters
	MetaInput *ActionVpsReplaceMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsReplaceInvocation) SetPathParamInt(param string, value int64) *ActionVpsReplaceInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsReplaceInvocation) SetPathParamString(param string, value string) *ActionVpsReplaceInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsReplaceInvocation) NewInput() *ActionVpsReplaceInput {
	inv.Input = &ActionVpsReplaceInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsReplaceInvocation) SetInput(input *ActionVpsReplaceInput) *ActionVpsReplaceInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsReplaceInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsReplaceInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsReplaceInvocation) NewMetaInput() *ActionVpsReplaceMetaGlobalInput {
	inv.MetaInput = &ActionVpsReplaceMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsReplaceInvocation) SetMetaInput(input *ActionVpsReplaceMetaGlobalInput) *ActionVpsReplaceInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsReplaceInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsReplaceInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsReplaceInvocation) Call() (*ActionVpsReplaceResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsReplaceInvocation) callAsBody() (*ActionVpsReplaceResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsReplaceResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Vps
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsReplaceResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsReplaceResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsReplaceResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsReplaceResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsReplaceResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsReplaceInvocation) makeAllInputParams() *ActionVpsReplaceRequest {
	return &ActionVpsReplaceRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsReplaceInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ExpirationDate") {
			ret["expiration_date"] = inv.Input.ExpirationDate
		}
		if inv.IsParameterSelected("Node") {
			if inv.IsParameterNil("Node") {
				ret["node"] = nil
			} else {
				ret["node"] = inv.Input.Node
			}
		}
		if inv.IsParameterSelected("Start") {
			ret["start"] = inv.Input.Start
		}
	}

	return ret
}

func (inv *ActionVpsReplaceInvocation) makeMetaInputParams() map[string]interface{} {
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
