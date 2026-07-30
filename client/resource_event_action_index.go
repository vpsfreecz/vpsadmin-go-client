package client

import ()

// ActionEventIndex is a type for action Event#Index
type ActionEventIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventIndex(client *Client) *ActionEventIndex {
	return &ActionEventIndex{
		Client: client,
	}
}

// ActionEventIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventIndexMetaGlobalInput) SetCount(value bool) *ActionEventIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventIndexMetaGlobalInput) SetIncludes(value string) *ActionEventIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventIndexMetaGlobalInput) SetNo(value bool) *ActionEventIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventIndexInput is a type for action input parameters
type ActionEventIndexInput struct {
	Action                       string "json:\"action\""
	Category                     string "json:\"category\""
	EventDeliveryGroup           int64  "json:\"event_delivery_group\""
	EventRouteId                 int64  "json:\"event_route_id\""
	EventType                    string "json:\"event_type\""
	FromId                       int64  "json:\"from_id\""
	GroupMembership              string "json:\"group_membership\""
	Limit                        int64  "json:\"limit\""
	NotificationReceiverId       int64  "json:\"notification_receiver_id\""
	NotificationReceiverTargetId int64  "json:\"notification_receiver_target_id\""
	NotificationTargetId         int64  "json:\"notification_target_id\""
	RoutingState                 string "json:\"routing_state\""
	Severity                     string "json:\"severity\""
	SubjectRelation              string "json:\"subject_relation\""
	User                         int64  "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAction sets parameter Action to value and selects it for sending
func (in *ActionEventIndexInput) SetAction(value string) *ActionEventIndexInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}

// SetCategory sets parameter Category to value and selects it for sending
func (in *ActionEventIndexInput) SetCategory(value string) *ActionEventIndexInput {
	in.Category = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Category"] = nil
	return in
}

// SetEventDeliveryGroup sets parameter EventDeliveryGroup to value and selects it for sending
func (in *ActionEventIndexInput) SetEventDeliveryGroup(value int64) *ActionEventIndexInput {
	in.EventDeliveryGroup = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventDeliveryGroupNil(false)
	in._selectedParameters["EventDeliveryGroup"] = nil
	return in
}

// SetEventDeliveryGroupNil sets parameter EventDeliveryGroup to nil and selects it for sending
func (in *ActionEventIndexInput) SetEventDeliveryGroupNil(set bool) *ActionEventIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["EventDeliveryGroup"] = nil
		in.SelectParameters("EventDeliveryGroup")
	} else {
		delete(in._nilParameters, "EventDeliveryGroup")
	}
	return in
}

// SetEventRouteId sets parameter EventRouteId to value and selects it for sending
func (in *ActionEventIndexInput) SetEventRouteId(value int64) *ActionEventIndexInput {
	in.EventRouteId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventRouteIdNil(false)
	in._selectedParameters["EventRouteId"] = nil
	return in
}

// SetEventRouteIdNil sets parameter EventRouteId to nil and selects it for sending
func (in *ActionEventIndexInput) SetEventRouteIdNil(set bool) *ActionEventIndexInput {
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
func (in *ActionEventIndexInput) SetEventType(value string) *ActionEventIndexInput {
	in.EventType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EventType"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEventIndexInput) SetFromId(value int64) *ActionEventIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetGroupMembership sets parameter GroupMembership to value and selects it for sending
func (in *ActionEventIndexInput) SetGroupMembership(value string) *ActionEventIndexInput {
	in.GroupMembership = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetGroupMembershipNil(false)
	in._selectedParameters["GroupMembership"] = nil
	return in
}

// SetGroupMembershipNil sets parameter GroupMembership to nil and selects it for sending
func (in *ActionEventIndexInput) SetGroupMembershipNil(set bool) *ActionEventIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["GroupMembership"] = nil
		in.SelectParameters("GroupMembership")
	} else {
		delete(in._nilParameters, "GroupMembership")
	}
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventIndexInput) SetLimit(value int64) *ActionEventIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNotificationReceiverId sets parameter NotificationReceiverId to value and selects it for sending
func (in *ActionEventIndexInput) SetNotificationReceiverId(value int64) *ActionEventIndexInput {
	in.NotificationReceiverId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationReceiverIdNil(false)
	in._selectedParameters["NotificationReceiverId"] = nil
	return in
}

