package client

import (
	"net/url"
	"strings"
)

// ActionOomReportMuteSimilar is a type for action Oom_report#Mute_similar
type ActionOomReportMuteSimilar struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportMuteSimilar(client *Client) *ActionOomReportMuteSimilar {
	return &ActionOomReportMuteSimilar{
		Client: client,
	}
}

// ActionOomReportMuteSimilarMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportMuteSimilarMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportMuteSimilarMetaGlobalInput) SetIncludes(value string) *ActionOomReportMuteSimilarMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportMuteSimilarMetaGlobalInput) SetNo(value bool) *ActionOomReportMuteSimilarMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportMuteSimilarMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportMuteSimilarMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportMuteSimilarMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportMuteSimilarMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportMuteSimilarInput is a type for action input parameters
type ActionOomReportMuteSimilarInput struct {
	ExpiresAt          string "json:\"expires_at\""
	MatchCgroup        bool   "json:\"match_cgroup\""
	MatchInvokedByName bool   "json:\"match_invoked_by_name\""
	MatchKilledName    bool   "json:\"match_killed_name\""
	MatchVps           bool   "json:\"match_vps\""
	RouteOwner         int64  "json:\"route_owner\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetExpiresAt sets parameter ExpiresAt to value and selects it for sending
func (in *ActionOomReportMuteSimilarInput) SetExpiresAt(value string) *ActionOomReportMuteSimilarInput {
	in.ExpiresAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetExpiresAtNil(false)
	in._selectedParameters["ExpiresAt"] = nil
	return in
}

// SetExpiresAtNil sets parameter ExpiresAt to nil and selects it for sending
func (in *ActionOomReportMuteSimilarInput) SetExpiresAtNil(set bool) *ActionOomReportMuteSimilarInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["ExpiresAt"] = nil
		in.SelectParameters("ExpiresAt")
	} else {
		delete(in._nilParameters, "ExpiresAt")
	}
	return in
}

// SetMatchCgroup sets parameter MatchCgroup to value and selects it for sending
func (in *ActionOomReportMuteSimilarInput) SetMatchCgroup(value bool) *ActionOomReportMuteSimilarInput {
	in.MatchCgroup = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MatchCgroup"] = nil
	return in
}

// SetMatchInvokedByName sets parameter MatchInvokedByName to value and selects it for sending
func (in *ActionOomReportMuteSimilarInput) SetMatchInvokedByName(value bool) *ActionOomReportMuteSimilarInput {
	in.MatchInvokedByName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MatchInvokedByName"] = nil
	return in
}

// SetMatchKilledName sets parameter MatchKilledName to value and selects it for sending
func (in *ActionOomReportMuteSimilarInput) SetMatchKilledName(value bool) *ActionOomReportMuteSimilarInput {
	in.MatchKilledName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MatchKilledName"] = nil
	return in
}

// SetMatchVps sets parameter MatchVps to value and selects it for sending
func (in *ActionOomReportMuteSimilarInput) SetMatchVps(value bool) *ActionOomReportMuteSimilarInput {
	in.MatchVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MatchVps"] = nil
	return in
}

// SetRouteOwner sets parameter RouteOwner to value and selects it for sending
func (in *ActionOomReportMuteSimilarInput) SetRouteOwner(value int64) *ActionOomReportMuteSimilarInput {
	in.RouteOwner = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetRouteOwnerNil(false)
	in._selectedParameters["RouteOwner"] = nil
	return in
}

// SetRouteOwnerNil sets parameter RouteOwner to nil and selects it for sending
func (in *ActionOomReportMuteSimilarInput) SetRouteOwnerNil(set bool) *ActionOomReportMuteSimilarInput {
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

// SelectParameters sets parameters from ActionOomReportMuteSimilarInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportMuteSimilarInput) SelectParameters(params ...string) *ActionOomReportMuteSimilarInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOomReportMuteSimilarInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOomReportMuteSimilarInput) UnselectParameters(params ...string) *ActionOomReportMuteSimilarInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOomReportMuteSimilarInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportMuteSimilarRequest is a type for the entire action request
type ActionOomReportMuteSimilarRequest struct {
	OomReport map[string]interface{} "json:\"oom_report\""
	Meta      map[string]interface{} "json:\"_meta\""
}

// ActionOomReportMuteSimilarOutput is a type for action output parameters
type ActionOomReportMuteSimilarOutput struct {
	Continue               bool   "json:\"continue\""
	CreatedAt              string "json:\"created_at\""
	DisplayLabel           string "json:\"display_label\""
	Enabled                bool   "json:\"enabled\""
	EventType              string "json:\"event_type\""
	ExpiresAt              string "json:\"expires_at\""
	GroupingEnabled        bool   "json:\"grouping_enabled\""
	Id                     int64  "json:\"id\""
	Label                  string "json:\"label\""
	MatcherCount           int64  "json:\"matcher_count\""
	MatcherSummary         string "json:\"matcher_summary\""
	NotificationReceiverId int64  "json:\"notification_receiver_id\""
	Position               int64  "json:\"position\""
	SingleUse              bool   "json:\"single_use\""
	SpentAt                string "json:\"spent_at\""
	SubjectScope           string "json:\"subject_scope\""
	UpdatedAt              string "json:\"updated_at\""
	UserId                 int64  "json:\"user_id\""
}

// Type for action response, including envelope
type ActionOomReportMuteSimilarResponse struct {
	Action *ActionOomReportMuteSimilar "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventRoute *ActionOomReportMuteSimilarOutput "json:\"event_route\""
	}

	// Action output without the namespace
	Output *ActionOomReportMuteSimilarOutput
}

