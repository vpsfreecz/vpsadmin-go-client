package client

import (
	"strings"
)

// ActionNodeTransferConnectionUpdate is a type for action Node_transfer_connection#Update
type ActionNodeTransferConnectionUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeTransferConnectionUpdate(client *Client) *ActionNodeTransferConnectionUpdate {
	return &ActionNodeTransferConnectionUpdate{
		Client: client,
	}
}

// ActionNodeTransferConnectionUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNodeTransferConnectionUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeTransferConnectionUpdateMetaGlobalInput) SetIncludes(value string) *ActionNodeTransferConnectionUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeTransferConnectionUpdateMetaGlobalInput) SetNo(value bool) *ActionNodeTransferConnectionUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeTransferConnectionUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNodeTransferConnectionUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeTransferConnectionUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeTransferConnectionUpdateInput is a type for action input parameters
type ActionNodeTransferConnectionUpdateInput struct {
	Enabled     bool   `json:"enabled"`
	NodeAIpAddr string `json:"node_a_ip_addr"`
	NodeBIpAddr string `json:"node_b_ip_addr"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionNodeTransferConnectionUpdateInput) SetEnabled(value bool) *ActionNodeTransferConnectionUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetNodeAIpAddr sets parameter NodeAIpAddr to value and selects it for sending
func (in *ActionNodeTransferConnectionUpdateInput) SetNodeAIpAddr(value string) *ActionNodeTransferConnectionUpdateInput {
	in.NodeAIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeAIpAddr"] = nil
	return in
}

// SetNodeBIpAddr sets parameter NodeBIpAddr to value and selects it for sending
func (in *ActionNodeTransferConnectionUpdateInput) SetNodeBIpAddr(value string) *ActionNodeTransferConnectionUpdateInput {
	in.NodeBIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeBIpAddr"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeTransferConnectionUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionUpdateInput) SelectParameters(params ...string) *ActionNodeTransferConnectionUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeTransferConnectionUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionUpdateInput) UnselectParameters(params ...string) *ActionNodeTransferConnectionUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeTransferConnectionUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeTransferConnectionUpdateRequest is a type for the entire action request
type ActionNodeTransferConnectionUpdateRequest struct {
	NodeTransferConnection map[string]interface{} `json:"node_transfer_connection"`
	Meta                   map[string]interface{} `json:"_meta"`
}

// ActionNodeTransferConnectionUpdateOutput is a type for action output parameters
type ActionNodeTransferConnectionUpdateOutput struct {
	CreatedAt   string                `json:"created_at"`
	Enabled     bool                  `json:"enabled"`
	Id          int64                 `json:"id"`
	NodeA       *ActionNodeShowOutput `json:"node_a"`
	NodeAIpAddr string                `json:"node_a_ip_addr"`
	NodeB       *ActionNodeShowOutput `json:"node_b"`
	NodeBIpAddr string                `json:"node_b_ip_addr"`
	UpdatedAt   string                `json:"updated_at"`
}

// Type for action response, including envelope
type ActionNodeTransferConnectionUpdateResponse struct {
	Action *ActionNodeTransferConnectionUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeTransferConnection *ActionNodeTransferConnectionUpdateOutput `json:"node_transfer_connection"`
	}

	// Action output without the namespace
	Output *ActionNodeTransferConnectionUpdateOutput
}

// Prepare the action for invocation
func (action *ActionNodeTransferConnectionUpdate) Prepare() *ActionNodeTransferConnectionUpdateInvocation {
	return &ActionNodeTransferConnectionUpdateInvocation{
		Action: action,
		Path:   "/v7.0/node_transfer_connections/{node_transfer_connection_id}",
	}
}

// ActionNodeTransferConnectionUpdateInvocation is used to configure action for invocation
type ActionNodeTransferConnectionUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNodeTransferConnectionUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeTransferConnectionUpdateInput
	// Global meta input parameters
	MetaInput *ActionNodeTransferConnectionUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeTransferConnectionUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNodeTransferConnectionUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeTransferConnectionUpdateInvocation) SetPathParamString(param string, value string) *ActionNodeTransferConnectionUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeTransferConnectionUpdateInvocation) NewInput() *ActionNodeTransferConnectionUpdateInput {
	inv.Input = &ActionNodeTransferConnectionUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeTransferConnectionUpdateInvocation) SetInput(input *ActionNodeTransferConnectionUpdateInput) *ActionNodeTransferConnectionUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeTransferConnectionUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeTransferConnectionUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeTransferConnectionUpdateInvocation) NewMetaInput() *ActionNodeTransferConnectionUpdateMetaGlobalInput {
	inv.MetaInput = &ActionNodeTransferConnectionUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeTransferConnectionUpdateInvocation) SetMetaInput(input *ActionNodeTransferConnectionUpdateMetaGlobalInput) *ActionNodeTransferConnectionUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeTransferConnectionUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeTransferConnectionUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeTransferConnectionUpdateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeTransferConnectionUpdateInvocation) Call() (*ActionNodeTransferConnectionUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNodeTransferConnectionUpdateInvocation) callAsBody() (*ActionNodeTransferConnectionUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNodeTransferConnectionUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeTransferConnection
	}
	return resp, err
}

func (inv *ActionNodeTransferConnectionUpdateInvocation) makeAllInputParams() *ActionNodeTransferConnectionUpdateRequest {
	return &ActionNodeTransferConnectionUpdateRequest{
		NodeTransferConnection: inv.makeInputParams(),
		Meta:                   inv.makeMetaInputParams(),
	}
}

func (inv *ActionNodeTransferConnectionUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("NodeAIpAddr") {
			ret["node_a_ip_addr"] = inv.Input.NodeAIpAddr
		}
		if inv.IsParameterSelected("NodeBIpAddr") {
			ret["node_b_ip_addr"] = inv.Input.NodeBIpAddr
		}
	}

	return ret
}

func (inv *ActionNodeTransferConnectionUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
