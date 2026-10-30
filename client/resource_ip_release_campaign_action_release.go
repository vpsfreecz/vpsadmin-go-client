package client

import (
	"net/url"
	"strings"
)

// ActionIpReleaseCampaignRelease is a type for action Ip_release_campaign#Release
type ActionIpReleaseCampaignRelease struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignRelease(client *Client) *ActionIpReleaseCampaignRelease {
	return &ActionIpReleaseCampaignRelease{
		Client: client,
	}
}

// ActionIpReleaseCampaignReleaseMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignReleaseMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignReleaseMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignReleaseMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignReleaseMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignReleaseMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignReleaseMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignReleaseMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignReleaseMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignReleaseMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignReleaseRequest is a type for the entire action request
type ActionIpReleaseCampaignReleaseRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// ActionIpReleaseCampaignReleaseOutput is a type for action output parameters
type ActionIpReleaseCampaignReleaseOutput struct {
	BlockedChainId    int64                             "json:\"blocked_chain_id\""
	BlockedResource   string                            "json:\"blocked_resource\""
	BlockedResourceId int64                             "json:\"blocked_resource_id\""
	CreatedAt         string                            "json:\"created_at\""
	CreatedById       int64                             "json:\"created_by_id\""
	CreatedByLogin    string                            "json:\"created_by_login\""
	Error             string                            "json:\"error\""
	Id                int64                             "json:\"id\""
	IpCount           int64                             "json:\"ip_count\""
	State             string                            "json:\"state\""
	TransactionChain  *ActionTransactionChainShowOutput "json:\"transaction_chain\""
}

// ActionIpReleaseCampaignReleaseMetaGlobalOutput is a type for global output metadata parameters
type ActionIpReleaseCampaignReleaseMetaGlobalOutput struct {
	ActionStateId int64 "json:\"action_state_id\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignReleaseResponse struct {
	Action *ActionIpReleaseCampaignRelease "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseCampaign *ActionIpReleaseCampaignReleaseOutput "json:\"ip_release_campaign\""
		// Global output metadata
		Meta *ActionIpReleaseCampaignReleaseMetaGlobalOutput "json:\"_meta\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseCampaignReleaseOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignRelease) Prepare() *ActionIpReleaseCampaignReleaseInvocation {
	return &ActionIpReleaseCampaignReleaseInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/{ip_release_campaign_id}/release",
	}
}

// ActionIpReleaseCampaignReleaseInvocation is used to configure action for invocation
type ActionIpReleaseCampaignReleaseInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignRelease

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignReleaseMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseCampaignReleaseInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseCampaignReleaseInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseCampaignReleaseInvocation) SetPathParamString(param string, value string) *ActionIpReleaseCampaignReleaseInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignReleaseInvocation) NewMetaInput() *ActionIpReleaseCampaignReleaseMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignReleaseMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignReleaseInvocation) SetMetaInput(input *ActionIpReleaseCampaignReleaseMetaGlobalInput) *ActionIpReleaseCampaignReleaseInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignReleaseInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignReleaseInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignReleaseInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpReleaseCampaignReleaseInvocation) Call() (*ActionIpReleaseCampaignReleaseResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionIpReleaseCampaignReleaseInvocation) callAsBody() (*ActionIpReleaseCampaignReleaseResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIpReleaseCampaignReleaseResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpReleaseCampaign
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionIpReleaseCampaignReleaseResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionIpReleaseCampaignReleaseResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionIpReleaseCampaignReleaseResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionIpReleaseCampaignReleaseResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionIpReleaseCampaignReleaseResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionIpReleaseCampaignReleaseInvocation) makeAllInputParams() *ActionIpReleaseCampaignReleaseRequest {
	return &ActionIpReleaseCampaignReleaseRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpReleaseCampaignReleaseInvocation) makeMetaInputParams() map[string]interface{} {
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
