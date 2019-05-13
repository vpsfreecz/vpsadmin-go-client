package client

import (
	"strings"
)

// ActionOutageUpdate is a type for action Outage#Update
type ActionOutageUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageUpdate(client *Client) *ActionOutageUpdate {
	return &ActionOutageUpdate{
		Client: client,
	}
}

// ActionOutageUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionOutageUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageUpdateMetaGlobalInput) SetNo(value bool) *ActionOutageUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageUpdateMetaGlobalInput) SetIncludes(value string) *ActionOutageUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionOutageUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageUpdateInput is a type for action input parameters
type ActionOutageUpdateInput struct {
	BeginsAt string `json:"begins_at"`
	FinishedAt string `json:"finished_at"`
	Duration int64 `json:"duration"`
	State string `json:"state"`
	Type string `json:"type"`
	EnSummary string `json:"en_summary"`
	EnDescription string `json:"en_description"`
	CsSummary string `json:"cs_summary"`
	CsDescription string `json:"cs_description"`
	SendMail bool `json:"send_mail"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetBeginsAt sets parameter BeginsAt to value and selects it for sending
func (in *ActionOutageUpdateInput) SetBeginsAt(value string) *ActionOutageUpdateInput {
	in.BeginsAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["BeginsAt"] = nil
	return in
}
// SetFinishedAt sets parameter FinishedAt to value and selects it for sending
func (in *ActionOutageUpdateInput) SetFinishedAt(value string) *ActionOutageUpdateInput {
	in.FinishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FinishedAt"] = nil
	return in
}
// SetDuration sets parameter Duration to value and selects it for sending
func (in *ActionOutageUpdateInput) SetDuration(value int64) *ActionOutageUpdateInput {
	in.Duration = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Duration"] = nil
	return in
}
// SetState sets parameter State to value and selects it for sending
func (in *ActionOutageUpdateInput) SetState(value string) *ActionOutageUpdateInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}
// SetType sets parameter Type to value and selects it for sending
func (in *ActionOutageUpdateInput) SetType(value string) *ActionOutageUpdateInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}
// SetEnSummary sets parameter EnSummary to value and selects it for sending
func (in *ActionOutageUpdateInput) SetEnSummary(value string) *ActionOutageUpdateInput {
	in.EnSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnSummary"] = nil
	return in
}
// SetEnDescription sets parameter EnDescription to value and selects it for sending
func (in *ActionOutageUpdateInput) SetEnDescription(value string) *ActionOutageUpdateInput {
	in.EnDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnDescription"] = nil
	return in
}
// SetCsSummary sets parameter CsSummary to value and selects it for sending
func (in *ActionOutageUpdateInput) SetCsSummary(value string) *ActionOutageUpdateInput {
	in.CsSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsSummary"] = nil
	return in
}
// SetCsDescription sets parameter CsDescription to value and selects it for sending
func (in *ActionOutageUpdateInput) SetCsDescription(value string) *ActionOutageUpdateInput {
	in.CsDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsDescription"] = nil
	return in
}
// SetSendMail sets parameter SendMail to value and selects it for sending
func (in *ActionOutageUpdateInput) SetSendMail(value bool) *ActionOutageUpdateInput {
	in.SendMail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SendMail"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageUpdateInput) SelectParameters(params ...string) *ActionOutageUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageUpdateRequest is a type for the entire action request
type ActionOutageUpdateRequest struct {
	Outage map[string]interface{} `json:"outage"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionOutageUpdateOutput is a type for action output parameters
type ActionOutageUpdateOutput struct {
	Id int64 `json:"id"`
	BeginsAt string `json:"begins_at"`
	FinishedAt string `json:"finished_at"`
	Duration int64 `json:"duration"`
	Planned bool `json:"planned"`
	State string `json:"state"`
	Type string `json:"type"`
	EnSummary string `json:"en_summary"`
	EnDescription string `json:"en_description"`
	CsSummary string `json:"cs_summary"`
	CsDescription string `json:"cs_description"`
	Affected bool `json:"affected"`
	AffectedUserCount int64 `json:"affected_user_count"`
	AffectedDirectVpsCount int64 `json:"affected_direct_vps_count"`
	AffectedIndirectVpsCount int64 `json:"affected_indirect_vps_count"`
}

// ActionOutageUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionOutageUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionOutageUpdateResponse struct {
	Action *ActionOutageUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Outage *ActionOutageUpdateOutput `json:"outage"`
		// Global output metadata
		Meta *ActionOutageUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionOutageUpdateOutput
}


// Prepare the action for invocation
func (action *ActionOutageUpdate) Prepare() *ActionOutageUpdateInvocation {
	return &ActionOutageUpdateInvocation{
		Action: action,
		Path: "/v5.0/outages/{outage_id}",
	}
}

// ActionOutageUpdateInvocation is used to configure action for invocation
type ActionOutageUpdateInvocation struct {
	// Pointer to the action
	Action *ActionOutageUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageUpdateInput
	// Global meta input parameters
	MetaInput *ActionOutageUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageUpdateInvocation) SetPathParamInt(param string, value int64) *ActionOutageUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageUpdateInvocation) SetPathParamString(param string, value string) *ActionOutageUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageUpdateInvocation) NewInput() *ActionOutageUpdateInput {
	inv.Input = &ActionOutageUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageUpdateInvocation) SetInput(input *ActionOutageUpdateInput) *ActionOutageUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageUpdateInvocation) NewMetaInput() *ActionOutageUpdateMetaGlobalInput {
	inv.MetaInput = &ActionOutageUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageUpdateInvocation) SetMetaInput(input *ActionOutageUpdateMetaGlobalInput) *ActionOutageUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageUpdateInvocation) Call() (*ActionOutageUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionOutageUpdateInvocation) callAsBody() (*ActionOutageUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Outage
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionOutageUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionOutageUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionOutageUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionOutageUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionOutageUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionOutageUpdateInvocation) makeAllInputParams() *ActionOutageUpdateRequest {
	return &ActionOutageUpdateRequest{
		Outage: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOutageUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("BeginsAt") {
			ret["begins_at"] = inv.Input.BeginsAt
		}
		if inv.IsParameterSelected("FinishedAt") {
			ret["finished_at"] = inv.Input.FinishedAt
		}
		if inv.IsParameterSelected("Duration") {
			ret["duration"] = inv.Input.Duration
		}
		if inv.IsParameterSelected("State") {
			ret["state"] = inv.Input.State
		}
		if inv.IsParameterSelected("Type") {
			ret["type"] = inv.Input.Type
		}
		if inv.IsParameterSelected("EnSummary") {
			ret["en_summary"] = inv.Input.EnSummary
		}
		if inv.IsParameterSelected("EnDescription") {
			ret["en_description"] = inv.Input.EnDescription
		}
		if inv.IsParameterSelected("CsSummary") {
			ret["cs_summary"] = inv.Input.CsSummary
		}
		if inv.IsParameterSelected("CsDescription") {
			ret["cs_description"] = inv.Input.CsDescription
		}
		if inv.IsParameterSelected("SendMail") {
			ret["send_mail"] = inv.Input.SendMail
		}
	}

	return ret
}

func (inv *ActionOutageUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
