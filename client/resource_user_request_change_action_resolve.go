package client

import (
	"strings"
)

// ActionUserRequestChangeResolve is a type for action User_request.Change#Resolve
type ActionUserRequestChangeResolve struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestChangeResolve(client *Client) *ActionUserRequestChangeResolve {
	return &ActionUserRequestChangeResolve{
		Client: client,
	}
}

// ActionUserRequestChangeResolveMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestChangeResolveMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestChangeResolveMetaGlobalInput) SetNo(value bool) *ActionUserRequestChangeResolveMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestChangeResolveMetaGlobalInput) SetIncludes(value string) *ActionUserRequestChangeResolveMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestChangeResolveMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestChangeResolveMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestChangeResolveMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestChangeResolveMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestChangeResolveInput is a type for action input parameters
type ActionUserRequestChangeResolveInput struct {
	Action string `json:"action"`
	Reason string `json:"reason"`
	ChangeReason string `json:"change_reason"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetAction sets parameter Action to value and selects it for sending
func (in *ActionUserRequestChangeResolveInput) SetAction(value string) *ActionUserRequestChangeResolveInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}
// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionUserRequestChangeResolveInput) SetReason(value string) *ActionUserRequestChangeResolveInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}
// SetChangeReason sets parameter ChangeReason to value and selects it for sending
func (in *ActionUserRequestChangeResolveInput) SetChangeReason(value string) *ActionUserRequestChangeResolveInput {
	in.ChangeReason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChangeReason"] = nil
	return in
}
// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserRequestChangeResolveInput) SetFullName(value string) *ActionUserRequestChangeResolveInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
	return in
}
// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionUserRequestChangeResolveInput) SetEmail(value string) *ActionUserRequestChangeResolveInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}
// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserRequestChangeResolveInput) SetAddress(value string) *ActionUserRequestChangeResolveInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestChangeResolveInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestChangeResolveInput) SelectParameters(params ...string) *ActionUserRequestChangeResolveInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestChangeResolveInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestChangeResolveRequest is a type for the entire action request
type ActionUserRequestChangeResolveRequest struct {
	Change map[string]interface{} `json:"change"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionUserRequestChangeResolveResponse struct {
	Action *ActionUserRequestChangeResolve `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionUserRequestChangeResolve) Prepare() *ActionUserRequestChangeResolveInvocation {
	return &ActionUserRequestChangeResolveInvocation{
		Action: action,
		Path: "/v5.0/user_request/changes/{change_id}/resolve",
	}
}

// ActionUserRequestChangeResolveInvocation is used to configure action for invocation
type ActionUserRequestChangeResolveInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestChangeResolve

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserRequestChangeResolveInput
	// Global meta input parameters
	MetaInput *ActionUserRequestChangeResolveMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserRequestChangeResolveInvocation) SetPathParamInt(param string, value int64) *ActionUserRequestChangeResolveInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserRequestChangeResolveInvocation) SetPathParamString(param string, value string) *ActionUserRequestChangeResolveInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserRequestChangeResolveInvocation) NewInput() *ActionUserRequestChangeResolveInput {
	inv.Input = &ActionUserRequestChangeResolveInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserRequestChangeResolveInvocation) SetInput(input *ActionUserRequestChangeResolveInput) *ActionUserRequestChangeResolveInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserRequestChangeResolveInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestChangeResolveInvocation) NewMetaInput() *ActionUserRequestChangeResolveMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestChangeResolveMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestChangeResolveInvocation) SetMetaInput(input *ActionUserRequestChangeResolveMetaGlobalInput) *ActionUserRequestChangeResolveInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestChangeResolveInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestChangeResolveInvocation) Call() (*ActionUserRequestChangeResolveResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserRequestChangeResolveInvocation) callAsBody() (*ActionUserRequestChangeResolveResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserRequestChangeResolveResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionUserRequestChangeResolveInvocation) makeAllInputParams() *ActionUserRequestChangeResolveRequest {
	return &ActionUserRequestChangeResolveRequest{
		Change: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserRequestChangeResolveInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Action") {
			ret["action"] = inv.Input.Action
		}
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
		if inv.IsParameterSelected("ChangeReason") {
			ret["change_reason"] = inv.Input.ChangeReason
		}
		if inv.IsParameterSelected("FullName") {
			ret["full_name"] = inv.Input.FullName
		}
		if inv.IsParameterSelected("Email") {
			ret["email"] = inv.Input.Email
		}
		if inv.IsParameterSelected("Address") {
			ret["address"] = inv.Input.Address
		}
	}

	return ret
}

func (inv *ActionUserRequestChangeResolveInvocation) makeMetaInputParams() map[string]interface{} {
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
