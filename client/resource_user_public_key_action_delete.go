package client

import (
	"strings"
)

// ActionUserPublicKeyDelete is a type for action User.Public_key#Delete
type ActionUserPublicKeyDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionUserPublicKeyDelete(client *Client) *ActionUserPublicKeyDelete {
	return &ActionUserPublicKeyDelete{
		Client: client,
	}
}

// ActionUserPublicKeyDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionUserPublicKeyDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserPublicKeyDeleteMetaGlobalInput) SetNo(value bool) *ActionUserPublicKeyDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserPublicKeyDeleteMetaGlobalInput) SetIncludes(value string) *ActionUserPublicKeyDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPublicKeyDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPublicKeyDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionUserPublicKeyDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPublicKeyDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserPublicKeyDeleteRequest is a type for the entire action request
type ActionUserPublicKeyDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionUserPublicKeyDeleteResponse struct {
	Action *ActionUserPublicKeyDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionUserPublicKeyDelete) Prepare() *ActionUserPublicKeyDeleteInvocation {
	return &ActionUserPublicKeyDeleteInvocation{
		Action: action,
		Path: "/v5.0/users/{user_id}/public_keys/{public_key_id}",
	}
}

// ActionUserPublicKeyDeleteInvocation is used to configure action for invocation
type ActionUserPublicKeyDeleteInvocation struct {
	// Pointer to the action
	Action *ActionUserPublicKeyDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserPublicKeyDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserPublicKeyDeleteInvocation) SetPathParamInt(param string, value int64) *ActionUserPublicKeyDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserPublicKeyDeleteInvocation) SetPathParamString(param string, value string) *ActionUserPublicKeyDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserPublicKeyDeleteInvocation) NewMetaInput() *ActionUserPublicKeyDeleteMetaGlobalInput {
	inv.MetaInput = &ActionUserPublicKeyDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserPublicKeyDeleteInvocation) SetMetaInput(input *ActionUserPublicKeyDeleteMetaGlobalInput) *ActionUserPublicKeyDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserPublicKeyDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserPublicKeyDeleteInvocation) Call() (*ActionUserPublicKeyDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserPublicKeyDeleteInvocation) callAsBody() (*ActionUserPublicKeyDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserPublicKeyDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionUserPublicKeyDeleteInvocation) makeAllInputParams() *ActionUserPublicKeyDeleteRequest {
	return &ActionUserPublicKeyDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionUserPublicKeyDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
