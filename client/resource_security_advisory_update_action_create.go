package client

import ()

// ActionSecurityAdvisoryUpdateCreate is a type for action Security_advisory_update#Create
type ActionSecurityAdvisoryUpdateCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryUpdateCreate(client *Client) *ActionSecurityAdvisoryUpdateCreate {
	return &ActionSecurityAdvisoryUpdateCreate{
		Client: client,
	}
}

// ActionSecurityAdvisoryUpdateCreateMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryUpdateCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateCreateInput is a type for action input parameters
type ActionSecurityAdvisoryUpdateCreateInput struct {
	CsMessage        string "json:\"cs_message\""
	CsSummary        string "json:\"cs_summary\""
	EnMessage        string "json:\"en_message\""
	EnSummary        string "json:\"en_summary\""
	PublishedAt      string "json:\"published_at\""
	SecurityAdvisory int64  "json:\"security_advisory\""
	SendMail         bool   "json:\"send_mail\""
	State            string "json:\"state\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCsMessage sets parameter CsMessage to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetCsMessage(value string) *ActionSecurityAdvisoryUpdateCreateInput {
	in.CsMessage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetCsMessageNil(false)
	in._selectedParameters["CsMessage"] = nil
	return in
}

// SetCsMessageNil sets parameter CsMessage to nil and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetCsMessageNil(set bool) *ActionSecurityAdvisoryUpdateCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["CsMessage"] = nil
		in.SelectParameters("CsMessage")
	} else {
		delete(in._nilParameters, "CsMessage")
	}
	return in
}

// SetCsSummary sets parameter CsSummary to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetCsSummary(value string) *ActionSecurityAdvisoryUpdateCreateInput {
	in.CsSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsSummary"] = nil
	return in
}

// SetEnMessage sets parameter EnMessage to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetEnMessage(value string) *ActionSecurityAdvisoryUpdateCreateInput {
	in.EnMessage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnMessageNil(false)
	in._selectedParameters["EnMessage"] = nil
	return in
}

// SetEnMessageNil sets parameter EnMessage to nil and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetEnMessageNil(set bool) *ActionSecurityAdvisoryUpdateCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["EnMessage"] = nil
		in.SelectParameters("EnMessage")
	} else {
		delete(in._nilParameters, "EnMessage")
	}
	return in
}

// SetEnSummary sets parameter EnSummary to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetEnSummary(value string) *ActionSecurityAdvisoryUpdateCreateInput {
	in.EnSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnSummary"] = nil
	return in
}

// SetPublishedAt sets parameter PublishedAt to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetPublishedAt(value string) *ActionSecurityAdvisoryUpdateCreateInput {
	in.PublishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetPublishedAtNil(false)
	in._selectedParameters["PublishedAt"] = nil
	return in
}

// SetPublishedAtNil sets parameter PublishedAt to nil and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetPublishedAtNil(set bool) *ActionSecurityAdvisoryUpdateCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["PublishedAt"] = nil
		in.SelectParameters("PublishedAt")
	} else {
		delete(in._nilParameters, "PublishedAt")
	}
	return in
}

// SetSecurityAdvisory sets parameter SecurityAdvisory to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetSecurityAdvisory(value int64) *ActionSecurityAdvisoryUpdateCreateInput {
	in.SecurityAdvisory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SecurityAdvisory"] = nil
	return in
}

// SetSendMail sets parameter SendMail to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetSendMail(value bool) *ActionSecurityAdvisoryUpdateCreateInput {
	in.SendMail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SendMail"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetState(value string) *ActionSecurityAdvisoryUpdateCreateInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetStateNil(false)
	in._selectedParameters["State"] = nil
	return in
}

