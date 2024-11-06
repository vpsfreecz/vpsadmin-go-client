package client

import ()

// ActionDnsZoneTransferCreate is a type for action Dns_zone_transfer#Create
type ActionDnsZoneTransferCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsZoneTransferCreate(client *Client) *ActionDnsZoneTransferCreate {
	return &ActionDnsZoneTransferCreate{
		Client: client,
	}
}

// ActionDnsZoneTransferCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsZoneTransferCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsZoneTransferCreateMetaGlobalInput) SetIncludes(value string) *ActionDnsZoneTransferCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsZoneTransferCreateMetaGlobalInput) SetNo(value bool) *ActionDnsZoneTransferCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneTransferCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneTransferCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsZoneTransferCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsZoneTransferCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneTransferCreateInput is a type for action input parameters
type ActionDnsZoneTransferCreateInput struct {
	DnsTsigKey    int64  `json:"dns_tsig_key"`
	DnsZone       int64  `json:"dns_zone"`
	HostIpAddress int64  `json:"host_ip_address"`
	PeerType      string `json:"peer_type"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDnsTsigKey sets parameter DnsTsigKey to value and selects it for sending
func (in *ActionDnsZoneTransferCreateInput) SetDnsTsigKey(value int64) *ActionDnsZoneTransferCreateInput {
	in.DnsTsigKey = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsTsigKeyNil(false)
	in._selectedParameters["DnsTsigKey"] = nil
	return in
}

// SetDnsTsigKeyNil sets parameter DnsTsigKey to nil and selects it for sending
func (in *ActionDnsZoneTransferCreateInput) SetDnsTsigKeyNil(set bool) *ActionDnsZoneTransferCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["DnsTsigKey"] = nil
		in.SelectParameters("DnsTsigKey")
	} else {
		delete(in._nilParameters, "DnsTsigKey")
	}
	return in
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnsZoneTransferCreateInput) SetDnsZone(value int64) *ActionDnsZoneTransferCreateInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsZoneNil(false)
	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetDnsZoneNil sets parameter DnsZone to nil and selects it for sending
func (in *ActionDnsZoneTransferCreateInput) SetDnsZoneNil(set bool) *ActionDnsZoneTransferCreateInput {
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

// SetHostIpAddress sets parameter HostIpAddress to value and selects it for sending
func (in *ActionDnsZoneTransferCreateInput) SetHostIpAddress(value int64) *ActionDnsZoneTransferCreateInput {
	in.HostIpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetHostIpAddressNil(false)
	in._selectedParameters["HostIpAddress"] = nil
	return in
}

// SetHostIpAddressNil sets parameter HostIpAddress to nil and selects it for sending
func (in *ActionDnsZoneTransferCreateInput) SetHostIpAddressNil(set bool) *ActionDnsZoneTransferCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["HostIpAddress"] = nil
		in.SelectParameters("HostIpAddress")
	} else {
		delete(in._nilParameters, "HostIpAddress")
	}
	return in
}

// SetPeerType sets parameter PeerType to value and selects it for sending
func (in *ActionDnsZoneTransferCreateInput) SetPeerType(value string) *ActionDnsZoneTransferCreateInput {
	in.PeerType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PeerType"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneTransferCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneTransferCreateInput) SelectParameters(params ...string) *ActionDnsZoneTransferCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsZoneTransferCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsZoneTransferCreateInput) UnselectParameters(params ...string) *ActionDnsZoneTransferCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsZoneTransferCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneTransferCreateRequest is a type for the entire action request
type ActionDnsZoneTransferCreateRequest struct {
	DnsZoneTransfer map[string]interface{} `json:"dns_zone_transfer"`
	Meta            map[string]interface{} `json:"_meta"`
}

// ActionDnsZoneTransferCreateOutput is a type for action output parameters
type ActionDnsZoneTransferCreateOutput struct {
	CreatedAt     string                         `json:"created_at"`
	DnsTsigKey    *ActionDnsTsigKeyShowOutput    `json:"dns_tsig_key"`
	DnsZone       *ActionDnsZoneShowOutput       `json:"dns_zone"`
	HostIpAddress *ActionHostIpAddressShowOutput `json:"host_ip_address"`
	Id            int64                          `json:"id"`
	PeerType      string                         `json:"peer_type"`
	UpdatedAt     string                         `json:"updated_at"`
}

// ActionDnsZoneTransferCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsZoneTransferCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsZoneTransferCreateResponse struct {
	Action *ActionDnsZoneTransferCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsZoneTransfer *ActionDnsZoneTransferCreateOutput `json:"dns_zone_transfer"`
		// Global output metadata
		Meta *ActionDnsZoneTransferCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDnsZoneTransferCreateOutput
}

// Prepare the action for invocation
func (action *ActionDnsZoneTransferCreate) Prepare() *ActionDnsZoneTransferCreateInvocation {
	return &ActionDnsZoneTransferCreateInvocation{
		Action: action,
		Path:   "/v7.0/dns_zone_transfers",
	}
}

// ActionDnsZoneTransferCreateInvocation is used to configure action for invocation
type ActionDnsZoneTransferCreateInvocation struct {
	// Pointer to the action
	Action *ActionDnsZoneTransferCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsZoneTransferCreateInput
	// Global meta input parameters
	MetaInput *ActionDnsZoneTransferCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsZoneTransferCreateInvocation) NewInput() *ActionDnsZoneTransferCreateInput {
	inv.Input = &ActionDnsZoneTransferCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsZoneTransferCreateInvocation) SetInput(input *ActionDnsZoneTransferCreateInput) *ActionDnsZoneTransferCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsZoneTransferCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsZoneTransferCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsZoneTransferCreateInvocation) NewMetaInput() *ActionDnsZoneTransferCreateMetaGlobalInput {
	inv.MetaInput = &ActionDnsZoneTransferCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsZoneTransferCreateInvocation) SetMetaInput(input *ActionDnsZoneTransferCreateMetaGlobalInput) *ActionDnsZoneTransferCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsZoneTransferCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsZoneTransferCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsZoneTransferCreateInvocation) Call() (*ActionDnsZoneTransferCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsZoneTransferCreateInvocation) callAsBody() (*ActionDnsZoneTransferCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsZoneTransferCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsZoneTransfer
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsZoneTransferCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsZoneTransferCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsZoneTransferCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsZoneTransferCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsZoneTransferCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsZoneTransferCreateInvocation) makeAllInputParams() *ActionDnsZoneTransferCreateRequest {
	return &ActionDnsZoneTransferCreateRequest{
		DnsZoneTransfer: inv.makeInputParams(),
		Meta:            inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsZoneTransferCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("DnsTsigKey") {
			if inv.IsParameterNil("DnsTsigKey") {
				ret["dns_tsig_key"] = nil
			} else {
				ret["dns_tsig_key"] = inv.Input.DnsTsigKey
			}
		}
		if inv.IsParameterSelected("DnsZone") {
			if inv.IsParameterNil("DnsZone") {
				ret["dns_zone"] = nil
			} else {
				ret["dns_zone"] = inv.Input.DnsZone
			}
		}
		if inv.IsParameterSelected("HostIpAddress") {
			if inv.IsParameterNil("HostIpAddress") {
				ret["host_ip_address"] = nil
			} else {
				ret["host_ip_address"] = inv.Input.HostIpAddress
			}
		}
		if inv.IsParameterSelected("PeerType") {
			ret["peer_type"] = inv.Input.PeerType
		}
	}

	return ret
}

func (inv *ActionDnsZoneTransferCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
