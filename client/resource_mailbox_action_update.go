package client

import (
	"strings"
)

// ActionMailboxUpdate is a type for action Mailbox#Update
type ActionMailboxUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxUpdate(client *Client) *ActionMailboxUpdate {
	return &ActionMailboxUpdate{
		Client: client,
	}
}

// ActionMailboxUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxUpdateMetaGlobalInput) SetIncludes(value string) *ActionMailboxUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxUpdateMetaGlobalInput) SetNo(value bool) *ActionMailboxUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxUpdateInput is a type for action input parameters
type ActionMailboxUpdateInput struct {
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
func (in *ActionMailboxUpdateInput) SetEnableSsl(value bool) *ActionMailboxUpdateInput {
	in.EnableSsl = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableSsl"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionMailboxUpdateInput) SetLabel(value string) *ActionMailboxUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetPassword sets parameter Password to value and selects it for sending
func (in *ActionMailboxUpdateInput) SetPassword(value string) *ActionMailboxUpdateInput {
	in.Password = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Password"] = nil
	return in
}

// SetPort sets parameter Port to value and selects it for sending
func (in *ActionMailboxUpdateInput) SetPort(value int64) *ActionMailboxUpdateInput {
	in.Port = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Port"] = nil
	return in
}

// SetServer sets parameter Server to value and selects it for sending
func (in *ActionMailboxUpdateInput) SetServer(value string) *ActionMailboxUpdateInput {
	in.Server = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Server"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionMailboxUpdateInput) SetUser(value string) *ActionMailboxUpdateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxUpdateInput) SelectParameters(params ...string) *ActionMailboxUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMailboxUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMailboxUpdateInput) UnselectParameters(params ...string) *ActionMailboxUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMailboxUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxUpdateRequest is a type for the entire action request
type ActionMailboxUpdateRequest struct {
	Mailbox map[string]interface{} `json:"mailbox"`
	Meta    map[string]interface{} `json:"_meta"`
}

// ActionMailboxUpdateOutput is a type for action output parameters
type ActionMailboxUpdateOutput struct {
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
type ActionMailboxUpdateResponse struct {
	Action *ActionMailboxUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Mailbox *ActionMailboxUpdateOutput `json:"mailbox"`
	}

	// Action output without the namespace
	Output *ActionMailboxUpdateOutput
}

// Prepare the action for invocation
func (action *ActionMailboxUpdate) Prepare() *ActionMailboxUpdateInvocation {
	return &ActionMailboxUpdateInvocation{
		Action: action,
		Path:   "/v6.0/mailboxes/{mailbox_id}",
	}
}

// ActionMailboxUpdateInvocation is used to configure action for invocation
type ActionMailboxUpdateInvocation struct {
	// Pointer to the action
	Action *ActionMailboxUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMailboxUpdateInput
	// Global meta input parameters
	MetaInput *ActionMailboxUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailboxUpdateInvocation) SetPathParamInt(param string, value int64) *ActionMailboxUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailboxUpdateInvocation) SetPathParamString(param string, value string) *ActionMailboxUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMailboxUpdateInvocation) NewInput() *ActionMailboxUpdateInput {
	inv.Input = &ActionMailboxUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMailboxUpdateInvocation) SetInput(input *ActionMailboxUpdateInput) *ActionMailboxUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMailboxUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMailboxUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxUpdateInvocation) NewMetaInput() *ActionMailboxUpdateMetaGlobalInput {
	inv.MetaInput = &ActionMailboxUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxUpdateInvocation) SetMetaInput(input *ActionMailboxUpdateMetaGlobalInput) *ActionMailboxUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxUpdateInvocation) Call() (*ActionMailboxUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMailboxUpdateInvocation) callAsBody() (*ActionMailboxUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMailboxUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Mailbox
	}
	return resp, err
}

func (inv *ActionMailboxUpdateInvocation) makeAllInputParams() *ActionMailboxUpdateRequest {
	return &ActionMailboxUpdateRequest{
		Mailbox: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionMailboxUpdateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionMailboxUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
