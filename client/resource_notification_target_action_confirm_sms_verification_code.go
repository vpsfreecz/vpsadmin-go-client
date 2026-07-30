package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTargetConfirmSmsVerificationCode is a type for action Notification_target#Confirm_sms_verification_code
type ActionNotificationTargetConfirmSmsVerificationCode struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetConfirmSmsVerificationCode(client *Client) *ActionNotificationTargetConfirmSmsVerificationCode {
	return &ActionNotificationTargetConfirmSmsVerificationCode{
		Client: client,
	}
}

// ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetConfirmSmsVerificationCodeInput is a type for action input parameters
type ActionNotificationTargetConfirmSmsVerificationCodeInput struct {
	Code string "json:\"code\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCode sets parameter Code to value and selects it for sending
func (in *ActionNotificationTargetConfirmSmsVerificationCodeInput) SetCode(value string) *ActionNotificationTargetConfirmSmsVerificationCodeInput {
	in.Code = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Code"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetConfirmSmsVerificationCodeInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetConfirmSmsVerificationCodeInput) SelectParameters(params ...string) *ActionNotificationTargetConfirmSmsVerificationCodeInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTargetConfirmSmsVerificationCodeInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTargetConfirmSmsVerificationCodeInput) UnselectParameters(params ...string) *ActionNotificationTargetConfirmSmsVerificationCodeInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTargetConfirmSmsVerificationCodeInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetConfirmSmsVerificationCodeRequest is a type for the entire action request
type ActionNotificationTargetConfirmSmsVerificationCodeRequest struct {
	NotificationTarget map[string]interface{} "json:\"notification_target\""
	Meta               map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTargetConfirmSmsVerificationCodeOutput is a type for action output parameters
type ActionNotificationTargetConfirmSmsVerificationCodeOutput struct {
	Action                 string                "json:\"action\""
	ConfigJson             string                "json:\"config_json\""
	CreatedAt              string                "json:\"created_at\""
	DeliveryMethodEnabled  bool                  "json:\"delivery_method_enabled\""
	DisplayTarget          string                "json:\"display_target\""
	Enabled                bool                  "json:\"enabled\""
	Id                     int64                 "json:\"id\""
	Label                  string                "json:\"label\""
	LastError              string                "json:\"last_error\""
	SecretPresent          bool                  "json:\"secret_present\""
	TargetKind             string                "json:\"target_kind\""
	TargetValue            string                "json:\"target_value\""
	TelegramBotName        string                "json:\"telegram_bot_name\""
	TelegramBotUrl         string                "json:\"telegram_bot_url\""
	TelegramPairingCommand string                "json:\"telegram_pairing_command\""
	TelegramPairingUrl     string                "json:\"telegram_pairing_url\""
	UpdatedAt              string                "json:\"updated_at\""
	User                   *ActionUserShowOutput "json:\"user\""
	VerificationToken      string                "json:\"verification_token\""
	Verified               bool                  "json:\"verified\""
	VerifiedAt             string                "json:\"verified_at\""
}

// Type for action response, including envelope
type ActionNotificationTargetConfirmSmsVerificationCodeResponse struct {
	Action *ActionNotificationTargetConfirmSmsVerificationCode "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTarget *ActionNotificationTargetConfirmSmsVerificationCodeOutput "json:\"notification_target\""
	}

	// Action output without the namespace
	Output *ActionNotificationTargetConfirmSmsVerificationCodeOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTargetConfirmSmsVerificationCode) Prepare() *ActionNotificationTargetConfirmSmsVerificationCodeInvocation {
	return &ActionNotificationTargetConfirmSmsVerificationCodeInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets/{notification_target_id}/confirm_sms_verification_code",
	}
}

// ActionNotificationTargetConfirmSmsVerificationCodeInvocation is used to configure action for invocation
type ActionNotificationTargetConfirmSmsVerificationCodeInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetConfirmSmsVerificationCode

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTargetConfirmSmsVerificationCodeInput
	// Global meta input parameters
	MetaInput *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTargetConfirmSmsVerificationCodeInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) SetPathParamString(param string, value string) *ActionNotificationTargetConfirmSmsVerificationCodeInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) NewInput() *ActionNotificationTargetConfirmSmsVerificationCodeInput {
	inv.Input = &ActionNotificationTargetConfirmSmsVerificationCodeInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) SetInput(input *ActionNotificationTargetConfirmSmsVerificationCodeInput) *ActionNotificationTargetConfirmSmsVerificationCodeInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) NewMetaInput() *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) SetMetaInput(input *ActionNotificationTargetConfirmSmsVerificationCodeMetaGlobalInput) *ActionNotificationTargetConfirmSmsVerificationCodeInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) Call() (*ActionNotificationTargetConfirmSmsVerificationCodeResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) callAsBody() (*ActionNotificationTargetConfirmSmsVerificationCodeResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTargetConfirmSmsVerificationCodeResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTarget
	}
	return resp, err
}

func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) makeAllInputParams() *ActionNotificationTargetConfirmSmsVerificationCodeRequest {
	return &ActionNotificationTargetConfirmSmsVerificationCodeRequest{
		NotificationTarget: inv.makeInputParams(),
		Meta:               inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Code") {
			ret["code"] = inv.Input.Code
		}
	}

	return ret
}

func (inv *ActionNotificationTargetConfirmSmsVerificationCodeInvocation) makeMetaInputParams() map[string]interface{} {
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
