package client

import ()

// ActionDnsServerZoneCreate is a type for action Dns_server_zone#Create
type ActionDnsServerZoneCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsServerZoneCreate(client *Client) *ActionDnsServerZoneCreate {
	return &ActionDnsServerZoneCreate{
		Client: client,
	}
}

// ActionDnsServerZoneCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsServerZoneCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsServerZoneCreateMetaGlobalInput) SetIncludes(value string) *ActionDnsServerZoneCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsServerZoneCreateMetaGlobalInput) SetNo(value bool) *ActionDnsServerZoneCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZoneCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZoneCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsServerZoneCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsServerZoneCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZoneCreateInput is a type for action input parameters
type ActionDnsServerZoneCreateInput struct {
	DnsServer int64  `json:"dns_server"`
	DnsZone   int64  `json:"dns_zone"`
	Type      string `json:"type"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDnsServer sets parameter DnsServer to value and selects it for sending
func (in *ActionDnsServerZoneCreateInput) SetDnsServer(value int64) *ActionDnsServerZoneCreateInput {
	in.DnsServer = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsServerNil(false)
	in._selectedParameters["DnsServer"] = nil
	return in
}

// SetDnsServerNil sets parameter DnsServer to nil and selects it for sending
func (in *ActionDnsServerZoneCreateInput) SetDnsServerNil(set bool) *ActionDnsServerZoneCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["DnsServer"] = nil
		in.SelectParameters("DnsServer")
	} else {
		delete(in._nilParameters, "DnsServer")
	}
	return in
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnsServerZoneCreateInput) SetDnsZone(value int64) *ActionDnsServerZoneCreateInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsZoneNil(false)
	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetDnsZoneNil sets parameter DnsZone to nil and selects it for sending
func (in *ActionDnsServerZoneCreateInput) SetDnsZoneNil(set bool) *ActionDnsServerZoneCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["DnsZone"] = nil
		in.SelectParameters("DnsZone")
	} else {
		delete(in._nilParameters, "DnsZone")
	}
	return in
}

// SetType sets parameter Type to value and selects it for sending
func (in *ActionDnsServerZoneCreateInput) SetType(value string) *ActionDnsServerZoneCreateInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsServerZoneCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsServerZoneCreateInput) SelectParameters(params ...string) *ActionDnsServerZoneCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsServerZoneCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsServerZoneCreateInput) UnselectParameters(params ...string) *ActionDnsServerZoneCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsServerZoneCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsServerZoneCreateRequest is a type for the entire action request
type ActionDnsServerZoneCreateRequest struct {
	DnsServerZone map[string]interface{} `json:"dns_server_zone"`
	Meta          map[string]interface{} `json:"_meta"`
}

// ActionDnsServerZoneCreateOutput is a type for action output parameters
type ActionDnsServerZoneCreateOutput struct {
	CreatedAt   string                     `json:"created_at"`
	DnsServer   *ActionDnsServerShowOutput `json:"dns_server"`
	DnsZone     *ActionDnsZoneShowOutput   `json:"dns_zone"`
	ExpiresAt   string                     `json:"expires_at"`
	Id          int64                      `json:"id"`
	LastCheckAt string                     `json:"last_check_at"`
	LoadedAt    string                     `json:"loaded_at"`
	RefreshAt   string                     `json:"refresh_at"`
	Serial      int64                      `json:"serial"`
	Type        string                     `json:"type"`
	UpdatedAt   string                     `json:"updated_at"`
}

// ActionDnsServerZoneCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsServerZoneCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsServerZoneCreateResponse struct {
	Action *ActionDnsServerZoneCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsServerZone *ActionDnsServerZoneCreateOutput `json:"dns_server_zone"`
		// Global output metadata
		Meta *ActionDnsServerZoneCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDnsServerZoneCreateOutput
}

// Prepare the action for invocation
func (action *ActionDnsServerZoneCreate) Prepare() *ActionDnsServerZoneCreateInvocation {
	return &ActionDnsServerZoneCreateInvocation{
		Action: action,
		Path:   "/v7.0/dns_server_zones",
	}
}

// ActionDnsServerZoneCreateInvocation is used to configure action for invocation
type ActionDnsServerZoneCreateInvocation struct {
	// Pointer to the action
	Action *ActionDnsServerZoneCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsServerZoneCreateInput
	// Global meta input parameters
	MetaInput *ActionDnsServerZoneCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsServerZoneCreateInvocation) NewInput() *ActionDnsServerZoneCreateInput {
	inv.Input = &ActionDnsServerZoneCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsServerZoneCreateInvocation) SetInput(input *ActionDnsServerZoneCreateInput) *ActionDnsServerZoneCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsServerZoneCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsServerZoneCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsServerZoneCreateInvocation) NewMetaInput() *ActionDnsServerZoneCreateMetaGlobalInput {
	inv.MetaInput = &ActionDnsServerZoneCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsServerZoneCreateInvocation) SetMetaInput(input *ActionDnsServerZoneCreateMetaGlobalInput) *ActionDnsServerZoneCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsServerZoneCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsServerZoneCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsServerZoneCreateInvocation) Call() (*ActionDnsServerZoneCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsServerZoneCreateInvocation) callAsBody() (*ActionDnsServerZoneCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsServerZoneCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsServerZone
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsServerZoneCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsServerZoneCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsServerZoneCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsServerZoneCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsServerZoneCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsServerZoneCreateInvocation) makeAllInputParams() *ActionDnsServerZoneCreateRequest {
	return &ActionDnsServerZoneCreateRequest{
		DnsServerZone: inv.makeInputParams(),
		Meta:          inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsServerZoneCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("DnsServer") {
			if inv.IsParameterNil("DnsServer") {
				ret["dns_server"] = nil
			} else {
				ret["dns_server"] = inv.Input.DnsServer
			}
		}
		if inv.IsParameterSelected("DnsZone") {
			if inv.IsParameterNil("DnsZone") {
				ret["dns_zone"] = nil
			} else {
				ret["dns_zone"] = inv.Input.DnsZone
			}
		}
		if inv.IsParameterSelected("Type") {
			ret["type"] = inv.Input.Type
		}
	}

	return ret
}

func (inv *ActionDnsServerZoneCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
