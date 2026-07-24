package client

import ()

// ActionEventRouteCreate is a type for action Event_route#Create
type ActionEventRouteCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteCreate(client *Client) *ActionEventRouteCreate {
	return &ActionEventRouteCreate{
		Client: client,
	}
}

// ActionEventRouteCreateMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteCreateMetaGlobalInput) SetIncludes(value string) *ActionEventRouteCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteCreateMetaGlobalInput) SetNo(value bool) *ActionEventRouteCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteCreateMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteCreateInput is a type for action input parameters
type ActionEventRouteCreateInput struct {
	Continue               bool        "json:\"continue\""
	Enabled                bool        "json:\"enabled\""
	EventType              string      "json:\"event_type\""
	EventTypePattern       string      "json:\"event_type_pattern\""
	GroupBy                interface{} "json:\"group_by\""
	GroupIntervalSeconds   int64       "json:\"group_interval_seconds\""
	GroupWaitSeconds       int64       "json:\"group_wait_seconds\""
	GroupingEnabled        bool        "json:\"grouping_enabled\""
	Label                  string      "json:\"label\""
	NotificationReceiverId int64       "json:\"notification_receiver_id\""
	ParentId               int64       "json:\"parent_id\""
	Position               int64       "json:\"position\""
	SubjectScope           string      "json:\"subject_scope\""
	User                   int64       "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetContinue sets parameter Continue to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetContinue(value bool) *ActionEventRouteCreateInput {
	in.Continue = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Continue"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetEnabled(value bool) *ActionEventRouteCreateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetEventType sets parameter EventType to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetEventType(value string) *ActionEventRouteCreateInput {
	in.EventType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventTypeNil(false)
	in._selectedParameters["EventType"] = nil
	return in
}

// SetEventTypeNil sets parameter EventType to nil and selects it for sending
func (in *ActionEventRouteCreateInput) SetEventTypeNil(set bool) *ActionEventRouteCreateInput {
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

// SetEventTypePattern sets parameter EventTypePattern to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetEventTypePattern(value string) *ActionEventRouteCreateInput {
	in.EventTypePattern = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventTypePatternNil(false)
	in._selectedParameters["EventTypePattern"] = nil
	return in
}

// SetEventTypePatternNil sets parameter EventTypePattern to nil and selects it for sending
func (in *ActionEventRouteCreateInput) SetEventTypePatternNil(set bool) *ActionEventRouteCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["EventTypePattern"] = nil
		in.SelectParameters("EventTypePattern")
	} else {
		delete(in._nilParameters, "EventTypePattern")
	}
	return in
}

// SetGroupBy sets parameter GroupBy to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetGroupBy(value interface{}) *ActionEventRouteCreateInput {
	in.GroupBy = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["GroupBy"] = nil
	return in
}

// SetGroupIntervalSeconds sets parameter GroupIntervalSeconds to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetGroupIntervalSeconds(value int64) *ActionEventRouteCreateInput {
	in.GroupIntervalSeconds = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetGroupIntervalSecondsNil(false)
	in._selectedParameters["GroupIntervalSeconds"] = nil
	return in
}

// SetGroupIntervalSecondsNil sets parameter GroupIntervalSeconds to nil and selects it for sending
func (in *ActionEventRouteCreateInput) SetGroupIntervalSecondsNil(set bool) *ActionEventRouteCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["GroupIntervalSeconds"] = nil
		in.SelectParameters("GroupIntervalSeconds")
	} else {
		delete(in._nilParameters, "GroupIntervalSeconds")
	}
	return in
}

// SetGroupWaitSeconds sets parameter GroupWaitSeconds to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetGroupWaitSeconds(value int64) *ActionEventRouteCreateInput {
	in.GroupWaitSeconds = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetGroupWaitSecondsNil(false)
	in._selectedParameters["GroupWaitSeconds"] = nil
	return in
}

// SetGroupWaitSecondsNil sets parameter GroupWaitSeconds to nil and selects it for sending
func (in *ActionEventRouteCreateInput) SetGroupWaitSecondsNil(set bool) *ActionEventRouteCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["GroupWaitSeconds"] = nil
		in.SelectParameters("GroupWaitSeconds")
	} else {
		delete(in._nilParameters, "GroupWaitSeconds")
	}
	return in
}

// SetGroupingEnabled sets parameter GroupingEnabled to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetGroupingEnabled(value bool) *ActionEventRouteCreateInput {
	in.GroupingEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["GroupingEnabled"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetLabel(value string) *ActionEventRouteCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLabelNil(false)
	in._selectedParameters["Label"] = nil
	return in
}

// SetLabelNil sets parameter Label to nil and selects it for sending
func (in *ActionEventRouteCreateInput) SetLabelNil(set bool) *ActionEventRouteCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Label"] = nil
		in.SelectParameters("Label")
	} else {
		delete(in._nilParameters, "Label")
	}
	return in
}

// SetNotificationReceiverId sets parameter NotificationReceiverId to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetNotificationReceiverId(value int64) *ActionEventRouteCreateInput {
	in.NotificationReceiverId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationReceiverIdNil(false)
	in._selectedParameters["NotificationReceiverId"] = nil
	return in
}

