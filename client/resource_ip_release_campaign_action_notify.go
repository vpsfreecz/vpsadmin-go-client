package client

import (
	"net/url"
	"strings"
)

// ActionIpReleaseCampaignNotify is a type for action Ip_release_campaign#Notify
type ActionIpReleaseCampaignNotify struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignNotify(client *Client) *ActionIpReleaseCampaignNotify {
	return &ActionIpReleaseCampaignNotify{
		Client: client,
	}
}

// ActionIpReleaseCampaignNotifyMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignNotifyMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignNotifyMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignNotifyMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignNotifyMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignNotifyMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignNotifyMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignNotifyMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignNotifyMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignNotifyMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignNotifyInput is a type for action input parameters
type ActionIpReleaseCampaignNotifyInput struct {
	Event string "json:\"event\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEvent sets parameter Event to value and selects it for sending
func (in *ActionIpReleaseCampaignNotifyInput) SetEvent(value string) *ActionIpReleaseCampaignNotifyInput {
	in.Event = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Event"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignNotifyInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignNotifyInput) SelectParameters(params ...string) *ActionIpReleaseCampaignNotifyInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseCampaignNotifyInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignNotifyInput) UnselectParameters(params ...string) *ActionIpReleaseCampaignNotifyInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseCampaignNotifyInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignNotifyRequest is a type for the entire action request
type ActionIpReleaseCampaignNotifyRequest struct {
	IpReleaseCampaign map[string]interface{} "json:\"ip_release_campaign\""
	Meta              map[string]interface{} "json:\"_meta\""
}

// ActionIpReleaseCampaignNotifyOutput is a type for action output parameters
type ActionIpReleaseCampaignNotifyOutput struct {
	AllowKeep bool                  "json:\"allow_keep\""
	ClosedAt  string                "json:\"closed_at\""
	ClosedBy  *ActionUserShowOutput "json:\"closed_by\""
	CreatedAt string                "json:\"created_at\""
	CreatedBy *ActionUserShowOutput "json:\"created_by\""
	Deadline  string                "json:\"deadline\""
	Id        int64                 "json:\"id\""
}

// ActionIpReleaseCampaignNotifyMetaGlobalOutput is a type for global output metadata parameters
type ActionIpReleaseCampaignNotifyMetaGlobalOutput struct {
	ActionStateId int64 "json:\"action_state_id\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignNotifyResponse struct {
	Action *ActionIpReleaseCampaignNotify "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseCampaign *ActionIpReleaseCampaignNotifyOutput "json:\"ip_release_campaign\""
		// Global output metadata
		Meta *ActionIpReleaseCampaignNotifyMetaGlobalOutput "json:\"_meta\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseCampaignNotifyOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignNotify) Prepare() *ActionIpReleaseCampaignNotifyInvocation {
	return &ActionIpReleaseCampaignNotifyInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/{ip_release_campaign_id}/notify",
	}
}

// ActionIpReleaseCampaignNotifyInvocation is used to configure action for invocation
type ActionIpReleaseCampaignNotifyInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignNotify

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseCampaignNotifyInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignNotifyMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseCampaignNotifyInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseCampaignNotifyInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseCampaignNotifyInvocation) SetPathParamString(param string, value string) *ActionIpReleaseCampaignNotifyInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseCampaignNotifyInvocation) NewInput() *ActionIpReleaseCampaignNotifyInput {
	inv.Input = &ActionIpReleaseCampaignNotifyInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseCampaignNotifyInvocation) SetInput(input *ActionIpReleaseCampaignNotifyInput) *ActionIpReleaseCampaignNotifyInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseCampaignNotifyInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignNotifyInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignNotifyInvocation) NewMetaInput() *ActionIpReleaseCampaignNotifyMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignNotifyMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignNotifyInvocation) SetMetaInput(input *ActionIpReleaseCampaignNotifyMetaGlobalInput) *ActionIpReleaseCampaignNotifyInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignNotifyInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignNotifyInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignNotifyInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpReleaseCampaignNotifyInvocation) Call() (*ActionIpReleaseCampaignNotifyResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionIpReleaseCampaignNotifyInvocation) callAsBody() (*ActionIpReleaseCampaignNotifyResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIpReleaseCampaignNotifyResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpReleaseCampaign
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionIpReleaseCampaignNotifyResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionIpReleaseCampaignNotifyResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionIpReleaseCampaignNotifyResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionIpReleaseCampaignNotifyResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionIpReleaseCampaignNotifyResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionIpReleaseCampaignNotifyInvocation) makeAllInputParams() *ActionIpReleaseCampaignNotifyRequest {
	return &ActionIpReleaseCampaignNotifyRequest{
		IpReleaseCampaign: inv.makeInputParams(),
		Meta:              inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpReleaseCampaignNotifyInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Event") {
			ret["event"] = inv.Input.Event
		}
	}

	return ret
}

func (inv *ActionIpReleaseCampaignNotifyInvocation) makeMetaInputParams() map[string]interface{} {
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
