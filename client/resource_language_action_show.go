package client

import (
	"strings"
)

// ActionLanguageShow is a type for action Language#Show
type ActionLanguageShow struct {
	// Pointer to client
	Client *Client
}

func NewActionLanguageShow(client *Client) *ActionLanguageShow {
	return &ActionLanguageShow{
		Client: client,
	}
}

// ActionLanguageShowMetaGlobalInput is a type for action global meta input parameters
type ActionLanguageShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLanguageShowMetaGlobalInput) SetIncludes(value string) *ActionLanguageShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLanguageShowMetaGlobalInput) SetNo(value bool) *ActionLanguageShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionLanguageShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLanguageShowMetaGlobalInput) SelectParameters(params ...string) *ActionLanguageShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLanguageShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLanguageShowOutput is a type for action output parameters
type ActionLanguageShowOutput struct {
	Code  string `json:"code"`
	Id    int64  `json:"id"`
	Label string `json:"label"`
}

// Type for action response, including envelope
type ActionLanguageShowResponse struct {
	Action *ActionLanguageShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Language *ActionLanguageShowOutput `json:"language"`
	}

	// Action output without the namespace
	Output *ActionLanguageShowOutput
}

// Prepare the action for invocation
func (action *ActionLanguageShow) Prepare() *ActionLanguageShowInvocation {
	return &ActionLanguageShowInvocation{
		Action: action,
		Path:   "/v6.0/languages/{language_id}",
	}
}

// ActionLanguageShowInvocation is used to configure action for invocation
type ActionLanguageShowInvocation struct {
	// Pointer to the action
	Action *ActionLanguageShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionLanguageShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionLanguageShowInvocation) SetPathParamInt(param string, value int64) *ActionLanguageShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionLanguageShowInvocation) SetPathParamString(param string, value string) *ActionLanguageShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLanguageShowInvocation) NewMetaInput() *ActionLanguageShowMetaGlobalInput {
	inv.MetaInput = &ActionLanguageShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLanguageShowInvocation) SetMetaInput(input *ActionLanguageShowMetaGlobalInput) *ActionLanguageShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLanguageShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLanguageShowInvocation) Call() (*ActionLanguageShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionLanguageShowInvocation) callAsQuery() (*ActionLanguageShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionLanguageShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Language
	}
	return resp, err
}

func (inv *ActionLanguageShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
