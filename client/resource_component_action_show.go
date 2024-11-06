package client

import (
	"strings"
)

// ActionComponentShow is a type for action Component#Show
type ActionComponentShow struct {
	// Pointer to client
	Client *Client
}

func NewActionComponentShow(client *Client) *ActionComponentShow {
	return &ActionComponentShow{
		Client: client,
	}
}

// ActionComponentShowMetaGlobalInput is a type for action global meta input parameters
type ActionComponentShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionComponentShowMetaGlobalInput) SetIncludes(value string) *ActionComponentShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionComponentShowMetaGlobalInput) SetNo(value bool) *ActionComponentShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionComponentShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionComponentShowMetaGlobalInput) SelectParameters(params ...string) *ActionComponentShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionComponentShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionComponentShowOutput is a type for action output parameters
type ActionComponentShowOutput struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Label       string `json:"label"`
	Name        string `json:"name"`
}

// Type for action response, including envelope
type ActionComponentShowResponse struct {
	Action *ActionComponentShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Component *ActionComponentShowOutput `json:"component"`
	}

	// Action output without the namespace
	Output *ActionComponentShowOutput
}

// Prepare the action for invocation
func (action *ActionComponentShow) Prepare() *ActionComponentShowInvocation {
	return &ActionComponentShowInvocation{
		Action: action,
		Path:   "/v7.0/components/{component_id}",
	}
}

// ActionComponentShowInvocation is used to configure action for invocation
type ActionComponentShowInvocation struct {
	// Pointer to the action
	Action *ActionComponentShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionComponentShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionComponentShowInvocation) SetPathParamInt(param string, value int64) *ActionComponentShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionComponentShowInvocation) SetPathParamString(param string, value string) *ActionComponentShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionComponentShowInvocation) NewMetaInput() *ActionComponentShowMetaGlobalInput {
	inv.MetaInput = &ActionComponentShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionComponentShowInvocation) SetMetaInput(input *ActionComponentShowMetaGlobalInput) *ActionComponentShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionComponentShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionComponentShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionComponentShowInvocation) Call() (*ActionComponentShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionComponentShowInvocation) callAsQuery() (*ActionComponentShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionComponentShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Component
	}
	return resp, err
}

func (inv *ActionComponentShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
