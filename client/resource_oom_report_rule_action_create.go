package client

import ()

// ActionOomReportRuleCreate is a type for action Oom_report_rule#Create
type ActionOomReportRuleCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportRuleCreate(client *Client) *ActionOomReportRuleCreate {
	return &ActionOomReportRuleCreate{
		Client: client,
	}
}

// ActionOomReportRuleCreateMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportRuleCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportRuleCreateMetaGlobalInput) SetIncludes(value string) *ActionOomReportRuleCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportRuleCreateMetaGlobalInput) SetNo(value bool) *ActionOomReportRuleCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportRuleCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportRuleCreateMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportRuleCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportRuleCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportRuleCreateInput is a type for action input parameters
type ActionOomReportRuleCreateInput struct {
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
func (in *ActionOomReportRuleCreateInput) SetAction(value string) *ActionOomReportRuleCreateInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}

// SetCgroupPattern sets parameter CgroupPattern to value and selects it for sending
func (in *ActionOomReportRuleCreateInput) SetCgroupPattern(value string) *ActionOomReportRuleCreateInput {
	in.CgroupPattern = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CgroupPattern"] = nil
	return in
}

// SetHitCount sets parameter HitCount to value and selects it for sending
func (in *ActionOomReportRuleCreateInput) SetHitCount(value int64) *ActionOomReportRuleCreateInput {
	in.HitCount = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HitCount"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionOomReportRuleCreateInput) SetVps(value int64) *ActionOomReportRuleCreateInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionOomReportRuleCreateInput) SetVpsNil(set bool) *ActionOomReportRuleCreateInput {
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

// SelectParameters sets parameters from ActionOomReportRuleCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportRuleCreateInput) SelectParameters(params ...string) *ActionOomReportRuleCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOomReportRuleCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOomReportRuleCreateInput) UnselectParameters(params ...string) *ActionOomReportRuleCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOomReportRuleCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportRuleCreateRequest is a type for the entire action request
type ActionOomReportRuleCreateRequest struct {
	OomReportRule map[string]interface{} `json:"oom_report_rule"`
	Meta          map[string]interface{} `json:"_meta"`
}

// ActionOomReportRuleCreateOutput is a type for action output parameters
type ActionOomReportRuleCreateOutput struct {
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
type ActionOomReportRuleCreateResponse struct {
	Action *ActionOomReportRuleCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OomReportRule *ActionOomReportRuleCreateOutput `json:"oom_report_rule"`
	}

	// Action output without the namespace
	Output *ActionOomReportRuleCreateOutput
}

// Prepare the action for invocation
func (action *ActionOomReportRuleCreate) Prepare() *ActionOomReportRuleCreateInvocation {
	return &ActionOomReportRuleCreateInvocation{
		Action: action,
		Path:   "/v7.0/oom_report_rules",
	}
}

// ActionOomReportRuleCreateInvocation is used to configure action for invocation
type ActionOomReportRuleCreateInvocation struct {
	// Pointer to the action
	Action *ActionOomReportRuleCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOomReportRuleCreateInput
	// Global meta input parameters
	MetaInput *ActionOomReportRuleCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOomReportRuleCreateInvocation) NewInput() *ActionOomReportRuleCreateInput {
	inv.Input = &ActionOomReportRuleCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOomReportRuleCreateInvocation) SetInput(input *ActionOomReportRuleCreateInput) *ActionOomReportRuleCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOomReportRuleCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOomReportRuleCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportRuleCreateInvocation) NewMetaInput() *ActionOomReportRuleCreateMetaGlobalInput {
	inv.MetaInput = &ActionOomReportRuleCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportRuleCreateInvocation) SetMetaInput(input *ActionOomReportRuleCreateMetaGlobalInput) *ActionOomReportRuleCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportRuleCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportRuleCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportRuleCreateInvocation) Call() (*ActionOomReportRuleCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOomReportRuleCreateInvocation) callAsBody() (*ActionOomReportRuleCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOomReportRuleCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OomReportRule
	}
	return resp, err
}

func (inv *ActionOomReportRuleCreateInvocation) makeAllInputParams() *ActionOomReportRuleCreateRequest {
	return &ActionOomReportRuleCreateRequest{
		OomReportRule: inv.makeInputParams(),
		Meta:          inv.makeMetaInputParams(),
	}
}

func (inv *ActionOomReportRuleCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionOomReportRuleCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
