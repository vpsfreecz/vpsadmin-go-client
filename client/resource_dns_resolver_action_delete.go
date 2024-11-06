package client

import (
	"strings"
)

// ActionDnsResolverDelete is a type for action Dns_resolver#Delete
type ActionDnsResolverDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsResolverDelete(client *Client) *ActionDnsResolverDelete {
	return &ActionDnsResolverDelete{
		Client: client,
	}
}

// ActionDnsResolverDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionDnsResolverDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsResolverDeleteMetaGlobalInput) SetIncludes(value string) *ActionDnsResolverDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsResolverDeleteMetaGlobalInput) SetNo(value bool) *ActionDnsResolverDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsResolverDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsResolverDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionDnsResolverDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsResolverDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsResolverDeleteInput is a type for action input parameters
type ActionDnsResolverDeleteInput struct {
	Force bool `json:"force"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetForce sets parameter Force to value and selects it for sending
func (in *ActionDnsResolverDeleteInput) SetForce(value bool) *ActionDnsResolverDeleteInput {
	in.Force = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Force"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsResolverDeleteInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsResolverDeleteInput) SelectParameters(params ...string) *ActionDnsResolverDeleteInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsResolverDeleteInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsResolverDeleteInput) UnselectParameters(params ...string) *ActionDnsResolverDeleteInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsResolverDeleteInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsResolverDeleteRequest is a type for the entire action request
type ActionDnsResolverDeleteRequest struct {
	DnsResolver map[string]interface{} `json:"dns_resolver"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionDnsResolverDeleteMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsResolverDeleteMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsResolverDeleteResponse struct {
	Action *ActionDnsResolverDelete `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		// Global output metadata
		Meta *ActionDnsResolverDeleteMetaGlobalOutput `json:"_meta"`
	}
}

// Prepare the action for invocation
func (action *ActionDnsResolverDelete) Prepare() *ActionDnsResolverDeleteInvocation {
	return &ActionDnsResolverDeleteInvocation{
		Action: action,
		Path:   "/v7.0/dns_resolvers/{dns_resolver_id}",
	}
}

// ActionDnsResolverDeleteInvocation is used to configure action for invocation
type ActionDnsResolverDeleteInvocation struct {
	// Pointer to the action
	Action *ActionDnsResolverDelete

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsResolverDeleteInput
	// Global meta input parameters
	MetaInput *ActionDnsResolverDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsResolverDeleteInvocation) SetPathParamInt(param string, value int64) *ActionDnsResolverDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsResolverDeleteInvocation) SetPathParamString(param string, value string) *ActionDnsResolverDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsResolverDeleteInvocation) NewInput() *ActionDnsResolverDeleteInput {
	inv.Input = &ActionDnsResolverDeleteInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsResolverDeleteInvocation) SetInput(input *ActionDnsResolverDeleteInput) *ActionDnsResolverDeleteInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsResolverDeleteInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsResolverDeleteInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsResolverDeleteInvocation) NewMetaInput() *ActionDnsResolverDeleteMetaGlobalInput {
	inv.MetaInput = &ActionDnsResolverDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsResolverDeleteInvocation) SetMetaInput(input *ActionDnsResolverDeleteMetaGlobalInput) *ActionDnsResolverDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsResolverDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsResolverDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsResolverDeleteInvocation) Call() (*ActionDnsResolverDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsResolverDeleteInvocation) callAsBody() (*ActionDnsResolverDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsResolverDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsResolverDeleteResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsResolverDeleteResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsResolverDeleteResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsResolverDeleteResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsResolverDeleteResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsResolverDeleteInvocation) makeAllInputParams() *ActionDnsResolverDeleteRequest {
	return &ActionDnsResolverDeleteRequest{
		DnsResolver: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsResolverDeleteInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Force") {
			ret["force"] = inv.Input.Force
		}
	}

	return ret
}

func (inv *ActionDnsResolverDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
