package client

import (
	"net/url"
	"strings"
)

// ActionEventRouteUpdate is a type for action Event_route#Update
type ActionEventRouteUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionEventRouteUpdate(client *Client) *ActionEventRouteUpdate {
	return &ActionEventRouteUpdate{
		Client: client,
	}
}

// ActionEventRouteUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionEventRouteUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventRouteUpdateMetaGlobalInput) SetIncludes(value string) *ActionEventRouteUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventRouteUpdateMetaGlobalInput) SetNo(value bool) *ActionEventRouteUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionEventRouteUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventRouteUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteUpdateInput is a type for action input parameters
type ActionEventRouteUpdateInput struct {
	Continue               bool   "json:\"continue\""
	Enabled                bool   "json:\"enabled\""
	EventType              string "json:\"event_type\""
	EventTypePattern       string "json:\"event_type_pattern\""
	Label                  string "json:\"label\""
	NotificationReceiverId int64  "json:\"notification_receiver_id\""
	ParentId               int64  "json:\"parent_id\""
	Position               int64  "json:\"position\""
	SubjectScope           string "json:\"subject_scope\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetContinue sets parameter Continue to value and selects it for sending
func (in *ActionEventRouteUpdateInput) SetContinue(value bool) *ActionEventRouteUpdateInput {
	in.Continue = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Continue"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionEventRouteUpdateInput) SetEnabled(value bool) *ActionEventRouteUpdateInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetEventType sets parameter EventType to value and selects it for sending
func (in *ActionEventRouteUpdateInput) SetEventType(value string) *ActionEventRouteUpdateInput {
	in.EventType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventTypeNil(false)
	in._selectedParameters["EventType"] = nil
	return in
}

// SetEventTypeNil sets parameter EventType to nil and selects it for sending
func (in *ActionEventRouteUpdateInput) SetEventTypeNil(set bool) *ActionEventRouteUpdateInput {
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
func (in *ActionEventRouteUpdateInput) SetEventTypePattern(value string) *ActionEventRouteUpdateInput {
	in.EventTypePattern = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventTypePatternNil(false)
	in._selectedParameters["EventTypePattern"] = nil
	return in
}

// SetEventTypePatternNil sets parameter EventTypePattern to nil and selects it for sending
func (in *ActionEventRouteUpdateInput) SetEventTypePatternNil(set bool) *ActionEventRouteUpdateInput {
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

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionEventRouteUpdateInput) SetLabel(value string) *ActionEventRouteUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLabelNil(false)
	in._selectedParameters["Label"] = nil
	return in
}

// SetLabelNil sets parameter Label to nil and selects it for sending
func (in *ActionEventRouteUpdateInput) SetLabelNil(set bool) *ActionEventRouteUpdateInput {
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
func (in *ActionEventRouteUpdateInput) SetNotificationReceiverId(value int64) *ActionEventRouteUpdateInput {
	in.NotificationReceiverId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationReceiverIdNil(false)
	in._selectedParameters["NotificationReceiverId"] = nil
	return in
}

// SetNotificationReceiverIdNil sets parameter NotificationReceiverId to nil and selects it for sending
func (in *ActionEventRouteUpdateInput) SetNotificationReceiverIdNil(set bool) *ActionEventRouteUpdateInput {
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
func (in *ActionEventRouteUpdateInput) SetParentId(value int64) *ActionEventRouteUpdateInput {
	in.ParentId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetParentIdNil(false)
	in._selectedParameters["ParentId"] = nil
	return in
}

// SetParentIdNil sets parameter ParentId to nil and selects it for sending
func (in *ActionEventRouteUpdateInput) SetParentIdNil(set bool) *ActionEventRouteUpdateInput {
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
func (in *ActionEventRouteUpdateInput) SetPosition(value int64) *ActionEventRouteUpdateInput {
	in.Position = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Position"] = nil
	return in
}

// SetSubjectScope sets parameter SubjectScope to value and selects it for sending
func (in *ActionEventRouteUpdateInput) SetSubjectScope(value string) *ActionEventRouteUpdateInput {
	in.SubjectScope = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SubjectScope"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventRouteUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventRouteUpdateInput) SelectParameters(params ...string) *ActionEventRouteUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventRouteUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventRouteUpdateInput) UnselectParameters(params ...string) *ActionEventRouteUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventRouteUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventRouteUpdateRequest is a type for the entire action request
type ActionEventRouteUpdateRequest struct {
	EventRoute map[string]interface{} "json:\"event_route\""
	Meta       map[string]interface{} "json:\"_meta\""
}

// ActionEventRouteUpdateOutput is a type for action output parameters
type ActionEventRouteUpdateOutput struct {
	Continue               bool                  "json:\"continue\""
	CreatedAt              string                "json:\"created_at\""
	DisplayLabel           string                "json:\"display_label\""
	Enabled                bool                  "json:\"enabled\""
	EventType              string                "json:\"event_type\""
	EventTypePattern       string                "json:\"event_type_pattern\""
	ExpiresAt              string                "json:\"expires_at\""
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
type ActionEventRouteUpdateResponse struct {
	Action *ActionEventRouteUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventRoute *ActionEventRouteUpdateOutput "json:\"event_route\""
	}

	// Action output without the namespace
	Output *ActionEventRouteUpdateOutput
}

// Prepare the action for invocation
func (action *ActionEventRouteUpdate) Prepare() *ActionEventRouteUpdateInvocation {
	return &ActionEventRouteUpdateInvocation{
		Action: action,
		Path:   "/v7.0/event_routes/{event_route_id}",
	}
}

// ActionEventRouteUpdateInvocation is used to configure action for invocation
type ActionEventRouteUpdateInvocation struct {
	// Pointer to the action
	Action *ActionEventRouteUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventRouteUpdateInput
	// Global meta input parameters
	MetaInput *ActionEventRouteUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventRouteUpdateInvocation) SetPathParamInt(param string, value int64) *ActionEventRouteUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventRouteUpdateInvocation) SetPathParamString(param string, value string) *ActionEventRouteUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventRouteUpdateInvocation) NewInput() *ActionEventRouteUpdateInput {
	inv.Input = &ActionEventRouteUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventRouteUpdateInvocation) SetInput(input *ActionEventRouteUpdateInput) *ActionEventRouteUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventRouteUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventRouteUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventRouteUpdateInvocation) NewMetaInput() *ActionEventRouteUpdateMetaGlobalInput {
	inv.MetaInput = &ActionEventRouteUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventRouteUpdateInvocation) SetMetaInput(input *ActionEventRouteUpdateMetaGlobalInput) *ActionEventRouteUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventRouteUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventRouteUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventRouteUpdateInvocation) validate() error {
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
func (inv *ActionEventRouteUpdateInvocation) Call() (*ActionEventRouteUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventRouteUpdateInvocation) callAsBody() (*ActionEventRouteUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventRouteUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventRoute
	}
	return resp, err
}

func (inv *ActionEventRouteUpdateInvocation) makeAllInputParams() *ActionEventRouteUpdateRequest {
	return &ActionEventRouteUpdateRequest{
		EventRoute: inv.makeInputParams(),
		Meta:       inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventRouteUpdateInvocation) makeInputParams() map[string]interface{} {
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
	}

	return ret
}

func (inv *ActionEventRouteUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
