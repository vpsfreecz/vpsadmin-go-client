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
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsTemplateUpdateMetaGlobalInput) SetNo(value bool) *ActionOsTemplateUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	Arch              string `json:"arch"`
	CgroupVersion     string `json:"cgroup_version"`
	Config            string `json:"config"`
	Distribution      string `json:"distribution"`
	EnableCloudInit   bool   `json:"enable_cloud_init"`
	EnableScript      bool   `json:"enable_script"`
	Enabled           bool   `json:"enabled"`
	HypervisorType    string `json:"hypervisor_type"`
	Info              string `json:"info"`
	Label             string `json:"label"`
	ManageDnsResolver bool   `json:"manage_dns_resolver"`
	ManageHostname    bool   `json:"manage_hostname"`
	Order             int64  `json:"order"`
	OsFamily          int64  `json:"os_family"`
	Supported         bool   `json:"supported"`
	Variant           string `json:"variant"`
	Vendor            string `json:"vendor"`
	Version           string `json:"version"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetArch sets parameter Arch to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetArch(value string) *ActionOsTemplateUpdateInput {
	in.Arch = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Arch"] = nil
	return in
}

// SetCgroupVersion sets parameter CgroupVersion to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetCgroupVersion(value string) *ActionOsTemplateUpdateInput {
	in.CgroupVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CgroupVersion"] = nil
	return in
}

// SetConfig sets parameter Config to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetConfig(value string) *ActionOsTemplateUpdateInput {
	in.Config = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Config"] = nil
	return in
}

// SetDistribution sets parameter Distribution to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetDistribution(value string) *ActionOsTemplateUpdateInput {
	in.Distribution = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Distribution"] = nil
	return in
}

// SetEnableCloudInit sets parameter EnableCloudInit to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetEnableCloudInit(value bool) *ActionOsTemplateUpdateInput {
	in.EnableCloudInit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableCloudInit"] = nil
	return in
}

// SetEnableScript sets parameter EnableScript to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetEnableScript(value bool) *ActionOsTemplateUpdateInput {
	in.EnableScript = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableScript"] = nil
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

// SetHypervisorType sets parameter HypervisorType to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetHypervisorType(value string) *ActionOsTemplateUpdateInput {
	in.HypervisorType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HypervisorType"] = nil
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

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetLabel(value string) *ActionOsTemplateUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetManageDnsResolver sets parameter ManageDnsResolver to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetManageDnsResolver(value bool) *ActionOsTemplateUpdateInput {
	in.ManageDnsResolver = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ManageDnsResolver"] = nil
	return in
}

// SetManageHostname sets parameter ManageHostname to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetManageHostname(value bool) *ActionOsTemplateUpdateInput {
	in.ManageHostname = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ManageHostname"] = nil
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

// SetOsFamily sets parameter OsFamily to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetOsFamily(value int64) *ActionOsTemplateUpdateInput {
	in.OsFamily = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetOsFamilyNil(false)
	in._selectedParameters["OsFamily"] = nil
	return in
}

// SetOsFamilyNil sets parameter OsFamily to nil and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetOsFamilyNil(set bool) *ActionOsTemplateUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["OsFamily"] = nil
		in.SelectParameters("OsFamily")
	} else {
		delete(in._nilParameters, "OsFamily")
	}
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

// SetVariant sets parameter Variant to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetVariant(value string) *ActionOsTemplateUpdateInput {
	in.Variant = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Variant"] = nil
	return in
}

// SetVendor sets parameter Vendor to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetVendor(value string) *ActionOsTemplateUpdateInput {
	in.Vendor = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vendor"] = nil
	return in
}

// SetVersion sets parameter Version to value and selects it for sending
func (in *ActionOsTemplateUpdateInput) SetVersion(value string) *ActionOsTemplateUpdateInput {
	in.Version = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Version"] = nil
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

// UnselectParameters unsets parameters from ActionOsTemplateUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOsTemplateUpdateInput) UnselectParameters(params ...string) *ActionOsTemplateUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Meta       map[string]interface{} `json:"_meta"`
}

// ActionOsTemplateUpdateOutput is a type for action output parameters
type ActionOsTemplateUpdateOutput struct {
	Arch              string                    `json:"arch"`
	CgroupVersion     string                    `json:"cgroup_version"`
	Config            string                    `json:"config"`
	Distribution      string                    `json:"distribution"`
	EnableCloudInit   bool                      `json:"enable_cloud_init"`
	EnableScript      bool                      `json:"enable_script"`
	Enabled           bool                      `json:"enabled"`
	HypervisorType    string                    `json:"hypervisor_type"`
	Id                int64                     `json:"id"`
	Info              string                    `json:"info"`
	Label             string                    `json:"label"`
	ManageDnsResolver bool                      `json:"manage_dns_resolver"`
	ManageHostname    bool                      `json:"manage_hostname"`
	Name              string                    `json:"name"`
	Order             int64                     `json:"order"`
	OsFamily          *ActionOsFamilyShowOutput `json:"os_family"`
	Supported         bool                      `json:"supported"`
	Variant           string                    `json:"variant"`
	Vendor            string                    `json:"vendor"`
	Version           string                    `json:"version"`
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
		Path:   "/v7.0/os_templates/{os_template_id}",
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOsTemplateUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsTemplateUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		Meta:       inv.makeMetaInputParams(),
	}
}

func (inv *ActionOsTemplateUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Arch") {
			ret["arch"] = inv.Input.Arch
		}
		if inv.IsParameterSelected("CgroupVersion") {
			ret["cgroup_version"] = inv.Input.CgroupVersion
		}
		if inv.IsParameterSelected("Config") {
			ret["config"] = inv.Input.Config
		}
		if inv.IsParameterSelected("Distribution") {
			ret["distribution"] = inv.Input.Distribution
		}
		if inv.IsParameterSelected("EnableCloudInit") {
			ret["enable_cloud_init"] = inv.Input.EnableCloudInit
		}
		if inv.IsParameterSelected("EnableScript") {
			ret["enable_script"] = inv.Input.EnableScript
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
		if inv.IsParameterSelected("ManageDnsResolver") {
			ret["manage_dns_resolver"] = inv.Input.ManageDnsResolver
		}
		if inv.IsParameterSelected("ManageHostname") {
			ret["manage_hostname"] = inv.Input.ManageHostname
		}
		if inv.IsParameterSelected("Order") {
			ret["order"] = inv.Input.Order
		}
		if inv.IsParameterSelected("OsFamily") {
			if inv.IsParameterNil("OsFamily") {
				ret["os_family"] = nil
			} else {
				ret["os_family"] = inv.Input.OsFamily
			}
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

func (inv *ActionOsTemplateUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
