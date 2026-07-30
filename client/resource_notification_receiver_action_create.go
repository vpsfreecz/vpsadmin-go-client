package client

import ()

// ActionNotificationReceiverCreate is a type for action Notification_receiver#Create
type ActionNotificationReceiverCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverCreate(client *Client) *ActionNotificationReceiverCreate {
	return &ActionNotificationReceiverCreate{
		Client: client,
	}
}

// ActionNotificationReceiverCreateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverCreateMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverCreateMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverCreateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverCreateInput is a type for action input parameters
type ActionNotificationReceiverCreateInput struct {
	Description string "json:\"description\""
	Enabled     bool   "json:\"enabled\""
	Label       string "json:\"label\""
	Mute        bool   "json:\"mute\""
	User        int64  "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDescription sets parameter Description to value and selects it for sending
func (in *ActionNotificationReceiverCreateInput) SetDescription(value string) *ActionNotificationReceiverCreateInput {
	in.Description = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDescriptionNil(false)
	in._selectedParameters["Description"] = nil
	return in
}

// SetDescriptionNil sets parameter Description to nil and selects it for sending
func (in *ActionNotificationReceiverCreateInput) SetDescriptionNil(set bool) *ActionNotificationReceiverCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Description"] = nil
		in.SelectParameters("Description")
	} else {
		delete(in._nilParameters, "Description")
	}
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionNotificationReceiverCreateInput) SetEnabled(value bool) *ActionNotificationReceiverCreateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionNotificationReceiverCreateInput) SetLabel(value string) *ActionNotificationReceiverCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetMute sets parameter Mute to value and selects it for sending
func (in *ActionNotificationReceiverCreateInput) SetMute(value bool) *ActionNotificationReceiverCreateInput {
	in.Mute = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Mute"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionNotificationReceiverCreateInput) SetUser(value int64) *ActionNotificationReceiverCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverCreateInput) SelectParameters(params ...string) *ActionNotificationReceiverCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationReceiverCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationReceiverCreateInput) UnselectParameters(params ...string) *ActionNotificationReceiverCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationReceiverCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverCreateRequest is a type for the entire action request
type ActionNotificationReceiverCreateRequest struct {
	NotificationReceiver map[string]interface{} "json:\"notification_receiver\""
	Meta                 map[string]interface{} "json:\"_meta\""
}

// ActionNotificationReceiverCreateOutput is a type for action output parameters
type ActionNotificationReceiverCreateOutput struct {
	CreatedAt            string                "json:\"created_at\""
	Description          string                "json:\"description\""
	DisplayActionSummary string                "json:\"display_action_summary\""
	Enabled              bool                  "json:\"enabled\""
	Id                   int64                 "json:\"id\""
	Label                string                "json:\"label\""
	Mute                 bool                  "json:\"mute\""
	UpdatedAt            string                "json:\"updated_at\""
	User                 *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionNotificationReceiverCreateResponse struct {
	Action *ActionNotificationReceiverCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationReceiver *ActionNotificationReceiverCreateOutput "json:\"notification_receiver\""
	}

	// Action output without the namespace
	Output *ActionNotificationReceiverCreateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverCreate) Prepare() *ActionNotificationReceiverCreateInvocation {
	return &ActionNotificationReceiverCreateInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers",
	}
}

// ActionNotificationReceiverCreateInvocation is used to configure action for invocation
type ActionNotificationReceiverCreateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationReceiverCreateInput
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationReceiverCreateInvocation) NewInput() *ActionNotificationReceiverCreateInput {
	inv.Input = &ActionNotificationReceiverCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationReceiverCreateInvocation) SetInput(input *ActionNotificationReceiverCreateInput) *ActionNotificationReceiverCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationReceiverCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationReceiverCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverCreateInvocation) NewMetaInput() *ActionNotificationReceiverCreateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverCreateInvocation) SetMetaInput(input *ActionNotificationReceiverCreateMetaGlobalInput) *ActionNotificationReceiverCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverCreateInvocation) validate() error {
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
func (inv *ActionNotificationReceiverCreateInvocation) Call() (*ActionNotificationReceiverCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationReceiverCreateInvocation) callAsBody() (*ActionNotificationReceiverCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationReceiverCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationReceiver
	}
	return resp, err
}

func (inv *ActionNotificationReceiverCreateInvocation) makeAllInputParams() *ActionNotificationReceiverCreateRequest {
	return &ActionNotificationReceiverCreateRequest{
		NotificationReceiver: inv.makeInputParams(),
		Meta:                 inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationReceiverCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Description") {
			if inv.IsParameterNil("Description") {
				ret["description"] = nil
			} else {
				ret["description"] = inv.Input.Description
			}
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Mute") {
			ret["mute"] = inv.Input.Mute
		}
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
	}

	return ret
}

func (inv *ActionNotificationReceiverCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
