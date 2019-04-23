package client

import (
)

// ActionClusterSetMaintenance is a type for action Cluster#Set_maintenance
type ActionClusterSetMaintenance struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterSetMaintenance(client *Client) *ActionClusterSetMaintenance {
	return &ActionClusterSetMaintenance{
		Client: client,
	}
}

// ActionClusterSetMaintenanceMetaGlobalInput is a type for action global meta input parameters
type ActionClusterSetMaintenanceMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterSetMaintenanceMetaGlobalInput) SetNo(value bool) *ActionClusterSetMaintenanceMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterSetMaintenanceMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterSetMaintenanceMetaGlobalInput) SelectParameters(params ...string) *ActionClusterSetMaintenanceMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterSetMaintenanceMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterSetMaintenanceInput is a type for action input parameters
type ActionClusterSetMaintenanceInput struct {
	Lock bool `json:"lock"`
	Reason string `json:"reason"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLock sets parameter Lock to value and selects it for sending
func (in *ActionClusterSetMaintenanceInput) SetLock(value bool) *ActionClusterSetMaintenanceInput {
	in.Lock = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lock"] = nil
	return in
}
// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionClusterSetMaintenanceInput) SetReason(value string) *ActionClusterSetMaintenanceInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterSetMaintenanceInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterSetMaintenanceInput) SelectParameters(params ...string) *ActionClusterSetMaintenanceInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterSetMaintenanceInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterSetMaintenanceRequest is a type for the entire action request
type ActionClusterSetMaintenanceRequest struct {
	Cluster map[string]interface{} `json:"cluster"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionClusterSetMaintenanceResponse struct {
	Action *ActionClusterSetMaintenance `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionClusterSetMaintenance) Prepare() *ActionClusterSetMaintenanceInvocation {
	return &ActionClusterSetMaintenanceInvocation{
		Action: action,
		Path: "/v5.0/cluster/set_maintenance",
	}
}

// ActionClusterSetMaintenanceInvocation is used to configure action for invocation
type ActionClusterSetMaintenanceInvocation struct {
	// Pointer to the action
	Action *ActionClusterSetMaintenance

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterSetMaintenanceInput
	// Global meta input parameters
	MetaInput *ActionClusterSetMaintenanceMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionClusterSetMaintenanceInvocation) SetInput(input *ActionClusterSetMaintenanceInput) *ActionClusterSetMaintenanceInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterSetMaintenanceInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterSetMaintenanceInvocation) SetMetaInput(input *ActionClusterSetMaintenanceMetaGlobalInput) *ActionClusterSetMaintenanceInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterSetMaintenanceInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterSetMaintenanceInvocation) Call() (*ActionClusterSetMaintenanceResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterSetMaintenanceInvocation) callAsBody() (*ActionClusterSetMaintenanceResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterSetMaintenanceResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionClusterSetMaintenanceInvocation) makeAllInputParams() *ActionClusterSetMaintenanceRequest {
	return &ActionClusterSetMaintenanceRequest{
		Cluster: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionClusterSetMaintenanceInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionClusterSetMaintenanceInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
