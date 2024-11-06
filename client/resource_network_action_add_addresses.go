package client

import (
	"strings"
)

// ActionNetworkAddAddresses is a type for action Network#Add_addresses
type ActionNetworkAddAddresses struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkAddAddresses(client *Client) *ActionNetworkAddAddresses {
	return &ActionNetworkAddAddresses{
		Client: client,
	}
}

// ActionNetworkAddAddressesMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkAddAddressesMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkAddAddressesMetaGlobalInput) SetNo(value bool) *ActionNetworkAddAddressesMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkAddAddressesMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkAddAddressesMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkAddAddressesMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkAddAddressesMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkAddAddressesInput is a type for action input parameters
type ActionNetworkAddAddressesInput struct {
	Count       int64 `json:"count"`
	Environment int64 `json:"environment"`
	User        int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNetworkAddAddressesInput) SetCount(value int64) *ActionNetworkAddAddressesInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionNetworkAddAddressesInput) SetEnvironment(value int64) *ActionNetworkAddAddressesInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionNetworkAddAddressesInput) SetEnvironmentNil(set bool) *ActionNetworkAddAddressesInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Environment"] = nil
		in.SelectParameters("Environment")
	} else {
		delete(in._nilParameters, "Environment")
	}
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionNetworkAddAddressesInput) SetUser(value int64) *ActionNetworkAddAddressesInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionNetworkAddAddressesInput) SetUserNil(set bool) *ActionNetworkAddAddressesInput {
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

// SelectParameters sets parameters from ActionNetworkAddAddressesInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkAddAddressesInput) SelectParameters(params ...string) *ActionNetworkAddAddressesInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNetworkAddAddressesInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNetworkAddAddressesInput) UnselectParameters(params ...string) *ActionNetworkAddAddressesInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNetworkAddAddressesInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkAddAddressesRequest is a type for the entire action request
type ActionNetworkAddAddressesRequest struct {
	Network map[string]interface{} `json:"network"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionNetworkAddAddressesOutput is a type for action output parameters
type ActionNetworkAddAddressesOutput struct {
	Count int64 `json:"count"`
}

// Type for action response, including envelope
type ActionNetworkAddAddressesResponse struct {
	Action *ActionNetworkAddAddresses `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Network *ActionNetworkAddAddressesOutput `json:"network"`
	}

	// Action output without the namespace
	Output *ActionNetworkAddAddressesOutput
}

// Prepare the action for invocation
func (action *ActionNetworkAddAddresses) Prepare() *ActionNetworkAddAddressesInvocation {
	return &ActionNetworkAddAddressesInvocation{
		Action: action,
		Path:   "/v7.0/networks/{network_id}/add_addresses",
	}
}

// ActionNetworkAddAddressesInvocation is used to configure action for invocation
type ActionNetworkAddAddressesInvocation struct {
	// Pointer to the action
	Action *ActionNetworkAddAddresses

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkAddAddressesInput
	// Global meta input parameters
	MetaInput *ActionNetworkAddAddressesMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNetworkAddAddressesInvocation) SetPathParamInt(param string, value int64) *ActionNetworkAddAddressesInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNetworkAddAddressesInvocation) SetPathParamString(param string, value string) *ActionNetworkAddAddressesInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkAddAddressesInvocation) NewInput() *ActionNetworkAddAddressesInput {
	inv.Input = &ActionNetworkAddAddressesInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkAddAddressesInvocation) SetInput(input *ActionNetworkAddAddressesInput) *ActionNetworkAddAddressesInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkAddAddressesInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNetworkAddAddressesInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkAddAddressesInvocation) NewMetaInput() *ActionNetworkAddAddressesMetaGlobalInput {
	inv.MetaInput = &ActionNetworkAddAddressesMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkAddAddressesInvocation) SetMetaInput(input *ActionNetworkAddAddressesMetaGlobalInput) *ActionNetworkAddAddressesInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkAddAddressesInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNetworkAddAddressesInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkAddAddressesInvocation) Call() (*ActionNetworkAddAddressesResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionNetworkAddAddressesInvocation) callAsBody() (*ActionNetworkAddAddressesResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNetworkAddAddressesResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Network
	}
	return resp, err
}

func (inv *ActionNetworkAddAddressesInvocation) makeAllInputParams() *ActionNetworkAddAddressesRequest {
	return &ActionNetworkAddAddressesRequest{
		Network: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionNetworkAddAddressesInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Count") {
			ret["count"] = inv.Input.Count
		}
		if inv.IsParameterSelected("Environment") {
			if inv.IsParameterNil("Environment") {
				ret["environment"] = nil
			} else {
				ret["environment"] = inv.Input.Environment
			}
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

func (inv *ActionNetworkAddAddressesInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
