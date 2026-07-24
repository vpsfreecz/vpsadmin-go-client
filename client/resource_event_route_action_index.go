package client

import ()

// ActionEventRouteIndex is a type for action Event_route#Index
type ActionEventRouteIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteIndex(client *Client) *ActionEventRouteIndex {
	return &ActionEventRouteIndex{
		Client: client,
	}
}

// ActionEventRouteIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEventRouteIndexMetaGlobalInput) SetCount(value bool) *ActionEventRouteIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteIndexMetaGlobalInput) SetIncludes(value string) *ActionEventRouteIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteIndexMetaGlobalInput) SetNo(value bool) *ActionEventRouteIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteIndexInput is a type for action input parameters
type ActionEventRouteIndexInput struct {
	Enabled                bool   "json:\"enabled\""
	EventType              string "json:\"event_type\""
	FromId                 int64  "json:\"from_id\""
	IncludeSpent           bool   "json:\"include_spent\""
	Limit                  int64  "json:\"limit\""
	NotificationReceiverId int64  "json:\"notification_receiver_id\""
	ParentId               int64  "json:\"parent_id\""
	SubjectScope           string "json:\"subject_scope\""
	User                   int64  "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionEventRouteIndexInput) SetEnabled(value bool) *ActionEventRouteIndexInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetEventType sets parameter EventType to value and selects it for sending
func (in *ActionEventRouteIndexInput) SetEventType(value string) *ActionEventRouteIndexInput {
	in.EventType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventTypeNil(false)
	in._selectedParameters["EventType"] = nil
	return in
}

// SetEventTypeNil sets parameter EventType to nil and selects it for sending
func (in *ActionEventRouteIndexInput) SetEventTypeNil(set bool) *ActionEventRouteIndexInput {
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
func (in *ActionEventRouteIndexInput) SetFromId(value int64) *ActionEventRouteIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetIncludeSpent sets parameter IncludeSpent to value and selects it for sending
func (in *ActionEventRouteIndexInput) SetIncludeSpent(value bool) *ActionEventRouteIndexInput {
	in.IncludeSpent = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IncludeSpent"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEventRouteIndexInput) SetLimit(value int64) *ActionEventRouteIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNotificationReceiverId sets parameter NotificationReceiverId to value and selects it for sending
func (in *ActionEventRouteIndexInput) SetNotificationReceiverId(value int64) *ActionEventRouteIndexInput {
	in.NotificationReceiverId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationReceiverIdNil(false)
	in._selectedParameters["NotificationReceiverId"] = nil
	return in
}

// SetNotificationReceiverIdNil sets parameter NotificationReceiverId to nil and selects it for sending
func (in *ActionEventRouteIndexInput) SetNotificationReceiverIdNil(set bool) *ActionEventRouteIndexInput {
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

// SetParentId sets parameter ParentId to value and selects it for sending
func (in *ActionEventRouteIndexInput) SetParentId(value int64) *ActionEventRouteIndexInput {
	in.ParentId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetParentIdNil(false)
	in._selectedParameters["ParentId"] = nil
	return in
}

// SetParentIdNil sets parameter ParentId to nil and selects it for sending
func (in *ActionEventRouteIndexInput) SetParentIdNil(set bool) *ActionEventRouteIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["ParentId"] = nil
		in.SelectParameters("ParentId")
	} else {
		delete(in._nilParameters, "ParentId")
	}
	return in
}

// SetSubjectScope sets parameter SubjectScope to value and selects it for sending
func (in *ActionEventRouteIndexInput) SetSubjectScope(value string) *ActionEventRouteIndexInput {
	in.SubjectScope = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SubjectScope"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionEventRouteIndexInput) SetUser(value int64) *ActionEventRouteIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteIndexInput) SelectParameters(params ...string) *ActionEventRouteIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteIndexInput) UnselectParameters(params ...string) *ActionEventRouteIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteIndexOutput is a type for action output parameters
type ActionEventRouteIndexOutput struct {
	Continue               bool                  "json:\"continue\""
	CreatedAt              string                "json:\"created_at\""
	DisplayLabel           string                "json:\"display_label\""
	Enabled                bool                  "json:\"enabled\""
	EventType              string                "json:\"event_type\""
	EventTypePattern       string                "json:\"event_type_pattern\""
	ExpiresAt              string                "json:\"expires_at\""
	GroupBy                interface{}           "json:\"group_by\""
	GroupIntervalSeconds   int64                 "json:\"group_interval_seconds\""
	GroupWaitSeconds       int64                 "json:\"group_wait_seconds\""
	GroupingEnabled        bool                  "json:\"grouping_enabled\""
	GroupingSummary        string                "json:\"grouping_summary\""
	HitCount               int64                 "json:\"hit_count\""
	Id                     int64                 "json:\"id\""
	Label                  string                "json:\"label\""
	MatcherSummary         string                "json:\"matcher_summary\""
	NotificationReceiverId int64                 "json:\"notification_receiver_id\""
	ParentId               int64                 "json:\"parent_id\""
	Position               int64                 "json:\"position\""
	SingleUse              bool                  "json:\"single_use\""
	SpentAt                string                "json:\"spent_at\""
	SubjectScope           string                "json:\"subject_scope\""
	UpdatedAt              string                "json:\"updated_at\""
	User                   *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionEventRouteIndexResponse struct {
	Action *ActionEventRouteIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventRoutes []*ActionEventRouteIndexOutput "json:\"event_routes\""
	}

	// Action output without the namespace
	Output []*ActionEventRouteIndexOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteIndex) Prepare() *ActionEventRouteIndexInvocation {
	return &ActionEventRouteIndexInvocation{
		Action: action,
		Path:   "/v7.0/event_routes",
	}
}

// ActionEventRouteIndexInvocation is used to configure action for invocation
type ActionEventRouteIndexInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteIndexInput
	// Global meta input parameters
	MetaInput *ActionEventRouteIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteIndexInvocation) NewInput() *ActionEventRouteIndexInput {
	inv.Input = &ActionEventRouteIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteIndexInvocation) SetInput(input *ActionEventRouteIndexInput) *ActionEventRouteIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteIndexInvocation) NewMetaInput() *ActionEventRouteIndexMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteIndexInvocation) SetMetaInput(input *ActionEventRouteIndexMetaGlobalInput) *ActionEventRouteIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
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
func (inv *ActionEventRouteIndexInvocation) Call() (*ActionEventRouteIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventRouteIndexInvocation) callAsQuery() (*ActionEventRouteIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventRouteIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventRoutes
	}
	return resp, err
}

