package client

import (
	"strings"
)

// ActionVpsMaintenanceWindowUpdateAll is a type for action Vps.Maintenance_window#Update_all
type ActionVpsMaintenanceWindowUpdateAll struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMaintenanceWindowUpdateAll(client *Client) *ActionVpsMaintenanceWindowUpdateAll {
	return &ActionVpsMaintenanceWindowUpdateAll{
		Client: client,
	}
}

// ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput) SetIncludes(value string) *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput) SetNo(value bool) *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMaintenanceWindowUpdateAllInput is a type for action input parameters
type ActionVpsMaintenanceWindowUpdateAllInput struct {
	ClosesAt int64 `json:"closes_at"`
	IsOpen   bool  `json:"is_open"`
	OpensAt  int64 `json:"opens_at"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetClosesAt sets parameter ClosesAt to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateAllInput) SetClosesAt(value int64) *ActionVpsMaintenanceWindowUpdateAllInput {
	in.ClosesAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClosesAt"] = nil
	return in
}

// SetIsOpen sets parameter IsOpen to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateAllInput) SetIsOpen(value bool) *ActionVpsMaintenanceWindowUpdateAllInput {
	in.IsOpen = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IsOpen"] = nil
	return in
}

// SetOpensAt sets parameter OpensAt to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateAllInput) SetOpensAt(value int64) *ActionVpsMaintenanceWindowUpdateAllInput {
	in.OpensAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OpensAt"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMaintenanceWindowUpdateAllInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMaintenanceWindowUpdateAllInput) SelectParameters(params ...string) *ActionVpsMaintenanceWindowUpdateAllInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsMaintenanceWindowUpdateAllInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsMaintenanceWindowUpdateAllInput) UnselectParameters(params ...string) *ActionVpsMaintenanceWindowUpdateAllInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsMaintenanceWindowUpdateAllInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMaintenanceWindowUpdateAllRequest is a type for the entire action request
type ActionVpsMaintenanceWindowUpdateAllRequest struct {
	MaintenanceWindow map[string]interface{} `json:"maintenance_window"`
	Meta              map[string]interface{} `json:"_meta"`
}

// ActionVpsMaintenanceWindowUpdateAllOutput is a type for action output parameters
type ActionVpsMaintenanceWindowUpdateAllOutput struct {
	ClosesAt int64 `json:"closes_at"`
	IsOpen   bool  `json:"is_open"`
	OpensAt  int64 `json:"opens_at"`
	Weekday  int64 `json:"weekday"`
}

// Type for action response, including envelope
type ActionVpsMaintenanceWindowUpdateAllResponse struct {
	Action *ActionVpsMaintenanceWindowUpdateAll `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MaintenanceWindows []*ActionVpsMaintenanceWindowUpdateAllOutput `json:"maintenance_windows"`
	}

	// Action output without the namespace
	Output []*ActionVpsMaintenanceWindowUpdateAllOutput
}

// Prepare the action for invocation
func (action *ActionVpsMaintenanceWindowUpdateAll) Prepare() *ActionVpsMaintenanceWindowUpdateAllInvocation {
	return &ActionVpsMaintenanceWindowUpdateAllInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/maintenance_windows",
	}
}

// ActionVpsMaintenanceWindowUpdateAllInvocation is used to configure action for invocation
type ActionVpsMaintenanceWindowUpdateAllInvocation struct {
	// Pointer to the action
	Action *ActionVpsMaintenanceWindowUpdateAll

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsMaintenanceWindowUpdateAllInput
	// Global meta input parameters
	MetaInput *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) SetPathParamInt(param string, value int64) *ActionVpsMaintenanceWindowUpdateAllInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) SetPathParamString(param string, value string) *ActionVpsMaintenanceWindowUpdateAllInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) NewInput() *ActionVpsMaintenanceWindowUpdateAllInput {
	inv.Input = &ActionVpsMaintenanceWindowUpdateAllInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) SetInput(input *ActionVpsMaintenanceWindowUpdateAllInput) *ActionVpsMaintenanceWindowUpdateAllInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) NewMetaInput() *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput {
	inv.MetaInput = &ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) SetMetaInput(input *ActionVpsMaintenanceWindowUpdateAllMetaGlobalInput) *ActionVpsMaintenanceWindowUpdateAllInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) Call() (*ActionVpsMaintenanceWindowUpdateAllResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) callAsBody() (*ActionVpsMaintenanceWindowUpdateAllResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsMaintenanceWindowUpdateAllResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MaintenanceWindows
	}
	return resp, err
}

func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) makeAllInputParams() *ActionVpsMaintenanceWindowUpdateAllRequest {
	return &ActionVpsMaintenanceWindowUpdateAllRequest{
		MaintenanceWindow: inv.makeInputParams(),
		Meta:              inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ClosesAt") {
			ret["closes_at"] = inv.Input.ClosesAt
		}
		if inv.IsParameterSelected("IsOpen") {
			ret["is_open"] = inv.Input.IsOpen
		}
		if inv.IsParameterSelected("OpensAt") {
			ret["opens_at"] = inv.Input.OpensAt
		}
	}

	return ret
}

func (inv *ActionVpsMaintenanceWindowUpdateAllInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
