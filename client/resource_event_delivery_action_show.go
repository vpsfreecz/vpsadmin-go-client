package client

import (
	"net/url"
	"strings"
)

// ActionEventDeliveryShow is a type for action Event.Delivery#Show
type ActionEventDeliveryShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEventDeliveryShow(client *Client) *ActionEventDeliveryShow {
	return &ActionEventDeliveryShow{
		Client: client,
	}
}

// ActionEventDeliveryShowMetaGlobalInput is a type for action global meta input parameters
type ActionEventDeliveryShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventDeliveryShowMetaGlobalInput) SetIncludes(value string) *ActionEventDeliveryShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventDeliveryShowMetaGlobalInput) SetNo(value bool) *ActionEventDeliveryShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventDeliveryShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryShowMetaGlobalInput) SelectParameters(params ...string) *ActionEventDeliveryShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventDeliveryShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryShowOutput is a type for action output parameters
type ActionEventDeliveryShowOutput struct {
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
	MailCc                                  string "json:\"mail_cc\""
	MailFrom                                string "json:\"mail_from\""
	MailLogId                               int64  "json:\"mail_log_id\""
	MailMessageId                           string "json:\"mail_message_id\""
	MailReplyTo                             string "json:\"mail_reply_to\""
	MailReturnPath                          string "json:\"mail_return_path\""
	MailSubject                             string "json:\"mail_subject\""
	MailTextHtml                            string "json:\"mail_text_html\""
	MailTextPlain                           string "json:\"mail_text_plain\""
	MailTo                                  string "json:\"mail_to\""
	NextAttemptAt                           string "json:\"next_attempt_at\""
	NotificationReceiverActionDisplayTarget string "json:\"notification_receiver_action_display_target\""
	NotificationReceiverActionLabel         string "json:\"notification_receiver_action_label\""
	NotificationReceiverId                  int64  "json:\"notification_receiver_id\""
	NotificationReceiverLabel               string "json:\"notification_receiver_label\""
	NotificationReceiverTargetId            int64  "json:\"notification_receiver_target_id\""
	NotificationTargetDisplayTarget         string "json:\"notification_target_display_target\""
	NotificationTargetId                    int64  "json:\"notification_target_id\""
	NotificationTargetLabel                 string "json:\"notification_target_label\""
	Payload                                 string "json:\"payload\""
	ProviderMessageId                       string "json:\"provider_message_id\""
	RecipientUserId                         int64  "json:\"recipient_user_id\""
	RecipientUserLogin                      string "json:\"recipient_user_login\""
	ReleasedAt                              string "json:\"released_at\""
	ResponseBody                            string "json:\"response_body\""
	ResponseHeadersJson                     string "json:\"response_headers_json\""
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
type ActionEventDeliveryShowResponse struct {
	Action *ActionEventDeliveryShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Delivery *ActionEventDeliveryShowOutput "json:\"delivery\""
	}

	// Action output without the namespace
	Output *ActionEventDeliveryShowOutput
}

// Prepare the action for invocation
func (action *ActionEventDeliveryShow) Prepare() *ActionEventDeliveryShowInvocation {
	return &ActionEventDeliveryShowInvocation{
		Action: action,
		Path:   "/v7.0/events/{event_id}/deliveries/{delivery_id}",
	}
}

// ActionEventDeliveryShowInvocation is used to configure action for invocation
type ActionEventDeliveryShowInvocation struct {
	// Pointer to the action
	Action *ActionEventDeliveryShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventDeliveryShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventDeliveryShowInvocation) SetPathParamInt(param string, value int64) *ActionEventDeliveryShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventDeliveryShowInvocation) SetPathParamString(param string, value string) *ActionEventDeliveryShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventDeliveryShowInvocation) NewMetaInput() *ActionEventDeliveryShowMetaGlobalInput {
	inv.MetaInput = &ActionEventDeliveryShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventDeliveryShowInvocation) SetMetaInput(input *ActionEventDeliveryShowMetaGlobalInput) *ActionEventDeliveryShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventDeliveryShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventDeliveryShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventDeliveryShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventDeliveryShowInvocation) Call() (*ActionEventDeliveryShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventDeliveryShowInvocation) callAsQuery() (*ActionEventDeliveryShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEventDeliveryShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Delivery
	}
	return resp, err
}

func (inv *ActionEventDeliveryShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
