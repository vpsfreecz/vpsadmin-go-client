package client

import (
	"strings"
)

// ActionDnsZoneUpdate is a type for action Dns_zone#Update
type ActionDnsZoneUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsZoneUpdate(client *Client) *ActionDnsZoneUpdate {
	return &ActionDnsZoneUpdate{
		Client: client,
	}
}

// ActionDnsZoneUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsZoneUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsZoneUpdateMetaGlobalInput) SetIncludes(value string) *ActionDnsZoneUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsZoneUpdateMetaGlobalInput) SetNo(value bool) *ActionDnsZoneUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsZoneUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsZoneUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneUpdateInput is a type for action input parameters
type ActionDnsZoneUpdateInput struct {
	DefaultTtl    int64  `json:"default_ttl"`
	DnssecEnabled bool   `json:"dnssec_enabled"`
	Email         string `json:"email"`
	Enabled       bool   `json:"enabled"`
	Label         string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDefaultTtl sets parameter DefaultTtl to value and selects it for sending
func (in *ActionDnsZoneUpdateInput) SetDefaultTtl(value int64) *ActionDnsZoneUpdateInput {
	in.DefaultTtl = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DefaultTtl"] = nil
	return in
}

// SetDnssecEnabled sets parameter DnssecEnabled to value and selects it for sending
func (in *ActionDnsZoneUpdateInput) SetDnssecEnabled(value bool) *ActionDnsZoneUpdateInput {
	in.DnssecEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnssecEnabled"] = nil
	return in
}

// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionDnsZoneUpdateInput) SetEmail(value string) *ActionDnsZoneUpdateInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionDnsZoneUpdateInput) SetEnabled(value bool) *ActionDnsZoneUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionDnsZoneUpdateInput) SetLabel(value string) *ActionDnsZoneUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneUpdateInput) SelectParameters(params ...string) *ActionDnsZoneUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsZoneUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsZoneUpdateInput) UnselectParameters(params ...string) *ActionDnsZoneUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsZoneUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneUpdateRequest is a type for the entire action request
type ActionDnsZoneUpdateRequest struct {
	DnsZone map[string]interface{} `json:"dns_zone"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionDnsZoneUpdateOutput is a type for action output parameters
type ActionDnsZoneUpdateOutput struct {
	CreatedAt             string                `json:"created_at"`
	DefaultTtl            int64                 `json:"default_ttl"`
	DnssecEnabled         bool                  `json:"dnssec_enabled"`
	Email                 string                `json:"email"`
	Enabled               bool                  `json:"enabled"`
	Id                    int64                 `json:"id"`
	Label                 string                `json:"label"`
	Name                  string                `json:"name"`
	ReverseNetworkAddress string                `json:"reverse_network_address"`
	ReverseNetworkPrefix  string                `json:"reverse_network_prefix"`
	Role                  string                `json:"role"`
	Serial                int64                 `json:"serial"`
	Source                string                `json:"source"`
	UpdatedAt             string                `json:"updated_at"`
	User                  *ActionUserShowOutput `json:"user"`
}

// ActionDnsZoneUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsZoneUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsZoneUpdateResponse struct {
	Action *ActionDnsZoneUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsZone *ActionDnsZoneUpdateOutput `json:"dns_zone"`
		// Global output metadata
		Meta *ActionDnsZoneUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDnsZoneUpdateOutput
}

// Prepare the action for invocation
func (action *ActionDnsZoneUpdate) Prepare() *ActionDnsZoneUpdateInvocation {
	return &ActionDnsZoneUpdateInvocation{
		Action: action,
		Path:   "/v7.0/dns_zones/{dns_zone_id}",
	}
}

// ActionDnsZoneUpdateInvocation is used to configure action for invocation
type ActionDnsZoneUpdateInvocation struct {
	// Pointer to the action
	Action *ActionDnsZoneUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsZoneUpdateInput
	// Global meta input parameters
	MetaInput *ActionDnsZoneUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsZoneUpdateInvocation) SetPathParamInt(param string, value int64) *ActionDnsZoneUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsZoneUpdateInvocation) SetPathParamString(param string, value string) *ActionDnsZoneUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsZoneUpdateInvocation) NewInput() *ActionDnsZoneUpdateInput {
	inv.Input = &ActionDnsZoneUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsZoneUpdateInvocation) SetInput(input *ActionDnsZoneUpdateInput) *ActionDnsZoneUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsZoneUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsZoneUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsZoneUpdateInvocation) NewMetaInput() *ActionDnsZoneUpdateMetaGlobalInput {
	inv.MetaInput = &ActionDnsZoneUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsZoneUpdateInvocation) SetMetaInput(input *ActionDnsZoneUpdateMetaGlobalInput) *ActionDnsZoneUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsZoneUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsZoneUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsZoneUpdateInvocation) Call() (*ActionDnsZoneUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsZoneUpdateInvocation) callAsBody() (*ActionDnsZoneUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsZoneUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsZone
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsZoneUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsZoneUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsZoneUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsZoneUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsZoneUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsZoneUpdateInvocation) makeAllInputParams() *ActionDnsZoneUpdateRequest {
	return &ActionDnsZoneUpdateRequest{
		DnsZone: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsZoneUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("DefaultTtl") {
			ret["default_ttl"] = inv.Input.DefaultTtl
		}
		if inv.IsParameterSelected("DnssecEnabled") {
			ret["dnssec_enabled"] = inv.Input.DnssecEnabled
		}
		if inv.IsParameterSelected("Email") {
			ret["email"] = inv.Input.Email
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionDnsZoneUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
