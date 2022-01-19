package client

import (
	"strings"
)

// ActionIncomingPaymentShow is a type for action Incoming_payment#Show
type ActionIncomingPaymentShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIncomingPaymentShow(client *Client) *ActionIncomingPaymentShow {
	return &ActionIncomingPaymentShow{
		Client: client,
	}
}

// ActionIncomingPaymentShowMetaGlobalInput is a type for action global meta input parameters
type ActionIncomingPaymentShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIncomingPaymentShowMetaGlobalInput) SetIncludes(value string) *ActionIncomingPaymentShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncomingPaymentShowMetaGlobalInput) SetNo(value bool) *ActionIncomingPaymentShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncomingPaymentShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncomingPaymentShowMetaGlobalInput) SelectParameters(params ...string) *ActionIncomingPaymentShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncomingPaymentShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionIncomingPaymentShowOutput is a type for action output parameters
type ActionIncomingPaymentShowOutput struct {
	AccountName string `json:"account_name"`
	Amount int64 `json:"amount"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
	Currency string `json:"currency"`
	Date string `json:"date"`
	Id int64 `json:"id"`
	Ks string `json:"ks"`
	SrcAmount int64 `json:"src_amount"`
	SrcCurrency string `json:"src_currency"`
	Ss string `json:"ss"`
	State string `json:"state"`
	TransactionId string `json:"transaction_id"`
	TransactionType string `json:"transaction_type"`
	UserIdent string `json:"user_ident"`
	UserMessage string `json:"user_message"`
	Vs string `json:"vs"`
}


// Type for action response, including envelope
type ActionIncomingPaymentShowResponse struct {
	Action *ActionIncomingPaymentShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncomingPayment *ActionIncomingPaymentShowOutput `json:"incoming_payment"`
	}

	// Action output without the namespace
	Output *ActionIncomingPaymentShowOutput
}


// Prepare the action for invocation
func (action *ActionIncomingPaymentShow) Prepare() *ActionIncomingPaymentShowInvocation {
	return &ActionIncomingPaymentShowInvocation{
		Action: action,
		Path: "/v6.0/incoming_payments/{incoming_payment_id}",
	}
}

// ActionIncomingPaymentShowInvocation is used to configure action for invocation
type ActionIncomingPaymentShowInvocation struct {
	// Pointer to the action
	Action *ActionIncomingPaymentShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIncomingPaymentShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIncomingPaymentShowInvocation) SetPathParamInt(param string, value int64) *ActionIncomingPaymentShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIncomingPaymentShowInvocation) SetPathParamString(param string, value string) *ActionIncomingPaymentShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncomingPaymentShowInvocation) NewMetaInput() *ActionIncomingPaymentShowMetaGlobalInput {
	inv.MetaInput = &ActionIncomingPaymentShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncomingPaymentShowInvocation) SetMetaInput(input *ActionIncomingPaymentShowMetaGlobalInput) *ActionIncomingPaymentShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncomingPaymentShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncomingPaymentShowInvocation) Call() (*ActionIncomingPaymentShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIncomingPaymentShowInvocation) callAsQuery() (*ActionIncomingPaymentShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncomingPaymentShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncomingPayment
	}
	return resp, err
}




func (inv *ActionIncomingPaymentShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

