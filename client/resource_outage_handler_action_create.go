package client

import (
	"strings"
)

// ActionOutageHandlerCreate is a type for action Outage.Handler#Create
type ActionOutageHandlerCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageHandlerCreate(client *Client) *ActionOutageHandlerCreate {
	return &ActionOutageHandlerCreate{
		Client: client,
	}
}

// ActionOutageHandlerCreateMetaGlobalInput is a type for action global meta input parameters
type ActionOutageHandlerCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageHandlerCreateMetaGlobalInput) SetNo(value bool) *ActionOutageHandlerCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageHandlerCreateMetaGlobalInput) SetIncludes(value string) *ActionOutageHandlerCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageHandlerCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageHandlerCreateMetaGlobalInput) SelectParameters(params ...string) *ActionOutageHandlerCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageHandlerCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageHandlerCreateInput is a type for action input parameters
type ActionOutageHandlerCreateInput struct {
	User int64 `json:"user"`
	FullName string `json:"full_name"`
	Note string `json:"note"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionOutageHandlerCreateInput) SetUser(value int64) *ActionOutageHandlerCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionOutageHandlerCreateInput) SetFullName(value string) *ActionOutageHandlerCreateInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
	return in
}
// SetNote sets parameter Note to value and selects it for sending
func (in *ActionOutageHandlerCreateInput) SetNote(value string) *ActionOutageHandlerCreateInput {
	in.Note = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Note"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageHandlerCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageHandlerCreateInput) SelectParameters(params ...string) *ActionOutageHandlerCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageHandlerCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageHandlerCreateRequest is a type for the entire action request
type ActionOutageHandlerCreateRequest struct {
	Handler map[string]interface{} `json:"handler"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionOutageHandlerCreateOutput is a type for action output parameters
type ActionOutageHandlerCreateOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	FullName string `json:"full_name"`
	Note string `json:"note"`
}


// Type for action response, including envelope
type ActionOutageHandlerCreateResponse struct {
	Action *ActionOutageHandlerCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Handler *ActionOutageHandlerCreateOutput `json:"handler"`
	}

	// Action output without the namespace
	Output *ActionOutageHandlerCreateOutput
}


// Prepare the action for invocation
func (action *ActionOutageHandlerCreate) Prepare() *ActionOutageHandlerCreateInvocation {
	return &ActionOutageHandlerCreateInvocation{
		Action: action,
		Path: "/v5.0/outages/{outage_id}/handlers",
	}
}

// ActionOutageHandlerCreateInvocation is used to configure action for invocation
type ActionOutageHandlerCreateInvocation struct {
	// Pointer to the action
	Action *ActionOutageHandlerCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageHandlerCreateInput
	// Global meta input parameters
	MetaInput *ActionOutageHandlerCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageHandlerCreateInvocation) SetPathParamInt(param string, value int64) *ActionOutageHandlerCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageHandlerCreateInvocation) SetPathParamString(param string, value string) *ActionOutageHandlerCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageHandlerCreateInvocation) NewInput() *ActionOutageHandlerCreateInput {
	inv.Input = &ActionOutageHandlerCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageHandlerCreateInvocation) SetInput(input *ActionOutageHandlerCreateInput) *ActionOutageHandlerCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageHandlerCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageHandlerCreateInvocation) NewMetaInput() *ActionOutageHandlerCreateMetaGlobalInput {
	inv.MetaInput = &ActionOutageHandlerCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageHandlerCreateInvocation) SetMetaInput(input *ActionOutageHandlerCreateMetaGlobalInput) *ActionOutageHandlerCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageHandlerCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageHandlerCreateInvocation) Call() (*ActionOutageHandlerCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionOutageHandlerCreateInvocation) callAsBody() (*ActionOutageHandlerCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageHandlerCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Handler
	}
	return resp, err
}




func (inv *ActionOutageHandlerCreateInvocation) makeAllInputParams() *ActionOutageHandlerCreateRequest {
	return &ActionOutageHandlerCreateRequest{
		Handler: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOutageHandlerCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
		if inv.IsParameterSelected("FullName") {
			ret["full_name"] = inv.Input.FullName
		}
		if inv.IsParameterSelected("Note") {
			ret["note"] = inv.Input.Note
		}
	}

	return ret
}

func (inv *ActionOutageHandlerCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
