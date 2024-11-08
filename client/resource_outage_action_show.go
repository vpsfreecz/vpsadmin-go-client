package client

import (
	"strings"
)

// ActionOutageShow is a type for action Outage#Show
type ActionOutageShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageShow(client *Client) *ActionOutageShow {
	return &ActionOutageShow{
		Client: client,
	}
}

// ActionOutageShowMetaGlobalInput is a type for action global meta input parameters
type ActionOutageShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageShowMetaGlobalInput) SetIncludes(value string) *ActionOutageShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageShowMetaGlobalInput) SetNo(value bool) *ActionOutageShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageShowMetaGlobalInput) SelectParameters(params ...string) *ActionOutageShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageShowOutput is a type for action output parameters
type ActionOutageShowOutput struct {
	Affected                 bool   `json:"affected"`
	AffectedDirectVpsCount   int64  `json:"affected_direct_vps_count"`
	AffectedExportCount      int64  `json:"affected_export_count"`
	AffectedIndirectVpsCount int64  `json:"affected_indirect_vps_count"`
	AffectedUserCount        int64  `json:"affected_user_count"`
	AutoResolve              bool   `json:"auto_resolve"`
	BeginsAt                 string `json:"begins_at"`
	CsDescription            string `json:"cs_description"`
	CsSummary                string `json:"cs_summary"`
	Duration                 int64  `json:"duration"`
	EnDescription            string `json:"en_description"`
	EnSummary                string `json:"en_summary"`
	FinishedAt               string `json:"finished_at"`
	Id                       int64  `json:"id"`
	Impact                   string `json:"impact"`
	State                    string `json:"state"`
	Type                     string `json:"type"`
}

// Type for action response, including envelope
type ActionOutageShowResponse struct {
	Action *ActionOutageShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Outage *ActionOutageShowOutput `json:"outage"`
	}

	// Action output without the namespace
	Output *ActionOutageShowOutput
}

// Prepare the action for invocation
func (action *ActionOutageShow) Prepare() *ActionOutageShowInvocation {
	return &ActionOutageShowInvocation{
		Action: action,
		Path:   "/v7.0/outages/{outage_id}",
	}
}

// ActionOutageShowInvocation is used to configure action for invocation
type ActionOutageShowInvocation struct {
	// Pointer to the action
	Action *ActionOutageShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOutageShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageShowInvocation) SetPathParamInt(param string, value int64) *ActionOutageShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageShowInvocation) SetPathParamString(param string, value string) *ActionOutageShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageShowInvocation) NewMetaInput() *ActionOutageShowMetaGlobalInput {
	inv.MetaInput = &ActionOutageShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageShowInvocation) SetMetaInput(input *ActionOutageShowMetaGlobalInput) *ActionOutageShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageShowInvocation) Call() (*ActionOutageShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOutageShowInvocation) callAsQuery() (*ActionOutageShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOutageShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Outage
	}
	return resp, err
}

func (inv *ActionOutageShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
