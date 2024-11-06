package client

import (
	"strings"
)

// ActionUserPublicKeyCreate is a type for action User.Public_key#Create
type ActionUserPublicKeyCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserPublicKeyCreate(client *Client) *ActionUserPublicKeyCreate {
	return &ActionUserPublicKeyCreate{
		Client: client,
	}
}

// ActionUserPublicKeyCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserPublicKeyCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserPublicKeyCreateMetaGlobalInput) SetIncludes(value string) *ActionUserPublicKeyCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserPublicKeyCreateMetaGlobalInput) SetNo(value bool) *ActionUserPublicKeyCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPublicKeyCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPublicKeyCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserPublicKeyCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPublicKeyCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserPublicKeyCreateInput is a type for action input parameters
type ActionUserPublicKeyCreateInput struct {
	AutoAdd bool   `json:"auto_add"`
	Key     string `json:"key"`
	Label   string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAutoAdd sets parameter AutoAdd to value and selects it for sending
func (in *ActionUserPublicKeyCreateInput) SetAutoAdd(value bool) *ActionUserPublicKeyCreateInput {
	in.AutoAdd = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AutoAdd"] = nil
	return in
}

// SetKey sets parameter Key to value and selects it for sending
func (in *ActionUserPublicKeyCreateInput) SetKey(value string) *ActionUserPublicKeyCreateInput {
	in.Key = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Key"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionUserPublicKeyCreateInput) SetLabel(value string) *ActionUserPublicKeyCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPublicKeyCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPublicKeyCreateInput) SelectParameters(params ...string) *ActionUserPublicKeyCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserPublicKeyCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserPublicKeyCreateInput) UnselectParameters(params ...string) *ActionUserPublicKeyCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserPublicKeyCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserPublicKeyCreateRequest is a type for the entire action request
type ActionUserPublicKeyCreateRequest struct {
	PublicKey map[string]interface{} `json:"public_key"`
	Meta      map[string]interface{} `json:"_meta"`
}

// ActionUserPublicKeyCreateOutput is a type for action output parameters
type ActionUserPublicKeyCreateOutput struct {
	AutoAdd     bool   `json:"auto_add"`
	Comment     string `json:"comment"`
	CreatedAt   string `json:"created_at"`
	Fingerprint string `json:"fingerprint"`
	Id          int64  `json:"id"`
	Key         string `json:"key"`
	Label       string `json:"label"`
	UpdatedAt   string `json:"updated_at"`
}

// Type for action response, including envelope
type ActionUserPublicKeyCreateResponse struct {
	Action *ActionUserPublicKeyCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PublicKey *ActionUserPublicKeyCreateOutput `json:"public_key"`
	}

	// Action output without the namespace
	Output *ActionUserPublicKeyCreateOutput
}

// Prepare the action for invocation
func (action *ActionUserPublicKeyCreate) Prepare() *ActionUserPublicKeyCreateInvocation {
	return &ActionUserPublicKeyCreateInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/public_keys",
	}
}

// ActionUserPublicKeyCreateInvocation is used to configure action for invocation
type ActionUserPublicKeyCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserPublicKeyCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserPublicKeyCreateInput
	// Global meta input parameters
	MetaInput *ActionUserPublicKeyCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserPublicKeyCreateInvocation) SetPathParamInt(param string, value int64) *ActionUserPublicKeyCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserPublicKeyCreateInvocation) SetPathParamString(param string, value string) *ActionUserPublicKeyCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserPublicKeyCreateInvocation) NewInput() *ActionUserPublicKeyCreateInput {
	inv.Input = &ActionUserPublicKeyCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserPublicKeyCreateInvocation) SetInput(input *ActionUserPublicKeyCreateInput) *ActionUserPublicKeyCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserPublicKeyCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserPublicKeyCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserPublicKeyCreateInvocation) NewMetaInput() *ActionUserPublicKeyCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserPublicKeyCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserPublicKeyCreateInvocation) SetMetaInput(input *ActionUserPublicKeyCreateMetaGlobalInput) *ActionUserPublicKeyCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserPublicKeyCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserPublicKeyCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserPublicKeyCreateInvocation) Call() (*ActionUserPublicKeyCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserPublicKeyCreateInvocation) callAsBody() (*ActionUserPublicKeyCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserPublicKeyCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.PublicKey
	}
	return resp, err
}

func (inv *ActionUserPublicKeyCreateInvocation) makeAllInputParams() *ActionUserPublicKeyCreateRequest {
	return &ActionUserPublicKeyCreateRequest{
		PublicKey: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserPublicKeyCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AutoAdd") {
			ret["auto_add"] = inv.Input.AutoAdd
		}
		if inv.IsParameterSelected("Key") {
			ret["key"] = inv.Input.Key
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
	}

	return ret
}

func (inv *ActionUserPublicKeyCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
