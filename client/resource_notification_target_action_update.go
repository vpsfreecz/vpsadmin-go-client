package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTargetUpdate is a type for action Notification_target#Update
type ActionNotificationTargetUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetUpdate(client *Client) *ActionNotificationTargetUpdate {
	return &ActionNotificationTargetUpdate{
		Client: client,
	}
}

// ActionNotificationTargetUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetUpdateMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetUpdateMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetUpdateInput is a type for action input parameters
type ActionNotificationTargetUpdateInput struct {
	Enabled     bool   "json:\"enabled\""
	Label       string "json:\"label\""
	Secret      string "json:\"secret\""
	TargetKind  string "json:\"target_kind\""
	TargetValue string "json:\"target_value\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionNotificationTargetUpdateInput) SetEnabled(value bool) *ActionNotificationTargetUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionNotificationTargetUpdateInput) SetLabel(value string) *ActionNotificationTargetUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLabelNil(false)
	in._selectedParameters["Label"] = nil
	return in
}

// SetLabelNil sets parameter Label to nil and selects it for sending
func (in *ActionNotificationTargetUpdateInput) SetLabelNil(set bool) *ActionNotificationTargetUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Label"] = nil
		in.SelectParameters("Label")
	} else {
		delete(in._nilParameters, "Label")
	}
	return in
}

// SetSecret sets parameter Secret to value and selects it for sending
func (in *ActionNotificationTargetUpdateInput) SetSecret(value string) *ActionNotificationTargetUpdateInput {
	in.Secret = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSecretNil(false)
	in._selectedParameters["Secret"] = nil
	return in
}

// SetSecretNil sets parameter Secret to nil and selects it for sending
func (in *ActionNotificationTargetUpdateInput) SetSecretNil(set bool) *ActionNotificationTargetUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Secret"] = nil
		in.SelectParameters("Secret")
	} else {
		delete(in._nilParameters, "Secret")
	}
	return in
}

// SetTargetKind sets parameter TargetKind to value and selects it for sending
func (in *ActionNotificationTargetUpdateInput) SetTargetKind(value string) *ActionNotificationTargetUpdateInput {
	in.TargetKind = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TargetKind"] = nil
	return in
}

// SetTargetValue sets parameter TargetValue to value and selects it for sending
func (in *ActionNotificationTargetUpdateInput) SetTargetValue(value string) *ActionNotificationTargetUpdateInput {
	in.TargetValue = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetTargetValueNil(false)
	in._selectedParameters["TargetValue"] = nil
	return in
}

// SetTargetValueNil sets parameter TargetValue to nil and selects it for sending
func (in *ActionNotificationTargetUpdateInput) SetTargetValueNil(set bool) *ActionNotificationTargetUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["TargetValue"] = nil
		in.SelectParameters("TargetValue")
	} else {
		delete(in._nilParameters, "TargetValue")
	}
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetUpdateInput) SelectParameters(params ...string) *ActionNotificationTargetUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTargetUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTargetUpdateInput) UnselectParameters(params ...string) *ActionNotificationTargetUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTargetUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetUpdateRequest is a type for the entire action request
type ActionNotificationTargetUpdateRequest struct {
	NotificationTarget map[string]interface{} "json:\"notification_target\""
	Meta               map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTargetUpdateOutput is a type for action output parameters
type ActionNotificationTargetUpdateOutput struct {
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
type ActionNotificationTargetUpdateResponse struct {
	Action *ActionNotificationTargetUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTarget *ActionNotificationTargetUpdateOutput "json:\"notification_target\""
	}

	// Action output without the namespace
	Output *ActionNotificationTargetUpdateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTargetUpdate) Prepare() *ActionNotificationTargetUpdateInvocation {
	return &ActionNotificationTargetUpdateInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets/{notification_target_id}",
	}
}

// ActionNotificationTargetUpdateInvocation is used to configure action for invocation
type ActionNotificationTargetUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTargetUpdateInput
	// Global meta input parameters
	MetaInput *ActionNotificationTargetUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTargetUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTargetUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTargetUpdateInvocation) SetPathParamString(param string, value string) *ActionNotificationTargetUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTargetUpdateInvocation) NewInput() *ActionNotificationTargetUpdateInput {
	inv.Input = &ActionNotificationTargetUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTargetUpdateInvocation) SetInput(input *ActionNotificationTargetUpdateInput) *ActionNotificationTargetUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTargetUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTargetUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetUpdateInvocation) NewMetaInput() *ActionNotificationTargetUpdateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetUpdateInvocation) SetMetaInput(input *ActionNotificationTargetUpdateMetaGlobalInput) *ActionNotificationTargetUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetUpdateInvocation) validate() error {
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
func (inv *ActionNotificationTargetUpdateInvocation) Call() (*ActionNotificationTargetUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTargetUpdateInvocation) callAsBody() (*ActionNotificationTargetUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTargetUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTarget
	}
	return resp, err
}

func (inv *ActionNotificationTargetUpdateInvocation) makeAllInputParams() *ActionNotificationTargetUpdateRequest {
	return &ActionNotificationTargetUpdateRequest{
		NotificationTarget: inv.makeInputParams(),
		Meta:               inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTargetUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("Label") {
			if inv.IsParameterNil("Label") {
				ret["label"] = nil
			} else {
				ret["label"] = inv.Input.Label
			}
		}
		if inv.IsParameterSelected("Secret") {
			if inv.IsParameterNil("Secret") {
				ret["secret"] = nil
			} else {
				ret["secret"] = inv.Input.Secret
			}
		}
		if inv.IsParameterSelected("TargetKind") {
			ret["target_kind"] = inv.Input.TargetKind
		}
		if inv.IsParameterSelected("TargetValue") {
			if inv.IsParameterNil("TargetValue") {
				ret["target_value"] = nil
			} else {
				ret["target_value"] = inv.Input.TargetValue
			}
		}
	}

	return ret
}

func (inv *ActionNotificationTargetUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
