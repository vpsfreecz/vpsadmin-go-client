package client

import (
	"net/url"
	"strings"
)

// ActionNotificationReceiverUpdate is a type for action Notification_receiver#Update
type ActionNotificationReceiverUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverUpdate(client *Client) *ActionNotificationReceiverUpdate {
	return &ActionNotificationReceiverUpdate{
		Client: client,
	}
}

// ActionNotificationReceiverUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverUpdateMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverUpdateMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverUpdateInput is a type for action input parameters
type ActionNotificationReceiverUpdateInput struct {
	Description string "json:\"description\""
	Enabled     bool   "json:\"enabled\""
	Label       string "json:\"label\""
	Mute        bool   "json:\"mute\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDescription sets parameter Description to value and selects it for sending
func (in *ActionNotificationReceiverUpdateInput) SetDescription(value string) *ActionNotificationReceiverUpdateInput {
	in.Description = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDescriptionNil(false)
	in._selectedParameters["Description"] = nil
	return in
}

// SetDescriptionNil sets parameter Description to nil and selects it for sending
func (in *ActionNotificationReceiverUpdateInput) SetDescriptionNil(set bool) *ActionNotificationReceiverUpdateInput {
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
func (in *ActionNotificationReceiverUpdateInput) SetEnabled(value bool) *ActionNotificationReceiverUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionNotificationReceiverUpdateInput) SetLabel(value string) *ActionNotificationReceiverUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetMute sets parameter Mute to value and selects it for sending
func (in *ActionNotificationReceiverUpdateInput) SetMute(value bool) *ActionNotificationReceiverUpdateInput {
	in.Mute = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Mute"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverUpdateInput) SelectParameters(params ...string) *ActionNotificationReceiverUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationReceiverUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationReceiverUpdateInput) UnselectParameters(params ...string) *ActionNotificationReceiverUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationReceiverUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverUpdateRequest is a type for the entire action request
type ActionNotificationReceiverUpdateRequest struct {
	NotificationReceiver map[string]interface{} "json:\"notification_receiver\""
	Meta                 map[string]interface{} "json:\"_meta\""
}

// ActionNotificationReceiverUpdateOutput is a type for action output parameters
type ActionNotificationReceiverUpdateOutput struct {
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
type ActionNotificationReceiverUpdateResponse struct {
	Action *ActionNotificationReceiverUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationReceiver *ActionNotificationReceiverUpdateOutput "json:\"notification_receiver\""
	}

	// Action output without the namespace
	Output *ActionNotificationReceiverUpdateOutput
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverUpdate) Prepare() *ActionNotificationReceiverUpdateInvocation {
	return &ActionNotificationReceiverUpdateInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers/{notification_receiver_id}",
	}
}

// ActionNotificationReceiverUpdateInvocation is used to configure action for invocation
type ActionNotificationReceiverUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationReceiverUpdateInput
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationReceiverUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNotificationReceiverUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationReceiverUpdateInvocation) SetPathParamString(param string, value string) *ActionNotificationReceiverUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationReceiverUpdateInvocation) NewInput() *ActionNotificationReceiverUpdateInput {
	inv.Input = &ActionNotificationReceiverUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationReceiverUpdateInvocation) SetInput(input *ActionNotificationReceiverUpdateInput) *ActionNotificationReceiverUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationReceiverUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationReceiverUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverUpdateInvocation) NewMetaInput() *ActionNotificationReceiverUpdateMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverUpdateInvocation) SetMetaInput(input *ActionNotificationReceiverUpdateMetaGlobalInput) *ActionNotificationReceiverUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverUpdateInvocation) validate() error {
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
func (inv *ActionNotificationReceiverUpdateInvocation) Call() (*ActionNotificationReceiverUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNotificationReceiverUpdateInvocation) callAsBody() (*ActionNotificationReceiverUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNotificationReceiverUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationReceiver
	}
	return resp, err
}

func (inv *ActionNotificationReceiverUpdateInvocation) makeAllInputParams() *ActionNotificationReceiverUpdateRequest {
	return &ActionNotificationReceiverUpdateRequest{
		NotificationReceiver: inv.makeInputParams(),
		Meta:                 inv.makeMetaInputParams(),
	}
}

func (inv *ActionNotificationReceiverUpdateInvocation) makeInputParams() map[string]interface{} {
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
	}

	return ret
}

func (inv *ActionNotificationReceiverUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
