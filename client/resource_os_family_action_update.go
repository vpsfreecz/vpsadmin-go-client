package client

import (
	"strings"
)

// ActionOsFamilyUpdate is a type for action Os_family#Update
type ActionOsFamilyUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionOsFamilyUpdate(client *Client) *ActionOsFamilyUpdate {
	return &ActionOsFamilyUpdate{
		Client: client,
	}
}

// ActionOsFamilyUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionOsFamilyUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsFamilyUpdateMetaGlobalInput) SetIncludes(value string) *ActionOsFamilyUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsFamilyUpdateMetaGlobalInput) SetNo(value bool) *ActionOsFamilyUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsFamilyUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsFamilyUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionOsFamilyUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsFamilyUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsFamilyUpdateInput is a type for action input parameters
type ActionOsFamilyUpdateInput struct {
	Description string `json:"description"`
	Label       string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDescription sets parameter Description to value and selects it for sending
func (in *ActionOsFamilyUpdateInput) SetDescription(value string) *ActionOsFamilyUpdateInput {
	in.Description = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Description"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionOsFamilyUpdateInput) SetLabel(value string) *ActionOsFamilyUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsFamilyUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsFamilyUpdateInput) SelectParameters(params ...string) *ActionOsFamilyUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOsFamilyUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOsFamilyUpdateInput) UnselectParameters(params ...string) *ActionOsFamilyUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOsFamilyUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsFamilyUpdateRequest is a type for the entire action request
type ActionOsFamilyUpdateRequest struct {
	OsFamily map[string]interface{} `json:"os_family"`
	Meta     map[string]interface{} `json:"_meta"`
}

// ActionOsFamilyUpdateOutput is a type for action output parameters
type ActionOsFamilyUpdateOutput struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Label       string `json:"label"`
}

// Type for action response, including envelope
type ActionOsFamilyUpdateResponse struct {
	Action *ActionOsFamilyUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OsFamily *ActionOsFamilyUpdateOutput `json:"os_family"`
	}

	// Action output without the namespace
	Output *ActionOsFamilyUpdateOutput
}

// Prepare the action for invocation
func (action *ActionOsFamilyUpdate) Prepare() *ActionOsFamilyUpdateInvocation {
	return &ActionOsFamilyUpdateInvocation{
		Action: action,
		Path:   "/v7.0/os_families/{os_family_id}",
	}
}

// ActionOsFamilyUpdateInvocation is used to configure action for invocation
type ActionOsFamilyUpdateInvocation struct {
	// Pointer to the action
	Action *ActionOsFamilyUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOsFamilyUpdateInput
	// Global meta input parameters
	MetaInput *ActionOsFamilyUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOsFamilyUpdateInvocation) SetPathParamInt(param string, value int64) *ActionOsFamilyUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOsFamilyUpdateInvocation) SetPathParamString(param string, value string) *ActionOsFamilyUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOsFamilyUpdateInvocation) NewInput() *ActionOsFamilyUpdateInput {
	inv.Input = &ActionOsFamilyUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOsFamilyUpdateInvocation) SetInput(input *ActionOsFamilyUpdateInput) *ActionOsFamilyUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOsFamilyUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOsFamilyUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsFamilyUpdateInvocation) NewMetaInput() *ActionOsFamilyUpdateMetaGlobalInput {
	inv.MetaInput = &ActionOsFamilyUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsFamilyUpdateInvocation) SetMetaInput(input *ActionOsFamilyUpdateMetaGlobalInput) *ActionOsFamilyUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsFamilyUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsFamilyUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsFamilyUpdateInvocation) Call() (*ActionOsFamilyUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOsFamilyUpdateInvocation) callAsBody() (*ActionOsFamilyUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOsFamilyUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OsFamily
	}
	return resp, err
}

func (inv *ActionOsFamilyUpdateInvocation) makeAllInputParams() *ActionOsFamilyUpdateRequest {
	return &ActionOsFamilyUpdateRequest{
		OsFamily: inv.makeInputParams(),
		Meta:     inv.makeMetaInputParams(),
	}
}

func (inv *ActionOsFamilyUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Description") {
			ret["description"] = inv.Input.Description
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionOsFamilyUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
