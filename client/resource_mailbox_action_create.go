package client

import ()

// ActionMailboxCreate is a type for action Mailbox#Create
type ActionMailboxCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxCreate(client *Client) *ActionMailboxCreate {
	return &ActionMailboxCreate{
		Client: client,
	}
}

// ActionMailboxCreateMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxCreateMetaGlobalInput) SetIncludes(value string) *ActionMailboxCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxCreateMetaGlobalInput) SetNo(value bool) *ActionMailboxCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxCreateMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxCreateInput is a type for action input parameters
type ActionMailboxCreateInput struct {
	EnableSsl bool   `json:"enable_ssl"`
	Label     string `json:"label"`
	Password  string `json:"password"`
	Port      int64  `json:"port"`
	Server    string `json:"server"`
	User      string `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnableSsl sets parameter EnableSsl to value and selects it for sending
func (in *ActionMailboxCreateInput) SetEnableSsl(value bool) *ActionMailboxCreateInput {
	in.EnableSsl = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableSsl"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionMailboxCreateInput) SetLabel(value string) *ActionMailboxCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetPassword sets parameter Password to value and selects it for sending
func (in *ActionMailboxCreateInput) SetPassword(value string) *ActionMailboxCreateInput {
	in.Password = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Password"] = nil
	return in
}

// SetPort sets parameter Port to value and selects it for sending
func (in *ActionMailboxCreateInput) SetPort(value int64) *ActionMailboxCreateInput {
	in.Port = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Port"] = nil
	return in
}

// SetServer sets parameter Server to value and selects it for sending
func (in *ActionMailboxCreateInput) SetServer(value string) *ActionMailboxCreateInput {
	in.Server = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Server"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionMailboxCreateInput) SetUser(value string) *ActionMailboxCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxCreateInput) SelectParameters(params ...string) *ActionMailboxCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMailboxCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMailboxCreateInput) UnselectParameters(params ...string) *ActionMailboxCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMailboxCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxCreateRequest is a type for the entire action request
type ActionMailboxCreateRequest struct {
	Mailbox map[string]interface{} `json:"mailbox"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionMailboxCreateOutput is a type for action output parameters
type ActionMailboxCreateOutput struct {
	CreatedAt string `json:"created_at"`
	EnableSsl bool   `json:"enable_ssl"`
	Id        int64  `json:"id"`
	Label     string `json:"label"`
	Port      int64  `json:"port"`
	Server    string `json:"server"`
	UpdatedAt string `json:"updated_at"`
	User      string `json:"user"`
}

// Type for action response, including envelope
type ActionMailboxCreateResponse struct {
	Action *ActionMailboxCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Mailbox *ActionMailboxCreateOutput `json:"mailbox"`
	}

	// Action output without the namespace
	Output *ActionMailboxCreateOutput
}

// Prepare the action for invocation
func (action *ActionMailboxCreate) Prepare() *ActionMailboxCreateInvocation {
	return &ActionMailboxCreateInvocation{
		Action: action,
		Path:   "/v7.0/mailboxes",
	}
}

// ActionMailboxCreateInvocation is used to configure action for invocation
type ActionMailboxCreateInvocation struct {
	// Pointer to the action
	Action *ActionMailboxCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailboxCreateInput
	// Global meta input parameters
	MetaInput *ActionMailboxCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailboxCreateInvocation) NewInput() *ActionMailboxCreateInput {
	inv.Input = &ActionMailboxCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailboxCreateInvocation) SetInput(input *ActionMailboxCreateInput) *ActionMailboxCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailboxCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMailboxCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxCreateInvocation) NewMetaInput() *ActionMailboxCreateMetaGlobalInput {
	inv.MetaInput = &ActionMailboxCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxCreateInvocation) SetMetaInput(input *ActionMailboxCreateMetaGlobalInput) *ActionMailboxCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxCreateInvocation) Call() (*ActionMailboxCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMailboxCreateInvocation) callAsBody() (*ActionMailboxCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailboxCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Mailbox
	}
	return resp, err
}

func (inv *ActionMailboxCreateInvocation) makeAllInputParams() *ActionMailboxCreateRequest {
	return &ActionMailboxCreateRequest{
		Mailbox: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailboxCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("EnableSsl") {
			ret["enable_ssl"] = inv.Input.EnableSsl
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Password") {
			ret["password"] = inv.Input.Password
		}
		if inv.IsParameterSelected("Port") {
			ret["port"] = inv.Input.Port
		}
		if inv.IsParameterSelected("Server") {
			ret["server"] = inv.Input.Server
		}
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
	}

	return ret
}

func (inv *ActionMailboxCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
