package client

import (
)

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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestChangeCreateMetaGlobalInput) SetIncludes(value string) *ActionUserRequestChangeCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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
	ChangeReason string `json:"change_reason"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserRequestChangeCreateInput) SetFullName(value string) *ActionUserRequestChangeCreateInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
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
// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserRequestChangeCreateInput) SetAddress(value string) *ActionUserRequestChangeCreateInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
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

func (in *ActionUserRequestChangeCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestChangeCreateRequest is a type for the entire action request
type ActionUserRequestChangeCreateRequest struct {
	Change map[string]interface{} `json:"change"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserRequestChangeCreateOutput is a type for action output parameters
type ActionUserRequestChangeCreateOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	State string `json:"state"`
	ApiIpAddr string `json:"api_ip_addr"`
	ApiIpPtr string `json:"api_ip_ptr"`
	ClientIpAddr string `json:"client_ip_addr"`
	ClientIpPtr string `json:"client_ip_ptr"`
	Admin *ActionUserShowOutput `json:"admin"`
	AdminResponse string `json:"admin_response"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Label string `json:"label"`
	ChangeReason string `json:"change_reason"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
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
		Path: "/v5.0/user_request/changes",
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
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserRequestChangeCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ChangeReason") {
			ret["change_reason"] = inv.Input.ChangeReason
		}
		if inv.IsParameterSelected("FullName") {
			ret["full_name"] = inv.Input.FullName
		}
		if inv.IsParameterSelected("Email") {
			ret["email"] = inv.Input.Email
		}
		if inv.IsParameterSelected("Address") {
			ret["address"] = inv.Input.Address
		}
	}

	return ret
}

func (inv *ActionUserRequestChangeCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
