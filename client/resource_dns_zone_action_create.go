package client

import ()

// ActionDnsZoneCreate is a type for action Dns_zone#Create
type ActionDnsZoneCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsZoneCreate(client *Client) *ActionDnsZoneCreate {
	return &ActionDnsZoneCreate{
		Client: client,
	}
}

// ActionDnsZoneCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsZoneCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsZoneCreateMetaGlobalInput) SetIncludes(value string) *ActionDnsZoneCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsZoneCreateMetaGlobalInput) SetNo(value bool) *ActionDnsZoneCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsZoneCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsZoneCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsZoneCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneCreateInput is a type for action input parameters
type ActionDnsZoneCreateInput struct {
	DefaultTtl            int64  `json:"default_ttl"`
	DnssecEnabled         bool   `json:"dnssec_enabled"`
	Email                 string `json:"email"`
	Enabled               bool   `json:"enabled"`
	Label                 string `json:"label"`
	Name                  string `json:"name"`
	ReverseNetworkAddress string `json:"reverse_network_address"`
	ReverseNetworkPrefix  string `json:"reverse_network_prefix"`
	Role                  string `json:"role"`
	SeedVps               int64  `json:"seed_vps"`
	Source                string `json:"source"`
	User                  int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDefaultTtl sets parameter DefaultTtl to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetDefaultTtl(value int64) *ActionDnsZoneCreateInput {
	in.DefaultTtl = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DefaultTtl"] = nil
	return in
}

// SetDnssecEnabled sets parameter DnssecEnabled to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetDnssecEnabled(value bool) *ActionDnsZoneCreateInput {
	in.DnssecEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DnssecEnabled"] = nil
	return in
}

// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetEmail(value string) *ActionDnsZoneCreateInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetEnabled(value bool) *ActionDnsZoneCreateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetLabel(value string) *ActionDnsZoneCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetName(value string) *ActionDnsZoneCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetReverseNetworkAddress sets parameter ReverseNetworkAddress to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetReverseNetworkAddress(value string) *ActionDnsZoneCreateInput {
	in.ReverseNetworkAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReverseNetworkAddress"] = nil
	return in
}

// SetReverseNetworkPrefix sets parameter ReverseNetworkPrefix to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetReverseNetworkPrefix(value string) *ActionDnsZoneCreateInput {
	in.ReverseNetworkPrefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReverseNetworkPrefix"] = nil
	return in
}

// SetRole sets parameter Role to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetRole(value string) *ActionDnsZoneCreateInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}

// SetSeedVps sets parameter SeedVps to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetSeedVps(value int64) *ActionDnsZoneCreateInput {
	in.SeedVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSeedVpsNil(false)
	in._selectedParameters["SeedVps"] = nil
	return in
}

// SetSeedVpsNil sets parameter SeedVps to nil and selects it for sending
func (in *ActionDnsZoneCreateInput) SetSeedVpsNil(set bool) *ActionDnsZoneCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["SeedVps"] = nil
		in.SelectParameters("SeedVps")
	} else {
		delete(in._nilParameters, "SeedVps")
	}
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetSource(value string) *ActionDnsZoneCreateInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionDnsZoneCreateInput) SetUser(value int64) *ActionDnsZoneCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionDnsZoneCreateInput) SetUserNil(set bool) *ActionDnsZoneCreateInput {
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

// SelectParameters sets parameters from ActionDnsZoneCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsZoneCreateInput) SelectParameters(params ...string) *ActionDnsZoneCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsZoneCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsZoneCreateInput) UnselectParameters(params ...string) *ActionDnsZoneCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsZoneCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsZoneCreateRequest is a type for the entire action request
type ActionDnsZoneCreateRequest struct {
	DnsZone map[string]interface{} `json:"dns_zone"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionDnsZoneCreateOutput is a type for action output parameters
type ActionDnsZoneCreateOutput struct {
	CreatedAt             string                `json:"created_at"`
	DefaultTtl            int64                 `json:"default_ttl"`
	DnssecEnabled         bool                  `json:"dnssec_enabled"`
	Email                 string                `json:"email"`
	Enabled               bool                  `json:"enabled"`
	Id                    int64                 `json:"id"`
	Label                 string                `json:"label"`
	Managed               bool                  `json:"managed"`
	Name                  string                `json:"name"`
	ReverseNetworkAddress string                `json:"reverse_network_address"`
	ReverseNetworkPrefix  string                `json:"reverse_network_prefix"`
	Role                  string                `json:"role"`
	Serial                int64                 `json:"serial"`
	Source                string                `json:"source"`
	UpdatedAt             string                `json:"updated_at"`
	User                  *ActionUserShowOutput `json:"user"`
}

// ActionDnsZoneCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsZoneCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsZoneCreateResponse struct {
	Action *ActionDnsZoneCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsZone *ActionDnsZoneCreateOutput `json:"dns_zone"`
		// Global output metadata
		Meta *ActionDnsZoneCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDnsZoneCreateOutput
}

// Prepare the action for invocation
func (action *ActionDnsZoneCreate) Prepare() *ActionDnsZoneCreateInvocation {
	return &ActionDnsZoneCreateInvocation{
		Action: action,
		Path:   "/v7.0/dns_zones",
	}
}

// ActionDnsZoneCreateInvocation is used to configure action for invocation
type ActionDnsZoneCreateInvocation struct {
	// Pointer to the action
	Action *ActionDnsZoneCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsZoneCreateInput
	// Global meta input parameters
	MetaInput *ActionDnsZoneCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsZoneCreateInvocation) NewInput() *ActionDnsZoneCreateInput {
	inv.Input = &ActionDnsZoneCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsZoneCreateInvocation) SetInput(input *ActionDnsZoneCreateInput) *ActionDnsZoneCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsZoneCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsZoneCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsZoneCreateInvocation) NewMetaInput() *ActionDnsZoneCreateMetaGlobalInput {
	inv.MetaInput = &ActionDnsZoneCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsZoneCreateInvocation) SetMetaInput(input *ActionDnsZoneCreateMetaGlobalInput) *ActionDnsZoneCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsZoneCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsZoneCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsZoneCreateInvocation) Call() (*ActionDnsZoneCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsZoneCreateInvocation) callAsBody() (*ActionDnsZoneCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsZoneCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsZone
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsZoneCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsZoneCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsZoneCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsZoneCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsZoneCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsZoneCreateInvocation) makeAllInputParams() *ActionDnsZoneCreateRequest {
	return &ActionDnsZoneCreateRequest{
		DnsZone: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsZoneCreateInvocation) makeInputParams() map[string]interface{} {
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
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("ReverseNetworkAddress") {
			ret["reverse_network_address"] = inv.Input.ReverseNetworkAddress
		}
		if inv.IsParameterSelected("ReverseNetworkPrefix") {
			ret["reverse_network_prefix"] = inv.Input.ReverseNetworkPrefix
		}
		if inv.IsParameterSelected("Role") {
			ret["role"] = inv.Input.Role
		}
		if inv.IsParameterSelected("SeedVps") {
			if inv.IsParameterNil("SeedVps") {
				ret["seed_vps"] = nil
			} else {
				ret["seed_vps"] = inv.Input.SeedVps
			}
		}
		if inv.IsParameterSelected("Source") {
			ret["source"] = inv.Input.Source
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

func (inv *ActionDnsZoneCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
