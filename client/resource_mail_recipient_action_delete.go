package client

import (
	"strings"
)

// ActionMailRecipientDelete is a type for action Mail_recipient#Delete
type ActionMailRecipientDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionMailRecipientDelete(client *Client) *ActionMailRecipientDelete {
	return &ActionMailRecipientDelete{
		Client: client,
	}
}

// ActionMailRecipientDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionMailRecipientDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailRecipientDeleteMetaGlobalInput) SetNo(value bool) *ActionMailRecipientDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailRecipientDeleteMetaGlobalInput) SetIncludes(value string) *ActionMailRecipientDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailRecipientDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailRecipientDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionMailRecipientDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailRecipientDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionMailRecipientDeleteRequest is a type for the entire action request
type ActionMailRecipientDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionMailRecipientDeleteResponse struct {
	Action *ActionMailRecipientDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionMailRecipientDelete) Prepare() *ActionMailRecipientDeleteInvocation {
	return &ActionMailRecipientDeleteInvocation{
		Action: action,
		Path: "/v6.0/mail_recipients/{mail_recipient_id}",
	}
}

// ActionMailRecipientDeleteInvocation is used to configure action for invocation
type ActionMailRecipientDeleteInvocation struct {
	// Pointer to the action
	Action *ActionMailRecipientDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailRecipientDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailRecipientDeleteInvocation) SetPathParamInt(param string, value int64) *ActionMailRecipientDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailRecipientDeleteInvocation) SetPathParamString(param string, value string) *ActionMailRecipientDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailRecipientDeleteInvocation) NewMetaInput() *ActionMailRecipientDeleteMetaGlobalInput {
	inv.MetaInput = &ActionMailRecipientDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailRecipientDeleteInvocation) SetMetaInput(input *ActionMailRecipientDeleteMetaGlobalInput) *ActionMailRecipientDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailRecipientDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailRecipientDeleteInvocation) Call() (*ActionMailRecipientDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMailRecipientDeleteInvocation) callAsBody() (*ActionMailRecipientDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailRecipientDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionMailRecipientDeleteInvocation) makeAllInputParams() *ActionMailRecipientDeleteRequest {
	return &ActionMailRecipientDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionMailRecipientDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
