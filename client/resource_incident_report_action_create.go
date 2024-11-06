package client

import ()

// ActionIncidentReportCreate is a type for action Incident_report#Create
type ActionIncidentReportCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentReportCreate(client *Client) *ActionIncidentReportCreate {
	return &ActionIncidentReportCreate{
		Client: client,
	}
}

// ActionIncidentReportCreateMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentReportCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIncidentReportCreateMetaGlobalInput) SetIncludes(value string) *ActionIncidentReportCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentReportCreateMetaGlobalInput) SetNo(value bool) *ActionIncidentReportCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentReportCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportCreateMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentReportCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentReportCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportCreateInput is a type for action input parameters
type ActionIncidentReportCreateInput struct {
	Codename            string `json:"codename"`
	CpuLimit            int64  `json:"cpu_limit"`
	DetectedAt          string `json:"detected_at"`
	IpAddressAssignment int64  `json:"ip_address_assignment"`
	Subject             string `json:"subject"`
	Text                string `json:"text"`
	Vps                 int64  `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCodename sets parameter Codename to value and selects it for sending
func (in *ActionIncidentReportCreateInput) SetCodename(value string) *ActionIncidentReportCreateInput {
	in.Codename = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Codename"] = nil
	return in
}

// SetCpuLimit sets parameter CpuLimit to value and selects it for sending
func (in *ActionIncidentReportCreateInput) SetCpuLimit(value int64) *ActionIncidentReportCreateInput {
	in.CpuLimit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CpuLimit"] = nil
	return in
}

// SetDetectedAt sets parameter DetectedAt to value and selects it for sending
func (in *ActionIncidentReportCreateInput) SetDetectedAt(value string) *ActionIncidentReportCreateInput {
	in.DetectedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DetectedAt"] = nil
	return in
}

// SetIpAddressAssignment sets parameter IpAddressAssignment to value and selects it for sending
func (in *ActionIncidentReportCreateInput) SetIpAddressAssignment(value int64) *ActionIncidentReportCreateInput {
	in.IpAddressAssignment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetIpAddressAssignmentNil(false)
	in._selectedParameters["IpAddressAssignment"] = nil
	return in
}

// SetIpAddressAssignmentNil sets parameter IpAddressAssignment to nil and selects it for sending
func (in *ActionIncidentReportCreateInput) SetIpAddressAssignmentNil(set bool) *ActionIncidentReportCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["IpAddressAssignment"] = nil
		in.SelectParameters("IpAddressAssignment")
	} else {
		delete(in._nilParameters, "IpAddressAssignment")
	}
	return in
}

// SetSubject sets parameter Subject to value and selects it for sending
func (in *ActionIncidentReportCreateInput) SetSubject(value string) *ActionIncidentReportCreateInput {
	in.Subject = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Subject"] = nil
	return in
}

// SetText sets parameter Text to value and selects it for sending
func (in *ActionIncidentReportCreateInput) SetText(value string) *ActionIncidentReportCreateInput {
	in.Text = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Text"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionIncidentReportCreateInput) SetVps(value int64) *ActionIncidentReportCreateInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionIncidentReportCreateInput) SetVpsNil(set bool) *ActionIncidentReportCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Vps"] = nil
		in.SelectParameters("Vps")
	} else {
		delete(in._nilParameters, "Vps")
	}
	return in
}

// SelectParameters sets parameters from ActionIncidentReportCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportCreateInput) SelectParameters(params ...string) *ActionIncidentReportCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIncidentReportCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIncidentReportCreateInput) UnselectParameters(params ...string) *ActionIncidentReportCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIncidentReportCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportCreateRequest is a type for the entire action request
type ActionIncidentReportCreateRequest struct {
	IncidentReport map[string]interface{} `json:"incident_report"`
	Meta           map[string]interface{} `json:"_meta"`
}

// ActionIncidentReportCreateOutput is a type for action output parameters
type ActionIncidentReportCreateOutput struct {
	Codename            string                               `json:"codename"`
	CpuLimit            int64                                `json:"cpu_limit"`
	CreatedAt           string                               `json:"created_at"`
	DetectedAt          string                               `json:"detected_at"`
	FiledBy             *ActionUserShowOutput                `json:"filed_by"`
	Id                  int64                                `json:"id"`
	IpAddressAssignment *ActionIpAddressAssignmentShowOutput `json:"ip_address_assignment"`
	Mailbox             *ActionMailboxShowOutput             `json:"mailbox"`
	RawUserId           int64                                `json:"raw_user_id"`
	RawVpsId            int64                                `json:"raw_vps_id"`
	ReportedAt          string                               `json:"reported_at"`
	Subject             string                               `json:"subject"`
	Text                string                               `json:"text"`
	User                *ActionUserShowOutput                `json:"user"`
	Vps                 *ActionVpsShowOutput                 `json:"vps"`
}

// ActionIncidentReportCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionIncidentReportCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionIncidentReportCreateResponse struct {
	Action *ActionIncidentReportCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncidentReport *ActionIncidentReportCreateOutput `json:"incident_report"`
		// Global output metadata
		Meta *ActionIncidentReportCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionIncidentReportCreateOutput
}

// Prepare the action for invocation
func (action *ActionIncidentReportCreate) Prepare() *ActionIncidentReportCreateInvocation {
	return &ActionIncidentReportCreateInvocation{
		Action: action,
		Path:   "/v7.0/incident_reports",
	}
}

// ActionIncidentReportCreateInvocation is used to configure action for invocation
type ActionIncidentReportCreateInvocation struct {
	// Pointer to the action
	Action *ActionIncidentReportCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIncidentReportCreateInput
	// Global meta input parameters
	MetaInput *ActionIncidentReportCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIncidentReportCreateInvocation) NewInput() *ActionIncidentReportCreateInput {
	inv.Input = &ActionIncidentReportCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIncidentReportCreateInvocation) SetInput(input *ActionIncidentReportCreateInput) *ActionIncidentReportCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIncidentReportCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIncidentReportCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentReportCreateInvocation) NewMetaInput() *ActionIncidentReportCreateMetaGlobalInput {
	inv.MetaInput = &ActionIncidentReportCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentReportCreateInvocation) SetMetaInput(input *ActionIncidentReportCreateMetaGlobalInput) *ActionIncidentReportCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentReportCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentReportCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncidentReportCreateInvocation) Call() (*ActionIncidentReportCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionIncidentReportCreateInvocation) callAsBody() (*ActionIncidentReportCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIncidentReportCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncidentReport
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionIncidentReportCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionIncidentReportCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionIncidentReportCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionIncidentReportCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionIncidentReportCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionIncidentReportCreateInvocation) makeAllInputParams() *ActionIncidentReportCreateRequest {
	return &ActionIncidentReportCreateRequest{
		IncidentReport: inv.makeInputParams(),
		Meta:           inv.makeMetaInputParams(),
	}
}

func (inv *ActionIncidentReportCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Codename") {
			ret["codename"] = inv.Input.Codename
		}
		if inv.IsParameterSelected("CpuLimit") {
			ret["cpu_limit"] = inv.Input.CpuLimit
		}
		if inv.IsParameterSelected("DetectedAt") {
			ret["detected_at"] = inv.Input.DetectedAt
		}
		if inv.IsParameterSelected("IpAddressAssignment") {
			if inv.IsParameterNil("IpAddressAssignment") {
				ret["ip_address_assignment"] = nil
			} else {
				ret["ip_address_assignment"] = inv.Input.IpAddressAssignment
			}
		}
		if inv.IsParameterSelected("Subject") {
			ret["subject"] = inv.Input.Subject
		}
		if inv.IsParameterSelected("Text") {
			ret["text"] = inv.Input.Text
		}
		if inv.IsParameterSelected("Vps") {
			if inv.IsParameterNil("Vps") {
				ret["vps"] = nil
			} else {
				ret["vps"] = inv.Input.Vps
			}
		}
	}

	return ret
}

func (inv *ActionIncidentReportCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
