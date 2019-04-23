package client

import (
	"strings"
)

// ActionEnvironmentSetMaintenance is a type for action Environment#Set_maintenance
type ActionEnvironmentSetMaintenance struct {
	// Pointer to client
	Client *Client
}

func NewActionEnvironmentSetMaintenance(client *Client) *ActionEnvironmentSetMaintenance {
	return &ActionEnvironmentSetMaintenance{
		Client: client,
	}
}

// ActionEnvironmentSetMaintenanceMetaGlobalInput is a type for action global meta input parameters
type ActionEnvironmentSetMaintenanceMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEnvironmentSetMaintenanceMetaGlobalInput) SetNo(value bool) *ActionEnvironmentSetMaintenanceMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentSetMaintenanceMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentSetMaintenanceMetaGlobalInput) SelectParameters(params ...string) *ActionEnvironmentSetMaintenanceMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentSetMaintenanceMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentSetMaintenanceInput is a type for action input parameters
type ActionEnvironmentSetMaintenanceInput struct {
	Lock bool `json:"lock"`
	Reason string `json:"reason"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLock sets parameter Lock to value and selects it for sending
func (in *ActionEnvironmentSetMaintenanceInput) SetLock(value bool) *ActionEnvironmentSetMaintenanceInput {
	in.Lock = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lock"] = nil
	return in
}
// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionEnvironmentSetMaintenanceInput) SetReason(value string) *ActionEnvironmentSetMaintenanceInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentSetMaintenanceInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentSetMaintenanceInput) SelectParameters(params ...string) *ActionEnvironmentSetMaintenanceInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentSetMaintenanceInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentSetMaintenanceRequest is a type for the entire action request
type ActionEnvironmentSetMaintenanceRequest struct {
	Environment map[string]interface{} `json:"environment"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionEnvironmentSetMaintenanceResponse struct {
	Action *ActionEnvironmentSetMaintenance `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionEnvironmentSetMaintenance) Prepare() *ActionEnvironmentSetMaintenanceInvocation {
	return &ActionEnvironmentSetMaintenanceInvocation{
		Action: action,
		Path: "/v5.0/environments/:environment_id/set_maintenance",
	}
}

// ActionEnvironmentSetMaintenanceInvocation is used to configure action for invocation
type ActionEnvironmentSetMaintenanceInvocation struct {
	// Pointer to the action
	Action *ActionEnvironmentSetMaintenance

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEnvironmentSetMaintenanceInput
	// Global meta input parameters
	MetaInput *ActionEnvironmentSetMaintenanceMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEnvironmentSetMaintenanceInvocation) SetPathParamInt(param string, value int64) *ActionEnvironmentSetMaintenanceInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEnvironmentSetMaintenanceInvocation) SetPathParamString(param string, value string) *ActionEnvironmentSetMaintenanceInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionEnvironmentSetMaintenanceInvocation) SetInput(input *ActionEnvironmentSetMaintenanceInput) *ActionEnvironmentSetMaintenanceInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEnvironmentSetMaintenanceInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEnvironmentSetMaintenanceInvocation) SetMetaInput(input *ActionEnvironmentSetMaintenanceMetaGlobalInput) *ActionEnvironmentSetMaintenanceInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEnvironmentSetMaintenanceInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEnvironmentSetMaintenanceInvocation) Call() (*ActionEnvironmentSetMaintenanceResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionEnvironmentSetMaintenanceInvocation) callAsBody() (*ActionEnvironmentSetMaintenanceResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEnvironmentSetMaintenanceResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionEnvironmentSetMaintenanceInvocation) makeAllInputParams() *ActionEnvironmentSetMaintenanceRequest {
	return &ActionEnvironmentSetMaintenanceRequest{
		Environment: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionEnvironmentSetMaintenanceInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionEnvironmentSetMaintenanceInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
