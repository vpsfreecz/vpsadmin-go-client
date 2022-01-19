package client

import (
	"strings"
)

// ActionLocationNetworkDelete is a type for action Location_network#Delete
type ActionLocationNetworkDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationNetworkDelete(client *Client) *ActionLocationNetworkDelete {
	return &ActionLocationNetworkDelete{
		Client: client,
	}
}

// ActionLocationNetworkDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionLocationNetworkDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationNetworkDeleteMetaGlobalInput) SetIncludes(value string) *ActionLocationNetworkDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationNetworkDeleteMetaGlobalInput) SetNo(value bool) *ActionLocationNetworkDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationNetworkDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationNetworkDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionLocationNetworkDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationNetworkDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationNetworkDeleteRequest is a type for the entire action request
type ActionLocationNetworkDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionLocationNetworkDeleteResponse struct {
	Action *ActionLocationNetworkDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionLocationNetworkDelete) Prepare() *ActionLocationNetworkDeleteInvocation {
	return &ActionLocationNetworkDeleteInvocation{
		Action: action,
		Path:   "/v6.0/location_networks/{location_network_id}",
	}
}

// ActionLocationNetworkDeleteInvocation is used to configure action for invocation
type ActionLocationNetworkDeleteInvocation struct {
	// Pointer to the action
	Action *ActionLocationNetworkDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionLocationNetworkDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionLocationNetworkDeleteInvocation) SetPathParamInt(param string, value int64) *ActionLocationNetworkDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionLocationNetworkDeleteInvocation) SetPathParamString(param string, value string) *ActionLocationNetworkDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLocationNetworkDeleteInvocation) NewMetaInput() *ActionLocationNetworkDeleteMetaGlobalInput {
	inv.MetaInput = &ActionLocationNetworkDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationNetworkDeleteInvocation) SetMetaInput(input *ActionLocationNetworkDeleteMetaGlobalInput) *ActionLocationNetworkDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationNetworkDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationNetworkDeleteInvocation) Call() (*ActionLocationNetworkDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionLocationNetworkDeleteInvocation) callAsBody() (*ActionLocationNetworkDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionLocationNetworkDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionLocationNetworkDeleteInvocation) makeAllInputParams() *ActionLocationNetworkDeleteRequest {
	return &ActionLocationNetworkDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionLocationNetworkDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
