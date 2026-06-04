package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryPublish is a type for action Security_advisory#Publish
type ActionSecurityAdvisoryPublish struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryPublish(client *Client) *ActionSecurityAdvisoryPublish {
	return &ActionSecurityAdvisoryPublish{
		Client: client,
	}
}

// ActionSecurityAdvisoryPublishMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryPublishMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryPublishMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryPublishMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryPublishMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryPublishMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryPublishMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryPublishMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryPublishMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryPublishMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryPublishInput is a type for action input parameters
type ActionSecurityAdvisoryPublishInput struct {
	PublishedAt string "json:\"published_at\""
	SendMail    bool   "json:\"send_mail\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetPublishedAt sets parameter PublishedAt to value and selects it for sending
func (in *ActionSecurityAdvisoryPublishInput) SetPublishedAt(value string) *ActionSecurityAdvisoryPublishInput {
	in.PublishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetPublishedAtNil(false)
	in._selectedParameters["PublishedAt"] = nil
	return in
}

// SetPublishedAtNil sets parameter PublishedAt to nil and selects it for sending
func (in *ActionSecurityAdvisoryPublishInput) SetPublishedAtNil(set bool) *ActionSecurityAdvisoryPublishInput {
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

// SetSendMail sets parameter SendMail to value and selects it for sending
func (in *ActionSecurityAdvisoryPublishInput) SetSendMail(value bool) *ActionSecurityAdvisoryPublishInput {
	in.SendMail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SendMail"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryPublishInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryPublishInput) SelectParameters(params ...string) *ActionSecurityAdvisoryPublishInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryPublishInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryPublishInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryPublishInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryPublishInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryPublishRequest is a type for the entire action request
type ActionSecurityAdvisoryPublishRequest struct {
	SecurityAdvisory map[string]interface{} "json:\"security_advisory\""
	Meta             map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryPublishOutput is a type for action output parameters
type ActionSecurityAdvisoryPublishOutput struct {
	Affected          bool                  "json:\"affected\""
	AffectedNodeCount int64                 "json:\"affected_node_count\""
	AffectedUserCount int64                 "json:\"affected_user_count\""
	AffectedVpsCount  int64                 "json:\"affected_vps_count\""
	CreatedAt         string                "json:\"created_at\""
	CreatedBy         *ActionUserShowOutput "json:\"created_by\""
	CsDescription     string                "json:\"cs_description\""
	CsResponse        string                "json:\"cs_response\""
	CsSummary         string                "json:\"cs_summary\""
	EnDescription     string                "json:\"en_description\""
	EnResponse        string                "json:\"en_response\""
	EnSummary         string                "json:\"en_summary\""
	Id                int64                 "json:\"id\""
	Name              string                "json:\"name\""
	PublishedAt       string                "json:\"published_at\""
	PublishedBy       *ActionUserShowOutput "json:\"published_by\""
	RetractedAt       string                "json:\"retracted_at\""
	State             string                "json:\"state\""
	UpdatedAt         string                "json:\"updated_at\""
}

// ActionSecurityAdvisoryPublishMetaGlobalOutput is a type for global output metadata parameters
type ActionSecurityAdvisoryPublishMetaGlobalOutput struct {
	ActionStateId int64 "json:\"action_state_id\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryPublishResponse struct {
	Action *ActionSecurityAdvisoryPublish "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisory *ActionSecurityAdvisoryPublishOutput "json:\"security_advisory\""
		// Global output metadata
		Meta *ActionSecurityAdvisoryPublishMetaGlobalOutput "json:\"_meta\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryPublishOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryPublish) Prepare() *ActionSecurityAdvisoryPublishInvocation {
	return &ActionSecurityAdvisoryPublishInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}/publish",
	}
}

// ActionSecurityAdvisoryPublishInvocation is used to configure action for invocation
type ActionSecurityAdvisoryPublishInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryPublish

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryPublishInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryPublishMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryPublishInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryPublishInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryPublishInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryPublishInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryPublishInvocation) NewInput() *ActionSecurityAdvisoryPublishInput {
	inv.Input = &ActionSecurityAdvisoryPublishInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryPublishInvocation) SetInput(input *ActionSecurityAdvisoryPublishInput) *ActionSecurityAdvisoryPublishInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryPublishInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryPublishInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryPublishInvocation) NewMetaInput() *ActionSecurityAdvisoryPublishMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryPublishMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryPublishInvocation) SetMetaInput(input *ActionSecurityAdvisoryPublishMetaGlobalInput) *ActionSecurityAdvisoryPublishInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryPublishInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryPublishInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryPublishInvocation) validate() error {
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
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSecurityAdvisoryPublishInvocation) Call() (*ActionSecurityAdvisoryPublishResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryPublishInvocation) callAsBody() (*ActionSecurityAdvisoryPublishResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryPublishResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisory
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionSecurityAdvisoryPublishResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionSecurityAdvisoryPublishResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionSecurityAdvisoryPublishResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionSecurityAdvisoryPublishResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionSecurityAdvisoryPublishResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionSecurityAdvisoryPublishInvocation) makeAllInputParams() *ActionSecurityAdvisoryPublishRequest {
	return &ActionSecurityAdvisoryPublishRequest{
		SecurityAdvisory: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryPublishInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("PublishedAt") {
			if inv.IsParameterNil("PublishedAt") {
				ret["published_at"] = nil
			} else {
				ret["published_at"] = inv.Input.PublishedAt
			}
		}
		if inv.IsParameterSelected("SendMail") {
			ret["send_mail"] = inv.Input.SendMail
		}
	}

	return ret
}

func (inv *ActionSecurityAdvisoryPublishInvocation) makeMetaInputParams() map[string]interface{} {
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
