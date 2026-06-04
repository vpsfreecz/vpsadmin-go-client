package client

import (
	"net/url"
	"strings"
)

// ActionNodeTransferConnectionShow is a type for action Node_transfer_connection#Show
type ActionNodeTransferConnectionShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeTransferConnectionShow(client *Client) *ActionNodeTransferConnectionShow {
	return &ActionNodeTransferConnectionShow{
		Client: client,
	}
}

// ActionNodeTransferConnectionShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeTransferConnectionShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeTransferConnectionShowMetaGlobalInput) SetIncludes(value string) *ActionNodeTransferConnectionShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeTransferConnectionShowMetaGlobalInput) SetNo(value bool) *ActionNodeTransferConnectionShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeTransferConnectionShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeTransferConnectionShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeTransferConnectionShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeTransferConnectionShowOutput is a type for action output parameters
type ActionNodeTransferConnectionShowOutput struct {
	CreatedAt   string                "json:\"created_at\""
	Enabled     bool                  "json:\"enabled\""
	Id          int64                 "json:\"id\""
	NodeA       *ActionNodeShowOutput "json:\"node_a\""
	NodeAIpAddr string                "json:\"node_a_ip_addr\""
	NodeB       *ActionNodeShowOutput "json:\"node_b\""
	NodeBIpAddr string                "json:\"node_b_ip_addr\""
	UpdatedAt   string                "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionNodeTransferConnectionShowResponse struct {
	Action *ActionNodeTransferConnectionShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeTransferConnection *ActionNodeTransferConnectionShowOutput "json:\"node_transfer_connection\""
	}

	// Action output without the namespace
	Output *ActionNodeTransferConnectionShowOutput
}

// Prepare the action for invocation
func (action *ActionNodeTransferConnectionShow) Prepare() *ActionNodeTransferConnectionShowInvocation {
	return &ActionNodeTransferConnectionShowInvocation{
		Action: action,
		Path:   "/v7.0/node_transfer_connections/{node_transfer_connection_id}",
	}
}

// ActionNodeTransferConnectionShowInvocation is used to configure action for invocation
type ActionNodeTransferConnectionShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeTransferConnectionShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeTransferConnectionShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeTransferConnectionShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeTransferConnectionShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeTransferConnectionShowInvocation) SetPathParamString(param string, value string) *ActionNodeTransferConnectionShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeTransferConnectionShowInvocation) NewMetaInput() *ActionNodeTransferConnectionShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeTransferConnectionShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeTransferConnectionShowInvocation) SetMetaInput(input *ActionNodeTransferConnectionShowMetaGlobalInput) *ActionNodeTransferConnectionShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeTransferConnectionShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeTransferConnectionShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeTransferConnectionShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeTransferConnectionShowInvocation) Call() (*ActionNodeTransferConnectionShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeTransferConnectionShowInvocation) callAsQuery() (*ActionNodeTransferConnectionShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeTransferConnectionShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeTransferConnection
	}
	return resp, err
}

func (inv *ActionNodeTransferConnectionShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
