package client

import ()

// ActionNodeTransferConnectionCreate is a type for action Node_transfer_connection#Create
type ActionNodeTransferConnectionCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeTransferConnectionCreate(client *Client) *ActionNodeTransferConnectionCreate {
	return &ActionNodeTransferConnectionCreate{
		Client: client,
	}
}

// ActionNodeTransferConnectionCreateMetaGlobalInput is a type for action global meta input parameters
type ActionNodeTransferConnectionCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeTransferConnectionCreateMetaGlobalInput) SetIncludes(value string) *ActionNodeTransferConnectionCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeTransferConnectionCreateMetaGlobalInput) SetNo(value bool) *ActionNodeTransferConnectionCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeTransferConnectionCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionCreateMetaGlobalInput) SelectParameters(params ...string) *ActionNodeTransferConnectionCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeTransferConnectionCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeTransferConnectionCreateInput is a type for action input parameters
type ActionNodeTransferConnectionCreateInput struct {
	Enabled     bool   "json:\"enabled\""
	NodeA       int64  "json:\"node_a\""
	NodeAIpAddr string "json:\"node_a_ip_addr\""
	NodeB       int64  "json:\"node_b\""
	NodeBIpAddr string "json:\"node_b_ip_addr\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionNodeTransferConnectionCreateInput) SetEnabled(value bool) *ActionNodeTransferConnectionCreateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetNodeA sets parameter NodeA to value and selects it for sending
func (in *ActionNodeTransferConnectionCreateInput) SetNodeA(value int64) *ActionNodeTransferConnectionCreateInput {
	in.NodeA = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeA"] = nil
	return in
}

// SetNodeAIpAddr sets parameter NodeAIpAddr to value and selects it for sending
func (in *ActionNodeTransferConnectionCreateInput) SetNodeAIpAddr(value string) *ActionNodeTransferConnectionCreateInput {
	in.NodeAIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeAIpAddr"] = nil
	return in
}

// SetNodeB sets parameter NodeB to value and selects it for sending
func (in *ActionNodeTransferConnectionCreateInput) SetNodeB(value int64) *ActionNodeTransferConnectionCreateInput {
	in.NodeB = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeB"] = nil
	return in
}

// SetNodeBIpAddr sets parameter NodeBIpAddr to value and selects it for sending
func (in *ActionNodeTransferConnectionCreateInput) SetNodeBIpAddr(value string) *ActionNodeTransferConnectionCreateInput {
	in.NodeBIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeBIpAddr"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeTransferConnectionCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionCreateInput) SelectParameters(params ...string) *ActionNodeTransferConnectionCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeTransferConnectionCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionCreateInput) UnselectParameters(params ...string) *ActionNodeTransferConnectionCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeTransferConnectionCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeTransferConnectionCreateRequest is a type for the entire action request
type ActionNodeTransferConnectionCreateRequest struct {
	NodeTransferConnection map[string]interface{} "json:\"node_transfer_connection\""
	Meta                   map[string]interface{} "json:\"_meta\""
}

// ActionNodeTransferConnectionCreateOutput is a type for action output parameters
type ActionNodeTransferConnectionCreateOutput struct {
	CreatedAt   string                "json:\"created_at\""
	Enabled     bool                  "json:\"enabled\""
	Id          int64                 "json:\"id\""
	NodeA       *ActionNodeShowOutput "json:\"node_a\""
	NodeAIpAddr string                "json:\"node_a_ip_addr\""
	NodeB       *ActionNodeShowOutput "json:\"node_b\""
	NodeBIpAddr string                "json:\"node_b_ip_addr\""
	UpdatedAt   string                "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionNodeTransferConnectionCreateResponse struct {
	Action *ActionNodeTransferConnectionCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeTransferConnection *ActionNodeTransferConnectionCreateOutput "json:\"node_transfer_connection\""
	}

	// Action output without the namespace
	Output *ActionNodeTransferConnectionCreateOutput
}

// Prepare the action for invocation
func (action *ActionNodeTransferConnectionCreate) Prepare() *ActionNodeTransferConnectionCreateInvocation {
	return &ActionNodeTransferConnectionCreateInvocation{
		Action: action,
		Path:   "/v7.0/node_transfer_connections",
	}
}

// ActionNodeTransferConnectionCreateInvocation is used to configure action for invocation
type ActionNodeTransferConnectionCreateInvocation struct {
	// Pointer to the action
	Action *ActionNodeTransferConnectionCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeTransferConnectionCreateInput
	// Global meta input parameters
	MetaInput *ActionNodeTransferConnectionCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeTransferConnectionCreateInvocation) NewInput() *ActionNodeTransferConnectionCreateInput {
	inv.Input = &ActionNodeTransferConnectionCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeTransferConnectionCreateInvocation) SetInput(input *ActionNodeTransferConnectionCreateInput) *ActionNodeTransferConnectionCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeTransferConnectionCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeTransferConnectionCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeTransferConnectionCreateInvocation) NewMetaInput() *ActionNodeTransferConnectionCreateMetaGlobalInput {
	inv.MetaInput = &ActionNodeTransferConnectionCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeTransferConnectionCreateInvocation) SetMetaInput(input *ActionNodeTransferConnectionCreateMetaGlobalInput) *ActionNodeTransferConnectionCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeTransferConnectionCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeTransferConnectionCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeTransferConnectionCreateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("NodeA") {
			if !inv.IsParameterNil("NodeA") {
				if inv.Input.NodeA < 0 {
					verr.Add("node_a", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("NodeB") {
			if !inv.IsParameterNil("NodeB") {
				if inv.Input.NodeB < 0 {
					verr.Add("node_b", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeTransferConnectionCreateInvocation) Call() (*ActionNodeTransferConnectionCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNodeTransferConnectionCreateInvocation) callAsBody() (*ActionNodeTransferConnectionCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNodeTransferConnectionCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeTransferConnection
	}
	return resp, err
}

func (inv *ActionNodeTransferConnectionCreateInvocation) makeAllInputParams() *ActionNodeTransferConnectionCreateRequest {
	return &ActionNodeTransferConnectionCreateRequest{
		NodeTransferConnection: inv.makeInputParams(),
		Meta:                   inv.makeMetaInputParams(),
	}
}

func (inv *ActionNodeTransferConnectionCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("NodeA") {
			ret["node_a"] = inv.Input.NodeA
		}
		if inv.IsParameterSelected("NodeAIpAddr") {
			ret["node_a_ip_addr"] = inv.Input.NodeAIpAddr
		}
		if inv.IsParameterSelected("NodeB") {
			ret["node_b"] = inv.Input.NodeB
		}
		if inv.IsParameterSelected("NodeBIpAddr") {
			ret["node_b_ip_addr"] = inv.Input.NodeBIpAddr
		}
	}

	return ret
}

func (inv *ActionNodeTransferConnectionCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
