package client

import ()

// ActionNotificationTargetIndex is a type for action Notification_target#Index
type ActionNotificationTargetIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTargetIndex(client *Client) *ActionNotificationTargetIndex {
	return &ActionNotificationTargetIndex{
		Client: client,
	}
}

// ActionNotificationTargetIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTargetIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNotificationTargetIndexMetaGlobalInput) SetCount(value bool) *ActionNotificationTargetIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTargetIndexMetaGlobalInput) SetIncludes(value string) *ActionNotificationTargetIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTargetIndexMetaGlobalInput) SetNo(value bool) *ActionNotificationTargetIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTargetIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTargetIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetIndexInput is a type for action input parameters
type ActionNotificationTargetIndexInput struct {
	Action  string "json:\"action\""
	Enabled bool   "json:\"enabled\""
	FromId  int64  "json:\"from_id\""
	Limit   int64  "json:\"limit\""
	User    int64  "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAction sets parameter Action to value and selects it for sending
func (in *ActionNotificationTargetIndexInput) SetAction(value string) *ActionNotificationTargetIndexInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionNotificationTargetIndexInput) SetEnabled(value bool) *ActionNotificationTargetIndexInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNotificationTargetIndexInput) SetFromId(value int64) *ActionNotificationTargetIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNotificationTargetIndexInput) SetLimit(value int64) *ActionNotificationTargetIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionNotificationTargetIndexInput) SetUser(value int64) *ActionNotificationTargetIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTargetIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTargetIndexInput) SelectParameters(params ...string) *ActionNotificationTargetIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTargetIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTargetIndexInput) UnselectParameters(params ...string) *ActionNotificationTargetIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTargetIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTargetIndexOutput is a type for action output parameters
type ActionNotificationTargetIndexOutput struct {
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
type ActionNotificationTargetIndexResponse struct {
	Action *ActionNotificationTargetIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTargets []*ActionNotificationTargetIndexOutput "json:\"notification_targets\""
	}

	// Action output without the namespace
	Output []*ActionNotificationTargetIndexOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTargetIndex) Prepare() *ActionNotificationTargetIndexInvocation {
	return &ActionNotificationTargetIndexInvocation{
		Action: action,
		Path:   "/v7.0/notification_targets",
	}
}

// ActionNotificationTargetIndexInvocation is used to configure action for invocation
type ActionNotificationTargetIndexInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTargetIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTargetIndexInput
	// Global meta input parameters
	MetaInput *ActionNotificationTargetIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTargetIndexInvocation) NewInput() *ActionNotificationTargetIndexInput {
	inv.Input = &ActionNotificationTargetIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTargetIndexInvocation) SetInput(input *ActionNotificationTargetIndexInput) *ActionNotificationTargetIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTargetIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTargetIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTargetIndexInvocation) NewMetaInput() *ActionNotificationTargetIndexMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTargetIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTargetIndexInvocation) SetMetaInput(input *ActionNotificationTargetIndexMetaGlobalInput) *ActionNotificationTargetIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTargetIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTargetIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTargetIndexInvocation) validate() error {
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
func (inv *ActionNotificationTargetIndexInvocation) Call() (*ActionNotificationTargetIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNotificationTargetIndexInvocation) callAsQuery() (*ActionNotificationTargetIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNotificationTargetIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTargets
	}
	return resp, err
}

func (inv *ActionNotificationTargetIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Action") {
			ret["notification_target[action]"] = inv.Input.Action
		}
		if inv.IsParameterSelected("Enabled") {
			ret["notification_target[enabled]"] = convertBoolToString(inv.Input.Enabled)
		}
		if inv.IsParameterSelected("FromId") {
			ret["notification_target[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["notification_target[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["notification_target[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionNotificationTargetIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
