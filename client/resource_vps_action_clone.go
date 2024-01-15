package client

import (
	"strings"
)

// ActionVpsClone is a type for action Vps#Clone
type ActionVpsClone struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsClone(client *Client) *ActionVpsClone {
	return &ActionVpsClone{
		Client: client,
	}
}

// ActionVpsCloneMetaGlobalInput is a type for action global meta input parameters
type ActionVpsCloneMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsCloneMetaGlobalInput) SetIncludes(value string) *ActionVpsCloneMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsCloneMetaGlobalInput) SetNo(value bool) *ActionVpsCloneMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsCloneMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsCloneMetaGlobalInput) SelectParameters(params ...string) *ActionVpsCloneMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsCloneMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsCloneInput is a type for action input parameters
type ActionVpsCloneInput struct {
	AddressLocation int64  `json:"address_location"`
	DatasetPlans    bool   `json:"dataset_plans"`
	Environment     int64  `json:"environment"`
	Features        bool   `json:"features"`
	Hostname        string `json:"hostname"`
	KeepSnapshots   bool   `json:"keep_snapshots"`
	Location        int64  `json:"location"`
	Node            int64  `json:"node"`
	Platform        string `json:"platform"`
	Resources       bool   `json:"resources"`
	Stop            bool   `json:"stop"`
	Subdatasets     bool   `json:"subdatasets"`
	User            int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddressLocation sets parameter AddressLocation to value and selects it for sending
func (in *ActionVpsCloneInput) SetAddressLocation(value int64) *ActionVpsCloneInput {
	in.AddressLocation = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetAddressLocationNil(false)
	in._selectedParameters["AddressLocation"] = nil
	return in
}

// SetAddressLocationNil sets parameter AddressLocation to nil and selects it for sending
func (in *ActionVpsCloneInput) SetAddressLocationNil(set bool) *ActionVpsCloneInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["AddressLocation"] = nil
		in.SelectParameters("AddressLocation")
	} else {
		delete(in._nilParameters, "AddressLocation")
	}
	return in
}

// SetDatasetPlans sets parameter DatasetPlans to value and selects it for sending
func (in *ActionVpsCloneInput) SetDatasetPlans(value bool) *ActionVpsCloneInput {
	in.DatasetPlans = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DatasetPlans"] = nil
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionVpsCloneInput) SetEnvironment(value int64) *ActionVpsCloneInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionVpsCloneInput) SetEnvironmentNil(set bool) *ActionVpsCloneInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Environment"] = nil
		in.SelectParameters("Environment")
	} else {
		delete(in._nilParameters, "Environment")
	}
	return in
}

// SetFeatures sets parameter Features to value and selects it for sending
func (in *ActionVpsCloneInput) SetFeatures(value bool) *ActionVpsCloneInput {
	in.Features = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Features"] = nil
	return in
}

// SetHostname sets parameter Hostname to value and selects it for sending
func (in *ActionVpsCloneInput) SetHostname(value string) *ActionVpsCloneInput {
	in.Hostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Hostname"] = nil
	return in
}