// SetNotificationReceiverIdNil sets parameter NotificationReceiverId to nil and selects it for sending
func (in *ActionEventIndexInput) SetNotificationReceiverIdNil(set bool) *ActionEventIndexInput {
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
func (in *ActionEventIndexInput) SetNotificationReceiverTargetId(value int64) *ActionEventIndexInput {
	in.NotificationReceiverTargetId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationReceiverTargetIdNil(false)
	in._selectedParameters["NotificationReceiverTargetId"] = nil
	return in
}

// SetNotificationReceiverTargetIdNil sets parameter NotificationReceiverTargetId to nil and selects it for sending
func (in *ActionEventIndexInput) SetNotificationReceiverTargetIdNil(set bool) *ActionEventIndexInput {
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
func (in *ActionEventIndexInput) SetNotificationTargetId(value int64) *ActionEventIndexInput {
	in.NotificationTargetId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationTargetIdNil(false)
	in._selectedParameters["NotificationTargetId"] = nil
	return in
}

// SetNotificationTargetIdNil sets parameter NotificationTargetId to nil and selects it for sending
func (in *ActionEventIndexInput) SetNotificationTargetIdNil(set bool) *ActionEventIndexInput {
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

// SetRoutingState sets parameter RoutingState to value and selects it for sending
func (in *ActionEventIndexInput) SetRoutingState(value string) *ActionEventIndexInput {
	in.RoutingState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RoutingState"] = nil
	return in
}

// SetSeverity sets parameter Severity to value and selects it for sending
func (in *ActionEventIndexInput) SetSeverity(value string) *ActionEventIndexInput {
	in.Severity = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Severity"] = nil
	return in
}

// SetSubjectRelation sets parameter SubjectRelation to value and selects it for sending
func (in *ActionEventIndexInput) SetSubjectRelation(value string) *ActionEventIndexInput {
	in.SubjectRelation = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSubjectRelationNil(false)
	in._selectedParameters["SubjectRelation"] = nil
	return in
}

// SetSubjectRelationNil sets parameter SubjectRelation to nil and selects it for sending
func (in *ActionEventIndexInput) SetSubjectRelationNil(set bool) *ActionEventIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["SubjectRelation"] = nil
		in.SelectParameters("SubjectRelation")
	} else {
		delete(in._nilParameters, "SubjectRelation")
	}
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionEventIndexInput) SetUser(value int64) *ActionEventIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionEventIndexInput) SetUserNil(set bool) *ActionEventIndexInput {
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

// SelectParameters sets parameters from ActionEventIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventIndexInput) SelectParameters(params ...string) *ActionEventIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventIndexInput) UnselectParameters(params ...string) *ActionEventIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventIndexOutput is a type for action output parameters
type ActionEventIndexOutput struct {
	Category        string                "json:\"category\""
	CreatedAt       string                "json:\"created_at\""
	EventType       string                "json:\"event_type\""
	Id              int64                 "json:\"id\""
	IpAddr          string                "json:\"ip_addr\""
	PayloadJson     string                "json:\"payload_json\""
	RoutingState    string                "json:\"routing_state\""
	Severity        string                "json:\"severity\""
	SourceClass     string                "json:\"source_class\""
	SourceId        int64                 "json:\"source_id\""
	Subject         string                "json:\"subject\""
	SubjectRelation string                "json:\"subject_relation\""
	Summary         string                "json:\"summary\""
	UpdatedAt       string                "json:\"updated_at\""
	User            *ActionUserShowOutput "json:\"user\""
	Vps             *ActionVpsShowOutput  "json:\"vps\""
}

// Type for action response, including envelope
type ActionEventIndexResponse struct {
	Action *ActionEventIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Events []*ActionEventIndexOutput "json:\"events\""
	}

	// Action output without the namespace
	Output []*ActionEventIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventIndex) Prepare() *ActionEventIndexInvocation {
	return &ActionEventIndexInvocation{
		Action: action,
		Path:   "/v7.0/events",
	}
}

