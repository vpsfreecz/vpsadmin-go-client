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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIncomingPaymentUpdateMetaGlobalInput) SetIncludes(value string) *ActionIncomingPaymentUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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

func (in *ActionIncomingPaymentUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncomingPaymentUpdateRequest is a type for the entire action request
type ActionIncomingPaymentUpdateRequest struct {
	IncomingPayment map[string]interface{} `json:"incoming_payment"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionIncomingPaymentUpdateOutput is a type for action output parameters
type ActionIncomingPaymentUpdateOutput struct {
	Id int64 `json:"id"`
	TransactionId string `json:"transaction_id"`
	State string `json:"state"`
	Date string `json:"date"`
	Amount int64 `json:"amount"`
	Currency string `json:"currency"`
	SrcAmount int64 `json:"src_amount"`
	SrcCurrency string `json:"src_currency"`
	AccountName string `json:"account_name"`
	UserIdent string `json:"user_ident"`
	UserMessage string `json:"user_message"`
	Vs string `json:"vs"`
	Ks string `json:"ks"`
	Ss string `json:"ss"`
	TransactionType string `json:"transaction_type"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
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
		Path: "/v5.0/incoming_payments/:incoming_payment_id",
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
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
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
		Meta: inv.makeMetaInputParams(),
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
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
