package client

import (
	"strings"
)

// ActionIncomingPaymentUpdate is a type for action Incoming_payment#Update
type ActionIncomingPaymentUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionIncomingPaymentUpdate(client *Client) *ActionIncomingPaymentUpdate {
	return &ActionIncomingPaymentUpdate{
		Client: client,
	}
}

// ActionIncomingPaymentUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionIncomingPaymentUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIncomingPaymentUpdateMetaGlobalInput) SetIncludes(value string) *ActionIncomingPaymentUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncomingPaymentUpdateMetaGlobalInput) SetNo(value bool) *ActionIncomingPaymentUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncomingPaymentUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncomingPaymentUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionIncomingPaymentUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncomingPaymentUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncomingPaymentUpdateInput is a type for action input parameters
type ActionIncomingPaymentUpdateInput struct {
	State string `json:"state"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionIncomingPaymentUpdateInput) SetState(value string) *ActionIncomingPaymentUpdateInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncomingPaymentUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncomingPaymentUpdateInput) SelectParameters(params ...string) *ActionIncomingPaymentUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIncomingPaymentUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIncomingPaymentUpdateInput) UnselectParameters(params ...string) *ActionIncomingPaymentUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIncomingPaymentUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncomingPaymentUpdateRequest is a type for the entire action request
type ActionIncomingPaymentUpdateRequest struct {
	IncomingPayment map[string]interface{} `json:"incoming_payment"`
	Meta            map[string]interface{} `json:"_meta"`
}

// ActionIncomingPaymentUpdateOutput is a type for action output parameters
type ActionIncomingPaymentUpdateOutput struct {
	AccountName     string `json:"account_name"`
	Amount          int64  `json:"amount"`
	Comment         string `json:"comment"`
	CreatedAt       string `json:"created_at"`
	Currency        string `json:"currency"`
	Date            string `json:"date"`
	Id              int64  `json:"id"`
	Ks              string `json:"ks"`
	SrcAmount       int64  `json:"src_amount"`
	SrcCurrency     string `json:"src_currency"`
	Ss              string `json:"ss"`
	State           string `json:"state"`
	TransactionId   string `json:"transaction_id"`
	TransactionType string `json:"transaction_type"`
	UserIdent       string `json:"user_ident"`
	UserMessage     string `json:"user_message"`
	Vs              string `json:"vs"`
}

// Type for action response, including envelope
type ActionIncomingPaymentUpdateResponse struct {
	Action *ActionIncomingPaymentUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncomingPayment *ActionIncomingPaymentUpdateOutput `json:"incoming_payment"`
	}

	// Action output without the namespace
	Output *ActionIncomingPaymentUpdateOutput
}

// Prepare the action for invocation
func (action *ActionIncomingPaymentUpdate) Prepare() *ActionIncomingPaymentUpdateInvocation {
	return &ActionIncomingPaymentUpdateInvocation{
		Action: action,
		Path:   "/v6.0/incoming_payments/{incoming_payment_id}",
	}
}

// ActionIncomingPaymentUpdateInvocation is used to configure action for invocation
type ActionIncomingPaymentUpdateInvocation struct {
	// Pointer to the action
	Action *ActionIncomingPaymentUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIncomingPaymentUpdateInput
	// Global meta input parameters
	MetaInput *ActionIncomingPaymentUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIncomingPaymentUpdateInvocation) SetPathParamInt(param string, value int64) *ActionIncomingPaymentUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIncomingPaymentUpdateInvocation) SetPathParamString(param string, value string) *ActionIncomingPaymentUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIncomingPaymentUpdateInvocation) NewInput() *ActionIncomingPaymentUpdateInput {
	inv.Input = &ActionIncomingPaymentUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIncomingPaymentUpdateInvocation) SetInput(input *ActionIncomingPaymentUpdateInput) *ActionIncomingPaymentUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIncomingPaymentUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIncomingPaymentUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncomingPaymentUpdateInvocation) NewMetaInput() *ActionIncomingPaymentUpdateMetaGlobalInput {
	inv.MetaInput = &ActionIncomingPaymentUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncomingPaymentUpdateInvocation) SetMetaInput(input *ActionIncomingPaymentUpdateMetaGlobalInput) *ActionIncomingPaymentUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncomingPaymentUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncomingPaymentUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncomingPaymentUpdateInvocation) Call() (*ActionIncomingPaymentUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionIncomingPaymentUpdateInvocation) callAsBody() (*ActionIncomingPaymentUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIncomingPaymentUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncomingPayment
	}
	return resp, err
}

func (inv *ActionIncomingPaymentUpdateInvocation) makeAllInputParams() *ActionIncomingPaymentUpdateRequest {
	return &ActionIncomingPaymentUpdateRequest{
		IncomingPayment: inv.makeInputParams(),
		Meta:            inv.makeMetaInputParams(),
	}
}

func (inv *ActionIncomingPaymentUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("State") {
			ret["state"] = inv.Input.State
		}
	}

	return ret
}

func (inv *ActionIncomingPaymentUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
