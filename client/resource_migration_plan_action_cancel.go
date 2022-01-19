package client

import (
	"strings"
)

// ActionMigrationPlanCancel is a type for action Migration_plan#Cancel
type ActionMigrationPlanCancel struct {
	// Pointer to client
	Client *Client
}

func NewActionMigrationPlanCancel(client *Client) *ActionMigrationPlanCancel {
	return &ActionMigrationPlanCancel{
		Client: client,
	}
}

// ActionMigrationPlanCancelMetaGlobalInput is a type for action global meta input parameters
type ActionMigrationPlanCancelMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMigrationPlanCancelMetaGlobalInput) SetIncludes(value string) *ActionMigrationPlanCancelMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanCancelMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanCancelMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanCancelMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanCancelMetaGlobalInput) SelectParameters(params ...string) *ActionMigrationPlanCancelMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanCancelMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionMigrationPlanCancelRequest is a type for the entire action request
type ActionMigrationPlanCancelRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionMigrationPlanCancelOutput is a type for action output parameters
type ActionMigrationPlanCancelOutput struct {
	Concurrency int64 `json:"concurrency"`
	CreatedAt string `json:"created_at"`
	FinishedAt string `json:"finished_at"`
	Id int64 `json:"id"`
	Reason string `json:"reason"`
	SendMail bool `json:"send_mail"`
	State string `json:"state"`
	StopOnError bool `json:"stop_on_error"`
	User *ActionUserShowOutput `json:"user"`
}


// Type for action response, including envelope
type ActionMigrationPlanCancelResponse struct {
	Action *ActionMigrationPlanCancel `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MigrationPlan *ActionMigrationPlanCancelOutput `json:"migration_plan"`
	}

	// Action output without the namespace
	Output *ActionMigrationPlanCancelOutput
}


// Prepare the action for invocation
func (action *ActionMigrationPlanCancel) Prepare() *ActionMigrationPlanCancelInvocation {
	return &ActionMigrationPlanCancelInvocation{
		Action: action,
		Path: "/v6.0/migration_plans/{migration_plan_id}/cancel",
	}
}

// ActionMigrationPlanCancelInvocation is used to configure action for invocation
type ActionMigrationPlanCancelInvocation struct {
	// Pointer to the action
	Action *ActionMigrationPlanCancel

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMigrationPlanCancelMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMigrationPlanCancelInvocation) SetPathParamInt(param string, value int64) *ActionMigrationPlanCancelInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMigrationPlanCancelInvocation) SetPathParamString(param string, value string) *ActionMigrationPlanCancelInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMigrationPlanCancelInvocation) NewMetaInput() *ActionMigrationPlanCancelMetaGlobalInput {
	inv.MetaInput = &ActionMigrationPlanCancelMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMigrationPlanCancelInvocation) SetMetaInput(input *ActionMigrationPlanCancelMetaGlobalInput) *ActionMigrationPlanCancelInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMigrationPlanCancelInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMigrationPlanCancelInvocation) Call() (*ActionMigrationPlanCancelResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionMigrationPlanCancelInvocation) callAsBody() (*ActionMigrationPlanCancelResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMigrationPlanCancelResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MigrationPlan
	}
	return resp, err
}




func (inv *ActionMigrationPlanCancelInvocation) makeAllInputParams() *ActionMigrationPlanCancelRequest {
	return &ActionMigrationPlanCancelRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionMigrationPlanCancelInvocation) makeMetaInputParams() map[string]interface{} {
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
