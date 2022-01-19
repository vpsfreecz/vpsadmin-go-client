package client

import (
	"strings"
)

// ActionVpsOutageWindowUpdateAll is a type for action Vps.Outage_window#Update_all
type ActionVpsOutageWindowUpdateAll struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsOutageWindowUpdateAll(client *Client) *ActionVpsOutageWindowUpdateAll {
	return &ActionVpsOutageWindowUpdateAll{
		Client: client,
	}
}

// ActionVpsOutageWindowUpdateAllMetaGlobalInput is a type for action global meta input parameters
type ActionVpsOutageWindowUpdateAllMetaGlobalInput struct {
	No       bool   `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateAllMetaGlobalInput) SetNo(value bool) *ActionVpsOutageWindowUpdateAllMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateAllMetaGlobalInput) SetIncludes(value string) *ActionVpsOutageWindowUpdateAllMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageWindowUpdateAllMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageWindowUpdateAllMetaGlobalInput) SelectParameters(params ...string) *ActionVpsOutageWindowUpdateAllMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageWindowUpdateAllMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsOutageWindowUpdateAllInput is a type for action input parameters
type ActionVpsOutageWindowUpdateAllInput struct {
	IsOpen   bool  `json:"is_open"`
	OpensAt  int64 `json:"opens_at"`
	ClosesAt int64 `json:"closes_at"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIsOpen sets parameter IsOpen to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateAllInput) SetIsOpen(value bool) *ActionVpsOutageWindowUpdateAllInput {
	in.IsOpen = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IsOpen"] = nil
	return in
}

// SetOpensAt sets parameter OpensAt to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateAllInput) SetOpensAt(value int64) *ActionVpsOutageWindowUpdateAllInput {
	in.OpensAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OpensAt"] = nil
	return in
}

// SetClosesAt sets parameter ClosesAt to value and selects it for sending
func (in *ActionVpsOutageWindowUpdateAllInput) SetClosesAt(value int64) *ActionVpsOutageWindowUpdateAllInput {
	in.ClosesAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClosesAt"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageWindowUpdateAllInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageWindowUpdateAllInput) SelectParameters(params ...string) *ActionVpsOutageWindowUpdateAllInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageWindowUpdateAllInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsOutageWindowUpdateAllRequest is a type for the entire action request
type ActionVpsOutageWindowUpdateAllRequest struct {
	OutageWindow map[string]interface{} `json:"outage_window"`
	Meta         map[string]interface{} `json:"_meta"`
}

// ActionVpsOutageWindowUpdateAllOutput is a type for action output parameters
type ActionVpsOutageWindowUpdateAllOutput struct {
	Weekday  int64 `json:"weekday"`
	IsOpen   bool  `json:"is_open"`
	OpensAt  int64 `json:"opens_at"`
	ClosesAt int64 `json:"closes_at"`
}

// Type for action response, including envelope
type ActionVpsOutageWindowUpdateAllResponse struct {
	Action *ActionVpsOutageWindowUpdateAll `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageWindows []*ActionVpsOutageWindowUpdateAllOutput `json:"outage_windows"`
	}

	// Action output without the namespace
	Output []*ActionVpsOutageWindowUpdateAllOutput
}

// Prepare the action for invocation
func (action *ActionVpsOutageWindowUpdateAll) Prepare() *ActionVpsOutageWindowUpdateAllInvocation {
	return &ActionVpsOutageWindowUpdateAllInvocation{
		Action: action,
		Path:   "/v5.0/vpses/{vps_id}/outage_windows",
	}
}

// ActionVpsOutageWindowUpdateAllInvocation is used to configure action for invocation
type ActionVpsOutageWindowUpdateAllInvocation struct {
	// Pointer to the action
	Action *ActionVpsOutageWindowUpdateAll

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsOutageWindowUpdateAllInput
	// Global meta input parameters
	MetaInput *ActionVpsOutageWindowUpdateAllMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsOutageWindowUpdateAllInvocation) SetPathParamInt(param string, value int64) *ActionVpsOutageWindowUpdateAllInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsOutageWindowUpdateAllInvocation) SetPathParamString(param string, value string) *ActionVpsOutageWindowUpdateAllInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsOutageWindowUpdateAllInvocation) NewInput() *ActionVpsOutageWindowUpdateAllInput {
	inv.Input = &ActionVpsOutageWindowUpdateAllInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsOutageWindowUpdateAllInvocation) SetInput(input *ActionVpsOutageWindowUpdateAllInput) *ActionVpsOutageWindowUpdateAllInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsOutageWindowUpdateAllInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsOutageWindowUpdateAllInvocation) NewMetaInput() *ActionVpsOutageWindowUpdateAllMetaGlobalInput {
	inv.MetaInput = &ActionVpsOutageWindowUpdateAllMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsOutageWindowUpdateAllInvocation) SetMetaInput(input *ActionVpsOutageWindowUpdateAllMetaGlobalInput) *ActionVpsOutageWindowUpdateAllInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsOutageWindowUpdateAllInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsOutageWindowUpdateAllInvocation) Call() (*ActionVpsOutageWindowUpdateAllResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsOutageWindowUpdateAllInvocation) callAsBody() (*ActionVpsOutageWindowUpdateAllResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsOutageWindowUpdateAllResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageWindows
	}
	return resp, err
}

func (inv *ActionVpsOutageWindowUpdateAllInvocation) makeAllInputParams() *ActionVpsOutageWindowUpdateAllRequest {
	return &ActionVpsOutageWindowUpdateAllRequest{
		OutageWindow: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsOutageWindowUpdateAllInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionVpsOutageWindowUpdateAllInvocation) makeMetaInputParams() map[string]interface{} {
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
