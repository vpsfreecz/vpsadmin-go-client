package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTargetSendSmsVerificationCode is a type for action Notification_target#Send_sms_verification_code
type ActionNotificationTargetSendSmsVerificationCode struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetSendSmsVerificationCode(client *Client) *ActionNotificationTargetSendSmsVerificationCode {
	return &ActionNotificationTargetSendSmsVerificationCode{
		Client: client,
	}
}

// ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetSendSmsVerificationCodeRequest is a type for the entire action request
type ActionNotificationTargetSendSmsVerificationCodeRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTargetSendSmsVerificationCodeOutput is a type for action output parameters
type ActionNotificationTargetSendSmsVerificationCodeOutput struct {
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
type ActionNotificationTargetSendSmsVerificationCodeResponse struct {
	Action *ActionNotificationTargetSendSmsVerificationCode "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTarget *ActionNotificationTargetSendSmsVerificationCodeOutput "json:\"notification_target\""
	}

	// Action output without the namespace
	Output *ActionNotificationTargetSendSmsVerificationCodeOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTargetSendSmsVerificationCode) Prepare() *ActionNotificationTargetSendSmsVerificationCodeInvocation {
	return &ActionNotificationTargetSendSmsVerificationCodeInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets/{notification_target_id}/send_sms_verification_code",
	}
}

// ActionNotificationTargetSendSmsVerificationCodeInvocation is used to configure action for invocation
type ActionNotificationTargetSendSmsVerificationCodeInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetSendSmsVerificationCode

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTargetSendSmsVerificationCodeInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) SetPathParamString(param string, value string) *ActionNotificationTargetSendSmsVerificationCodeInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) NewMetaInput() *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) SetMetaInput(input *ActionNotificationTargetSendSmsVerificationCodeMetaGlobalInput) *ActionNotificationTargetSendSmsVerificationCodeInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) Call() (*ActionNotificationTargetSendSmsVerificationCodeResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) callAsBody() (*ActionNotificationTargetSendSmsVerificationCodeResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTargetSendSmsVerificationCodeResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTarget
	}
	return resp, err
}

func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) makeAllInputParams() *ActionNotificationTargetSendSmsVerificationCodeRequest {
	return &ActionNotificationTargetSendSmsVerificationCodeRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTargetSendSmsVerificationCodeInvocation) makeMetaInputParams() map[string]interface{} {
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
