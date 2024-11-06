package client

import (
	"strings"
)

// ActionOsTemplateShow is a type for action Os_template#Show
type ActionOsTemplateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOsTemplateShow(client *Client) *ActionOsTemplateShow {
	return &ActionOsTemplateShow{
		Client: client,
	}
}

// ActionOsTemplateShowMetaGlobalInput is a type for action global meta input parameters
type ActionOsTemplateShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsTemplateShowMetaGlobalInput) SetIncludes(value string) *ActionOsTemplateShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsTemplateShowMetaGlobalInput) SetNo(value bool) *ActionOsTemplateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsTemplateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsTemplateShowMetaGlobalInput) SelectParameters(params ...string) *ActionOsTemplateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsTemplateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsTemplateShowOutput is a type for action output parameters
type ActionOsTemplateShowOutput struct {
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
type ActionOsTemplateShowResponse struct {
	Action *ActionOsTemplateShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OsTemplate *ActionOsTemplateShowOutput `json:"os_template"`
	}

	// Action output without the namespace
	Output *ActionOsTemplateShowOutput
}

// Prepare the action for invocation
func (action *ActionOsTemplateShow) Prepare() *ActionOsTemplateShowInvocation {
	return &ActionOsTemplateShowInvocation{
		Action: action,
		Path:   "/v7.0/os_templates/{os_template_id}",
	}
}

// ActionOsTemplateShowInvocation is used to configure action for invocation
type ActionOsTemplateShowInvocation struct {
	// Pointer to the action
	Action *ActionOsTemplateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOsTemplateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOsTemplateShowInvocation) SetPathParamInt(param string, value int64) *ActionOsTemplateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOsTemplateShowInvocation) SetPathParamString(param string, value string) *ActionOsTemplateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsTemplateShowInvocation) NewMetaInput() *ActionOsTemplateShowMetaGlobalInput {
	inv.MetaInput = &ActionOsTemplateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsTemplateShowInvocation) SetMetaInput(input *ActionOsTemplateShowMetaGlobalInput) *ActionOsTemplateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsTemplateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsTemplateShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsTemplateShowInvocation) Call() (*ActionOsTemplateShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOsTemplateShowInvocation) callAsQuery() (*ActionOsTemplateShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOsTemplateShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OsTemplate
	}
	return resp, err
}

func (inv *ActionOsTemplateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
