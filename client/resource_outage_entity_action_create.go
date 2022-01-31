package client

import (
	"strings"
)

// ActionOutageEntityCreate is a type for action Outage.Entity#Create
type ActionOutageEntityCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageEntityCreate(client *Client) *ActionOutageEntityCreate {
	return &ActionOutageEntityCreate{
		Client: client,
	}
}

// ActionOutageEntityCreateMetaGlobalInput is a type for action global meta input parameters
type ActionOutageEntityCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageEntityCreateMetaGlobalInput) SetIncludes(value string) *ActionOutageEntityCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageEntityCreateMetaGlobalInput) SetNo(value bool) *ActionOutageEntityCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageEntityCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageEntityCreateMetaGlobalInput) SelectParameters(params ...string) *ActionOutageEntityCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageEntityCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageEntityCreateInput is a type for action input parameters
type ActionOutageEntityCreateInput struct {
	EntityId int64  `json:"entity_id"`
	Name     string `json:"name"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEntityId sets parameter EntityId to value and selects it for sending
func (in *ActionOutageEntityCreateInput) SetEntityId(value int64) *ActionOutageEntityCreateInput {
	in.EntityId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EntityId"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionOutageEntityCreateInput) SetName(value string) *ActionOutageEntityCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageEntityCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageEntityCreateInput) SelectParameters(params ...string) *ActionOutageEntityCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOutageEntityCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOutageEntityCreateInput) UnselectParameters(params ...string) *ActionOutageEntityCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOutageEntityCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageEntityCreateRequest is a type for the entire action request
type ActionOutageEntityCreateRequest struct {
	Entity map[string]interface{} `json:"entity"`
	Meta   map[string]interface{} `json:"_meta"`
}

// ActionOutageEntityCreateOutput is a type for action output parameters
type ActionOutageEntityCreateOutput struct {
	EntityId int64  `json:"entity_id"`
	Id       int64  `json:"id"`
	Label    string `json:"label"`
	Name     string `json:"name"`
}

// Type for action response, including envelope
type ActionOutageEntityCreateResponse struct {
	Action *ActionOutageEntityCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Entity *ActionOutageEntityCreateOutput `json:"entity"`
	}

	// Action output without the namespace
	Output *ActionOutageEntityCreateOutput
}

// Prepare the action for invocation
func (action *ActionOutageEntityCreate) Prepare() *ActionOutageEntityCreateInvocation {
	return &ActionOutageEntityCreateInvocation{
		Action: action,
		Path:   "/v6.0/outages/{outage_id}/entities",
	}
}

// ActionOutageEntityCreateInvocation is used to configure action for invocation
type ActionOutageEntityCreateInvocation struct {
	// Pointer to the action
	Action *ActionOutageEntityCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageEntityCreateInput
	// Global meta input parameters
	MetaInput *ActionOutageEntityCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageEntityCreateInvocation) SetPathParamInt(param string, value int64) *ActionOutageEntityCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageEntityCreateInvocation) SetPathParamString(param string, value string) *ActionOutageEntityCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageEntityCreateInvocation) NewInput() *ActionOutageEntityCreateInput {
	inv.Input = &ActionOutageEntityCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageEntityCreateInvocation) SetInput(input *ActionOutageEntityCreateInput) *ActionOutageEntityCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageEntityCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOutageEntityCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageEntityCreateInvocation) NewMetaInput() *ActionOutageEntityCreateMetaGlobalInput {
	inv.MetaInput = &ActionOutageEntityCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageEntityCreateInvocation) SetMetaInput(input *ActionOutageEntityCreateMetaGlobalInput) *ActionOutageEntityCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageEntityCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageEntityCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageEntityCreateInvocation) Call() (*ActionOutageEntityCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOutageEntityCreateInvocation) callAsBody() (*ActionOutageEntityCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageEntityCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Entity
	}
	return resp, err
}

func (inv *ActionOutageEntityCreateInvocation) makeAllInputParams() *ActionOutageEntityCreateRequest {
	return &ActionOutageEntityCreateRequest{
		Entity: inv.makeInputParams(),
		Meta:   inv.makeMetaInputParams(),
	}
}

func (inv *ActionOutageEntityCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("EntityId") {
			ret["entity_id"] = inv.Input.EntityId
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
	}

	return ret
}

func (inv *ActionOutageEntityCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
