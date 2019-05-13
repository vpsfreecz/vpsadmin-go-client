package client

import (
	"strings"
)

// ActionNetworkShow is a type for action Network#Show
type ActionNetworkShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkShow(client *Client) *ActionNetworkShow {
	return &ActionNetworkShow{
		Client: client,
	}
}

// ActionNetworkShowMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkShowMetaGlobalInput) SetNo(value bool) *ActionNetworkShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkShowMetaGlobalInput) SetIncludes(value string) *ActionNetworkShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkShowMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionNetworkShowOutput is a type for action output parameters
type ActionNetworkShowOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Location *ActionLocationShowOutput `json:"location"`
	IpVersion int64 `json:"ip_version"`
	Address string `json:"address"`
	Prefix int64 `json:"prefix"`
	Role string `json:"role"`
	Managed bool `json:"managed"`
	SplitAccess string `json:"split_access"`
	SplitPrefix int64 `json:"split_prefix"`
	Autopick bool `json:"autopick"`
	Size int64 `json:"size"`
	Used int64 `json:"used"`
	Assigned int64 `json:"assigned"`
	Owned int64 `json:"owned"`
}


// Type for action response, including envelope
type ActionNetworkShowResponse struct {
	Action *ActionNetworkShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Network *ActionNetworkShowOutput `json:"network"`
	}

	// Action output without the namespace
	Output *ActionNetworkShowOutput
}


// Prepare the action for invocation
func (action *ActionNetworkShow) Prepare() *ActionNetworkShowInvocation {
	return &ActionNetworkShowInvocation{
		Action: action,
		Path: "/v5.0/networks/{network_id}",
	}
}

// ActionNetworkShowInvocation is used to configure action for invocation
type ActionNetworkShowInvocation struct {
	// Pointer to the action
	Action *ActionNetworkShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNetworkShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNetworkShowInvocation) SetPathParamInt(param string, value int64) *ActionNetworkShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNetworkShowInvocation) SetPathParamString(param string, value string) *ActionNetworkShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkShowInvocation) NewMetaInput() *ActionNetworkShowMetaGlobalInput {
	inv.MetaInput = &ActionNetworkShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkShowInvocation) SetMetaInput(input *ActionNetworkShowMetaGlobalInput) *ActionNetworkShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNetworkShowInvocation) Call() (*ActionNetworkShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNetworkShowInvocation) callAsQuery() (*ActionNetworkShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNetworkShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Network
	}
	return resp, err
}




func (inv *ActionNetworkShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

