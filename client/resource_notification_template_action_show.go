package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTemplateShow is a type for action Notification_template#Show
type ActionNotificationTemplateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTemplateShow(client *Client) *ActionNotificationTemplateShow {
	return &ActionNotificationTemplateShow{
		Client: client,
	}
}

// ActionNotificationTemplateShowMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTemplateShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTemplateShowMetaGlobalInput) SetIncludes(value string) *ActionNotificationTemplateShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTemplateShowMetaGlobalInput) SetNo(value bool) *ActionNotificationTemplateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateShowMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTemplateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTemplateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateShowOutput is a type for action output parameters
type ActionNotificationTemplateShowOutput struct {
	CreatedAt      string "json:\"created_at\""
	Id             int64  "json:\"id\""
	Label          string "json:\"label\""
	Name           string "json:\"name\""
	TemplateId     string "json:\"template_id\""
	UpdatedAt      string "json:\"updated_at\""
	UserVisibility string "json:\"user_visibility\""
}

// Type for action response, including envelope
type ActionNotificationTemplateShowResponse struct {
	Action *ActionNotificationTemplateShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationTemplate *ActionNotificationTemplateShowOutput "json:\"notification_template\""
	}

	// Action output without the namespace
	Output *ActionNotificationTemplateShowOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTemplateShow) Prepare() *ActionNotificationTemplateShowInvocation {
	return &ActionNotificationTemplateShowInvocation{
		Action: action,
		Path:   "/v7.0/notification_templates/{notification_template_id}",
	}
}

// ActionNotificationTemplateShowInvocation is used to configure action for invocation
type ActionNotificationTemplateShowInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTemplateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNotificationTemplateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTemplateShowInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTemplateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTemplateShowInvocation) SetPathParamString(param string, value string) *ActionNotificationTemplateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTemplateShowInvocation) NewMetaInput() *ActionNotificationTemplateShowMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTemplateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTemplateShowInvocation) SetMetaInput(input *ActionNotificationTemplateShowMetaGlobalInput) *ActionNotificationTemplateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTemplateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTemplateShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTemplateShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationTemplateShowInvocation) Call() (*ActionNotificationTemplateShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNotificationTemplateShowInvocation) callAsQuery() (*ActionNotificationTemplateShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNotificationTemplateShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationTemplate
	}
	return resp, err
}

func (inv *ActionNotificationTemplateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