func (inv *ActionEventRouteIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["event_route[enabled]"] = convertBoolToString(inv.Input.Enabled)
		}
		if inv.IsParameterSelected("EventType") {
			if inv.IsParameterNil("EventType") {
				ret["event_route[event_type]"] = ""
			} else {
				ret["event_route[event_type]"] = inv.Input.EventType
			}
		}
		if inv.IsParameterSelected("FromId") {
			ret["event_route[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("IncludeSpent") {
			ret["event_route[include_spent]"] = convertBoolToString(inv.Input.IncludeSpent)
		}
		if inv.IsParameterSelected("Limit") {
			ret["event_route[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("NotificationReceiverId") {
			if inv.IsParameterNil("NotificationReceiverId") {
				ret["event_route[notification_receiver_id]"] = ""
			} else {
				ret["event_route[notification_receiver_id]"] = convertInt64ToString(inv.Input.NotificationReceiverId)
			}
		}
		if inv.IsParameterSelected("ParentId") {
			if inv.IsParameterNil("ParentId") {
				ret["event_route[parent_id]"] = ""
			} else {
				ret["event_route[parent_id]"] = convertInt64ToString(inv.Input.ParentId)
			}
		}
		if inv.IsParameterSelected("SubjectScope") {
			ret["event_route[subject_scope]"] = inv.Input.SubjectScope
		}
		if inv.IsParameterSelected("User") {
			ret["event_route[user]"] = convertInt64ToString(inv.Input.User)
		}
	}

	return nil
}

func (inv *ActionEventRouteIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