// SetNotificationReceiverIdNil sets parameter NotificationReceiverId to nil and selects it for sending
func (in *ActionEventRouteCreateInput) SetNotificationReceiverIdNil(set bool) *ActionEventRouteCreateInput {
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
func (in *ActionEventRouteCreateInput) SetParentId(value int64) *ActionEventRouteCreateInput {
	in.ParentId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetParentIdNil(false)
	in._selectedParameters["ParentId"] = nil
	return in
}

// SetParentIdNil sets parameter ParentId to nil and selects it for sending
func (in *ActionEventRouteCreateInput) SetParentIdNil(set bool) *ActionEventRouteCreateInput {
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

// SetPosition sets parameter Position to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetPosition(value int64) *ActionEventRouteCreateInput {
	in.Position = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Position"] = nil
	return in
}

// SetSubjectScope sets parameter SubjectScope to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetSubjectScope(value string) *ActionEventRouteCreateInput {
	in.SubjectScope = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SubjectScope"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionEventRouteCreateInput) SetUser(value int64) *ActionEventRouteCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteCreateInput) SelectParameters(params ...string) *ActionEventRouteCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteCreateInput) UnselectParameters(params ...string) *ActionEventRouteCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteCreateRequest is a type for the entire action request
type ActionEventRouteCreateRequest struct {
	EventRoute map[string]interface{} "json:\"event_route\""
	Meta       map[string]interface{} "json:\"_meta\""
}

// ActionEventRouteCreateOutput is a type for action output parameters
type ActionEventRouteCreateOutput struct {
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
type ActionEventRouteCreateResponse struct {
	Action *ActionEventRouteCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventRoute *ActionEventRouteCreateOutput "json:\"event_route\""
	}

	// Action output without the namespace
	Output *ActionEventRouteCreateOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteCreate) Prepare() *ActionEventRouteCreateInvocation {
	return &ActionEventRouteCreateInvocation{
		Action: action,
		Path:   "/v7.0/event_routes",
	}
}

// ActionEventRouteCreateInvocation is used to configure action for invocation
type ActionEventRouteCreateInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteCreateInput
	// Global meta input parameters
	MetaInput *ActionEventRouteCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteCreateInvocation) NewInput() *ActionEventRouteCreateInput {
	inv.Input = &ActionEventRouteCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteCreateInvocation) SetInput(input *ActionEventRouteCreateInput) *ActionEventRouteCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteCreateInvocation) NewMetaInput() *ActionEventRouteCreateMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteCreateInvocation) SetMetaInput(input *ActionEventRouteCreateMetaGlobalInput) *ActionEventRouteCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteCreateInvocation) validate() error {
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
func (inv *ActionEventRouteCreateInvocation) Call() (*ActionEventRouteCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventRouteCreateInvocation) callAsBody() (*ActionEventRouteCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventRouteCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventRoute
	}
	return resp, err
}

func (inv *ActionEventRouteCreateInvocation) makeAllInputParams() *ActionEventRouteCreateRequest {
	return &ActionEventRouteCreateRequest{
		EventRoute: inv.makeInputParams(),
		Meta:       inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventRouteCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Continue") {
			ret["continue"] = inv.Input.Continue
		}
		if inv.IsParameterSelected("Enabled") {
			ret["enabled"] = inv.Input.Enabled
		}
		if inv.IsParameterSelected("EventType") {
			if inv.IsParameterNil("EventType") {
				ret["event_type"] = nil
			} else {
				ret["event_type"] = inv.Input.EventType
			}
		}
		if inv.IsParameterSelected("EventTypePattern") {
			if inv.IsParameterNil("EventTypePattern") {
				ret["event_type_pattern"] = nil
			} else {
				ret["event_type_pattern"] = inv.Input.EventTypePattern
			}
		}
		if inv.IsParameterSelected("GroupBy") {
			ret["group_by"] = inv.Input.GroupBy
		}
		if inv.IsParameterSelected("GroupIntervalSeconds") {
			if inv.IsParameterNil("GroupIntervalSeconds") {
				ret["group_interval_seconds"] = nil
			} else {
				ret["group_interval_seconds"] = inv.Input.GroupIntervalSeconds
			}
		}
		if inv.IsParameterSelected("GroupWaitSeconds") {
			if inv.IsParameterNil("GroupWaitSeconds") {
				ret["group_wait_seconds"] = nil
			} else {
				ret["group_wait_seconds"] = inv.Input.GroupWaitSeconds
			}
		}
		if inv.IsParameterSelected("GroupingEnabled") {
			ret["grouping_enabled"] = inv.Input.GroupingEnabled
		}
		if inv.IsParameterSelected("Label") {
			if inv.IsParameterNil("Label") {
				ret["label"] = nil
			} else {
				ret["label"] = inv.Input.Label
			}
		}
		if inv.IsParameterSelected("NotificationReceiverId") {
			if inv.IsParameterNil("NotificationReceiverId") {
				ret["notification_receiver_id"] = nil
			} else {
				ret["notification_receiver_id"] = inv.Input.NotificationReceiverId
			}
		}
		if inv.IsParameterSelected("ParentId") {
			if inv.IsParameterNil("ParentId") {
				ret["parent_id"] = nil
			} else {
				ret["parent_id"] = inv.Input.ParentId
			}
		}
		if inv.IsParameterSelected("Position") {
			ret["position"] = inv.Input.Position
		}
		if inv.IsParameterSelected("SubjectScope") {
			ret["subject_scope"] = inv.Input.SubjectScope
		}
		if inv.IsParameterSelected("User") {
			ret["user"] = inv.Input.User
		}
	}

	return ret
}

func (inv *ActionEventRouteCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
