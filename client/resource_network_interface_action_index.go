package client

import (
)

// ActionNetworkInterfaceIndex is a type for action Network_interface#Index
type ActionNetworkInterfaceIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkInterfaceIndex(client *Client) *ActionNetworkInterfaceIndex {
	return &ActionNetworkInterfaceIndex{
		Client: client,
	}
}

// ActionNetworkInterfaceIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkInterfaceIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkInterfaceIndexMetaGlobalInput) SetNo(value bool) *ActionNetworkInterfaceIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNetworkInterfaceIndexMetaGlobalInput) SetCount(value bool) *ActionNetworkInterfaceIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkInterfaceIndexMetaGlobalInput) SetIncludes(value string) *ActionNetworkInterfaceIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkInterfaceIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkInterfaceIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceIndexInput is a type for action input parameters
type ActionNetworkInterfaceIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Vps int64 `json:"vps"`
	Location int64 `json:"location"`
	User int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionNetworkInterfaceIndexInput) SetOffset(value int64) *ActionNetworkInterfaceIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNetworkInterfaceIndexInput) SetLimit(value int64) *ActionNetworkInterfaceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionNetworkInterfaceIndexInput) SetVps(value int64) *ActionNetworkInterfaceIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNetworkInterfaceIndexInput) SetLocation(value int64) *ActionNetworkInterfaceIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionNetworkInterfaceIndexInput) SetUser(value int64) *ActionNetworkInterfaceIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceIndexInput) SelectParameters(params ...string) *ActionNetworkInterfaceIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkInterfaceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionNetworkInterfaceIndexOutput is a type for action output parameters
type ActionNetworkInterfaceIndexOutput struct {
	Id int64 `json:"id"`
	Vps *ActionVpsShowOutput `json:"vps"`
	Name string `json:"name"`
	Type string `json:"type"`
	Mac string `json:"mac"`
}


// Type for action response, including envelope
type ActionNetworkInterfaceIndexResponse struct {
	Action *ActionNetworkInterfaceIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NetworkInterfaces []*ActionNetworkInterfaceIndexOutput `json:"network_interfaces"`
	}

	// Action output without the namespace
	Output []*ActionNetworkInterfaceIndexOutput
}


// Prepare the action for invocation
func (action *ActionNetworkInterfaceIndex) Prepare() *ActionNetworkInterfaceIndexInvocation {
	return &ActionNetworkInterfaceIndexInvocation{
		Action: action,
		Path: "/v5.0/network_interfaces",
	}
}

// ActionNetworkInterfaceIndexInvocation is used to configure action for invocation
type ActionNetworkInterfaceIndexInvocation struct {
	// Pointer to the action
	Action *ActionNetworkInterfaceIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkInterfaceIndexInput
	// Global meta input parameters
	MetaInput *ActionNetworkInterfaceIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkInterfaceIndexInvocation) NewInput() *ActionNetworkInterfaceIndexInput {
	inv.Input = &ActionNetworkInterfaceIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkInterfaceIndexInvocation) SetInput(input *ActionNetworkInterfaceIndexInput) *ActionNetworkInterfaceIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkInterfaceIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkInterfaceIndexInvocation) NewMetaInput() *ActionNetworkInterfaceIndexMetaGlobalInput {
	inv.MetaInput = &ActionNetworkInterfaceIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkInterfaceIndexInvocation) SetMetaInput(input *ActionNetworkInterfaceIndexMetaGlobalInput) *ActionNetworkInterfaceIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkInterfaceIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkInterfaceIndexInvocation) Call() (*ActionNetworkInterfaceIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNetworkInterfaceIndexInvocation) callAsQuery() (*ActionNetworkInterfaceIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNetworkInterfaceIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NetworkInterfaces
	}
	return resp, err
}



func (inv *ActionNetworkInterfaceIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["network_interface[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["network_interface[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Vps") {
			ret["network_interface[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
		if inv.IsParameterSelected("Location") {
			ret["network_interface[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("User") {
			ret["network_interface[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionNetworkInterfaceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

