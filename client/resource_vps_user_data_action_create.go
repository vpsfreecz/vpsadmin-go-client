package client

import ()

// ActionVpsUserDataCreate is a type for action Vps_user_data#Create
type ActionVpsUserDataCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsUserDataCreate(client *Client) *ActionVpsUserDataCreate {
	return &ActionVpsUserDataCreate{
		Client: client,
	}
}

// ActionVpsUserDataCreateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsUserDataCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsUserDataCreateMetaGlobalInput) SetIncludes(value string) *ActionVpsUserDataCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsUserDataCreateMetaGlobalInput) SetNo(value bool) *ActionVpsUserDataCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUserDataCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataCreateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsUserDataCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsUserDataCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataCreateInput is a type for action input parameters
type ActionVpsUserDataCreateInput struct {
	Content string `json:"content"`
	Format  string `json:"format"`
	Label   string `json:"label"`
	User    int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetContent sets parameter Content to value and selects it for sending
func (in *ActionVpsUserDataCreateInput) SetContent(value string) *ActionVpsUserDataCreateInput {
	in.Content = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Content"] = nil
	return in
}

// SetFormat sets parameter Format to value and selects it for sending
func (in *ActionVpsUserDataCreateInput) SetFormat(value string) *ActionVpsUserDataCreateInput {
	in.Format = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Format"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionVpsUserDataCreateInput) SetLabel(value string) *ActionVpsUserDataCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsUserDataCreateInput) SetUser(value int64) *ActionVpsUserDataCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionVpsUserDataCreateInput) SetUserNil(set bool) *ActionVpsUserDataCreateInput {
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

// SelectParameters sets parameters from ActionVpsUserDataCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataCreateInput) SelectParameters(params ...string) *ActionVpsUserDataCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsUserDataCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsUserDataCreateInput) UnselectParameters(params ...string) *ActionVpsUserDataCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsUserDataCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataCreateRequest is a type for the entire action request
type ActionVpsUserDataCreateRequest struct {
	VpsUserData map[string]interface{} `json:"vps_user_data"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionVpsUserDataCreateOutput is a type for action output parameters
type ActionVpsUserDataCreateOutput struct {
	Content   string                `json:"content"`
	CreatedAt string                `json:"created_at"`
	Format    string                `json:"format"`
	Id        int64                 `json:"id"`
	Label     string                `json:"label"`
	UpdatedAt string                `json:"updated_at"`
	User      *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionVpsUserDataCreateResponse struct {
	Action *ActionVpsUserDataCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsUserData *ActionVpsUserDataCreateOutput `json:"vps_user_data"`
	}

	// Action output without the namespace
	Output *ActionVpsUserDataCreateOutput
}

// Prepare the action for invocation
func (action *ActionVpsUserDataCreate) Prepare() *ActionVpsUserDataCreateInvocation {
	return &ActionVpsUserDataCreateInvocation{
		Action: action,
		Path:   "/v7.0/vps_user_data",
	}
}

// ActionVpsUserDataCreateInvocation is used to configure action for invocation
type ActionVpsUserDataCreateInvocation struct {
	// Pointer to the action
	Action *ActionVpsUserDataCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsUserDataCreateInput
	// Global meta input parameters
	MetaInput *ActionVpsUserDataCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsUserDataCreateInvocation) NewInput() *ActionVpsUserDataCreateInput {
	inv.Input = &ActionVpsUserDataCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsUserDataCreateInvocation) SetInput(input *ActionVpsUserDataCreateInput) *ActionVpsUserDataCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsUserDataCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsUserDataCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsUserDataCreateInvocation) NewMetaInput() *ActionVpsUserDataCreateMetaGlobalInput {
	inv.MetaInput = &ActionVpsUserDataCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsUserDataCreateInvocation) SetMetaInput(input *ActionVpsUserDataCreateMetaGlobalInput) *ActionVpsUserDataCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsUserDataCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsUserDataCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsUserDataCreateInvocation) Call() (*ActionVpsUserDataCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsUserDataCreateInvocation) callAsBody() (*ActionVpsUserDataCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsUserDataCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsUserData
	}
	return resp, err
}

func (inv *ActionVpsUserDataCreateInvocation) makeAllInputParams() *ActionVpsUserDataCreateRequest {
	return &ActionVpsUserDataCreateRequest{
		VpsUserData: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsUserDataCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Content") {
			ret["content"] = inv.Input.Content
		}
		if inv.IsParameterSelected("Format") {
			ret["format"] = inv.Input.Format
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
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

func (inv *ActionVpsUserDataCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
