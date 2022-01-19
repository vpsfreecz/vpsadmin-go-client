package client

import (
	"strings"
)

// ActionNewsLogDelete is a type for action News_log#Delete
type ActionNewsLogDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionNewsLogDelete(client *Client) *ActionNewsLogDelete {
	return &ActionNewsLogDelete{
		Client: client,
	}
}

// ActionNewsLogDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionNewsLogDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNewsLogDeleteMetaGlobalInput) SetIncludes(value string) *ActionNewsLogDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNewsLogDeleteMetaGlobalInput) SetNo(value bool) *ActionNewsLogDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionNewsLogDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNewsLogDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogDeleteRequest is a type for the entire action request
type ActionNewsLogDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionNewsLogDeleteResponse struct {
	Action *ActionNewsLogDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionNewsLogDelete) Prepare() *ActionNewsLogDeleteInvocation {
	return &ActionNewsLogDeleteInvocation{
		Action: action,
		Path:   "/v6.0/news_logs/{news_log_id}",
	}
}

// ActionNewsLogDeleteInvocation is used to configure action for invocation
type ActionNewsLogDeleteInvocation struct {
	// Pointer to the action
	Action *ActionNewsLogDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNewsLogDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNewsLogDeleteInvocation) SetPathParamInt(param string, value int64) *ActionNewsLogDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNewsLogDeleteInvocation) SetPathParamString(param string, value string) *ActionNewsLogDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNewsLogDeleteInvocation) NewMetaInput() *ActionNewsLogDeleteMetaGlobalInput {
	inv.MetaInput = &ActionNewsLogDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNewsLogDeleteInvocation) SetMetaInput(input *ActionNewsLogDeleteMetaGlobalInput) *ActionNewsLogDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNewsLogDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNewsLogDeleteInvocation) Call() (*ActionNewsLogDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionNewsLogDeleteInvocation) callAsBody() (*ActionNewsLogDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNewsLogDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionNewsLogDeleteInvocation) makeAllInputParams() *ActionNewsLogDeleteRequest {
	return &ActionNewsLogDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNewsLogDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
