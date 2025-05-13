package client

import (
	"strings"
)

// ActionVpsUserDataDelete is a type for action Vps_user_data#Delete
type ActionVpsUserDataDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsUserDataDelete(client *Client) *ActionVpsUserDataDelete {
	return &ActionVpsUserDataDelete{
		Client: client,
	}
}

// ActionVpsUserDataDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionVpsUserDataDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsUserDataDeleteMetaGlobalInput) SetIncludes(value string) *ActionVpsUserDataDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsUserDataDeleteMetaGlobalInput) SetNo(value bool) *ActionVpsUserDataDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUserDataDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionVpsUserDataDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsUserDataDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataDeleteRequest is a type for the entire action request
type ActionVpsUserDataDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionVpsUserDataDeleteResponse struct {
	Action *ActionVpsUserDataDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionVpsUserDataDelete) Prepare() *ActionVpsUserDataDeleteInvocation {
	return &ActionVpsUserDataDeleteInvocation{
		Action: action,
		Path:   "/v7.0/vps_user_data/{vps_user_data_id}",
	}
}

// ActionVpsUserDataDeleteInvocation is used to configure action for invocation
type ActionVpsUserDataDeleteInvocation struct {
	// Pointer to the action
	Action *ActionVpsUserDataDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsUserDataDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsUserDataDeleteInvocation) SetPathParamInt(param string, value int64) *ActionVpsUserDataDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsUserDataDeleteInvocation) SetPathParamString(param string, value string) *ActionVpsUserDataDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsUserDataDeleteInvocation) NewMetaInput() *ActionVpsUserDataDeleteMetaGlobalInput {
	inv.MetaInput = &ActionVpsUserDataDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsUserDataDeleteInvocation) SetMetaInput(input *ActionVpsUserDataDeleteMetaGlobalInput) *ActionVpsUserDataDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsUserDataDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsUserDataDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsUserDataDeleteInvocation) Call() (*ActionVpsUserDataDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionVpsUserDataDeleteInvocation) callAsBody() (*ActionVpsUserDataDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsUserDataDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionVpsUserDataDeleteInvocation) makeAllInputParams() *ActionVpsUserDataDeleteRequest {
	return &ActionVpsUserDataDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionVpsUserDataDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
