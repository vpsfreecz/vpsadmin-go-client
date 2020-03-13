package client

import (
)

// ActionUserNamespaceMapCreate is a type for action User_namespace_map#Create
type ActionUserNamespaceMapCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapCreate(client *Client) *ActionUserNamespaceMapCreate {
	return &ActionUserNamespaceMapCreate{
		Client: client,
	}
}

// ActionUserNamespaceMapCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapCreateMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapCreateMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapCreateInput is a type for action input parameters
type ActionUserNamespaceMapCreateInput struct {
	UserNamespace int64 `json:"user_namespace"`
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetUserNamespace sets parameter UserNamespace to value and selects it for sending
func (in *ActionUserNamespaceMapCreateInput) SetUserNamespace(value int64) *ActionUserNamespaceMapCreateInput {
	in.UserNamespace = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserNamespace"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionUserNamespaceMapCreateInput) SetLabel(value string) *ActionUserNamespaceMapCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapCreateInput) SelectParameters(params ...string) *ActionUserNamespaceMapCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapCreateRequest is a type for the entire action request
type ActionUserNamespaceMapCreateRequest struct {
	UserNamespaceMap map[string]interface{} `json:"user_namespace_map"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserNamespaceMapCreateOutput is a type for action output parameters
type ActionUserNamespaceMapCreateOutput struct {
	Id int64 `json:"id"`
	UserNamespace *ActionUserNamespaceShowOutput `json:"user_namespace"`
	Label string `json:"label"`
}


// Type for action response, including envelope
type ActionUserNamespaceMapCreateResponse struct {
	Action *ActionUserNamespaceMapCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserNamespaceMap *ActionUserNamespaceMapCreateOutput `json:"user_namespace_map"`
	}

	// Action output without the namespace
	Output *ActionUserNamespaceMapCreateOutput
}


// Prepare the action for invocation
func (action *ActionUserNamespaceMapCreate) Prepare() *ActionUserNamespaceMapCreateInvocation {
	return &ActionUserNamespaceMapCreateInvocation{
		Action: action,
		Path: "/v6.0/user_namespace_maps",
	}
}

// ActionUserNamespaceMapCreateInvocation is used to configure action for invocation
type ActionUserNamespaceMapCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNamespaceMapCreateInput
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNamespaceMapCreateInvocation) NewInput() *ActionUserNamespaceMapCreateInput {
	inv.Input = &ActionUserNamespaceMapCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNamespaceMapCreateInvocation) SetInput(input *ActionUserNamespaceMapCreateInput) *ActionUserNamespaceMapCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNamespaceMapCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapCreateInvocation) NewMetaInput() *ActionUserNamespaceMapCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapCreateInvocation) SetMetaInput(input *ActionUserNamespaceMapCreateMetaGlobalInput) *ActionUserNamespaceMapCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceMapCreateInvocation) Call() (*ActionUserNamespaceMapCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserNamespaceMapCreateInvocation) callAsBody() (*ActionUserNamespaceMapCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserNamespaceMapCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserNamespaceMap
	}
	return resp, err
}




func (inv *ActionUserNamespaceMapCreateInvocation) makeAllInputParams() *ActionUserNamespaceMapCreateRequest {
	return &ActionUserNamespaceMapCreateRequest{
		UserNamespaceMap: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNamespaceMapCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("UserNamespace") {
			ret["user_namespace"] = inv.Input.UserNamespace
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionUserNamespaceMapCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
