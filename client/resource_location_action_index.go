package client

import (
)

// ActionLocationIndex is a type for action Location#Index
type ActionLocationIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationIndex(client *Client) *ActionLocationIndex {
	return &ActionLocationIndex{
		Client: client,
	}
}

// ActionLocationIndexMetaGlobalInput is a type for action global meta input parameters
type ActionLocationIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationIndexMetaGlobalInput) SetNo(value bool) *ActionLocationIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionLocationIndexMetaGlobalInput) SetCount(value bool) *ActionLocationIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationIndexMetaGlobalInput) SetIncludes(value string) *ActionLocationIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationIndexMetaGlobalInput) SelectParameters(params ...string) *ActionLocationIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationIndexInput is a type for action input parameters
type ActionLocationIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Environment int64 `json:"environment"`
	HasHypervisor bool `json:"has_hypervisor"`
	HasStorage bool `json:"has_storage"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionLocationIndexInput) SetOffset(value int64) *ActionLocationIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionLocationIndexInput) SetLimit(value int64) *ActionLocationIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionLocationIndexInput) SetEnvironment(value int64) *ActionLocationIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetHasHypervisor sets parameter HasHypervisor to value and selects it for sending
func (in *ActionLocationIndexInput) SetHasHypervisor(value bool) *ActionLocationIndexInput {
	in.HasHypervisor = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HasHypervisor"] = nil
	return in
}
// SetHasStorage sets parameter HasStorage to value and selects it for sending
func (in *ActionLocationIndexInput) SetHasStorage(value bool) *ActionLocationIndexInput {
	in.HasStorage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HasStorage"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationIndexInput) SelectParameters(params ...string) *ActionLocationIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionLocationIndexOutput is a type for action output parameters
type ActionLocationIndexOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	HasIpv6 bool `json:"has_ipv6"`
	VpsOnboot bool `json:"vps_onboot"`
	RemoteConsoleServer string `json:"remote_console_server"`
	Domain string `json:"domain"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	MaintenanceLock string `json:"maintenance_lock"`
	MaintenanceLockReason string `json:"maintenance_lock_reason"`
}


// Type for action response, including envelope
type ActionLocationIndexResponse struct {
	Action *ActionLocationIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Locations []*ActionLocationIndexOutput `json:"locations"`
	}

	// Action output without the namespace
	Output []*ActionLocationIndexOutput
}


// Prepare the action for invocation
func (action *ActionLocationIndex) Prepare() *ActionLocationIndexInvocation {
	return &ActionLocationIndexInvocation{
		Action: action,
		Path: "/v5.0/locations",
	}
}

// ActionLocationIndexInvocation is used to configure action for invocation
type ActionLocationIndexInvocation struct {
	// Pointer to the action
	Action *ActionLocationIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionLocationIndexInput
	// Global meta input parameters
	MetaInput *ActionLocationIndexMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionLocationIndexInvocation) SetInput(input *ActionLocationIndexInput) *ActionLocationIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionLocationIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationIndexInvocation) SetMetaInput(input *ActionLocationIndexMetaGlobalInput) *ActionLocationIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationIndexInvocation) Call() (*ActionLocationIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionLocationIndexInvocation) callAsQuery() (*ActionLocationIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionLocationIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Locations
	}
	return resp, err
}



func (inv *ActionLocationIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["location[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["location[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Environment") {
			ret["location[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("HasHypervisor") {
			ret["location[has_hypervisor]"] = convertBoolToString(inv.Input.HasHypervisor)
		}
		if inv.IsParameterSelected("HasStorage") {
			ret["location[has_storage]"] = convertBoolToString(inv.Input.HasStorage)
		}
	}
}

func (inv *ActionLocationIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

