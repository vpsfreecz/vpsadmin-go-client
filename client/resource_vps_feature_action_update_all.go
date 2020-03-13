package client

import (
	"strings"
)

// ActionVpsFeatureUpdateAll is a type for action Vps.Feature#Update_all
type ActionVpsFeatureUpdateAll struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsFeatureUpdateAll(client *Client) *ActionVpsFeatureUpdateAll {
	return &ActionVpsFeatureUpdateAll{
		Client: client,
	}
}

// ActionVpsFeatureUpdateAllMetaGlobalInput is a type for action global meta input parameters
type ActionVpsFeatureUpdateAllMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllMetaGlobalInput) SetNo(value bool) *ActionVpsFeatureUpdateAllMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllMetaGlobalInput) SetIncludes(value string) *ActionVpsFeatureUpdateAllMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsFeatureUpdateAllMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsFeatureUpdateAllMetaGlobalInput) SelectParameters(params ...string) *ActionVpsFeatureUpdateAllMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsFeatureUpdateAllMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsFeatureUpdateAllInput is a type for action input parameters
type ActionVpsFeatureUpdateAllInput struct {
	Iptables bool `json:"iptables"`
	Tun bool `json:"tun"`
	Fuse bool `json:"fuse"`
	Nfs bool `json:"nfs"`
	Ppp bool `json:"ppp"`
	Bridge bool `json:"bridge"`
	Kvm bool `json:"kvm"`
	Lxc bool `json:"lxc"`
	Docker bool `json:"docker"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIptables sets parameter Iptables to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllInput) SetIptables(value bool) *ActionVpsFeatureUpdateAllInput {
	in.Iptables = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Iptables"] = nil
	return in
}
// SetTun sets parameter Tun to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllInput) SetTun(value bool) *ActionVpsFeatureUpdateAllInput {
	in.Tun = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Tun"] = nil
	return in
}
// SetFuse sets parameter Fuse to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllInput) SetFuse(value bool) *ActionVpsFeatureUpdateAllInput {
	in.Fuse = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Fuse"] = nil
	return in
}
// SetNfs sets parameter Nfs to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllInput) SetNfs(value bool) *ActionVpsFeatureUpdateAllInput {
	in.Nfs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Nfs"] = nil
	return in
}
// SetPpp sets parameter Ppp to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllInput) SetPpp(value bool) *ActionVpsFeatureUpdateAllInput {
	in.Ppp = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ppp"] = nil
	return in
}
// SetBridge sets parameter Bridge to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllInput) SetBridge(value bool) *ActionVpsFeatureUpdateAllInput {
	in.Bridge = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Bridge"] = nil
	return in
}
// SetKvm sets parameter Kvm to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllInput) SetKvm(value bool) *ActionVpsFeatureUpdateAllInput {
	in.Kvm = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Kvm"] = nil
	return in
}
// SetLxc sets parameter Lxc to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllInput) SetLxc(value bool) *ActionVpsFeatureUpdateAllInput {
	in.Lxc = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lxc"] = nil
	return in
}
// SetDocker sets parameter Docker to value and selects it for sending
func (in *ActionVpsFeatureUpdateAllInput) SetDocker(value bool) *ActionVpsFeatureUpdateAllInput {
	in.Docker = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Docker"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsFeatureUpdateAllInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsFeatureUpdateAllInput) SelectParameters(params ...string) *ActionVpsFeatureUpdateAllInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsFeatureUpdateAllInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsFeatureUpdateAllRequest is a type for the entire action request
type ActionVpsFeatureUpdateAllRequest struct {
	Feature map[string]interface{} `json:"feature"`
	Meta map[string]interface{} `json:"_meta"`
}


// ActionVpsFeatureUpdateAllMetaGlobalOutput is a type for global output metadata parameters
type ActionVpsFeatureUpdateAllMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionVpsFeatureUpdateAllResponse struct {
	Action *ActionVpsFeatureUpdateAll `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionVpsFeatureUpdateAllMetaGlobalOutput `json:"_meta"`
	}
}


// Prepare the action for invocation
func (action *ActionVpsFeatureUpdateAll) Prepare() *ActionVpsFeatureUpdateAllInvocation {
	return &ActionVpsFeatureUpdateAllInvocation{
		Action: action,
		Path: "/v6.0/vpses/{vps_id}/features/update_all",
	}
}

// ActionVpsFeatureUpdateAllInvocation is used to configure action for invocation
type ActionVpsFeatureUpdateAllInvocation struct {
	// Pointer to the action
	Action *ActionVpsFeatureUpdateAll

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsFeatureUpdateAllInput
	// Global meta input parameters
	MetaInput *ActionVpsFeatureUpdateAllMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsFeatureUpdateAllInvocation) SetPathParamInt(param string, value int64) *ActionVpsFeatureUpdateAllInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsFeatureUpdateAllInvocation) SetPathParamString(param string, value string) *ActionVpsFeatureUpdateAllInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsFeatureUpdateAllInvocation) NewInput() *ActionVpsFeatureUpdateAllInput {
	inv.Input = &ActionVpsFeatureUpdateAllInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsFeatureUpdateAllInvocation) SetInput(input *ActionVpsFeatureUpdateAllInput) *ActionVpsFeatureUpdateAllInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsFeatureUpdateAllInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsFeatureUpdateAllInvocation) NewMetaInput() *ActionVpsFeatureUpdateAllMetaGlobalInput {
	inv.MetaInput = &ActionVpsFeatureUpdateAllMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsFeatureUpdateAllInvocation) SetMetaInput(input *ActionVpsFeatureUpdateAllMetaGlobalInput) *ActionVpsFeatureUpdateAllInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsFeatureUpdateAllInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsFeatureUpdateAllInvocation) Call() (*ActionVpsFeatureUpdateAllResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsFeatureUpdateAllInvocation) callAsBody() (*ActionVpsFeatureUpdateAllResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsFeatureUpdateAllResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionVpsFeatureUpdateAllResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionVpsFeatureUpdateAllResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionVpsFeatureUpdateAllResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionVpsFeatureUpdateAllResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionVpsFeatureUpdateAllResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionVpsFeatureUpdateAllInvocation) makeAllInputParams() *ActionVpsFeatureUpdateAllRequest {
	return &ActionVpsFeatureUpdateAllRequest{
		Feature: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsFeatureUpdateAllInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Iptables") {
			ret["iptables"] = inv.Input.Iptables
		}
		if inv.IsParameterSelected("Tun") {
			ret["tun"] = inv.Input.Tun
		}
		if inv.IsParameterSelected("Fuse") {
			ret["fuse"] = inv.Input.Fuse
		}
		if inv.IsParameterSelected("Nfs") {
			ret["nfs"] = inv.Input.Nfs
		}
		if inv.IsParameterSelected("Ppp") {
			ret["ppp"] = inv.Input.Ppp
		}
		if inv.IsParameterSelected("Bridge") {
			ret["bridge"] = inv.Input.Bridge
		}
		if inv.IsParameterSelected("Kvm") {
			ret["kvm"] = inv.Input.Kvm
		}
		if inv.IsParameterSelected("Lxc") {
			ret["lxc"] = inv.Input.Lxc
		}
		if inv.IsParameterSelected("Docker") {
			ret["docker"] = inv.Input.Docker
		}
	}

	return ret
}

func (inv *ActionVpsFeatureUpdateAllInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
