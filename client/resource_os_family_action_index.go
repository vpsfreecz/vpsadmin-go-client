package client

import ()

// ActionOsFamilyIndex is a type for action Os_family#Index
type ActionOsFamilyIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOsFamilyIndex(client *Client) *ActionOsFamilyIndex {
	return &ActionOsFamilyIndex{
		Client: client,
	}
}

// ActionOsFamilyIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOsFamilyIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOsFamilyIndexMetaGlobalInput) SetCount(value bool) *ActionOsFamilyIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsFamilyIndexMetaGlobalInput) SetIncludes(value string) *ActionOsFamilyIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsFamilyIndexMetaGlobalInput) SetNo(value bool) *ActionOsFamilyIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsFamilyIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsFamilyIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOsFamilyIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsFamilyIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsFamilyIndexInput is a type for action input parameters
type ActionOsFamilyIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionOsFamilyIndexInput) SetFromId(value int64) *ActionOsFamilyIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOsFamilyIndexInput) SetLimit(value int64) *ActionOsFamilyIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsFamilyIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsFamilyIndexInput) SelectParameters(params ...string) *ActionOsFamilyIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOsFamilyIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOsFamilyIndexInput) UnselectParameters(params ...string) *ActionOsFamilyIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOsFamilyIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsFamilyIndexOutput is a type for action output parameters
type ActionOsFamilyIndexOutput struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Label       string `json:"label"`
}

// Type for action response, including envelope
type ActionOsFamilyIndexResponse struct {
	Action *ActionOsFamilyIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OsFamilies []*ActionOsFamilyIndexOutput `json:"os_families"`
	}

	// Action output without the namespace
	Output []*ActionOsFamilyIndexOutput
}

// Prepare the action for invocation
func (action *ActionOsFamilyIndex) Prepare() *ActionOsFamilyIndexInvocation {
	return &ActionOsFamilyIndexInvocation{
		Action: action,
		Path:   "/v7.0/os_families",
	}
}

// ActionOsFamilyIndexInvocation is used to configure action for invocation
type ActionOsFamilyIndexInvocation struct {
	// Pointer to the action
	Action *ActionOsFamilyIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOsFamilyIndexInput
	// Global meta input parameters
	MetaInput *ActionOsFamilyIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOsFamilyIndexInvocation) NewInput() *ActionOsFamilyIndexInput {
	inv.Input = &ActionOsFamilyIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOsFamilyIndexInvocation) SetInput(input *ActionOsFamilyIndexInput) *ActionOsFamilyIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOsFamilyIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOsFamilyIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsFamilyIndexInvocation) NewMetaInput() *ActionOsFamilyIndexMetaGlobalInput {
	inv.MetaInput = &ActionOsFamilyIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsFamilyIndexInvocation) SetMetaInput(input *ActionOsFamilyIndexMetaGlobalInput) *ActionOsFamilyIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsFamilyIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsFamilyIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsFamilyIndexInvocation) Call() (*ActionOsFamilyIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOsFamilyIndexInvocation) callAsQuery() (*ActionOsFamilyIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOsFamilyIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OsFamilies
	}
	return resp, err
}

func (inv *ActionOsFamilyIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["os_family[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["os_family[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionOsFamilyIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
