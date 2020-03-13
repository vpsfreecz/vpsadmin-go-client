package client

import (
	"strings"
)

// ActionVpsMaintenanceWindowUpdate is a type for action Vps.Maintenance_window#Update
type ActionVpsMaintenanceWindowUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMaintenanceWindowUpdate(client *Client) *ActionVpsMaintenanceWindowUpdate {
	return &ActionVpsMaintenanceWindowUpdate{
		Client: client,
	}
}

// ActionVpsMaintenanceWindowUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMaintenanceWindowUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateMetaGlobalInput) SetNo(value bool) *ActionVpsMaintenanceWindowUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateMetaGlobalInput) SetIncludes(value string) *ActionVpsMaintenanceWindowUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMaintenanceWindowUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMaintenanceWindowUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMaintenanceWindowUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMaintenanceWindowUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMaintenanceWindowUpdateInput is a type for action input parameters
type ActionVpsMaintenanceWindowUpdateInput struct {
	IsOpen bool `json:"is_open"`
	OpensAt int64 `json:"opens_at"`
	ClosesAt int64 `json:"closes_at"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIsOpen sets parameter IsOpen to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateInput) SetIsOpen(value bool) *ActionVpsMaintenanceWindowUpdateInput {
	in.IsOpen = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IsOpen"] = nil
	return in
}
// SetOpensAt sets parameter OpensAt to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateInput) SetOpensAt(value int64) *ActionVpsMaintenanceWindowUpdateInput {
	in.OpensAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OpensAt"] = nil
	return in
}
// SetClosesAt sets parameter ClosesAt to value and selects it for sending
func (in *ActionVpsMaintenanceWindowUpdateInput) SetClosesAt(value int64) *ActionVpsMaintenanceWindowUpdateInput {
	in.ClosesAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClosesAt"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMaintenanceWindowUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMaintenanceWindowUpdateInput) SelectParameters(params ...string) *ActionVpsMaintenanceWindowUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMaintenanceWindowUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMaintenanceWindowUpdateRequest is a type for the entire action request
type ActionVpsMaintenanceWindowUpdateRequest struct {
	MaintenanceWindow map[string]interface{} `json:"maintenance_window"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsMaintenanceWindowUpdateOutput is a type for action output parameters
type ActionVpsMaintenanceWindowUpdateOutput struct {
	Weekday int64 `json:"weekday"`
	IsOpen bool `json:"is_open"`
	OpensAt int64 `json:"opens_at"`
	ClosesAt int64 `json:"closes_at"`
}


// Type for action response, including envelope
type ActionVpsMaintenanceWindowUpdateResponse struct {
	Action *ActionVpsMaintenanceWindowUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MaintenanceWindow *ActionVpsMaintenanceWindowUpdateOutput `json:"maintenance_window"`
	}

	// Action output without the namespace
	Output *ActionVpsMaintenanceWindowUpdateOutput
}


// Prepare the action for invocation
func (action *ActionVpsMaintenanceWindowUpdate) Prepare() *ActionVpsMaintenanceWindowUpdateInvocation {
	return &ActionVpsMaintenanceWindowUpdateInvocation{
		Action: action,
		Path: "/v6.0/vpses/{vps_id}/maintenance_windows/{maintenance_window_id}",
	}
}

// ActionVpsMaintenanceWindowUpdateInvocation is used to configure action for invocation
type ActionVpsMaintenanceWindowUpdateInvocation struct {
	// Pointer to the action
	Action *ActionVpsMaintenanceWindowUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsMaintenanceWindowUpdateInput
	// Global meta input parameters
	MetaInput *ActionVpsMaintenanceWindowUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMaintenanceWindowUpdateInvocation) SetPathParamInt(param string, value int64) *ActionVpsMaintenanceWindowUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMaintenanceWindowUpdateInvocation) SetPathParamString(param string, value string) *ActionVpsMaintenanceWindowUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsMaintenanceWindowUpdateInvocation) NewInput() *ActionVpsMaintenanceWindowUpdateInput {
	inv.Input = &ActionVpsMaintenanceWindowUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsMaintenanceWindowUpdateInvocation) SetInput(input *ActionVpsMaintenanceWindowUpdateInput) *ActionVpsMaintenanceWindowUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsMaintenanceWindowUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMaintenanceWindowUpdateInvocation) NewMetaInput() *ActionVpsMaintenanceWindowUpdateMetaGlobalInput {
	inv.MetaInput = &ActionVpsMaintenanceWindowUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMaintenanceWindowUpdateInvocation) SetMetaInput(input *ActionVpsMaintenanceWindowUpdateMetaGlobalInput) *ActionVpsMaintenanceWindowUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMaintenanceWindowUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMaintenanceWindowUpdateInvocation) Call() (*ActionVpsMaintenanceWindowUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsMaintenanceWindowUpdateInvocation) callAsBody() (*ActionVpsMaintenanceWindowUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsMaintenanceWindowUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MaintenanceWindow
	}
	return resp, err
}




func (inv *ActionVpsMaintenanceWindowUpdateInvocation) makeAllInputParams() *ActionVpsMaintenanceWindowUpdateRequest {
	return &ActionVpsMaintenanceWindowUpdateRequest{
		MaintenanceWindow: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsMaintenanceWindowUpdateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionVpsMaintenanceWindowUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
