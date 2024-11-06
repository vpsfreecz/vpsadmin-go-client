package client

import ()

// ActionDnsTsigKeyCreate is a type for action Dns_tsig_key#Create
type ActionDnsTsigKeyCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsTsigKeyCreate(client *Client) *ActionDnsTsigKeyCreate {
	return &ActionDnsTsigKeyCreate{
		Client: client,
	}
}

// ActionDnsTsigKeyCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsTsigKeyCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsTsigKeyCreateMetaGlobalInput) SetIncludes(value string) *ActionDnsTsigKeyCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsTsigKeyCreateMetaGlobalInput) SetNo(value bool) *ActionDnsTsigKeyCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsTsigKeyCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsTsigKeyCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsTsigKeyCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsTsigKeyCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsTsigKeyCreateInput is a type for action input parameters
type ActionDnsTsigKeyCreateInput struct {
	Algorithm string `json:"algorithm"`
	Name      string `json:"name"`
	User      int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAlgorithm sets parameter Algorithm to value and selects it for sending
func (in *ActionDnsTsigKeyCreateInput) SetAlgorithm(value string) *ActionDnsTsigKeyCreateInput {
	in.Algorithm = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Algorithm"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionDnsTsigKeyCreateInput) SetName(value string) *ActionDnsTsigKeyCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionDnsTsigKeyCreateInput) SetUser(value int64) *ActionDnsTsigKeyCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionDnsTsigKeyCreateInput) SetUserNil(set bool) *ActionDnsTsigKeyCreateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionDnsTsigKeyCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsTsigKeyCreateInput) SelectParameters(params ...string) *ActionDnsTsigKeyCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsTsigKeyCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsTsigKeyCreateInput) UnselectParameters(params ...string) *ActionDnsTsigKeyCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsTsigKeyCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsTsigKeyCreateRequest is a type for the entire action request
type ActionDnsTsigKeyCreateRequest struct {
	DnsTsigKey map[string]interface{} `json:"dns_tsig_key"`
	Meta       map[string]interface{} `json:"_meta"`
}

// ActionDnsTsigKeyCreateOutput is a type for action output parameters
type ActionDnsTsigKeyCreateOutput struct {
	Algorithm string                `json:"algorithm"`
	CreatedAt string                `json:"created_at"`
	Id        int64                 `json:"id"`
	Name      string                `json:"name"`
	Secret    string                `json:"secret"`
	UpdatedAt string                `json:"updated_at"`
	User      *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionDnsTsigKeyCreateResponse struct {
	Action *ActionDnsTsigKeyCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsTsigKey *ActionDnsTsigKeyCreateOutput `json:"dns_tsig_key"`
	}

	// Action output without the namespace
	Output *ActionDnsTsigKeyCreateOutput
}

// Prepare the action for invocation
func (action *ActionDnsTsigKeyCreate) Prepare() *ActionDnsTsigKeyCreateInvocation {
	return &ActionDnsTsigKeyCreateInvocation{
		Action: action,
		Path:   "/v7.0/dns_tsig_keys",
	}
}

// ActionDnsTsigKeyCreateInvocation is used to configure action for invocation
type ActionDnsTsigKeyCreateInvocation struct {
	// Pointer to the action
	Action *ActionDnsTsigKeyCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsTsigKeyCreateInput
	// Global meta input parameters
	MetaInput *ActionDnsTsigKeyCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsTsigKeyCreateInvocation) NewInput() *ActionDnsTsigKeyCreateInput {
	inv.Input = &ActionDnsTsigKeyCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsTsigKeyCreateInvocation) SetInput(input *ActionDnsTsigKeyCreateInput) *ActionDnsTsigKeyCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsTsigKeyCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsTsigKeyCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsTsigKeyCreateInvocation) NewMetaInput() *ActionDnsTsigKeyCreateMetaGlobalInput {
	inv.MetaInput = &ActionDnsTsigKeyCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsTsigKeyCreateInvocation) SetMetaInput(input *ActionDnsTsigKeyCreateMetaGlobalInput) *ActionDnsTsigKeyCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsTsigKeyCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsTsigKeyCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsTsigKeyCreateInvocation) Call() (*ActionDnsTsigKeyCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsTsigKeyCreateInvocation) callAsBody() (*ActionDnsTsigKeyCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsTsigKeyCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsTsigKey
	}
	return resp, err
}

func (inv *ActionDnsTsigKeyCreateInvocation) makeAllInputParams() *ActionDnsTsigKeyCreateRequest {
	return &ActionDnsTsigKeyCreateRequest{
		DnsTsigKey: inv.makeInputParams(),
		Meta:       inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsTsigKeyCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Algorithm") {
			ret["algorithm"] = inv.Input.Algorithm
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["user"] = nil
			} else {
				ret["user"] = inv.Input.User
			}
		}
	}

	return ret
}

func (inv *ActionDnsTsigKeyCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
