package client

import (
	"strings"
)

// ActionDnsRecordDynamicUpdate is a type for action Dns_record#Dynamic_update
type ActionDnsRecordDynamicUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsRecordDynamicUpdate(client *Client) *ActionDnsRecordDynamicUpdate {
	return &ActionDnsRecordDynamicUpdate{
		Client: client,
	}
}

// ActionDnsRecordDynamicUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsRecordDynamicUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsRecordDynamicUpdateMetaGlobalInput) SetIncludes(value string) *ActionDnsRecordDynamicUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsRecordDynamicUpdateMetaGlobalInput) SetNo(value bool) *ActionDnsRecordDynamicUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsRecordDynamicUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordDynamicUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsRecordDynamicUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsRecordDynamicUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordDynamicUpdateOutput is a type for action output parameters
type ActionDnsRecordDynamicUpdateOutput struct {
	Content string `json:"content"`
}

// ActionDnsRecordDynamicUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsRecordDynamicUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsRecordDynamicUpdateResponse struct {
	Action *ActionDnsRecordDynamicUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsRecord *ActionDnsRecordDynamicUpdateOutput `json:"dns_record"`
		// Global output metadata
		Meta *ActionDnsRecordDynamicUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDnsRecordDynamicUpdateOutput
}

// Prepare the action for invocation
func (action *ActionDnsRecordDynamicUpdate) Prepare() *ActionDnsRecordDynamicUpdateInvocation {
	return &ActionDnsRecordDynamicUpdateInvocation{
		Action: action,
		Path:   "/v7.0/dns_records/dynamic_update/{access_token}",
	}
}

// ActionDnsRecordDynamicUpdateInvocation is used to configure action for invocation
type ActionDnsRecordDynamicUpdateInvocation struct {
	// Pointer to the action
	Action *ActionDnsRecordDynamicUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsRecordDynamicUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsRecordDynamicUpdateInvocation) SetPathParamInt(param string, value int64) *ActionDnsRecordDynamicUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsRecordDynamicUpdateInvocation) SetPathParamString(param string, value string) *ActionDnsRecordDynamicUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsRecordDynamicUpdateInvocation) NewMetaInput() *ActionDnsRecordDynamicUpdateMetaGlobalInput {
	inv.MetaInput = &ActionDnsRecordDynamicUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsRecordDynamicUpdateInvocation) SetMetaInput(input *ActionDnsRecordDynamicUpdateMetaGlobalInput) *ActionDnsRecordDynamicUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsRecordDynamicUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsRecordDynamicUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsRecordDynamicUpdateInvocation) Call() (*ActionDnsRecordDynamicUpdateResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsRecordDynamicUpdateInvocation) callAsQuery() (*ActionDnsRecordDynamicUpdateResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsRecordDynamicUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsRecord
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsRecordDynamicUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsRecordDynamicUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsRecordDynamicUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsRecordDynamicUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsRecordDynamicUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsRecordDynamicUpdateInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
