package client

import (
	"strings"
)

// ActionMigrationPlanStart is a type for action Migration_plan#Start
type ActionMigrationPlanStart struct {
	// Pointer to client
	Client *Client
}

func NewActionMigrationPlanStart(client *Client) *ActionMigrationPlanStart {
	return &ActionMigrationPlanStart{
		Client: client,
	}
}

// ActionMigrationPlanStartMetaGlobalInput is a type for action global meta input parameters
type ActionMigrationPlanStartMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanStartMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanStartMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMigrationPlanStartMetaGlobalInput) SetIncludes(value string) *ActionMigrationPlanStartMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanStartMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanStartMetaGlobalInput) SelectParameters(params ...string) *ActionMigrationPlanStartMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanStartMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionMigrationPlanStartRequest is a type for the entire action request
type ActionMigrationPlanStartRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMigrationPlanStartOutput is a type for action output parameters
type ActionMigrationPlanStartOutput struct {
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
type ActionMigrationPlanStartResponse struct {
	Action *ActionMigrationPlanStart `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MigrationPlan *ActionMigrationPlanStartOutput `json:"migration_plan"`
	}

	// Action output without the namespace
	Output *ActionMigrationPlanStartOutput
}


// Prepare the action for invocation
func (action *ActionMigrationPlanStart) Prepare() *ActionMigrationPlanStartInvocation {
	return &ActionMigrationPlanStartInvocation{
		Action: action,
		Path: "/v5.0/migration_plans/:migration_plan_id/start",
	}
}

// ActionMigrationPlanStartInvocation is used to configure action for invocation
type ActionMigrationPlanStartInvocation struct {
	// Pointer to the action
	Action *ActionMigrationPlanStart

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMigrationPlanStartMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMigrationPlanStartInvocation) SetPathParamInt(param string, value int64) *ActionMigrationPlanStartInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMigrationPlanStartInvocation) SetPathParamString(param string, value string) *ActionMigrationPlanStartInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMigrationPlanStartInvocation) SetMetaInput(input *ActionMigrationPlanStartMetaGlobalInput) *ActionMigrationPlanStartInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMigrationPlanStartInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMigrationPlanStartInvocation) Call() (*ActionMigrationPlanStartResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMigrationPlanStartInvocation) callAsBody() (*ActionMigrationPlanStartResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMigrationPlanStartResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MigrationPlan
	}
	return resp, err
}




func (inv *ActionMigrationPlanStartInvocation) makeAllInputParams() *ActionMigrationPlanStartRequest {
	return &ActionMigrationPlanStartRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionMigrationPlanStartInvocation) makeMetaInputParams() map[string]interface{} {
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
