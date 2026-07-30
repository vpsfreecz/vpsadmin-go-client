package client

import (
	"net/url"
	"strings"
)

// ActionNotificationReceiverTargetCreate is a type for action Notification_receiver.Target#Create
type ActionNotificationReceiverTargetCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverTargetCreate(client *Client) *ActionNotificationReceiverTargetCreate {
	return &ActionNotificationReceiverTargetCreate{
		Client: client,
	}
}

// ActionNotificationReceiverTargetCreateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverTargetCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverTargetCreateMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverTargetCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverTargetCreateMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverTargetCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverTargetCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetCreateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverTargetCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverTargetCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverTargetCreateInput is a type for action input parameters
type ActionNotificationReceiverTargetCreateInput struct {
	NotificationTargetId int64 "json:\"notification_target_id\""
	Position             int64 "json:\"position\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNotificationTargetId sets parameter NotificationTargetId to value and selects it for sending
func (in *ActionNotificationReceiverTargetCreateInput) SetNotificationTargetId(value int64) *ActionNotificationReceiverTargetCreateInput {
	in.NotificationTargetId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NotificationTargetId"] = nil
	return in
}

// SetPosition sets parameter Position to value and selects it for sending
func (in *ActionNotificationReceiverTargetCreateInput) SetPosition(value int64) *ActionNotificationReceiverTargetCreateInput {
	in.Position = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetPositionNil(false)
	in._selectedParameters["Position"] = nil
	return in
}

// SetPositionNil sets parameter Position to nil and selects it for sending
func (in *ActionNotificationReceiverTargetCreateInput) SetPositionNil(set bool) *ActionNotificationReceiverTargetCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Position"] = nil
		in.SelectParameters("Position")
	} else {
		delete(in._nilParameters, "Position")
	}
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverTargetCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetCreateInput) SelectParameters(params ...string) *ActionNotificationReceiverTargetCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationReceiverTargetCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetCreateInput) UnselectParameters(params ...string) *ActionNotificationReceiverTargetCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationReceiverTargetCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverTargetCreateRequest is a type for the entire action request
type ActionNotificationReceiverTargetCreateRequest struct {
	Target map[string]interface{} "json:\"target\""
	Meta   map[string]interface{} "json:\"_meta\""
}

// ActionNotificationReceiverTargetCreateOutput is a type for action output parameters
type ActionNotificationReceiverTargetCreateOutput struct {
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
type ActionNotificationReceiverTargetCreateResponse struct {
	Action *ActionNotificationReceiverTargetCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Target *ActionNotificationReceiverTargetCreateOutput "json:\"target\""
	}

	// Action output without the namespace
	Output *ActionNotificationReceiverTargetCreateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverTargetCreate) Prepare() *ActionNotificationReceiverTargetCreateInvocation {
	return &ActionNotificationReceiverTargetCreateInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers/{notification_receiver_id}/target",
	}
}

// ActionNotificationReceiverTargetCreateInvocation is used to configure action for invocation
type ActionNotificationReceiverTargetCreateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverTargetCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationReceiverTargetCreateInput
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverTargetCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationReceiverTargetCreateInvocation) SetPathParamInt(param string, value int64) *ActionNotificationReceiverTargetCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationReceiverTargetCreateInvocation) SetPathParamString(param string, value string) *ActionNotificationReceiverTargetCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationReceiverTargetCreateInvocation) NewInput() *ActionNotificationReceiverTargetCreateInput {
	inv.Input = &ActionNotificationReceiverTargetCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationReceiverTargetCreateInvocation) SetInput(input *ActionNotificationReceiverTargetCreateInput) *ActionNotificationReceiverTargetCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationReceiverTargetCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationReceiverTargetCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverTargetCreateInvocation) NewMetaInput() *ActionNotificationReceiverTargetCreateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverTargetCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverTargetCreateInvocation) SetMetaInput(input *ActionNotificationReceiverTargetCreateMetaGlobalInput) *ActionNotificationReceiverTargetCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverTargetCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverTargetCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverTargetCreateInvocation) validate() error {
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
func (inv *ActionNotificationReceiverTargetCreateInvocation) Call() (*ActionNotificationReceiverTargetCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationReceiverTargetCreateInvocation) callAsBody() (*ActionNotificationReceiverTargetCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationReceiverTargetCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Target
	}
	return resp, err
}

func (inv *ActionNotificationReceiverTargetCreateInvocation) makeAllInputParams() *ActionNotificationReceiverTargetCreateRequest {
	return &ActionNotificationReceiverTargetCreateRequest{
		Target: inv.makeInputParams(),
		Meta:   inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationReceiverTargetCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("NotificationTargetId") {
			ret["notification_target_id"] = inv.Input.NotificationTargetId
		}
		if inv.IsParameterSelected("Position") {
			if inv.IsParameterNil("Position") {
				ret["position"] = nil
			} else {
				ret["position"] = inv.Input.Position
			}
		}
	}

	return ret
}

func (inv *ActionNotificationReceiverTargetCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
