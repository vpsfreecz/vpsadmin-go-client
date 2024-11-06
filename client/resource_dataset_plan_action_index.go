package client

import ()

// ActionDatasetPlanIndex is a type for action Dataset_plan#Index
type ActionDatasetPlanIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetPlanIndex(client *Client) *ActionDatasetPlanIndex {
	return &ActionDatasetPlanIndex{
		Client: client,
	}
}

// ActionDatasetPlanIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetPlanIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDatasetPlanIndexMetaGlobalInput) SetCount(value bool) *ActionDatasetPlanIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetPlanIndexMetaGlobalInput) SetIncludes(value string) *ActionDatasetPlanIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetPlanIndexMetaGlobalInput) SetNo(value bool) *ActionDatasetPlanIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPlanIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPlanIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetPlanIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetPlanIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetPlanIndexInput is a type for action input parameters
type ActionDatasetPlanIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDatasetPlanIndexInput) SetFromId(value int64) *ActionDatasetPlanIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDatasetPlanIndexInput) SetLimit(value int64) *ActionDatasetPlanIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPlanIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPlanIndexInput) SelectParameters(params ...string) *ActionDatasetPlanIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetPlanIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetPlanIndexInput) UnselectParameters(params ...string) *ActionDatasetPlanIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetPlanIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetPlanIndexOutput is a type for action output parameters
type ActionDatasetPlanIndexOutput struct {
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Label       string `json:"label"`
}

// Type for action response, including envelope
type ActionDatasetPlanIndexResponse struct {
	Action *ActionDatasetPlanIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DatasetPlans []*ActionDatasetPlanIndexOutput `json:"dataset_plans"`
	}

	// Action output without the namespace
	Output []*ActionDatasetPlanIndexOutput
}

// Prepare the action for invocation
func (action *ActionDatasetPlanIndex) Prepare() *ActionDatasetPlanIndexInvocation {
	return &ActionDatasetPlanIndexInvocation{
		Action: action,
		Path:   "/v7.0/dataset_plans",
	}
}

// ActionDatasetPlanIndexInvocation is used to configure action for invocation
type ActionDatasetPlanIndexInvocation struct {
	// Pointer to the action
	Action *ActionDatasetPlanIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetPlanIndexInput
	// Global meta input parameters
	MetaInput *ActionDatasetPlanIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetPlanIndexInvocation) NewInput() *ActionDatasetPlanIndexInput {
	inv.Input = &ActionDatasetPlanIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetPlanIndexInvocation) SetInput(input *ActionDatasetPlanIndexInput) *ActionDatasetPlanIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetPlanIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetPlanIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetPlanIndexInvocation) NewMetaInput() *ActionDatasetPlanIndexMetaGlobalInput {
	inv.MetaInput = &ActionDatasetPlanIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetPlanIndexInvocation) SetMetaInput(input *ActionDatasetPlanIndexMetaGlobalInput) *ActionDatasetPlanIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetPlanIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetPlanIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetPlanIndexInvocation) Call() (*ActionDatasetPlanIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetPlanIndexInvocation) callAsQuery() (*ActionDatasetPlanIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetPlanIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DatasetPlans
	}
	return resp, err
}

func (inv *ActionDatasetPlanIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["dataset_plan[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dataset_plan[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionDatasetPlanIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
