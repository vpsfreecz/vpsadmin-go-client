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
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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
	Arch           string `json:"arch"`
	CgroupVersion  string `json:"cgroup_version"`
	Distribution   string `json:"distribution"`
	Enabled        bool   `json:"enabled"`
	HypervisorType string `json:"hypervisor_type"`
	Info           string `json:"info"`
	Label          string `json:"label"`
	Order          int64  `json:"order"`
	Supported      bool   `json:"supported"`
	Variant        string `json:"variant"`
	Vendor         string `json:"vendor"`
	Version        string `json:"version"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetArch sets parameter Arch to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetArch(value string) *ActionOsTemplateCreateInput {
	in.Arch = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Arch"] = nil
	return in
}

// SetCgroupVersion sets parameter CgroupVersion to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetCgroupVersion(value string) *ActionOsTemplateCreateInput {
	in.CgroupVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CgroupVersion"] = nil
	return in
}

// SetDistribution sets parameter Distribution to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetDistribution(value string) *ActionOsTemplateCreateInput {
	in.Distribution = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Distribution"] = nil
	return in
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

// SetVariant sets parameter Variant to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetVariant(value string) *ActionOsTemplateCreateInput {
	in.Variant = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Variant"] = nil
	return in
}

// SetVendor sets parameter Vendor to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetVendor(value string) *ActionOsTemplateCreateInput {
	in.Vendor = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vendor"] = nil
	return in
}

// SetVersion sets parameter Version to value and selects it for sending
func (in *ActionOsTemplateCreateInput) SetVersion(value string) *ActionOsTemplateCreateInput {
	in.Version = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Version"] = nil
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

// UnselectParameters unsets parameters from ActionOsTemplateCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOsTemplateCreateInput) UnselectParameters(params ...string) *ActionOsTemplateCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Arch           string `json:"arch"`
	CgroupVersion  string `json:"cgroup_version"`
	Distribution   string `json:"distribution"`
	Enabled        bool   `json:"enabled"`
	HypervisorType string `json:"hypervisor_type"`
	Id             int64  `json:"id"`
	Info           string `json:"info"`
	Label          string `json:"label"`
	Name           string `json:"name"`
	Order          int64  `json:"order"`
	Supported      bool   `json:"supported"`
	Variant        string `json:"variant"`
	Vendor         string `json:"vendor"`
	Version        string `json:"version"`
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOsTemplateCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsTemplateCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		if inv.IsParameterSelected("Arch") {
			ret["arch"] = inv.Input.Arch
		}
		if inv.IsParameterSelected("CgroupVersion") {
			ret["cgroup_version"] = inv.Input.CgroupVersion
		}
		if inv.IsParameterSelected("Distribution") {
			ret["distribution"] = inv.Input.Distribution
		}
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
		if inv.IsParameterSelected("Order") {
			ret["order"] = inv.Input.Order
		}
		if inv.IsParameterSelected("Supported") {
			ret["supported"] = inv.Input.Supported
		}
		if inv.IsParameterSelected("Variant") {
			ret["variant"] = inv.Input.Variant
		}
		if inv.IsParameterSelected("Vendor") {
			ret["vendor"] = inv.Input.Vendor
		}
		if inv.IsParameterSelected("Version") {
			ret["version"] = inv.Input.Version
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
