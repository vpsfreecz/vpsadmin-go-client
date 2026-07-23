package client

import ()

// ActionEventDeliveryIndex is a type for action Event_delivery#Index
type ActionEventDeliveryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventDeliveryIndex(client *Client) *ActionEventDeliveryIndex {
	return &ActionEventDeliveryIndex{
		Client: client,
	}
}

// ActionEventDeliveryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventDeliveryIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventDeliveryIndexMetaGlobalInput) SetCount(value bool) *ActionEventDeliveryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventDeliveryIndexMetaGlobalInput) SetIncludes(value string) *ActionEventDeliveryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventDeliveryIndexMetaGlobalInput) SetNo(value bool) *ActionEventDeliveryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventDeliveryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventDeliveryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventDeliveryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryIndexInput is a type for action input parameters
type ActionEventDeliveryIndexInput struct {
	Action                       string "json:\"action\""
	EventRouteId                 int64  "json:\"event_route_id\""
	EventType                    string "json:\"event_type\""
	FromId                       int64  "json:\"from_id\""
	Limit                        int64  "json:\"limit\""
	NotificationReceiverId       int64  "json:\"notification_receiver_id\""
	NotificationReceiverTargetId int64  "json:\"notification_receiver_target_id\""
	NotificationTargetId         int64  "json:\"notification_target_id\""
	RecipientUser                int64  "json:\"recipient_user\""
	State                        string "json:\"state\""
	StateGroup                   string "json:\"state_group\""
	User                         int64  "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAction sets parameter Action to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetAction(value string) *ActionEventDeliveryIndexInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetActionNil(false)
	in._selectedParameters["Action"] = nil
	return in
}

// SetActionNil sets parameter Action to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetActionNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Action"] = nil
		in.SelectParameters("Action")
	} else {
		delete(in._nilParameters, "Action")
	}
	return in
}

// SetEventRouteId sets parameter EventRouteId to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetEventRouteId(value int64) *ActionEventDeliveryIndexInput {
	in.EventRouteId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventRouteIdNil(false)
	in._selectedParameters["EventRouteId"] = nil
	return in
}

// SetEventRouteIdNil sets parameter EventRouteId to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetEventRouteIdNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["EventRouteId"] = nil
		in.SelectParameters("EventRouteId")
	} else {
		delete(in._nilParameters, "EventRouteId")
	}
	return in
}

// SetEventType sets parameter EventType to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetEventType(value string) *ActionEventDeliveryIndexInput {
	in.EventType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventTypeNil(false)
	in._selectedParameters["EventType"] = nil
	return in
}

// SetEventTypeNil sets parameter EventType to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetEventTypeNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["EventType"] = nil
		in.SelectParameters("EventType")
	} else {
		delete(in._nilParameters, "EventType")
	}
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetFromId(value int64) *ActionEventDeliveryIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetLimit(value int64) *ActionEventDeliveryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNotificationReceiverId sets parameter NotificationReceiverId to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetNotificationReceiverId(value int64) *ActionEventDeliveryIndexInput {
	in.NotificationReceiverId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationReceiverIdNil(false)
	in._selectedParameters["NotificationReceiverId"] = nil
	return in
}

// SetNotificationReceiverIdNil sets parameter NotificationReceiverId to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetNotificationReceiverIdNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["NotificationReceiverId"] = nil
		in.SelectParameters("NotificationReceiverId")
	} else {
		delete(in._nilParameters, "NotificationReceiverId")
	}
	return in
}

// SetNotificationReceiverTargetId sets parameter NotificationReceiverTargetId to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetNotificationReceiverTargetId(value int64) *ActionEventDeliveryIndexInput {
	in.NotificationReceiverTargetId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationReceiverTargetIdNil(false)
	in._selectedParameters["NotificationReceiverTargetId"] = nil
	return in
}

// SetNotificationReceiverTargetIdNil sets parameter NotificationReceiverTargetId to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetNotificationReceiverTargetIdNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["NotificationReceiverTargetId"] = nil
		in.SelectParameters("NotificationReceiverTargetId")
	} else {
		delete(in._nilParameters, "NotificationReceiverTargetId")
	}
	return in
}

