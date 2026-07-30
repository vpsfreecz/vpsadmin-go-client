package client

import (
	"net/url"
	"strings"
)

// ActionEventDeliveryAttemptIndex is a type for action Event.Delivery.Attempt#Index
type ActionEventDeliveryAttemptIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventDeliveryAttemptIndex(client *Client) *ActionEventDeliveryAttemptIndex {
	return &ActionEventDeliveryAttemptIndex{
		Client: client,
	}
}

// ActionEventDeliveryAttemptIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventDeliveryAttemptIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventDeliveryAttemptIndexMetaGlobalInput) SetCount(value bool) *ActionEventDeliveryAttemptIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventDeliveryAttemptIndexMetaGlobalInput) SetIncludes(value string) *ActionEventDeliveryAttemptIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventDeliveryAttemptIndexMetaGlobalInput) SetNo(value bool) *ActionEventDeliveryAttemptIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventDeliveryAttemptIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryAttemptIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventDeliveryAttemptIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventDeliveryAttemptIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryAttemptIndexInput is a type for action input parameters
type ActionEventDeliveryAttemptIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEventDeliveryAttemptIndexInput) SetFromId(value int64) *ActionEventDeliveryAttemptIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventDeliveryAttemptIndexInput) SetLimit(value int64) *ActionEventDeliveryAttemptIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventDeliveryAttemptIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryAttemptIndexInput) SelectParameters(params ...string) *ActionEventDeliveryAttemptIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventDeliveryAttemptIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventDeliveryAttemptIndexInput) UnselectParameters(params ...string) *ActionEventDeliveryAttemptIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventDeliveryAttemptIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryAttemptIndexOutput is a type for action output parameters
type ActionEventDeliveryAttemptIndexOutput struct {
	Action              string "json:\"action\""
	AttemptNumber       int64  "json:\"attempt_number\""
	CreatedAt           string "json:\"created_at\""
	ErrorSummary        string "json:\"error_summary\""
	EventDeliveryId     int64  "json:\"event_delivery_id\""
	FinishedAt          string "json:\"finished_at\""
	Id                  int64  "json:\"id\""
	ProviderMessageId   string "json:\"provider_message_id\""
	RecipientUserId     int64  "json:\"recipient_user_id\""
	RecipientUserLogin  string "json:\"recipient_user_login\""
	ResponseBody        string "json:\"response_body\""
	ResponseHeadersJson string "json:\"response_headers_json\""
	ResponseStatus      int64  "json:\"response_status\""
	StartedAt           string "json:\"started_at\""
	State               string "json:\"state\""
	UpdatedAt           string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionEventDeliveryAttemptIndexResponse struct {
	Action *ActionEventDeliveryAttemptIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Attempts []*ActionEventDeliveryAttemptIndexOutput "json:\"attempts\""
	}

	// Action output without the namespace
	Output []*ActionEventDeliveryAttemptIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventDeliveryAttemptIndex) Prepare() *ActionEventDeliveryAttemptIndexInvocation {
	return &ActionEventDeliveryAttemptIndexInvocation{
		Action: action,
		Path:   "/v7.0/events/{event_id}/deliveries/{delivery_id}/attempts",
	}
}

// ActionEventDeliveryAttemptIndexInvocation is used to configure action for invocation
type ActionEventDeliveryAttemptIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventDeliveryAttemptIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventDeliveryAttemptIndexInput
	// Global meta input parameters
	MetaInput *ActionEventDeliveryAttemptIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventDeliveryAttemptIndexInvocation) SetPathParamInt(param string, value int64) *ActionEventDeliveryAttemptIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventDeliveryAttemptIndexInvocation) SetPathParamString(param string, value string) *ActionEventDeliveryAttemptIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventDeliveryAttemptIndexInvocation) NewInput() *ActionEventDeliveryAttemptIndexInput {
	inv.Input = &ActionEventDeliveryAttemptIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventDeliveryAttemptIndexInvocation) SetInput(input *ActionEventDeliveryAttemptIndexInput) *ActionEventDeliveryAttemptIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventDeliveryAttemptIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventDeliveryAttemptIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventDeliveryAttemptIndexInvocation) NewMetaInput() *ActionEventDeliveryAttemptIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventDeliveryAttemptIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventDeliveryAttemptIndexInvocation) SetMetaInput(input *ActionEventDeliveryAttemptIndexMetaGlobalInput) *ActionEventDeliveryAttemptIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventDeliveryAttemptIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventDeliveryAttemptIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventDeliveryAttemptIndexInvocation) validate() error {
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
func (inv *ActionEventDeliveryAttemptIndexInvocation) Call() (*ActionEventDeliveryAttemptIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventDeliveryAttemptIndexInvocation) callAsQuery() (*ActionEventDeliveryAttemptIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventDeliveryAttemptIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Attempts
	}
	return resp, err
}

func (inv *ActionEventDeliveryAttemptIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["attempt[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["attempt[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}

	return nil
}

func (inv *ActionEventDeliveryAttemptIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
