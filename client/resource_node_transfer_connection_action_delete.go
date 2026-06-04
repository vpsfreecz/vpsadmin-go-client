package client

import (
	"net/url"
	"strings"
)

// ActionNodeTransferConnectionDelete is a type for action Node_transfer_connection#Delete
type ActionNodeTransferConnectionDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeTransferConnectionDelete(client *Client) *ActionNodeTransferConnectionDelete {
	return &ActionNodeTransferConnectionDelete{
		Client: client,
	}
}

// ActionNodeTransferConnectionDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionNodeTransferConnectionDeleteMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeTransferConnectionDeleteMetaGlobalInput) SetIncludes(value string) *ActionNodeTransferConnectionDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeTransferConnectionDeleteMetaGlobalInput) SetNo(value bool) *ActionNodeTransferConnectionDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeTransferConnectionDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionNodeTransferConnectionDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeTransferConnectionDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeTransferConnectionDeleteRequest is a type for the entire action request
type ActionNodeTransferConnectionDeleteRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionNodeTransferConnectionDeleteResponse struct {
	Action *ActionNodeTransferConnectionDelete "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionNodeTransferConnectionDelete) Prepare() *ActionNodeTransferConnectionDeleteInvocation {
	return &ActionNodeTransferConnectionDeleteInvocation{
		Action: action,
		Path:   "/v7.0/node_transfer_connections/{node_transfer_connection_id}",
	}
}

// ActionNodeTransferConnectionDeleteInvocation is used to configure action for invocation
type ActionNodeTransferConnectionDeleteInvocation struct {
	// Pointer to the action
	Action *ActionNodeTransferConnectionDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeTransferConnectionDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeTransferConnectionDeleteInvocation) SetPathParamInt(param string, value int64) *ActionNodeTransferConnectionDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeTransferConnectionDeleteInvocation) SetPathParamString(param string, value string) *ActionNodeTransferConnectionDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeTransferConnectionDeleteInvocation) NewMetaInput() *ActionNodeTransferConnectionDeleteMetaGlobalInput {
	inv.MetaInput = &ActionNodeTransferConnectionDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeTransferConnectionDeleteInvocation) SetMetaInput(input *ActionNodeTransferConnectionDeleteMetaGlobalInput) *ActionNodeTransferConnectionDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeTransferConnectionDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeTransferConnectionDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeTransferConnectionDeleteInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeTransferConnectionDeleteInvocation) Call() (*ActionNodeTransferConnectionDeleteResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionNodeTransferConnectionDeleteInvocation) callAsBody() (*ActionNodeTransferConnectionDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNodeTransferConnectionDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionNodeTransferConnectionDeleteInvocation) makeAllInputParams() *ActionNodeTransferConnectionDeleteRequest {
	return &ActionNodeTransferConnectionDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNodeTransferConnectionDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
