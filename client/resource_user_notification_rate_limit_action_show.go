package client

import (
	"net/url"
	"strings"
)

// ActionUserNotificationRateLimitShow is a type for action User.Notification_rate_limit#Show
type ActionUserNotificationRateLimitShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNotificationRateLimitShow(client *Client) *ActionUserNotificationRateLimitShow {
	return &ActionUserNotificationRateLimitShow{
		Client: client,
	}
}

// ActionUserNotificationRateLimitShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserNotificationRateLimitShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNotificationRateLimitShowMetaGlobalInput) SetIncludes(value string) *ActionUserNotificationRateLimitShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNotificationRateLimitShowMetaGlobalInput) SetNo(value bool) *ActionUserNotificationRateLimitShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationRateLimitShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationRateLimitShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserNotificationRateLimitShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNotificationRateLimitShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationRateLimitShowOutput is a type for action output parameters
type ActionUserNotificationRateLimitShowOutput struct {
	CreatedAt          string "json:\"created_at\""
	DefaultLimitCount  int64  "json:\"default_limit_count\""
	DeliveryMethod     string "json:\"delivery_method\""
	Id                 string "json:\"id\""
	Label              string "json:\"label\""
	LimitCount         int64  "json:\"limit_count\""
	OverrideLimitCount int64  "json:\"override_limit_count\""
	Period             string "json:\"period\""
	PeriodLabel        string "json:\"period_label\""
	RemainingCount     int64  "json:\"remaining_count\""
	ResetsAt           string "json:\"resets_at\""
	Source             string "json:\"source\""
	UpdatedAt          string "json:\"updated_at\""
	UsedCount          int64  "json:\"used_count\""
}

// Type for action response, including envelope
type ActionUserNotificationRateLimitShowResponse struct {
	Action *ActionUserNotificationRateLimitShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationRateLimit *ActionUserNotificationRateLimitShowOutput "json:\"notification_rate_limit\""
	}

	// Action output without the namespace
	Output *ActionUserNotificationRateLimitShowOutput
}

// Prepare the action for invocation
func (action *ActionUserNotificationRateLimitShow) Prepare() *ActionUserNotificationRateLimitShowInvocation {
	return &ActionUserNotificationRateLimitShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/notification_rate_limits/{notification_rate_limit_id}",
	}
}

// ActionUserNotificationRateLimitShowInvocation is used to configure action for invocation
type ActionUserNotificationRateLimitShowInvocation struct {
	// Pointer to the action
	Action *ActionUserNotificationRateLimitShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserNotificationRateLimitShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNotificationRateLimitShowInvocation) SetPathParamInt(param string, value int64) *ActionUserNotificationRateLimitShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNotificationRateLimitShowInvocation) SetPathParamString(param string, value string) *ActionUserNotificationRateLimitShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNotificationRateLimitShowInvocation) NewMetaInput() *ActionUserNotificationRateLimitShowMetaGlobalInput {
	inv.MetaInput = &ActionUserNotificationRateLimitShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNotificationRateLimitShowInvocation) SetMetaInput(input *ActionUserNotificationRateLimitShowMetaGlobalInput) *ActionUserNotificationRateLimitShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNotificationRateLimitShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNotificationRateLimitShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserNotificationRateLimitShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNotificationRateLimitShowInvocation) Call() (*ActionUserNotificationRateLimitShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionUserNotificationRateLimitShowInvocation) callAsQuery() (*ActionUserNotificationRateLimitShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNotificationRateLimitShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationRateLimit
	}
	return resp, err
}

func (inv *ActionUserNotificationRateLimitShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
