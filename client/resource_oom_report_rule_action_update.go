package client

import (
	"strings"
)

// ActionOomReportRuleUpdate is a type for action Oom_report_rule#Update
type ActionOomReportRuleUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportRuleUpdate(client *Client) *ActionOomReportRuleUpdate {
	return &ActionOomReportRuleUpdate{
		Client: client,
	}
}

// ActionOomReportRuleUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportRuleUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportRuleUpdateMetaGlobalInput) SetIncludes(value string) *ActionOomReportRuleUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportRuleUpdateMetaGlobalInput) SetNo(value bool) *ActionOomReportRuleUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportRuleUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportRuleUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportRuleUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportRuleUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportRuleUpdateInput is a type for action input parameters
type ActionOomReportRuleUpdateInput struct {
	Action        string `json:"action"`
	CgroupPattern string `json:"cgroup_pattern"`
	HitCount      int64  `json:"hit_count"`
	Vps           int64  `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAction sets parameter Action to value and selects it for sending
func (in *ActionOomReportRuleUpdateInput) SetAction(value string) *ActionOomReportRuleUpdateInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}

// SetCgroupPattern sets parameter CgroupPattern to value and selects it for sending
func (in *ActionOomReportRuleUpdateInput) SetCgroupPattern(value string) *ActionOomReportRuleUpdateInput {
	in.CgroupPattern = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CgroupPattern"] = nil
	return in
}

// SetHitCount sets parameter HitCount to value and selects it for sending
func (in *ActionOomReportRuleUpdateInput) SetHitCount(value int64) *ActionOomReportRuleUpdateInput {
	in.HitCount = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HitCount"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionOomReportRuleUpdateInput) SetVps(value int64) *ActionOomReportRuleUpdateInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionOomReportRuleUpdateInput) SetVpsNil(set bool) *ActionOomReportRuleUpdateInput {
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

// SelectParameters sets parameters from ActionOomReportRuleUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportRuleUpdateInput) SelectParameters(params ...string) *ActionOomReportRuleUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOomReportRuleUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOomReportRuleUpdateInput) UnselectParameters(params ...string) *ActionOomReportRuleUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOomReportRuleUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportRuleUpdateRequest is a type for the entire action request
type ActionOomReportRuleUpdateRequest struct {
	OomReportRule map[string]interface{} `json:"oom_report_rule"`
	Meta          map[string]interface{} `json:"_meta"`
}

// ActionOomReportRuleUpdateOutput is a type for action output parameters
type ActionOomReportRuleUpdateOutput struct {
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
type ActionOomReportRuleUpdateResponse struct {
	Action *ActionOomReportRuleUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OomReportRule *ActionOomReportRuleUpdateOutput `json:"oom_report_rule"`
	}

	// Action output without the namespace
	Output *ActionOomReportRuleUpdateOutput
}

// Prepare the action for invocation
func (action *ActionOomReportRuleUpdate) Prepare() *ActionOomReportRuleUpdateInvocation {
	return &ActionOomReportRuleUpdateInvocation{
		Action: action,
		Path:   "/v7.0/oom_report_rules/{oom_report_rule_id}",
	}
}

// ActionOomReportRuleUpdateInvocation is used to configure action for invocation
type ActionOomReportRuleUpdateInvocation struct {
	// Pointer to the action
	Action *ActionOomReportRuleUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOomReportRuleUpdateInput
	// Global meta input parameters
	MetaInput *ActionOomReportRuleUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportRuleUpdateInvocation) SetPathParamInt(param string, value int64) *ActionOomReportRuleUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportRuleUpdateInvocation) SetPathParamString(param string, value string) *ActionOomReportRuleUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOomReportRuleUpdateInvocation) NewInput() *ActionOomReportRuleUpdateInput {
	inv.Input = &ActionOomReportRuleUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOomReportRuleUpdateInvocation) SetInput(input *ActionOomReportRuleUpdateInput) *ActionOomReportRuleUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOomReportRuleUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOomReportRuleUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportRuleUpdateInvocation) NewMetaInput() *ActionOomReportRuleUpdateMetaGlobalInput {
	inv.MetaInput = &ActionOomReportRuleUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportRuleUpdateInvocation) SetMetaInput(input *ActionOomReportRuleUpdateMetaGlobalInput) *ActionOomReportRuleUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportRuleUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportRuleUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportRuleUpdateInvocation) Call() (*ActionOomReportRuleUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOomReportRuleUpdateInvocation) callAsBody() (*ActionOomReportRuleUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOomReportRuleUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OomReportRule
	}
	return resp, err
}

func (inv *ActionOomReportRuleUpdateInvocation) makeAllInputParams() *ActionOomReportRuleUpdateRequest {
	return &ActionOomReportRuleUpdateRequest{
		OomReportRule: inv.makeInputParams(),
		Meta:          inv.makeMetaInputParams(),
	}
}

func (inv *ActionOomReportRuleUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Action") {
			ret["action"] = inv.Input.Action
		}
		if inv.IsParameterSelected("CgroupPattern") {
			ret["cgroup_pattern"] = inv.Input.CgroupPattern
		}
		if inv.IsParameterSelected("HitCount") {
			ret["hit_count"] = inv.Input.HitCount
		}
		if inv.IsParameterSelected("Vps") {
			if inv.IsParameterNil("Vps") {
				ret["vps"] = nil
			} else {
				ret["vps"] = inv.Input.Vps
			}
		}
	}

	return ret
}

func (inv *ActionOomReportRuleUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
