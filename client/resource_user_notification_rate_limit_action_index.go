package client

import (
	"net/url"
	"strings"
)

// ActionUserNotificationRateLimitIndex is a type for action User.Notification_rate_limit#Index
type ActionUserNotificationRateLimitIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNotificationRateLimitIndex(client *Client) *ActionUserNotificationRateLimitIndex {
	return &ActionUserNotificationRateLimitIndex{
		Client: client,
	}
}

// ActionUserNotificationRateLimitIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserNotificationRateLimitIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserNotificationRateLimitIndexMetaGlobalInput) SetCount(value bool) *ActionUserNotificationRateLimitIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNotificationRateLimitIndexMetaGlobalInput) SetIncludes(value string) *ActionUserNotificationRateLimitIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNotificationRateLimitIndexMetaGlobalInput) SetNo(value bool) *ActionUserNotificationRateLimitIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationRateLimitIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationRateLimitIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserNotificationRateLimitIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNotificationRateLimitIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationRateLimitIndexInput is a type for action input parameters
type ActionUserNotificationRateLimitIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserNotificationRateLimitIndexInput) SetFromId(value int64) *ActionUserNotificationRateLimitIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserNotificationRateLimitIndexInput) SetLimit(value int64) *ActionUserNotificationRateLimitIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationRateLimitIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationRateLimitIndexInput) SelectParameters(params ...string) *ActionUserNotificationRateLimitIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserNotificationRateLimitIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserNotificationRateLimitIndexInput) UnselectParameters(params ...string) *ActionUserNotificationRateLimitIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserNotificationRateLimitIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationRateLimitIndexOutput is a type for action output parameters
type ActionUserNotificationRateLimitIndexOutput struct {
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
type ActionUserNotificationRateLimitIndexResponse struct {
	Action *ActionUserNotificationRateLimitIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationRateLimits []*ActionUserNotificationRateLimitIndexOutput "json:\"notification_rate_limits\""
	}

	// Action output without the namespace
	Output []*ActionUserNotificationRateLimitIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserNotificationRateLimitIndex) Prepare() *ActionUserNotificationRateLimitIndexInvocation {
	return &ActionUserNotificationRateLimitIndexInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/notification_rate_limits",
	}
}

// ActionUserNotificationRateLimitIndexInvocation is used to configure action for invocation
type ActionUserNotificationRateLimitIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserNotificationRateLimitIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNotificationRateLimitIndexInput
	// Global meta input parameters
	MetaInput *ActionUserNotificationRateLimitIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNotificationRateLimitIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserNotificationRateLimitIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNotificationRateLimitIndexInvocation) SetPathParamString(param string, value string) *ActionUserNotificationRateLimitIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNotificationRateLimitIndexInvocation) NewInput() *ActionUserNotificationRateLimitIndexInput {
	inv.Input = &ActionUserNotificationRateLimitIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNotificationRateLimitIndexInvocation) SetInput(input *ActionUserNotificationRateLimitIndexInput) *ActionUserNotificationRateLimitIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNotificationRateLimitIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserNotificationRateLimitIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNotificationRateLimitIndexInvocation) NewMetaInput() *ActionUserNotificationRateLimitIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserNotificationRateLimitIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNotificationRateLimitIndexInvocation) SetMetaInput(input *ActionUserNotificationRateLimitIndexMetaGlobalInput) *ActionUserNotificationRateLimitIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNotificationRateLimitIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNotificationRateLimitIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserNotificationRateLimitIndexInvocation) validate() error {
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
func (inv *ActionUserNotificationRateLimitIndexInvocation) Call() (*ActionUserNotificationRateLimitIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionUserNotificationRateLimitIndexInvocation) callAsQuery() (*ActionUserNotificationRateLimitIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNotificationRateLimitIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationRateLimits
	}
	return resp, err
}

func (inv *ActionUserNotificationRateLimitIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["notification_rate_limit[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["notification_rate_limit[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserNotificationRateLimitIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
