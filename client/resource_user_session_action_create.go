package client

import ()

// ActionUserSessionCreate is a type for action User_session#Create
type ActionUserSessionCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserSessionCreate(client *Client) *ActionUserSessionCreate {
	return &ActionUserSessionCreate{
		Client: client,
	}
}

// ActionUserSessionCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserSessionCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserSessionCreateMetaGlobalInput) SetIncludes(value string) *ActionUserSessionCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserSessionCreateMetaGlobalInput) SetNo(value bool) *ActionUserSessionCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserSessionCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSessionCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserSessionCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserSessionCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserSessionCreateInput is a type for action input parameters
type ActionUserSessionCreateInput struct {
	Label         string `json:"label"`
	Scope         string `json:"scope"`
	TokenInterval int64  `json:"token_interval"`
	TokenLifetime string `json:"token_lifetime"`
	User          int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionUserSessionCreateInput) SetLabel(value string) *ActionUserSessionCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetScope sets parameter Scope to value and selects it for sending
func (in *ActionUserSessionCreateInput) SetScope(value string) *ActionUserSessionCreateInput {
	in.Scope = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Scope"] = nil
	return in
}

// SetTokenInterval sets parameter TokenInterval to value and selects it for sending
func (in *ActionUserSessionCreateInput) SetTokenInterval(value int64) *ActionUserSessionCreateInput {
	in.TokenInterval = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TokenInterval"] = nil
	return in
}

// SetTokenLifetime sets parameter TokenLifetime to value and selects it for sending
func (in *ActionUserSessionCreateInput) SetTokenLifetime(value string) *ActionUserSessionCreateInput {
	in.TokenLifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TokenLifetime"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserSessionCreateInput) SetUser(value int64) *ActionUserSessionCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionUserSessionCreateInput) SetUserNil(set bool) *ActionUserSessionCreateInput {
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

// SelectParameters sets parameters from ActionUserSessionCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSessionCreateInput) SelectParameters(params ...string) *ActionUserSessionCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserSessionCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserSessionCreateInput) UnselectParameters(params ...string) *ActionUserSessionCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserSessionCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserSessionCreateRequest is a type for the entire action request
type ActionUserSessionCreateRequest struct {
	UserSession map[string]interface{} `json:"user_session"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionUserSessionCreateOutput is a type for action output parameters
type ActionUserSessionCreateOutput struct {
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
	TokenFull     string                `json:"token_full"`
	TokenInterval int64                 `json:"token_interval"`
	TokenLifetime string                `json:"token_lifetime"`
	User          *ActionUserShowOutput `json:"user"`
	UserAgent     string                `json:"user_agent"`
}

// Type for action response, including envelope
type ActionUserSessionCreateResponse struct {
	Action *ActionUserSessionCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserSession *ActionUserSessionCreateOutput `json:"user_session"`
	}

	// Action output without the namespace
	Output *ActionUserSessionCreateOutput
}

// Prepare the action for invocation
func (action *ActionUserSessionCreate) Prepare() *ActionUserSessionCreateInvocation {
	return &ActionUserSessionCreateInvocation{
		Action: action,
		Path:   "/v6.0/user_sessions",
	}
}

// ActionUserSessionCreateInvocation is used to configure action for invocation
type ActionUserSessionCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserSessionCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserSessionCreateInput
	// Global meta input parameters
	MetaInput *ActionUserSessionCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserSessionCreateInvocation) NewInput() *ActionUserSessionCreateInput {
	inv.Input = &ActionUserSessionCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserSessionCreateInvocation) SetInput(input *ActionUserSessionCreateInput) *ActionUserSessionCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserSessionCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserSessionCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserSessionCreateInvocation) NewMetaInput() *ActionUserSessionCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserSessionCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserSessionCreateInvocation) SetMetaInput(input *ActionUserSessionCreateMetaGlobalInput) *ActionUserSessionCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserSessionCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserSessionCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserSessionCreateInvocation) Call() (*ActionUserSessionCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserSessionCreateInvocation) callAsBody() (*ActionUserSessionCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserSessionCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserSession
	}
	return resp, err
}

func (inv *ActionUserSessionCreateInvocation) makeAllInputParams() *ActionUserSessionCreateRequest {
	return &ActionUserSessionCreateRequest{
		UserSession: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserSessionCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Scope") {
			ret["scope"] = inv.Input.Scope
		}
		if inv.IsParameterSelected("TokenInterval") {
			ret["token_interval"] = inv.Input.TokenInterval
		}
		if inv.IsParameterSelected("TokenLifetime") {
			ret["token_lifetime"] = inv.Input.TokenLifetime
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

func (inv *ActionUserSessionCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
