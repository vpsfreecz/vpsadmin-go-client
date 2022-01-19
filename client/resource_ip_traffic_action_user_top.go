package client

import ()

// ActionIpTrafficUserTop is a type for action Ip_traffic#User_top
type ActionIpTrafficUserTop struct {
	// Pointer to client
	Client *Client
}

func NewActionIpTrafficUserTop(client *Client) *ActionIpTrafficUserTop {
	return &ActionIpTrafficUserTop{
		Client: client,
	}
}

// ActionIpTrafficUserTopMetaGlobalInput is a type for action global meta input parameters
type ActionIpTrafficUserTopMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpTrafficUserTopMetaGlobalInput) SetCount(value bool) *ActionIpTrafficUserTopMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpTrafficUserTopMetaGlobalInput) SetIncludes(value string) *ActionIpTrafficUserTopMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpTrafficUserTopMetaGlobalInput) SetNo(value bool) *ActionIpTrafficUserTopMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpTrafficUserTopMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpTrafficUserTopMetaGlobalInput) SelectParameters(params ...string) *ActionIpTrafficUserTopMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpTrafficUserTopMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpTrafficUserTopInput is a type for action input parameters
type ActionIpTrafficUserTopInput struct {
	Accumulate  string `json:"accumulate"`
	Environment int64  `json:"environment"`
	From        string `json:"from"`
	IpVersion   int64  `json:"ip_version"`
	Limit       int64  `json:"limit"`
	Location    int64  `json:"location"`
	Month       int64  `json:"month"`
	Network     int64  `json:"network"`
	Node        int64  `json:"node"`
	Offset      int64  `json:"offset"`
	Protocol    string `json:"protocol"`
	Role        string `json:"role"`
	To          string `json:"to"`
	Year        int64  `json:"year"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetAccumulate sets parameter Accumulate to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetAccumulate(value string) *ActionIpTrafficUserTopInput {
	in.Accumulate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Accumulate"] = nil
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetEnvironment(value int64) *ActionIpTrafficUserTopInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetFrom(value string) *ActionIpTrafficUserTopInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetIpVersion sets parameter IpVersion to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetIpVersion(value int64) *ActionIpTrafficUserTopInput {
	in.IpVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpVersion"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetLimit(value int64) *ActionIpTrafficUserTopInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetLocation(value int64) *ActionIpTrafficUserTopInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}

// SetMonth sets parameter Month to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetMonth(value int64) *ActionIpTrafficUserTopInput {
	in.Month = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Month"] = nil
	return in
}

// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetNetwork(value int64) *ActionIpTrafficUserTopInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Network"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetNode(value int64) *ActionIpTrafficUserTopInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetOffset(value int64) *ActionIpTrafficUserTopInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetProtocol sets parameter Protocol to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetProtocol(value string) *ActionIpTrafficUserTopInput {
	in.Protocol = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Protocol"] = nil
	return in
}

// SetRole sets parameter Role to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetRole(value string) *ActionIpTrafficUserTopInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetTo(value string) *ActionIpTrafficUserTopInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SetYear sets parameter Year to value and selects it for sending
func (in *ActionIpTrafficUserTopInput) SetYear(value int64) *ActionIpTrafficUserTopInput {
	in.Year = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Year"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpTrafficUserTopInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpTrafficUserTopInput) SelectParameters(params ...string) *ActionIpTrafficUserTopInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpTrafficUserTopInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpTrafficUserTopOutput is a type for action output parameters
type ActionIpTrafficUserTopOutput struct {
	BytesIn    int64                 `json:"bytes_in"`
	BytesOut   int64                 `json:"bytes_out"`
	CreatedAt  string                `json:"created_at"`
	PacketsIn  int64                 `json:"packets_in"`
	PacketsOut int64                 `json:"packets_out"`
	User       *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionIpTrafficUserTopResponse struct {
	Action *ActionIpTrafficUserTop `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpTraffics []*ActionIpTrafficUserTopOutput `json:"ip_traffics"`
	}

	// Action output without the namespace
	Output []*ActionIpTrafficUserTopOutput
}

// Prepare the action for invocation
func (action *ActionIpTrafficUserTop) Prepare() *ActionIpTrafficUserTopInvocation {
	return &ActionIpTrafficUserTopInvocation{
		Action: action,
		Path:   "/v6.0/ip_traffics/user_top",
	}
}

// ActionIpTrafficUserTopInvocation is used to configure action for invocation
type ActionIpTrafficUserTopInvocation struct {
	// Pointer to the action
	Action *ActionIpTrafficUserTop

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpTrafficUserTopInput
	// Global meta input parameters
	MetaInput *ActionIpTrafficUserTopMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpTrafficUserTopInvocation) NewInput() *ActionIpTrafficUserTopInput {
	inv.Input = &ActionIpTrafficUserTopInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpTrafficUserTopInvocation) SetInput(input *ActionIpTrafficUserTopInput) *ActionIpTrafficUserTopInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpTrafficUserTopInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpTrafficUserTopInvocation) NewMetaInput() *ActionIpTrafficUserTopMetaGlobalInput {
	inv.MetaInput = &ActionIpTrafficUserTopMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpTrafficUserTopInvocation) SetMetaInput(input *ActionIpTrafficUserTopMetaGlobalInput) *ActionIpTrafficUserTopInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpTrafficUserTopInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpTrafficUserTopInvocation) Call() (*ActionIpTrafficUserTopResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIpTrafficUserTopInvocation) callAsQuery() (*ActionIpTrafficUserTopResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpTrafficUserTopResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpTraffics
	}
	return resp, err
}

func (inv *ActionIpTrafficUserTopInvocation) convertInputToQueryParams(ret map[string]string) {
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
		if inv.IsParameterSelected("Protocol") {
			ret["ip_traffic[protocol]"] = inv.Input.Protocol
		}
		if inv.IsParameterSelected("Role") {
			ret["ip_traffic[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("To") {
			ret["ip_traffic[to]"] = inv.Input.To
		}
		if inv.IsParameterSelected("Year") {
			ret["ip_traffic[year]"] = convertInt64ToString(inv.Input.Year)
		}
	}
}

func (inv *ActionIpTrafficUserTopInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
