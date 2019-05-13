package client

import (
	"strings"
)

// ActionOsTemplateUpdate is a type for action Os_template#Update
type ActionOsTemplateUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionOsTemplateUpdate(client *Client) *ActionOsTemplateUpdate {
	return &ActionOsTemplateUpdate{
		Client: client,
	}
}

// ActionOsTemplateUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionOsTemplateUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsTemplateUpdateMetaGlobalInput) SetNo(value bool) *ActionOsTemplateUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsTemplateUpdateMetaGlobalInput) SetIncludes(value string) *ActionOsTemplateUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsTemplateUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionOsTemplateUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsTemplateUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsTemplateUpdateInput is a type for action input parameters
type ActionOsTemplateUpdateInput struct {
	Label string `json:"label"`
	Info string `json:"info"`
	Enabled bool `json:"enabled"`
	Supported bool `json:"supported"`
	Order int64 `json:"order"`
	HypervisorType string `json:"hypervisor_type"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetLabel(value string) *ActionOsTemplateUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetInfo sets parameter Info to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetInfo(value string) *ActionOsTemplateUpdateInput {
	in.Info = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Info"] = nil
	return in
}
// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetEnabled(value bool) *ActionOsTemplateUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}
// SetSupported sets parameter Supported to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetSupported(value bool) *ActionOsTemplateUpdateInput {
	in.Supported = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Supported"] = nil
	return in
}
// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetOrder(value int64) *ActionOsTemplateUpdateInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}
// SetHypervisorType sets parameter HypervisorType to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetHypervisorType(value string) *ActionOsTemplateUpdateInput {
	in.HypervisorType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HypervisorType"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsTemplateUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateUpdateInput) SelectParameters(params ...string) *ActionOsTemplateUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsTemplateUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsTemplateUpdateRequest is a type for the entire action request
type ActionOsTemplateUpdateRequest struct {
	OsTemplate map[string]interface{} `json:"os_template"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionOsTemplateUpdateOutput is a type for action output parameters
type ActionOsTemplateUpdateOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	Info string `json:"info"`
	Enabled bool `json:"enabled"`
	Supported bool `json:"supported"`
	Order int64 `json:"order"`
	HypervisorType string `json:"hypervisor_type"`
}


// Type for action response, including envelope
type ActionOsTemplateUpdateResponse struct {
	Action *ActionOsTemplateUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OsTemplate *ActionOsTemplateUpdateOutput `json:"os_template"`
	}

	// Action output without the namespace
	Output *ActionOsTemplateUpdateOutput
}


// Prepare the action for invocation
func (action *ActionOsTemplateUpdate) Prepare() *ActionOsTemplateUpdateInvocation {
	return &ActionOsTemplateUpdateInvocation{
		Action: action,
		Path: "/v5.0/os_templates/{os_template_id}",
	}
}

// ActionOsTemplateUpdateInvocation is used to configure action for invocation
type ActionOsTemplateUpdateInvocation struct {
	// Pointer to the action
	Action *ActionOsTemplateUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOsTemplateUpdateInput
	// Global meta input parameters
	MetaInput *ActionOsTemplateUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOsTemplateUpdateInvocation) SetPathParamInt(param string, value int64) *ActionOsTemplateUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOsTemplateUpdateInvocation) SetPathParamString(param string, value string) *ActionOsTemplateUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOsTemplateUpdateInvocation) NewInput() *ActionOsTemplateUpdateInput {
	inv.Input = &ActionOsTemplateUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOsTemplateUpdateInvocation) SetInput(input *ActionOsTemplateUpdateInput) *ActionOsTemplateUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOsTemplateUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsTemplateUpdateInvocation) NewMetaInput() *ActionOsTemplateUpdateMetaGlobalInput {
	inv.MetaInput = &ActionOsTemplateUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsTemplateUpdateInvocation) SetMetaInput(input *ActionOsTemplateUpdateMetaGlobalInput) *ActionOsTemplateUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsTemplateUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsTemplateUpdateInvocation) Call() (*ActionOsTemplateUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionOsTemplateUpdateInvocation) callAsBody() (*ActionOsTemplateUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOsTemplateUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OsTemplate
	}
	return resp, err
}




func (inv *ActionOsTemplateUpdateInvocation) makeAllInputParams() *ActionOsTemplateUpdateRequest {
	return &ActionOsTemplateUpdateRequest{
		OsTemplate: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOsTemplateUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Info") {
			ret["info"] = inv.Input.Info
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("Supported") {
			ret["supported"] = inv.Input.Supported
		}
		if inv.IsParameterSelected("Order") {
			ret["order"] = inv.Input.Order
		}
		if inv.IsParameterSelected("HypervisorType") {
			ret["hypervisor_type"] = inv.Input.HypervisorType
		}
	}

	return ret
}

func (inv *ActionOsTemplateUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
