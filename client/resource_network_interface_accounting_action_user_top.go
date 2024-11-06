package client

import ()

// ActionNetworkInterfaceAccountingUserTop is a type for action Network_interface_accounting#User_top
type ActionNetworkInterfaceAccountingUserTop struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkInterfaceAccountingUserTop(client *Client) *ActionNetworkInterfaceAccountingUserTop {
	return &ActionNetworkInterfaceAccountingUserTop{
		Client: client,
	}
}

// ActionNetworkInterfaceAccountingUserTopMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkInterfaceAccountingUserTopMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput) SetCount(value bool) *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput) SetIncludes(value string) *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput) SetNo(value bool) *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceAccountingUserTopMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceAccountingUserTopInput is a type for action input parameters
type ActionNetworkInterfaceAccountingUserTopInput struct {
	Environment int64  `json:"environment"`
	From        string `json:"from"`
	FromBytes   int64  `json:"from_bytes"`
	Limit       int64  `json:"limit"`
	Location    int64  `json:"location"`
	Month       int64  `json:"month"`
	Node        int64  `json:"node"`
	To          string `json:"to"`
	Year        int64  `json:"year"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetEnvironment(value int64) *ActionNetworkInterfaceAccountingUserTopInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetEnvironmentNil(set bool) *ActionNetworkInterfaceAccountingUserTopInput {
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
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetFrom(value string) *ActionNetworkInterfaceAccountingUserTopInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromBytes sets parameter FromBytes to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetFromBytes(value int64) *ActionNetworkInterfaceAccountingUserTopInput {
	in.FromBytes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromBytes"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetLimit(value int64) *ActionNetworkInterfaceAccountingUserTopInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetLocation(value int64) *ActionNetworkInterfaceAccountingUserTopInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetLocationNil(set bool) *ActionNetworkInterfaceAccountingUserTopInput {
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
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetMonth(value int64) *ActionNetworkInterfaceAccountingUserTopInput {
	in.Month = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Month"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetNode(value int64) *ActionNetworkInterfaceAccountingUserTopInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetNodeNil(set bool) *ActionNetworkInterfaceAccountingUserTopInput {
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

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetTo(value string) *ActionNetworkInterfaceAccountingUserTopInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SetYear sets parameter Year to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingUserTopInput) SetYear(value int64) *ActionNetworkInterfaceAccountingUserTopInput {
	in.Year = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Year"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceAccountingUserTopInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceAccountingUserTopInput) SelectParameters(params ...string) *ActionNetworkInterfaceAccountingUserTopInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNetworkInterfaceAccountingUserTopInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceAccountingUserTopInput) UnselectParameters(params ...string) *ActionNetworkInterfaceAccountingUserTopInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNetworkInterfaceAccountingUserTopInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceAccountingUserTopOutput is a type for action output parameters
type ActionNetworkInterfaceAccountingUserTopOutput struct {
	Bytes      int64                 `json:"bytes"`
	BytesIn    int64                 `json:"bytes_in"`
	BytesOut   int64                 `json:"bytes_out"`
	Month      int64                 `json:"month"`
	Packets    int64                 `json:"packets"`
	PacketsIn  int64                 `json:"packets_in"`
	PacketsOut int64                 `json:"packets_out"`
	User       *ActionUserShowOutput `json:"user"`
	Year       int64                 `json:"year"`
}

// Type for action response, including envelope
type ActionNetworkInterfaceAccountingUserTopResponse struct {
	Action *ActionNetworkInterfaceAccountingUserTop `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NetworkInterfaceAccountings []*ActionNetworkInterfaceAccountingUserTopOutput `json:"network_interface_accountings"`
	}

	// Action output without the namespace
	Output []*ActionNetworkInterfaceAccountingUserTopOutput
}

// Prepare the action for invocation
func (action *ActionNetworkInterfaceAccountingUserTop) Prepare() *ActionNetworkInterfaceAccountingUserTopInvocation {
	return &ActionNetworkInterfaceAccountingUserTopInvocation{
		Action: action,
		Path:   "/v7.0/network_interface_accountings/user_top",
	}
}

// ActionNetworkInterfaceAccountingUserTopInvocation is used to configure action for invocation
type ActionNetworkInterfaceAccountingUserTopInvocation struct {
	// Pointer to the action
	Action *ActionNetworkInterfaceAccountingUserTop

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkInterfaceAccountingUserTopInput
	// Global meta input parameters
	MetaInput *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) NewInput() *ActionNetworkInterfaceAccountingUserTopInput {
	inv.Input = &ActionNetworkInterfaceAccountingUserTopInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) SetInput(input *ActionNetworkInterfaceAccountingUserTopInput) *ActionNetworkInterfaceAccountingUserTopInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) NewMetaInput() *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput {
	inv.MetaInput = &ActionNetworkInterfaceAccountingUserTopMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) SetMetaInput(input *ActionNetworkInterfaceAccountingUserTopMetaGlobalInput) *ActionNetworkInterfaceAccountingUserTopInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) Call() (*ActionNetworkInterfaceAccountingUserTopResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) callAsQuery() (*ActionNetworkInterfaceAccountingUserTopResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNetworkInterfaceAccountingUserTopResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NetworkInterfaceAccountings
	}
	return resp, err
}

func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["network_interface_accounting[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("From") {
			ret["network_interface_accounting[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromBytes") {
			ret["network_interface_accounting[from_bytes]"] = convertInt64ToString(inv.Input.FromBytes)
		}
		if inv.IsParameterSelected("Limit") {
			ret["network_interface_accounting[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["network_interface_accounting[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Month") {
			ret["network_interface_accounting[month]"] = convertInt64ToString(inv.Input.Month)
		}
		if inv.IsParameterSelected("Node") {
			ret["network_interface_accounting[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("To") {
			ret["network_interface_accounting[to]"] = inv.Input.To
		}
		if inv.IsParameterSelected("Year") {
			ret["network_interface_accounting[year]"] = convertInt64ToString(inv.Input.Year)
		}
	}
}

func (inv *ActionNetworkInterfaceAccountingUserTopInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
