package client

import (
	"strings"
)

// ActionUserSessionUpdate is a type for action User_session#Update
type ActionUserSessionUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserSessionUpdate(client *Client) *ActionUserSessionUpdate {
	return &ActionUserSessionUpdate{
		Client: client,
	}
}

// ActionUserSessionUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserSessionUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserSessionUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserSessionUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserSessionUpdateMetaGlobalInput) SetNo(value bool) *ActionUserSessionUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserSessionUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSessionUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserSessionUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserSessionUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserSessionUpdateInput is a type for action input parameters
type ActionUserSessionUpdateInput struct {
	Label string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionUserSessionUpdateInput) SetLabel(value string) *ActionUserSessionUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserSessionUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSessionUpdateInput) SelectParameters(params ...string) *ActionUserSessionUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserSessionUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserSessionUpdateInput) UnselectParameters(params ...string) *ActionUserSessionUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserSessionUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserSessionUpdateRequest is a type for the entire action request
type ActionUserSessionUpdateRequest struct {
	UserSession map[string]interface{} `json:"user_session"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionUserSessionUpdateOutput is a type for action output parameters
type ActionUserSessionUpdateOutput struct {
	Admin         *ActionUserShowOutput `json:"admin"`
	ApiIpAddr     string                `json:"api_ip_addr"`
	ApiIpPtr      string                `json:"api_ip_ptr"`
	AuthType      string                `json:"auth_type"`
	ClientIpAddr  string                `json:"client_ip_addr"`
	ClientIpPtr   string                `json:"client_ip_ptr"`
	ClientVersion string                `json:"client_version"`
	ClosedAt      string                `json:"closed_at"`
	CreatedAt     string                `json:"created_at"`
	Id            int64                 `json:"id"`
	Label         string                `json:"label"`
	LastRequestAt string                `json:"last_request_at"`
	RequestCount  int64                 `json:"request_count"`
	Scope         string                `json:"scope"`
	TokenFragment string                `json:"token_fragment"`
	TokenInterval int64                 `json:"token_interval"`
	TokenLifetime string                `json:"token_lifetime"`
	User          *ActionUserShowOutput `json:"user"`
	UserAgent     string                `json:"user_agent"`
}

// Type for action response, including envelope
type ActionUserSessionUpdateResponse struct {
	Action *ActionUserSessionUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserSession *ActionUserSessionUpdateOutput `json:"user_session"`
	}

	// Action output without the namespace
	Output *ActionUserSessionUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserSessionUpdate) Prepare() *ActionUserSessionUpdateInvocation {
	return &ActionUserSessionUpdateInvocation{
		Action: action,
		Path:   "/v6.0/user_sessions/{user_session_id}",
	}
}

// ActionUserSessionUpdateInvocation is used to configure action for invocation
type ActionUserSessionUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserSessionUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserSessionUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserSessionUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserSessionUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserSessionUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserSessionUpdateInvocation) SetPathParamString(param string, value string) *ActionUserSessionUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserSessionUpdateInvocation) NewInput() *ActionUserSessionUpdateInput {
	inv.Input = &ActionUserSessionUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserSessionUpdateInvocation) SetInput(input *ActionUserSessionUpdateInput) *ActionUserSessionUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserSessionUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserSessionUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserSessionUpdateInvocation) NewMetaInput() *ActionUserSessionUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserSessionUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserSessionUpdateInvocation) SetMetaInput(input *ActionUserSessionUpdateMetaGlobalInput) *ActionUserSessionUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserSessionUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserSessionUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserSessionUpdateInvocation) Call() (*ActionUserSessionUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserSessionUpdateInvocation) callAsBody() (*ActionUserSessionUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserSessionUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserSession
	}
	return resp, err
}

func (inv *ActionUserSessionUpdateInvocation) makeAllInputParams() *ActionUserSessionUpdateRequest {
	return &ActionUserSessionUpdateRequest{
		UserSession: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserSessionUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionUserSessionUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
