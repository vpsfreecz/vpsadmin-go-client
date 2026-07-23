package client

import (
	"net/url"
	"strings"
)

// ActionNotificationReceiverTargetIndex is a type for action Notification_receiver.Target#Index
type ActionNotificationReceiverTargetIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverTargetIndex(client *Client) *ActionNotificationReceiverTargetIndex {
	return &ActionNotificationReceiverTargetIndex{
		Client: client,
	}
}

// ActionNotificationReceiverTargetIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverTargetIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNotificationReceiverTargetIndexMetaGlobalInput) SetCount(value bool) *ActionNotificationReceiverTargetIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverTargetIndexMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverTargetIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverTargetIndexMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverTargetIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverTargetIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverTargetIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverTargetIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverTargetIndexInput is a type for action input parameters
type ActionNotificationReceiverTargetIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNotificationReceiverTargetIndexInput) SetFromId(value int64) *ActionNotificationReceiverTargetIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNotificationReceiverTargetIndexInput) SetLimit(value int64) *ActionNotificationReceiverTargetIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverTargetIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetIndexInput) SelectParameters(params ...string) *ActionNotificationReceiverTargetIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationReceiverTargetIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetIndexInput) UnselectParameters(params ...string) *ActionNotificationReceiverTargetIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationReceiverTargetIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverTargetIndexOutput is a type for action output parameters
type ActionNotificationReceiverTargetIndexOutput struct {
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
type ActionNotificationReceiverTargetIndexResponse struct {
	Action *ActionNotificationReceiverTargetIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Targets []*ActionNotificationReceiverTargetIndexOutput "json:\"targets\""
	}

	// Action output without the namespace
	Output []*ActionNotificationReceiverTargetIndexOutput
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverTargetIndex) Prepare() *ActionNotificationReceiverTargetIndexInvocation {
	return &ActionNotificationReceiverTargetIndexInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers/{notification_receiver_id}/target",
	}
}

// ActionNotificationReceiverTargetIndexInvocation is used to configure action for invocation
type ActionNotificationReceiverTargetIndexInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverTargetIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationReceiverTargetIndexInput
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverTargetIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationReceiverTargetIndexInvocation) SetPathParamInt(param string, value int64) *ActionNotificationReceiverTargetIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationReceiverTargetIndexInvocation) SetPathParamString(param string, value string) *ActionNotificationReceiverTargetIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationReceiverTargetIndexInvocation) NewInput() *ActionNotificationReceiverTargetIndexInput {
	inv.Input = &ActionNotificationReceiverTargetIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationReceiverTargetIndexInvocation) SetInput(input *ActionNotificationReceiverTargetIndexInput) *ActionNotificationReceiverTargetIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationReceiverTargetIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationReceiverTargetIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverTargetIndexInvocation) NewMetaInput() *ActionNotificationReceiverTargetIndexMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverTargetIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverTargetIndexInvocation) SetMetaInput(input *ActionNotificationReceiverTargetIndexMetaGlobalInput) *ActionNotificationReceiverTargetIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverTargetIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverTargetIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverTargetIndexInvocation) validate() error {
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
func (inv *ActionNotificationReceiverTargetIndexInvocation) Call() (*ActionNotificationReceiverTargetIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNotificationReceiverTargetIndexInvocation) callAsQuery() (*ActionNotificationReceiverTargetIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNotificationReceiverTargetIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Targets
	}
	return resp, err
}

func (inv *ActionNotificationReceiverTargetIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["target[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["target[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionNotificationReceiverTargetIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
