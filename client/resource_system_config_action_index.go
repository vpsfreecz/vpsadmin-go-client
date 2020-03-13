package client

import (
)

// ActionSystemConfigIndex is a type for action System_config#Index
type ActionSystemConfigIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionSystemConfigIndex(client *Client) *ActionSystemConfigIndex {
	return &ActionSystemConfigIndex{
		Client: client,
	}
}

// ActionSystemConfigIndexMetaGlobalInput is a type for action global meta input parameters
type ActionSystemConfigIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSystemConfigIndexMetaGlobalInput) SetNo(value bool) *ActionSystemConfigIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionSystemConfigIndexMetaGlobalInput) SetCount(value bool) *ActionSystemConfigIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSystemConfigIndexMetaGlobalInput) SetIncludes(value string) *ActionSystemConfigIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionSystemConfigIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSystemConfigIndexMetaGlobalInput) SelectParameters(params ...string) *ActionSystemConfigIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSystemConfigIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSystemConfigIndexInput is a type for action input parameters
type ActionSystemConfigIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Category string `json:"category"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionSystemConfigIndexInput) SetOffset(value int64) *ActionSystemConfigIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionSystemConfigIndexInput) SetLimit(value int64) *ActionSystemConfigIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetCategory sets parameter Category to value and selects it for sending
func (in *ActionSystemConfigIndexInput) SetCategory(value string) *ActionSystemConfigIndexInput {
	in.Category = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Category"] = nil
	return in
}

// SelectParameters sets parameters from ActionSystemConfigIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSystemConfigIndexInput) SelectParameters(params ...string) *ActionSystemConfigIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSystemConfigIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionSystemConfigIndexOutput is a type for action output parameters
type ActionSystemConfigIndexOutput struct {
	Category string `json:"category"`
	Name string `json:"name"`
	Type string `json:"type"`
	Label string `json:"label"`
	Description string `json:"description"`
	MinUserLevel int64 `json:"min_user_level"`
}


// Type for action response, including envelope
type ActionSystemConfigIndexResponse struct {
	Action *ActionSystemConfigIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SystemConfigs []*ActionSystemConfigIndexOutput `json:"system_configs"`
	}

	// Action output without the namespace
	Output []*ActionSystemConfigIndexOutput
}


// Prepare the action for invocation
func (action *ActionSystemConfigIndex) Prepare() *ActionSystemConfigIndexInvocation {
	return &ActionSystemConfigIndexInvocation{
		Action: action,
		Path: "/v6.0/system_configs",
	}
}

// ActionSystemConfigIndexInvocation is used to configure action for invocation
type ActionSystemConfigIndexInvocation struct {
	// Pointer to the action
	Action *ActionSystemConfigIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSystemConfigIndexInput
	// Global meta input parameters
	MetaInput *ActionSystemConfigIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSystemConfigIndexInvocation) NewInput() *ActionSystemConfigIndexInput {
	inv.Input = &ActionSystemConfigIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSystemConfigIndexInvocation) SetInput(input *ActionSystemConfigIndexInput) *ActionSystemConfigIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSystemConfigIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSystemConfigIndexInvocation) NewMetaInput() *ActionSystemConfigIndexMetaGlobalInput {
	inv.MetaInput = &ActionSystemConfigIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSystemConfigIndexInvocation) SetMetaInput(input *ActionSystemConfigIndexMetaGlobalInput) *ActionSystemConfigIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSystemConfigIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSystemConfigIndexInvocation) Call() (*ActionSystemConfigIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionSystemConfigIndexInvocation) callAsQuery() (*ActionSystemConfigIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSystemConfigIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SystemConfigs
	}
	return resp, err
}



func (inv *ActionSystemConfigIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["system_config[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["system_config[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Category") {
			ret["system_config[category]"] = inv.Input.Category
		}
	}
}

func (inv *ActionSystemConfigIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

