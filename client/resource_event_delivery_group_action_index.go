package client

import ()

// ActionEventDeliveryGroupIndex is a type for action Event_delivery_group#Index
type ActionEventDeliveryGroupIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventDeliveryGroupIndex(client *Client) *ActionEventDeliveryGroupIndex {
	return &ActionEventDeliveryGroupIndex{
		Client: client,
	}
}

// ActionEventDeliveryGroupIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventDeliveryGroupIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventDeliveryGroupIndexMetaGlobalInput) SetCount(value bool) *ActionEventDeliveryGroupIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventDeliveryGroupIndexMetaGlobalInput) SetIncludes(value string) *ActionEventDeliveryGroupIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventDeliveryGroupIndexMetaGlobalInput) SetNo(value bool) *ActionEventDeliveryGroupIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventDeliveryGroupIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryGroupIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventDeliveryGroupIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventDeliveryGroupIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryGroupIndexInput is a type for action input parameters
type ActionEventDeliveryGroupIndexInput struct {
	EventRouteId           int64  "json:\"event_route_id\""
	FromId                 int64  "json:\"from_id\""
	Limit                  int64  "json:\"limit\""
	NotificationReceiverId int64  "json:\"notification_receiver_id\""
	RouteOwner             int64  "json:\"route_owner\""
	StateGroup             string "json:\"state_group\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEventRouteId sets parameter EventRouteId to value and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetEventRouteId(value int64) *ActionEventDeliveryGroupIndexInput {
	in.EventRouteId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventRouteIdNil(false)
	in._selectedParameters["EventRouteId"] = nil
	return in
}

// SetEventRouteIdNil sets parameter EventRouteId to nil and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetEventRouteIdNil(set bool) *ActionEventDeliveryGroupIndexInput {
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

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetFromId(value int64) *ActionEventDeliveryGroupIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetLimit(value int64) *ActionEventDeliveryGroupIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNotificationReceiverId sets parameter NotificationReceiverId to value and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetNotificationReceiverId(value int64) *ActionEventDeliveryGroupIndexInput {
	in.NotificationReceiverId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationReceiverIdNil(false)
	in._selectedParameters["NotificationReceiverId"] = nil
	return in
}

// SetNotificationReceiverIdNil sets parameter NotificationReceiverId to nil and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetNotificationReceiverIdNil(set bool) *ActionEventDeliveryGroupIndexInput {
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

// SetRouteOwner sets parameter RouteOwner to value and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetRouteOwner(value int64) *ActionEventDeliveryGroupIndexInput {
	in.RouteOwner = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetRouteOwnerNil(false)
	in._selectedParameters["RouteOwner"] = nil
	return in
}

// SetRouteOwnerNil sets parameter RouteOwner to nil and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetRouteOwnerNil(set bool) *ActionEventDeliveryGroupIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["RouteOwner"] = nil
		in.SelectParameters("RouteOwner")
	} else {
		delete(in._nilParameters, "RouteOwner")
	}
	return in
}

// SetStateGroup sets parameter StateGroup to value and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetStateGroup(value string) *ActionEventDeliveryGroupIndexInput {
	in.StateGroup = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetStateGroupNil(false)
	in._selectedParameters["StateGroup"] = nil
	return in
}

