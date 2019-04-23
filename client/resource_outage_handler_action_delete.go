package client

import (
	"strings"
)

// ActionOutageHandlerDelete is a type for action Outage.Handler#Delete
type ActionOutageHandlerDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageHandlerDelete(client *Client) *ActionOutageHandlerDelete {
	return &ActionOutageHandlerDelete{
		Client: client,
	}
}

// ActionOutageHandlerDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionOutageHandlerDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageHandlerDeleteMetaGlobalInput) SetNo(value bool) *ActionOutageHandlerDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageHandlerDeleteMetaGlobalInput) SetIncludes(value string) *ActionOutageHandlerDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageHandlerDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageHandlerDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionOutageHandlerDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageHandlerDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionOutageHandlerDeleteRequest is a type for the entire action request
type ActionOutageHandlerDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionOutageHandlerDeleteResponse struct {
	Action *ActionOutageHandlerDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionOutageHandlerDelete) Prepare() *ActionOutageHandlerDeleteInvocation {
	return &ActionOutageHandlerDeleteInvocation{
		Action: action,
		Path: "/v5.0/outages/:outage_id/handlers/:handler_id",
	}
}

// ActionOutageHandlerDeleteInvocation is used to configure action for invocation
type ActionOutageHandlerDeleteInvocation struct {
	// Pointer to the action
	Action *ActionOutageHandlerDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOutageHandlerDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageHandlerDeleteInvocation) SetPathParamInt(param string, value int64) *ActionOutageHandlerDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageHandlerDeleteInvocation) SetPathParamString(param string, value string) *ActionOutageHandlerDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageHandlerDeleteInvocation) SetMetaInput(input *ActionOutageHandlerDeleteMetaGlobalInput) *ActionOutageHandlerDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageHandlerDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageHandlerDeleteInvocation) Call() (*ActionOutageHandlerDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionOutageHandlerDeleteInvocation) callAsBody() (*ActionOutageHandlerDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageHandlerDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionOutageHandlerDeleteInvocation) makeAllInputParams() *ActionOutageHandlerDeleteRequest {
	return &ActionOutageHandlerDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionOutageHandlerDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
