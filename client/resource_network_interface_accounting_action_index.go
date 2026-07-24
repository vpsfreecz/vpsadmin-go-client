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
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
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
	Environment int64  "json:\"environment\""
	From        string "json:\"from\""
	FromBytes   int64  "json:\"from_bytes\""
	FromDate    string "json:\"from_date\""
	Limit       int64  "json:\"limit\""
	Location    int64  "json:\"location\""
	Month       int64  "json:\"month\""
	Node        int64  "json:\"node\""
	Order       string "json:\"order\""
	To          string "json:\"to\""
	User        int64  "json:\"user\""
	Vps         int64  "json:\"vps\""
	Year        int64  "json:\"year\""
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

	in._selectedParameters["Environment"] = nil
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

// SetFromBytes sets parameter FromBytes to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetFromBytes(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.FromBytes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromBytes"] = nil
	return in
}

// SetFromDate sets parameter FromDate to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetFromDate(value string) *ActionNetworkInterfaceAccountingIndexInput {
	in.FromDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromDate"] = nil
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

	in._selectedParameters["Location"] = nil
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

	in._selectedParameters["Node"] = nil
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

	in._selectedParameters["User"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionNetworkInterfaceAccountingIndexInput) SetVps(value int64) *ActionNetworkInterfaceAccountingIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
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
	Bytes            int64                             "json:\"bytes\""
	BytesIn          int64                             "json:\"bytes_in\""
	BytesOut         int64                             "json:\"bytes_out\""
	CreatedAt        string                            "json:\"created_at\""
	Month            int64                             "json:\"month\""
	NetworkInterface *ActionNetworkInterfaceShowOutput "json:\"network_interface\""
	Packets          int64                             "json:\"packets\""
	PacketsIn        int64                             "json:\"packets_in\""
	PacketsOut       int64                             "json:\"packets_out\""
	UpdatedAt        string                            "json:\"updated_at\""
	Year             int64                             "json:\"year\""
}

// Type for action response, including envelope
type ActionNetworkInterfaceAccountingIndexResponse struct {
	Action *ActionNetworkInterfaceAccountingIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NetworkInterfaceAccountings []*ActionNetworkInterfaceAccountingIndexOutput "json:\"network_interface_accountings\""
	}

	// Action output without the namespace
	Output []*ActionNetworkInterfaceAccountingIndexOutput
}

// Prepare the action for invocation
func (action *ActionNetworkInterfaceAccountingIndex) Prepare() *ActionNetworkInterfaceAccountingIndexInvocation {
	return &ActionNetworkInterfaceAccountingIndexInvocation{
		Action: action,
		Path:   "/v7.0/network_interface_accountings",
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

func (inv *ActionNetworkInterfaceAccountingIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			if !inv.IsParameterNil("Environment") {
				if inv.Input.Environment < 0 {
					verr.Add("environment", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("From") {
			if !inv.IsParameterNil("From") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.From)
				if !ok {
					verr.Add("from", "not a valid datetime")
				} else {
					inv.Input.From = normalized
				}
			}
		}
		if inv.IsParameterSelected("FromDate") {
			if !inv.IsParameterNil("FromDate") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.FromDate)
				if !ok {
					verr.Add("from_date", "not a valid datetime")
				} else {
					inv.Input.FromDate = normalized
				}
			}
		}
		if inv.IsParameterSelected("Location") {
			if !inv.IsParameterNil("Location") {
				if inv.Input.Location < 0 {
					verr.Add("location", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("To") {
			if !inv.IsParameterNil("To") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.To)
				if !ok {
					verr.Add("to", "not a valid datetime")
				} else {
					inv.Input.To = normalized
				}
			}
		}
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Vps") {
			if !inv.IsParameterNil("Vps") {
				if inv.Input.Vps < 0 {
					verr.Add("vps", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkInterfaceAccountingIndexInvocation) Call() (*ActionNetworkInterfaceAccountingIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNetworkInterfaceAccountingIndexInvocation) callAsQuery() (*ActionNetworkInterfaceAccountingIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNetworkInterfaceAccountingIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NetworkInterfaceAccountings
	}
	return resp, err
}

func (inv *ActionNetworkInterfaceAccountingIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
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
		if inv.IsParameterSelected("FromDate") {
			ret["network_interface_accounting[from_date]"] = inv.Input.FromDate
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

	return nil
}

func (inv *ActionNetworkInterfaceAccountingIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
