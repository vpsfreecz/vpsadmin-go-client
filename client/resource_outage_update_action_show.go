package client

import (
	"strings"
)

// ActionOutageUpdateShow is a type for action Outage_update#Show
type ActionOutageUpdateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageUpdateShow(client *Client) *ActionOutageUpdateShow {
	return &ActionOutageUpdateShow{
		Client: client,
	}
}

// ActionOutageUpdateShowMetaGlobalInput is a type for action global meta input parameters
type ActionOutageUpdateShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageUpdateShowMetaGlobalInput) SetIncludes(value string) *ActionOutageUpdateShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageUpdateShowMetaGlobalInput) SetNo(value bool) *ActionOutageUpdateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageUpdateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageUpdateShowMetaGlobalInput) SelectParameters(params ...string) *ActionOutageUpdateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageUpdateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageUpdateShowOutput is a type for action output parameters
type ActionOutageUpdateShowOutput struct {
	BeginsAt      string                  `json:"begins_at"`
	CreatedAt     string                  `json:"created_at"`
	CsDescription string                  `json:"cs_description"`
	CsSummary     string                  `json:"cs_summary"`
	Duration      int64                   `json:"duration"`
	EnDescription string                  `json:"en_description"`
	EnSummary     string                  `json:"en_summary"`
	FinishedAt    string                  `json:"finished_at"`
	Id            int64                   `json:"id"`
	Outage        *ActionOutageShowOutput `json:"outage"`
	ReportedBy    *ActionUserShowOutput   `json:"reported_by"`
	ReporterName  string                  `json:"reporter_name"`
	State         string                  `json:"state"`
	Type          string                  `json:"type"`
}

// Type for action response, including envelope
type ActionOutageUpdateShowResponse struct {
	Action *ActionOutageUpdateShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageUpdate *ActionOutageUpdateShowOutput `json:"outage_update"`
	}

	// Action output without the namespace
	Output *ActionOutageUpdateShowOutput
}

// Prepare the action for invocation
func (action *ActionOutageUpdateShow) Prepare() *ActionOutageUpdateShowInvocation {
	return &ActionOutageUpdateShowInvocation{
		Action: action,
		Path:   "/v6.0/outage_updates/{outage_update_id}",
	}
}

// ActionOutageUpdateShowInvocation is used to configure action for invocation
type ActionOutageUpdateShowInvocation struct {
	// Pointer to the action
	Action *ActionOutageUpdateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOutageUpdateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageUpdateShowInvocation) SetPathParamInt(param string, value int64) *ActionOutageUpdateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageUpdateShowInvocation) SetPathParamString(param string, value string) *ActionOutageUpdateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageUpdateShowInvocation) NewMetaInput() *ActionOutageUpdateShowMetaGlobalInput {
	inv.MetaInput = &ActionOutageUpdateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageUpdateShowInvocation) SetMetaInput(input *ActionOutageUpdateShowMetaGlobalInput) *ActionOutageUpdateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageUpdateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageUpdateShowInvocation) Call() (*ActionOutageUpdateShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOutageUpdateShowInvocation) callAsQuery() (*ActionOutageUpdateShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOutageUpdateShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageUpdate
	}
	return resp, err
}

func (inv *ActionOutageUpdateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
