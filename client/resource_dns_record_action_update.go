package client

import (
	"strings"
)

// ActionDnsRecordUpdate is a type for action Dns_record#Update
type ActionDnsRecordUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsRecordUpdate(client *Client) *ActionDnsRecordUpdate {
	return &ActionDnsRecordUpdate{
		Client: client,
	}
}

// ActionDnsRecordUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsRecordUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsRecordUpdateMetaGlobalInput) SetIncludes(value string) *ActionDnsRecordUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsRecordUpdateMetaGlobalInput) SetNo(value bool) *ActionDnsRecordUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsRecordUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsRecordUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsRecordUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordUpdateInput is a type for action input parameters
type ActionDnsRecordUpdateInput struct {
	Comment              string `json:"comment"`
	Content              string `json:"content"`
	DynamicUpdateEnabled bool   `json:"dynamic_update_enabled"`
	Enabled              bool   `json:"enabled"`
	Priority             int64  `json:"priority"`
	Ttl                  int64  `json:"ttl"`
	User                 int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetComment sets parameter Comment to value and selects it for sending
func (in *ActionDnsRecordUpdateInput) SetComment(value string) *ActionDnsRecordUpdateInput {
	in.Comment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Comment"] = nil
	return in
}

// SetContent sets parameter Content to value and selects it for sending
func (in *ActionDnsRecordUpdateInput) SetContent(value string) *ActionDnsRecordUpdateInput {
	in.Content = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Content"] = nil
	return in
}

// SetDynamicUpdateEnabled sets parameter DynamicUpdateEnabled to value and selects it for sending
func (in *ActionDnsRecordUpdateInput) SetDynamicUpdateEnabled(value bool) *ActionDnsRecordUpdateInput {
	in.DynamicUpdateEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DynamicUpdateEnabled"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionDnsRecordUpdateInput) SetEnabled(value bool) *ActionDnsRecordUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetPriority sets parameter Priority to value and selects it for sending
func (in *ActionDnsRecordUpdateInput) SetPriority(value int64) *ActionDnsRecordUpdateInput {
	in.Priority = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Priority"] = nil
	return in
}

// SetTtl sets parameter Ttl to value and selects it for sending
func (in *ActionDnsRecordUpdateInput) SetTtl(value int64) *ActionDnsRecordUpdateInput {
	in.Ttl = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ttl"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionDnsRecordUpdateInput) SetUser(value int64) *ActionDnsRecordUpdateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionDnsRecordUpdateInput) SetUserNil(set bool) *ActionDnsRecordUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionDnsRecordUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordUpdateInput) SelectParameters(params ...string) *ActionDnsRecordUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsRecordUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsRecordUpdateInput) UnselectParameters(params ...string) *ActionDnsRecordUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsRecordUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordUpdateRequest is a type for the entire action request
type ActionDnsRecordUpdateRequest struct {
	DnsRecord map[string]interface{} `json:"dns_record"`
	Meta      map[string]interface{} `json:"_meta"`
}

// ActionDnsRecordUpdateOutput is a type for action output parameters
type ActionDnsRecordUpdateOutput struct {
	Comment              string                   `json:"comment"`
	Content              string                   `json:"content"`
	CreatedAt            string                   `json:"created_at"`
	DnsZone              *ActionDnsZoneShowOutput `json:"dns_zone"`
	DynamicUpdateEnabled bool                     `json:"dynamic_update_enabled"`
	DynamicUpdateUrl     string                   `json:"dynamic_update_url"`
	Enabled              bool                     `json:"enabled"`
	Id                   int64                    `json:"id"`
	Managed              bool                     `json:"managed"`
	Name                 string                   `json:"name"`
	Priority             int64                    `json:"priority"`
	Ttl                  int64                    `json:"ttl"`
	Type                 string                   `json:"type"`
	UpdatedAt            string                   `json:"updated_at"`
	User                 *ActionUserShowOutput    `json:"user"`
}

// ActionDnsRecordUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsRecordUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsRecordUpdateResponse struct {
	Action *ActionDnsRecordUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsRecord *ActionDnsRecordUpdateOutput `json:"dns_record"`
		// Global output metadata
		Meta *ActionDnsRecordUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDnsRecordUpdateOutput
}

// Prepare the action for invocation
func (action *ActionDnsRecordUpdate) Prepare() *ActionDnsRecordUpdateInvocation {
	return &ActionDnsRecordUpdateInvocation{
		Action: action,
		Path:   "/v7.0/dns_records/{dns_record_id}",
	}
}

// ActionDnsRecordUpdateInvocation is used to configure action for invocation
type ActionDnsRecordUpdateInvocation struct {
	// Pointer to the action
	Action *ActionDnsRecordUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsRecordUpdateInput
	// Global meta input parameters
	MetaInput *ActionDnsRecordUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsRecordUpdateInvocation) SetPathParamInt(param string, value int64) *ActionDnsRecordUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsRecordUpdateInvocation) SetPathParamString(param string, value string) *ActionDnsRecordUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsRecordUpdateInvocation) NewInput() *ActionDnsRecordUpdateInput {
	inv.Input = &ActionDnsRecordUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsRecordUpdateInvocation) SetInput(input *ActionDnsRecordUpdateInput) *ActionDnsRecordUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsRecordUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsRecordUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsRecordUpdateInvocation) NewMetaInput() *ActionDnsRecordUpdateMetaGlobalInput {
	inv.MetaInput = &ActionDnsRecordUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsRecordUpdateInvocation) SetMetaInput(input *ActionDnsRecordUpdateMetaGlobalInput) *ActionDnsRecordUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsRecordUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsRecordUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsRecordUpdateInvocation) Call() (*ActionDnsRecordUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsRecordUpdateInvocation) callAsBody() (*ActionDnsRecordUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsRecordUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsRecord
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsRecordUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsRecordUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsRecordUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsRecordUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsRecordUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsRecordUpdateInvocation) makeAllInputParams() *ActionDnsRecordUpdateRequest {
	return &ActionDnsRecordUpdateRequest{
		DnsRecord: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsRecordUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Comment") {
			ret["comment"] = inv.Input.Comment
		}
		if inv.IsParameterSelected("Content") {
			ret["content"] = inv.Input.Content
		}
		if inv.IsParameterSelected("DynamicUpdateEnabled") {
			ret["dynamic_update_enabled"] = inv.Input.DynamicUpdateEnabled
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("Priority") {
			ret["priority"] = inv.Input.Priority
		}
		if inv.IsParameterSelected("Ttl") {
			ret["ttl"] = inv.Input.Ttl
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["user"] = nil
			} else {
				ret["user"] = inv.Input.User
			}
		}
	}

	return ret
}

func (inv *ActionDnsRecordUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
