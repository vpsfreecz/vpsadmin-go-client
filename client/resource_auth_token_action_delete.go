package client

import (
	"strings"
)

// ActionAuthTokenDelete is a type for action Auth_token#Delete
type ActionAuthTokenDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionAuthTokenDelete(client *Client) *ActionAuthTokenDelete {
	return &ActionAuthTokenDelete{
		Client: client,
	}
}

// ActionAuthTokenDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionAuthTokenDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionAuthTokenDeleteMetaGlobalInput) SetNo(value bool) *ActionAuthTokenDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionAuthTokenDeleteMetaGlobalInput) SetIncludes(value string) *ActionAuthTokenDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionAuthTokenDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionAuthTokenDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionAuthTokenDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionAuthTokenDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionAuthTokenDeleteRequest is a type for the entire action request
type ActionAuthTokenDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionAuthTokenDeleteResponse struct {
	Action *ActionAuthTokenDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionAuthTokenDelete) Prepare() *ActionAuthTokenDeleteInvocation {
	return &ActionAuthTokenDeleteInvocation{
		Action: action,
		Path: "/v5.0/auth_tokens/:auth_token_id",
	}
}

// ActionAuthTokenDeleteInvocation is used to configure action for invocation
type ActionAuthTokenDeleteInvocation struct {
	// Pointer to the action
	Action *ActionAuthTokenDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionAuthTokenDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionAuthTokenDeleteInvocation) SetPathParamInt(param string, value int64) *ActionAuthTokenDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionAuthTokenDeleteInvocation) SetPathParamString(param string, value string) *ActionAuthTokenDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionAuthTokenDeleteInvocation) NewMetaInput() *ActionAuthTokenDeleteMetaGlobalInput {
	inv.MetaInput = &ActionAuthTokenDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionAuthTokenDeleteInvocation) SetMetaInput(input *ActionAuthTokenDeleteMetaGlobalInput) *ActionAuthTokenDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionAuthTokenDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionAuthTokenDeleteInvocation) Call() (*ActionAuthTokenDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionAuthTokenDeleteInvocation) callAsBody() (*ActionAuthTokenDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionAuthTokenDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionAuthTokenDeleteInvocation) makeAllInputParams() *ActionAuthTokenDeleteRequest {
	return &ActionAuthTokenDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionAuthTokenDeleteInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
