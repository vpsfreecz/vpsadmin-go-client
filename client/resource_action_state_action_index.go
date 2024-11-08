package client

import ()

// ActionActionStateIndex is a type for action Action_state#Index
type ActionActionStateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionActionStateIndex(client *Client) *ActionActionStateIndex {
	return &ActionActionStateIndex{
		Client: client,
	}
}

// ActionActionStateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionActionStateIndexMetaGlobalInput struct {
	Count bool `json:"count"`
	No    bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionActionStateIndexMetaGlobalInput) SetCount(value bool) *ActionActionStateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionActionStateIndexMetaGlobalInput) SetNo(value bool) *ActionActionStateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionActionStateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionActionStateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionActionStateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionActionStateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionActionStateIndexInput is a type for action input parameters
type ActionActionStateIndexInput struct {
	FromId int64  `json:"from_id"`
	Limit  int64  `json:"limit"`
	Order  string `json:"order"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionActionStateIndexInput) SetFromId(value int64) *ActionActionStateIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionActionStateIndexInput) SetLimit(value int64) *ActionActionStateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionActionStateIndexInput) SetOrder(value string) *ActionActionStateIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SelectParameters sets parameters from ActionActionStateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionActionStateIndexInput) SelectParameters(params ...string) *ActionActionStateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionActionStateIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionActionStateIndexInput) UnselectParameters(params ...string) *ActionActionStateIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionActionStateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionActionStateIndexOutput is a type for action output parameters
type ActionActionStateIndexOutput struct {
	CanCancel bool   `json:"can_cancel"`
	CreatedAt string `json:"created_at"`
	Current   int64  `json:"current"`
	Finished  bool   `json:"finished"`
	Id        int64  `json:"id"`
	Label     string `json:"label"`
	Status    bool   `json:"status"`
	Total     int64  `json:"total"`
	Unit      string `json:"unit"`
	UpdatedAt string `json:"updated_at"`
}

// Type for action response, including envelope
type ActionActionStateIndexResponse struct {
	Action *ActionActionStateIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ActionStates []*ActionActionStateIndexOutput `json:"action_states"`
	}

	// Action output without the namespace
	Output []*ActionActionStateIndexOutput
}

// Prepare the action for invocation
func (action *ActionActionStateIndex) Prepare() *ActionActionStateIndexInvocation {
	return &ActionActionStateIndexInvocation{
		Action: action,
		Path:   "/v7.0/action_states",
	}
}

// ActionActionStateIndexInvocation is used to configure action for invocation
type ActionActionStateIndexInvocation struct {
	// Pointer to the action
	Action *ActionActionStateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionActionStateIndexInput
	// Global meta input parameters
	MetaInput *ActionActionStateIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionActionStateIndexInvocation) NewInput() *ActionActionStateIndexInput {
	inv.Input = &ActionActionStateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionActionStateIndexInvocation) SetInput(input *ActionActionStateIndexInput) *ActionActionStateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionActionStateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionActionStateIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionActionStateIndexInvocation) NewMetaInput() *ActionActionStateIndexMetaGlobalInput {
	inv.MetaInput = &ActionActionStateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionActionStateIndexInvocation) SetMetaInput(input *ActionActionStateIndexMetaGlobalInput) *ActionActionStateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionActionStateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionActionStateIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionActionStateIndexInvocation) Call() (*ActionActionStateIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionActionStateIndexInvocation) callAsQuery() (*ActionActionStateIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionActionStateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ActionStates
	}
	return resp, err
}

func (inv *ActionActionStateIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["action_state[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["action_state[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Order") {
			ret["action_state[order]"] = inv.Input.Order
		}
	}
}

func (inv *ActionActionStateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
