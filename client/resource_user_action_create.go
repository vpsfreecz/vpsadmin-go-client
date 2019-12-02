package client

import (
)

// ActionUserCreate is a type for action User#Create
type ActionUserCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserCreate(client *Client) *ActionUserCreate {
	return &ActionUserCreate{
		Client: client,
	}
}

// ActionUserCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserCreateMetaGlobalInput) SetNo(value bool) *ActionUserCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserCreateMetaGlobalInput) SetIncludes(value string) *ActionUserCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserCreateInput is a type for action input parameters
type ActionUserCreateInput struct {
	Login string `json:"login"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
	Level int64 `json:"level"`
	Info string `json:"info"`
	MailerEnabled bool `json:"mailer_enabled"`
	PasswordReset bool `json:"password_reset"`
	Lockout bool `json:"lockout"`
	Language int64 `json:"language"`
	Password string `json:"password"`
	Vps bool `json:"vps"`
	Node int64 `json:"node"`
	Environment int64 `json:"environment"`
	Location int64 `json:"location"`
	OsTemplate int64 `json:"os_template"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLogin sets parameter Login to value and selects it for sending
func (in *ActionUserCreateInput) SetLogin(value string) *ActionUserCreateInput {
	in.Login = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Login"] = nil
	return in
}
// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserCreateInput) SetFullName(value string) *ActionUserCreateInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
	return in
}
// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionUserCreateInput) SetEmail(value string) *ActionUserCreateInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}
// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserCreateInput) SetAddress(value string) *ActionUserCreateInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}
// SetLevel sets parameter Level to value and selects it for sending
func (in *ActionUserCreateInput) SetLevel(value int64) *ActionUserCreateInput {
	in.Level = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Level"] = nil
	return in
}
// SetInfo sets parameter Info to value and selects it for sending
func (in *ActionUserCreateInput) SetInfo(value string) *ActionUserCreateInput {
	in.Info = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Info"] = nil
	return in
}
// SetMailerEnabled sets parameter MailerEnabled to value and selects it for sending
func (in *ActionUserCreateInput) SetMailerEnabled(value bool) *ActionUserCreateInput {
	in.MailerEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MailerEnabled"] = nil
	return in
}
// SetPasswordReset sets parameter PasswordReset to value and selects it for sending
func (in *ActionUserCreateInput) SetPasswordReset(value bool) *ActionUserCreateInput {
	in.PasswordReset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PasswordReset"] = nil
	return in
}
// SetLockout sets parameter Lockout to value and selects it for sending
func (in *ActionUserCreateInput) SetLockout(value bool) *ActionUserCreateInput {
	in.Lockout = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lockout"] = nil
	return in
}
// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionUserCreateInput) SetLanguage(value int64) *ActionUserCreateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Language"] = nil
	return in
}
// SetPassword sets parameter Password to value and selects it for sending
func (in *ActionUserCreateInput) SetPassword(value string) *ActionUserCreateInput {
	in.Password = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Password"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionUserCreateInput) SetVps(value bool) *ActionUserCreateInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionUserCreateInput) SetNode(value int64) *ActionUserCreateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionUserCreateInput) SetEnvironment(value int64) *ActionUserCreateInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionUserCreateInput) SetLocation(value int64) *ActionUserCreateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionUserCreateInput) SetOsTemplate(value int64) *ActionUserCreateInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OsTemplate"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserCreateInput) SelectParameters(params ...string) *ActionUserCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserCreateRequest is a type for the entire action request
type ActionUserCreateRequest struct {
	User map[string]interface{} `json:"user"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserCreateOutput is a type for action output parameters
type ActionUserCreateOutput struct {
	Id int64 `json:"id"`
	Login string `json:"login"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
	Level int64 `json:"level"`
	Info string `json:"info"`
	MailerEnabled bool `json:"mailer_enabled"`
	PasswordReset bool `json:"password_reset"`
	Lockout bool `json:"lockout"`
	Language *ActionLanguageShowOutput `json:"language"`
	LastActivityAt string `json:"last_activity_at"`
	CreatedAt string `json:"created_at"`
	ObjectState string `json:"object_state"`
	ExpirationDate string `json:"expiration_date"`
}

// ActionUserCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionUserCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionUserCreateResponse struct {
	Action *ActionUserCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		User *ActionUserCreateOutput `json:"user"`
		// Global output metadata
		Meta *ActionUserCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionUserCreateOutput
}


// Prepare the action for invocation
func (action *ActionUserCreate) Prepare() *ActionUserCreateInvocation {
	return &ActionUserCreateInvocation{
		Action: action,
		Path: "/v5.0/users",
	}
}

// ActionUserCreateInvocation is used to configure action for invocation
type ActionUserCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserCreateInput
	// Global meta input parameters
	MetaInput *ActionUserCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserCreateInvocation) NewInput() *ActionUserCreateInput {
	inv.Input = &ActionUserCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserCreateInvocation) SetInput(input *ActionUserCreateInput) *ActionUserCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserCreateInvocation) NewMetaInput() *ActionUserCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserCreateInvocation) SetMetaInput(input *ActionUserCreateMetaGlobalInput) *ActionUserCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserCreateInvocation) Call() (*ActionUserCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserCreateInvocation) callAsBody() (*ActionUserCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.User
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionUserCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionUserCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionUserCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionUserCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
			Timeout: timeout,
			UpdateIn: updateIn,
			Status: pollResp.Output.Status,
			Current: pollResp.Output.Current,
			Total: pollResp.Output.Total,
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
func (resp *ActionUserCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionUserCreateInvocation) makeAllInputParams() *ActionUserCreateRequest {
	return &ActionUserCreateRequest{
		User: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Login") {
			ret["login"] = inv.Input.Login
		}
		if inv.IsParameterSelected("FullName") {
			ret["full_name"] = inv.Input.FullName
		}
		if inv.IsParameterSelected("Email") {
			ret["email"] = inv.Input.Email
		}
		if inv.IsParameterSelected("Address") {
			ret["address"] = inv.Input.Address
		}
		if inv.IsParameterSelected("Level") {
			ret["level"] = inv.Input.Level
		}
		if inv.IsParameterSelected("Info") {
			ret["info"] = inv.Input.Info
		}
		if inv.IsParameterSelected("MailerEnabled") {
			ret["mailer_enabled"] = inv.Input.MailerEnabled
		}
		if inv.IsParameterSelected("PasswordReset") {
			ret["password_reset"] = inv.Input.PasswordReset
		}
		if inv.IsParameterSelected("Lockout") {
			ret["lockout"] = inv.Input.Lockout
		}
		if inv.IsParameterSelected("Language") {
			ret["language"] = inv.Input.Language
		}
		if inv.IsParameterSelected("Password") {
			ret["password"] = inv.Input.Password
		}
		if inv.IsParameterSelected("Vps") {
			ret["vps"] = inv.Input.Vps
		}
		if inv.IsParameterSelected("Node") {
			ret["node"] = inv.Input.Node
		}
		if inv.IsParameterSelected("Environment") {
			ret["environment"] = inv.Input.Environment
		}
		if inv.IsParameterSelected("Location") {
			ret["location"] = inv.Input.Location
		}
		if inv.IsParameterSelected("OsTemplate") {
			ret["os_template"] = inv.Input.OsTemplate
		}
	}

	return ret
}

func (inv *ActionUserCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
