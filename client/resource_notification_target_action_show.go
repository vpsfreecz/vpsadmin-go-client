package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTargetShow is a type for action Notification_target#Show
type ActionNotificationTargetShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetShow(client *Client) *ActionNotificationTargetShow {
	return &ActionNotificationTargetShow{
		Client: client,
	}
}

// ActionNotificationTargetShowMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetShowMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetShowMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetShowMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetShowOutput is a type for action output parameters
type ActionNotificationTargetShowOutput struct {
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
type ActionNotificationTargetShowResponse struct {
	Action *ActionNotificationTargetShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTarget *ActionNotificationTargetShowOutput "json:\"notification_target\""
	}

	// Action output without the namespace
	Output *ActionNotificationTargetShowOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTargetShow) Prepare() *ActionNotificationTargetShowInvocation {
	return &ActionNotificationTargetShowInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets/{notification_target_id}",
	}
}

// ActionNotificationTargetShowInvocation is used to configure action for invocation
type ActionNotificationTargetShowInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationTargetShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTargetShowInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTargetShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTargetShowInvocation) SetPathParamString(param string, value string) *ActionNotificationTargetShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetShowInvocation) NewMetaInput() *ActionNotificationTargetShowMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetShowInvocation) SetMetaInput(input *ActionNotificationTargetShowMetaGlobalInput) *ActionNotificationTargetShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTargetShowInvocation) Call() (*ActionNotificationTargetShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNotificationTargetShowInvocation) callAsQuery() (*ActionNotificationTargetShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNotificationTargetShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTarget
	}
	return resp, err
}

func (inv *ActionNotificationTargetShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
