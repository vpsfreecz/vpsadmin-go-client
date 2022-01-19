package client

import (
)

// ActionEnvironmentIndex is a type for action Environment#Index
type ActionEnvironmentIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEnvironmentIndex(client *Client) *ActionEnvironmentIndex {
	return &ActionEnvironmentIndex{
		Client: client,
	}
}

// ActionEnvironmentIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEnvironmentIndexMetaGlobalInput struct {
	Count bool `json:"count"`
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEnvironmentIndexMetaGlobalInput) SetCount(value bool) *ActionEnvironmentIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEnvironmentIndexMetaGlobalInput) SetIncludes(value string) *ActionEnvironmentIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionEnvironmentIndexMetaGlobalInput) SetNo(value bool) *ActionEnvironmentIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEnvironmentIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentIndexInput is a type for action input parameters
type ActionEnvironmentIndexInput struct {
	HasHypervisor bool `json:"has_hypervisor"`
	HasStorage bool `json:"has_storage"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetHasHypervisor sets parameter HasHypervisor to value and selects it for sending
func (in *ActionEnvironmentIndexInput) SetHasHypervisor(value bool) *ActionEnvironmentIndexInput {
	in.HasHypervisor = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HasHypervisor"] = nil
	return in
}
// SetHasStorage sets parameter HasStorage to value and selects it for sending
func (in *ActionEnvironmentIndexInput) SetHasStorage(value bool) *ActionEnvironmentIndexInput {
	in.HasStorage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HasStorage"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEnvironmentIndexInput) SetLimit(value int64) *ActionEnvironmentIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionEnvironmentIndexInput) SetOffset(value int64) *ActionEnvironmentIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentIndexInput) SelectParameters(params ...string) *ActionEnvironmentIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionEnvironmentIndexOutput is a type for action output parameters
type ActionEnvironmentIndexOutput struct {
	CanCreateVps bool `json:"can_create_vps"`
	CanDestroyVps bool `json:"can_destroy_vps"`
	Description string `json:"description"`
	Domain string `json:"domain"`
	Id int64 `json:"id"`
	Label string `json:"label"`
	MaintenanceLock string `json:"maintenance_lock"`
	MaintenanceLockReason string `json:"maintenance_lock_reason"`
	MaxVpsCount int64 `json:"max_vps_count"`
	UserIpOwnership bool `json:"user_ip_ownership"`
	VpsLifetime int64 `json:"vps_lifetime"`
}


// Type for action response, including envelope
type ActionEnvironmentIndexResponse struct {
	Action *ActionEnvironmentIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Environments []*ActionEnvironmentIndexOutput `json:"environments"`
	}

	// Action output without the namespace
	Output []*ActionEnvironmentIndexOutput
}


// Prepare the action for invocation
func (action *ActionEnvironmentIndex) Prepare() *ActionEnvironmentIndexInvocation {
	return &ActionEnvironmentIndexInvocation{
		Action: action,
		Path: "/v6.0/environments",
	}
}

// ActionEnvironmentIndexInvocation is used to configure action for invocation
type ActionEnvironmentIndexInvocation struct {
	// Pointer to the action
	Action *ActionEnvironmentIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEnvironmentIndexInput
	// Global meta input parameters
	MetaInput *ActionEnvironmentIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEnvironmentIndexInvocation) NewInput() *ActionEnvironmentIndexInput {
	inv.Input = &ActionEnvironmentIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEnvironmentIndexInvocation) SetInput(input *ActionEnvironmentIndexInput) *ActionEnvironmentIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEnvironmentIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEnvironmentIndexInvocation) NewMetaInput() *ActionEnvironmentIndexMetaGlobalInput {
	inv.MetaInput = &ActionEnvironmentIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEnvironmentIndexInvocation) SetMetaInput(input *ActionEnvironmentIndexMetaGlobalInput) *ActionEnvironmentIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEnvironmentIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEnvironmentIndexInvocation) Call() (*ActionEnvironmentIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionEnvironmentIndexInvocation) callAsQuery() (*ActionEnvironmentIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEnvironmentIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Environments
	}
	return resp, err
}



func (inv *ActionEnvironmentIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("HasHypervisor") {
			ret["environment[has_hypervisor]"] = convertBoolToString(inv.Input.HasHypervisor)
		}
		if inv.IsParameterSelected("HasStorage") {
			ret["environment[has_storage]"] = convertBoolToString(inv.Input.HasStorage)
		}
		if inv.IsParameterSelected("Limit") {
			ret["environment[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["environment[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionEnvironmentIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