// SetStateNil sets parameter State to nil and selects it for sending
func (in *ActionSecurityAdvisoryUpdateCreateInput) SetStateNil(set bool) *ActionSecurityAdvisoryUpdateCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["State"] = nil
		in.SelectParameters("State")
	} else {
		delete(in._nilParameters, "State")
	}
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateCreateInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryUpdateCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateCreateInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryUpdateCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateCreateRequest is a type for the entire action request
type ActionSecurityAdvisoryUpdateCreateRequest struct {
	SecurityAdvisoryUpdate map[string]interface{} "json:\"security_advisory_update\""
	Meta                   map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryUpdateCreateOutput is a type for action output parameters
type ActionSecurityAdvisoryUpdateCreateOutput struct {
	CreatedAt        string                            "json:\"created_at\""
	CsMessage        string                            "json:\"cs_message\""
	CsSummary        string                            "json:\"cs_summary\""
	EnMessage        string                            "json:\"en_message\""
	EnSummary        string                            "json:\"en_summary\""
	Id               int64                             "json:\"id\""
	Name             string                            "json:\"name\""
	ReportedBy       *ActionUserShowOutput             "json:\"reported_by\""
	ReporterName     string                            "json:\"reporter_name\""
	SecurityAdvisory *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	State            string                            "json:\"state\""
	UpdatedAt        string                            "json:\"updated_at\""
}

// ActionSecurityAdvisoryUpdateCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionSecurityAdvisoryUpdateCreateMetaGlobalOutput struct {
	ActionStateId int64 "json:\"action_state_id\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryUpdateCreateResponse struct {
	Action *ActionSecurityAdvisoryUpdateCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisoryUpdate *ActionSecurityAdvisoryUpdateCreateOutput "json:\"security_advisory_update\""
		// Global output metadata
		Meta *ActionSecurityAdvisoryUpdateCreateMetaGlobalOutput "json:\"_meta\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryUpdateCreateOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryUpdateCreate) Prepare() *ActionSecurityAdvisoryUpdateCreateInvocation {
	return &ActionSecurityAdvisoryUpdateCreateInvocation{
		Action: action,
		Path:   "/v7.0/security_advisory_updates",
	}
}

// ActionSecurityAdvisoryUpdateCreateInvocation is used to configure action for invocation
type ActionSecurityAdvisoryUpdateCreateInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryUpdateCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryUpdateCreateInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) NewInput() *ActionSecurityAdvisoryUpdateCreateInput {
	inv.Input = &ActionSecurityAdvisoryUpdateCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) SetInput(input *ActionSecurityAdvisoryUpdateCreateInput) *ActionSecurityAdvisoryUpdateCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) NewMetaInput() *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryUpdateCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) SetMetaInput(input *ActionSecurityAdvisoryUpdateCreateMetaGlobalInput) *ActionSecurityAdvisoryUpdateCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("PublishedAt") {
			if !inv.IsParameterNil("PublishedAt") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.PublishedAt)
				if !ok {
					verr.Add("published_at", "not a valid datetime")
				} else {
					inv.Input.PublishedAt = normalized
				}
			}
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			if !inv.IsParameterNil("SecurityAdvisory") {
				if inv.Input.SecurityAdvisory < 0 {
					verr.Add("security_advisory", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) Call() (*ActionSecurityAdvisoryUpdateCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) callAsBody() (*ActionSecurityAdvisoryUpdateCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryUpdateCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisoryUpdate
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionSecurityAdvisoryUpdateCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionSecurityAdvisoryUpdateCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionSecurityAdvisoryUpdateCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionSecurityAdvisoryUpdateCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionSecurityAdvisoryUpdateCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) makeAllInputParams() *ActionSecurityAdvisoryUpdateCreateRequest {
	return &ActionSecurityAdvisoryUpdateCreateRequest{
		SecurityAdvisoryUpdate: inv.makeInputParams(),
		Meta:                   inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("CsMessage") {
			if inv.IsParameterNil("CsMessage") {
				ret["cs_message"] = nil
			} else {
				ret["cs_message"] = inv.Input.CsMessage
			}
		}
		if inv.IsParameterSelected("CsSummary") {
			ret["cs_summary"] = inv.Input.CsSummary
		}
		if inv.IsParameterSelected("EnMessage") {
			if inv.IsParameterNil("EnMessage") {
				ret["en_message"] = nil
			} else {
				ret["en_message"] = inv.Input.EnMessage
			}
		}
		if inv.IsParameterSelected("EnSummary") {
			ret["en_summary"] = inv.Input.EnSummary
		}
		if inv.IsParameterSelected("PublishedAt") {
			if inv.IsParameterNil("PublishedAt") {
				ret["published_at"] = nil
			} else {
				ret["published_at"] = inv.Input.PublishedAt
			}
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			ret["security_advisory"] = inv.Input.SecurityAdvisory
		}
		if inv.IsParameterSelected("SendMail") {
			ret["send_mail"] = inv.Input.SendMail
		}
		if inv.IsParameterSelected("State") {
			if inv.IsParameterNil("State") {
				ret["state"] = nil
			} else {
				ret["state"] = inv.Input.State
			}
		}
	}

	return ret
}

func (inv *ActionSecurityAdvisoryUpdateCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
