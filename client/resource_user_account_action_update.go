package client

import (
	"strings"
)

// ActionUserAccountUpdate is a type for action User_account#Update
type ActionUserAccountUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserAccountUpdate(client *Client) *ActionUserAccountUpdate {
	return &ActionUserAccountUpdate{
		Client: client,
	}
}

// ActionUserAccountUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserAccountUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserAccountUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserAccountUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserAccountUpdateMetaGlobalInput) SetNo(value bool) *ActionUserAccountUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserAccountUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserAccountUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserAccountUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserAccountUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserAccountUpdateInput is a type for action input parameters
type ActionUserAccountUpdateInput struct {
	MonthlyPayment int64  `json:"monthly_payment"`
	PaidUntil      string `json:"paid_until"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetMonthlyPayment sets parameter MonthlyPayment to value and selects it for sending
func (in *ActionUserAccountUpdateInput) SetMonthlyPayment(value int64) *ActionUserAccountUpdateInput {
	in.MonthlyPayment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MonthlyPayment"] = nil
	return in
}

// SetPaidUntil sets parameter PaidUntil to value and selects it for sending
func (in *ActionUserAccountUpdateInput) SetPaidUntil(value string) *ActionUserAccountUpdateInput {
	in.PaidUntil = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PaidUntil"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserAccountUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserAccountUpdateInput) SelectParameters(params ...string) *ActionUserAccountUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserAccountUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserAccountUpdateInput) UnselectParameters(params ...string) *ActionUserAccountUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserAccountUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserAccountUpdateRequest is a type for the entire action request
type ActionUserAccountUpdateRequest struct {
	UserAccount map[string]interface{} `json:"user_account"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionUserAccountUpdateOutput is a type for action output parameters
type ActionUserAccountUpdateOutput struct {
	Id             int64  `json:"id"`
	MonthlyPayment int64  `json:"monthly_payment"`
	PaidUntil      string `json:"paid_until"`
}

// Type for action response, including envelope
type ActionUserAccountUpdateResponse struct {
	Action *ActionUserAccountUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserAccount *ActionUserAccountUpdateOutput `json:"user_account"`
	}

	// Action output without the namespace
	Output *ActionUserAccountUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserAccountUpdate) Prepare() *ActionUserAccountUpdateInvocation {
	return &ActionUserAccountUpdateInvocation{
		Action: action,
		Path:   "/v6.0/user_accounts/{user_account_id}",
	}
}

// ActionUserAccountUpdateInvocation is used to configure action for invocation
type ActionUserAccountUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserAccountUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserAccountUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserAccountUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserAccountUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserAccountUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserAccountUpdateInvocation) SetPathParamString(param string, value string) *ActionUserAccountUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserAccountUpdateInvocation) NewInput() *ActionUserAccountUpdateInput {
	inv.Input = &ActionUserAccountUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserAccountUpdateInvocation) SetInput(input *ActionUserAccountUpdateInput) *ActionUserAccountUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserAccountUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserAccountUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserAccountUpdateInvocation) NewMetaInput() *ActionUserAccountUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserAccountUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserAccountUpdateInvocation) SetMetaInput(input *ActionUserAccountUpdateMetaGlobalInput) *ActionUserAccountUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserAccountUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserAccountUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserAccountUpdateInvocation) Call() (*ActionUserAccountUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserAccountUpdateInvocation) callAsBody() (*ActionUserAccountUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserAccountUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserAccount
	}
	return resp, err
}

func (inv *ActionUserAccountUpdateInvocation) makeAllInputParams() *ActionUserAccountUpdateRequest {
	return &ActionUserAccountUpdateRequest{
		UserAccount: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserAccountUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("MonthlyPayment") {
			ret["monthly_payment"] = inv.Input.MonthlyPayment
		}
		if inv.IsParameterSelected("PaidUntil") {
			ret["paid_until"] = inv.Input.PaidUntil
		}
	}

	return ret
}

func (inv *ActionUserAccountUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
