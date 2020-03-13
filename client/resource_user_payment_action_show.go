package client

import (
	"strings"
)

// ActionUserPaymentShow is a type for action User_payment#Show
type ActionUserPaymentShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserPaymentShow(client *Client) *ActionUserPaymentShow {
	return &ActionUserPaymentShow{
		Client: client,
	}
}

// ActionUserPaymentShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserPaymentShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserPaymentShowMetaGlobalInput) SetNo(value bool) *ActionUserPaymentShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserPaymentShowMetaGlobalInput) SetIncludes(value string) *ActionUserPaymentShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPaymentShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPaymentShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserPaymentShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPaymentShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserPaymentShowOutput is a type for action output parameters
type ActionUserPaymentShowOutput struct {
	Id int64 `json:"id"`
	IncomingPayment *ActionIncomingPaymentShowOutput `json:"incoming_payment"`
	User *ActionUserShowOutput `json:"user"`
	Amount int64 `json:"amount"`
	AccountedBy *ActionUserShowOutput `json:"accounted_by"`
	FromDate string `json:"from_date"`
	ToDate string `json:"to_date"`
	CreatedAt string `json:"created_at"`
}


// Type for action response, including envelope
type ActionUserPaymentShowResponse struct {
	Action *ActionUserPaymentShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserPayment *ActionUserPaymentShowOutput `json:"user_payment"`
	}

	// Action output without the namespace
	Output *ActionUserPaymentShowOutput
}


// Prepare the action for invocation
func (action *ActionUserPaymentShow) Prepare() *ActionUserPaymentShowInvocation {
	return &ActionUserPaymentShowInvocation{
		Action: action,
		Path: "/v6.0/user_payments/{user_payment_id}",
	}
}

// ActionUserPaymentShowInvocation is used to configure action for invocation
type ActionUserPaymentShowInvocation struct {
	// Pointer to the action
	Action *ActionUserPaymentShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserPaymentShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserPaymentShowInvocation) SetPathParamInt(param string, value int64) *ActionUserPaymentShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserPaymentShowInvocation) SetPathParamString(param string, value string) *ActionUserPaymentShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserPaymentShowInvocation) NewMetaInput() *ActionUserPaymentShowMetaGlobalInput {
	inv.MetaInput = &ActionUserPaymentShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserPaymentShowInvocation) SetMetaInput(input *ActionUserPaymentShowMetaGlobalInput) *ActionUserPaymentShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserPaymentShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserPaymentShowInvocation) Call() (*ActionUserPaymentShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserPaymentShowInvocation) callAsQuery() (*ActionUserPaymentShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserPaymentShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserPayment
	}
	return resp, err
}




func (inv *ActionUserPaymentShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