// ActionEventIndexInvocation is used to configure action for invocation
type ActionEventIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventIndexInput
	// Global meta input parameters
	MetaInput *ActionEventIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventIndexInvocation) NewInput() *ActionEventIndexInput {
	inv.Input = &ActionEventIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventIndexInvocation) SetInput(input *ActionEventIndexInput) *ActionEventIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventIndexInvocation) NewMetaInput() *ActionEventIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventIndexInvocation) SetMetaInput(input *ActionEventIndexMetaGlobalInput) *ActionEventIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("EventDeliveryGroup") {
			if !inv.IsParameterNil("EventDeliveryGroup") {
				if inv.Input.EventDeliveryGroup < 0 {
					verr.Add("event_delivery_group", "not a valid resource id")
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
func (inv *ActionEventIndexInvocation) Call() (*ActionEventIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventIndexInvocation) callAsQuery() (*ActionEventIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Events
	}
	return resp, err
}

func (inv *ActionEventIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("Action") {
			ret["event[action]"] = inv.Input.Action
		}
		if inv.IsParameterSelected("Category") {
			ret["event[category]"] = inv.Input.Category
		}
		if inv.IsParameterSelected("EventDeliveryGroup") {
			if inv.IsParameterNil("EventDeliveryGroup") {
				ret["event[event_delivery_group]"] = ""
			} else {
				ret["event[event_delivery_group]"] = convertInt64ToString(inv.Input.EventDeliveryGroup)
			}
		}
		if inv.IsParameterSelected("EventRouteId") {
			if inv.IsParameterNil("EventRouteId") {
				ret["event[event_route_id]"] = ""
			} else {
				ret["event[event_route_id]"] = convertInt64ToString(inv.Input.EventRouteId)
			}
		}
		if inv.IsParameterSelected("EventType") {
			ret["event[event_type]"] = inv.Input.EventType
		}
		if inv.IsParameterSelected("FromId") {
			ret["event[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("GroupMembership") {
			if inv.IsParameterNil("GroupMembership") {
				ret["event[group_membership]"] = ""
			} else {
				ret["event[group_membership]"] = inv.Input.GroupMembership
			}
		}
		if inv.IsParameterSelected("Limit") {
			ret["event[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("NotificationReceiverId") {
			if inv.IsParameterNil("NotificationReceiverId") {
				ret["event[notification_receiver_id]"] = ""
			} else {
				ret["event[notification_receiver_id]"] = convertInt64ToString(inv.Input.NotificationReceiverId)
			}
		}
		if inv.IsParameterSelected("NotificationReceiverTargetId") {
			if inv.IsParameterNil("NotificationReceiverTargetId") {
				ret["event[notification_receiver_target_id]"] = ""
			} else {
				ret["event[notification_receiver_target_id]"] = convertInt64ToString(inv.Input.NotificationReceiverTargetId)
			}
		}
		if inv.IsParameterSelected("NotificationTargetId") {
			if inv.IsParameterNil("NotificationTargetId") {
				ret["event[notification_target_id]"] = ""
			} else {
				ret["event[notification_target_id]"] = convertInt64ToString(inv.Input.NotificationTargetId)
			}
		}
		if inv.IsParameterSelected("RoutingState") {
			ret["event[routing_state]"] = inv.Input.RoutingState
		}
		if inv.IsParameterSelected("Severity") {
			ret["event[severity]"] = inv.Input.Severity
		}
		if inv.IsParameterSelected("SubjectRelation") {
			if inv.IsParameterNil("SubjectRelation") {
				ret["event[subject_relation]"] = ""
			} else {
				ret["event[subject_relation]"] = inv.Input.SubjectRelation
			}
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["event[user]"] = ""
			} else {
				ret["event[user]"] = convertInt64ToString(inv.Input.User)
			}
		}
	}

	return nil
}

func (inv *ActionEventIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