// SetKeepSnapshots sets parameter KeepSnapshots to value and selects it for sending
func (in *ActionVpsCloneInput) SetKeepSnapshots(value bool) *ActionVpsCloneInput {
	in.KeepSnapshots = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["KeepSnapshots"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionVpsCloneInput) SetLocation(value int64) *ActionVpsCloneInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionVpsCloneInput) SetLocationNil(set bool) *ActionVpsCloneInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Location"] = nil
		in.SelectParameters("Location")
	} else {
		delete(in._nilParameters, "Location")
	}
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsCloneInput) SetNode(value int64) *ActionVpsCloneInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionVpsCloneInput) SetNodeNil(set bool) *ActionVpsCloneInput {
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

// SetPlatform sets parameter Platform to value and selects it for sending
func (in *ActionVpsCloneInput) SetPlatform(value string) *ActionVpsCloneInput {
	in.Platform = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Platform"] = nil
	return in
}

// SetResources sets parameter Resources to value and selects it for sending
func (in *ActionVpsCloneInput) SetResources(value bool) *ActionVpsCloneInput {
	in.Resources = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Resources"] = nil
	return in
}

// SetStop sets parameter Stop to value and selects it for sending
func (in *ActionVpsCloneInput) SetStop(value bool) *ActionVpsCloneInput {
	in.Stop = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Stop"] = nil
	return in
}

// SetSubdatasets sets parameter Subdatasets to value and selects it for sending
func (in *ActionVpsCloneInput) SetSubdatasets(value bool) *ActionVpsCloneInput {
	in.Subdatasets = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Subdatasets"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsCloneInput) SetUser(value int64) *ActionVpsCloneInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionVpsCloneInput) SetUserNil(set bool) *ActionVpsCloneInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionVpsCloneInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsCloneInput) SelectParameters(params ...string) *ActionVpsCloneInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsCloneInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsCloneInput) UnselectParameters(params ...string) *ActionVpsCloneInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsCloneInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsCloneRequest is a type for the entire action request
type ActionVpsCloneRequest struct {
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsCloneOutput is a type for action output parameters
type ActionVpsCloneOutput struct {
	AllowAdminModifications bool                              `json:"allow_admin_modifications"`
	AutostartEnable         bool                              `json:"autostart_enable"`
	AutostartPriority       int64                             `json:"autostart_priority"`
	CgroupVersion           string                            `json:"cgroup_version"`
	Config                  string                            `json:"config"`
	Cpu                     int64                             `json:"cpu"`
	CpuIdle                 float64                           `json:"cpu_idle"`
	CpuIowait               float64                           `json:"cpu_iowait"`
	CpuIrq                  float64                           `json:"cpu_irq"`
	CpuLimit                int64                             `json:"cpu_limit"`
	CpuNice                 float64                           `json:"cpu_nice"`
	CpuSoftirq              float64                           `json:"cpu_softirq"`
	CpuSystem               float64                           `json:"cpu_system"`
	CpuUser                 float64                           `json:"cpu_user"`
	CreatedAt               string                            `json:"created_at"`
	Dataset                 *ActionDatasetShowOutput          `json:"dataset"`
	Diskspace               int64                             `json:"diskspace"`
	DnsResolver             *ActionDnsResolverShowOutput      `json:"dns_resolver"`
	Hostname                string                            `json:"hostname"`
	Id                      int64                             `json:"id"`
	InRescueMode            bool                              `json:"in_rescue_mode"`
	Info                    string                            `json:"info"`
	IsRunning               bool                              `json:"is_running"`
	Loadavg                 float64                           `json:"loadavg"`
	ManageHostname          bool                              `json:"manage_hostname"`
	Memory                  int64                             `json:"memory"`
	Node                    *ActionNodeShowOutput             `json:"node"`
	Onstartall              bool                              `json:"onstartall"`
	OsTemplate              *ActionOsTemplateShowOutput       `json:"os_template"`
	Pool                    *ActionPoolShowOutput             `json:"pool"`
	ProcessCount            int64                             `json:"process_count"`
	StartMenuTimeout        int64                             `json:"start_menu_timeout"`
	Swap                    int64                             `json:"swap"`
	Uptime                  int64                             `json:"uptime"`
	UsedDiskspace           int64                             `json:"used_diskspace"`
	UsedMemory              int64                             `json:"used_memory"`
	UsedSwap                int64                             `json:"used_swap"`
	User                    *ActionUserShowOutput             `json:"user"`
	UserNamespaceMap        *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
}

// ActionVpsCloneMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsCloneMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsCloneResponse struct {
	Action *ActionVpsClone `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Vps *ActionVpsCloneOutput `json:"vps"`
		// Global output metadata
		Meta *ActionVpsCloneMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionVpsCloneOutput
}

// Prepare the action for invocation
func (action *ActionVpsClone) Prepare() *ActionVpsCloneInvocation {
	return &ActionVpsCloneInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/clone",
	}
}

// ActionVpsCloneInvocation is used to configure action for invocation
type ActionVpsCloneInvocation struct {
	// Pointer to the action
	Action *ActionVpsClone

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsCloneInput
	// Global meta input parameters
	MetaInput *ActionVpsCloneMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsCloneInvocation) SetPathParamInt(param string, value int64) *ActionVpsCloneInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsCloneInvocation) SetPathParamString(param string, value string) *ActionVpsCloneInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsCloneInvocation) NewInput() *ActionVpsCloneInput {
	inv.Input = &ActionVpsCloneInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsCloneInvocation) SetInput(input *ActionVpsCloneInput) *ActionVpsCloneInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsCloneInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsCloneInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsCloneInvocation) NewMetaInput() *ActionVpsCloneMetaGlobalInput {
	inv.MetaInput = &ActionVpsCloneMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsCloneInvocation) SetMetaInput(input *ActionVpsCloneMetaGlobalInput) *ActionVpsCloneInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsCloneInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsCloneInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsCloneInvocation) Call() (*ActionVpsCloneResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsCloneInvocation) callAsBody() (*ActionVpsCloneResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsCloneResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Vps
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsCloneResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsCloneResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsCloneResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsCloneResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsCloneResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsCloneInvocation) makeAllInputParams() *ActionVpsCloneRequest {
	return &ActionVpsCloneRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsCloneInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AddressLocation") {
			if inv.IsParameterNil("AddressLocation") {
				ret["address_location"] = nil
			} else {
				ret["address_location"] = inv.Input.AddressLocation
			}
		}
		if inv.IsParameterSelected("DatasetPlans") {
			ret["dataset_plans"] = inv.Input.DatasetPlans
		}
		if inv.IsParameterSelected("Environment") {
			if inv.IsParameterNil("Environment") {
				ret["environment"] = nil
			} else {
				ret["environment"] = inv.Input.Environment
			}
		}
		if inv.IsParameterSelected("Features") {
			ret["features"] = inv.Input.Features
		}
		if inv.IsParameterSelected("Hostname") {
			ret["hostname"] = inv.Input.Hostname
		}
		if inv.IsParameterSelected("KeepSnapshots") {
			ret["keep_snapshots"] = inv.Input.KeepSnapshots
		}
		if inv.IsParameterSelected("Location") {
			if inv.IsParameterNil("Location") {
				ret["location"] = nil
			} else {
				ret["location"] = inv.Input.Location
			}
		}
		if inv.IsParameterSelected("Node") {
			if inv.IsParameterNil("Node") {
				ret["node"] = nil
			} else {
				ret["node"] = inv.Input.Node
			}
		}
		if inv.IsParameterSelected("Platform") {
			ret["platform"] = inv.Input.Platform
		}
		if inv.IsParameterSelected("Resources") {
			ret["resources"] = inv.Input.Resources
		}
		if inv.IsParameterSelected("Stop") {
			ret["stop"] = inv.Input.Stop
		}
		if inv.IsParameterSelected("Subdatasets") {
			ret["subdatasets"] = inv.Input.Subdatasets
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["user"] = nil
			} else {
				ret["user"] = inv.Input.User
			}
		}
	}

	return ret
}

func (inv *ActionVpsCloneInvocation) makeMetaInputParams() map[string]interface{} {
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
