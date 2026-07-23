package client

import (
	"net/url"
	"strings"
)

// ActionUserNotificationDeliveryMethodShow is a type for action User.Notification_delivery_method#Show
type ActionUserNotificationDeliveryMethodShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNotificationDeliveryMethodShow(client *Client) *ActionUserNotificationDeliveryMethodShow {
	return &ActionUserNotificationDeliveryMethodShow{
		Client: client,
	}
}

// ActionUserNotificationDeliveryMethodShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserNotificationDeliveryMethodShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodShowMetaGlobalInput) SetIncludes(value string) *ActionUserNotificationDeliveryMethodShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNotificationDeliveryMethodShowMetaGlobalInput) SetNo(value bool) *ActionUserNotificationDeliveryMethodShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationDeliveryMethodShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationDeliveryMethodShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserNotificationDeliveryMethodShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNotificationDeliveryMethodShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationDeliveryMethodShowOutput is a type for action output parameters
type ActionUserNotificationDeliveryMethodShowOutput struct {
	CreatedAt      string "json:\"created_at\""
	DeliveryMethod string "json:\"delivery_method\""
	Enabled        bool   "json:\"enabled\""
	Id             string "json:\"id\""
	Label          string "json:\"label\""
	UpdatedAt      string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionUserNotificationDeliveryMethodShowResponse struct {
	Action *ActionUserNotificationDeliveryMethodShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationDeliveryMethod *ActionUserNotificationDeliveryMethodShowOutput "json:\"notification_delivery_method\""
	}

	// Action output without the namespace
	Output *ActionUserNotificationDeliveryMethodShowOutput
}

// Prepare the action for invocation
func (action *ActionUserNotificationDeliveryMethodShow) Prepare() *ActionUserNotificationDeliveryMethodShowInvocation {
	return &ActionUserNotificationDeliveryMethodShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/notification_delivery_methods/{notification_delivery_method_id}",
	}
}

// ActionUserNotificationDeliveryMethodShowInvocation is used to configure action for invocation
type ActionUserNotificationDeliveryMethodShowInvocation struct {
	// Pointer to the action
	Action *ActionUserNotificationDeliveryMethodShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserNotificationDeliveryMethodShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNotificationDeliveryMethodShowInvocation) SetPathParamInt(param string, value int64) *ActionUserNotificationDeliveryMethodShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNotificationDeliveryMethodShowInvocation) SetPathParamString(param string, value string) *ActionUserNotificationDeliveryMethodShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNotificationDeliveryMethodShowInvocation) NewMetaInput() *ActionUserNotificationDeliveryMethodShowMetaGlobalInput {
	inv.MetaInput = &ActionUserNotificationDeliveryMethodShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNotificationDeliveryMethodShowInvocation) SetMetaInput(input *ActionUserNotificationDeliveryMethodShowMetaGlobalInput) *ActionUserNotificationDeliveryMethodShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNotificationDeliveryMethodShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNotificationDeliveryMethodShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserNotificationDeliveryMethodShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNotificationDeliveryMethodShowInvocation) Call() (*ActionUserNotificationDeliveryMethodShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionUserNotificationDeliveryMethodShowInvocation) callAsQuery() (*ActionUserNotificationDeliveryMethodShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNotificationDeliveryMethodShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationDeliveryMethod
	}
	return resp, err
}

func (inv *ActionUserNotificationDeliveryMethodShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
