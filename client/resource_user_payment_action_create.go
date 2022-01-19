package client

import (
)

// ActionUserPaymentCreate is a type for action User_payment#Create
type ActionUserPaymentCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserPaymentCreate(client *Client) *ActionUserPaymentCreate {
	return &ActionUserPaymentCreate{
		Client: client,
	}
}

// ActionUserPaymentCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserPaymentCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserPaymentCreateMetaGlobalInput) SetIncludes(value string) *ActionUserPaymentCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserPaymentCreateMetaGlobalInput) SetNo(value bool) *ActionUserPaymentCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPaymentCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPaymentCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserPaymentCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPaymentCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserPaymentCreateInput is a type for action input parameters
type ActionUserPaymentCreateInput struct {
	Amount int64 `json:"amount"`
	IncomingPayment int64 `json:"incoming_payment"`
	User int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetAmount sets parameter Amount to value and selects it for sending
func (in *ActionUserPaymentCreateInput) SetAmount(value int64) *ActionUserPaymentCreateInput {
	in.Amount = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Amount"] = nil
	return in
}
// SetIncomingPayment sets parameter IncomingPayment to value and selects it for sending
func (in *ActionUserPaymentCreateInput) SetIncomingPayment(value int64) *ActionUserPaymentCreateInput {
	in.IncomingPayment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IncomingPayment"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserPaymentCreateInput) SetUser(value int64) *ActionUserPaymentCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPaymentCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPaymentCreateInput) SelectParameters(params ...string) *ActionUserPaymentCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPaymentCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserPaymentCreateRequest is a type for the entire action request
type ActionUserPaymentCreateRequest struct {
	UserPayment map[string]interface{} `json:"user_payment"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserPaymentCreateOutput is a type for action output parameters
type ActionUserPaymentCreateOutput struct {
	AccountedBy *ActionUserShowOutput `json:"accounted_by"`
	Amount int64 `json:"amount"`
	CreatedAt string `json:"created_at"`
	FromDate string `json:"from_date"`
	Id int64 `json:"id"`
	IncomingPayment *ActionIncomingPaymentShowOutput `json:"incoming_payment"`
	ToDate string `json:"to_date"`
	User *ActionUserShowOutput `json:"user"`
}

// ActionUserPaymentCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionUserPaymentCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionUserPaymentCreateResponse struct {
	Action *ActionUserPaymentCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserPayment *ActionUserPaymentCreateOutput `json:"user_payment"`
		// Global output metadata
		Meta *ActionUserPaymentCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionUserPaymentCreateOutput
}


// Prepare the action for invocation
func (action *ActionUserPaymentCreate) Prepare() *ActionUserPaymentCreateInvocation {
	return &ActionUserPaymentCreateInvocation{
		Action: action,
		Path: "/v6.0/user_payments",
	}
}

// ActionUserPaymentCreateInvocation is used to configure action for invocation
type ActionUserPaymentCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserPaymentCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserPaymentCreateInput
	// Global meta input parameters
	MetaInput *ActionUserPaymentCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserPaymentCreateInvocation) NewInput() *ActionUserPaymentCreateInput {
	inv.Input = &ActionUserPaymentCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserPaymentCreateInvocation) SetInput(input *ActionUserPaymentCreateInput) *ActionUserPaymentCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserPaymentCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserPaymentCreateInvocation) NewMetaInput() *ActionUserPaymentCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserPaymentCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserPaymentCreateInvocation) SetMetaInput(input *ActionUserPaymentCreateMetaGlobalInput) *ActionUserPaymentCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserPaymentCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserPaymentCreateInvocation) Call() (*ActionUserPaymentCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserPaymentCreateInvocation) callAsBody() (*ActionUserPaymentCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserPaymentCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserPayment
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionUserPaymentCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionUserPaymentCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionUserPaymentCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionUserPaymentCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)
	input.SetUpdateIn(updateIn)

	pollResp, err := req.Call()

	if err != nil {
		return pollResp, err
	} else if pollResp.Output.Finished {
		return pollResp, nil
	}

	if callback(pollResp.Output) == StopWatching {
		return pollResp, nil
	}

	for {
		req = resp.Action.Client.ActionState.Poll.Prepare()
		req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
		req.SetInput(&ActionActionStatePollInput{
			Timeout: timeout,
			UpdateIn: updateIn,
			Status: pollResp.Output.Status,
			Current: pollResp.Output.Current,
			Total: pollResp.Output.Total,
		})
		pollResp, err = req.Call()

		if err != nil {
			return pollResp, err
		} else if pollResp.Output.Finished {
			return pollResp, nil
		}

		if callback(pollResp.Output) == StopWatching {
			return pollResp, nil
		}
	}
}

// CancelOperation cancels the current blocking operation
func (resp *ActionUserPaymentCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionUserPaymentCreateInvocation) makeAllInputParams() *ActionUserPaymentCreateRequest {
	return &ActionUserPaymentCreateRequest{
		UserPayment: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserPaymentCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Amount") {
			ret["amount"] = inv.Input.Amount
		}
		if inv.IsParameterSelected("IncomingPayment") {
			ret["incoming_payment"] = inv.Input.IncomingPayment
		}
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
	}

	return ret
}

func (inv *ActionUserPaymentCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
