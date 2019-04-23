package client

import (
	"strings"
)

// ActionPoolSetMaintenance is a type for action Pool#Set_maintenance
type ActionPoolSetMaintenance struct {
	// Pointer to client
	Client *Client
}

func NewActionPoolSetMaintenance(client *Client) *ActionPoolSetMaintenance {
	return &ActionPoolSetMaintenance{
		Client: client,
	}
}

// ActionPoolSetMaintenanceMetaGlobalInput is a type for action global meta input parameters
type ActionPoolSetMaintenanceMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionPoolSetMaintenanceMetaGlobalInput) SetNo(value bool) *ActionPoolSetMaintenanceMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionPoolSetMaintenanceMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPoolSetMaintenanceMetaGlobalInput) SelectParameters(params ...string) *ActionPoolSetMaintenanceMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPoolSetMaintenanceMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPoolSetMaintenanceInput is a type for action input parameters
type ActionPoolSetMaintenanceInput struct {
	Lock bool `json:"lock"`
	Reason string `json:"reason"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLock sets parameter Lock to value and selects it for sending
func (in *ActionPoolSetMaintenanceInput) SetLock(value bool) *ActionPoolSetMaintenanceInput {
	in.Lock = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lock"] = nil
	return in
}
// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionPoolSetMaintenanceInput) SetReason(value string) *ActionPoolSetMaintenanceInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SelectParameters sets parameters from ActionPoolSetMaintenanceInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPoolSetMaintenanceInput) SelectParameters(params ...string) *ActionPoolSetMaintenanceInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPoolSetMaintenanceInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPoolSetMaintenanceRequest is a type for the entire action request
type ActionPoolSetMaintenanceRequest struct {
	Pool map[string]interface{} `json:"pool"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionPoolSetMaintenanceResponse struct {
	Action *ActionPoolSetMaintenance `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionPoolSetMaintenance) Prepare() *ActionPoolSetMaintenanceInvocation {
	return &ActionPoolSetMaintenanceInvocation{
		Action: action,
		Path: "/v5.0/pools/:pool_id/set_maintenance",
	}
}

// ActionPoolSetMaintenanceInvocation is used to configure action for invocation
type ActionPoolSetMaintenanceInvocation struct {
	// Pointer to the action
	Action *ActionPoolSetMaintenance

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionPoolSetMaintenanceInput
	// Global meta input parameters
	MetaInput *ActionPoolSetMaintenanceMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionPoolSetMaintenanceInvocation) SetPathParamInt(param string, value int64) *ActionPoolSetMaintenanceInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionPoolSetMaintenanceInvocation) SetPathParamString(param string, value string) *ActionPoolSetMaintenanceInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionPoolSetMaintenanceInvocation) SetInput(input *ActionPoolSetMaintenanceInput) *ActionPoolSetMaintenanceInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionPoolSetMaintenanceInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionPoolSetMaintenanceInvocation) SetMetaInput(input *ActionPoolSetMaintenanceMetaGlobalInput) *ActionPoolSetMaintenanceInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionPoolSetMaintenanceInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionPoolSetMaintenanceInvocation) Call() (*ActionPoolSetMaintenanceResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionPoolSetMaintenanceInvocation) callAsBody() (*ActionPoolSetMaintenanceResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionPoolSetMaintenanceResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionPoolSetMaintenanceInvocation) makeAllInputParams() *ActionPoolSetMaintenanceRequest {
	return &ActionPoolSetMaintenanceRequest{
		Pool: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionPoolSetMaintenanceInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionPoolSetMaintenanceInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
