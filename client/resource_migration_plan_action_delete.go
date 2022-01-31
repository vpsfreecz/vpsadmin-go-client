package client

import (
	"strings"
)

// ActionMigrationPlanDelete is a type for action Migration_plan#Delete
type ActionMigrationPlanDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionMigrationPlanDelete(client *Client) *ActionMigrationPlanDelete {
	return &ActionMigrationPlanDelete{
		Client: client,
	}
}

// ActionMigrationPlanDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionMigrationPlanDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMigrationPlanDeleteMetaGlobalInput) SetIncludes(value string) *ActionMigrationPlanDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMigrationPlanDeleteMetaGlobalInput) SetNo(value bool) *ActionMigrationPlanDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMigrationPlanDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMigrationPlanDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionMigrationPlanDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMigrationPlanDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMigrationPlanDeleteRequest is a type for the entire action request
type ActionMigrationPlanDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionMigrationPlanDeleteResponse struct {
	Action *ActionMigrationPlanDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionMigrationPlanDelete) Prepare() *ActionMigrationPlanDeleteInvocation {
	return &ActionMigrationPlanDeleteInvocation{
		Action: action,
		Path:   "/v6.0/migration_plans/{migration_plan_id}",
	}
}

// ActionMigrationPlanDeleteInvocation is used to configure action for invocation
type ActionMigrationPlanDeleteInvocation struct {
	// Pointer to the action
	Action *ActionMigrationPlanDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMigrationPlanDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMigrationPlanDeleteInvocation) SetPathParamInt(param string, value int64) *ActionMigrationPlanDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMigrationPlanDeleteInvocation) SetPathParamString(param string, value string) *ActionMigrationPlanDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMigrationPlanDeleteInvocation) NewMetaInput() *ActionMigrationPlanDeleteMetaGlobalInput {
	inv.MetaInput = &ActionMigrationPlanDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMigrationPlanDeleteInvocation) SetMetaInput(input *ActionMigrationPlanDeleteMetaGlobalInput) *ActionMigrationPlanDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMigrationPlanDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMigrationPlanDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMigrationPlanDeleteInvocation) Call() (*ActionMigrationPlanDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMigrationPlanDeleteInvocation) callAsBody() (*ActionMigrationPlanDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMigrationPlanDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionMigrationPlanDeleteInvocation) makeAllInputParams() *ActionMigrationPlanDeleteRequest {
	return &ActionMigrationPlanDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionMigrationPlanDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