// Prepare the action for invocation
func (action *ActionOomReportMuteSimilar) Prepare() *ActionOomReportMuteSimilarInvocation {
	return &ActionOomReportMuteSimilarInvocation{
		Action: action,
		Path:   "/v7.0/oom_reports/{oom_report_id}/mute_similar",
	}
}

// ActionOomReportMuteSimilarInvocation is used to configure action for invocation
type ActionOomReportMuteSimilarInvocation struct {
	// Pointer to the action
	Action *ActionOomReportMuteSimilar

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOomReportMuteSimilarInput
	// Global meta input parameters
	MetaInput *ActionOomReportMuteSimilarMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportMuteSimilarInvocation) SetPathParamInt(param string, value int64) *ActionOomReportMuteSimilarInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportMuteSimilarInvocation) SetPathParamString(param string, value string) *ActionOomReportMuteSimilarInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOomReportMuteSimilarInvocation) NewInput() *ActionOomReportMuteSimilarInput {
	inv.Input = &ActionOomReportMuteSimilarInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOomReportMuteSimilarInvocation) SetInput(input *ActionOomReportMuteSimilarInput) *ActionOomReportMuteSimilarInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOomReportMuteSimilarInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOomReportMuteSimilarInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportMuteSimilarInvocation) NewMetaInput() *ActionOomReportMuteSimilarMetaGlobalInput {
	inv.MetaInput = &ActionOomReportMuteSimilarMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportMuteSimilarInvocation) SetMetaInput(input *ActionOomReportMuteSimilarMetaGlobalInput) *ActionOomReportMuteSimilarInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportMuteSimilarInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportMuteSimilarInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionOomReportMuteSimilarInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("ExpiresAt") {
			if !inv.IsParameterNil("ExpiresAt") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.ExpiresAt)
				if !ok {
					verr.Add("expires_at", "not a valid datetime")
				} else {
					inv.Input.ExpiresAt = normalized
				}
			}
		}
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
func (inv *ActionOomReportMuteSimilarInvocation) Call() (*ActionOomReportMuteSimilarResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionOomReportMuteSimilarInvocation) callAsBody() (*ActionOomReportMuteSimilarResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOomReportMuteSimilarResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventRoute
	}
	return resp, err
}

func (inv *ActionOomReportMuteSimilarInvocation) makeAllInputParams() *ActionOomReportMuteSimilarRequest {
	return &ActionOomReportMuteSimilarRequest{
		OomReport: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionOomReportMuteSimilarInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ExpiresAt") {
			if inv.IsParameterNil("ExpiresAt") {
				ret["expires_at"] = nil
			} else {
				ret["expires_at"] = inv.Input.ExpiresAt
			}
		}
		if inv.IsParameterSelected("MatchCgroup") {
			ret["match_cgroup"] = inv.Input.MatchCgroup
		}
		if inv.IsParameterSelected("MatchInvokedByName") {
			ret["match_invoked_by_name"] = inv.Input.MatchInvokedByName
		}
		if inv.IsParameterSelected("MatchKilledName") {
			ret["match_killed_name"] = inv.Input.MatchKilledName
		}
		if inv.IsParameterSelected("MatchVps") {
			ret["match_vps"] = inv.Input.MatchVps
		}
		if inv.IsParameterSelected("RouteOwner") {
			if inv.IsParameterNil("RouteOwner") {
				ret["route_owner"] = nil
			} else {
				ret["route_owner"] = inv.Input.RouteOwner
			}
		}
	}

	return ret
}

func (inv *ActionOomReportMuteSimilarInvocation) makeMetaInputParams() map[string]interface{} {
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
