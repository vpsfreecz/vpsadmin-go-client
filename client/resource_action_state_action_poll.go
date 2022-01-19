package client

import (
	"strings"
)

// ActionActionStatePoll is a type for action Action_state#Poll
type ActionActionStatePoll struct {
	// Pointer to client
	Client *Client
}

func NewActionActionStatePoll(client *Client) *ActionActionStatePoll {
	return &ActionActionStatePoll{
		Client: client,
	}
}

// ActionActionStatePollMetaGlobalInput is a type for action global meta input parameters
type ActionActionStatePollMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionActionStatePollMetaGlobalInput) SetNo(value bool) *ActionActionStatePollMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionActionStatePollMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionActionStatePollMetaGlobalInput) SelectParameters(params ...string) *ActionActionStatePollMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionActionStatePollMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionActionStatePollInput is a type for action input parameters
type ActionActionStatePollInput struct {
	Current int64 `json:"current"`
	Status bool `json:"status"`
	Timeout float64 `json:"timeout"`
	Total int64 `json:"total"`
	UpdateIn float64 `json:"update_in"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCurrent sets parameter Current to value and selects it for sending
func (in *ActionActionStatePollInput) SetCurrent(value int64) *ActionActionStatePollInput {
	in.Current = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Current"] = nil
	return in
}
// SetStatus sets parameter Status to value and selects it for sending
func (in *ActionActionStatePollInput) SetStatus(value bool) *ActionActionStatePollInput {
	in.Status = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Status"] = nil
	return in
}
// SetTimeout sets parameter Timeout to value and selects it for sending
func (in *ActionActionStatePollInput) SetTimeout(value float64) *ActionActionStatePollInput {
	in.Timeout = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Timeout"] = nil
	return in
}
// SetTotal sets parameter Total to value and selects it for sending
func (in *ActionActionStatePollInput) SetTotal(value int64) *ActionActionStatePollInput {
	in.Total = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Total"] = nil
	return in
}
// SetUpdateIn sets parameter UpdateIn to value and selects it for sending
func (in *ActionActionStatePollInput) SetUpdateIn(value float64) *ActionActionStatePollInput {
	in.UpdateIn = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UpdateIn"] = nil
	return in
}

// SelectParameters sets parameters from ActionActionStatePollInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionActionStatePollInput) SelectParameters(params ...string) *ActionActionStatePollInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionActionStatePollInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionActionStatePollOutput is a type for action output parameters
type ActionActionStatePollOutput struct {
	CanCancel bool `json:"can_cancel"`
	CreatedAt string `json:"created_at"`
	Current int64 `json:"current"`
	Finished bool `json:"finished"`
	Id int64 `json:"id"`
	Label string `json:"label"`
	Status bool `json:"status"`
	Total int64 `json:"total"`
	Unit string `json:"unit"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionActionStatePollResponse struct {
	Action *ActionActionStatePoll `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ActionState *ActionActionStatePollOutput `json:"action_state"`
	}

	// Action output without the namespace
	Output *ActionActionStatePollOutput
}


// Prepare the action for invocation
func (action *ActionActionStatePoll) Prepare() *ActionActionStatePollInvocation {
	return &ActionActionStatePollInvocation{
		Action: action,
		Path: "/v6.0/action_states/{action_state_id}/poll",
	}
}

// ActionActionStatePollInvocation is used to configure action for invocation
type ActionActionStatePollInvocation struct {
	// Pointer to the action
	Action *ActionActionStatePoll

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionActionStatePollInput
	// Global meta input parameters
	MetaInput *ActionActionStatePollMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionActionStatePollInvocation) SetPathParamInt(param string, value int64) *ActionActionStatePollInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionActionStatePollInvocation) SetPathParamString(param string, value string) *ActionActionStatePollInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionActionStatePollInvocation) NewInput() *ActionActionStatePollInput {
	inv.Input = &ActionActionStatePollInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionActionStatePollInvocation) SetInput(input *ActionActionStatePollInput) *ActionActionStatePollInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionActionStatePollInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionActionStatePollInvocation) NewMetaInput() *ActionActionStatePollMetaGlobalInput {
	inv.MetaInput = &ActionActionStatePollMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionActionStatePollInvocation) SetMetaInput(input *ActionActionStatePollMetaGlobalInput) *ActionActionStatePollInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionActionStatePollInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionActionStatePollInvocation) Call() (*ActionActionStatePollResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionActionStatePollInvocation) callAsQuery() (*ActionActionStatePollResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionActionStatePollResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ActionState
	}
	return resp, err
}



func (inv *ActionActionStatePollInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Current") {
			ret["action_state[current]"] = convertInt64ToString(inv.Input.Current)
		}
		if inv.IsParameterSelected("Status") {
			ret["action_state[status]"] = convertBoolToString(inv.Input.Status)
		}
		if inv.IsParameterSelected("Timeout") {
			ret["action_state[timeout]"] = convertFloat64ToString(inv.Input.Timeout)
		}
		if inv.IsParameterSelected("Total") {
			ret["action_state[total]"] = convertInt64ToString(inv.Input.Total)
		}
		if inv.IsParameterSelected("UpdateIn") {
			ret["action_state[update_in]"] = convertFloat64ToString(inv.Input.UpdateIn)
		}
	}
}

func (inv *ActionActionStatePollInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

