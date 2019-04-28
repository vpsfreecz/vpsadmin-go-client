package client

import (
	"strings"
)

// ActionNodeUpdate is a type for action Node#Update
type ActionNodeUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeUpdate(client *Client) *ActionNodeUpdate {
	return &ActionNodeUpdate{
		Client: client,
	}
}

// ActionNodeUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNodeUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeUpdateMetaGlobalInput) SetNo(value bool) *ActionNodeUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeUpdateMetaGlobalInput) SetIncludes(value string) *ActionNodeUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNodeUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeUpdateInput is a type for action input parameters
type ActionNodeUpdateInput struct {
	Name string `json:"name"`
	Type string `json:"type"`
	HypervisorType string `json:"hypervisor_type"`
	Location int64 `json:"location"`
	IpAddr string `json:"ip_addr"`
	NetInterface string `json:"net_interface"`
	MaxTx int64 `json:"max_tx"`
	MaxRx int64 `json:"max_rx"`
	Cpus int64 `json:"cpus"`
	TotalMemory int64 `json:"total_memory"`
	TotalSwap int64 `json:"total_swap"`
	MaxVps int64 `json:"max_vps"`
	VePrivate string `json:"ve_private"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeUpdateInput) SetName(value string) *ActionNodeUpdateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}
// SetType sets parameter Type to value and selects it for sending
func (in *ActionNodeUpdateInput) SetType(value string) *ActionNodeUpdateInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}
// SetHypervisorType sets parameter HypervisorType to value and selects it for sending
func (in *ActionNodeUpdateInput) SetHypervisorType(value string) *ActionNodeUpdateInput {
	in.HypervisorType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HypervisorType"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNodeUpdateInput) SetLocation(value int64) *ActionNodeUpdateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetIpAddr sets parameter IpAddr to value and selects it for sending
func (in *ActionNodeUpdateInput) SetIpAddr(value string) *ActionNodeUpdateInput {
	in.IpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpAddr"] = nil
	return in
}
// SetNetInterface sets parameter NetInterface to value and selects it for sending
func (in *ActionNodeUpdateInput) SetNetInterface(value string) *ActionNodeUpdateInput {
	in.NetInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NetInterface"] = nil
	return in
}
// SetMaxTx sets parameter MaxTx to value and selects it for sending
func (in *ActionNodeUpdateInput) SetMaxTx(value int64) *ActionNodeUpdateInput {
	in.MaxTx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxTx"] = nil
	return in
}
// SetMaxRx sets parameter MaxRx to value and selects it for sending
func (in *ActionNodeUpdateInput) SetMaxRx(value int64) *ActionNodeUpdateInput {
	in.MaxRx = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxRx"] = nil
	return in
}
// SetCpus sets parameter Cpus to value and selects it for sending
func (in *ActionNodeUpdateInput) SetCpus(value int64) *ActionNodeUpdateInput {
	in.Cpus = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Cpus"] = nil
	return in
}
// SetTotalMemory sets parameter TotalMemory to value and selects it for sending
func (in *ActionNodeUpdateInput) SetTotalMemory(value int64) *ActionNodeUpdateInput {
	in.TotalMemory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TotalMemory"] = nil
	return in
}
// SetTotalSwap sets parameter TotalSwap to value and selects it for sending
func (in *ActionNodeUpdateInput) SetTotalSwap(value int64) *ActionNodeUpdateInput {
	in.TotalSwap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TotalSwap"] = nil
	return in
}
// SetMaxVps sets parameter MaxVps to value and selects it for sending
func (in *ActionNodeUpdateInput) SetMaxVps(value int64) *ActionNodeUpdateInput {
	in.MaxVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxVps"] = nil
	return in
}
// SetVePrivate sets parameter VePrivate to value and selects it for sending
func (in *ActionNodeUpdateInput) SetVePrivate(value string) *ActionNodeUpdateInput {
	in.VePrivate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VePrivate"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeUpdateInput) SelectParameters(params ...string) *ActionNodeUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeUpdateRequest is a type for the entire action request
type ActionNodeUpdateRequest struct {
	Node map[string]interface{} `json:"node"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionNodeUpdateResponse struct {
	Action *ActionNodeUpdate `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionNodeUpdate) Prepare() *ActionNodeUpdateInvocation {
	return &ActionNodeUpdateInvocation{
		Action: action,
		Path: "/v5.0/nodes/:node_id",
	}
}

// ActionNodeUpdateInvocation is used to configure action for invocation
type ActionNodeUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNodeUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeUpdateInput
	// Global meta input parameters
	MetaInput *ActionNodeUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNodeUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeUpdateInvocation) SetPathParamString(param string, value string) *ActionNodeUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeUpdateInvocation) NewInput() *ActionNodeUpdateInput {
	inv.Input = &ActionNodeUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeUpdateInvocation) SetInput(input *ActionNodeUpdateInput) *ActionNodeUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeUpdateInvocation) NewMetaInput() *ActionNodeUpdateMetaGlobalInput {
	inv.MetaInput = &ActionNodeUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeUpdateInvocation) SetMetaInput(input *ActionNodeUpdateMetaGlobalInput) *ActionNodeUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeUpdateInvocation) Call() (*ActionNodeUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionNodeUpdateInvocation) callAsBody() (*ActionNodeUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNodeUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionNodeUpdateInvocation) makeAllInputParams() *ActionNodeUpdateRequest {
	return &ActionNodeUpdateRequest{
		Node: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNodeUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Type") {
			ret["type"] = inv.Input.Type
		}
		if inv.IsParameterSelected("HypervisorType") {
			ret["hypervisor_type"] = inv.Input.HypervisorType
		}
		if inv.IsParameterSelected("Location") {
			ret["location"] = inv.Input.Location
		}
		if inv.IsParameterSelected("IpAddr") {
			ret["ip_addr"] = inv.Input.IpAddr
		}
		if inv.IsParameterSelected("NetInterface") {
			ret["net_interface"] = inv.Input.NetInterface
		}
		if inv.IsParameterSelected("MaxTx") {
			ret["max_tx"] = inv.Input.MaxTx
		}
		if inv.IsParameterSelected("MaxRx") {
			ret["max_rx"] = inv.Input.MaxRx
		}
		if inv.IsParameterSelected("Cpus") {
			ret["cpus"] = inv.Input.Cpus
		}
		if inv.IsParameterSelected("TotalMemory") {
			ret["total_memory"] = inv.Input.TotalMemory
		}
		if inv.IsParameterSelected("TotalSwap") {
			ret["total_swap"] = inv.Input.TotalSwap
		}
		if inv.IsParameterSelected("MaxVps") {
			ret["max_vps"] = inv.Input.MaxVps
		}
		if inv.IsParameterSelected("VePrivate") {
			ret["ve_private"] = inv.Input.VePrivate
		}
	}

	return ret
}

func (inv *ActionNodeUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
