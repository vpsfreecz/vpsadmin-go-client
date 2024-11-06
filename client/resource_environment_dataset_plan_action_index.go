package client

import (
	"strings"
)

// ActionEnvironmentDatasetPlanIndex is a type for action Environment.Dataset_plan#Index
type ActionEnvironmentDatasetPlanIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEnvironmentDatasetPlanIndex(client *Client) *ActionEnvironmentDatasetPlanIndex {
	return &ActionEnvironmentDatasetPlanIndex{
		Client: client,
	}
}

// ActionEnvironmentDatasetPlanIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEnvironmentDatasetPlanIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEnvironmentDatasetPlanIndexMetaGlobalInput) SetCount(value bool) *ActionEnvironmentDatasetPlanIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEnvironmentDatasetPlanIndexMetaGlobalInput) SetIncludes(value string) *ActionEnvironmentDatasetPlanIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEnvironmentDatasetPlanIndexMetaGlobalInput) SetNo(value bool) *ActionEnvironmentDatasetPlanIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentDatasetPlanIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentDatasetPlanIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEnvironmentDatasetPlanIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentDatasetPlanIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentDatasetPlanIndexInput is a type for action input parameters
type ActionEnvironmentDatasetPlanIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEnvironmentDatasetPlanIndexInput) SetFromId(value int64) *ActionEnvironmentDatasetPlanIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEnvironmentDatasetPlanIndexInput) SetLimit(value int64) *ActionEnvironmentDatasetPlanIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentDatasetPlanIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentDatasetPlanIndexInput) SelectParameters(params ...string) *ActionEnvironmentDatasetPlanIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEnvironmentDatasetPlanIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEnvironmentDatasetPlanIndexInput) UnselectParameters(params ...string) *ActionEnvironmentDatasetPlanIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEnvironmentDatasetPlanIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentDatasetPlanIndexOutput is a type for action output parameters
type ActionEnvironmentDatasetPlanIndexOutput struct {
	DatasetPlan *ActionDatasetPlanShowOutput `json:"dataset_plan"`
	Id          int64                        `json:"id"`
	Label       string                       `json:"label"`
	UserAdd     bool                         `json:"user_add"`
	UserRemove  bool                         `json:"user_remove"`
}

// Type for action response, including envelope
type ActionEnvironmentDatasetPlanIndexResponse struct {
	Action *ActionEnvironmentDatasetPlanIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DatasetPlans []*ActionEnvironmentDatasetPlanIndexOutput `json:"dataset_plans"`
	}

	// Action output without the namespace
	Output []*ActionEnvironmentDatasetPlanIndexOutput
}

// Prepare the action for invocation
func (action *ActionEnvironmentDatasetPlanIndex) Prepare() *ActionEnvironmentDatasetPlanIndexInvocation {
	return &ActionEnvironmentDatasetPlanIndexInvocation{
		Action: action,
		Path:   "/v7.0/environments/{environment_id}/dataset_plans",
	}
}

// ActionEnvironmentDatasetPlanIndexInvocation is used to configure action for invocation
type ActionEnvironmentDatasetPlanIndexInvocation struct {
	// Pointer to the action
	Action *ActionEnvironmentDatasetPlanIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEnvironmentDatasetPlanIndexInput
	// Global meta input parameters
	MetaInput *ActionEnvironmentDatasetPlanIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) SetPathParamInt(param string, value int64) *ActionEnvironmentDatasetPlanIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) SetPathParamString(param string, value string) *ActionEnvironmentDatasetPlanIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) NewInput() *ActionEnvironmentDatasetPlanIndexInput {
	inv.Input = &ActionEnvironmentDatasetPlanIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) SetInput(input *ActionEnvironmentDatasetPlanIndexInput) *ActionEnvironmentDatasetPlanIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) NewMetaInput() *ActionEnvironmentDatasetPlanIndexMetaGlobalInput {
	inv.MetaInput = &ActionEnvironmentDatasetPlanIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) SetMetaInput(input *ActionEnvironmentDatasetPlanIndexMetaGlobalInput) *ActionEnvironmentDatasetPlanIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEnvironmentDatasetPlanIndexInvocation) Call() (*ActionEnvironmentDatasetPlanIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionEnvironmentDatasetPlanIndexInvocation) callAsQuery() (*ActionEnvironmentDatasetPlanIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEnvironmentDatasetPlanIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DatasetPlans
	}
	return resp, err
}

func (inv *ActionEnvironmentDatasetPlanIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["dataset_plan[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dataset_plan[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionEnvironmentDatasetPlanIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
