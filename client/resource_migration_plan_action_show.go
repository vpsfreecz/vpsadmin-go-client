package client

import (
	"strings"
)

// ActionMigrationPlanShow is a type for action Migration_plan#Show
type ActionMigrationPlanShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMigrationPlanShow(client *Client) *ActionMigrationPlanShow {
	return &ActionMigrationPlanShow{
		Client: client,
	}
}

// ActionMigrationPlanShowMetaGlobalInput is a type for action global meta input parameters
type ActionMigrationPlanShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanShowMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMigrationPlanShowMetaGlobalInput) SetIncludes(value string) *ActionMigrationPlanShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanShowMetaGlobalInput) SelectParameters(params ...string) *ActionMigrationPlanShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionMigrationPlanShowOutput is a type for action output parameters
type ActionMigrationPlanShowOutput struct {
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
type ActionMigrationPlanShowResponse struct {
	Action *ActionMigrationPlanShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MigrationPlan *ActionMigrationPlanShowOutput `json:"migration_plan"`
	}

	// Action output without the namespace
	Output *ActionMigrationPlanShowOutput
}


// Prepare the action for invocation
func (action *ActionMigrationPlanShow) Prepare() *ActionMigrationPlanShowInvocation {
	return &ActionMigrationPlanShowInvocation{
		Action: action,
		Path: "/v5.0/migration_plans/{migration_plan_id}",
	}
}

// ActionMigrationPlanShowInvocation is used to configure action for invocation
type ActionMigrationPlanShowInvocation struct {
	// Pointer to the action
	Action *ActionMigrationPlanShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMigrationPlanShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMigrationPlanShowInvocation) SetPathParamInt(param string, value int64) *ActionMigrationPlanShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMigrationPlanShowInvocation) SetPathParamString(param string, value string) *ActionMigrationPlanShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMigrationPlanShowInvocation) NewMetaInput() *ActionMigrationPlanShowMetaGlobalInput {
	inv.MetaInput = &ActionMigrationPlanShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMigrationPlanShowInvocation) SetMetaInput(input *ActionMigrationPlanShowMetaGlobalInput) *ActionMigrationPlanShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMigrationPlanShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMigrationPlanShowInvocation) Call() (*ActionMigrationPlanShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMigrationPlanShowInvocation) callAsQuery() (*ActionMigrationPlanShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMigrationPlanShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MigrationPlan
	}
	return resp, err
}




func (inv *ActionMigrationPlanShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

