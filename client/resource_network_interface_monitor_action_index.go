package client

import ()

// ActionNetworkInterfaceMonitorIndex is a type for action Network_interface_monitor#Index
type ActionNetworkInterfaceMonitorIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkInterfaceMonitorIndex(client *Client) *ActionNetworkInterfaceMonitorIndex {
	return &ActionNetworkInterfaceMonitorIndex{
		Client: client,
	}
}

// ActionNetworkInterfaceMonitorIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkInterfaceMonitorIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexMetaGlobalInput) SetCount(value bool) *ActionNetworkInterfaceMonitorIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexMetaGlobalInput) SetIncludes(value string) *ActionNetworkInterfaceMonitorIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexMetaGlobalInput) SetNo(value bool) *ActionNetworkInterfaceMonitorIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceMonitorIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceMonitorIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkInterfaceMonitorIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkInterfaceMonitorIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceMonitorIndexInput is a type for action input parameters
type ActionNetworkInterfaceMonitorIndexInput struct {
	Environment      int64  `json:"environment"`
	FromId           int64  `json:"from_id"`
	Limit            int64  `json:"limit"`
	Location         int64  `json:"location"`
	NetworkInterface int64  `json:"network_interface"`
	Node             int64  `json:"node"`
	Order            string `json:"order"`
	User             int64  `json:"user"`
	Vps              int64  `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetEnvironment(value int64) *ActionNetworkInterfaceMonitorIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetEnvironmentNil(set bool) *ActionNetworkInterfaceMonitorIndexInput {
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

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetFromId(value int64) *ActionNetworkInterfaceMonitorIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetLimit(value int64) *ActionNetworkInterfaceMonitorIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetLocation(value int64) *ActionNetworkInterfaceMonitorIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetLocationNil(set bool) *ActionNetworkInterfaceMonitorIndexInput {
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

// SetNetworkInterface sets parameter NetworkInterface to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetNetworkInterface(value int64) *ActionNetworkInterfaceMonitorIndexInput {
	in.NetworkInterface = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkInterfaceNil(false)
	in._selectedParameters["NetworkInterface"] = nil
	return in
}

// SetNetworkInterfaceNil sets parameter NetworkInterface to nil and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetNetworkInterfaceNil(set bool) *ActionNetworkInterfaceMonitorIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["NetworkInterface"] = nil
		in.SelectParameters("NetworkInterface")
	} else {
		delete(in._nilParameters, "NetworkInterface")
	}
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetNode(value int64) *ActionNetworkInterfaceMonitorIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetNodeNil(set bool) *ActionNetworkInterfaceMonitorIndexInput {
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

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetOrder(value string) *ActionNetworkInterfaceMonitorIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetUser(value int64) *ActionNetworkInterfaceMonitorIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetUserNil(set bool) *ActionNetworkInterfaceMonitorIndexInput {
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

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetVps(value int64) *ActionNetworkInterfaceMonitorIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionNetworkInterfaceMonitorIndexInput) SetVpsNil(set bool) *ActionNetworkInterfaceMonitorIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Vps"] = nil
		in.SelectParameters("Vps")
	} else {
		delete(in._nilParameters, "Vps")
	}
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceMonitorIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceMonitorIndexInput) SelectParameters(params ...string) *ActionNetworkInterfaceMonitorIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNetworkInterfaceMonitorIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceMonitorIndexInput) UnselectParameters(params ...string) *ActionNetworkInterfaceMonitorIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNetworkInterfaceMonitorIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceMonitorIndexOutput is a type for action output parameters
type ActionNetworkInterfaceMonitorIndexOutput struct {
	Bytes            int64                             `json:"bytes"`
	BytesIn          int64                             `json:"bytes_in"`
	BytesOut         int64                             `json:"bytes_out"`
	Delta            int64                             `json:"delta"`
	Id               int64                             `json:"id"`
	NetworkInterface *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	Packets          int64                             `json:"packets"`
	PacketsIn        int64                             `json:"packets_in"`
	PacketsOut       int64                             `json:"packets_out"`
	UpdatedAt        string                            `json:"updated_at"`
}

// Type for action response, including envelope
type ActionNetworkInterfaceMonitorIndexResponse struct {
	Action *ActionNetworkInterfaceMonitorIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NetworkInterfaceMonitors []*ActionNetworkInterfaceMonitorIndexOutput `json:"network_interface_monitors"`
	}

	// Action output without the namespace
	Output []*ActionNetworkInterfaceMonitorIndexOutput
}

// Prepare the action for invocation
func (action *ActionNetworkInterfaceMonitorIndex) Prepare() *ActionNetworkInterfaceMonitorIndexInvocation {
	return &ActionNetworkInterfaceMonitorIndexInvocation{
		Action: action,
		Path:   "/v7.0/network_interface_monitors",
	}
}

// ActionNetworkInterfaceMonitorIndexInvocation is used to configure action for invocation
type ActionNetworkInterfaceMonitorIndexInvocation struct {
	// Pointer to the action
	Action *ActionNetworkInterfaceMonitorIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkInterfaceMonitorIndexInput
	// Global meta input parameters
	MetaInput *ActionNetworkInterfaceMonitorIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkInterfaceMonitorIndexInvocation) NewInput() *ActionNetworkInterfaceMonitorIndexInput {
	inv.Input = &ActionNetworkInterfaceMonitorIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkInterfaceMonitorIndexInvocation) SetInput(input *ActionNetworkInterfaceMonitorIndexInput) *ActionNetworkInterfaceMonitorIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkInterfaceMonitorIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceMonitorIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkInterfaceMonitorIndexInvocation) NewMetaInput() *ActionNetworkInterfaceMonitorIndexMetaGlobalInput {
	inv.MetaInput = &ActionNetworkInterfaceMonitorIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkInterfaceMonitorIndexInvocation) SetMetaInput(input *ActionNetworkInterfaceMonitorIndexMetaGlobalInput) *ActionNetworkInterfaceMonitorIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkInterfaceMonitorIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceMonitorIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkInterfaceMonitorIndexInvocation) Call() (*ActionNetworkInterfaceMonitorIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNetworkInterfaceMonitorIndexInvocation) callAsQuery() (*ActionNetworkInterfaceMonitorIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNetworkInterfaceMonitorIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NetworkInterfaceMonitors
	}
	return resp, err
}

func (inv *ActionNetworkInterfaceMonitorIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["network_interface_monitor[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("FromId") {
			ret["network_interface_monitor[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["network_interface_monitor[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["network_interface_monitor[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("NetworkInterface") {
			ret["network_interface_monitor[network_interface]"] = convertInt64ToString(inv.Input.NetworkInterface)
		}
		if inv.IsParameterSelected("Node") {
			ret["network_interface_monitor[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Order") {
			ret["network_interface_monitor[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("User") {
			ret["network_interface_monitor[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["network_interface_monitor[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionNetworkInterfaceMonitorIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
