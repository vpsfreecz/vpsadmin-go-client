package client

import ()

// ActionSecurityAdvisoryIndex is a type for action Security_advisory#Index
type ActionSecurityAdvisoryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryIndex(client *Client) *ActionSecurityAdvisoryIndex {
	return &ActionSecurityAdvisoryIndex{
		Client: client,
	}
}

// ActionSecurityAdvisoryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexMetaGlobalInput) SetCount(value bool) *ActionSecurityAdvisoryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryIndexInput is a type for action input parameters
type ActionSecurityAdvisoryIndexInput struct {
	Affected    bool   "json:\"affected\""
	Cve         string "json:\"cve\""
	ExternalId  string "json:\"external_id\""
	FromId      int64  "json:\"from_id\""
	Limit       int64  "json:\"limit\""
	Node        int64  "json:\"node\""
	Order       string "json:\"order\""
	RecentSince string "json:\"recent_since\""
	Since       string "json:\"since\""
	State       string "json:\"state\""
	User        int64  "json:\"user\""
	Vps         int64  "json:\"vps\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAffected sets parameter Affected to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetAffected(value bool) *ActionSecurityAdvisoryIndexInput {
	in.Affected = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Affected"] = nil
	return in
}

// SetCve sets parameter Cve to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetCve(value string) *ActionSecurityAdvisoryIndexInput {
	in.Cve = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Cve"] = nil
	return in
}

// SetExternalId sets parameter ExternalId to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetExternalId(value string) *ActionSecurityAdvisoryIndexInput {
	in.ExternalId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ExternalId"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetFromId(value int64) *ActionSecurityAdvisoryIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetLimit(value int64) *ActionSecurityAdvisoryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetNode(value int64) *ActionSecurityAdvisoryIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetOrder(value string) *ActionSecurityAdvisoryIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetRecentSince sets parameter RecentSince to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetRecentSince(value string) *ActionSecurityAdvisoryIndexInput {
	in.RecentSince = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RecentSince"] = nil
	return in
}

// SetSince sets parameter Since to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetSince(value string) *ActionSecurityAdvisoryIndexInput {
	in.Since = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Since"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetState(value string) *ActionSecurityAdvisoryIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetUser(value int64) *ActionSecurityAdvisoryIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionSecurityAdvisoryIndexInput) SetVps(value int64) *ActionSecurityAdvisoryIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryIndexInput) SelectParameters(params ...string) *ActionSecurityAdvisoryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryIndexInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryIndexOutput is a type for action output parameters
type ActionSecurityAdvisoryIndexOutput struct {
	Affected          bool                  "json:\"affected\""
	AffectedNodeCount int64                 "json:\"affected_node_count\""
	AffectedUserCount int64                 "json:\"affected_user_count\""
	AffectedVpsCount  int64                 "json:\"affected_vps_count\""
	ContentRevision   int64                 "json:\"content_revision\""
	CreatedAt         string                "json:\"created_at\""
	CreatedBy         *ActionUserShowOutput "json:\"created_by\""
	CsDescription     string                "json:\"cs_description\""
	CsResponse        string                "json:\"cs_response\""
	CsSummary         string                "json:\"cs_summary\""
	EnDescription     string                "json:\"en_description\""
	EnResponse        string                "json:\"en_response\""
	EnSummary         string                "json:\"en_summary\""
	ExternalId        string                "json:\"external_id\""
	Id                int64                 "json:\"id\""
	Name              string                "json:\"name\""
	PublishedAt       string                "json:\"published_at\""
	PublishedBy       *ActionUserShowOutput "json:\"published_by\""
	RetractedAt       string                "json:\"retracted_at\""
	State             string                "json:\"state\""
	UpdatedAt         string                "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryIndexResponse struct {
	Action *ActionSecurityAdvisoryIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisories []*ActionSecurityAdvisoryIndexOutput "json:\"security_advisories\""
	}

	// Action output without the namespace
	Output []*ActionSecurityAdvisoryIndexOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryIndex) Prepare() *ActionSecurityAdvisoryIndexInvocation {
	return &ActionSecurityAdvisoryIndexInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories",
	}
}

// ActionSecurityAdvisoryIndexInvocation is used to configure action for invocation
type ActionSecurityAdvisoryIndexInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryIndexInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryIndexInvocation) NewInput() *ActionSecurityAdvisoryIndexInput {
	inv.Input = &ActionSecurityAdvisoryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryIndexInvocation) SetInput(input *ActionSecurityAdvisoryIndexInput) *ActionSecurityAdvisoryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryIndexInvocation) NewMetaInput() *ActionSecurityAdvisoryIndexMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryIndexInvocation) SetMetaInput(input *ActionSecurityAdvisoryIndexMetaGlobalInput) *ActionSecurityAdvisoryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("RecentSince") {
			if !inv.IsParameterNil("RecentSince") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.RecentSince)
				if !ok {
					verr.Add("recent_since", "not a valid datetime")
				} else {
					inv.Input.RecentSince = normalized
				}
			}
		}
		if inv.IsParameterSelected("Since") {
			if !inv.IsParameterNil("Since") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.Since)
				if !ok {
					verr.Add("since", "not a valid datetime")
				} else {
					inv.Input.Since = normalized
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
		if inv.IsParameterSelected("Vps") {
			if !inv.IsParameterNil("Vps") {
				if inv.Input.Vps < 0 {
					verr.Add("vps", "not a valid resource id")
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
func (inv *ActionSecurityAdvisoryIndexInvocation) Call() (*ActionSecurityAdvisoryIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionSecurityAdvisoryIndexInvocation) callAsQuery() (*ActionSecurityAdvisoryIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSecurityAdvisoryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisories
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Affected") {
			ret["security_advisory[affected]"] = convertBoolToString(inv.Input.Affected)
		}
		if inv.IsParameterSelected("Cve") {
			ret["security_advisory[cve]"] = inv.Input.Cve
		}
		if inv.IsParameterSelected("ExternalId") {
			ret["security_advisory[external_id]"] = inv.Input.ExternalId
		}
		if inv.IsParameterSelected("FromId") {
			ret["security_advisory[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["security_advisory[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["security_advisory[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Order") {
			ret["security_advisory[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("RecentSince") {
			ret["security_advisory[recent_since]"] = inv.Input.RecentSince
		}
		if inv.IsParameterSelected("Since") {
			ret["security_advisory[since]"] = inv.Input.Since
		}
		if inv.IsParameterSelected("State") {
			ret["security_advisory[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("User") {
			ret["security_advisory[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["security_advisory[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionSecurityAdvisoryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
