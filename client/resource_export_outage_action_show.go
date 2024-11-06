package client

import (
	"strings"
)

// ActionExportOutageShow is a type for action Export_outage#Show
type ActionExportOutageShow struct {
	// Pointer to client
	Client *Client
}

func NewActionExportOutageShow(client *Client) *ActionExportOutageShow {
	return &ActionExportOutageShow{
		Client: client,
	}
}

// ActionExportOutageShowMetaGlobalInput is a type for action global meta input parameters
type ActionExportOutageShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportOutageShowMetaGlobalInput) SetIncludes(value string) *ActionExportOutageShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportOutageShowMetaGlobalInput) SetNo(value bool) *ActionExportOutageShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportOutageShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportOutageShowMetaGlobalInput) SelectParameters(params ...string) *ActionExportOutageShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportOutageShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportOutageShowOutput is a type for action output parameters
type ActionExportOutageShowOutput struct {
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	Export      *ActionExportShowOutput      `json:"export"`
	Id          int64                        `json:"id"`
	Location    *ActionLocationShowOutput    `json:"location"`
	Node        *ActionNodeShowOutput        `json:"node"`
	Outage      *ActionOutageShowOutput      `json:"outage"`
	User        *ActionUserShowOutput        `json:"user"`
}

// Type for action response, including envelope
type ActionExportOutageShowResponse struct {
	Action *ActionExportOutageShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ExportOutage *ActionExportOutageShowOutput `json:"export_outage"`
	}

	// Action output without the namespace
	Output *ActionExportOutageShowOutput
}

// Prepare the action for invocation
func (action *ActionExportOutageShow) Prepare() *ActionExportOutageShowInvocation {
	return &ActionExportOutageShowInvocation{
		Action: action,
		Path:   "/v7.0/export_outages/{export_outage_id}",
	}
}

// ActionExportOutageShowInvocation is used to configure action for invocation
type ActionExportOutageShowInvocation struct {
	// Pointer to the action
	Action *ActionExportOutageShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionExportOutageShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionExportOutageShowInvocation) SetPathParamInt(param string, value int64) *ActionExportOutageShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionExportOutageShowInvocation) SetPathParamString(param string, value string) *ActionExportOutageShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportOutageShowInvocation) NewMetaInput() *ActionExportOutageShowMetaGlobalInput {
	inv.MetaInput = &ActionExportOutageShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportOutageShowInvocation) SetMetaInput(input *ActionExportOutageShowMetaGlobalInput) *ActionExportOutageShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportOutageShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionExportOutageShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionExportOutageShowInvocation) Call() (*ActionExportOutageShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionExportOutageShowInvocation) callAsQuery() (*ActionExportOutageShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionExportOutageShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ExportOutage
	}
	return resp, err
}

func (inv *ActionExportOutageShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
