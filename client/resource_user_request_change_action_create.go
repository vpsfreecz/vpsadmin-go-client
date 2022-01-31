package client

import ()

// ActionUserRequestChangeCreate is a type for action User_request.Change#Create
type ActionUserRequestChangeCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestChangeCreate(client *Client) *ActionUserRequestChangeCreate {
	return &ActionUserRequestChangeCreate{
		Client: client,
	}
}

// ActionUserRequestChangeCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestChangeCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestChangeCreateMetaGlobalInput) SetIncludes(value string) *ActionUserRequestChangeCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestChangeCreateMetaGlobalInput) SetNo(value bool) *ActionUserRequestChangeCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestChangeCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestChangeCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestChangeCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestChangeCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestChangeCreateInput is a type for action input parameters
type ActionUserRequestChangeCreateInput struct {
	Address      string `json:"address"`
	ChangeReason string `json:"change_reason"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserRequestChangeCreateInput) SetAddress(value string) *ActionUserRequestChangeCreateInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}

// SetChangeReason sets parameter ChangeReason to value and selects it for sending
func (in *ActionUserRequestChangeCreateInput) SetChangeReason(value string) *ActionUserRequestChangeCreateInput {
	in.ChangeReason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ChangeReason"] = nil
	return in
}

// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionUserRequestChangeCreateInput) SetEmail(value string) *ActionUserRequestChangeCreateInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}

// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserRequestChangeCreateInput) SetFullName(value string) *ActionUserRequestChangeCreateInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestChangeCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestChangeCreateInput) SelectParameters(params ...string) *ActionUserRequestChangeCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserRequestChangeCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserRequestChangeCreateInput) UnselectParameters(params ...string) *ActionUserRequestChangeCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserRequestChangeCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestChangeCreateRequest is a type for the entire action request
type ActionUserRequestChangeCreateRequest struct {
	Change map[string]interface{} `json:"change"`
	Meta   map[string]interface{} `json:"_meta"`
}

// ActionUserRequestChangeCreateOutput is a type for action output parameters
type ActionUserRequestChangeCreateOutput struct {
	Address       string                `json:"address"`
	Admin         *ActionUserShowOutput `json:"admin"`
	AdminResponse string                `json:"admin_response"`
	ApiIpAddr     string                `json:"api_ip_addr"`
	ApiIpPtr      string                `json:"api_ip_ptr"`
	ChangeReason  string                `json:"change_reason"`
	ClientIpAddr  string                `json:"client_ip_addr"`
	ClientIpPtr   string                `json:"client_ip_ptr"`
	CreatedAt     string                `json:"created_at"`
	Email         string                `json:"email"`
	FullName      string                `json:"full_name"`
	Id            int64                 `json:"id"`
	Label         string                `json:"label"`
	State         string                `json:"state"`
	UpdatedAt     string                `json:"updated_at"`
	User          *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionUserRequestChangeCreateResponse struct {
	Action *ActionUserRequestChangeCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Change *ActionUserRequestChangeCreateOutput `json:"change"`
	}

	// Action output without the namespace
	Output *ActionUserRequestChangeCreateOutput
}

// Prepare the action for invocation
func (action *ActionUserRequestChangeCreate) Prepare() *ActionUserRequestChangeCreateInvocation {
	return &ActionUserRequestChangeCreateInvocation{
		Action: action,
		Path:   "/v6.0/user_request/changes",
	}
}

// ActionUserRequestChangeCreateInvocation is used to configure action for invocation
type ActionUserRequestChangeCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestChangeCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserRequestChangeCreateInput
	// Global meta input parameters
	MetaInput *ActionUserRequestChangeCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserRequestChangeCreateInvocation) NewInput() *ActionUserRequestChangeCreateInput {
	inv.Input = &ActionUserRequestChangeCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserRequestChangeCreateInvocation) SetInput(input *ActionUserRequestChangeCreateInput) *ActionUserRequestChangeCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserRequestChangeCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserRequestChangeCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestChangeCreateInvocation) NewMetaInput() *ActionUserRequestChangeCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestChangeCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestChangeCreateInvocation) SetMetaInput(input *ActionUserRequestChangeCreateMetaGlobalInput) *ActionUserRequestChangeCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestChangeCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserRequestChangeCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestChangeCreateInvocation) Call() (*ActionUserRequestChangeCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserRequestChangeCreateInvocation) callAsBody() (*ActionUserRequestChangeCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserRequestChangeCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Change
	}
	return resp, err
}

func (inv *ActionUserRequestChangeCreateInvocation) makeAllInputParams() *ActionUserRequestChangeCreateRequest {
	return &ActionUserRequestChangeCreateRequest{
		Change: inv.makeInputParams(),
		Meta:   inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserRequestChangeCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Address") {
			ret["address"] = inv.Input.Address
		}
		if inv.IsParameterSelected("ChangeReason") {
			ret["change_reason"] = inv.Input.ChangeReason
		}
		if inv.IsParameterSelected("Email") {
			ret["email"] = inv.Input.Email
		}
		if inv.IsParameterSelected("FullName") {
			ret["full_name"] = inv.Input.FullName
		}
	}

	return ret
}

func (inv *ActionUserRequestChangeCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
