package client

import ()

// ActionMigrationPlanIndex is a type for action Migration_plan#Index
type ActionMigrationPlanIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMigrationPlanIndex(client *Client) *ActionMigrationPlanIndex {
	return &ActionMigrationPlanIndex{
		Client: client,
	}
}

// ActionMigrationPlanIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMigrationPlanIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMigrationPlanIndexMetaGlobalInput) SetCount(value bool) *ActionMigrationPlanIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMigrationPlanIndexMetaGlobalInput) SetIncludes(value string) *ActionMigrationPlanIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanIndexMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMigrationPlanIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMigrationPlanIndexInput is a type for action input parameters
type ActionMigrationPlanIndexInput struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	State  string `json:"state"`
	User   int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMigrationPlanIndexInput) SetLimit(value int64) *ActionMigrationPlanIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionMigrationPlanIndexInput) SetOffset(value int64) *ActionMigrationPlanIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionMigrationPlanIndexInput) SetState(value string) *ActionMigrationPlanIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionMigrationPlanIndexInput) SetUser(value int64) *ActionMigrationPlanIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanIndexInput) SelectParameters(params ...string) *ActionMigrationPlanIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMigrationPlanIndexOutput is a type for action output parameters
type ActionMigrationPlanIndexOutput struct {
	Concurrency int64                 `json:"concurrency"`
	CreatedAt   string                `json:"created_at"`
	FinishedAt  string                `json:"finished_at"`
	Id          int64                 `json:"id"`
	Reason      string                `json:"reason"`
	SendMail    bool                  `json:"send_mail"`
	State       string                `json:"state"`
	StopOnError bool                  `json:"stop_on_error"`
	User        *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionMigrationPlanIndexResponse struct {
	Action *ActionMigrationPlanIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MigrationPlans []*ActionMigrationPlanIndexOutput `json:"migration_plans"`
	}

	// Action output without the namespace
	Output []*ActionMigrationPlanIndexOutput
}

// Prepare the action for invocation
func (action *ActionMigrationPlanIndex) Prepare() *ActionMigrationPlanIndexInvocation {
	return &ActionMigrationPlanIndexInvocation{
		Action: action,
		Path:   "/v6.0/migration_plans",
	}
}

// ActionMigrationPlanIndexInvocation is used to configure action for invocation
type ActionMigrationPlanIndexInvocation struct {
	// Pointer to the action
	Action *ActionMigrationPlanIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMigrationPlanIndexInput
	// Global meta input parameters
	MetaInput *ActionMigrationPlanIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMigrationPlanIndexInvocation) NewInput() *ActionMigrationPlanIndexInput {
	inv.Input = &ActionMigrationPlanIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMigrationPlanIndexInvocation) SetInput(input *ActionMigrationPlanIndexInput) *ActionMigrationPlanIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMigrationPlanIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMigrationPlanIndexInvocation) NewMetaInput() *ActionMigrationPlanIndexMetaGlobalInput {
	inv.MetaInput = &ActionMigrationPlanIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMigrationPlanIndexInvocation) SetMetaInput(input *ActionMigrationPlanIndexMetaGlobalInput) *ActionMigrationPlanIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMigrationPlanIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMigrationPlanIndexInvocation) Call() (*ActionMigrationPlanIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMigrationPlanIndexInvocation) callAsQuery() (*ActionMigrationPlanIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMigrationPlanIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MigrationPlans
	}
	return resp, err
}

func (inv *ActionMigrationPlanIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["migration_plan[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["migration_plan[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("State") {
			ret["migration_plan[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("User") {
			ret["migration_plan[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionMigrationPlanIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
