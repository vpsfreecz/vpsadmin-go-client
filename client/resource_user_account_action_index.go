package client

import ()

// ActionUserAccountIndex is a type for action User_account#Index
type ActionUserAccountIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserAccountIndex(client *Client) *ActionUserAccountIndex {
	return &ActionUserAccountIndex{
		Client: client,
	}
}

// ActionUserAccountIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserAccountIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserAccountIndexMetaGlobalInput) SetCount(value bool) *ActionUserAccountIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserAccountIndexMetaGlobalInput) SetIncludes(value string) *ActionUserAccountIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserAccountIndexMetaGlobalInput) SetNo(value bool) *ActionUserAccountIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserAccountIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserAccountIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserAccountIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserAccountIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserAccountIndexInput is a type for action input parameters
type ActionUserAccountIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserAccountIndexInput) SetFromId(value int64) *ActionUserAccountIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserAccountIndexInput) SetLimit(value int64) *ActionUserAccountIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserAccountIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserAccountIndexInput) SelectParameters(params ...string) *ActionUserAccountIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserAccountIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserAccountIndexInput) UnselectParameters(params ...string) *ActionUserAccountIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserAccountIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserAccountIndexOutput is a type for action output parameters
type ActionUserAccountIndexOutput struct {
	Id             int64  `json:"id"`
	MonthlyPayment int64  `json:"monthly_payment"`
	PaidUntil      string `json:"paid_until"`
}

// Type for action response, including envelope
type ActionUserAccountIndexResponse struct {
	Action *ActionUserAccountIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserAccounts []*ActionUserAccountIndexOutput `json:"user_accounts"`
	}

	// Action output without the namespace
	Output []*ActionUserAccountIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserAccountIndex) Prepare() *ActionUserAccountIndexInvocation {
	return &ActionUserAccountIndexInvocation{
		Action: action,
		Path:   "/v7.0/user_accounts",
	}
}

// ActionUserAccountIndexInvocation is used to configure action for invocation
type ActionUserAccountIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserAccountIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserAccountIndexInput
	// Global meta input parameters
	MetaInput *ActionUserAccountIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserAccountIndexInvocation) NewInput() *ActionUserAccountIndexInput {
	inv.Input = &ActionUserAccountIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserAccountIndexInvocation) SetInput(input *ActionUserAccountIndexInput) *ActionUserAccountIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserAccountIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserAccountIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserAccountIndexInvocation) NewMetaInput() *ActionUserAccountIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserAccountIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserAccountIndexInvocation) SetMetaInput(input *ActionUserAccountIndexMetaGlobalInput) *ActionUserAccountIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserAccountIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserAccountIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserAccountIndexInvocation) Call() (*ActionUserAccountIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserAccountIndexInvocation) callAsQuery() (*ActionUserAccountIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserAccountIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserAccounts
	}
	return resp, err
}

func (inv *ActionUserAccountIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["user_account[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["user_account[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserAccountIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
