package client

import ()

// ActionOsFamilyCreate is a type for action Os_family#Create
type ActionOsFamilyCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionOsFamilyCreate(client *Client) *ActionOsFamilyCreate {
	return &ActionOsFamilyCreate{
		Client: client,
	}
}

// ActionOsFamilyCreateMetaGlobalInput is a type for action global meta input parameters
type ActionOsFamilyCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsFamilyCreateMetaGlobalInput) SetIncludes(value string) *ActionOsFamilyCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsFamilyCreateMetaGlobalInput) SetNo(value bool) *ActionOsFamilyCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsFamilyCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsFamilyCreateMetaGlobalInput) SelectParameters(params ...string) *ActionOsFamilyCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsFamilyCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsFamilyCreateInput is a type for action input parameters
type ActionOsFamilyCreateInput struct {
	Description string `json:"description"`
	Label       string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDescription sets parameter Description to value and selects it for sending
func (in *ActionOsFamilyCreateInput) SetDescription(value string) *ActionOsFamilyCreateInput {
	in.Description = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Description"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionOsFamilyCreateInput) SetLabel(value string) *ActionOsFamilyCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsFamilyCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsFamilyCreateInput) SelectParameters(params ...string) *ActionOsFamilyCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOsFamilyCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOsFamilyCreateInput) UnselectParameters(params ...string) *ActionOsFamilyCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOsFamilyCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsFamilyCreateRequest is a type for the entire action request
type ActionOsFamilyCreateRequest struct {
	OsFamily map[string]interface{} `json:"os_family"`
	Meta     map[string]interface{} `json:"_meta"`
}

// ActionOsFamilyCreateOutput is a type for action output parameters
type ActionOsFamilyCreateOutput struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Label       string `json:"label"`
}

// Type for action response, including envelope
type ActionOsFamilyCreateResponse struct {
	Action *ActionOsFamilyCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OsFamily *ActionOsFamilyCreateOutput `json:"os_family"`
	}

	// Action output without the namespace
	Output *ActionOsFamilyCreateOutput
}

// Prepare the action for invocation
func (action *ActionOsFamilyCreate) Prepare() *ActionOsFamilyCreateInvocation {
	return &ActionOsFamilyCreateInvocation{
		Action: action,
		Path:   "/v7.0/os_families",
	}
}

// ActionOsFamilyCreateInvocation is used to configure action for invocation
type ActionOsFamilyCreateInvocation struct {
	// Pointer to the action
	Action *ActionOsFamilyCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOsFamilyCreateInput
	// Global meta input parameters
	MetaInput *ActionOsFamilyCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOsFamilyCreateInvocation) NewInput() *ActionOsFamilyCreateInput {
	inv.Input = &ActionOsFamilyCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOsFamilyCreateInvocation) SetInput(input *ActionOsFamilyCreateInput) *ActionOsFamilyCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOsFamilyCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOsFamilyCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsFamilyCreateInvocation) NewMetaInput() *ActionOsFamilyCreateMetaGlobalInput {
	inv.MetaInput = &ActionOsFamilyCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsFamilyCreateInvocation) SetMetaInput(input *ActionOsFamilyCreateMetaGlobalInput) *ActionOsFamilyCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsFamilyCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsFamilyCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsFamilyCreateInvocation) Call() (*ActionOsFamilyCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOsFamilyCreateInvocation) callAsBody() (*ActionOsFamilyCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOsFamilyCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OsFamily
	}
	return resp, err
}

func (inv *ActionOsFamilyCreateInvocation) makeAllInputParams() *ActionOsFamilyCreateRequest {
	return &ActionOsFamilyCreateRequest{
		OsFamily: inv.makeInputParams(),
		Meta:     inv.makeMetaInputParams(),
	}
}

func (inv *ActionOsFamilyCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionOsFamilyCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
