package client

import ()

// ActionOutageUpdateCreate is a type for action Outage_update#Create
type ActionOutageUpdateCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageUpdateCreate(client *Client) *ActionOutageUpdateCreate {
	return &ActionOutageUpdateCreate{
		Client: client,
	}
}

// ActionOutageUpdateCreateMetaGlobalInput is a type for action global meta input parameters
type ActionOutageUpdateCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageUpdateCreateMetaGlobalInput) SetIncludes(value string) *ActionOutageUpdateCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageUpdateCreateMetaGlobalInput) SetNo(value bool) *ActionOutageUpdateCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageUpdateCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageUpdateCreateMetaGlobalInput) SelectParameters(params ...string) *ActionOutageUpdateCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageUpdateCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageUpdateCreateInput is a type for action input parameters
type ActionOutageUpdateCreateInput struct {
	BeginsAt      string `json:"begins_at"`
	CsDescription string `json:"cs_description"`
	CsSummary     string `json:"cs_summary"`
	Duration      int64  `json:"duration"`
	EnDescription string `json:"en_description"`
	EnSummary     string `json:"en_summary"`
	FinishedAt    string `json:"finished_at"`
	Impact        string `json:"impact"`
	Outage        int64  `json:"outage"`
	SendMail      bool   `json:"send_mail"`
	State         string `json:"state"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetBeginsAt sets parameter BeginsAt to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetBeginsAt(value string) *ActionOutageUpdateCreateInput {
	in.BeginsAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["BeginsAt"] = nil
	return in
}

// SetCsDescription sets parameter CsDescription to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetCsDescription(value string) *ActionOutageUpdateCreateInput {
	in.CsDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsDescription"] = nil
	return in
}

// SetCsSummary sets parameter CsSummary to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetCsSummary(value string) *ActionOutageUpdateCreateInput {
	in.CsSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsSummary"] = nil
	return in
}

// SetDuration sets parameter Duration to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetDuration(value int64) *ActionOutageUpdateCreateInput {
	in.Duration = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Duration"] = nil
	return in
}

// SetEnDescription sets parameter EnDescription to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetEnDescription(value string) *ActionOutageUpdateCreateInput {
	in.EnDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnDescription"] = nil
	return in
}

// SetEnSummary sets parameter EnSummary to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetEnSummary(value string) *ActionOutageUpdateCreateInput {
	in.EnSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnSummary"] = nil
	return in
}

// SetFinishedAt sets parameter FinishedAt to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetFinishedAt(value string) *ActionOutageUpdateCreateInput {
	in.FinishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FinishedAt"] = nil
	return in
}

// SetImpact sets parameter Impact to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetImpact(value string) *ActionOutageUpdateCreateInput {
	in.Impact = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Impact"] = nil
	return in
}

// SetOutage sets parameter Outage to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetOutage(value int64) *ActionOutageUpdateCreateInput {
	in.Outage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetOutageNil(false)
	in._selectedParameters["Outage"] = nil
	return in
}

// SetOutageNil sets parameter Outage to nil and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetOutageNil(set bool) *ActionOutageUpdateCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Outage"] = nil
		in.SelectParameters("Outage")
	} else {
		delete(in._nilParameters, "Outage")
	}
	return in
}

// SetSendMail sets parameter SendMail to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetSendMail(value bool) *ActionOutageUpdateCreateInput {
	in.SendMail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SendMail"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionOutageUpdateCreateInput) SetState(value string) *ActionOutageUpdateCreateInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageUpdateCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageUpdateCreateInput) SelectParameters(params ...string) *ActionOutageUpdateCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOutageUpdateCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOutageUpdateCreateInput) UnselectParameters(params ...string) *ActionOutageUpdateCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOutageUpdateCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageUpdateCreateRequest is a type for the entire action request
type ActionOutageUpdateCreateRequest struct {
	OutageUpdate map[string]interface{} `json:"outage_update"`
	Meta         map[string]interface{} `json:"_meta"`
}

// ActionOutageUpdateCreateOutput is a type for action output parameters
type ActionOutageUpdateCreateOutput struct {
	BeginsAt      string                  `json:"begins_at"`
	CreatedAt     string                  `json:"created_at"`
	CsDescription string                  `json:"cs_description"`
	CsSummary     string                  `json:"cs_summary"`
	Duration      int64                   `json:"duration"`
	EnDescription string                  `json:"en_description"`
	EnSummary     string                  `json:"en_summary"`
	FinishedAt    string                  `json:"finished_at"`
	Id            int64                   `json:"id"`
	Impact        string                  `json:"impact"`
	Outage        *ActionOutageShowOutput `json:"outage"`
	ReportedBy    *ActionUserShowOutput   `json:"reported_by"`
	ReporterName  string                  `json:"reporter_name"`
	State         string                  `json:"state"`
	Type          string                  `json:"type"`
}

// ActionOutageUpdateCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionOutageUpdateCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionOutageUpdateCreateResponse struct {
	Action *ActionOutageUpdateCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageUpdate *ActionOutageUpdateCreateOutput `json:"outage_update"`
		// Global output metadata
		Meta *ActionOutageUpdateCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionOutageUpdateCreateOutput
}

// Prepare the action for invocation
func (action *ActionOutageUpdateCreate) Prepare() *ActionOutageUpdateCreateInvocation {
	return &ActionOutageUpdateCreateInvocation{
		Action: action,
		Path:   "/v7.0/outage_updates",
	}
}

// ActionOutageUpdateCreateInvocation is used to configure action for invocation
type ActionOutageUpdateCreateInvocation struct {
	// Pointer to the action
	Action *ActionOutageUpdateCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageUpdateCreateInput
	// Global meta input parameters
	MetaInput *ActionOutageUpdateCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageUpdateCreateInvocation) NewInput() *ActionOutageUpdateCreateInput {
	inv.Input = &ActionOutageUpdateCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageUpdateCreateInvocation) SetInput(input *ActionOutageUpdateCreateInput) *ActionOutageUpdateCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageUpdateCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOutageUpdateCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageUpdateCreateInvocation) NewMetaInput() *ActionOutageUpdateCreateMetaGlobalInput {
	inv.MetaInput = &ActionOutageUpdateCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageUpdateCreateInvocation) SetMetaInput(input *ActionOutageUpdateCreateMetaGlobalInput) *ActionOutageUpdateCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageUpdateCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageUpdateCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageUpdateCreateInvocation) Call() (*ActionOutageUpdateCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOutageUpdateCreateInvocation) callAsBody() (*ActionOutageUpdateCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageUpdateCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageUpdate
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionOutageUpdateCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionOutageUpdateCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionOutageUpdateCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionOutageUpdateCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
			Timeout:  timeout,
			UpdateIn: updateIn,
			Status:   pollResp.Output.Status,
			Current:  pollResp.Output.Current,
			Total:    pollResp.Output.Total,
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
func (resp *ActionOutageUpdateCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionOutageUpdateCreateInvocation) makeAllInputParams() *ActionOutageUpdateCreateRequest {
	return &ActionOutageUpdateCreateRequest{
		OutageUpdate: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionOutageUpdateCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("BeginsAt") {
			ret["begins_at"] = inv.Input.BeginsAt
		}
		if inv.IsParameterSelected("CsDescription") {
			ret["cs_description"] = inv.Input.CsDescription
		}
		if inv.IsParameterSelected("CsSummary") {
			ret["cs_summary"] = inv.Input.CsSummary
		}
		if inv.IsParameterSelected("Duration") {
			ret["duration"] = inv.Input.Duration
		}
		if inv.IsParameterSelected("EnDescription") {
			ret["en_description"] = inv.Input.EnDescription
		}
		if inv.IsParameterSelected("EnSummary") {
			ret["en_summary"] = inv.Input.EnSummary
		}
		if inv.IsParameterSelected("FinishedAt") {
			ret["finished_at"] = inv.Input.FinishedAt
		}
		if inv.IsParameterSelected("Impact") {
			ret["impact"] = inv.Input.Impact
		}
		if inv.IsParameterSelected("Outage") {
			if inv.IsParameterNil("Outage") {
				ret["outage"] = nil
			} else {
				ret["outage"] = inv.Input.Outage
			}
		}
		if inv.IsParameterSelected("SendMail") {
			ret["send_mail"] = inv.Input.SendMail
		}
		if inv.IsParameterSelected("State") {
			ret["state"] = inv.Input.State
		}
	}

	return ret
}

func (inv *ActionOutageUpdateCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
