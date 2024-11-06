package client

import ()

// ActionComponentIndex is a type for action Component#Index
type ActionComponentIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionComponentIndex(client *Client) *ActionComponentIndex {
	return &ActionComponentIndex{
		Client: client,
	}
}

// ActionComponentIndexMetaGlobalInput is a type for action global meta input parameters
type ActionComponentIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionComponentIndexMetaGlobalInput) SetCount(value bool) *ActionComponentIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionComponentIndexMetaGlobalInput) SetIncludes(value string) *ActionComponentIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionComponentIndexMetaGlobalInput) SetNo(value bool) *ActionComponentIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionComponentIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionComponentIndexMetaGlobalInput) SelectParameters(params ...string) *ActionComponentIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionComponentIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionComponentIndexInput is a type for action input parameters
type ActionComponentIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionComponentIndexInput) SetFromId(value int64) *ActionComponentIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionComponentIndexInput) SetLimit(value int64) *ActionComponentIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionComponentIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionComponentIndexInput) SelectParameters(params ...string) *ActionComponentIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionComponentIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionComponentIndexInput) UnselectParameters(params ...string) *ActionComponentIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionComponentIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionComponentIndexOutput is a type for action output parameters
type ActionComponentIndexOutput struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Label       string `json:"label"`
	Name        string `json:"name"`
}

// Type for action response, including envelope
type ActionComponentIndexResponse struct {
	Action *ActionComponentIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Components []*ActionComponentIndexOutput `json:"components"`
	}

	// Action output without the namespace
	Output []*ActionComponentIndexOutput
}

// Prepare the action for invocation
func (action *ActionComponentIndex) Prepare() *ActionComponentIndexInvocation {
	return &ActionComponentIndexInvocation{
		Action: action,
		Path:   "/v7.0/components",
	}
}

// ActionComponentIndexInvocation is used to configure action for invocation
type ActionComponentIndexInvocation struct {
	// Pointer to the action
	Action *ActionComponentIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionComponentIndexInput
	// Global meta input parameters
	MetaInput *ActionComponentIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionComponentIndexInvocation) NewInput() *ActionComponentIndexInput {
	inv.Input = &ActionComponentIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionComponentIndexInvocation) SetInput(input *ActionComponentIndexInput) *ActionComponentIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionComponentIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionComponentIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionComponentIndexInvocation) NewMetaInput() *ActionComponentIndexMetaGlobalInput {
	inv.MetaInput = &ActionComponentIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionComponentIndexInvocation) SetMetaInput(input *ActionComponentIndexMetaGlobalInput) *ActionComponentIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionComponentIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionComponentIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionComponentIndexInvocation) Call() (*ActionComponentIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionComponentIndexInvocation) callAsQuery() (*ActionComponentIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionComponentIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Components
	}
	return resp, err
}

func (inv *ActionComponentIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["component[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["component[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionComponentIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
