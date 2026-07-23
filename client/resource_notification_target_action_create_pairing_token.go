package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTargetCreatePairingToken is a type for action Notification_target#Create_pairing_token
type ActionNotificationTargetCreatePairingToken struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetCreatePairingToken(client *Client) *ActionNotificationTargetCreatePairingToken {
	return &ActionNotificationTargetCreatePairingToken{
		Client: client,
	}
}

// ActionNotificationTargetCreatePairingTokenMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetCreatePairingTokenMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetCreatePairingTokenMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetCreatePairingTokenMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetCreatePairingTokenMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetCreatePairingTokenMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetCreatePairingTokenMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetCreatePairingTokenMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetCreatePairingTokenMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetCreatePairingTokenMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetCreatePairingTokenRequest is a type for the entire action request
type ActionNotificationTargetCreatePairingTokenRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTargetCreatePairingTokenOutput is a type for action output parameters
type ActionNotificationTargetCreatePairingTokenOutput struct {
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
type ActionNotificationTargetCreatePairingTokenResponse struct {
	Action *ActionNotificationTargetCreatePairingToken "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTarget *ActionNotificationTargetCreatePairingTokenOutput "json:\"notification_target\""
	}

	// Action output without the namespace
	Output *ActionNotificationTargetCreatePairingTokenOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTargetCreatePairingToken) Prepare() *ActionNotificationTargetCreatePairingTokenInvocation {
	return &ActionNotificationTargetCreatePairingTokenInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets/{notification_target_id}/create_pairing_token",
	}
}

// ActionNotificationTargetCreatePairingTokenInvocation is used to configure action for invocation
type ActionNotificationTargetCreatePairingTokenInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetCreatePairingToken

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationTargetCreatePairingTokenMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTargetCreatePairingTokenInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTargetCreatePairingTokenInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTargetCreatePairingTokenInvocation) SetPathParamString(param string, value string) *ActionNotificationTargetCreatePairingTokenInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetCreatePairingTokenInvocation) NewMetaInput() *ActionNotificationTargetCreatePairingTokenMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetCreatePairingTokenMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetCreatePairingTokenInvocation) SetMetaInput(input *ActionNotificationTargetCreatePairingTokenMetaGlobalInput) *ActionNotificationTargetCreatePairingTokenInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetCreatePairingTokenInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetCreatePairingTokenInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetCreatePairingTokenInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTargetCreatePairingTokenInvocation) Call() (*ActionNotificationTargetCreatePairingTokenResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTargetCreatePairingTokenInvocation) callAsBody() (*ActionNotificationTargetCreatePairingTokenResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTargetCreatePairingTokenResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTarget
	}
	return resp, err
}

func (inv *ActionNotificationTargetCreatePairingTokenInvocation) makeAllInputParams() *ActionNotificationTargetCreatePairingTokenRequest {
	return &ActionNotificationTargetCreatePairingTokenRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTargetCreatePairingTokenInvocation) makeMetaInputParams() map[string]interface{} {
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
