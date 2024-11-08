package client

import (
	"strings"
)

// ActionNodeEvacuate is a type for action Node#Evacuate
type ActionNodeEvacuate struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeEvacuate(client *Client) *ActionNodeEvacuate {
	return &ActionNodeEvacuate{
		Client: client,
	}
}

// ActionNodeEvacuateMetaGlobalInput is a type for action global meta input parameters
type ActionNodeEvacuateMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeEvacuateMetaGlobalInput) SetNo(value bool) *ActionNodeEvacuateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeEvacuateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeEvacuateMetaGlobalInput) SelectParameters(params ...string) *ActionNodeEvacuateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeEvacuateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeEvacuateInput is a type for action input parameters
type ActionNodeEvacuateInput struct {
	CleanupData       bool   `json:"cleanup_data"`
	Concurrency       int64  `json:"concurrency"`
	DstNode           int64  `json:"dst_node"`
	MaintenanceWindow bool   `json:"maintenance_window"`
	Reason            string `json:"reason"`
	SendMail          bool   `json:"send_mail"`
	StopOnError       bool   `json:"stop_on_error"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCleanupData sets parameter CleanupData to value and selects it for sending
func (in *ActionNodeEvacuateInput) SetCleanupData(value bool) *ActionNodeEvacuateInput {
	in.CleanupData = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CleanupData"] = nil
	return in
}

// SetConcurrency sets parameter Concurrency to value and selects it for sending
func (in *ActionNodeEvacuateInput) SetConcurrency(value int64) *ActionNodeEvacuateInput {
	in.Concurrency = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Concurrency"] = nil
	return in
}

// SetDstNode sets parameter DstNode to value and selects it for sending
func (in *ActionNodeEvacuateInput) SetDstNode(value int64) *ActionNodeEvacuateInput {
	in.DstNode = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDstNodeNil(false)
	in._selectedParameters["DstNode"] = nil
	return in
}

// SetDstNodeNil sets parameter DstNode to nil and selects it for sending
func (in *ActionNodeEvacuateInput) SetDstNodeNil(set bool) *ActionNodeEvacuateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["DstNode"] = nil
		in.SelectParameters("DstNode")
	} else {
		delete(in._nilParameters, "DstNode")
	}
	return in
}

// SetMaintenanceWindow sets parameter MaintenanceWindow to value and selects it for sending
func (in *ActionNodeEvacuateInput) SetMaintenanceWindow(value bool) *ActionNodeEvacuateInput {
	in.MaintenanceWindow = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaintenanceWindow"] = nil
	return in
}

// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionNodeEvacuateInput) SetReason(value string) *ActionNodeEvacuateInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SetSendMail sets parameter SendMail to value and selects it for sending
func (in *ActionNodeEvacuateInput) SetSendMail(value bool) *ActionNodeEvacuateInput {
	in.SendMail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SendMail"] = nil
	return in
}

// SetStopOnError sets parameter StopOnError to value and selects it for sending
func (in *ActionNodeEvacuateInput) SetStopOnError(value bool) *ActionNodeEvacuateInput {
	in.StopOnError = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["StopOnError"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeEvacuateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeEvacuateInput) SelectParameters(params ...string) *ActionNodeEvacuateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeEvacuateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeEvacuateInput) UnselectParameters(params ...string) *ActionNodeEvacuateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeEvacuateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeEvacuateRequest is a type for the entire action request
type ActionNodeEvacuateRequest struct {
	Node map[string]interface{} `json:"node"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionNodeEvacuateOutput is a type for action output parameters
type ActionNodeEvacuateOutput struct {
	MigrationPlanId int64 `json:"migration_plan_id"`
}

// Type for action response, including envelope
type ActionNodeEvacuateResponse struct {
	Action *ActionNodeEvacuate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Node *ActionNodeEvacuateOutput `json:"node"`
	}

	// Action output without the namespace
	Output *ActionNodeEvacuateOutput
}

// Prepare the action for invocation
func (action *ActionNodeEvacuate) Prepare() *ActionNodeEvacuateInvocation {
	return &ActionNodeEvacuateInvocation{
		Action: action,
		Path:   "/v7.0/nodes/{node_id}/evacuate",
	}
}

// ActionNodeEvacuateInvocation is used to configure action for invocation
type ActionNodeEvacuateInvocation struct {
	// Pointer to the action
	Action *ActionNodeEvacuate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeEvacuateInput
	// Global meta input parameters
	MetaInput *ActionNodeEvacuateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeEvacuateInvocation) SetPathParamInt(param string, value int64) *ActionNodeEvacuateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeEvacuateInvocation) SetPathParamString(param string, value string) *ActionNodeEvacuateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeEvacuateInvocation) NewInput() *ActionNodeEvacuateInput {
	inv.Input = &ActionNodeEvacuateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeEvacuateInvocation) SetInput(input *ActionNodeEvacuateInput) *ActionNodeEvacuateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeEvacuateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeEvacuateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeEvacuateInvocation) NewMetaInput() *ActionNodeEvacuateMetaGlobalInput {
	inv.MetaInput = &ActionNodeEvacuateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeEvacuateInvocation) SetMetaInput(input *ActionNodeEvacuateMetaGlobalInput) *ActionNodeEvacuateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeEvacuateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeEvacuateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeEvacuateInvocation) Call() (*ActionNodeEvacuateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionNodeEvacuateInvocation) callAsBody() (*ActionNodeEvacuateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNodeEvacuateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Node
	}
	return resp, err
}

func (inv *ActionNodeEvacuateInvocation) makeAllInputParams() *ActionNodeEvacuateRequest {
	return &ActionNodeEvacuateRequest{
		Node: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNodeEvacuateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("CleanupData") {
			ret["cleanup_data"] = inv.Input.CleanupData
		}
		if inv.IsParameterSelected("Concurrency") {
			ret["concurrency"] = inv.Input.Concurrency
		}
		if inv.IsParameterSelected("DstNode") {
			if inv.IsParameterNil("DstNode") {
				ret["dst_node"] = nil
			} else {
				ret["dst_node"] = inv.Input.DstNode
			}
		}
		if inv.IsParameterSelected("MaintenanceWindow") {
			ret["maintenance_window"] = inv.Input.MaintenanceWindow
		}
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
		if inv.IsParameterSelected("SendMail") {
			ret["send_mail"] = inv.Input.SendMail
		}
		if inv.IsParameterSelected("StopOnError") {
			ret["stop_on_error"] = inv.Input.StopOnError
		}
	}

	return ret
}

func (inv *ActionNodeEvacuateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
