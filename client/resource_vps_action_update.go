package client

import (
	"strings"
)

// ActionVpsUpdate is a type for action Vps#Update
type ActionVpsUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsUpdate(client *Client) *ActionVpsUpdate {
	return &ActionVpsUpdate{
		Client: client,
	}
}

// ActionVpsUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsUpdateMetaGlobalInput) SetIncludes(value string) *ActionVpsUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsUpdateMetaGlobalInput) SetNo(value bool) *ActionVpsUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUpdateInput is a type for action input parameters
type ActionVpsUpdateInput struct {
	AdminLockType              string `json:"admin_lock_type"`
	AdminOverride              bool   `json:"admin_override"`
	AllowAdminModifications    bool   `json:"allow_admin_modifications"`
	AutostartEnable            bool   `json:"autostart_enable"`
	AutostartPriority          int64  `json:"autostart_priority"`
	CgroupVersion              string `json:"cgroup_version"`
	ChangeReason               string `json:"change_reason"`
	Config                     string `json:"config"`
	Cpu                        int64  `json:"cpu"`
	CpuLimit                   int64  `json:"cpu_limit"`
	DnsResolver                int64  `json:"dns_resolver"`
	EnableNetwork              bool   `json:"enable_network"`
	EnableOsTemplateAutoUpdate bool   `json:"enable_os_template_auto_update"`
	ExpirationDate             string `json:"expiration_date"`
	Hostname                   string `json:"hostname"`
	Info                       string `json:"info"`
	ManageHostname             bool   `json:"manage_hostname"`
	MapMode                    string `json:"map_mode"`
	Memory                     int64  `json:"memory"`
	Node                       int64  `json:"node"`
	ObjectState                string `json:"object_state"`
	Onstartall                 bool   `json:"onstartall"`
	OsTemplate                 int64  `json:"os_template"`
	RemindAfterDate            string `json:"remind_after_date"`
	StartMenuTimeout           int64  `json:"start_menu_timeout"`
	Swap                       int64  `json:"swap"`
	User                       int64  `json:"user"`
	UserNamespaceMap           int64  `json:"user_namespace_map"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAdminLockType sets parameter AdminLockType to value and selects it for sending
func (in *ActionVpsUpdateInput) SetAdminLockType(value string) *ActionVpsUpdateInput {
	in.AdminLockType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AdminLockType"] = nil
	return in
}

// SetAdminOverride sets parameter AdminOverride to value and selects it for sending
func (in *ActionVpsUpdateInput) SetAdminOverride(value bool) *ActionVpsUpdateInput {
	in.AdminOverride = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AdminOverride"] = nil
	return in
}

// SetAllowAdminModifications sets parameter AllowAdminModifications to value and selects it for sending
func (in *ActionVpsUpdateInput) SetAllowAdminModifications(value bool) *ActionVpsUpdateInput {
	in.AllowAdminModifications = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AllowAdminModifications"] = nil
	return in
}

// SetAutostartEnable sets parameter AutostartEnable to value and selects it for sending
func (in *ActionVpsUpdateInput) SetAutostartEnable(value bool) *ActionVpsUpdateInput {
	in.AutostartEnable = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AutostartEnable"] = nil
	return in
}

// SetAutostartPriority sets parameter AutostartPriority to value and selects it for sending
func (in *ActionVpsUpdateInput) SetAutostartPriority(value int64) *ActionVpsUpdateInput {
	in.AutostartPriority = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AutostartPriority"] = nil
	return in
}

// SetCgroupVersion sets parameter CgroupVersion to value and selects it for sending
func (in *ActionVpsUpdateInput) SetCgroupVersion(value string) *ActionVpsUpdateInput {
	in.CgroupVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CgroupVersion"] = nil
	return in
}

// SetChangeReason sets parameter ChangeReason to value and selects it for sending
func (in *ActionVpsUpdateInput) SetChangeReason(value string) *ActionVpsUpdateInput {
	in.ChangeReason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChangeReason"] = nil
	return in
}

// SetConfig sets parameter Config to value and selects it for sending
func (in *ActionVpsUpdateInput) SetConfig(value string) *ActionVpsUpdateInput {
	in.Config = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Config"] = nil
	return in
}

// SetCpu sets parameter Cpu to value and selects it for sending
func (in *ActionVpsUpdateInput) SetCpu(value int64) *ActionVpsUpdateInput {
	in.Cpu = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Cpu"] = nil
	return in
}

// SetCpuLimit sets parameter CpuLimit to value and selects it for sending
func (in *ActionVpsUpdateInput) SetCpuLimit(value int64) *ActionVpsUpdateInput {
	in.CpuLimit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CpuLimit"] = nil
	return in
}

// SetDnsResolver sets parameter DnsResolver to value and selects it for sending
func (in *ActionVpsUpdateInput) SetDnsResolver(value int64) *ActionVpsUpdateInput {
	in.DnsResolver = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsResolverNil(false)
	in._selectedParameters["DnsResolver"] = nil
	return in
}

// SetDnsResolverNil sets parameter DnsResolver to nil and selects it for sending
func (in *ActionVpsUpdateInput) SetDnsResolverNil(set bool) *ActionVpsUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["DnsResolver"] = nil
		in.SelectParameters("DnsResolver")
	} else {
		delete(in._nilParameters, "DnsResolver")
	}
	return in
}

// SetEnableNetwork sets parameter EnableNetwork to value and selects it for sending
func (in *ActionVpsUpdateInput) SetEnableNetwork(value bool) *ActionVpsUpdateInput {
	in.EnableNetwork = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableNetwork"] = nil
	return in
}

// SetEnableOsTemplateAutoUpdate sets parameter EnableOsTemplateAutoUpdate to value and selects it for sending
func (in *ActionVpsUpdateInput) SetEnableOsTemplateAutoUpdate(value bool) *ActionVpsUpdateInput {
	in.EnableOsTemplateAutoUpdate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableOsTemplateAutoUpdate"] = nil
	return in
}

// SetExpirationDate sets parameter ExpirationDate to value and selects it for sending
func (in *ActionVpsUpdateInput) SetExpirationDate(value string) *ActionVpsUpdateInput {
	in.ExpirationDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpirationDate"] = nil
	return in
}

// SetHostname sets parameter Hostname to value and selects it for sending
func (in *ActionVpsUpdateInput) SetHostname(value string) *ActionVpsUpdateInput {
	in.Hostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Hostname"] = nil
	return in
}

// SetInfo sets parameter Info to value and selects it for sending
func (in *ActionVpsUpdateInput) SetInfo(value string) *ActionVpsUpdateInput {
	in.Info = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Info"] = nil
	return in
}

// SetManageHostname sets parameter ManageHostname to value and selects it for sending
func (in *ActionVpsUpdateInput) SetManageHostname(value bool) *ActionVpsUpdateInput {
	in.ManageHostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ManageHostname"] = nil
	return in
}

// SetMapMode sets parameter MapMode to value and selects it for sending
func (in *ActionVpsUpdateInput) SetMapMode(value string) *ActionVpsUpdateInput {
	in.MapMode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MapMode"] = nil
	return in
}

// SetMemory sets parameter Memory to value and selects it for sending
func (in *ActionVpsUpdateInput) SetMemory(value int64) *ActionVpsUpdateInput {
	in.Memory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Memory"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsUpdateInput) SetNode(value int64) *ActionVpsUpdateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionVpsUpdateInput) SetNodeNil(set bool) *ActionVpsUpdateInput {
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

// SetObjectState sets parameter ObjectState to value and selects it for sending
func (in *ActionVpsUpdateInput) SetObjectState(value string) *ActionVpsUpdateInput {
	in.ObjectState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectState"] = nil
	return in
}

// SetOnstartall sets parameter Onstartall to value and selects it for sending
func (in *ActionVpsUpdateInput) SetOnstartall(value bool) *ActionVpsUpdateInput {
	in.Onstartall = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Onstartall"] = nil
	return in
}

// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionVpsUpdateInput) SetOsTemplate(value int64) *ActionVpsUpdateInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetOsTemplateNil(false)
	in._selectedParameters["OsTemplate"] = nil
	return in
}

// SetOsTemplateNil sets parameter OsTemplate to nil and selects it for sending
func (in *ActionVpsUpdateInput) SetOsTemplateNil(set bool) *ActionVpsUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["OsTemplate"] = nil
		in.SelectParameters("OsTemplate")
	} else {
		delete(in._nilParameters, "OsTemplate")
	}
	return in
}

// SetRemindAfterDate sets parameter RemindAfterDate to value and selects it for sending
func (in *ActionVpsUpdateInput) SetRemindAfterDate(value string) *ActionVpsUpdateInput {
	in.RemindAfterDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RemindAfterDate"] = nil
	return in
}

// SetStartMenuTimeout sets parameter StartMenuTimeout to value and selects it for sending
func (in *ActionVpsUpdateInput) SetStartMenuTimeout(value int64) *ActionVpsUpdateInput {
	in.StartMenuTimeout = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["StartMenuTimeout"] = nil
	return in
}

// SetSwap sets parameter Swap to value and selects it for sending
func (in *ActionVpsUpdateInput) SetSwap(value int64) *ActionVpsUpdateInput {
	in.Swap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Swap"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsUpdateInput) SetUser(value int64) *ActionVpsUpdateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionVpsUpdateInput) SetUserNil(set bool) *ActionVpsUpdateInput {
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

// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionVpsUpdateInput) SetUserNamespaceMap(value int64) *ActionVpsUpdateInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNamespaceMapNil(false)
	in._selectedParameters["UserNamespaceMap"] = nil
	return in
}

// SetUserNamespaceMapNil sets parameter UserNamespaceMap to nil and selects it for sending
func (in *ActionVpsUpdateInput) SetUserNamespaceMapNil(set bool) *ActionVpsUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["UserNamespaceMap"] = nil
		in.SelectParameters("UserNamespaceMap")
	} else {
		delete(in._nilParameters, "UserNamespaceMap")
	}
	return in
}

// SelectParameters sets parameters from ActionVpsUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUpdateInput) SelectParameters(params ...string) *ActionVpsUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsUpdateInput) UnselectParameters(params ...string) *ActionVpsUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUpdateRequest is a type for the entire action request
type ActionVpsUpdateRequest struct {
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsUpdateResponse struct {
	Action *ActionVpsUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsUpdateMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionVpsUpdate) Prepare() *ActionVpsUpdateInvocation {
	return &ActionVpsUpdateInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}",
	}
}

// ActionVpsUpdateInvocation is used to configure action for invocation
type ActionVpsUpdateInvocation struct {
	// Pointer to the action
	Action *ActionVpsUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsUpdateInput
	// Global meta input parameters
	MetaInput *ActionVpsUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsUpdateInvocation) SetPathParamInt(param string, value int64) *ActionVpsUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsUpdateInvocation) SetPathParamString(param string, value string) *ActionVpsUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsUpdateInvocation) NewInput() *ActionVpsUpdateInput {
	inv.Input = &ActionVpsUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsUpdateInvocation) SetInput(input *ActionVpsUpdateInput) *ActionVpsUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsUpdateInvocation) NewMetaInput() *ActionVpsUpdateMetaGlobalInput {
	inv.MetaInput = &ActionVpsUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsUpdateInvocation) SetMetaInput(input *ActionVpsUpdateMetaGlobalInput) *ActionVpsUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsUpdateInvocation) Call() (*ActionVpsUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsUpdateInvocation) callAsBody() (*ActionVpsUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionVpsUpdateInvocation) makeAllInputParams() *ActionVpsUpdateRequest {
	return &ActionVpsUpdateRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AdminLockType") {
			ret["admin_lock_type"] = inv.Input.AdminLockType
		}
		if inv.IsParameterSelected("AdminOverride") {
			ret["admin_override"] = inv.Input.AdminOverride
		}
		if inv.IsParameterSelected("AllowAdminModifications") {
			ret["allow_admin_modifications"] = inv.Input.AllowAdminModifications
		}
		if inv.IsParameterSelected("AutostartEnable") {
			ret["autostart_enable"] = inv.Input.AutostartEnable
		}
		if inv.IsParameterSelected("AutostartPriority") {
			ret["autostart_priority"] = inv.Input.AutostartPriority
		}
		if inv.IsParameterSelected("CgroupVersion") {
			ret["cgroup_version"] = inv.Input.CgroupVersion
		}
		if inv.IsParameterSelected("ChangeReason") {
			ret["change_reason"] = inv.Input.ChangeReason
		}
		if inv.IsParameterSelected("Config") {
			ret["config"] = inv.Input.Config
		}
		if inv.IsParameterSelected("Cpu") {
			ret["cpu"] = inv.Input.Cpu
		}
		if inv.IsParameterSelected("CpuLimit") {
			ret["cpu_limit"] = inv.Input.CpuLimit
		}
		if inv.IsParameterSelected("DnsResolver") {
			if inv.IsParameterNil("DnsResolver") {
				ret["dns_resolver"] = nil
			} else {
				ret["dns_resolver"] = inv.Input.DnsResolver
			}
		}
		if inv.IsParameterSelected("EnableNetwork") {
			ret["enable_network"] = inv.Input.EnableNetwork
		}
		if inv.IsParameterSelected("EnableOsTemplateAutoUpdate") {
			ret["enable_os_template_auto_update"] = inv.Input.EnableOsTemplateAutoUpdate
		}
		if inv.IsParameterSelected("ExpirationDate") {
			ret["expiration_date"] = inv.Input.ExpirationDate
		}
		if inv.IsParameterSelected("Hostname") {
			ret["hostname"] = inv.Input.Hostname
		}
		if inv.IsParameterSelected("Info") {
			ret["info"] = inv.Input.Info
		}
		if inv.IsParameterSelected("ManageHostname") {
			ret["manage_hostname"] = inv.Input.ManageHostname
		}
		if inv.IsParameterSelected("MapMode") {
			ret["map_mode"] = inv.Input.MapMode
		}
		if inv.IsParameterSelected("Memory") {
			ret["memory"] = inv.Input.Memory
		}
		if inv.IsParameterSelected("Node") {
			if inv.IsParameterNil("Node") {
				ret["node"] = nil
			} else {
				ret["node"] = inv.Input.Node
			}
		}
		if inv.IsParameterSelected("ObjectState") {
			ret["object_state"] = inv.Input.ObjectState
		}
		if inv.IsParameterSelected("Onstartall") {
			ret["onstartall"] = inv.Input.Onstartall
		}
		if inv.IsParameterSelected("OsTemplate") {
			if inv.IsParameterNil("OsTemplate") {
				ret["os_template"] = nil
			} else {
				ret["os_template"] = inv.Input.OsTemplate
			}
		}
		if inv.IsParameterSelected("RemindAfterDate") {
			ret["remind_after_date"] = inv.Input.RemindAfterDate
		}
		if inv.IsParameterSelected("StartMenuTimeout") {
			ret["start_menu_timeout"] = inv.Input.StartMenuTimeout
		}
		if inv.IsParameterSelected("Swap") {
			ret["swap"] = inv.Input.Swap
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["user"] = nil
			} else {
				ret["user"] = inv.Input.User
			}
		}
		if inv.IsParameterSelected("UserNamespaceMap") {
			if inv.IsParameterNil("UserNamespaceMap") {
				ret["user_namespace_map"] = nil
			} else {
				ret["user_namespace_map"] = inv.Input.UserNamespaceMap
			}
		}
	}

	return ret
}

func (inv *ActionVpsUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
