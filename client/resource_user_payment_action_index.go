package client

import (
)

// ActionUserPaymentIndex is a type for action User_payment#Index
type ActionUserPaymentIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserPaymentIndex(client *Client) *ActionUserPaymentIndex {
	return &ActionUserPaymentIndex{
		Client: client,
	}
}

// ActionUserPaymentIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserPaymentIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserPaymentIndexMetaGlobalInput) SetNo(value bool) *ActionUserPaymentIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserPaymentIndexMetaGlobalInput) SetCount(value bool) *ActionUserPaymentIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserPaymentIndexMetaGlobalInput) SetIncludes(value string) *ActionUserPaymentIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPaymentIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPaymentIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserPaymentIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPaymentIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserPaymentIndexInput is a type for action input parameters
type ActionUserPaymentIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	User int64 `json:"user"`
	AccountedBy int64 `json:"accounted_by"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserPaymentIndexInput) SetOffset(value int64) *ActionUserPaymentIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserPaymentIndexInput) SetLimit(value int64) *ActionUserPaymentIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserPaymentIndexInput) SetUser(value int64) *ActionUserPaymentIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetAccountedBy sets parameter AccountedBy to value and selects it for sending
func (in *ActionUserPaymentIndexInput) SetAccountedBy(value int64) *ActionUserPaymentIndexInput {
	in.AccountedBy = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AccountedBy"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPaymentIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPaymentIndexInput) SelectParameters(params ...string) *ActionUserPaymentIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPaymentIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserPaymentIndexOutput is a type for action output parameters
type ActionUserPaymentIndexOutput struct {
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
type ActionUserPaymentIndexResponse struct {
	Action *ActionUserPaymentIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserPayments []*ActionUserPaymentIndexOutput `json:"user_payments"`
	}

	// Action output without the namespace
	Output []*ActionUserPaymentIndexOutput
}


// Prepare the action for invocation
func (action *ActionUserPaymentIndex) Prepare() *ActionUserPaymentIndexInvocation {
	return &ActionUserPaymentIndexInvocation{
		Action: action,
		Path: "/v5.0/user_payments",
	}
}

// ActionUserPaymentIndexInvocation is used to configure action for invocation
type ActionUserPaymentIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserPaymentIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserPaymentIndexInput
	// Global meta input parameters
	MetaInput *ActionUserPaymentIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserPaymentIndexInvocation) NewInput() *ActionUserPaymentIndexInput {
	inv.Input = &ActionUserPaymentIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserPaymentIndexInvocation) SetInput(input *ActionUserPaymentIndexInput) *ActionUserPaymentIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserPaymentIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserPaymentIndexInvocation) NewMetaInput() *ActionUserPaymentIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserPaymentIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserPaymentIndexInvocation) SetMetaInput(input *ActionUserPaymentIndexMetaGlobalInput) *ActionUserPaymentIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserPaymentIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserPaymentIndexInvocation) Call() (*ActionUserPaymentIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserPaymentIndexInvocation) callAsQuery() (*ActionUserPaymentIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserPaymentIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserPayments
	}
	return resp, err
}



func (inv *ActionUserPaymentIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["user_payment[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["user_payment[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["user_payment[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("AccountedBy") {
			ret["user_payment[accounted_by]"] = convertInt64ToString(inv.Input.AccountedBy)
		}
	}
}

func (inv *ActionUserPaymentIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

