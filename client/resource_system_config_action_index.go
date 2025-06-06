package client

import ()

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
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSystemConfigIndexMetaGlobalInput) SetNo(value bool) *ActionSystemConfigIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	Category string `json:"category"`
	FromId   int64  `json:"from_id"`
	Limit    int64  `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionSystemConfigIndexInput) SetFromId(value int64) *ActionSystemConfigIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
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

// UnselectParameters unsets parameters from ActionSystemConfigIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSystemConfigIndexInput) UnselectParameters(params ...string) *ActionSystemConfigIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Category     string `json:"category"`
	Description  string `json:"description"`
	Label        string `json:"label"`
	MinUserLevel int64  `json:"min_user_level"`
	Name         string `json:"name"`
	Type         string `json:"type"`
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
		Path:   "/v7.0/system_configs",
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSystemConfigIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSystemConfigIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		if inv.IsParameterSelected("Category") {
			ret["system_config[category]"] = inv.Input.Category
		}
		if inv.IsParameterSelected("FromId") {
			ret["system_config[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["system_config[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionSystemConfigIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
