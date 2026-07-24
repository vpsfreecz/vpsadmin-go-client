package client

import (
	"net/url"
	"strings"
)

// ActionNotificationReceiverTargetShow is a type for action Notification_receiver.Target#Show
type ActionNotificationReceiverTargetShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverTargetShow(client *Client) *ActionNotificationReceiverTargetShow {
	return &ActionNotificationReceiverTargetShow{
		Client: client,
	}
}

// ActionNotificationReceiverTargetShowMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverTargetShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverTargetShowMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverTargetShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverTargetShowMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverTargetShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverTargetShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetShowMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverTargetShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverTargetShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverTargetShowOutput is a type for action output parameters
type ActionNotificationReceiverTargetShowOutput struct {
	Action                 string "json:\"action\""
	CreatedAt              string "json:\"created_at\""
	DeliveryMethodEnabled  bool   "json:\"delivery_method_enabled\""
	DisplayTarget          string "json:\"display_target\""
	Id                     int64  "json:\"id\""
	Label                  string "json:\"label\""
	LastError              string "json:\"last_error\""
	NotificationTargetId   int64  "json:\"notification_target_id\""
	Position               int64  "json:\"position\""
	SecretPresent          bool   "json:\"secret_present\""
	TargetEnabled          bool   "json:\"target_enabled\""
	TargetKind             string "json:\"target_kind\""
	TargetValue            string "json:\"target_value\""
	TelegramBotName        string "json:\"telegram_bot_name\""
	TelegramBotUrl         string "json:\"telegram_bot_url\""
	TelegramPairingCommand string "json:\"telegram_pairing_command\""
	TelegramPairingUrl     string "json:\"telegram_pairing_url\""
	UpdatedAt              string "json:\"updated_at\""
	Verified               bool   "json:\"verified\""
	VerifiedAt             string "json:\"verified_at\""
}

// Type for action response, including envelope
type ActionNotificationReceiverTargetShowResponse struct {
	Action *ActionNotificationReceiverTargetShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Target *ActionNotificationReceiverTargetShowOutput "json:\"target\""
	}

	// Action output without the namespace
	Output *ActionNotificationReceiverTargetShowOutput
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverTargetShow) Prepare() *ActionNotificationReceiverTargetShowInvocation {
	return &ActionNotificationReceiverTargetShowInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers/{notification_receiver_id}/target/{target_id}",
	}
}

// ActionNotificationReceiverTargetShowInvocation is used to configure action for invocation
type ActionNotificationReceiverTargetShowInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverTargetShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverTargetShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationReceiverTargetShowInvocation) SetPathParamInt(param string, value int64) *ActionNotificationReceiverTargetShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationReceiverTargetShowInvocation) SetPathParamString(param string, value string) *ActionNotificationReceiverTargetShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverTargetShowInvocation) NewMetaInput() *ActionNotificationReceiverTargetShowMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverTargetShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverTargetShowInvocation) SetMetaInput(input *ActionNotificationReceiverTargetShowMetaGlobalInput) *ActionNotificationReceiverTargetShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverTargetShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverTargetShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverTargetShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationReceiverTargetShowInvocation) Call() (*ActionNotificationReceiverTargetShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNotificationReceiverTargetShowInvocation) callAsQuery() (*ActionNotificationReceiverTargetShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNotificationReceiverTargetShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Target
	}
	return resp, err
}

func (inv *ActionNotificationReceiverTargetShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
