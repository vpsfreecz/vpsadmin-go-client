package client

import ()

// ActionOsTemplateIndex is a type for action Os_template#Index
type ActionOsTemplateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOsTemplateIndex(client *Client) *ActionOsTemplateIndex {
	return &ActionOsTemplateIndex{
		Client: client,
	}
}

// ActionOsTemplateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOsTemplateIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOsTemplateIndexMetaGlobalInput) SetCount(value bool) *ActionOsTemplateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsTemplateIndexMetaGlobalInput) SetIncludes(value string) *ActionOsTemplateIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsTemplateIndexMetaGlobalInput) SetNo(value bool) *ActionOsTemplateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsTemplateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOsTemplateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsTemplateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsTemplateIndexInput is a type for action input parameters
type ActionOsTemplateIndexInput struct {
	CgroupVersion  string `json:"cgroup_version"`
	FromId         int64  `json:"from_id"`
	HypervisorType string `json:"hypervisor_type"`
	Limit          int64  `json:"limit"`
	Location       int64  `json:"location"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCgroupVersion sets parameter CgroupVersion to value and selects it for sending
func (in *ActionOsTemplateIndexInput) SetCgroupVersion(value string) *ActionOsTemplateIndexInput {
	in.CgroupVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CgroupVersion"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionOsTemplateIndexInput) SetFromId(value int64) *ActionOsTemplateIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetHypervisorType sets parameter HypervisorType to value and selects it for sending
func (in *ActionOsTemplateIndexInput) SetHypervisorType(value string) *ActionOsTemplateIndexInput {
	in.HypervisorType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HypervisorType"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOsTemplateIndexInput) SetLimit(value int64) *ActionOsTemplateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionOsTemplateIndexInput) SetLocation(value int64) *ActionOsTemplateIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionOsTemplateIndexInput) SetLocationNil(set bool) *ActionOsTemplateIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Location"] = nil
		in.SelectParameters("Location")
	} else {
		delete(in._nilParameters, "Location")
	}
	return in
}

// SelectParameters sets parameters from ActionOsTemplateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateIndexInput) SelectParameters(params ...string) *ActionOsTemplateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOsTemplateIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOsTemplateIndexInput) UnselectParameters(params ...string) *ActionOsTemplateIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOsTemplateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsTemplateIndexOutput is a type for action output parameters
type ActionOsTemplateIndexOutput struct {
	Arch           string                    `json:"arch"`
	CgroupVersion  string                    `json:"cgroup_version"`
	Config         string                    `json:"config"`
	Distribution   string                    `json:"distribution"`
	Enabled        bool                      `json:"enabled"`
	HypervisorType string                    `json:"hypervisor_type"`
	Id             int64                     `json:"id"`
	Info           string                    `json:"info"`
	Label          string                    `json:"label"`
	Name           string                    `json:"name"`
	Order          int64                     `json:"order"`
	OsFamily       *ActionOsFamilyShowOutput `json:"os_family"`
	Supported      bool                      `json:"supported"`
	Variant        string                    `json:"variant"`
	Vendor         string                    `json:"vendor"`
	Version        string                    `json:"version"`
}

// Type for action response, including envelope
type ActionOsTemplateIndexResponse struct {
	Action *ActionOsTemplateIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OsTemplates []*ActionOsTemplateIndexOutput `json:"os_templates"`
	}

	// Action output without the namespace
	Output []*ActionOsTemplateIndexOutput
}

// Prepare the action for invocation
func (action *ActionOsTemplateIndex) Prepare() *ActionOsTemplateIndexInvocation {
	return &ActionOsTemplateIndexInvocation{
		Action: action,
		Path:   "/v7.0/os_templates",
	}
}

// ActionOsTemplateIndexInvocation is used to configure action for invocation
type ActionOsTemplateIndexInvocation struct {
	// Pointer to the action
	Action *ActionOsTemplateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOsTemplateIndexInput
	// Global meta input parameters
	MetaInput *ActionOsTemplateIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOsTemplateIndexInvocation) NewInput() *ActionOsTemplateIndexInput {
	inv.Input = &ActionOsTemplateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOsTemplateIndexInvocation) SetInput(input *ActionOsTemplateIndexInput) *ActionOsTemplateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOsTemplateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOsTemplateIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsTemplateIndexInvocation) NewMetaInput() *ActionOsTemplateIndexMetaGlobalInput {
	inv.MetaInput = &ActionOsTemplateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsTemplateIndexInvocation) SetMetaInput(input *ActionOsTemplateIndexMetaGlobalInput) *ActionOsTemplateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsTemplateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsTemplateIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsTemplateIndexInvocation) Call() (*ActionOsTemplateIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOsTemplateIndexInvocation) callAsQuery() (*ActionOsTemplateIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOsTemplateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OsTemplates
	}
	return resp, err
}

func (inv *ActionOsTemplateIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("CgroupVersion") {
			ret["os_template[cgroup_version]"] = inv.Input.CgroupVersion
		}
		if inv.IsParameterSelected("FromId") {
			ret["os_template[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("HypervisorType") {
			ret["os_template[hypervisor_type]"] = inv.Input.HypervisorType
		}
		if inv.IsParameterSelected("Limit") {
			ret["os_template[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["os_template[location]"] = convertInt64ToString(inv.Input.Location)
		}
	}
}

func (inv *ActionOsTemplateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
