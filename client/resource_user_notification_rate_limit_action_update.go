package client

import (
	"net/url"
	"strings"
)

// ActionUserNotificationRateLimitUpdate is a type for action User.Notification_rate_limit#Update
type ActionUserNotificationRateLimitUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNotificationRateLimitUpdate(client *Client) *ActionUserNotificationRateLimitUpdate {
	return &ActionUserNotificationRateLimitUpdate{
		Client: client,
	}
}

// ActionUserNotificationRateLimitUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserNotificationRateLimitUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNotificationRateLimitUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserNotificationRateLimitUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNotificationRateLimitUpdateMetaGlobalInput) SetNo(value bool) *ActionUserNotificationRateLimitUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationRateLimitUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationRateLimitUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserNotificationRateLimitUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNotificationRateLimitUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationRateLimitUpdateInput is a type for action input parameters
type ActionUserNotificationRateLimitUpdateInput struct {
	LimitCount int64 "json:\"limit_count\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimitCount sets parameter LimitCount to value and selects it for sending
func (in *ActionUserNotificationRateLimitUpdateInput) SetLimitCount(value int64) *ActionUserNotificationRateLimitUpdateInput {
	in.LimitCount = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["LimitCount"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNotificationRateLimitUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNotificationRateLimitUpdateInput) SelectParameters(params ...string) *ActionUserNotificationRateLimitUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserNotificationRateLimitUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserNotificationRateLimitUpdateInput) UnselectParameters(params ...string) *ActionUserNotificationRateLimitUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserNotificationRateLimitUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNotificationRateLimitUpdateRequest is a type for the entire action request
type ActionUserNotificationRateLimitUpdateRequest struct {
	NotificationRateLimit map[string]interface{} "json:\"notification_rate_limit\""
	Meta                  map[string]interface{} "json:\"_meta\""
}

// ActionUserNotificationRateLimitUpdateOutput is a type for action output parameters
type ActionUserNotificationRateLimitUpdateOutput struct {
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
type ActionUserNotificationRateLimitUpdateResponse struct {
	Action *ActionUserNotificationRateLimitUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationRateLimit *ActionUserNotificationRateLimitUpdateOutput "json:\"notification_rate_limit\""
	}

	// Action output without the namespace
	Output *ActionUserNotificationRateLimitUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserNotificationRateLimitUpdate) Prepare() *ActionUserNotificationRateLimitUpdateInvocation {
	return &ActionUserNotificationRateLimitUpdateInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/notification_rate_limits/{notification_rate_limit_id}",
	}
}

// ActionUserNotificationRateLimitUpdateInvocation is used to configure action for invocation
type ActionUserNotificationRateLimitUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserNotificationRateLimitUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNotificationRateLimitUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserNotificationRateLimitUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNotificationRateLimitUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserNotificationRateLimitUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNotificationRateLimitUpdateInvocation) SetPathParamString(param string, value string) *ActionUserNotificationRateLimitUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNotificationRateLimitUpdateInvocation) NewInput() *ActionUserNotificationRateLimitUpdateInput {
	inv.Input = &ActionUserNotificationRateLimitUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNotificationRateLimitUpdateInvocation) SetInput(input *ActionUserNotificationRateLimitUpdateInput) *ActionUserNotificationRateLimitUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNotificationRateLimitUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserNotificationRateLimitUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNotificationRateLimitUpdateInvocation) NewMetaInput() *ActionUserNotificationRateLimitUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserNotificationRateLimitUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNotificationRateLimitUpdateInvocation) SetMetaInput(input *ActionUserNotificationRateLimitUpdateMetaGlobalInput) *ActionUserNotificationRateLimitUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNotificationRateLimitUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNotificationRateLimitUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserNotificationRateLimitUpdateInvocation) validate() error {
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
func (inv *ActionUserNotificationRateLimitUpdateInvocation) Call() (*ActionUserNotificationRateLimitUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionUserNotificationRateLimitUpdateInvocation) callAsBody() (*ActionUserNotificationRateLimitUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserNotificationRateLimitUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationRateLimit
	}
	return resp, err
}

func (inv *ActionUserNotificationRateLimitUpdateInvocation) makeAllInputParams() *ActionUserNotificationRateLimitUpdateRequest {
	return &ActionUserNotificationRateLimitUpdateRequest{
		NotificationRateLimit: inv.makeInputParams(),
		Meta:                  inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNotificationRateLimitUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("LimitCount") {
			ret["limit_count"] = inv.Input.LimitCount
		}
	}

	return ret
}

func (inv *ActionUserNotificationRateLimitUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
