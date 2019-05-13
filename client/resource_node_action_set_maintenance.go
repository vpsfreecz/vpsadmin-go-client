package client

import (
	"strings"
)

// ActionNodeSetMaintenance is a type for action Node#Set_maintenance
type ActionNodeSetMaintenance struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeSetMaintenance(client *Client) *ActionNodeSetMaintenance {
	return &ActionNodeSetMaintenance{
		Client: client,
	}
}

// ActionNodeSetMaintenanceMetaGlobalInput is a type for action global meta input parameters
type ActionNodeSetMaintenanceMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeSetMaintenanceMetaGlobalInput) SetNo(value bool) *ActionNodeSetMaintenanceMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSetMaintenanceMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSetMaintenanceMetaGlobalInput) SelectParameters(params ...string) *ActionNodeSetMaintenanceMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeSetMaintenanceMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSetMaintenanceInput is a type for action input parameters
type ActionNodeSetMaintenanceInput struct {
	Lock bool `json:"lock"`
	Reason string `json:"reason"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLock sets parameter Lock to value and selects it for sending
func (in *ActionNodeSetMaintenanceInput) SetLock(value bool) *ActionNodeSetMaintenanceInput {
	in.Lock = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lock"] = nil
	return in
}
// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionNodeSetMaintenanceInput) SetReason(value string) *ActionNodeSetMaintenanceInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSetMaintenanceInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSetMaintenanceInput) SelectParameters(params ...string) *ActionNodeSetMaintenanceInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeSetMaintenanceInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSetMaintenanceRequest is a type for the entire action request
type ActionNodeSetMaintenanceRequest struct {
	Node map[string]interface{} `json:"node"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionNodeSetMaintenanceResponse struct {
	Action *ActionNodeSetMaintenance `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionNodeSetMaintenance) Prepare() *ActionNodeSetMaintenanceInvocation {
	return &ActionNodeSetMaintenanceInvocation{
		Action: action,
		Path: "/v5.0/nodes/{node_id}/set_maintenance",
	}
}

// ActionNodeSetMaintenanceInvocation is used to configure action for invocation
type ActionNodeSetMaintenanceInvocation struct {
	// Pointer to the action
	Action *ActionNodeSetMaintenance

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeSetMaintenanceInput
	// Global meta input parameters
	MetaInput *ActionNodeSetMaintenanceMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeSetMaintenanceInvocation) SetPathParamInt(param string, value int64) *ActionNodeSetMaintenanceInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeSetMaintenanceInvocation) SetPathParamString(param string, value string) *ActionNodeSetMaintenanceInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeSetMaintenanceInvocation) NewInput() *ActionNodeSetMaintenanceInput {
	inv.Input = &ActionNodeSetMaintenanceInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeSetMaintenanceInvocation) SetInput(input *ActionNodeSetMaintenanceInput) *ActionNodeSetMaintenanceInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeSetMaintenanceInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeSetMaintenanceInvocation) NewMetaInput() *ActionNodeSetMaintenanceMetaGlobalInput {
	inv.MetaInput = &ActionNodeSetMaintenanceMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeSetMaintenanceInvocation) SetMetaInput(input *ActionNodeSetMaintenanceMetaGlobalInput) *ActionNodeSetMaintenanceInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeSetMaintenanceInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeSetMaintenanceInvocation) Call() (*ActionNodeSetMaintenanceResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionNodeSetMaintenanceInvocation) callAsBody() (*ActionNodeSetMaintenanceResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNodeSetMaintenanceResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionNodeSetMaintenanceInvocation) makeAllInputParams() *ActionNodeSetMaintenanceRequest {
	return &ActionNodeSetMaintenanceRequest{
		Node: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNodeSetMaintenanceInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionNodeSetMaintenanceInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
