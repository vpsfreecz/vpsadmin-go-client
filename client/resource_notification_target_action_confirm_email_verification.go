package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTargetConfirmEmailVerification is a type for action Notification_target#Confirm_email_verification
type ActionNotificationTargetConfirmEmailVerification struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetConfirmEmailVerification(client *Client) *ActionNotificationTargetConfirmEmailVerification {
	return &ActionNotificationTargetConfirmEmailVerification{
		Client: client,
	}
}

// ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetConfirmEmailVerificationInput is a type for action input parameters
type ActionNotificationTargetConfirmEmailVerificationInput struct {
	Token string "json:\"token\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetToken sets parameter Token to value and selects it for sending
func (in *ActionNotificationTargetConfirmEmailVerificationInput) SetToken(value string) *ActionNotificationTargetConfirmEmailVerificationInput {
	in.Token = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Token"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetConfirmEmailVerificationInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetConfirmEmailVerificationInput) SelectParameters(params ...string) *ActionNotificationTargetConfirmEmailVerificationInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTargetConfirmEmailVerificationInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTargetConfirmEmailVerificationInput) UnselectParameters(params ...string) *ActionNotificationTargetConfirmEmailVerificationInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTargetConfirmEmailVerificationInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetConfirmEmailVerificationRequest is a type for the entire action request
type ActionNotificationTargetConfirmEmailVerificationRequest struct {
	NotificationTarget map[string]interface{} "json:\"notification_target\""
	Meta               map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTargetConfirmEmailVerificationOutput is a type for action output parameters
type ActionNotificationTargetConfirmEmailVerificationOutput struct {
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
type ActionNotificationTargetConfirmEmailVerificationResponse struct {
	Action *ActionNotificationTargetConfirmEmailVerification "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTarget *ActionNotificationTargetConfirmEmailVerificationOutput "json:\"notification_target\""
	}

	// Action output without the namespace
	Output *ActionNotificationTargetConfirmEmailVerificationOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTargetConfirmEmailVerification) Prepare() *ActionNotificationTargetConfirmEmailVerificationInvocation {
	return &ActionNotificationTargetConfirmEmailVerificationInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets/{notification_target_id}/confirm_email_verification",
	}
}

// ActionNotificationTargetConfirmEmailVerificationInvocation is used to configure action for invocation
type ActionNotificationTargetConfirmEmailVerificationInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetConfirmEmailVerification

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTargetConfirmEmailVerificationInput
	// Global meta input parameters
	MetaInput *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTargetConfirmEmailVerificationInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) SetPathParamString(param string, value string) *ActionNotificationTargetConfirmEmailVerificationInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) NewInput() *ActionNotificationTargetConfirmEmailVerificationInput {
	inv.Input = &ActionNotificationTargetConfirmEmailVerificationInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) SetInput(input *ActionNotificationTargetConfirmEmailVerificationInput) *ActionNotificationTargetConfirmEmailVerificationInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) NewMetaInput() *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) SetMetaInput(input *ActionNotificationTargetConfirmEmailVerificationMetaGlobalInput) *ActionNotificationTargetConfirmEmailVerificationInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) validate() error {
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
func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) Call() (*ActionNotificationTargetConfirmEmailVerificationResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) callAsBody() (*ActionNotificationTargetConfirmEmailVerificationResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTargetConfirmEmailVerificationResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTarget
	}
	return resp, err
}

func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) makeAllInputParams() *ActionNotificationTargetConfirmEmailVerificationRequest {
	return &ActionNotificationTargetConfirmEmailVerificationRequest{
		NotificationTarget: inv.makeInputParams(),
		Meta:               inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Token") {
			ret["token"] = inv.Input.Token
		}
	}

	return ret
}

func (inv *ActionNotificationTargetConfirmEmailVerificationInvocation) makeMetaInputParams() map[string]interface{} {
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
