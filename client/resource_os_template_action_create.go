package client

import ()

// ActionOsTemplateCreate is a type for action Os_template#Create
type ActionOsTemplateCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionOsTemplateCreate(client *Client) *ActionOsTemplateCreate {
	return &ActionOsTemplateCreate{
		Client: client,
	}
}

// ActionOsTemplateCreateMetaGlobalInput is a type for action global meta input parameters
type ActionOsTemplateCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsTemplateCreateMetaGlobalInput) SetIncludes(value string) *ActionOsTemplateCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsTemplateCreateMetaGlobalInput) SetNo(value bool) *ActionOsTemplateCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsTemplateCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateCreateMetaGlobalInput) SelectParameters(params ...string) *ActionOsTemplateCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsTemplateCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsTemplateCreateInput is a type for action input parameters
type ActionOsTemplateCreateInput struct {
	Enabled        bool   `json:"enabled"`
	HypervisorType string `json:"hypervisor_type"`
	Info           string `json:"info"`
	Label          string `json:"label"`
	Name           string `json:"name"`
	Order          int64  `json:"order"`
	Supported      bool   `json:"supported"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetEnabled(value bool) *ActionOsTemplateCreateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetHypervisorType sets parameter HypervisorType to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetHypervisorType(value string) *ActionOsTemplateCreateInput {
	in.HypervisorType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HypervisorType"] = nil
	return in
}

// SetInfo sets parameter Info to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetInfo(value string) *ActionOsTemplateCreateInput {
	in.Info = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Info"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetLabel(value string) *ActionOsTemplateCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetName(value string) *ActionOsTemplateCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetOrder(value int64) *ActionOsTemplateCreateInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetSupported sets parameter Supported to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetSupported(value bool) *ActionOsTemplateCreateInput {
	in.Supported = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Supported"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsTemplateCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateCreateInput) SelectParameters(params ...string) *ActionOsTemplateCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsTemplateCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsTemplateCreateRequest is a type for the entire action request
type ActionOsTemplateCreateRequest struct {
	OsTemplate map[string]interface{} `json:"os_template"`
	Meta       map[string]interface{} `json:"_meta"`
}

// ActionOsTemplateCreateOutput is a type for action output parameters
type ActionOsTemplateCreateOutput struct {
	Enabled        bool   `json:"enabled"`
	HypervisorType string `json:"hypervisor_type"`
	Id             int64  `json:"id"`
	Info           string `json:"info"`
	Label          string `json:"label"`
	Name           string `json:"name"`
	Order          int64  `json:"order"`
	Supported      bool   `json:"supported"`
}

// Type for action response, including envelope
type ActionOsTemplateCreateResponse struct {
	Action *ActionOsTemplateCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OsTemplate *ActionOsTemplateCreateOutput `json:"os_template"`
	}

	// Action output without the namespace
	Output *ActionOsTemplateCreateOutput
}

// Prepare the action for invocation
func (action *ActionOsTemplateCreate) Prepare() *ActionOsTemplateCreateInvocation {
	return &ActionOsTemplateCreateInvocation{
		Action: action,
		Path:   "/v6.0/os_templates",
	}
}

// ActionOsTemplateCreateInvocation is used to configure action for invocation
type ActionOsTemplateCreateInvocation struct {
	// Pointer to the action
	Action *ActionOsTemplateCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOsTemplateCreateInput
	// Global meta input parameters
	MetaInput *ActionOsTemplateCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOsTemplateCreateInvocation) NewInput() *ActionOsTemplateCreateInput {
	inv.Input = &ActionOsTemplateCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOsTemplateCreateInvocation) SetInput(input *ActionOsTemplateCreateInput) *ActionOsTemplateCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOsTemplateCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsTemplateCreateInvocation) NewMetaInput() *ActionOsTemplateCreateMetaGlobalInput {
	inv.MetaInput = &ActionOsTemplateCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsTemplateCreateInvocation) SetMetaInput(input *ActionOsTemplateCreateMetaGlobalInput) *ActionOsTemplateCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsTemplateCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsTemplateCreateInvocation) Call() (*ActionOsTemplateCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOsTemplateCreateInvocation) callAsBody() (*ActionOsTemplateCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOsTemplateCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OsTemplate
	}
	return resp, err
}

func (inv *ActionOsTemplateCreateInvocation) makeAllInputParams() *ActionOsTemplateCreateRequest {
	return &ActionOsTemplateCreateRequest{
		OsTemplate: inv.makeInputParams(),
		Meta:       inv.makeMetaInputParams(),
	}
}

func (inv *ActionOsTemplateCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("HypervisorType") {
			ret["hypervisor_type"] = inv.Input.HypervisorType
		}
		if inv.IsParameterSelected("Info") {
			ret["info"] = inv.Input.Info
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Order") {
			ret["order"] = inv.Input.Order
		}
		if inv.IsParameterSelected("Supported") {
			ret["supported"] = inv.Input.Supported
		}
	}

	return ret
}

func (inv *ActionOsTemplateCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
