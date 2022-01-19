package client

import (
	"strings"
)

// ActionUserOutageShow is a type for action User_outage#Show
type ActionUserOutageShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserOutageShow(client *Client) *ActionUserOutageShow {
	return &ActionUserOutageShow{
		Client: client,
	}
}

// ActionUserOutageShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserOutageShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserOutageShowMetaGlobalInput) SetIncludes(value string) *ActionUserOutageShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserOutageShowMetaGlobalInput) SetNo(value bool) *ActionUserOutageShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserOutageShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserOutageShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserOutageShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserOutageShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserOutageShowOutput is a type for action output parameters
type ActionUserOutageShowOutput struct {
	ExportCount int64                   `json:"export_count"`
	Id          int64                   `json:"id"`
	Outage      *ActionOutageShowOutput `json:"outage"`
	User        *ActionUserShowOutput   `json:"user"`
	VpsCount    int64                   `json:"vps_count"`
}

// Type for action response, including envelope
type ActionUserOutageShowResponse struct {
	Action *ActionUserOutageShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserOutage *ActionUserOutageShowOutput `json:"user_outage"`
	}

	// Action output without the namespace
	Output *ActionUserOutageShowOutput
}

// Prepare the action for invocation
func (action *ActionUserOutageShow) Prepare() *ActionUserOutageShowInvocation {
	return &ActionUserOutageShowInvocation{
		Action: action,
		Path:   "/v6.0/user_outages/{user_outage_id}",
	}
}

// ActionUserOutageShowInvocation is used to configure action for invocation
type ActionUserOutageShowInvocation struct {
	// Pointer to the action
	Action *ActionUserOutageShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserOutageShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserOutageShowInvocation) SetPathParamInt(param string, value int64) *ActionUserOutageShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserOutageShowInvocation) SetPathParamString(param string, value string) *ActionUserOutageShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserOutageShowInvocation) NewMetaInput() *ActionUserOutageShowMetaGlobalInput {
	inv.MetaInput = &ActionUserOutageShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserOutageShowInvocation) SetMetaInput(input *ActionUserOutageShowMetaGlobalInput) *ActionUserOutageShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserOutageShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserOutageShowInvocation) Call() (*ActionUserOutageShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserOutageShowInvocation) callAsQuery() (*ActionUserOutageShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserOutageShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserOutage
	}
	return resp, err
}

func (inv *ActionUserOutageShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
