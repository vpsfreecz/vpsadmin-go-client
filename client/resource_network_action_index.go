package client

import (
)

// ActionNetworkIndex is a type for action Network#Index
type ActionNetworkIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkIndex(client *Client) *ActionNetworkIndex {
	return &ActionNetworkIndex{
		Client: client,
	}
}

// ActionNetworkIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkIndexMetaGlobalInput) SetNo(value bool) *ActionNetworkIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNetworkIndexMetaGlobalInput) SetCount(value bool) *ActionNetworkIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkIndexMetaGlobalInput) SetIncludes(value string) *ActionNetworkIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkIndexInput is a type for action input parameters
type ActionNetworkIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Location int64 `json:"location"`
	Purpose string `json:"purpose"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionNetworkIndexInput) SetOffset(value int64) *ActionNetworkIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNetworkIndexInput) SetLimit(value int64) *ActionNetworkIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNetworkIndexInput) SetLocation(value int64) *ActionNetworkIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetPurpose sets parameter Purpose to value and selects it for sending
func (in *ActionNetworkIndexInput) SetPurpose(value string) *ActionNetworkIndexInput {
	in.Purpose = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Purpose"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkIndexInput) SelectParameters(params ...string) *ActionNetworkIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionNetworkIndexOutput is a type for action output parameters
type ActionNetworkIndexOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Location *ActionLocationShowOutput `json:"location"`
	IpVersion int64 `json:"ip_version"`
	Address string `json:"address"`
	Prefix int64 `json:"prefix"`
	Role string `json:"role"`
	Managed bool `json:"managed"`
	SplitAccess string `json:"split_access"`
	SplitPrefix int64 `json:"split_prefix"`
	Autopick bool `json:"autopick"`
	Purpose string `json:"purpose"`
	Size int64 `json:"size"`
	Used int64 `json:"used"`
	Assigned int64 `json:"assigned"`
	Owned int64 `json:"owned"`
}


// Type for action response, including envelope
type ActionNetworkIndexResponse struct {
	Action *ActionNetworkIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Networks []*ActionNetworkIndexOutput `json:"networks"`
	}

	// Action output without the namespace
	Output []*ActionNetworkIndexOutput
}


// Prepare the action for invocation
func (action *ActionNetworkIndex) Prepare() *ActionNetworkIndexInvocation {
	return &ActionNetworkIndexInvocation{
		Action: action,
		Path: "/v5.0/networks",
	}
}

// ActionNetworkIndexInvocation is used to configure action for invocation
type ActionNetworkIndexInvocation struct {
	// Pointer to the action
	Action *ActionNetworkIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkIndexInput
	// Global meta input parameters
	MetaInput *ActionNetworkIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkIndexInvocation) NewInput() *ActionNetworkIndexInput {
	inv.Input = &ActionNetworkIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkIndexInvocation) SetInput(input *ActionNetworkIndexInput) *ActionNetworkIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkIndexInvocation) NewMetaInput() *ActionNetworkIndexMetaGlobalInput {
	inv.MetaInput = &ActionNetworkIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkIndexInvocation) SetMetaInput(input *ActionNetworkIndexMetaGlobalInput) *ActionNetworkIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkIndexInvocation) Call() (*ActionNetworkIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNetworkIndexInvocation) callAsQuery() (*ActionNetworkIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNetworkIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Networks
	}
	return resp, err
}



func (inv *ActionNetworkIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["network[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["network[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["network[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Purpose") {
			ret["network[purpose]"] = inv.Input.Purpose
		}
	}
}

func (inv *ActionNetworkIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

