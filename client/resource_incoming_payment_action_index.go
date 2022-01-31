package client

import ()

// ActionIncomingPaymentIndex is a type for action Incoming_payment#Index
type ActionIncomingPaymentIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIncomingPaymentIndex(client *Client) *ActionIncomingPaymentIndex {
	return &ActionIncomingPaymentIndex{
		Client: client,
	}
}

// ActionIncomingPaymentIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIncomingPaymentIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIncomingPaymentIndexMetaGlobalInput) SetCount(value bool) *ActionIncomingPaymentIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIncomingPaymentIndexMetaGlobalInput) SetIncludes(value string) *ActionIncomingPaymentIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncomingPaymentIndexMetaGlobalInput) SetNo(value bool) *ActionIncomingPaymentIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncomingPaymentIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncomingPaymentIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIncomingPaymentIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncomingPaymentIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncomingPaymentIndexInput is a type for action input parameters
type ActionIncomingPaymentIndexInput struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	State  string `json:"state"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIncomingPaymentIndexInput) SetLimit(value int64) *ActionIncomingPaymentIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIncomingPaymentIndexInput) SetOffset(value int64) *ActionIncomingPaymentIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionIncomingPaymentIndexInput) SetState(value string) *ActionIncomingPaymentIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncomingPaymentIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncomingPaymentIndexInput) SelectParameters(params ...string) *ActionIncomingPaymentIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIncomingPaymentIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIncomingPaymentIndexInput) UnselectParameters(params ...string) *ActionIncomingPaymentIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIncomingPaymentIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncomingPaymentIndexOutput is a type for action output parameters
type ActionIncomingPaymentIndexOutput struct {
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
type ActionIncomingPaymentIndexResponse struct {
	Action *ActionIncomingPaymentIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncomingPayments []*ActionIncomingPaymentIndexOutput `json:"incoming_payments"`
	}

	// Action output without the namespace
	Output []*ActionIncomingPaymentIndexOutput
}

// Prepare the action for invocation
func (action *ActionIncomingPaymentIndex) Prepare() *ActionIncomingPaymentIndexInvocation {
	return &ActionIncomingPaymentIndexInvocation{
		Action: action,
		Path:   "/v6.0/incoming_payments",
	}
}

// ActionIncomingPaymentIndexInvocation is used to configure action for invocation
type ActionIncomingPaymentIndexInvocation struct {
	// Pointer to the action
	Action *ActionIncomingPaymentIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIncomingPaymentIndexInput
	// Global meta input parameters
	MetaInput *ActionIncomingPaymentIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIncomingPaymentIndexInvocation) NewInput() *ActionIncomingPaymentIndexInput {
	inv.Input = &ActionIncomingPaymentIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIncomingPaymentIndexInvocation) SetInput(input *ActionIncomingPaymentIndexInput) *ActionIncomingPaymentIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIncomingPaymentIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIncomingPaymentIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncomingPaymentIndexInvocation) NewMetaInput() *ActionIncomingPaymentIndexMetaGlobalInput {
	inv.MetaInput = &ActionIncomingPaymentIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncomingPaymentIndexInvocation) SetMetaInput(input *ActionIncomingPaymentIndexMetaGlobalInput) *ActionIncomingPaymentIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncomingPaymentIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncomingPaymentIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncomingPaymentIndexInvocation) Call() (*ActionIncomingPaymentIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIncomingPaymentIndexInvocation) callAsQuery() (*ActionIncomingPaymentIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncomingPaymentIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncomingPayments
	}
	return resp, err
}

func (inv *ActionIncomingPaymentIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["incoming_payment[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["incoming_payment[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("State") {
			ret["incoming_payment[state]"] = inv.Input.State
		}
	}
}

func (inv *ActionIncomingPaymentIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