// SetNotificationTargetId sets parameter NotificationTargetId to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetNotificationTargetId(value int64) *ActionEventDeliveryIndexInput {
	in.NotificationTargetId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationTargetIdNil(false)
	in._selectedParameters["NotificationTargetId"] = nil
	return in
}

// SetNotificationTargetIdNil sets parameter NotificationTargetId to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetNotificationTargetIdNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["NotificationTargetId"] = nil
		in.SelectParameters("NotificationTargetId")
	} else {
		delete(in._nilParameters, "NotificationTargetId")
	}
	return in
}

// SetRecipientUser sets parameter RecipientUser to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetRecipientUser(value int64) *ActionEventDeliveryIndexInput {
	in.RecipientUser = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetRecipientUserNil(false)
	in._selectedParameters["RecipientUser"] = nil
	return in
}

// SetRecipientUserNil sets parameter RecipientUser to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetRecipientUserNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["RecipientUser"] = nil
		in.SelectParameters("RecipientUser")
	} else {
		delete(in._nilParameters, "RecipientUser")
	}
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetState(value string) *ActionEventDeliveryIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetStateNil(false)
	in._selectedParameters["State"] = nil
	return in
}

// SetStateNil sets parameter State to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetStateNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["State"] = nil
		in.SelectParameters("State")
	} else {
		delete(in._nilParameters, "State")
	}
	return in
}

// SetStateGroup sets parameter StateGroup to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetStateGroup(value string) *ActionEventDeliveryIndexInput {
	in.StateGroup = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetStateGroupNil(false)
	in._selectedParameters["StateGroup"] = nil
	return in
}

// SetStateGroupNil sets parameter StateGroup to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetStateGroupNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["StateGroup"] = nil
		in.SelectParameters("StateGroup")
	} else {
		delete(in._nilParameters, "StateGroup")
	}
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetUser(value int64) *ActionEventDeliveryIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionEventDeliveryIndexInput) SetUserNil(set bool) *ActionEventDeliveryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionEventDeliveryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryIndexInput) SelectParameters(params ...string) *ActionEventDeliveryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventDeliveryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventDeliveryIndexInput) UnselectParameters(params ...string) *ActionEventDeliveryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventDeliveryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryIndexOutput is a type for action output parameters
type ActionEventDeliveryIndexOutput struct {
	Action                                  string "json:\"action\""
	AttemptCount                            int64  "json:\"attempt_count\""
	CreatedAt                               string "json:\"created_at\""
	DeliveryTransactionChainId              int64  "json:\"delivery_transaction_chain_id\""
	DeliveryTransactionChainLabel           string "json:\"delivery_transaction_chain_label\""
	ErrorSummary                            string "json:\"error_summary\""
	EventCreatedAt                          string "json:\"event_created_at\""
	EventId                                 int64  "json:\"event_id\""
	EventRouteId                            int64  "json:\"event_route_id\""
	EventRouteLabel                         string "json:\"event_route_label\""
	EventRoutingContextId                   int64  "json:\"event_routing_context_id\""
	EventSeverity                           string "json:\"event_severity\""
	EventSubject                            string "json:\"event_subject\""
	EventType                               string "json:\"event_type\""
	EventUserId                             int64  "json:\"event_user_id\""
	EventUserLogin                          string "json:\"event_user_login\""
	EventVpsHostname                        string "json:\"event_vps_hostname\""
	EventVpsId                              int64  "json:\"event_vps_id\""
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
type ActionEventDeliveryIndexResponse struct {
	Action *ActionEventDeliveryIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventDeliveries []*ActionEventDeliveryIndexOutput "json:\"event_deliveries\""
	}

	// Action output without the namespace
	Output []*ActionEventDeliveryIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventDeliveryIndex) Prepare() *ActionEventDeliveryIndexInvocation {
	return &ActionEventDeliveryIndexInvocation{
		Action: action,
		Path:   "/v7.0/event_deliveries",
	}
}

