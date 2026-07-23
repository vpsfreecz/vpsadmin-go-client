package client

import (
	"net/url"
	"strings"
)

// ActionUserNotificationDeliveryMethodIndex is a type for action User.Notification_delivery_method#Index
type ActionUserNotificationDeliveryMethodIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNotificationDeliveryMethodIndex(client *Client) *ActionUserNotificationDeliveryMethodIndex {
	return &ActionUserNotificationDeliveryMethodIndex{
		Client: client,
	}
}

// ActionUserNotificationDeliveryMethodIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserNotificationDeliveryMethodIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput) SetCount(value bool) *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput) SetIncludes(value string) *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput) SetNo(value bool) *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationDeliveryMethodIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationDeliveryMethodIndexInput is a type for action input parameters
type ActionUserNotificationDeliveryMethodIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodIndexInput) SetFromId(value int64) *ActionUserNotificationDeliveryMethodIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodIndexInput) SetLimit(value int64) *ActionUserNotificationDeliveryMethodIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationDeliveryMethodIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationDeliveryMethodIndexInput) SelectParameters(params ...string) *ActionUserNotificationDeliveryMethodIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserNotificationDeliveryMethodIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserNotificationDeliveryMethodIndexInput) UnselectParameters(params ...string) *ActionUserNotificationDeliveryMethodIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserNotificationDeliveryMethodIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationDeliveryMethodIndexOutput is a type for action output parameters
type ActionUserNotificationDeliveryMethodIndexOutput struct {
	CreatedAt      string "json:\"created_at\""
	DeliveryMethod string "json:\"delivery_method\""
	Enabled        bool   "json:\"enabled\""
	Id             string "json:\"id\""
	Label          string "json:\"label\""
	UpdatedAt      string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionUserNotificationDeliveryMethodIndexResponse struct {
	Action *ActionUserNotificationDeliveryMethodIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationDeliveryMethods []*ActionUserNotificationDeliveryMethodIndexOutput "json:\"notification_delivery_methods\""
	}

	// Action output without the namespace
	Output []*ActionUserNotificationDeliveryMethodIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserNotificationDeliveryMethodIndex) Prepare() *ActionUserNotificationDeliveryMethodIndexInvocation {
	return &ActionUserNotificationDeliveryMethodIndexInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/notification_delivery_methods",
	}
}

// ActionUserNotificationDeliveryMethodIndexInvocation is used to configure action for invocation
type ActionUserNotificationDeliveryMethodIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserNotificationDeliveryMethodIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNotificationDeliveryMethodIndexInput
	// Global meta input parameters
	MetaInput *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserNotificationDeliveryMethodIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) SetPathParamString(param string, value string) *ActionUserNotificationDeliveryMethodIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) NewInput() *ActionUserNotificationDeliveryMethodIndexInput {
	inv.Input = &ActionUserNotificationDeliveryMethodIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) SetInput(input *ActionUserNotificationDeliveryMethodIndexInput) *ActionUserNotificationDeliveryMethodIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) NewMetaInput() *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserNotificationDeliveryMethodIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) SetMetaInput(input *ActionUserNotificationDeliveryMethodIndexMetaGlobalInput) *ActionUserNotificationDeliveryMethodIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) validate() error {
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
func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) Call() (*ActionUserNotificationDeliveryMethodIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) callAsQuery() (*ActionUserNotificationDeliveryMethodIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNotificationDeliveryMethodIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationDeliveryMethods
	}
	return resp, err
}

func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["notification_delivery_method[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["notification_delivery_method[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserNotificationDeliveryMethodIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
