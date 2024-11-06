package client

import (
	"strings"
)

// ActionLocationSetMaintenance is a type for action Location#Set_maintenance
type ActionLocationSetMaintenance struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationSetMaintenance(client *Client) *ActionLocationSetMaintenance {
	return &ActionLocationSetMaintenance{
		Client: client,
	}
}

// ActionLocationSetMaintenanceMetaGlobalInput is a type for action global meta input parameters
type ActionLocationSetMaintenanceMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationSetMaintenanceMetaGlobalInput) SetNo(value bool) *ActionLocationSetMaintenanceMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationSetMaintenanceMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationSetMaintenanceMetaGlobalInput) SelectParameters(params ...string) *ActionLocationSetMaintenanceMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationSetMaintenanceMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationSetMaintenanceInput is a type for action input parameters
type ActionLocationSetMaintenanceInput struct {
	Lock   bool   `json:"lock"`
	Reason string `json:"reason"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLock sets parameter Lock to value and selects it for sending
func (in *ActionLocationSetMaintenanceInput) SetLock(value bool) *ActionLocationSetMaintenanceInput {
	in.Lock = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lock"] = nil
	return in
}

// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionLocationSetMaintenanceInput) SetReason(value string) *ActionLocationSetMaintenanceInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationSetMaintenanceInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationSetMaintenanceInput) SelectParameters(params ...string) *ActionLocationSetMaintenanceInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionLocationSetMaintenanceInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionLocationSetMaintenanceInput) UnselectParameters(params ...string) *ActionLocationSetMaintenanceInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionLocationSetMaintenanceInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationSetMaintenanceRequest is a type for the entire action request
type ActionLocationSetMaintenanceRequest struct {
	Location map[string]interface{} `json:"location"`
	Meta     map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionLocationSetMaintenanceResponse struct {
	Action *ActionLocationSetMaintenance `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionLocationSetMaintenance) Prepare() *ActionLocationSetMaintenanceInvocation {
	return &ActionLocationSetMaintenanceInvocation{
		Action: action,
		Path:   "/v7.0/locations/{location_id}/set_maintenance",
	}
}

// ActionLocationSetMaintenanceInvocation is used to configure action for invocation
type ActionLocationSetMaintenanceInvocation struct {
	// Pointer to the action
	Action *ActionLocationSetMaintenance

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionLocationSetMaintenanceInput
	// Global meta input parameters
	MetaInput *ActionLocationSetMaintenanceMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionLocationSetMaintenanceInvocation) SetPathParamInt(param string, value int64) *ActionLocationSetMaintenanceInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionLocationSetMaintenanceInvocation) SetPathParamString(param string, value string) *ActionLocationSetMaintenanceInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionLocationSetMaintenanceInvocation) NewInput() *ActionLocationSetMaintenanceInput {
	inv.Input = &ActionLocationSetMaintenanceInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionLocationSetMaintenanceInvocation) SetInput(input *ActionLocationSetMaintenanceInput) *ActionLocationSetMaintenanceInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionLocationSetMaintenanceInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionLocationSetMaintenanceInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLocationSetMaintenanceInvocation) NewMetaInput() *ActionLocationSetMaintenanceMetaGlobalInput {
	inv.MetaInput = &ActionLocationSetMaintenanceMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationSetMaintenanceInvocation) SetMetaInput(input *ActionLocationSetMaintenanceMetaGlobalInput) *ActionLocationSetMaintenanceInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationSetMaintenanceInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionLocationSetMaintenanceInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationSetMaintenanceInvocation) Call() (*ActionLocationSetMaintenanceResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionLocationSetMaintenanceInvocation) callAsBody() (*ActionLocationSetMaintenanceResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionLocationSetMaintenanceResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionLocationSetMaintenanceInvocation) makeAllInputParams() *ActionLocationSetMaintenanceRequest {
	return &ActionLocationSetMaintenanceRequest{
		Location: inv.makeInputParams(),
		Meta:     inv.makeMetaInputParams(),
	}
}

func (inv *ActionLocationSetMaintenanceInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionLocationSetMaintenanceInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
