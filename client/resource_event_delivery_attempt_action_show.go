package client

import (
	"net/url"
	"strings"
)

// ActionEventDeliveryAttemptShow is a type for action Event.Delivery.Attempt#Show
type ActionEventDeliveryAttemptShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEventDeliveryAttemptShow(client *Client) *ActionEventDeliveryAttemptShow {
	return &ActionEventDeliveryAttemptShow{
		Client: client,
	}
}

// ActionEventDeliveryAttemptShowMetaGlobalInput is a type for action global meta input parameters
type ActionEventDeliveryAttemptShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventDeliveryAttemptShowMetaGlobalInput) SetIncludes(value string) *ActionEventDeliveryAttemptShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventDeliveryAttemptShowMetaGlobalInput) SetNo(value bool) *ActionEventDeliveryAttemptShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventDeliveryAttemptShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryAttemptShowMetaGlobalInput) SelectParameters(params ...string) *ActionEventDeliveryAttemptShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventDeliveryAttemptShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryAttemptShowOutput is a type for action output parameters
type ActionEventDeliveryAttemptShowOutput struct {
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
type ActionEventDeliveryAttemptShowResponse struct {
	Action *ActionEventDeliveryAttemptShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Attempt *ActionEventDeliveryAttemptShowOutput "json:\"attempt\""
	}

	// Action output without the namespace
	Output *ActionEventDeliveryAttemptShowOutput
}

// Prepare the action for invocation
func (action *ActionEventDeliveryAttemptShow) Prepare() *ActionEventDeliveryAttemptShowInvocation {
	return &ActionEventDeliveryAttemptShowInvocation{
		Action: action,
		Path:   "/v7.0/events/{event_id}/deliveries/{delivery_id}/attempts/{attempt_id}",
	}
}

// ActionEventDeliveryAttemptShowInvocation is used to configure action for invocation
type ActionEventDeliveryAttemptShowInvocation struct {
	// Pointer to the action
	Action *ActionEventDeliveryAttemptShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventDeliveryAttemptShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventDeliveryAttemptShowInvocation) SetPathParamInt(param string, value int64) *ActionEventDeliveryAttemptShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventDeliveryAttemptShowInvocation) SetPathParamString(param string, value string) *ActionEventDeliveryAttemptShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventDeliveryAttemptShowInvocation) NewMetaInput() *ActionEventDeliveryAttemptShowMetaGlobalInput {
	inv.MetaInput = &ActionEventDeliveryAttemptShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventDeliveryAttemptShowInvocation) SetMetaInput(input *ActionEventDeliveryAttemptShowMetaGlobalInput) *ActionEventDeliveryAttemptShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventDeliveryAttemptShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventDeliveryAttemptShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventDeliveryAttemptShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventDeliveryAttemptShowInvocation) Call() (*ActionEventDeliveryAttemptShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventDeliveryAttemptShowInvocation) callAsQuery() (*ActionEventDeliveryAttemptShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventDeliveryAttemptShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Attempt
	}
	return resp, err
}

func (inv *ActionEventDeliveryAttemptShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
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
