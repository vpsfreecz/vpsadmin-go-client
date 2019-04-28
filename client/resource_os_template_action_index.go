package client

import (
)

// ActionOsTemplateIndex is a type for action Os_template#Index
type ActionOsTemplateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOsTemplateIndex(client *Client) *ActionOsTemplateIndex {
	return &ActionOsTemplateIndex{
		Client: client,
	}
}

// ActionOsTemplateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOsTemplateIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsTemplateIndexMetaGlobalInput) SetNo(value bool) *ActionOsTemplateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOsTemplateIndexMetaGlobalInput) SetCount(value bool) *ActionOsTemplateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsTemplateIndexMetaGlobalInput) SetIncludes(value string) *ActionOsTemplateIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsTemplateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOsTemplateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsTemplateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsTemplateIndexInput is a type for action input parameters
type ActionOsTemplateIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Location int64 `json:"location"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionOsTemplateIndexInput) SetOffset(value int64) *ActionOsTemplateIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOsTemplateIndexInput) SetLimit(value int64) *ActionOsTemplateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionOsTemplateIndexInput) SetLocation(value int64) *ActionOsTemplateIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsTemplateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateIndexInput) SelectParameters(params ...string) *ActionOsTemplateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsTemplateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionOsTemplateIndexOutput is a type for action output parameters
type ActionOsTemplateIndexOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	Info string `json:"info"`
	Enabled bool `json:"enabled"`
	Supported bool `json:"supported"`
	Order int64 `json:"order"`
	HypervisorType string `json:"hypervisor_type"`
}


// Type for action response, including envelope
type ActionOsTemplateIndexResponse struct {
	Action *ActionOsTemplateIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OsTemplates []*ActionOsTemplateIndexOutput `json:"os_templates"`
	}

	// Action output without the namespace
	Output []*ActionOsTemplateIndexOutput
}


// Prepare the action for invocation
func (action *ActionOsTemplateIndex) Prepare() *ActionOsTemplateIndexInvocation {
	return &ActionOsTemplateIndexInvocation{
		Action: action,
		Path: "/v5.0/os_templates",
	}
}

// ActionOsTemplateIndexInvocation is used to configure action for invocation
type ActionOsTemplateIndexInvocation struct {
	// Pointer to the action
	Action *ActionOsTemplateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOsTemplateIndexInput
	// Global meta input parameters
	MetaInput *ActionOsTemplateIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOsTemplateIndexInvocation) NewInput() *ActionOsTemplateIndexInput {
	inv.Input = &ActionOsTemplateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOsTemplateIndexInvocation) SetInput(input *ActionOsTemplateIndexInput) *ActionOsTemplateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOsTemplateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsTemplateIndexInvocation) NewMetaInput() *ActionOsTemplateIndexMetaGlobalInput {
	inv.MetaInput = &ActionOsTemplateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsTemplateIndexInvocation) SetMetaInput(input *ActionOsTemplateIndexMetaGlobalInput) *ActionOsTemplateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsTemplateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsTemplateIndexInvocation) Call() (*ActionOsTemplateIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOsTemplateIndexInvocation) callAsQuery() (*ActionOsTemplateIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOsTemplateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OsTemplates
	}
	return resp, err
}



func (inv *ActionOsTemplateIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["os_template[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["os_template[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["os_template[location]"] = convertInt64ToString(inv.Input.Location)
		}
	}
}

func (inv *ActionOsTemplateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

