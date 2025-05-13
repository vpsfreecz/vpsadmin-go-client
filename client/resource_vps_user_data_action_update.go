package client

import (
	"strings"
)

// ActionVpsUserDataUpdate is a type for action Vps_user_data#Update
type ActionVpsUserDataUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsUserDataUpdate(client *Client) *ActionVpsUserDataUpdate {
	return &ActionVpsUserDataUpdate{
		Client: client,
	}
}

// ActionVpsUserDataUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsUserDataUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsUserDataUpdateMetaGlobalInput) SetIncludes(value string) *ActionVpsUserDataUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsUserDataUpdateMetaGlobalInput) SetNo(value bool) *ActionVpsUserDataUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUserDataUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsUserDataUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsUserDataUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataUpdateInput is a type for action input parameters
type ActionVpsUserDataUpdateInput struct {
	Content string `json:"content"`
	Format  string `json:"format"`
	Label   string `json:"label"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetContent sets parameter Content to value and selects it for sending
func (in *ActionVpsUserDataUpdateInput) SetContent(value string) *ActionVpsUserDataUpdateInput {
	in.Content = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Content"] = nil
	return in
}

// SetFormat sets parameter Format to value and selects it for sending
func (in *ActionVpsUserDataUpdateInput) SetFormat(value string) *ActionVpsUserDataUpdateInput {
	in.Format = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Format"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionVpsUserDataUpdateInput) SetLabel(value string) *ActionVpsUserDataUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUserDataUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataUpdateInput) SelectParameters(params ...string) *ActionVpsUserDataUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsUserDataUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsUserDataUpdateInput) UnselectParameters(params ...string) *ActionVpsUserDataUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsUserDataUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataUpdateRequest is a type for the entire action request
type ActionVpsUserDataUpdateRequest struct {
	VpsUserData map[string]interface{} `json:"vps_user_data"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionVpsUserDataUpdateOutput is a type for action output parameters
type ActionVpsUserDataUpdateOutput struct {
	Content   string                `json:"content"`
	CreatedAt string                `json:"created_at"`
	Format    string                `json:"format"`
	Id        int64                 `json:"id"`
	Label     string                `json:"label"`
	UpdatedAt string                `json:"updated_at"`
	User      *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionVpsUserDataUpdateResponse struct {
	Action *ActionVpsUserDataUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsUserData *ActionVpsUserDataUpdateOutput `json:"vps_user_data"`
	}

	// Action output without the namespace
	Output *ActionVpsUserDataUpdateOutput
}

// Prepare the action for invocation
func (action *ActionVpsUserDataUpdate) Prepare() *ActionVpsUserDataUpdateInvocation {
	return &ActionVpsUserDataUpdateInvocation{
		Action: action,
		Path:   "/v7.0/vps_user_data/{vps_user_data_id}",
	}
}

// ActionVpsUserDataUpdateInvocation is used to configure action for invocation
type ActionVpsUserDataUpdateInvocation struct {
	// Pointer to the action
	Action *ActionVpsUserDataUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsUserDataUpdateInput
	// Global meta input parameters
	MetaInput *ActionVpsUserDataUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsUserDataUpdateInvocation) SetPathParamInt(param string, value int64) *ActionVpsUserDataUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsUserDataUpdateInvocation) SetPathParamString(param string, value string) *ActionVpsUserDataUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsUserDataUpdateInvocation) NewInput() *ActionVpsUserDataUpdateInput {
	inv.Input = &ActionVpsUserDataUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsUserDataUpdateInvocation) SetInput(input *ActionVpsUserDataUpdateInput) *ActionVpsUserDataUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsUserDataUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsUserDataUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsUserDataUpdateInvocation) NewMetaInput() *ActionVpsUserDataUpdateMetaGlobalInput {
	inv.MetaInput = &ActionVpsUserDataUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsUserDataUpdateInvocation) SetMetaInput(input *ActionVpsUserDataUpdateMetaGlobalInput) *ActionVpsUserDataUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsUserDataUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsUserDataUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsUserDataUpdateInvocation) Call() (*ActionVpsUserDataUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsUserDataUpdateInvocation) callAsBody() (*ActionVpsUserDataUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsUserDataUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsUserData
	}
	return resp, err
}

func (inv *ActionVpsUserDataUpdateInvocation) makeAllInputParams() *ActionVpsUserDataUpdateRequest {
	return &ActionVpsUserDataUpdateRequest{
		VpsUserData: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsUserDataUpdateInvocation) makeInputParams() map[string]interface{} {
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
	}

	return ret
}

func (inv *ActionVpsUserDataUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
