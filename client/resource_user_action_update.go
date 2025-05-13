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
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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
	Address                    string `json:"address"`
	ChangeReason               string `json:"change_reason"`
	Email                      string `json:"email"`
	EnableBasicAuth            bool   `json:"enable_basic_auth"`
	EnableMultiFactorAuth      bool   `json:"enable_multi_factor_auth"`
	EnableNewLoginNotification bool   `json:"enable_new_login_notification"`
	EnableOauth2Auth           bool   `json:"enable_oauth2_auth"`
	EnableSingleSignOn         bool   `json:"enable_single_sign_on"`
	EnableTokenAuth            bool   `json:"enable_token_auth"`
	ExpirationDate             string `json:"expiration_date"`
	FullName                   string `json:"full_name"`
	Info                       string `json:"info"`
	Language                   int64  `json:"language"`
	Level                      int64  `json:"level"`
	Lockout                    bool   `json:"lockout"`
	Login                      string `json:"login"`
	LogoutSessions             bool   `json:"logout_sessions"`
	MailerEnabled              bool   `json:"mailer_enabled"`
	NewPassword                string `json:"new_password"`
	ObjectState                string `json:"object_state"`
	Password                   string `json:"password"`
	PasswordReset              bool   `json:"password_reset"`
	PreferredLogoutAll         bool   `json:"preferred_logout_all"`
	PreferredSessionLength     int64  `json:"preferred_session_length"`
	RemindAfterDate            string `json:"remind_after_date"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetEnableBasicAuth sets parameter EnableBasicAuth to value and selects it for sending
func (in *ActionUserUpdateInput) SetEnableBasicAuth(value bool) *ActionUserUpdateInput {
	in.EnableBasicAuth = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableBasicAuth"] = nil
	return in
}

// SetEnableMultiFactorAuth sets parameter EnableMultiFactorAuth to value and selects it for sending
func (in *ActionUserUpdateInput) SetEnableMultiFactorAuth(value bool) *ActionUserUpdateInput {
	in.EnableMultiFactorAuth = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableMultiFactorAuth"] = nil
	return in
}

// SetEnableNewLoginNotification sets parameter EnableNewLoginNotification to value and selects it for sending
func (in *ActionUserUpdateInput) SetEnableNewLoginNotification(value bool) *ActionUserUpdateInput {
	in.EnableNewLoginNotification = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableNewLoginNotification"] = nil
	return in
}

// SetEnableOauth2Auth sets parameter EnableOauth2Auth to value and selects it for sending
func (in *ActionUserUpdateInput) SetEnableOauth2Auth(value bool) *ActionUserUpdateInput {
	in.EnableOauth2Auth = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableOauth2Auth"] = nil
	return in
}

// SetEnableSingleSignOn sets parameter EnableSingleSignOn to value and selects it for sending
func (in *ActionUserUpdateInput) SetEnableSingleSignOn(value bool) *ActionUserUpdateInput {
	in.EnableSingleSignOn = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableSingleSignOn"] = nil
	return in
}

// SetEnableTokenAuth sets parameter EnableTokenAuth to value and selects it for sending
func (in *ActionUserUpdateInput) SetEnableTokenAuth(value bool) *ActionUserUpdateInput {
	in.EnableTokenAuth = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableTokenAuth"] = nil
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

	in.SetLanguageNil(false)
	in._selectedParameters["Language"] = nil
	return in
}

// SetLanguageNil sets parameter Language to nil and selects it for sending
func (in *ActionUserUpdateInput) SetLanguageNil(set bool) *ActionUserUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Language"] = nil
		in.SelectParameters("Language")
	} else {
		delete(in._nilParameters, "Language")
	}
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

// SetLogoutSessions sets parameter LogoutSessions to value and selects it for sending
func (in *ActionUserUpdateInput) SetLogoutSessions(value bool) *ActionUserUpdateInput {
	in.LogoutSessions = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["LogoutSessions"] = nil
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

// SetPreferredLogoutAll sets parameter PreferredLogoutAll to value and selects it for sending
func (in *ActionUserUpdateInput) SetPreferredLogoutAll(value bool) *ActionUserUpdateInput {
	in.PreferredLogoutAll = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PreferredLogoutAll"] = nil
	return in
}

// SetPreferredSessionLength sets parameter PreferredSessionLength to value and selects it for sending
func (in *ActionUserUpdateInput) SetPreferredSessionLength(value int64) *ActionUserUpdateInput {
	in.PreferredSessionLength = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PreferredSessionLength"] = nil
	return in
}

// SetRemindAfterDate sets parameter RemindAfterDate to value and selects it for sending
func (in *ActionUserUpdateInput) SetRemindAfterDate(value string) *ActionUserUpdateInput {
	in.RemindAfterDate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RemindAfterDate"] = nil
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

// UnselectParameters unsets parameters from ActionUserUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserUpdateInput) UnselectParameters(params ...string) *ActionUserUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Address                    string                    `json:"address"`
	CreatedAt                  string                    `json:"created_at"`
	DokuwikiGroups             string                    `json:"dokuwiki_groups"`
	Email                      string                    `json:"email"`
	EnableBasicAuth            bool                      `json:"enable_basic_auth"`
	EnableMultiFactorAuth      bool                      `json:"enable_multi_factor_auth"`
	EnableNewLoginNotification bool                      `json:"enable_new_login_notification"`
	EnableOauth2Auth           bool                      `json:"enable_oauth2_auth"`
	EnableSingleSignOn         bool                      `json:"enable_single_sign_on"`
	EnableTokenAuth            bool                      `json:"enable_token_auth"`
	FullName                   string                    `json:"full_name"`
	Id                         int64                     `json:"id"`
	Info                       string                    `json:"info"`
	Language                   *ActionLanguageShowOutput `json:"language"`
	LastActivityAt             string                    `json:"last_activity_at"`
	Level                      int64                     `json:"level"`
	Lockout                    bool                      `json:"lockout"`
	Login                      string                    `json:"login"`
	MailerEnabled              bool                      `json:"mailer_enabled"`
	PasswordReset              bool                      `json:"password_reset"`
	PreferredLogoutAll         bool                      `json:"preferred_logout_all"`
	PreferredSessionLength     int64                     `json:"preferred_session_length"`
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
		Path:   "/v7.0/users/{user_id}",
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		if inv.IsParameterSelected("EnableBasicAuth") {
			ret["enable_basic_auth"] = inv.Input.EnableBasicAuth
		}
		if inv.IsParameterSelected("EnableMultiFactorAuth") {
			ret["enable_multi_factor_auth"] = inv.Input.EnableMultiFactorAuth
		}
		if inv.IsParameterSelected("EnableNewLoginNotification") {
			ret["enable_new_login_notification"] = inv.Input.EnableNewLoginNotification
		}
		if inv.IsParameterSelected("EnableOauth2Auth") {
			ret["enable_oauth2_auth"] = inv.Input.EnableOauth2Auth
		}
		if inv.IsParameterSelected("EnableSingleSignOn") {
			ret["enable_single_sign_on"] = inv.Input.EnableSingleSignOn
		}
		if inv.IsParameterSelected("EnableTokenAuth") {
			ret["enable_token_auth"] = inv.Input.EnableTokenAuth
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
			if inv.IsParameterNil("Language") {
				ret["language"] = nil
			} else {
				ret["language"] = inv.Input.Language
			}
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
		if inv.IsParameterSelected("LogoutSessions") {
			ret["logout_sessions"] = inv.Input.LogoutSessions
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
		if inv.IsParameterSelected("PreferredLogoutAll") {
			ret["preferred_logout_all"] = inv.Input.PreferredLogoutAll
		}
		if inv.IsParameterSelected("PreferredSessionLength") {
			ret["preferred_session_length"] = inv.Input.PreferredSessionLength
		}
		if inv.IsParameterSelected("RemindAfterDate") {
			ret["remind_after_date"] = inv.Input.RemindAfterDate
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
