package client

import (
	"strings"
)

// ActionUserPublicKeyUpdate is a type for action User.Public_key#Update
type ActionUserPublicKeyUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserPublicKeyUpdate(client *Client) *ActionUserPublicKeyUpdate {
	return &ActionUserPublicKeyUpdate{
		Client: client,
	}
}

// ActionUserPublicKeyUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserPublicKeyUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserPublicKeyUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserPublicKeyUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserPublicKeyUpdateMetaGlobalInput) SetNo(value bool) *ActionUserPublicKeyUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPublicKeyUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPublicKeyUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserPublicKeyUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPublicKeyUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserPublicKeyUpdateInput is a type for action input parameters
type ActionUserPublicKeyUpdateInput struct {
	AutoAdd bool   `json:"auto_add"`
	Key     string `json:"key"`
	Label   string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAutoAdd sets parameter AutoAdd to value and selects it for sending
func (in *ActionUserPublicKeyUpdateInput) SetAutoAdd(value bool) *ActionUserPublicKeyUpdateInput {
	in.AutoAdd = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AutoAdd"] = nil
	return in
}

// SetKey sets parameter Key to value and selects it for sending
func (in *ActionUserPublicKeyUpdateInput) SetKey(value string) *ActionUserPublicKeyUpdateInput {
	in.Key = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Key"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionUserPublicKeyUpdateInput) SetLabel(value string) *ActionUserPublicKeyUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPublicKeyUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPublicKeyUpdateInput) SelectParameters(params ...string) *ActionUserPublicKeyUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserPublicKeyUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserPublicKeyUpdateInput) UnselectParameters(params ...string) *ActionUserPublicKeyUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserPublicKeyUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserPublicKeyUpdateRequest is a type for the entire action request
type ActionUserPublicKeyUpdateRequest struct {
	PublicKey map[string]interface{} `json:"public_key"`
	Meta      map[string]interface{} `json:"_meta"`
}

// ActionUserPublicKeyUpdateOutput is a type for action output parameters
type ActionUserPublicKeyUpdateOutput struct {
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
type ActionUserPublicKeyUpdateResponse struct {
	Action *ActionUserPublicKeyUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PublicKey *ActionUserPublicKeyUpdateOutput `json:"public_key"`
	}

	// Action output without the namespace
	Output *ActionUserPublicKeyUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserPublicKeyUpdate) Prepare() *ActionUserPublicKeyUpdateInvocation {
	return &ActionUserPublicKeyUpdateInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/public_keys/{public_key_id}",
	}
}

// ActionUserPublicKeyUpdateInvocation is used to configure action for invocation
type ActionUserPublicKeyUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserPublicKeyUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserPublicKeyUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserPublicKeyUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserPublicKeyUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserPublicKeyUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserPublicKeyUpdateInvocation) SetPathParamString(param string, value string) *ActionUserPublicKeyUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserPublicKeyUpdateInvocation) NewInput() *ActionUserPublicKeyUpdateInput {
	inv.Input = &ActionUserPublicKeyUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserPublicKeyUpdateInvocation) SetInput(input *ActionUserPublicKeyUpdateInput) *ActionUserPublicKeyUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserPublicKeyUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserPublicKeyUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserPublicKeyUpdateInvocation) NewMetaInput() *ActionUserPublicKeyUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserPublicKeyUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserPublicKeyUpdateInvocation) SetMetaInput(input *ActionUserPublicKeyUpdateMetaGlobalInput) *ActionUserPublicKeyUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserPublicKeyUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserPublicKeyUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserPublicKeyUpdateInvocation) Call() (*ActionUserPublicKeyUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserPublicKeyUpdateInvocation) callAsBody() (*ActionUserPublicKeyUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserPublicKeyUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.PublicKey
	}
	return resp, err
}

func (inv *ActionUserPublicKeyUpdateInvocation) makeAllInputParams() *ActionUserPublicKeyUpdateRequest {
	return &ActionUserPublicKeyUpdateRequest{
		PublicKey: inv.makeInputParams(),
		Meta:      inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserPublicKeyUpdateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionUserPublicKeyUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
