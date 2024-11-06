package client

import (
	"strings"
)

// ActionOsFamilyShow is a type for action Os_family#Show
type ActionOsFamilyShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOsFamilyShow(client *Client) *ActionOsFamilyShow {
	return &ActionOsFamilyShow{
		Client: client,
	}
}

// ActionOsFamilyShowMetaGlobalInput is a type for action global meta input parameters
type ActionOsFamilyShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOsFamilyShowMetaGlobalInput) SetIncludes(value string) *ActionOsFamilyShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOsFamilyShowMetaGlobalInput) SetNo(value bool) *ActionOsFamilyShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOsFamilyShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOsFamilyShowMetaGlobalInput) SelectParameters(params ...string) *ActionOsFamilyShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOsFamilyShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOsFamilyShowOutput is a type for action output parameters
type ActionOsFamilyShowOutput struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Label       string `json:"label"`
}

// Type for action response, including envelope
type ActionOsFamilyShowResponse struct {
	Action *ActionOsFamilyShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OsFamily *ActionOsFamilyShowOutput `json:"os_family"`
	}

	// Action output without the namespace
	Output *ActionOsFamilyShowOutput
}

// Prepare the action for invocation
func (action *ActionOsFamilyShow) Prepare() *ActionOsFamilyShowInvocation {
	return &ActionOsFamilyShowInvocation{
		Action: action,
		Path:   "/v7.0/os_families/{os_family_id}",
	}
}

// ActionOsFamilyShowInvocation is used to configure action for invocation
type ActionOsFamilyShowInvocation struct {
	// Pointer to the action
	Action *ActionOsFamilyShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOsFamilyShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOsFamilyShowInvocation) SetPathParamInt(param string, value int64) *ActionOsFamilyShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOsFamilyShowInvocation) SetPathParamString(param string, value string) *ActionOsFamilyShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOsFamilyShowInvocation) NewMetaInput() *ActionOsFamilyShowMetaGlobalInput {
	inv.MetaInput = &ActionOsFamilyShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOsFamilyShowInvocation) SetMetaInput(input *ActionOsFamilyShowMetaGlobalInput) *ActionOsFamilyShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOsFamilyShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOsFamilyShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOsFamilyShowInvocation) Call() (*ActionOsFamilyShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOsFamilyShowInvocation) callAsQuery() (*ActionOsFamilyShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOsFamilyShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OsFamily
	}
	return resp, err
}

func (inv *ActionOsFamilyShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
