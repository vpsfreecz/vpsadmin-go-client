package client

import (
	"strings"
)

// ActionUserTouch is a type for action User#Touch
type ActionUserTouch struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTouch(client *Client) *ActionUserTouch {
	return &ActionUserTouch{
		Client: client,
	}
}

// ActionUserTouchMetaGlobalInput is a type for action global meta input parameters
type ActionUserTouchMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTouchMetaGlobalInput) SetNo(value bool) *ActionUserTouchMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserTouchMetaGlobalInput) SetIncludes(value string) *ActionUserTouchMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTouchMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTouchMetaGlobalInput) SelectParameters(params ...string) *ActionUserTouchMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTouchMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}





// Type for action response, including envelope
type ActionUserTouchResponse struct {
	Action *ActionUserTouch `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionUserTouch) Prepare() *ActionUserTouchInvocation {
	return &ActionUserTouchInvocation{
		Action: action,
		Path: "/v5.0/users/{user_id}/touch",
	}
}

// ActionUserTouchInvocation is used to configure action for invocation
type ActionUserTouchInvocation struct {
	// Pointer to the action
	Action *ActionUserTouch

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserTouchMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTouchInvocation) SetPathParamInt(param string, value int64) *ActionUserTouchInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTouchInvocation) SetPathParamString(param string, value string) *ActionUserTouchInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTouchInvocation) NewMetaInput() *ActionUserTouchMetaGlobalInput {
	inv.MetaInput = &ActionUserTouchMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTouchInvocation) SetMetaInput(input *ActionUserTouchMetaGlobalInput) *ActionUserTouchInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTouchInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserTouchInvocation) Call() (*ActionUserTouchResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserTouchInvocation) callAsQuery() (*ActionUserTouchResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserTouchResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	return resp, err
}




func (inv *ActionUserTouchInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

