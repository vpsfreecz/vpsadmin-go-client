package client

import (
	"net/url"
	"strings"
)

// ActionEventDeliveryRetry is a type for action Event.Delivery#Retry
type ActionEventDeliveryRetry struct {
	// Pointer to client
	Client *Client
}

func NewActionEventDeliveryRetry(client *Client) *ActionEventDeliveryRetry {
	return &ActionEventDeliveryRetry{
		Client: client,
	}
}

// ActionEventDeliveryRetryMetaGlobalInput is a type for action global meta input parameters
type ActionEventDeliveryRetryMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventDeliveryRetryMetaGlobalInput) SetIncludes(value string) *ActionEventDeliveryRetryMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventDeliveryRetryMetaGlobalInput) SetNo(value bool) *ActionEventDeliveryRetryMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventDeliveryRetryMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryRetryMetaGlobalInput) SelectParameters(params ...string) *ActionEventDeliveryRetryMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventDeliveryRetryMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryRetryRequest is a type for the entire action request
type ActionEventDeliveryRetryRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// ActionEventDeliveryRetryOutput is a type for action output parameters
type ActionEventDeliveryRetryOutput struct {
	Action                                  string "json:\"action\""
	AttemptCount                            int64  "json:\"attempt_count\""
	CreatedAt                               string "json:\"created_at\""
	DeliveryTransactionChainId              int64  "json:\"delivery_transaction_chain_id\""
	DeliveryTransactionChainLabel           string "json:\"delivery_transaction_chain_label\""
	ErrorSummary                            string "json:\"error_summary\""
	EventRouteId                            int64  "json:\"event_route_id\""
	EventRouteLabel                         string "json:\"event_route_label\""
	EventRoutingContextId                   int64  "json:\"event_routing_context_id\""
	Id                                      int64  "json:\"id\""
	LastAttemptAt                           string "json:\"last_attempt_at\""
	MailLogId                               int64  "json:\"mail_log_id\""
	NextAttemptAt                           string "json:\"next_attempt_at\""
	NotificationReceiverActionDisplayTarget string "json:\"notification_receiver_action_display_target\""
	NotificationReceiverActionLabel         string "json:\"notification_receiver_action_label\""
	NotificationReceiverId                  int64  "json:\"notification_receiver_id\""
	NotificationReceiverLabel               string "json:\"notification_receiver_label\""
	NotificationReceiverTargetId            int64  "json:\"notification_receiver_target_id\""
	NotificationTargetDisplayTarget         string "json:\"notification_target_display_target\""
	NotificationTargetId                    int64  "json:\"notification_target_id\""
	NotificationTargetLabel                 string "json:\"notification_target_label\""
	ProviderMessageId                       string "json:\"provider_message_id\""
	RecipientUserId                         int64  "json:\"recipient_user_id\""
	RecipientUserLogin                      string "json:\"recipient_user_login\""
	ReleasedAt                              string "json:\"released_at\""
	ResponseBody                            string "json:\"response_body\""
	ResponseStatus                          int64  "json:\"response_status\""
	State                                   string "json:\"state\""
	TargetKind                              string "json:\"target_kind\""
	TargetLabel                             string "json:\"target_label\""
	TargetValue                             string "json:\"target_value\""
	TemplateName                            string "json:\"template_name\""
	TransactionId                           int64  "json:\"transaction_id\""
	UpdatedAt                               string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionEventDeliveryRetryResponse struct {
	Action *ActionEventDeliveryRetry "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Delivery *ActionEventDeliveryRetryOutput "json:\"delivery\""
	}

	// Action output without the namespace
	Output *ActionEventDeliveryRetryOutput
}

// Prepare the action for invocation
func (action *ActionEventDeliveryRetry) Prepare() *ActionEventDeliveryRetryInvocation {
	return &ActionEventDeliveryRetryInvocation{
		Action: action,
		Path:   "/v7.0/events/{event_id}/deliveries/{delivery_id}/retry",
	}
}

// ActionEventDeliveryRetryInvocation is used to configure action for invocation
type ActionEventDeliveryRetryInvocation struct {
	// Pointer to the action
	Action *ActionEventDeliveryRetry

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventDeliveryRetryMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventDeliveryRetryInvocation) SetPathParamInt(param string, value int64) *ActionEventDeliveryRetryInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventDeliveryRetryInvocation) SetPathParamString(param string, value string) *ActionEventDeliveryRetryInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventDeliveryRetryInvocation) NewMetaInput() *ActionEventDeliveryRetryMetaGlobalInput {
	inv.MetaInput = &ActionEventDeliveryRetryMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventDeliveryRetryInvocation) SetMetaInput(input *ActionEventDeliveryRetryMetaGlobalInput) *ActionEventDeliveryRetryInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventDeliveryRetryInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventDeliveryRetryInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventDeliveryRetryInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventDeliveryRetryInvocation) Call() (*ActionEventDeliveryRetryResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventDeliveryRetryInvocation) callAsBody() (*ActionEventDeliveryRetryResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventDeliveryRetryResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Delivery
	}
	return resp, err
}

func (inv *ActionEventDeliveryRetryInvocation) makeAllInputParams() *ActionEventDeliveryRetryRequest {
	return &ActionEventDeliveryRetryRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventDeliveryRetryInvocation) makeMetaInputParams() map[string]interface{} {
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
