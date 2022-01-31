package client

import (
	"strings"
)

// ActionVpsSetMaintenance is a type for action Vps#Set_maintenance
type ActionVpsSetMaintenance struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsSetMaintenance(client *Client) *ActionVpsSetMaintenance {
	return &ActionVpsSetMaintenance{
		Client: client,
	}
}

// ActionVpsSetMaintenanceMetaGlobalInput is a type for action global meta input parameters
type ActionVpsSetMaintenanceMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsSetMaintenanceMetaGlobalInput) SetNo(value bool) *ActionVpsSetMaintenanceMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsSetMaintenanceMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsSetMaintenanceMetaGlobalInput) SelectParameters(params ...string) *ActionVpsSetMaintenanceMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsSetMaintenanceMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsSetMaintenanceInput is a type for action input parameters
type ActionVpsSetMaintenanceInput struct {
	Lock   bool   `json:"lock"`
	Reason string `json:"reason"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLock sets parameter Lock to value and selects it for sending
func (in *ActionVpsSetMaintenanceInput) SetLock(value bool) *ActionVpsSetMaintenanceInput {
	in.Lock = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lock"] = nil
	return in
}

// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionVpsSetMaintenanceInput) SetReason(value string) *ActionVpsSetMaintenanceInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsSetMaintenanceInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsSetMaintenanceInput) SelectParameters(params ...string) *ActionVpsSetMaintenanceInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsSetMaintenanceInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsSetMaintenanceInput) UnselectParameters(params ...string) *ActionVpsSetMaintenanceInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsSetMaintenanceInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsSetMaintenanceRequest is a type for the entire action request
type ActionVpsSetMaintenanceRequest struct {
	Vps  map[string]interface{} `json:"vps"`
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionVpsSetMaintenanceResponse struct {
	Action *ActionVpsSetMaintenance `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionVpsSetMaintenance) Prepare() *ActionVpsSetMaintenanceInvocation {
	return &ActionVpsSetMaintenanceInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/set_maintenance",
	}
}

// ActionVpsSetMaintenanceInvocation is used to configure action for invocation
type ActionVpsSetMaintenanceInvocation struct {
	// Pointer to the action
	Action *ActionVpsSetMaintenance

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsSetMaintenanceInput
	// Global meta input parameters
	MetaInput *ActionVpsSetMaintenanceMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsSetMaintenanceInvocation) SetPathParamInt(param string, value int64) *ActionVpsSetMaintenanceInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsSetMaintenanceInvocation) SetPathParamString(param string, value string) *ActionVpsSetMaintenanceInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsSetMaintenanceInvocation) NewInput() *ActionVpsSetMaintenanceInput {
	inv.Input = &ActionVpsSetMaintenanceInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsSetMaintenanceInvocation) SetInput(input *ActionVpsSetMaintenanceInput) *ActionVpsSetMaintenanceInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsSetMaintenanceInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsSetMaintenanceInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsSetMaintenanceInvocation) NewMetaInput() *ActionVpsSetMaintenanceMetaGlobalInput {
	inv.MetaInput = &ActionVpsSetMaintenanceMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsSetMaintenanceInvocation) SetMetaInput(input *ActionVpsSetMaintenanceMetaGlobalInput) *ActionVpsSetMaintenanceInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsSetMaintenanceInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsSetMaintenanceInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsSetMaintenanceInvocation) Call() (*ActionVpsSetMaintenanceResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsSetMaintenanceInvocation) callAsBody() (*ActionVpsSetMaintenanceResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsSetMaintenanceResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionVpsSetMaintenanceInvocation) makeAllInputParams() *ActionVpsSetMaintenanceRequest {
	return &ActionVpsSetMaintenanceRequest{
		Vps:  inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsSetMaintenanceInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Lock") {
			ret["lock"] = inv.Input.Lock
		}
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
	}

	return ret
}

func (inv *ActionVpsSetMaintenanceInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