// ActionEventDeliveryIndexInvocation is used to configure action for invocation
type ActionEventDeliveryIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventDeliveryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventDeliveryIndexInput
	// Global meta input parameters
	MetaInput *ActionEventDeliveryIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventDeliveryIndexInvocation) NewInput() *ActionEventDeliveryIndexInput {
	inv.Input = &ActionEventDeliveryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventDeliveryIndexInvocation) SetInput(input *ActionEventDeliveryIndexInput) *ActionEventDeliveryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventDeliveryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventDeliveryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventDeliveryIndexInvocation) NewMetaInput() *ActionEventDeliveryIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventDeliveryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventDeliveryIndexInvocation) SetMetaInput(input *ActionEventDeliveryIndexMetaGlobalInput) *ActionEventDeliveryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventDeliveryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventDeliveryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventDeliveryIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("RecipientUser") {
			if !inv.IsParameterNil("RecipientUser") {
				if inv.Input.RecipientUser < 0 {
					verr.Add("recipient_user", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventDeliveryIndexInvocation) Call() (*ActionEventDeliveryIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventDeliveryIndexInvocation) callAsQuery() (*ActionEventDeliveryIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEventDeliveryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventDeliveries
	}
	return resp, err
}

func (inv *ActionEventDeliveryIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Action") {
			if inv.IsParameterNil("Action") {
				ret["event_delivery[action]"] = ""
			} else {
				ret["event_delivery[action]"] = inv.Input.Action
			}
		}
		if inv.IsParameterSelected("EventRouteId") {
			if inv.IsParameterNil("EventRouteId") {
				ret["event_delivery[event_route_id]"] = ""
			} else {
				ret["event_delivery[event_route_id]"] = convertInt64ToString(inv.Input.EventRouteId)
			}
		}
		if inv.IsParameterSelected("EventType") {
			if inv.IsParameterNil("EventType") {
				ret["event_delivery[event_type]"] = ""
			} else {
				ret["event_delivery[event_type]"] = inv.Input.EventType
			}
		}
		if inv.IsParameterSelected("FromId") {
			ret["event_delivery[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["event_delivery[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("NotificationReceiverId") {
			if inv.IsParameterNil("NotificationReceiverId") {
				ret["event_delivery[notification_receiver_id]"] = ""
			} else {
				ret["event_delivery[notification_receiver_id]"] = convertInt64ToString(inv.Input.NotificationReceiverId)
			}
		}
		if inv.IsParameterSelected("NotificationReceiverTargetId") {
			if inv.IsParameterNil("NotificationReceiverTargetId") {
				ret["event_delivery[notification_receiver_target_id]"] = ""
			} else {
				ret["event_delivery[notification_receiver_target_id]"] = convertInt64ToString(inv.Input.NotificationReceiverTargetId)
			}
		}
		if inv.IsParameterSelected("NotificationTargetId") {
			if inv.IsParameterNil("NotificationTargetId") {
				ret["event_delivery[notification_target_id]"] = ""
			} else {
				ret["event_delivery[notification_target_id]"] = convertInt64ToString(inv.Input.NotificationTargetId)
			}
		}
		if inv.IsParameterSelected("RecipientUser") {
			if inv.IsParameterNil("RecipientUser") {
				ret["event_delivery[recipient_user]"] = ""
			} else {
				ret["event_delivery[recipient_user]"] = convertInt64ToString(inv.Input.RecipientUser)
			}
		}
		if inv.IsParameterSelected("State") {
			if inv.IsParameterNil("State") {
				ret["event_delivery[state]"] = ""
			} else {
				ret["event_delivery[state]"] = inv.Input.State
			}
		}
		if inv.IsParameterSelected("StateGroup") {
			if inv.IsParameterNil("StateGroup") {
				ret["event_delivery[state_group]"] = ""
			} else {
				ret["event_delivery[state_group]"] = inv.Input.StateGroup
			}
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["event_delivery[user]"] = ""
			} else {
				ret["event_delivery[user]"] = convertInt64ToString(inv.Input.User)
			}
		}
	}
}

func (inv *ActionEventDeliveryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
