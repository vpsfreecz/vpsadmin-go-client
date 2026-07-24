package client

import (
	"net/url"
	"strings"
)

// ActionNotificationTemplateVariantIndex is a type for action Notification_template.Variant#Index
type ActionNotificationTemplateVariantIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationTemplateVariantIndex(client *Client) *ActionNotificationTemplateVariantIndex {
	return &ActionNotificationTemplateVariantIndex{
		Client: client,
	}
}

// ActionNotificationTemplateVariantIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationTemplateVariantIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNotificationTemplateVariantIndexMetaGlobalInput) SetCount(value bool) *ActionNotificationTemplateVariantIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationTemplateVariantIndexMetaGlobalInput) SetIncludes(value string) *ActionNotificationTemplateVariantIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationTemplateVariantIndexMetaGlobalInput) SetNo(value bool) *ActionNotificationTemplateVariantIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateVariantIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationTemplateVariantIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationTemplateVariantIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateVariantIndexInput is a type for action input parameters
type ActionNotificationTemplateVariantIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNotificationTemplateVariantIndexInput) SetFromId(value int64) *ActionNotificationTemplateVariantIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNotificationTemplateVariantIndexInput) SetLimit(value int64) *ActionNotificationTemplateVariantIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationTemplateVariantIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantIndexInput) SelectParameters(params ...string) *ActionNotificationTemplateVariantIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationTemplateVariantIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationTemplateVariantIndexInput) UnselectParameters(params ...string) *ActionNotificationTemplateVariantIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationTemplateVariantIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationTemplateVariantIndexOutput is a type for action output parameters
type ActionNotificationTemplateVariantIndexOutput struct {
	CreatedAt  string                    "json:\"created_at\""
	From       string                    "json:\"from\""
	Html       string                    "json:\"html\""
	Id         int64                     "json:\"id\""
	Language   *ActionLanguageShowOutput "json:\"language\""
	Options    interface{}               "json:\"options\""
	Protocol   string                    "json:\"protocol\""
	ReplyTo    string                    "json:\"reply_to\""
	ReturnPath string                    "json:\"return_path\""
	Subject    string                    "json:\"subject\""
	Text       string                    "json:\"text\""
	UpdatedAt  string                    "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionNotificationTemplateVariantIndexResponse struct {
	Action *ActionNotificationTemplateVariantIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Variants []*ActionNotificationTemplateVariantIndexOutput "json:\"variants\""
	}

	// Action output without the namespace
	Output []*ActionNotificationTemplateVariantIndexOutput
}

// Prepare the action for invocation
func (action *ActionNotificationTemplateVariantIndex) Prepare() *ActionNotificationTemplateVariantIndexInvocation {
	return &ActionNotificationTemplateVariantIndexInvocation{
		Action: action,
		Path:   "/v7.0/notification_templates/{notification_template_id}/variants",
	}
}

// ActionNotificationTemplateVariantIndexInvocation is used to configure action for invocation
type ActionNotificationTemplateVariantIndexInvocation struct {
	// Pointer to the action
	Action *ActionNotificationTemplateVariantIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationTemplateVariantIndexInput
	// Global meta input parameters
	MetaInput *ActionNotificationTemplateVariantIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNotificationTemplateVariantIndexInvocation) SetPathParamInt(param string, value int64) *ActionNotificationTemplateVariantIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNotificationTemplateVariantIndexInvocation) SetPathParamString(param string, value string) *ActionNotificationTemplateVariantIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationTemplateVariantIndexInvocation) NewInput() *ActionNotificationTemplateVariantIndexInput {
	inv.Input = &ActionNotificationTemplateVariantIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationTemplateVariantIndexInvocation) SetInput(input *ActionNotificationTemplateVariantIndexInput) *ActionNotificationTemplateVariantIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationTemplateVariantIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationTemplateVariantIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationTemplateVariantIndexInvocation) NewMetaInput() *ActionNotificationTemplateVariantIndexMetaGlobalInput {
	inv.MetaInput = &ActionNotificationTemplateVariantIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationTemplateVariantIndexInvocation) SetMetaInput(input *ActionNotificationTemplateVariantIndexMetaGlobalInput) *ActionNotificationTemplateVariantIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationTemplateVariantIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationTemplateVariantIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationTemplateVariantIndexInvocation) validate() error {
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
func (inv *ActionNotificationTemplateVariantIndexInvocation) Call() (*ActionNotificationTemplateVariantIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNotificationTemplateVariantIndexInvocation) callAsQuery() (*ActionNotificationTemplateVariantIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNotificationTemplateVariantIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Variants
	}
	return resp, err
}

func (inv *ActionNotificationTemplateVariantIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["variant[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["variant[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}

	return nil
}

func (inv *ActionNotificationTemplateVariantIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
