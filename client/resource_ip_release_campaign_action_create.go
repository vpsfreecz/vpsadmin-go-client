package client

import ()

// ActionIpReleaseCampaignCreate is a type for action Ip_release_campaign#Create
type ActionIpReleaseCampaignCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignCreate(client *Client) *ActionIpReleaseCampaignCreate {
	return &ActionIpReleaseCampaignCreate{
		Client: client,
	}
}

// ActionIpReleaseCampaignCreateMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignCreateMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignCreateMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignCreateMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignCreateInput is a type for action input parameters
type ActionIpReleaseCampaignCreateInput struct {
	AllowKeep bool   "json:\"allow_keep\""
	Deadline  string "json:\"deadline\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAllowKeep sets parameter AllowKeep to value and selects it for sending
func (in *ActionIpReleaseCampaignCreateInput) SetAllowKeep(value bool) *ActionIpReleaseCampaignCreateInput {
	in.AllowKeep = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AllowKeep"] = nil
	return in
}

// SetDeadline sets parameter Deadline to value and selects it for sending
func (in *ActionIpReleaseCampaignCreateInput) SetDeadline(value string) *ActionIpReleaseCampaignCreateInput {
	in.Deadline = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Deadline"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignCreateInput) SelectParameters(params ...string) *ActionIpReleaseCampaignCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseCampaignCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignCreateInput) UnselectParameters(params ...string) *ActionIpReleaseCampaignCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseCampaignCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignCreateRequest is a type for the entire action request
type ActionIpReleaseCampaignCreateRequest struct {
	IpReleaseCampaign map[string]interface{} "json:\"ip_release_campaign\""
	Meta              map[string]interface{} "json:\"_meta\""
}

// ActionIpReleaseCampaignCreateOutput is a type for action output parameters
type ActionIpReleaseCampaignCreateOutput struct {
	AllowKeep bool                  "json:\"allow_keep\""
	ClosedAt  string                "json:\"closed_at\""
	ClosedBy  *ActionUserShowOutput "json:\"closed_by\""
	CreatedAt string                "json:\"created_at\""
	CreatedBy *ActionUserShowOutput "json:\"created_by\""
	Deadline  string                "json:\"deadline\""
	Id        int64                 "json:\"id\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignCreateResponse struct {
	Action *ActionIpReleaseCampaignCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseCampaign *ActionIpReleaseCampaignCreateOutput "json:\"ip_release_campaign\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseCampaignCreateOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignCreate) Prepare() *ActionIpReleaseCampaignCreateInvocation {
	return &ActionIpReleaseCampaignCreateInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns",
	}
}

// ActionIpReleaseCampaignCreateInvocation is used to configure action for invocation
type ActionIpReleaseCampaignCreateInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseCampaignCreateInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseCampaignCreateInvocation) NewInput() *ActionIpReleaseCampaignCreateInput {
	inv.Input = &ActionIpReleaseCampaignCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseCampaignCreateInvocation) SetInput(input *ActionIpReleaseCampaignCreateInput) *ActionIpReleaseCampaignCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseCampaignCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignCreateInvocation) NewMetaInput() *ActionIpReleaseCampaignCreateMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignCreateInvocation) SetMetaInput(input *ActionIpReleaseCampaignCreateMetaGlobalInput) *ActionIpReleaseCampaignCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignCreateInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("Deadline") {
			if !inv.IsParameterNil("Deadline") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.Deadline)
				if !ok {
					verr.Add("deadline", t(inv.Action.Client, "validation.invalid_datetime", nil))
				} else {
					inv.Input.Deadline = normalized
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
func (inv *ActionIpReleaseCampaignCreateInvocation) Call() (*ActionIpReleaseCampaignCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionIpReleaseCampaignCreateInvocation) callAsBody() (*ActionIpReleaseCampaignCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIpReleaseCampaignCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpReleaseCampaign
	}
	return resp, err
}

func (inv *ActionIpReleaseCampaignCreateInvocation) makeAllInputParams() *ActionIpReleaseCampaignCreateRequest {
	return &ActionIpReleaseCampaignCreateRequest{
		IpReleaseCampaign: inv.makeInputParams(),
		Meta:              inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpReleaseCampaignCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AllowKeep") {
			ret["allow_keep"] = inv.Input.AllowKeep
		}
		if inv.IsParameterSelected("Deadline") {
			ret["deadline"] = inv.Input.Deadline
		}
	}

	return ret
}

func (inv *ActionIpReleaseCampaignCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
