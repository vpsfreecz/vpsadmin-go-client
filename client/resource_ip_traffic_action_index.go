package client

import ()

// ActionIpTrafficIndex is a type for action Ip_traffic#Index
type ActionIpTrafficIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpTrafficIndex(client *Client) *ActionIpTrafficIndex {
	return &ActionIpTrafficIndex{
		Client: client,
	}
}

// ActionIpTrafficIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpTrafficIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpTrafficIndexMetaGlobalInput) SetCount(value bool) *ActionIpTrafficIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpTrafficIndexMetaGlobalInput) SetIncludes(value string) *ActionIpTrafficIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpTrafficIndexMetaGlobalInput) SetNo(value bool) *ActionIpTrafficIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpTrafficIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpTrafficIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpTrafficIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpTrafficIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpTrafficIndexInput is a type for action input parameters
type ActionIpTrafficIndexInput struct {
	Accumulate  string `json:"accumulate"`
	Environment int64  `json:"environment"`
	From        string `json:"from"`
	IpAddress   int64  `json:"ip_address"`
	IpVersion   int64  `json:"ip_version"`
	Limit       int64  `json:"limit"`
	Location    int64  `json:"location"`
	Month       int64  `json:"month"`
	Network     int64  `json:"network"`
	Node        int64  `json:"node"`
	Offset      int64  `json:"offset"`
	Order       string `json:"order"`
	Protocol    string `json:"protocol"`
	Role        string `json:"role"`
	To          string `json:"to"`
	User        int64  `json:"user"`
	Vps         int64  `json:"vps"`
	Year        int64  `json:"year"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAccumulate sets parameter Accumulate to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetAccumulate(value string) *ActionIpTrafficIndexInput {
	in.Accumulate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Accumulate"] = nil
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetEnvironment(value int64) *ActionIpTrafficIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionIpTrafficIndexInput) SetEnvironmentNil(set bool) *ActionIpTrafficIndexInput {
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

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetFrom(value string) *ActionIpTrafficIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetIpAddress sets parameter IpAddress to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetIpAddress(value int64) *ActionIpTrafficIndexInput {
	in.IpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetIpAddressNil(false)
	in._selectedParameters["IpAddress"] = nil
	return in
}

// SetIpAddressNil sets parameter IpAddress to nil and selects it for sending
func (in *ActionIpTrafficIndexInput) SetIpAddressNil(set bool) *ActionIpTrafficIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["IpAddress"] = nil
		in.SelectParameters("IpAddress")
	} else {
		delete(in._nilParameters, "IpAddress")
	}
	return in
}

// SetIpVersion sets parameter IpVersion to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetIpVersion(value int64) *ActionIpTrafficIndexInput {
	in.IpVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpVersion"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetLimit(value int64) *ActionIpTrafficIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetLocation(value int64) *ActionIpTrafficIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionIpTrafficIndexInput) SetLocationNil(set bool) *ActionIpTrafficIndexInput {
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

// SetMonth sets parameter Month to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetMonth(value int64) *ActionIpTrafficIndexInput {
	in.Month = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Month"] = nil
	return in
}

// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetNetwork(value int64) *ActionIpTrafficIndexInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkNil(false)
	in._selectedParameters["Network"] = nil
	return in
}

// SetNetworkNil sets parameter Network to nil and selects it for sending
func (in *ActionIpTrafficIndexInput) SetNetworkNil(set bool) *ActionIpTrafficIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Network"] = nil
		in.SelectParameters("Network")
	} else {
		delete(in._nilParameters, "Network")
	}
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetNode(value int64) *ActionIpTrafficIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionIpTrafficIndexInput) SetNodeNil(set bool) *ActionIpTrafficIndexInput {
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

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetOffset(value int64) *ActionIpTrafficIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetOrder(value string) *ActionIpTrafficIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetProtocol sets parameter Protocol to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetProtocol(value string) *ActionIpTrafficIndexInput {
	in.Protocol = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Protocol"] = nil
	return in
}

// SetRole sets parameter Role to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetRole(value string) *ActionIpTrafficIndexInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetTo(value string) *ActionIpTrafficIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetUser(value int64) *ActionIpTrafficIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionIpTrafficIndexInput) SetUserNil(set bool) *ActionIpTrafficIndexInput {
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
func (in *ActionIpTrafficIndexInput) SetVps(value int64) *ActionIpTrafficIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionIpTrafficIndexInput) SetVpsNil(set bool) *ActionIpTrafficIndexInput {
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

// SetYear sets parameter Year to value and selects it for sending
func (in *ActionIpTrafficIndexInput) SetYear(value int64) *ActionIpTrafficIndexInput {
	in.Year = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Year"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpTrafficIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpTrafficIndexInput) SelectParameters(params ...string) *ActionIpTrafficIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpTrafficIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpTrafficIndexInput) UnselectParameters(params ...string) *ActionIpTrafficIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpTrafficIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpTrafficIndexOutput is a type for action output parameters
type ActionIpTrafficIndexOutput struct {
	BytesIn    int64                      `json:"bytes_in"`
	BytesOut   int64                      `json:"bytes_out"`
	CreatedAt  string                     `json:"created_at"`
	Id         int64                      `json:"id"`
	IpAddress  *ActionIpAddressShowOutput `json:"ip_address"`
	PacketsIn  int64                      `json:"packets_in"`
	PacketsOut int64                      `json:"packets_out"`
	Protocol   string                     `json:"protocol"`
	Role       string                     `json:"role"`
	User       *ActionUserShowOutput      `json:"user"`
}

// Type for action response, including envelope
type ActionIpTrafficIndexResponse struct {
	Action *ActionIpTrafficIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpTraffics []*ActionIpTrafficIndexOutput `json:"ip_traffics"`
	}

	// Action output without the namespace
	Output []*ActionIpTrafficIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpTrafficIndex) Prepare() *ActionIpTrafficIndexInvocation {
	return &ActionIpTrafficIndexInvocation{
		Action: action,
		Path:   "/v6.0/ip_traffics",
	}
}

// ActionIpTrafficIndexInvocation is used to configure action for invocation
type ActionIpTrafficIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpTrafficIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpTrafficIndexInput
	// Global meta input parameters
	MetaInput *ActionIpTrafficIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpTrafficIndexInvocation) NewInput() *ActionIpTrafficIndexInput {
	inv.Input = &ActionIpTrafficIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpTrafficIndexInvocation) SetInput(input *ActionIpTrafficIndexInput) *ActionIpTrafficIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpTrafficIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpTrafficIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpTrafficIndexInvocation) NewMetaInput() *ActionIpTrafficIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpTrafficIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpTrafficIndexInvocation) SetMetaInput(input *ActionIpTrafficIndexMetaGlobalInput) *ActionIpTrafficIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpTrafficIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpTrafficIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpTrafficIndexInvocation) Call() (*ActionIpTrafficIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIpTrafficIndexInvocation) callAsQuery() (*ActionIpTrafficIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpTrafficIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpTraffics
	}
	return resp, err
}

func (inv *ActionIpTrafficIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Accumulate") {
			ret["ip_traffic[accumulate]"] = inv.Input.Accumulate
		}
		if inv.IsParameterSelected("Environment") {
			ret["ip_traffic[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("From") {
			ret["ip_traffic[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("IpAddress") {
			ret["ip_traffic[ip_address]"] = convertInt64ToString(inv.Input.IpAddress)
		}
		if inv.IsParameterSelected("IpVersion") {
			ret["ip_traffic[ip_version]"] = convertInt64ToString(inv.Input.IpVersion)
		}
		if inv.IsParameterSelected("Limit") {
			ret["ip_traffic[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["ip_traffic[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Month") {
			ret["ip_traffic[month]"] = convertInt64ToString(inv.Input.Month)
		}
		if inv.IsParameterSelected("Network") {
			ret["ip_traffic[network]"] = convertInt64ToString(inv.Input.Network)
		}
		if inv.IsParameterSelected("Node") {
			ret["ip_traffic[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Offset") {
			ret["ip_traffic[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Order") {
			ret["ip_traffic[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("Protocol") {
			ret["ip_traffic[protocol]"] = inv.Input.Protocol
		}
		if inv.IsParameterSelected("Role") {
			ret["ip_traffic[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("To") {
			ret["ip_traffic[to]"] = inv.Input.To
		}
		if inv.IsParameterSelected("User") {
			ret["ip_traffic[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["ip_traffic[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
		if inv.IsParameterSelected("Year") {
			ret["ip_traffic[year]"] = convertInt64ToString(inv.Input.Year)
		}
	}
}

func (inv *ActionIpTrafficIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
