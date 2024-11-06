package client

import ()

// ActionDnsRecordCreate is a type for action Dns_record#Create
type ActionDnsRecordCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsRecordCreate(client *Client) *ActionDnsRecordCreate {
	return &ActionDnsRecordCreate{
		Client: client,
	}
}

// ActionDnsRecordCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsRecordCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsRecordCreateMetaGlobalInput) SetIncludes(value string) *ActionDnsRecordCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsRecordCreateMetaGlobalInput) SetNo(value bool) *ActionDnsRecordCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsRecordCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsRecordCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsRecordCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordCreateInput is a type for action input parameters
type ActionDnsRecordCreateInput struct {
	Comment              string `json:"comment"`
	Content              string `json:"content"`
	DnsZone              int64  `json:"dns_zone"`
	DynamicUpdateEnabled bool   `json:"dynamic_update_enabled"`
	Enabled              bool   `json:"enabled"`
	Name                 string `json:"name"`
	Priority             int64  `json:"priority"`
	Ttl                  int64  `json:"ttl"`
	Type                 string `json:"type"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetComment sets parameter Comment to value and selects it for sending
func (in *ActionDnsRecordCreateInput) SetComment(value string) *ActionDnsRecordCreateInput {
	in.Comment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Comment"] = nil
	return in
}

// SetContent sets parameter Content to value and selects it for sending
func (in *ActionDnsRecordCreateInput) SetContent(value string) *ActionDnsRecordCreateInput {
	in.Content = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Content"] = nil
	return in
}

// SetDnsZone sets parameter DnsZone to value and selects it for sending
func (in *ActionDnsRecordCreateInput) SetDnsZone(value int64) *ActionDnsRecordCreateInput {
	in.DnsZone = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDnsZoneNil(false)
	in._selectedParameters["DnsZone"] = nil
	return in
}

// SetDnsZoneNil sets parameter DnsZone to nil and selects it for sending
func (in *ActionDnsRecordCreateInput) SetDnsZoneNil(set bool) *ActionDnsRecordCreateInput {
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

// SetDynamicUpdateEnabled sets parameter DynamicUpdateEnabled to value and selects it for sending
func (in *ActionDnsRecordCreateInput) SetDynamicUpdateEnabled(value bool) *ActionDnsRecordCreateInput {
	in.DynamicUpdateEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DynamicUpdateEnabled"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionDnsRecordCreateInput) SetEnabled(value bool) *ActionDnsRecordCreateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionDnsRecordCreateInput) SetName(value string) *ActionDnsRecordCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetPriority sets parameter Priority to value and selects it for sending
func (in *ActionDnsRecordCreateInput) SetPriority(value int64) *ActionDnsRecordCreateInput {
	in.Priority = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Priority"] = nil
	return in
}

// SetTtl sets parameter Ttl to value and selects it for sending
func (in *ActionDnsRecordCreateInput) SetTtl(value int64) *ActionDnsRecordCreateInput {
	in.Ttl = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Ttl"] = nil
	return in
}

// SetType sets parameter Type to value and selects it for sending
func (in *ActionDnsRecordCreateInput) SetType(value string) *ActionDnsRecordCreateInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsRecordCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsRecordCreateInput) SelectParameters(params ...string) *ActionDnsRecordCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsRecordCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsRecordCreateInput) UnselectParameters(params ...string) *ActionDnsRecordCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsRecordCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsRecordCreateRequest is a type for the entire action request
type ActionDnsRecordCreateRequest struct {
	DnsRecord map[string]interface{} `json:"dns_record"`
	Meta      map[string]interface{} `json:"_meta"`
}

// ActionDnsRecordCreateOutput is a type for action output parameters
type ActionDnsRecordCreateOutput struct {
	Comment              string                   `json:"comment"`
	Content              string                   `json:"content"`
	CreatedAt            string                   `json:"created_at"`
	DnsZone              *ActionDnsZoneShowOutput `json:"dns_zone"`
	DynamicUpdateEnabled bool                     `json:"dynamic_update_enabled"`
	DynamicUpdateUrl     string                   `json:"dynamic_update_url"`
	Enabled              bool                     `json:"enabled"`
	Id                   int64                    `json:"id"`
	Name                 string                   `json:"name"`
	Priority             int64                    `json:"priority"`
	Ttl                  int64                    `json:"ttl"`
	Type                 string                   `json:"type"`
	UpdatedAt            string                   `json:"updated_at"`
}

// ActionDnsRecordCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionDnsRecordCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionDnsRecordCreateResponse struct {
	Action *ActionDnsRecordCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsRecord *ActionDnsRecordCreateOutput `json:"dns_record"`
		// Global output metadata
		Meta *ActionDnsRecordCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionDnsRecordCreateOutput
}

// Prepare the action for invocation
func (action *ActionDnsRecordCreate) Prepare() *ActionDnsRecordCreateInvocation {
	return &ActionDnsRecordCreateInvocation{
		Action: action,
		Path:   "/v7.0/dns_records",
	}
}

// ActionDnsRecordCreateInvocation is used to configure action for invocation
type ActionDnsRecordCreateInvocation struct {
	// Pointer to the action
	Action *ActionDnsRecordCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsRecordCreateInput
	// Global meta input parameters
	MetaInput *ActionDnsRecordCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsRecordCreateInvocation) NewInput() *ActionDnsRecordCreateInput {
	inv.Input = &ActionDnsRecordCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsRecordCreateInvocation) SetInput(input *ActionDnsRecordCreateInput) *ActionDnsRecordCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsRecordCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsRecordCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsRecordCreateInvocation) NewMetaInput() *ActionDnsRecordCreateMetaGlobalInput {
	inv.MetaInput = &ActionDnsRecordCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsRecordCreateInvocation) SetMetaInput(input *ActionDnsRecordCreateMetaGlobalInput) *ActionDnsRecordCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsRecordCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsRecordCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsRecordCreateInvocation) Call() (*ActionDnsRecordCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsRecordCreateInvocation) callAsBody() (*ActionDnsRecordCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsRecordCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsRecord
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionDnsRecordCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionDnsRecordCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionDnsRecordCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionDnsRecordCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
func (resp *ActionDnsRecordCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionDnsRecordCreateInvocation) makeAllInputParams() *ActionDnsRecordCreateRequest {
	return &ActionDnsRecordCreateRequest{
		DnsRecord: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsRecordCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Comment") {
			ret["comment"] = inv.Input.Comment
		}
		if inv.IsParameterSelected("Content") {
			ret["content"] = inv.Input.Content
		}
		if inv.IsParameterSelected("DnsZone") {
			if inv.IsParameterNil("DnsZone") {
				ret["dns_zone"] = nil
			} else {
				ret["dns_zone"] = inv.Input.DnsZone
			}
		}
		if inv.IsParameterSelected("DynamicUpdateEnabled") {
			ret["dynamic_update_enabled"] = inv.Input.DynamicUpdateEnabled
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Priority") {
			ret["priority"] = inv.Input.Priority
		}
		if inv.IsParameterSelected("Ttl") {
			ret["ttl"] = inv.Input.Ttl
		}
		if inv.IsParameterSelected("Type") {
			ret["type"] = inv.Input.Type
		}
	}

	return ret
}

func (inv *ActionDnsRecordCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
