package client

import (
	"strings"
)

// ActionUserUpdate is a type for action User#Update
type ActionUserUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserUpdate(client *Client) *ActionUserUpdate {
	return &ActionUserUpdate{
		Client: client,
	}
}

// ActionUserUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserUpdateMetaGlobalInput) SetNo(value bool) *ActionUserUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserUpdateInput is a type for action input parameters
type ActionUserUpdateInput struct {
	Address        string `json:"address"`
	ChangeReason   string `json:"change_reason"`
	Email          string `json:"email"`
	ExpirationDate string `json:"expiration_date"`
	FullName       string `json:"full_name"`
	Info           string `json:"info"`
	Language       int64  `json:"language"`
	Level          int64  `json:"level"`
	Lockout        bool   `json:"lockout"`
	Login          string `json:"login"`
	MailerEnabled  bool   `json:"mailer_enabled"`
	NewPassword    string `json:"new_password"`
	ObjectState    string `json:"object_state"`
	Password       string `json:"password"`
	PasswordReset  bool   `json:"password_reset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserUpdateInput) SetAddress(value string) *ActionUserUpdateInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}

// SetChangeReason sets parameter ChangeReason to value and selects it for sending
func (in *ActionUserUpdateInput) SetChangeReason(value string) *ActionUserUpdateInput {
	in.ChangeReason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChangeReason"] = nil
	return in
}

// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionUserUpdateInput) SetEmail(value string) *ActionUserUpdateInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}

// SetExpirationDate sets parameter ExpirationDate to value and selects it for sending
func (in *ActionUserUpdateInput) SetExpirationDate(value string) *ActionUserUpdateInput {
	in.ExpirationDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExpirationDate"] = nil
	return in
}

// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserUpdateInput) SetFullName(value string) *ActionUserUpdateInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
	return in
}

// SetInfo sets parameter Info to value and selects it for sending
func (in *ActionUserUpdateInput) SetInfo(value string) *ActionUserUpdateInput {
	in.Info = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Info"] = nil
	return in
}

// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionUserUpdateInput) SetLanguage(value int64) *ActionUserUpdateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Language"] = nil
	return in
}

// SetLevel sets parameter Level to value and selects it for sending
func (in *ActionUserUpdateInput) SetLevel(value int64) *ActionUserUpdateInput {
	in.Level = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Level"] = nil
	return in
}

// SetLockout sets parameter Lockout to value and selects it for sending
func (in *ActionUserUpdateInput) SetLockout(value bool) *ActionUserUpdateInput {
	in.Lockout = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lockout"] = nil
	return in
}

// SetLogin sets parameter Login to value and selects it for sending
func (in *ActionUserUpdateInput) SetLogin(value string) *ActionUserUpdateInput {
	in.Login = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Login"] = nil
	return in
}

// SetMailerEnabled sets parameter MailerEnabled to value and selects it for sending
func (in *ActionUserUpdateInput) SetMailerEnabled(value bool) *ActionUserUpdateInput {
	in.MailerEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MailerEnabled"] = nil
	return in
}

// SetNewPassword sets parameter NewPassword to value and selects it for sending
func (in *ActionUserUpdateInput) SetNewPassword(value string) *ActionUserUpdateInput {
	in.NewPassword = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NewPassword"] = nil
	return in
}

// SetObjectState sets parameter ObjectState to value and selects it for sending
func (in *ActionUserUpdateInput) SetObjectState(value string) *ActionUserUpdateInput {
	in.ObjectState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectState"] = nil
	return in
}

// SetPassword sets parameter Password to value and selects it for sending
func (in *ActionUserUpdateInput) SetPassword(value string) *ActionUserUpdateInput {
	in.Password = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Password"] = nil
	return in
}

// SetPasswordReset sets parameter PasswordReset to value and selects it for sending
func (in *ActionUserUpdateInput) SetPasswordReset(value bool) *ActionUserUpdateInput {
	in.PasswordReset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PasswordReset"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserUpdateInput) SelectParameters(params ...string) *ActionUserUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserUpdateRequest is a type for the entire action request
type ActionUserUpdateRequest struct {
	User map[string]interface{} `json:"user"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserUpdateOutput is a type for action output parameters
type ActionUserUpdateOutput struct {
	Address        string                    `json:"address"`
	CreatedAt      string                    `json:"created_at"`
	Email          string                    `json:"email"`
	FullName       string                    `json:"full_name"`
	Id             int64                     `json:"id"`
	Info           string                    `json:"info"`
	Language       *ActionLanguageShowOutput `json:"language"`
	LastActivityAt string                    `json:"last_activity_at"`
	Level          int64                     `json:"level"`
	Lockout        bool                      `json:"lockout"`
	Login          string                    `json:"login"`
	MailerEnabled  bool                      `json:"mailer_enabled"`
	PasswordReset  bool                      `json:"password_reset"`
}

// ActionUserUpdateMetaGlobalOutput is a type for global output metadata parameters
type ActionUserUpdateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionUserUpdateResponse struct {
	Action *ActionUserUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		User *ActionUserUpdateOutput `json:"user"`
		// Global output metadata
		Meta *ActionUserUpdateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionUserUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserUpdate) Prepare() *ActionUserUpdateInvocation {
	return &ActionUserUpdateInvocation{
		Action: action,
		Path:   "/v6.0/users/{user_id}",
	}
}

// ActionUserUpdateInvocation is used to configure action for invocation
type ActionUserUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserUpdateInvocation) SetPathParamString(param string, value string) *ActionUserUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserUpdateInvocation) NewInput() *ActionUserUpdateInput {
	inv.Input = &ActionUserUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserUpdateInvocation) SetInput(input *ActionUserUpdateInput) *ActionUserUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserUpdateInvocation) NewMetaInput() *ActionUserUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserUpdateInvocation) SetMetaInput(input *ActionUserUpdateMetaGlobalInput) *ActionUserUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserUpdateInvocation) Call() (*ActionUserUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserUpdateInvocation) callAsBody() (*ActionUserUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.User
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionUserUpdateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionUserUpdateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionUserUpdateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)

	input := req.NewInput()
	input.SetTimeout(timeout)

	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionUserUpdateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
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
			Timeout:  timeout,
			UpdateIn: updateIn,
			Status:   pollResp.Output.Status,
			Current:  pollResp.Output.Current,
			Total:    pollResp.Output.Total,
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
func (resp *ActionUserUpdateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

func (inv *ActionUserUpdateInvocation) makeAllInputParams() *ActionUserUpdateRequest {
	return &ActionUserUpdateRequest{
		User: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Address") {
			ret["address"] = inv.Input.Address
		}
		if inv.IsParameterSelected("ChangeReason") {
			ret["change_reason"] = inv.Input.ChangeReason
		}
		if inv.IsParameterSelected("Email") {
			ret["email"] = inv.Input.Email
		}
		if inv.IsParameterSelected("ExpirationDate") {
			ret["expiration_date"] = inv.Input.ExpirationDate
		}
		if inv.IsParameterSelected("FullName") {
			ret["full_name"] = inv.Input.FullName
		}
		if inv.IsParameterSelected("Info") {
			ret["info"] = inv.Input.Info
		}
		if inv.IsParameterSelected("Language") {
			ret["language"] = inv.Input.Language
		}
		if inv.IsParameterSelected("Level") {
			ret["level"] = inv.Input.Level
		}
		if inv.IsParameterSelected("Lockout") {
			ret["lockout"] = inv.Input.Lockout
		}
		if inv.IsParameterSelected("Login") {
			ret["login"] = inv.Input.Login
		}
		if inv.IsParameterSelected("MailerEnabled") {
			ret["mailer_enabled"] = inv.Input.MailerEnabled
		}
		if inv.IsParameterSelected("NewPassword") {
			ret["new_password"] = inv.Input.NewPassword
		}
		if inv.IsParameterSelected("ObjectState") {
			ret["object_state"] = inv.Input.ObjectState
		}
		if inv.IsParameterSelected("Password") {
			ret["password"] = inv.Input.Password
		}
		if inv.IsParameterSelected("PasswordReset") {
			ret["password_reset"] = inv.Input.PasswordReset
		}
	}

	return ret
}

func (inv *ActionUserUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
