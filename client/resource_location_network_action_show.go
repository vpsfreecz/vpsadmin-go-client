package client

import (
	"strings"
)

// ActionLocationNetworkShow is a type for action Location_network#Show
type ActionLocationNetworkShow struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationNetworkShow(client *Client) *ActionLocationNetworkShow {
	return &ActionLocationNetworkShow{
		Client: client,
	}
}

// ActionLocationNetworkShowMetaGlobalInput is a type for action global meta input parameters
type ActionLocationNetworkShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationNetworkShowMetaGlobalInput) SetIncludes(value string) *ActionLocationNetworkShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationNetworkShowMetaGlobalInput) SetNo(value bool) *ActionLocationNetworkShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationNetworkShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationNetworkShowMetaGlobalInput) SelectParameters(params ...string) *ActionLocationNetworkShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationNetworkShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationNetworkShowOutput is a type for action output parameters
type ActionLocationNetworkShowOutput struct {
	Autopick bool                      `json:"autopick"`
	Id       int64                     `json:"id"`
	Location *ActionLocationShowOutput `json:"location"`
	Network  *ActionNetworkShowOutput  `json:"network"`
	Primary  bool                      `json:"primary"`
	Priority int64                     `json:"priority"`
	Userpick bool                      `json:"userpick"`
}

// Type for action response, including envelope
type ActionLocationNetworkShowResponse struct {
	Action *ActionLocationNetworkShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		LocationNetwork *ActionLocationNetworkShowOutput `json:"location_network"`
	}

	// Action output without the namespace
	Output *ActionLocationNetworkShowOutput
}

// Prepare the action for invocation
func (action *ActionLocationNetworkShow) Prepare() *ActionLocationNetworkShowInvocation {
	return &ActionLocationNetworkShowInvocation{
		Action: action,
		Path:   "/v6.0/location_networks/{location_network_id}",
	}
}

// ActionLocationNetworkShowInvocation is used to configure action for invocation
type ActionLocationNetworkShowInvocation struct {
	// Pointer to the action
	Action *ActionLocationNetworkShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionLocationNetworkShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionLocationNetworkShowInvocation) SetPathParamInt(param string, value int64) *ActionLocationNetworkShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionLocationNetworkShowInvocation) SetPathParamString(param string, value string) *ActionLocationNetworkShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLocationNetworkShowInvocation) NewMetaInput() *ActionLocationNetworkShowMetaGlobalInput {
	inv.MetaInput = &ActionLocationNetworkShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationNetworkShowInvocation) SetMetaInput(input *ActionLocationNetworkShowMetaGlobalInput) *ActionLocationNetworkShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationNetworkShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationNetworkShowInvocation) Call() (*ActionLocationNetworkShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionLocationNetworkShowInvocation) callAsQuery() (*ActionLocationNetworkShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionLocationNetworkShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.LocationNetwork
	}
	return resp, err
}

func (inv *ActionLocationNetworkShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
