package client

import ()

// ActionSessionTokenCreate is a type for action Session_token#Create
type ActionSessionTokenCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionSessionTokenCreate(client *Client) *ActionSessionTokenCreate {
	return &ActionSessionTokenCreate{
		Client: client,
	}
}

// ActionSessionTokenCreateMetaGlobalInput is a type for action global meta input parameters
type ActionSessionTokenCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSessionTokenCreateMetaGlobalInput) SetIncludes(value string) *ActionSessionTokenCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSessionTokenCreateMetaGlobalInput) SetNo(value bool) *ActionSessionTokenCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSessionTokenCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenCreateMetaGlobalInput) SelectParameters(params ...string) *ActionSessionTokenCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSessionTokenCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSessionTokenCreateInput is a type for action input parameters
type ActionSessionTokenCreateInput struct {
	Interval int64  `json:"interval"`
	Label    string `json:"label"`
	Lifetime string `json:"lifetime"`
	User     int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetInterval sets parameter Interval to value and selects it for sending
func (in *ActionSessionTokenCreateInput) SetInterval(value int64) *ActionSessionTokenCreateInput {
	in.Interval = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Interval"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionSessionTokenCreateInput) SetLabel(value string) *ActionSessionTokenCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetLifetime sets parameter Lifetime to value and selects it for sending
func (in *ActionSessionTokenCreateInput) SetLifetime(value string) *ActionSessionTokenCreateInput {
	in.Lifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lifetime"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionSessionTokenCreateInput) SetUser(value int64) *ActionSessionTokenCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionSessionTokenCreateInput) SetUserNil(set bool) *ActionSessionTokenCreateInput {
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

// SelectParameters sets parameters from ActionSessionTokenCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenCreateInput) SelectParameters(params ...string) *ActionSessionTokenCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSessionTokenCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSessionTokenCreateInput) UnselectParameters(params ...string) *ActionSessionTokenCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSessionTokenCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSessionTokenCreateRequest is a type for the entire action request
type ActionSessionTokenCreateRequest struct {
	SessionToken map[string]interface{} `json:"session_token"`
	Meta         map[string]interface{} `json:"_meta"`
}

// ActionSessionTokenCreateOutput is a type for action output parameters
type ActionSessionTokenCreateOutput struct {
	CreatedAt string                `json:"created_at"`
	Id        int64                 `json:"id"`
	Interval  int64                 `json:"interval"`
	Label     string                `json:"label"`
	Lifetime  string                `json:"lifetime"`
	Token     string                `json:"token"`
	UseCount  int64                 `json:"use_count"`
	User      *ActionUserShowOutput `json:"user"`
	ValidTo   string                `json:"valid_to"`
}

// Type for action response, including envelope
type ActionSessionTokenCreateResponse struct {
	Action *ActionSessionTokenCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SessionToken *ActionSessionTokenCreateOutput `json:"session_token"`
	}

	// Action output without the namespace
	Output *ActionSessionTokenCreateOutput
}

// Prepare the action for invocation
func (action *ActionSessionTokenCreate) Prepare() *ActionSessionTokenCreateInvocation {
	return &ActionSessionTokenCreateInvocation{
		Action: action,
		Path:   "/v6.0/session_tokens",
	}
}

// ActionSessionTokenCreateInvocation is used to configure action for invocation
type ActionSessionTokenCreateInvocation struct {
	// Pointer to the action
	Action *ActionSessionTokenCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSessionTokenCreateInput
	// Global meta input parameters
	MetaInput *ActionSessionTokenCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSessionTokenCreateInvocation) NewInput() *ActionSessionTokenCreateInput {
	inv.Input = &ActionSessionTokenCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSessionTokenCreateInvocation) SetInput(input *ActionSessionTokenCreateInput) *ActionSessionTokenCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSessionTokenCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSessionTokenCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSessionTokenCreateInvocation) NewMetaInput() *ActionSessionTokenCreateMetaGlobalInput {
	inv.MetaInput = &ActionSessionTokenCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSessionTokenCreateInvocation) SetMetaInput(input *ActionSessionTokenCreateMetaGlobalInput) *ActionSessionTokenCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSessionTokenCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSessionTokenCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSessionTokenCreateInvocation) Call() (*ActionSessionTokenCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionSessionTokenCreateInvocation) callAsBody() (*ActionSessionTokenCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSessionTokenCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SessionToken
	}
	return resp, err
}

func (inv *ActionSessionTokenCreateInvocation) makeAllInputParams() *ActionSessionTokenCreateRequest {
	return &ActionSessionTokenCreateRequest{
		SessionToken: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionSessionTokenCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Interval") {
			ret["interval"] = inv.Input.Interval
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Lifetime") {
			ret["lifetime"] = inv.Input.Lifetime
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

func (inv *ActionSessionTokenCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
