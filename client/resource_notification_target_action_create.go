package client

import ()

// ActionNotificationTargetCreate is a type for action Notification_target#Create
type ActionNotificationTargetCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetCreate(client *Client) *ActionNotificationTargetCreate {
	return &ActionNotificationTargetCreate{
		Client: client,
	}
}

// ActionNotificationTargetCreateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetCreateMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetCreateMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetCreateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetCreateInput is a type for action input parameters
type ActionNotificationTargetCreateInput struct {
	Action      string "json:\"action\""
	Enabled     bool   "json:\"enabled\""
	Label       string "json:\"label\""
	Secret      string "json:\"secret\""
	TargetKind  string "json:\"target_kind\""
	TargetValue string "json:\"target_value\""
	User        int64  "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAction sets parameter Action to value and selects it for sending
func (in *ActionNotificationTargetCreateInput) SetAction(value string) *ActionNotificationTargetCreateInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionNotificationTargetCreateInput) SetEnabled(value bool) *ActionNotificationTargetCreateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionNotificationTargetCreateInput) SetLabel(value string) *ActionNotificationTargetCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLabelNil(false)
	in._selectedParameters["Label"] = nil
	return in
}

// SetLabelNil sets parameter Label to nil and selects it for sending
func (in *ActionNotificationTargetCreateInput) SetLabelNil(set bool) *ActionNotificationTargetCreateInput {
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
func (in *ActionNotificationTargetCreateInput) SetSecret(value string) *ActionNotificationTargetCreateInput {
	in.Secret = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSecretNil(false)
	in._selectedParameters["Secret"] = nil
	return in
}

// SetSecretNil sets parameter Secret to nil and selects it for sending
func (in *ActionNotificationTargetCreateInput) SetSecretNil(set bool) *ActionNotificationTargetCreateInput {
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
func (in *ActionNotificationTargetCreateInput) SetTargetKind(value string) *ActionNotificationTargetCreateInput {
	in.TargetKind = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TargetKind"] = nil
	return in
}

// SetTargetValue sets parameter TargetValue to value and selects it for sending
func (in *ActionNotificationTargetCreateInput) SetTargetValue(value string) *ActionNotificationTargetCreateInput {
	in.TargetValue = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetTargetValueNil(false)
	in._selectedParameters["TargetValue"] = nil
	return in
}

// SetTargetValueNil sets parameter TargetValue to nil and selects it for sending
func (in *ActionNotificationTargetCreateInput) SetTargetValueNil(set bool) *ActionNotificationTargetCreateInput {
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

// SetUser sets parameter User to value and selects it for sending
func (in *ActionNotificationTargetCreateInput) SetUser(value int64) *ActionNotificationTargetCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetCreateInput) SelectParameters(params ...string) *ActionNotificationTargetCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTargetCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTargetCreateInput) UnselectParameters(params ...string) *ActionNotificationTargetCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTargetCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetCreateRequest is a type for the entire action request
type ActionNotificationTargetCreateRequest struct {
	NotificationTarget map[string]interface{} "json:\"notification_target\""
	Meta               map[string]interface{} "json:\"_meta\""
}

// ActionNotificationTargetCreateOutput is a type for action output parameters
type ActionNotificationTargetCreateOutput struct {
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
type ActionNotificationTargetCreateResponse struct {
	Action *ActionNotificationTargetCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTarget *ActionNotificationTargetCreateOutput "json:\"notification_target\""
	}

	// Action output without the namespace
	Output *ActionNotificationTargetCreateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTargetCreate) Prepare() *ActionNotificationTargetCreateInvocation {
	return &ActionNotificationTargetCreateInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets",
	}
}

// ActionNotificationTargetCreateInvocation is used to configure action for invocation
type ActionNotificationTargetCreateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTargetCreateInput
	// Global meta input parameters
	MetaInput *ActionNotificationTargetCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTargetCreateInvocation) NewInput() *ActionNotificationTargetCreateInput {
	inv.Input = &ActionNotificationTargetCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTargetCreateInvocation) SetInput(input *ActionNotificationTargetCreateInput) *ActionNotificationTargetCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTargetCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTargetCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetCreateInvocation) NewMetaInput() *ActionNotificationTargetCreateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetCreateInvocation) SetMetaInput(input *ActionNotificationTargetCreateMetaGlobalInput) *ActionNotificationTargetCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetCreateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTargetCreateInvocation) Call() (*ActionNotificationTargetCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationTargetCreateInvocation) callAsBody() (*ActionNotificationTargetCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationTargetCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTarget
	}
	return resp, err
}

func (inv *ActionNotificationTargetCreateInvocation) makeAllInputParams() *ActionNotificationTargetCreateRequest {
	return &ActionNotificationTargetCreateRequest{
		NotificationTarget: inv.makeInputParams(),
		Meta:               inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationTargetCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Action") {
			ret["action"] = inv.Input.Action
		}
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
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
	}

	return ret
}

func (inv *ActionNotificationTargetCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
