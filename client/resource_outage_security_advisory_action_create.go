package client

import ()

// ActionOutageSecurityAdvisoryCreate is a type for action Outage_security_advisory#Create
type ActionOutageSecurityAdvisoryCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageSecurityAdvisoryCreate(client *Client) *ActionOutageSecurityAdvisoryCreate {
	return &ActionOutageSecurityAdvisoryCreate{
		Client: client,
	}
}

// ActionOutageSecurityAdvisoryCreateMetaGlobalInput is a type for action global meta input parameters
type ActionOutageSecurityAdvisoryCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryCreateMetaGlobalInput) SetIncludes(value string) *ActionOutageSecurityAdvisoryCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryCreateMetaGlobalInput) SetNo(value bool) *ActionOutageSecurityAdvisoryCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageSecurityAdvisoryCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageSecurityAdvisoryCreateMetaGlobalInput) SelectParameters(params ...string) *ActionOutageSecurityAdvisoryCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageSecurityAdvisoryCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageSecurityAdvisoryCreateInput is a type for action input parameters
type ActionOutageSecurityAdvisoryCreateInput struct {
	Outage           int64 "json:\"outage\""
	SecurityAdvisory int64 "json:\"security_advisory\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetOutage sets parameter Outage to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryCreateInput) SetOutage(value int64) *ActionOutageSecurityAdvisoryCreateInput {
	in.Outage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Outage"] = nil
	return in
}

// SetSecurityAdvisory sets parameter SecurityAdvisory to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryCreateInput) SetSecurityAdvisory(value int64) *ActionOutageSecurityAdvisoryCreateInput {
	in.SecurityAdvisory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SecurityAdvisory"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageSecurityAdvisoryCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageSecurityAdvisoryCreateInput) SelectParameters(params ...string) *ActionOutageSecurityAdvisoryCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOutageSecurityAdvisoryCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOutageSecurityAdvisoryCreateInput) UnselectParameters(params ...string) *ActionOutageSecurityAdvisoryCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOutageSecurityAdvisoryCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageSecurityAdvisoryCreateRequest is a type for the entire action request
type ActionOutageSecurityAdvisoryCreateRequest struct {
	OutageSecurityAdvisory map[string]interface{} "json:\"outage_security_advisory\""
	Meta                   map[string]interface{} "json:\"_meta\""
}

// ActionOutageSecurityAdvisoryCreateOutput is a type for action output parameters
type ActionOutageSecurityAdvisoryCreateOutput struct {
	Id                 int64                             "json:\"id\""
	Outage             *ActionOutageShowOutput           "json:\"outage\""
	OutageId           int64                             "json:\"outage_id\""
	SecurityAdvisory   *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	SecurityAdvisoryId int64                             "json:\"security_advisory_id\""
}

// Type for action response, including envelope
type ActionOutageSecurityAdvisoryCreateResponse struct {
	Action *ActionOutageSecurityAdvisoryCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageSecurityAdvisory *ActionOutageSecurityAdvisoryCreateOutput "json:\"outage_security_advisory\""
	}

	// Action output without the namespace
	Output *ActionOutageSecurityAdvisoryCreateOutput
}

// Prepare the action for invocation
func (action *ActionOutageSecurityAdvisoryCreate) Prepare() *ActionOutageSecurityAdvisoryCreateInvocation {
	return &ActionOutageSecurityAdvisoryCreateInvocation{
		Action: action,
		Path:   "/v7.0/outage_security_advisories",
	}
}

// ActionOutageSecurityAdvisoryCreateInvocation is used to configure action for invocation
type ActionOutageSecurityAdvisoryCreateInvocation struct {
	// Pointer to the action
	Action *ActionOutageSecurityAdvisoryCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageSecurityAdvisoryCreateInput
	// Global meta input parameters
	MetaInput *ActionOutageSecurityAdvisoryCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageSecurityAdvisoryCreateInvocation) NewInput() *ActionOutageSecurityAdvisoryCreateInput {
	inv.Input = &ActionOutageSecurityAdvisoryCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageSecurityAdvisoryCreateInvocation) SetInput(input *ActionOutageSecurityAdvisoryCreateInput) *ActionOutageSecurityAdvisoryCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageSecurityAdvisoryCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOutageSecurityAdvisoryCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageSecurityAdvisoryCreateInvocation) NewMetaInput() *ActionOutageSecurityAdvisoryCreateMetaGlobalInput {
	inv.MetaInput = &ActionOutageSecurityAdvisoryCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageSecurityAdvisoryCreateInvocation) SetMetaInput(input *ActionOutageSecurityAdvisoryCreateMetaGlobalInput) *ActionOutageSecurityAdvisoryCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageSecurityAdvisoryCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageSecurityAdvisoryCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionOutageSecurityAdvisoryCreateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Outage") {
			if !inv.IsParameterNil("Outage") {
				if inv.Input.Outage < 0 {
					verr.Add("outage", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			if !inv.IsParameterNil("SecurityAdvisory") {
				if inv.Input.SecurityAdvisory < 0 {
					verr.Add("security_advisory", "not a valid resource id")
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
func (inv *ActionOutageSecurityAdvisoryCreateInvocation) Call() (*ActionOutageSecurityAdvisoryCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionOutageSecurityAdvisoryCreateInvocation) callAsBody() (*ActionOutageSecurityAdvisoryCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageSecurityAdvisoryCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageSecurityAdvisory
	}
	return resp, err
}

func (inv *ActionOutageSecurityAdvisoryCreateInvocation) makeAllInputParams() *ActionOutageSecurityAdvisoryCreateRequest {
	return &ActionOutageSecurityAdvisoryCreateRequest{
		OutageSecurityAdvisory: inv.makeInputParams(),
		Meta:                   inv.makeMetaInputParams(),
	}
}

func (inv *ActionOutageSecurityAdvisoryCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Outage") {
			ret["outage"] = inv.Input.Outage
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			ret["security_advisory"] = inv.Input.SecurityAdvisory
		}
	}

	return ret
}

func (inv *ActionOutageSecurityAdvisoryCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
