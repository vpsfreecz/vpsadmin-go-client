package client

import (
	"strings"
)

// ActionNetworkInterfaceShow is a type for action Network_interface#Show
type ActionNetworkInterfaceShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkInterfaceShow(client *Client) *ActionNetworkInterfaceShow {
	return &ActionNetworkInterfaceShow{
		Client: client,
	}
}

// ActionNetworkInterfaceShowMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkInterfaceShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkInterfaceShowMetaGlobalInput) SetIncludes(value string) *ActionNetworkInterfaceShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkInterfaceShowMetaGlobalInput) SetNo(value bool) *ActionNetworkInterfaceShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkInterfaceShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkInterfaceShowMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkInterfaceShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkInterfaceShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkInterfaceShowOutput is a type for action output parameters
type ActionNetworkInterfaceShowOutput struct {
	Id   int64                `json:"id"`
	Mac  string               `json:"mac"`
	Name string               `json:"name"`
	Type string               `json:"type"`
	Vps  *ActionVpsShowOutput `json:"vps"`
}

// Type for action response, including envelope
type ActionNetworkInterfaceShowResponse struct {
	Action *ActionNetworkInterfaceShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NetworkInterface *ActionNetworkInterfaceShowOutput `json:"network_interface"`
	}

	// Action output without the namespace
	Output *ActionNetworkInterfaceShowOutput
}

// Prepare the action for invocation
func (action *ActionNetworkInterfaceShow) Prepare() *ActionNetworkInterfaceShowInvocation {
	return &ActionNetworkInterfaceShowInvocation{
		Action: action,
		Path:   "/v6.0/network_interfaces/{network_interface_id}",
	}
}

// ActionNetworkInterfaceShowInvocation is used to configure action for invocation
type ActionNetworkInterfaceShowInvocation struct {
	// Pointer to the action
	Action *ActionNetworkInterfaceShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNetworkInterfaceShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNetworkInterfaceShowInvocation) SetPathParamInt(param string, value int64) *ActionNetworkInterfaceShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNetworkInterfaceShowInvocation) SetPathParamString(param string, value string) *ActionNetworkInterfaceShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkInterfaceShowInvocation) NewMetaInput() *ActionNetworkInterfaceShowMetaGlobalInput {
	inv.MetaInput = &ActionNetworkInterfaceShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkInterfaceShowInvocation) SetMetaInput(input *ActionNetworkInterfaceShowMetaGlobalInput) *ActionNetworkInterfaceShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkInterfaceShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNetworkInterfaceShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkInterfaceShowInvocation) Call() (*ActionNetworkInterfaceShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNetworkInterfaceShowInvocation) callAsQuery() (*ActionNetworkInterfaceShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNetworkInterfaceShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NetworkInterface
	}
	return resp, err
}

func (inv *ActionNetworkInterfaceShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
