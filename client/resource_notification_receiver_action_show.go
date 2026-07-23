package client

import (
	"net/url"
	"strings"
)

// ActionNotificationReceiverShow is a type for action Notification_receiver#Show
type ActionNotificationReceiverShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverShow(client *Client) *ActionNotificationReceiverShow {
	return &ActionNotificationReceiverShow{
		Client: client,
	}
}

// ActionNotificationReceiverShowMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverShowMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverShowMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverShowMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverShowOutput is a type for action output parameters
type ActionNotificationReceiverShowOutput struct {
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
type ActionNotificationReceiverShowResponse struct {
	Action *ActionNotificationReceiverShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationReceiver *ActionNotificationReceiverShowOutput "json:\"notification_receiver\""
	}

	// Action output without the namespace
	Output *ActionNotificationReceiverShowOutput
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverShow) Prepare() *ActionNotificationReceiverShowInvocation {
	return &ActionNotificationReceiverShowInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers/{notification_receiver_id}",
	}
}

// ActionNotificationReceiverShowInvocation is used to configure action for invocation
type ActionNotificationReceiverShowInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationReceiverShowInvocation) SetPathParamInt(param string, value int64) *ActionNotificationReceiverShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationReceiverShowInvocation) SetPathParamString(param string, value string) *ActionNotificationReceiverShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverShowInvocation) NewMetaInput() *ActionNotificationReceiverShowMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverShowInvocation) SetMetaInput(input *ActionNotificationReceiverShowMetaGlobalInput) *ActionNotificationReceiverShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationReceiverShowInvocation) Call() (*ActionNotificationReceiverShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNotificationReceiverShowInvocation) callAsQuery() (*ActionNotificationReceiverShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNotificationReceiverShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationReceiver
	}
	return resp, err
}

func (inv *ActionNotificationReceiverShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
