package client

import (
	"strings"
)

// ActionVpsConfigShow is a type for action Vps_config#Show
type ActionVpsConfigShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsConfigShow(client *Client) *ActionVpsConfigShow {
	return &ActionVpsConfigShow{
		Client: client,
	}
}

// ActionVpsConfigShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsConfigShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsConfigShowMetaGlobalInput) SetIncludes(value string) *ActionVpsConfigShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsConfigShowMetaGlobalInput) SetNo(value bool) *ActionVpsConfigShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConfigShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConfigShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsConfigShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConfigShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConfigShowOutput is a type for action output parameters
type ActionVpsConfigShowOutput struct {
	Config string `json:"config"`
	Id     int64  `json:"id"`
	Label  string `json:"label"`
	Name   string `json:"name"`
}

// Type for action response, including envelope
type ActionVpsConfigShowResponse struct {
	Action *ActionVpsConfigShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsConfig *ActionVpsConfigShowOutput `json:"vps_config"`
	}

	// Action output without the namespace
	Output *ActionVpsConfigShowOutput
}

// Prepare the action for invocation
func (action *ActionVpsConfigShow) Prepare() *ActionVpsConfigShowInvocation {
	return &ActionVpsConfigShowInvocation{
		Action: action,
		Path:   "/v6.0/vps_configs/{vps_config_id}",
	}
}

// ActionVpsConfigShowInvocation is used to configure action for invocation
type ActionVpsConfigShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsConfigShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsConfigShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsConfigShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsConfigShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsConfigShowInvocation) SetPathParamString(param string, value string) *ActionVpsConfigShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsConfigShowInvocation) NewMetaInput() *ActionVpsConfigShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsConfigShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsConfigShowInvocation) SetMetaInput(input *ActionVpsConfigShowMetaGlobalInput) *ActionVpsConfigShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsConfigShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsConfigShowInvocation) Call() (*ActionVpsConfigShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsConfigShowInvocation) callAsQuery() (*ActionVpsConfigShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsConfigShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsConfig
	}
	return resp, err
}

func (inv *ActionVpsConfigShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
