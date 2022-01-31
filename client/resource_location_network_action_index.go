package client

import ()

// ActionLocationNetworkIndex is a type for action Location_network#Index
type ActionLocationNetworkIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationNetworkIndex(client *Client) *ActionLocationNetworkIndex {
	return &ActionLocationNetworkIndex{
		Client: client,
	}
}

// ActionLocationNetworkIndexMetaGlobalInput is a type for action global meta input parameters
type ActionLocationNetworkIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionLocationNetworkIndexMetaGlobalInput) SetCount(value bool) *ActionLocationNetworkIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationNetworkIndexMetaGlobalInput) SetIncludes(value string) *ActionLocationNetworkIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationNetworkIndexMetaGlobalInput) SetNo(value bool) *ActionLocationNetworkIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationNetworkIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationNetworkIndexMetaGlobalInput) SelectParameters(params ...string) *ActionLocationNetworkIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationNetworkIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationNetworkIndexInput is a type for action input parameters
type ActionLocationNetworkIndexInput struct {
	Limit    int64 `json:"limit"`
	Location int64 `json:"location"`
	Network  int64 `json:"network"`
	Offset   int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionLocationNetworkIndexInput) SetLimit(value int64) *ActionLocationNetworkIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionLocationNetworkIndexInput) SetLocation(value int64) *ActionLocationNetworkIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionLocationNetworkIndexInput) SetLocationNil(set bool) *ActionLocationNetworkIndexInput {
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

// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionLocationNetworkIndexInput) SetNetwork(value int64) *ActionLocationNetworkIndexInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkNil(false)
	in._selectedParameters["Network"] = nil
	return in
}

// SetNetworkNil sets parameter Network to nil and selects it for sending
func (in *ActionLocationNetworkIndexInput) SetNetworkNil(set bool) *ActionLocationNetworkIndexInput {
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

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionLocationNetworkIndexInput) SetOffset(value int64) *ActionLocationNetworkIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationNetworkIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationNetworkIndexInput) SelectParameters(params ...string) *ActionLocationNetworkIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionLocationNetworkIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionLocationNetworkIndexInput) UnselectParameters(params ...string) *ActionLocationNetworkIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionLocationNetworkIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationNetworkIndexOutput is a type for action output parameters
type ActionLocationNetworkIndexOutput struct {
	Autopick bool                      `json:"autopick"`
	Id       int64                     `json:"id"`
	Location *ActionLocationShowOutput `json:"location"`
	Network  *ActionNetworkShowOutput  `json:"network"`
	Primary  bool                      `json:"primary"`
	Priority int64                     `json:"priority"`
	Userpick bool                      `json:"userpick"`
}

// Type for action response, including envelope
type ActionLocationNetworkIndexResponse struct {
	Action *ActionLocationNetworkIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		LocationNetworks []*ActionLocationNetworkIndexOutput `json:"location_networks"`
	}

	// Action output without the namespace
	Output []*ActionLocationNetworkIndexOutput
}

// Prepare the action for invocation
func (action *ActionLocationNetworkIndex) Prepare() *ActionLocationNetworkIndexInvocation {
	return &ActionLocationNetworkIndexInvocation{
		Action: action,
		Path:   "/v6.0/location_networks",
	}
}

// ActionLocationNetworkIndexInvocation is used to configure action for invocation
type ActionLocationNetworkIndexInvocation struct {
	// Pointer to the action
	Action *ActionLocationNetworkIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionLocationNetworkIndexInput
	// Global meta input parameters
	MetaInput *ActionLocationNetworkIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionLocationNetworkIndexInvocation) NewInput() *ActionLocationNetworkIndexInput {
	inv.Input = &ActionLocationNetworkIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionLocationNetworkIndexInvocation) SetInput(input *ActionLocationNetworkIndexInput) *ActionLocationNetworkIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionLocationNetworkIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionLocationNetworkIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLocationNetworkIndexInvocation) NewMetaInput() *ActionLocationNetworkIndexMetaGlobalInput {
	inv.MetaInput = &ActionLocationNetworkIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationNetworkIndexInvocation) SetMetaInput(input *ActionLocationNetworkIndexMetaGlobalInput) *ActionLocationNetworkIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationNetworkIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionLocationNetworkIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationNetworkIndexInvocation) Call() (*ActionLocationNetworkIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionLocationNetworkIndexInvocation) callAsQuery() (*ActionLocationNetworkIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionLocationNetworkIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.LocationNetworks
	}
	return resp, err
}

func (inv *ActionLocationNetworkIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["location_network[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["location_network[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Network") {
			ret["location_network[network]"] = convertInt64ToString(inv.Input.Network)
		}
		if inv.IsParameterSelected("Offset") {
			ret["location_network[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionLocationNetworkIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
