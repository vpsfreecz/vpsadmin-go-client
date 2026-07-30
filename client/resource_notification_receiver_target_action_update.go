package client

import (
	"net/url"
	"strings"
)

// ActionNotificationReceiverTargetUpdate is a type for action Notification_receiver.Target#Update
type ActionNotificationReceiverTargetUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverTargetUpdate(client *Client) *ActionNotificationReceiverTargetUpdate {
	return &ActionNotificationReceiverTargetUpdate{
		Client: client,
	}
}

// ActionNotificationReceiverTargetUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverTargetUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverTargetUpdateMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverTargetUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverTargetUpdateMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverTargetUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverTargetUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverTargetUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverTargetUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverTargetUpdateInput is a type for action input parameters
type ActionNotificationReceiverTargetUpdateInput struct {
	Position int64 "json:\"position\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetPosition sets parameter Position to value and selects it for sending
func (in *ActionNotificationReceiverTargetUpdateInput) SetPosition(value int64) *ActionNotificationReceiverTargetUpdateInput {
	in.Position = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetPositionNil(false)
	in._selectedParameters["Position"] = nil
	return in
}

// SetPositionNil sets parameter Position to nil and selects it for sending
func (in *ActionNotificationReceiverTargetUpdateInput) SetPositionNil(set bool) *ActionNotificationReceiverTargetUpdateInput {
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

// SelectParameters sets parameters from ActionNotificationReceiverTargetUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetUpdateInput) SelectParameters(params ...string) *ActionNotificationReceiverTargetUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationReceiverTargetUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationReceiverTargetUpdateInput) UnselectParameters(params ...string) *ActionNotificationReceiverTargetUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationReceiverTargetUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverTargetUpdateRequest is a type for the entire action request
type ActionNotificationReceiverTargetUpdateRequest struct {
	Target map[string]interface{} "json:\"target\""
	Meta   map[string]interface{} "json:\"_meta\""
}

// ActionNotificationReceiverTargetUpdateOutput is a type for action output parameters
type ActionNotificationReceiverTargetUpdateOutput struct {
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
type ActionNotificationReceiverTargetUpdateResponse struct {
	Action *ActionNotificationReceiverTargetUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Target *ActionNotificationReceiverTargetUpdateOutput "json:\"target\""
	}

	// Action output without the namespace
	Output *ActionNotificationReceiverTargetUpdateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverTargetUpdate) Prepare() *ActionNotificationReceiverTargetUpdateInvocation {
	return &ActionNotificationReceiverTargetUpdateInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers/{notification_receiver_id}/target/{target_id}",
	}
}

// ActionNotificationReceiverTargetUpdateInvocation is used to configure action for invocation
type ActionNotificationReceiverTargetUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverTargetUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationReceiverTargetUpdateInput
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverTargetUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationReceiverTargetUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNotificationReceiverTargetUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationReceiverTargetUpdateInvocation) SetPathParamString(param string, value string) *ActionNotificationReceiverTargetUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationReceiverTargetUpdateInvocation) NewInput() *ActionNotificationReceiverTargetUpdateInput {
	inv.Input = &ActionNotificationReceiverTargetUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationReceiverTargetUpdateInvocation) SetInput(input *ActionNotificationReceiverTargetUpdateInput) *ActionNotificationReceiverTargetUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationReceiverTargetUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationReceiverTargetUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverTargetUpdateInvocation) NewMetaInput() *ActionNotificationReceiverTargetUpdateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverTargetUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverTargetUpdateInvocation) SetMetaInput(input *ActionNotificationReceiverTargetUpdateMetaGlobalInput) *ActionNotificationReceiverTargetUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverTargetUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverTargetUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverTargetUpdateInvocation) validate() error {
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
func (inv *ActionNotificationReceiverTargetUpdateInvocation) Call() (*ActionNotificationReceiverTargetUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationReceiverTargetUpdateInvocation) callAsBody() (*ActionNotificationReceiverTargetUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationReceiverTargetUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Target
	}
	return resp, err
}

func (inv *ActionNotificationReceiverTargetUpdateInvocation) makeAllInputParams() *ActionNotificationReceiverTargetUpdateRequest {
	return &ActionNotificationReceiverTargetUpdateRequest{
		Target: inv.makeInputParams(),
		Meta:   inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationReceiverTargetUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
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

func (inv *ActionNotificationReceiverTargetUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