// SetStateGroupNil sets parameter StateGroup to nil and selects it for sending
func (in *ActionEventDeliveryGroupIndexInput) SetStateGroupNil(set bool) *ActionEventDeliveryGroupIndexInput {
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

// SelectParameters sets parameters from ActionEventDeliveryGroupIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryGroupIndexInput) SelectParameters(params ...string) *ActionEventDeliveryGroupIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventDeliveryGroupIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventDeliveryGroupIndexInput) UnselectParameters(params ...string) *ActionEventDeliveryGroupIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventDeliveryGroupIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryGroupIndexOutput is a type for action output parameters
type ActionEventDeliveryGroupIndexOutput struct {
	EventRouteId              int64       "json:\"event_route_id\""
	EventRouteLabel           string      "json:\"event_route_label\""
	GroupBy                   interface{} "json:\"group_by\""
	GroupIntervalSeconds      int64       "json:\"group_interval_seconds\""
	GroupWaitSeconds          int64       "json:\"group_wait_seconds\""
	Id                        int64       "json:\"id\""
	Labels                    interface{} "json:\"labels\""
	LastSealedAt              string      "json:\"last_sealed_at\""
	NextFlushAt               string      "json:\"next_flush_at\""
	NotificationReceiverId    int64       "json:\"notification_receiver_id\""
	NotificationReceiverLabel string      "json:\"notification_receiver_label\""
	PendingEventCount         int64       "json:\"pending_event_count\""
	RouteOwnerId              int64       "json:\"route_owner_id\""
	RouteOwnerLogin           string      "json:\"route_owner_login\""
	State                     string      "json:\"state\""
	StreamCount               int64       "json:\"stream_count\""
}

// Type for action response, including envelope
type ActionEventDeliveryGroupIndexResponse struct {
	Action *ActionEventDeliveryGroupIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventDeliveryGroups []*ActionEventDeliveryGroupIndexOutput "json:\"event_delivery_groups\""
	}

	// Action output without the namespace
	Output []*ActionEventDeliveryGroupIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventDeliveryGroupIndex) Prepare() *ActionEventDeliveryGroupIndexInvocation {
	return &ActionEventDeliveryGroupIndexInvocation{
		Action: action,
		Path:   "/v7.0/event_delivery_groups",
	}
}

// ActionEventDeliveryGroupIndexInvocation is used to configure action for invocation
type ActionEventDeliveryGroupIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventDeliveryGroupIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventDeliveryGroupIndexInput
	// Global meta input parameters
	MetaInput *ActionEventDeliveryGroupIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventDeliveryGroupIndexInvocation) NewInput() *ActionEventDeliveryGroupIndexInput {
	inv.Input = &ActionEventDeliveryGroupIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventDeliveryGroupIndexInvocation) SetInput(input *ActionEventDeliveryGroupIndexInput) *ActionEventDeliveryGroupIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventDeliveryGroupIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventDeliveryGroupIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventDeliveryGroupIndexInvocation) NewMetaInput() *ActionEventDeliveryGroupIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventDeliveryGroupIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventDeliveryGroupIndexInvocation) SetMetaInput(input *ActionEventDeliveryGroupIndexMetaGlobalInput) *ActionEventDeliveryGroupIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventDeliveryGroupIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventDeliveryGroupIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventDeliveryGroupIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("RouteOwner") {
			if !inv.IsParameterNil("RouteOwner") {
				if inv.Input.RouteOwner < 0 {
					verr.Add("route_owner", "not a valid resource id")
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
func (inv *ActionEventDeliveryGroupIndexInvocation) Call() (*ActionEventDeliveryGroupIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventDeliveryGroupIndexInvocation) callAsQuery() (*ActionEventDeliveryGroupIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventDeliveryGroupIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventDeliveryGroups
	}
	return resp, err
}

func (inv *ActionEventDeliveryGroupIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("EventRouteId") {
			if inv.IsParameterNil("EventRouteId") {
				ret["event_delivery_group[event_route_id]"] = ""
			} else {
				ret["event_delivery_group[event_route_id]"] = convertInt64ToString(inv.Input.EventRouteId)
			}
		}
		if inv.IsParameterSelected("FromId") {
			ret["event_delivery_group[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["event_delivery_group[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("NotificationReceiverId") {
			if inv.IsParameterNil("NotificationReceiverId") {
				ret["event_delivery_group[notification_receiver_id]"] = ""
			} else {
				ret["event_delivery_group[notification_receiver_id]"] = convertInt64ToString(inv.Input.NotificationReceiverId)
			}
		}
		if inv.IsParameterSelected("RouteOwner") {
			if inv.IsParameterNil("RouteOwner") {
				ret["event_delivery_group[route_owner]"] = ""
			} else {
				ret["event_delivery_group[route_owner]"] = convertInt64ToString(inv.Input.RouteOwner)
			}
		}
		if inv.IsParameterSelected("StateGroup") {
			if inv.IsParameterNil("StateGroup") {
				ret["event_delivery_group[state_group]"] = ""
			} else {
				ret["event_delivery_group[state_group]"] = inv.Input.StateGroup
			}
		}
	}

	return nil
}

func (inv *ActionEventDeliveryGroupIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
