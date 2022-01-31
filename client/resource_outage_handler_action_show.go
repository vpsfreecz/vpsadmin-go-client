package client

import (
	"strings"
)

// ActionOutageHandlerShow is a type for action Outage.Handler#Show
type ActionOutageHandlerShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageHandlerShow(client *Client) *ActionOutageHandlerShow {
	return &ActionOutageHandlerShow{
		Client: client,
	}
}

// ActionOutageHandlerShowMetaGlobalInput is a type for action global meta input parameters
type ActionOutageHandlerShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageHandlerShowMetaGlobalInput) SetIncludes(value string) *ActionOutageHandlerShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageHandlerShowMetaGlobalInput) SetNo(value bool) *ActionOutageHandlerShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageHandlerShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageHandlerShowMetaGlobalInput) SelectParameters(params ...string) *ActionOutageHandlerShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageHandlerShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageHandlerShowOutput is a type for action output parameters
type ActionOutageHandlerShowOutput struct {
	FullName string                `json:"full_name"`
	Id       int64                 `json:"id"`
	Note     string                `json:"note"`
	User     *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionOutageHandlerShowResponse struct {
	Action *ActionOutageHandlerShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Handler *ActionOutageHandlerShowOutput `json:"handler"`
	}

	// Action output without the namespace
	Output *ActionOutageHandlerShowOutput
}

// Prepare the action for invocation
func (action *ActionOutageHandlerShow) Prepare() *ActionOutageHandlerShowInvocation {
	return &ActionOutageHandlerShowInvocation{
		Action: action,
		Path:   "/v6.0/outages/{outage_id}/handlers/{handler_id}",
	}
}

// ActionOutageHandlerShowInvocation is used to configure action for invocation
type ActionOutageHandlerShowInvocation struct {
	// Pointer to the action
	Action *ActionOutageHandlerShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOutageHandlerShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageHandlerShowInvocation) SetPathParamInt(param string, value int64) *ActionOutageHandlerShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageHandlerShowInvocation) SetPathParamString(param string, value string) *ActionOutageHandlerShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageHandlerShowInvocation) NewMetaInput() *ActionOutageHandlerShowMetaGlobalInput {
	inv.MetaInput = &ActionOutageHandlerShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageHandlerShowInvocation) SetMetaInput(input *ActionOutageHandlerShowMetaGlobalInput) *ActionOutageHandlerShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageHandlerShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageHandlerShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageHandlerShowInvocation) Call() (*ActionOutageHandlerShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOutageHandlerShowInvocation) callAsQuery() (*ActionOutageHandlerShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOutageHandlerShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Handler
	}
	return resp, err
}

func (inv *ActionOutageHandlerShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
