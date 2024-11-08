package client

import (
	"strings"
)

// ActionNetworkUpdate is a type for action Network#Update
type ActionNetworkUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkUpdate(client *Client) *ActionNetworkUpdate {
	return &ActionNetworkUpdate{
		Client: client,
	}
}

// ActionNetworkUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkUpdateMetaGlobalInput) SetIncludes(value string) *ActionNetworkUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkUpdateMetaGlobalInput) SetNo(value bool) *ActionNetworkUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkUpdateInput is a type for action input parameters
type ActionNetworkUpdateInput struct {
	Address     string `json:"address"`
	IpVersion   int64  `json:"ip_version"`
	Label       string `json:"label"`
	Managed     bool   `json:"managed"`
	Prefix      int64  `json:"prefix"`
	Purpose     string `json:"purpose"`
	Role        string `json:"role"`
	SplitAccess string `json:"split_access"`
	SplitPrefix int64  `json:"split_prefix"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionNetworkUpdateInput) SetAddress(value string) *ActionNetworkUpdateInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}

// SetIpVersion sets parameter IpVersion to value and selects it for sending
func (in *ActionNetworkUpdateInput) SetIpVersion(value int64) *ActionNetworkUpdateInput {
	in.IpVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpVersion"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionNetworkUpdateInput) SetLabel(value string) *ActionNetworkUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetManaged sets parameter Managed to value and selects it for sending
func (in *ActionNetworkUpdateInput) SetManaged(value bool) *ActionNetworkUpdateInput {
	in.Managed = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Managed"] = nil
	return in
}

// SetPrefix sets parameter Prefix to value and selects it for sending
func (in *ActionNetworkUpdateInput) SetPrefix(value int64) *ActionNetworkUpdateInput {
	in.Prefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Prefix"] = nil
	return in
}

// SetPurpose sets parameter Purpose to value and selects it for sending
func (in *ActionNetworkUpdateInput) SetPurpose(value string) *ActionNetworkUpdateInput {
	in.Purpose = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Purpose"] = nil
	return in
}

// SetRole sets parameter Role to value and selects it for sending
func (in *ActionNetworkUpdateInput) SetRole(value string) *ActionNetworkUpdateInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}

// SetSplitAccess sets parameter SplitAccess to value and selects it for sending
func (in *ActionNetworkUpdateInput) SetSplitAccess(value string) *ActionNetworkUpdateInput {
	in.SplitAccess = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SplitAccess"] = nil
	return in
}

// SetSplitPrefix sets parameter SplitPrefix to value and selects it for sending
func (in *ActionNetworkUpdateInput) SetSplitPrefix(value int64) *ActionNetworkUpdateInput {
	in.SplitPrefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SplitPrefix"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkUpdateInput) SelectParameters(params ...string) *ActionNetworkUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNetworkUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNetworkUpdateInput) UnselectParameters(params ...string) *ActionNetworkUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNetworkUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkUpdateRequest is a type for the entire action request
type ActionNetworkUpdateRequest struct {
	Network map[string]interface{} `json:"network"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionNetworkUpdateOutput is a type for action output parameters
type ActionNetworkUpdateOutput struct {
	Address         string                    `json:"address"`
	Assigned        int64                     `json:"assigned"`
	Id              int64                     `json:"id"`
	IpVersion       int64                     `json:"ip_version"`
	Label           string                    `json:"label"`
	Managed         bool                      `json:"managed"`
	Owned           int64                     `json:"owned"`
	Prefix          int64                     `json:"prefix"`
	PrimaryLocation *ActionLocationShowOutput `json:"primary_location"`
	Purpose         string                    `json:"purpose"`
	Role            string                    `json:"role"`
	Size            int64                     `json:"size"`
	SplitAccess     string                    `json:"split_access"`
	SplitPrefix     int64                     `json:"split_prefix"`
	Taken           int64                     `json:"taken"`
	Used            int64                     `json:"used"`
}

// Type for action response, including envelope
type ActionNetworkUpdateResponse struct {
	Action *ActionNetworkUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Network *ActionNetworkUpdateOutput `json:"network"`
	}

	// Action output without the namespace
	Output *ActionNetworkUpdateOutput
}

// Prepare the action for invocation
func (action *ActionNetworkUpdate) Prepare() *ActionNetworkUpdateInvocation {
	return &ActionNetworkUpdateInvocation{
		Action: action,
		Path:   "/v7.0/networks/{network_id}",
	}
}

// ActionNetworkUpdateInvocation is used to configure action for invocation
type ActionNetworkUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNetworkUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkUpdateInput
	// Global meta input parameters
	MetaInput *ActionNetworkUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNetworkUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNetworkUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNetworkUpdateInvocation) SetPathParamString(param string, value string) *ActionNetworkUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkUpdateInvocation) NewInput() *ActionNetworkUpdateInput {
	inv.Input = &ActionNetworkUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkUpdateInvocation) SetInput(input *ActionNetworkUpdateInput) *ActionNetworkUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNetworkUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkUpdateInvocation) NewMetaInput() *ActionNetworkUpdateMetaGlobalInput {
	inv.MetaInput = &ActionNetworkUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkUpdateInvocation) SetMetaInput(input *ActionNetworkUpdateMetaGlobalInput) *ActionNetworkUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNetworkUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkUpdateInvocation) Call() (*ActionNetworkUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionNetworkUpdateInvocation) callAsBody() (*ActionNetworkUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNetworkUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Network
	}
	return resp, err
}

func (inv *ActionNetworkUpdateInvocation) makeAllInputParams() *ActionNetworkUpdateRequest {
	return &ActionNetworkUpdateRequest{
		Network: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionNetworkUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Address") {
			ret["address"] = inv.Input.Address
		}
		if inv.IsParameterSelected("IpVersion") {
			ret["ip_version"] = inv.Input.IpVersion
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Managed") {
			ret["managed"] = inv.Input.Managed
		}
		if inv.IsParameterSelected("Prefix") {
			ret["prefix"] = inv.Input.Prefix
		}
		if inv.IsParameterSelected("Purpose") {
			ret["purpose"] = inv.Input.Purpose
		}
		if inv.IsParameterSelected("Role") {
			ret["role"] = inv.Input.Role
		}
		if inv.IsParameterSelected("SplitAccess") {
			ret["split_access"] = inv.Input.SplitAccess
		}
		if inv.IsParameterSelected("SplitPrefix") {
			ret["split_prefix"] = inv.Input.SplitPrefix
		}
	}

	return ret
}

func (inv *ActionNetworkUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
