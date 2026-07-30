package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTargetSendEmailVerification is a type for action Notification_target#Send_email_verification
type ActionNotificationTargetSendEmailVerification struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetSendEmailVerification(client *Client) *ActionNotificationTargetSendEmailVerification {
	return &ActionNotificationTargetSendEmailVerification{
		Client: client,
	}
}

// ActionNotificationTargetSendEmailVerificationMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetSendEmailVerificationMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetSendEmailVerificationMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetSendEmailVerificationMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetSendEmailVerificationMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetSendEmailVerificationMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetSendEmailVerificationMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetSendEmailVerificationMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetSendEmailVerificationMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetSendEmailVerificationMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetSendEmailVerificationRequest is a type for the entire action request
type ActionNotificationTargetSendEmailVerificationRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTargetSendEmailVerificationOutput is a type for action output parameters
type ActionNotificationTargetSendEmailVerificationOutput struct {
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
type ActionNotificationTargetSendEmailVerificationResponse struct {
	Action *ActionNotificationTargetSendEmailVerification "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTarget *ActionNotificationTargetSendEmailVerificationOutput "json:\"notification_target\""
	}

	// Action output without the namespace
	Output *ActionNotificationTargetSendEmailVerificationOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTargetSendEmailVerification) Prepare() *ActionNotificationTargetSendEmailVerificationInvocation {
	return &ActionNotificationTargetSendEmailVerificationInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets/{notification_target_id}/send_email_verification",
	}
}

// ActionNotificationTargetSendEmailVerificationInvocation is used to configure action for invocation
type ActionNotificationTargetSendEmailVerificationInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetSendEmailVerification

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationTargetSendEmailVerificationMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTargetSendEmailVerificationInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTargetSendEmailVerificationInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTargetSendEmailVerificationInvocation) SetPathParamString(param string, value string) *ActionNotificationTargetSendEmailVerificationInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetSendEmailVerificationInvocation) NewMetaInput() *ActionNotificationTargetSendEmailVerificationMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetSendEmailVerificationMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetSendEmailVerificationInvocation) SetMetaInput(input *ActionNotificationTargetSendEmailVerificationMetaGlobalInput) *ActionNotificationTargetSendEmailVerificationInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetSendEmailVerificationInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetSendEmailVerificationInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetSendEmailVerificationInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTargetSendEmailVerificationInvocation) Call() (*ActionNotificationTargetSendEmailVerificationResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTargetSendEmailVerificationInvocation) callAsBody() (*ActionNotificationTargetSendEmailVerificationResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTargetSendEmailVerificationResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTarget
	}
	return resp, err
}

func (inv *ActionNotificationTargetSendEmailVerificationInvocation) makeAllInputParams() *ActionNotificationTargetSendEmailVerificationRequest {
	return &ActionNotificationTargetSendEmailVerificationRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTargetSendEmailVerificationInvocation) makeMetaInputParams() map[string]interface{} {
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
