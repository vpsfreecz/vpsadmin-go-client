package client

import (
	"net/url"
	"strings"
)

// ActionIncidentReportMuteSimilar is a type for action Incident_report#Mute_similar
type ActionIncidentReportMuteSimilar struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentReportMuteSimilar(client *Client) *ActionIncidentReportMuteSimilar {
	return &ActionIncidentReportMuteSimilar{
		Client: client,
	}
}

// ActionIncidentReportMuteSimilarMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentReportMuteSimilarMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIncidentReportMuteSimilarMetaGlobalInput) SetIncludes(value string) *ActionIncidentReportMuteSimilarMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentReportMuteSimilarMetaGlobalInput) SetNo(value bool) *ActionIncidentReportMuteSimilarMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentReportMuteSimilarMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportMuteSimilarMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentReportMuteSimilarMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentReportMuteSimilarMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportMuteSimilarInput is a type for action input parameters
type ActionIncidentReportMuteSimilarInput struct {
	ExpiresAt     string "json:\"expires_at\""
	MatchCodename bool   "json:\"match_codename\""
	MatchIpAddr   bool   "json:\"match_ip_addr\""
	MatchSubject  bool   "json:\"match_subject\""
	MatchVps      bool   "json:\"match_vps\""
	RouteOwner    int64  "json:\"route_owner\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetExpiresAt sets parameter ExpiresAt to value and selects it for sending
func (in *ActionIncidentReportMuteSimilarInput) SetExpiresAt(value string) *ActionIncidentReportMuteSimilarInput {
	in.ExpiresAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetExpiresAtNil(false)
	in._selectedParameters["ExpiresAt"] = nil
	return in
}

// SetExpiresAtNil sets parameter ExpiresAt to nil and selects it for sending
func (in *ActionIncidentReportMuteSimilarInput) SetExpiresAtNil(set bool) *ActionIncidentReportMuteSimilarInput {
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

// SetMatchCodename sets parameter MatchCodename to value and selects it for sending
func (in *ActionIncidentReportMuteSimilarInput) SetMatchCodename(value bool) *ActionIncidentReportMuteSimilarInput {
	in.MatchCodename = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MatchCodename"] = nil
	return in
}

// SetMatchIpAddr sets parameter MatchIpAddr to value and selects it for sending
func (in *ActionIncidentReportMuteSimilarInput) SetMatchIpAddr(value bool) *ActionIncidentReportMuteSimilarInput {
	in.MatchIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MatchIpAddr"] = nil
	return in
}

// SetMatchSubject sets parameter MatchSubject to value and selects it for sending
func (in *ActionIncidentReportMuteSimilarInput) SetMatchSubject(value bool) *ActionIncidentReportMuteSimilarInput {
	in.MatchSubject = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MatchSubject"] = nil
	return in
}

// SetMatchVps sets parameter MatchVps to value and selects it for sending
func (in *ActionIncidentReportMuteSimilarInput) SetMatchVps(value bool) *ActionIncidentReportMuteSimilarInput {
	in.MatchVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MatchVps"] = nil
	return in
}

// SetRouteOwner sets parameter RouteOwner to value and selects it for sending
func (in *ActionIncidentReportMuteSimilarInput) SetRouteOwner(value int64) *ActionIncidentReportMuteSimilarInput {
	in.RouteOwner = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetRouteOwnerNil(false)
	in._selectedParameters["RouteOwner"] = nil
	return in
}

// SetRouteOwnerNil sets parameter RouteOwner to nil and selects it for sending
func (in *ActionIncidentReportMuteSimilarInput) SetRouteOwnerNil(set bool) *ActionIncidentReportMuteSimilarInput {
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

// SelectParameters sets parameters from ActionIncidentReportMuteSimilarInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportMuteSimilarInput) SelectParameters(params ...string) *ActionIncidentReportMuteSimilarInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIncidentReportMuteSimilarInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIncidentReportMuteSimilarInput) UnselectParameters(params ...string) *ActionIncidentReportMuteSimilarInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIncidentReportMuteSimilarInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportMuteSimilarRequest is a type for the entire action request
type ActionIncidentReportMuteSimilarRequest struct {
	IncidentReport map[string]interface{} "json:\"incident_report\""
	Meta           map[string]interface{} "json:\"_meta\""
}

// ActionIncidentReportMuteSimilarOutput is a type for action output parameters
type ActionIncidentReportMuteSimilarOutput struct {
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
type ActionIncidentReportMuteSimilarResponse struct {
	Action *ActionIncidentReportMuteSimilar "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventRoute *ActionIncidentReportMuteSimilarOutput "json:\"event_route\""
	}

	// Action output without the namespace
	Output *ActionIncidentReportMuteSimilarOutput
}

// Prepare the action for invocation
func (action *ActionIncidentReportMuteSimilar) Prepare() *ActionIncidentReportMuteSimilarInvocation {
	return &ActionIncidentReportMuteSimilarInvocation{
		Action: action,
		Path:   "/v7.0/incident_reports/{incident_report_id}/mute_similar",
	}
}

// ActionIncidentReportMuteSimilarInvocation is used to configure action for invocation
type ActionIncidentReportMuteSimilarInvocation struct {
	// Pointer to the action
	Action *ActionIncidentReportMuteSimilar

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIncidentReportMuteSimilarInput
	// Global meta input parameters
	MetaInput *ActionIncidentReportMuteSimilarMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIncidentReportMuteSimilarInvocation) SetPathParamInt(param string, value int64) *ActionIncidentReportMuteSimilarInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIncidentReportMuteSimilarInvocation) SetPathParamString(param string, value string) *ActionIncidentReportMuteSimilarInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIncidentReportMuteSimilarInvocation) NewInput() *ActionIncidentReportMuteSimilarInput {
	inv.Input = &ActionIncidentReportMuteSimilarInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIncidentReportMuteSimilarInvocation) SetInput(input *ActionIncidentReportMuteSimilarInput) *ActionIncidentReportMuteSimilarInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIncidentReportMuteSimilarInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIncidentReportMuteSimilarInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentReportMuteSimilarInvocation) NewMetaInput() *ActionIncidentReportMuteSimilarMetaGlobalInput {
	inv.MetaInput = &ActionIncidentReportMuteSimilarMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentReportMuteSimilarInvocation) SetMetaInput(input *ActionIncidentReportMuteSimilarMetaGlobalInput) *ActionIncidentReportMuteSimilarInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentReportMuteSimilarInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentReportMuteSimilarInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIncidentReportMuteSimilarInvocation) validate() error {
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
func (inv *ActionIncidentReportMuteSimilarInvocation) Call() (*ActionIncidentReportMuteSimilarResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionIncidentReportMuteSimilarInvocation) callAsBody() (*ActionIncidentReportMuteSimilarResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIncidentReportMuteSimilarResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventRoute
	}
	return resp, err
}

func (inv *ActionIncidentReportMuteSimilarInvocation) makeAllInputParams() *ActionIncidentReportMuteSimilarRequest {
	return &ActionIncidentReportMuteSimilarRequest{
		IncidentReport: inv.makeInputParams(),
		Meta:           inv.makeMetaInputParams(),
	}
}

func (inv *ActionIncidentReportMuteSimilarInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ExpiresAt") {
			if inv.IsParameterNil("ExpiresAt") {
				ret["expires_at"] = nil
			} else {
				ret["expires_at"] = inv.Input.ExpiresAt
			}
		}
		if inv.IsParameterSelected("MatchCodename") {
			ret["match_codename"] = inv.Input.MatchCodename
		}
		if inv.IsParameterSelected("MatchIpAddr") {
			ret["match_ip_addr"] = inv.Input.MatchIpAddr
		}
		if inv.IsParameterSelected("MatchSubject") {
			ret["match_subject"] = inv.Input.MatchSubject
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

func (inv *ActionIncidentReportMuteSimilarInvocation) makeMetaInputParams() map[string]interface{} {
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
