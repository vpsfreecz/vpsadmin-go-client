package client

import (
	"net/url"
	"strings"
)

// ActionUserNotificationDeliveryMethodUpdate is a type for action User.Notification_delivery_method#Update
type ActionUserNotificationDeliveryMethodUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNotificationDeliveryMethodUpdate(client *Client) *ActionUserNotificationDeliveryMethodUpdate {
	return &ActionUserNotificationDeliveryMethodUpdate{
		Client: client,
	}
}

// ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput) SetNo(value bool) *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationDeliveryMethodUpdateInput is a type for action input parameters
type ActionUserNotificationDeliveryMethodUpdateInput struct {
	Enabled bool "json:\"enabled\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodUpdateInput) SetEnabled(value bool) *ActionUserNotificationDeliveryMethodUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationDeliveryMethodUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationDeliveryMethodUpdateInput) SelectParameters(params ...string) *ActionUserNotificationDeliveryMethodUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserNotificationDeliveryMethodUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserNotificationDeliveryMethodUpdateInput) UnselectParameters(params ...string) *ActionUserNotificationDeliveryMethodUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserNotificationDeliveryMethodUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationDeliveryMethodUpdateRequest is a type for the entire action request
type ActionUserNotificationDeliveryMethodUpdateRequest struct {
	NotificationDeliveryMethod map[string]interface{} "json:\"notification_delivery_method\""
	Meta                       map[string]interface{} "json:\"_meta\""
}

// ActionUserNotificationDeliveryMethodUpdateOutput is a type for action output parameters
type ActionUserNotificationDeliveryMethodUpdateOutput struct {
	CreatedAt      string "json:\"created_at\""
	DeliveryMethod string "json:\"delivery_method\""
	Enabled        bool   "json:\"enabled\""
	Id             string "json:\"id\""
	Label          string "json:\"label\""
	UpdatedAt      string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionUserNotificationDeliveryMethodUpdateResponse struct {
	Action *ActionUserNotificationDeliveryMethodUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationDeliveryMethod *ActionUserNotificationDeliveryMethodUpdateOutput "json:\"notification_delivery_method\""
	}

	// Action output without the namespace
	Output *ActionUserNotificationDeliveryMethodUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserNotificationDeliveryMethodUpdate) Prepare() *ActionUserNotificationDeliveryMethodUpdateInvocation {
	return &ActionUserNotificationDeliveryMethodUpdateInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/notification_delivery_methods/{notification_delivery_method_id}",
	}
}

// ActionUserNotificationDeliveryMethodUpdateInvocation is used to configure action for invocation
type ActionUserNotificationDeliveryMethodUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserNotificationDeliveryMethodUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNotificationDeliveryMethodUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserNotificationDeliveryMethodUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) SetPathParamString(param string, value string) *ActionUserNotificationDeliveryMethodUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) NewInput() *ActionUserNotificationDeliveryMethodUpdateInput {
	inv.Input = &ActionUserNotificationDeliveryMethodUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) SetInput(input *ActionUserNotificationDeliveryMethodUpdateInput) *ActionUserNotificationDeliveryMethodUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) NewMetaInput() *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) SetMetaInput(input *ActionUserNotificationDeliveryMethodUpdateMetaGlobalInput) *ActionUserNotificationDeliveryMethodUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) validate() error {
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
func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) Call() (*ActionUserNotificationDeliveryMethodUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) callAsBody() (*ActionUserNotificationDeliveryMethodUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserNotificationDeliveryMethodUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationDeliveryMethod
	}
	return resp, err
}

func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) makeAllInputParams() *ActionUserNotificationDeliveryMethodUpdateRequest {
	return &ActionUserNotificationDeliveryMethodUpdateRequest{
		NotificationDeliveryMethod: inv.makeInputParams(),
		Meta:                       inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
	}

	return ret
}

func (inv *ActionUserNotificationDeliveryMethodUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
