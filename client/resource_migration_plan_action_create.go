package client

import (
)

// ActionMigrationPlanCreate is a type for action Migration_plan#Create
type ActionMigrationPlanCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionMigrationPlanCreate(client *Client) *ActionMigrationPlanCreate {
	return &ActionMigrationPlanCreate{
		Client: client,
	}
}

// ActionMigrationPlanCreateMetaGlobalInput is a type for action global meta input parameters
type ActionMigrationPlanCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanCreateMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMigrationPlanCreateMetaGlobalInput) SetIncludes(value string) *ActionMigrationPlanCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanCreateMetaGlobalInput) SelectParameters(params ...string) *ActionMigrationPlanCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMigrationPlanCreateInput is a type for action input parameters
type ActionMigrationPlanCreateInput struct {
	StopOnError bool `json:"stop_on_error"`
	SendMail bool `json:"send_mail"`
	Concurrency int64 `json:"concurrency"`
	Reason string `json:"reason"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetStopOnError sets parameter StopOnError to value and selects it for sending
func (in *ActionMigrationPlanCreateInput) SetStopOnError(value bool) *ActionMigrationPlanCreateInput {
	in.StopOnError = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["StopOnError"] = nil
	return in
}
// SetSendMail sets parameter SendMail to value and selects it for sending
func (in *ActionMigrationPlanCreateInput) SetSendMail(value bool) *ActionMigrationPlanCreateInput {
	in.SendMail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SendMail"] = nil
	return in
}
// SetConcurrency sets parameter Concurrency to value and selects it for sending
func (in *ActionMigrationPlanCreateInput) SetConcurrency(value int64) *ActionMigrationPlanCreateInput {
	in.Concurrency = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Concurrency"] = nil
	return in
}
// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionMigrationPlanCreateInput) SetReason(value string) *ActionMigrationPlanCreateInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanCreateInput) SelectParameters(params ...string) *ActionMigrationPlanCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMigrationPlanCreateRequest is a type for the entire action request
type ActionMigrationPlanCreateRequest struct {
	MigrationPlan map[string]interface{} `json:"migration_plan"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMigrationPlanCreateOutput is a type for action output parameters
type ActionMigrationPlanCreateOutput struct {
	Id int64 `json:"id"`
	State string `json:"state"`
	StopOnError bool `json:"stop_on_error"`
	SendMail bool `json:"send_mail"`
	Concurrency int64 `json:"concurrency"`
	Reason string `json:"reason"`
	User *ActionUserShowOutput `json:"user"`
	CreatedAt string `json:"created_at"`
	FinishedAt string `json:"finished_at"`
}


// Type for action response, including envelope
type ActionMigrationPlanCreateResponse struct {
	Action *ActionMigrationPlanCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MigrationPlan *ActionMigrationPlanCreateOutput `json:"migration_plan"`
	}

	// Action output without the namespace
	Output *ActionMigrationPlanCreateOutput
}


// Prepare the action for invocation
func (action *ActionMigrationPlanCreate) Prepare() *ActionMigrationPlanCreateInvocation {
	return &ActionMigrationPlanCreateInvocation{
		Action: action,
		Path: "/v5.0/migration_plans",
	}
}

// ActionMigrationPlanCreateInvocation is used to configure action for invocation
type ActionMigrationPlanCreateInvocation struct {
	// Pointer to the action
	Action *ActionMigrationPlanCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMigrationPlanCreateInput
	// Global meta input parameters
	MetaInput *ActionMigrationPlanCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMigrationPlanCreateInvocation) NewInput() *ActionMigrationPlanCreateInput {
	inv.Input = &ActionMigrationPlanCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMigrationPlanCreateInvocation) SetInput(input *ActionMigrationPlanCreateInput) *ActionMigrationPlanCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMigrationPlanCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMigrationPlanCreateInvocation) NewMetaInput() *ActionMigrationPlanCreateMetaGlobalInput {
	inv.MetaInput = &ActionMigrationPlanCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMigrationPlanCreateInvocation) SetMetaInput(input *ActionMigrationPlanCreateMetaGlobalInput) *ActionMigrationPlanCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMigrationPlanCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMigrationPlanCreateInvocation) Call() (*ActionMigrationPlanCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMigrationPlanCreateInvocation) callAsBody() (*ActionMigrationPlanCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMigrationPlanCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MigrationPlan
	}
	return resp, err
}




func (inv *ActionMigrationPlanCreateInvocation) makeAllInputParams() *ActionMigrationPlanCreateRequest {
	return &ActionMigrationPlanCreateRequest{
		MigrationPlan: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMigrationPlanCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("StopOnError") {
			ret["stop_on_error"] = inv.Input.StopOnError
		}
		if inv.IsParameterSelected("SendMail") {
			ret["send_mail"] = inv.Input.SendMail
		}
		if inv.IsParameterSelected("Concurrency") {
			ret["concurrency"] = inv.Input.Concurrency
		}
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
	}

	return ret
}

func (inv *ActionMigrationPlanCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
