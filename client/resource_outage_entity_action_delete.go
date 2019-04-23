package client

import (
	"strings"
)

// ActionOutageEntityDelete is a type for action Outage.Entity#Delete
type ActionOutageEntityDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageEntityDelete(client *Client) *ActionOutageEntityDelete {
	return &ActionOutageEntityDelete{
		Client: client,
	}
}

// ActionOutageEntityDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionOutageEntityDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageEntityDeleteMetaGlobalInput) SetNo(value bool) *ActionOutageEntityDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageEntityDeleteMetaGlobalInput) SetIncludes(value string) *ActionOutageEntityDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageEntityDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageEntityDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionOutageEntityDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageEntityDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionOutageEntityDeleteRequest is a type for the entire action request
type ActionOutageEntityDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionOutageEntityDeleteResponse struct {
	Action *ActionOutageEntityDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionOutageEntityDelete) Prepare() *ActionOutageEntityDeleteInvocation {
	return &ActionOutageEntityDeleteInvocation{
		Action: action,
		Path: "/v5.0/outages/:outage_id/entities/:entity_id",
	}
}

// ActionOutageEntityDeleteInvocation is used to configure action for invocation
type ActionOutageEntityDeleteInvocation struct {
	// Pointer to the action
	Action *ActionOutageEntityDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOutageEntityDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageEntityDeleteInvocation) SetPathParamInt(param string, value int64) *ActionOutageEntityDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageEntityDeleteInvocation) SetPathParamString(param string, value string) *ActionOutageEntityDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageEntityDeleteInvocation) SetMetaInput(input *ActionOutageEntityDeleteMetaGlobalInput) *ActionOutageEntityDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageEntityDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageEntityDeleteInvocation) Call() (*ActionOutageEntityDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionOutageEntityDeleteInvocation) callAsBody() (*ActionOutageEntityDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageEntityDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionOutageEntityDeleteInvocation) makeAllInputParams() *ActionOutageEntityDeleteRequest {
	return &ActionOutageEntityDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionOutageEntityDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
