package client

import ()

// ActionNetworkInterfaceAccountingIndex is a type for action Network_interface_accounting#Index
type ActionNetworkInterfaceAccountingIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkInterfaceAccountingIndex(client *Client) *ActionNetworkInterfaceAccountingIndex {
	return &ActionNetworkInterfaceAccountingIndex{
		Client: client,
	}
}

// ActionNetworkInterfaceAccountingIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkInterfaceAccountingIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexMetaGlobalInput) SetCount(value bool) *ActionNetworkInterfaceAccountingIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexMetaGlobalInput) SetIncludes(value string) *ActionNetworkInterfaceAccountingIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexMetaGlobalInput) SetNo(value bool) *ActionNetworkInterfaceAccountingIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceAccountingIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceAccountingIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkInterfaceAccountingIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkInterfaceAccountingIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceAccountingIndexInput is a type for action input parameters
type ActionNetworkInterfaceAccountingIndexInput struct {
	Environment int64  `json:"environment"`
	From        string `json:"from"`
	Limit       int64  `json:"limit"`
	Location    int64  `json:"location"`
	Month       int64  `json:"month"`
	Node        int64  `json:"node"`
	Offset      int64  `json:"offset"`
	Order       string `json:"order"`
	To          string `json:"to"`
	User        int64  `json:"user"`
	Vps         int64  `json:"vps"`
	Year        int64  `json:"year"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetEnvironment(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetEnvironmentNil(set bool) *ActionNetworkInterfaceAccountingIndexInput {
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
func (in *ActionNetworkInterfaceAccountingIndexInput) SetFrom(value string) *ActionNetworkInterfaceAccountingIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetLimit(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetLocation(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetLocationNil(set bool) *ActionNetworkInterfaceAccountingIndexInput {
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
func (in *ActionNetworkInterfaceAccountingIndexInput) SetMonth(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.Month = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Month"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetNode(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetNodeNil(set bool) *ActionNetworkInterfaceAccountingIndexInput {
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
func (in *ActionNetworkInterfaceAccountingIndexInput) SetOffset(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetOrder(value string) *ActionNetworkInterfaceAccountingIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetTo(value string) *ActionNetworkInterfaceAccountingIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetUser(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetUserNil(set bool) *ActionNetworkInterfaceAccountingIndexInput {
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
func (in *ActionNetworkInterfaceAccountingIndexInput) SetVps(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetVpsNil(set bool) *ActionNetworkInterfaceAccountingIndexInput {
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
func (in *ActionNetworkInterfaceAccountingIndexInput) SetYear(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.Year = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Year"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceAccountingIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceAccountingIndexInput) SelectParameters(params ...string) *ActionNetworkInterfaceAccountingIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNetworkInterfaceAccountingIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceAccountingIndexInput) UnselectParameters(params ...string) *ActionNetworkInterfaceAccountingIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNetworkInterfaceAccountingIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceAccountingIndexOutput is a type for action output parameters
type ActionNetworkInterfaceAccountingIndexOutput struct {
	BytesIn          int64                             `json:"bytes_in"`
	BytesOut         int64                             `json:"bytes_out"`
	CreatedAt        string                            `json:"created_at"`
	Month            int64                             `json:"month"`
	NetworkInterface *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	PacketsIn        int64                             `json:"packets_in"`
	PacketsOut       int64                             `json:"packets_out"`
	UpdatedAt        string                            `json:"updated_at"`
	Year             int64                             `json:"year"`
}

// Type for action response, including envelope
type ActionNetworkInterfaceAccountingIndexResponse struct {
	Action *ActionNetworkInterfaceAccountingIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NetworkInterfaceAccountings []*ActionNetworkInterfaceAccountingIndexOutput `json:"network_interface_accountings"`
	}

	// Action output without the namespace
	Output []*ActionNetworkInterfaceAccountingIndexOutput
}

// Prepare the action for invocation
func (action *ActionNetworkInterfaceAccountingIndex) Prepare() *ActionNetworkInterfaceAccountingIndexInvocation {
	return &ActionNetworkInterfaceAccountingIndexInvocation{
		Action: action,
		Path:   "/v6.0/network_interface_accountings",
	}
}

// ActionNetworkInterfaceAccountingIndexInvocation is used to configure action for invocation
type ActionNetworkInterfaceAccountingIndexInvocation struct {
	// Pointer to the action
	Action *ActionNetworkInterfaceAccountingIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkInterfaceAccountingIndexInput
	// Global meta input parameters
	MetaInput *ActionNetworkInterfaceAccountingIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) NewInput() *ActionNetworkInterfaceAccountingIndexInput {
	inv.Input = &ActionNetworkInterfaceAccountingIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) SetInput(input *ActionNetworkInterfaceAccountingIndexInput) *ActionNetworkInterfaceAccountingIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) NewMetaInput() *ActionNetworkInterfaceAccountingIndexMetaGlobalInput {
	inv.MetaInput = &ActionNetworkInterfaceAccountingIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) SetMetaInput(input *ActionNetworkInterfaceAccountingIndexMetaGlobalInput) *ActionNetworkInterfaceAccountingIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) Call() (*ActionNetworkInterfaceAccountingIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNetworkInterfaceAccountingIndexInvocation) callAsQuery() (*ActionNetworkInterfaceAccountingIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNetworkInterfaceAccountingIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NetworkInterfaceAccountings
	}
	return resp, err
}

func (inv *ActionNetworkInterfaceAccountingIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["network_interface_accounting[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("From") {
			ret["network_interface_accounting[from]"] = inv.Input.From
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
		if inv.IsParameterSelected("Offset") {
			ret["network_interface_accounting[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Order") {
			ret["network_interface_accounting[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("To") {
			ret["network_interface_accounting[to]"] = inv.Input.To
		}
		if inv.IsParameterSelected("User") {
			ret["network_interface_accounting[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["network_interface_accounting[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
		if inv.IsParameterSelected("Year") {
			ret["network_interface_accounting[year]"] = convertInt64ToString(inv.Input.Year)
		}
	}
}

func (inv *ActionNetworkInterfaceAccountingIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
