package client

import ()

// ActionOomReportRuleIndex is a type for action Oom_report_rule#Index
type ActionOomReportRuleIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportRuleIndex(client *Client) *ActionOomReportRuleIndex {
	return &ActionOomReportRuleIndex{
		Client: client,
	}
}

// ActionOomReportRuleIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportRuleIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOomReportRuleIndexMetaGlobalInput) SetCount(value bool) *ActionOomReportRuleIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportRuleIndexMetaGlobalInput) SetIncludes(value string) *ActionOomReportRuleIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportRuleIndexMetaGlobalInput) SetNo(value bool) *ActionOomReportRuleIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportRuleIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportRuleIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportRuleIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportRuleIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportRuleIndexInput is a type for action input parameters
type ActionOomReportRuleIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	Vps    int64 `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionOomReportRuleIndexInput) SetFromId(value int64) *ActionOomReportRuleIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOomReportRuleIndexInput) SetLimit(value int64) *ActionOomReportRuleIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionOomReportRuleIndexInput) SetVps(value int64) *ActionOomReportRuleIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionOomReportRuleIndexInput) SetVpsNil(set bool) *ActionOomReportRuleIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Vps"] = nil
		in.SelectParameters("Vps")
	} else {
		delete(in._nilParameters, "Vps")
	}
	return in
}

// SelectParameters sets parameters from ActionOomReportRuleIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportRuleIndexInput) SelectParameters(params ...string) *ActionOomReportRuleIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOomReportRuleIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOomReportRuleIndexInput) UnselectParameters(params ...string) *ActionOomReportRuleIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOomReportRuleIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportRuleIndexOutput is a type for action output parameters
type ActionOomReportRuleIndexOutput struct {
	Action        string               `json:"action"`
	CgroupPattern string               `json:"cgroup_pattern"`
	CreatedAt     string               `json:"created_at"`
	HitCount      int64                `json:"hit_count"`
	Id            int64                `json:"id"`
	Label         string               `json:"label"`
	UpdatedAt     string               `json:"updated_at"`
	Vps           *ActionVpsShowOutput `json:"vps"`
}

// Type for action response, including envelope
type ActionOomReportRuleIndexResponse struct {
	Action *ActionOomReportRuleIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OomReportRules []*ActionOomReportRuleIndexOutput `json:"oom_report_rules"`
	}

	// Action output without the namespace
	Output []*ActionOomReportRuleIndexOutput
}

// Prepare the action for invocation
func (action *ActionOomReportRuleIndex) Prepare() *ActionOomReportRuleIndexInvocation {
	return &ActionOomReportRuleIndexInvocation{
		Action: action,
		Path:   "/v7.0/oom_report_rules",
	}
}

// ActionOomReportRuleIndexInvocation is used to configure action for invocation
type ActionOomReportRuleIndexInvocation struct {
	// Pointer to the action
	Action *ActionOomReportRuleIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOomReportRuleIndexInput
	// Global meta input parameters
	MetaInput *ActionOomReportRuleIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOomReportRuleIndexInvocation) NewInput() *ActionOomReportRuleIndexInput {
	inv.Input = &ActionOomReportRuleIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOomReportRuleIndexInvocation) SetInput(input *ActionOomReportRuleIndexInput) *ActionOomReportRuleIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOomReportRuleIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOomReportRuleIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportRuleIndexInvocation) NewMetaInput() *ActionOomReportRuleIndexMetaGlobalInput {
	inv.MetaInput = &ActionOomReportRuleIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportRuleIndexInvocation) SetMetaInput(input *ActionOomReportRuleIndexMetaGlobalInput) *ActionOomReportRuleIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportRuleIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportRuleIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportRuleIndexInvocation) Call() (*ActionOomReportRuleIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOomReportRuleIndexInvocation) callAsQuery() (*ActionOomReportRuleIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOomReportRuleIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OomReportRules
	}
	return resp, err
}

func (inv *ActionOomReportRuleIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["oom_report_rule[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["oom_report_rule[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Vps") {
			ret["oom_report_rule[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionOomReportRuleIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
