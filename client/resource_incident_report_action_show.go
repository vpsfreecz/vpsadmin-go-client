package client

import (
	"strings"
)

// ActionIncidentReportShow is a type for action Incident_report#Show
type ActionIncidentReportShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentReportShow(client *Client) *ActionIncidentReportShow {
	return &ActionIncidentReportShow{
		Client: client,
	}
}

// ActionIncidentReportShowMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentReportShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIncidentReportShowMetaGlobalInput) SetIncludes(value string) *ActionIncidentReportShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentReportShowMetaGlobalInput) SetNo(value bool) *ActionIncidentReportShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentReportShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportShowMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentReportShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentReportShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportShowOutput is a type for action output parameters
type ActionIncidentReportShowOutput struct {
	Codename            string                               `json:"codename"`
	CpuLimit            int64                                `json:"cpu_limit"`
	CreatedAt           string                               `json:"created_at"`
	DetectedAt          string                               `json:"detected_at"`
	FiledBy             *ActionUserShowOutput                `json:"filed_by"`
	Id                  int64                                `json:"id"`
	IpAddressAssignment *ActionIpAddressAssignmentShowOutput `json:"ip_address_assignment"`
	Mailbox             *ActionMailboxShowOutput             `json:"mailbox"`
	RawUserId           int64                                `json:"raw_user_id"`
	RawVpsId            int64                                `json:"raw_vps_id"`
	ReportedAt          string                               `json:"reported_at"`
	Subject             string                               `json:"subject"`
	Text                string                               `json:"text"`
	User                *ActionUserShowOutput                `json:"user"`
	Vps                 *ActionVpsShowOutput                 `json:"vps"`
}

// Type for action response, including envelope
type ActionIncidentReportShowResponse struct {
	Action *ActionIncidentReportShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncidentReport *ActionIncidentReportShowOutput `json:"incident_report"`
	}

	// Action output without the namespace
	Output *ActionIncidentReportShowOutput
}

// Prepare the action for invocation
func (action *ActionIncidentReportShow) Prepare() *ActionIncidentReportShowInvocation {
	return &ActionIncidentReportShowInvocation{
		Action: action,
		Path:   "/v7.0/incident_reports/{incident_report_id}",
	}
}

// ActionIncidentReportShowInvocation is used to configure action for invocation
type ActionIncidentReportShowInvocation struct {
	// Pointer to the action
	Action *ActionIncidentReportShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIncidentReportShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIncidentReportShowInvocation) SetPathParamInt(param string, value int64) *ActionIncidentReportShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIncidentReportShowInvocation) SetPathParamString(param string, value string) *ActionIncidentReportShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentReportShowInvocation) NewMetaInput() *ActionIncidentReportShowMetaGlobalInput {
	inv.MetaInput = &ActionIncidentReportShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentReportShowInvocation) SetMetaInput(input *ActionIncidentReportShowMetaGlobalInput) *ActionIncidentReportShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentReportShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentReportShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncidentReportShowInvocation) Call() (*ActionIncidentReportShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIncidentReportShowInvocation) callAsQuery() (*ActionIncidentReportShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncidentReportShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncidentReport
	}
	return resp, err
}

func (inv *ActionIncidentReportShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
