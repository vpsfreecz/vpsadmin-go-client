package client

import (
	"strings"
)

// ActionVpsOutageWindowUpdate is a type for action Vps.Outage_window#Update
type ActionVpsOutageWindowUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsOutageWindowUpdate(client *Client) *ActionVpsOutageWindowUpdate {
	return &ActionVpsOutageWindowUpdate{
		Client: client,
	}
}

// ActionVpsOutageWindowUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsOutageWindowUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateMetaGlobalInput) SetNo(value bool) *ActionVpsOutageWindowUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateMetaGlobalInput) SetIncludes(value string) *ActionVpsOutageWindowUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageWindowUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageWindowUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsOutageWindowUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageWindowUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsOutageWindowUpdateInput is a type for action input parameters
type ActionVpsOutageWindowUpdateInput struct {
	IsOpen bool `json:"is_open"`
	OpensAt int64 `json:"opens_at"`
	ClosesAt int64 `json:"closes_at"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIsOpen sets parameter IsOpen to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateInput) SetIsOpen(value bool) *ActionVpsOutageWindowUpdateInput {
	in.IsOpen = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IsOpen"] = nil
	return in
}
// SetOpensAt sets parameter OpensAt to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateInput) SetOpensAt(value int64) *ActionVpsOutageWindowUpdateInput {
	in.OpensAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OpensAt"] = nil
	return in
}
// SetClosesAt sets parameter ClosesAt to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateInput) SetClosesAt(value int64) *ActionVpsOutageWindowUpdateInput {
	in.ClosesAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClosesAt"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageWindowUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageWindowUpdateInput) SelectParameters(params ...string) *ActionVpsOutageWindowUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageWindowUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsOutageWindowUpdateRequest is a type for the entire action request
type ActionVpsOutageWindowUpdateRequest struct {
	OutageWindow map[string]interface{} `json:"outage_window"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsOutageWindowUpdateOutput is a type for action output parameters
type ActionVpsOutageWindowUpdateOutput struct {
	Weekday int64 `json:"weekday"`
	IsOpen bool `json:"is_open"`
	OpensAt int64 `json:"opens_at"`
	ClosesAt int64 `json:"closes_at"`
}


// Type for action response, including envelope
type ActionVpsOutageWindowUpdateResponse struct {
	Action *ActionVpsOutageWindowUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageWindow *ActionVpsOutageWindowUpdateOutput `json:"outage_window"`
	}

	// Action output without the namespace
	Output *ActionVpsOutageWindowUpdateOutput
}


// Prepare the action for invocation
func (action *ActionVpsOutageWindowUpdate) Prepare() *ActionVpsOutageWindowUpdateInvocation {
	return &ActionVpsOutageWindowUpdateInvocation{
		Action: action,
		Path: "/v5.0/vpses/:vps_id/outage_windows/:outage_window_id",
	}
}

// ActionVpsOutageWindowUpdateInvocation is used to configure action for invocation
type ActionVpsOutageWindowUpdateInvocation struct {
	// Pointer to the action
	Action *ActionVpsOutageWindowUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsOutageWindowUpdateInput
	// Global meta input parameters
	MetaInput *ActionVpsOutageWindowUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsOutageWindowUpdateInvocation) SetPathParamInt(param string, value int64) *ActionVpsOutageWindowUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsOutageWindowUpdateInvocation) SetPathParamString(param string, value string) *ActionVpsOutageWindowUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsOutageWindowUpdateInvocation) NewInput() *ActionVpsOutageWindowUpdateInput {
	inv.Input = &ActionVpsOutageWindowUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsOutageWindowUpdateInvocation) SetInput(input *ActionVpsOutageWindowUpdateInput) *ActionVpsOutageWindowUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsOutageWindowUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsOutageWindowUpdateInvocation) NewMetaInput() *ActionVpsOutageWindowUpdateMetaGlobalInput {
	inv.MetaInput = &ActionVpsOutageWindowUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsOutageWindowUpdateInvocation) SetMetaInput(input *ActionVpsOutageWindowUpdateMetaGlobalInput) *ActionVpsOutageWindowUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsOutageWindowUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsOutageWindowUpdateInvocation) Call() (*ActionVpsOutageWindowUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsOutageWindowUpdateInvocation) callAsBody() (*ActionVpsOutageWindowUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsOutageWindowUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageWindow
	}
	return resp, err
}




func (inv *ActionVpsOutageWindowUpdateInvocation) makeAllInputParams() *ActionVpsOutageWindowUpdateRequest {
	return &ActionVpsOutageWindowUpdateRequest{
		OutageWindow: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsOutageWindowUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("IsOpen") {
			ret["is_open"] = inv.Input.IsOpen
		}
		if inv.IsParameterSelected("OpensAt") {
			ret["opens_at"] = inv.Input.OpensAt
		}
		if inv.IsParameterSelected("ClosesAt") {
			ret["closes_at"] = inv.Input.ClosesAt
		}
	}

	return ret
}

func (inv *ActionVpsOutageWindowUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
