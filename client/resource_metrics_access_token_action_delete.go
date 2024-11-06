package client

import (
	"strings"
)

// ActionMetricsAccessTokenDelete is a type for action Metrics_access_token#Delete
type ActionMetricsAccessTokenDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionMetricsAccessTokenDelete(client *Client) *ActionMetricsAccessTokenDelete {
	return &ActionMetricsAccessTokenDelete{
		Client: client,
	}
}

// ActionMetricsAccessTokenDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionMetricsAccessTokenDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMetricsAccessTokenDeleteMetaGlobalInput) SetIncludes(value string) *ActionMetricsAccessTokenDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMetricsAccessTokenDeleteMetaGlobalInput) SetNo(value bool) *ActionMetricsAccessTokenDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMetricsAccessTokenDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMetricsAccessTokenDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionMetricsAccessTokenDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMetricsAccessTokenDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMetricsAccessTokenDeleteRequest is a type for the entire action request
type ActionMetricsAccessTokenDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionMetricsAccessTokenDeleteResponse struct {
	Action *ActionMetricsAccessTokenDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionMetricsAccessTokenDelete) Prepare() *ActionMetricsAccessTokenDeleteInvocation {
	return &ActionMetricsAccessTokenDeleteInvocation{
		Action: action,
		Path:   "/v7.0/metrics_access_tokens/{metrics_access_token_id}",
	}
}

// ActionMetricsAccessTokenDeleteInvocation is used to configure action for invocation
type ActionMetricsAccessTokenDeleteInvocation struct {
	// Pointer to the action
	Action *ActionMetricsAccessTokenDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMetricsAccessTokenDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMetricsAccessTokenDeleteInvocation) SetPathParamInt(param string, value int64) *ActionMetricsAccessTokenDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMetricsAccessTokenDeleteInvocation) SetPathParamString(param string, value string) *ActionMetricsAccessTokenDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMetricsAccessTokenDeleteInvocation) NewMetaInput() *ActionMetricsAccessTokenDeleteMetaGlobalInput {
	inv.MetaInput = &ActionMetricsAccessTokenDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMetricsAccessTokenDeleteInvocation) SetMetaInput(input *ActionMetricsAccessTokenDeleteMetaGlobalInput) *ActionMetricsAccessTokenDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMetricsAccessTokenDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMetricsAccessTokenDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMetricsAccessTokenDeleteInvocation) Call() (*ActionMetricsAccessTokenDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMetricsAccessTokenDeleteInvocation) callAsBody() (*ActionMetricsAccessTokenDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMetricsAccessTokenDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionMetricsAccessTokenDeleteInvocation) makeAllInputParams() *ActionMetricsAccessTokenDeleteRequest {
	return &ActionMetricsAccessTokenDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMetricsAccessTokenDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
