package client

import (
	"strings"
)

// ActionDatasetPlanCreate is a type for action Dataset.Plan#Create
type ActionDatasetPlanCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetPlanCreate(client *Client) *ActionDatasetPlanCreate {
	return &ActionDatasetPlanCreate{
		Client: client,
	}
}

// ActionDatasetPlanCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetPlanCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetPlanCreateMetaGlobalInput) SetIncludes(value string) *ActionDatasetPlanCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetPlanCreateMetaGlobalInput) SetNo(value bool) *ActionDatasetPlanCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPlanCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPlanCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetPlanCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetPlanCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetPlanCreateInput is a type for action input parameters
type ActionDatasetPlanCreateInput struct {
	EnvironmentDatasetPlan int64 `json:"environment_dataset_plan"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnvironmentDatasetPlan sets parameter EnvironmentDatasetPlan to value and selects it for sending
func (in *ActionDatasetPlanCreateInput) SetEnvironmentDatasetPlan(value int64) *ActionDatasetPlanCreateInput {
	in.EnvironmentDatasetPlan = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnvironmentDatasetPlan"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPlanCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPlanCreateInput) SelectParameters(params ...string) *ActionDatasetPlanCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetPlanCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetPlanCreateRequest is a type for the entire action request
type ActionDatasetPlanCreateRequest struct {
	Plan map[string]interface{} `json:"plan"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionDatasetPlanCreateOutput is a type for action output parameters
type ActionDatasetPlanCreateOutput struct {
	EnvironmentDatasetPlan *ActionEnvironmentDatasetPlanShowOutput `json:"environment_dataset_plan"`
	Id                     int64                                   `json:"id"`
}

// Type for action response, including envelope
type ActionDatasetPlanCreateResponse struct {
	Action *ActionDatasetPlanCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Plan *ActionDatasetPlanCreateOutput `json:"plan"`
	}

	// Action output without the namespace
	Output *ActionDatasetPlanCreateOutput
}

// Prepare the action for invocation
func (action *ActionDatasetPlanCreate) Prepare() *ActionDatasetPlanCreateInvocation {
	return &ActionDatasetPlanCreateInvocation{
		Action: action,
		Path:   "/v6.0/datasets/{dataset_id}/plans",
	}
}

// ActionDatasetPlanCreateInvocation is used to configure action for invocation
type ActionDatasetPlanCreateInvocation struct {
	// Pointer to the action
	Action *ActionDatasetPlanCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetPlanCreateInput
	// Global meta input parameters
	MetaInput *ActionDatasetPlanCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetPlanCreateInvocation) SetPathParamInt(param string, value int64) *ActionDatasetPlanCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetPlanCreateInvocation) SetPathParamString(param string, value string) *ActionDatasetPlanCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetPlanCreateInvocation) NewInput() *ActionDatasetPlanCreateInput {
	inv.Input = &ActionDatasetPlanCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetPlanCreateInvocation) SetInput(input *ActionDatasetPlanCreateInput) *ActionDatasetPlanCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetPlanCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetPlanCreateInvocation) NewMetaInput() *ActionDatasetPlanCreateMetaGlobalInput {
	inv.MetaInput = &ActionDatasetPlanCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetPlanCreateInvocation) SetMetaInput(input *ActionDatasetPlanCreateMetaGlobalInput) *ActionDatasetPlanCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetPlanCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetPlanCreateInvocation) Call() (*ActionDatasetPlanCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetPlanCreateInvocation) callAsBody() (*ActionDatasetPlanCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetPlanCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Plan
	}
	return resp, err
}

func (inv *ActionDatasetPlanCreateInvocation) makeAllInputParams() *ActionDatasetPlanCreateRequest {
	return &ActionDatasetPlanCreateRequest{
		Plan: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetPlanCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("EnvironmentDatasetPlan") {
			ret["environment_dataset_plan"] = inv.Input.EnvironmentDatasetPlan
		}
	}

	return ret
}

func (inv *ActionDatasetPlanCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
