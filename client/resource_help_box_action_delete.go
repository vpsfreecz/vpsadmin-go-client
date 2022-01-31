package client

import (
	"strings"
)

// ActionHelpBoxDelete is a type for action Help_box#Delete
type ActionHelpBoxDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionHelpBoxDelete(client *Client) *ActionHelpBoxDelete {
	return &ActionHelpBoxDelete{
		Client: client,
	}
}

// ActionHelpBoxDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionHelpBoxDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHelpBoxDeleteMetaGlobalInput) SetIncludes(value string) *ActionHelpBoxDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHelpBoxDeleteMetaGlobalInput) SetNo(value bool) *ActionHelpBoxDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionHelpBoxDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHelpBoxDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionHelpBoxDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHelpBoxDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHelpBoxDeleteRequest is a type for the entire action request
type ActionHelpBoxDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionHelpBoxDeleteResponse struct {
	Action *ActionHelpBoxDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionHelpBoxDelete) Prepare() *ActionHelpBoxDeleteInvocation {
	return &ActionHelpBoxDeleteInvocation{
		Action: action,
		Path:   "/v6.0/help_boxes/{help_box_id}",
	}
}

// ActionHelpBoxDeleteInvocation is used to configure action for invocation
type ActionHelpBoxDeleteInvocation struct {
	// Pointer to the action
	Action *ActionHelpBoxDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionHelpBoxDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionHelpBoxDeleteInvocation) SetPathParamInt(param string, value int64) *ActionHelpBoxDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionHelpBoxDeleteInvocation) SetPathParamString(param string, value string) *ActionHelpBoxDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHelpBoxDeleteInvocation) NewMetaInput() *ActionHelpBoxDeleteMetaGlobalInput {
	inv.MetaInput = &ActionHelpBoxDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHelpBoxDeleteInvocation) SetMetaInput(input *ActionHelpBoxDeleteMetaGlobalInput) *ActionHelpBoxDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHelpBoxDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionHelpBoxDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHelpBoxDeleteInvocation) Call() (*ActionHelpBoxDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionHelpBoxDeleteInvocation) callAsBody() (*ActionHelpBoxDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionHelpBoxDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionHelpBoxDeleteInvocation) makeAllInputParams() *ActionHelpBoxDeleteRequest {
	return &ActionHelpBoxDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionHelpBoxDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
