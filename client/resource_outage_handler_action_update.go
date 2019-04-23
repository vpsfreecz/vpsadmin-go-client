package client

import (
	"strings"
)

// ActionOutageHandlerUpdate is a type for action Outage.Handler#Update
type ActionOutageHandlerUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageHandlerUpdate(client *Client) *ActionOutageHandlerUpdate {
	return &ActionOutageHandlerUpdate{
		Client: client,
	}
}

// ActionOutageHandlerUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionOutageHandlerUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageHandlerUpdateMetaGlobalInput) SetNo(value bool) *ActionOutageHandlerUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageHandlerUpdateMetaGlobalInput) SetIncludes(value string) *ActionOutageHandlerUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageHandlerUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageHandlerUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionOutageHandlerUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageHandlerUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageHandlerUpdateInput is a type for action input parameters
type ActionOutageHandlerUpdateInput struct {
	Note string `json:"note"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNote sets parameter Note to value and selects it for sending
func (in *ActionOutageHandlerUpdateInput) SetNote(value string) *ActionOutageHandlerUpdateInput {
	in.Note = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Note"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageHandlerUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageHandlerUpdateInput) SelectParameters(params ...string) *ActionOutageHandlerUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageHandlerUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageHandlerUpdateRequest is a type for the entire action request
type ActionOutageHandlerUpdateRequest struct {
	Handler map[string]interface{} `json:"handler"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionOutageHandlerUpdateOutput is a type for action output parameters
type ActionOutageHandlerUpdateOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	FullName string `json:"full_name"`
	Note string `json:"note"`
}


// Type for action response, including envelope
type ActionOutageHandlerUpdateResponse struct {
	Action *ActionOutageHandlerUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Handler *ActionOutageHandlerUpdateOutput `json:"handler"`
	}

	// Action output without the namespace
	Output *ActionOutageHandlerUpdateOutput
}


// Prepare the action for invocation
func (action *ActionOutageHandlerUpdate) Prepare() *ActionOutageHandlerUpdateInvocation {
	return &ActionOutageHandlerUpdateInvocation{
		Action: action,
		Path: "/v5.0/outages/:outage_id/handlers/:handler_id",
	}
}

// ActionOutageHandlerUpdateInvocation is used to configure action for invocation
type ActionOutageHandlerUpdateInvocation struct {
	// Pointer to the action
	Action *ActionOutageHandlerUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageHandlerUpdateInput
	// Global meta input parameters
	MetaInput *ActionOutageHandlerUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageHandlerUpdateInvocation) SetPathParamInt(param string, value int64) *ActionOutageHandlerUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageHandlerUpdateInvocation) SetPathParamString(param string, value string) *ActionOutageHandlerUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageHandlerUpdateInvocation) SetInput(input *ActionOutageHandlerUpdateInput) *ActionOutageHandlerUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageHandlerUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageHandlerUpdateInvocation) SetMetaInput(input *ActionOutageHandlerUpdateMetaGlobalInput) *ActionOutageHandlerUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageHandlerUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageHandlerUpdateInvocation) Call() (*ActionOutageHandlerUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionOutageHandlerUpdateInvocation) callAsBody() (*ActionOutageHandlerUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageHandlerUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Handler
	}
	return resp, err
}




func (inv *ActionOutageHandlerUpdateInvocation) makeAllInputParams() *ActionOutageHandlerUpdateRequest {
	return &ActionOutageHandlerUpdateRequest{
		Handler: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOutageHandlerUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Note") {
			ret["note"] = inv.Input.Note
		}
	}

	return ret
}

func (inv *ActionOutageHandlerUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
