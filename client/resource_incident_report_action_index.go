package client

import ()

// ActionIncidentReportIndex is a type for action Incident_report#Index
type ActionIncidentReportIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIncidentReportIndex(client *Client) *ActionIncidentReportIndex {
	return &ActionIncidentReportIndex{
		Client: client,
	}
}

// ActionIncidentReportIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIncidentReportIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIncidentReportIndexMetaGlobalInput) SetCount(value bool) *ActionIncidentReportIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIncidentReportIndexMetaGlobalInput) SetIncludes(value string) *ActionIncidentReportIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIncidentReportIndexMetaGlobalInput) SetNo(value bool) *ActionIncidentReportIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIncidentReportIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIncidentReportIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIncidentReportIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportIndexInput is a type for action input parameters
type ActionIncidentReportIndexInput struct {
	Codename            string `json:"codename"`
	FiledBy             int64  `json:"filed_by"`
	FromId              int64  `json:"from_id"`
	IpAddr              string `json:"ip_addr"`
	IpAddressAssignment int64  `json:"ip_address_assignment"`
	Limit               int64  `json:"limit"`
	Mailbox             int64  `json:"mailbox"`
	User                int64  `json:"user"`
	Vps                 int64  `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCodename sets parameter Codename to value and selects it for sending
func (in *ActionIncidentReportIndexInput) SetCodename(value string) *ActionIncidentReportIndexInput {
	in.Codename = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Codename"] = nil
	return in
}

// SetFiledBy sets parameter FiledBy to value and selects it for sending
func (in *ActionIncidentReportIndexInput) SetFiledBy(value int64) *ActionIncidentReportIndexInput {
	in.FiledBy = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetFiledByNil(false)
	in._selectedParameters["FiledBy"] = nil
	return in
}

// SetFiledByNil sets parameter FiledBy to nil and selects it for sending
func (in *ActionIncidentReportIndexInput) SetFiledByNil(set bool) *ActionIncidentReportIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["FiledBy"] = nil
		in.SelectParameters("FiledBy")
	} else {
		delete(in._nilParameters, "FiledBy")
	}
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIncidentReportIndexInput) SetFromId(value int64) *ActionIncidentReportIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetIpAddr sets parameter IpAddr to value and selects it for sending
func (in *ActionIncidentReportIndexInput) SetIpAddr(value string) *ActionIncidentReportIndexInput {
	in.IpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpAddr"] = nil
	return in
}

// SetIpAddressAssignment sets parameter IpAddressAssignment to value and selects it for sending
func (in *ActionIncidentReportIndexInput) SetIpAddressAssignment(value int64) *ActionIncidentReportIndexInput {
	in.IpAddressAssignment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetIpAddressAssignmentNil(false)
	in._selectedParameters["IpAddressAssignment"] = nil
	return in
}

// SetIpAddressAssignmentNil sets parameter IpAddressAssignment to nil and selects it for sending
func (in *ActionIncidentReportIndexInput) SetIpAddressAssignmentNil(set bool) *ActionIncidentReportIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["IpAddressAssignment"] = nil
		in.SelectParameters("IpAddressAssignment")
	} else {
		delete(in._nilParameters, "IpAddressAssignment")
	}
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIncidentReportIndexInput) SetLimit(value int64) *ActionIncidentReportIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetMailbox sets parameter Mailbox to value and selects it for sending
func (in *ActionIncidentReportIndexInput) SetMailbox(value int64) *ActionIncidentReportIndexInput {
	in.Mailbox = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetMailboxNil(false)
	in._selectedParameters["Mailbox"] = nil
	return in
}

// SetMailboxNil sets parameter Mailbox to nil and selects it for sending
func (in *ActionIncidentReportIndexInput) SetMailboxNil(set bool) *ActionIncidentReportIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Mailbox"] = nil
		in.SelectParameters("Mailbox")
	} else {
		delete(in._nilParameters, "Mailbox")
	}
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionIncidentReportIndexInput) SetUser(value int64) *ActionIncidentReportIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionIncidentReportIndexInput) SetUserNil(set bool) *ActionIncidentReportIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionIncidentReportIndexInput) SetVps(value int64) *ActionIncidentReportIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionIncidentReportIndexInput) SetVpsNil(set bool) *ActionIncidentReportIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Vps"] = nil
		in.SelectParameters("Vps")
	} else {
		delete(in._nilParameters, "Vps")
	}
	return in
}

// SelectParameters sets parameters from ActionIncidentReportIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIncidentReportIndexInput) SelectParameters(params ...string) *ActionIncidentReportIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIncidentReportIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIncidentReportIndexInput) UnselectParameters(params ...string) *ActionIncidentReportIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIncidentReportIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIncidentReportIndexOutput is a type for action output parameters
type ActionIncidentReportIndexOutput struct {
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
	VpsAction           string                               `json:"vps_action"`
}

// Type for action response, including envelope
type ActionIncidentReportIndexResponse struct {
	Action *ActionIncidentReportIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IncidentReports []*ActionIncidentReportIndexOutput `json:"incident_reports"`
	}

	// Action output without the namespace
	Output []*ActionIncidentReportIndexOutput
}

// Prepare the action for invocation
func (action *ActionIncidentReportIndex) Prepare() *ActionIncidentReportIndexInvocation {
	return &ActionIncidentReportIndexInvocation{
		Action: action,
		Path:   "/v7.0/incident_reports",
	}
}

// ActionIncidentReportIndexInvocation is used to configure action for invocation
type ActionIncidentReportIndexInvocation struct {
	// Pointer to the action
	Action *ActionIncidentReportIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIncidentReportIndexInput
	// Global meta input parameters
	MetaInput *ActionIncidentReportIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIncidentReportIndexInvocation) NewInput() *ActionIncidentReportIndexInput {
	inv.Input = &ActionIncidentReportIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIncidentReportIndexInvocation) SetInput(input *ActionIncidentReportIndexInput) *ActionIncidentReportIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIncidentReportIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIncidentReportIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIncidentReportIndexInvocation) NewMetaInput() *ActionIncidentReportIndexMetaGlobalInput {
	inv.MetaInput = &ActionIncidentReportIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIncidentReportIndexInvocation) SetMetaInput(input *ActionIncidentReportIndexMetaGlobalInput) *ActionIncidentReportIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIncidentReportIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIncidentReportIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIncidentReportIndexInvocation) Call() (*ActionIncidentReportIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIncidentReportIndexInvocation) callAsQuery() (*ActionIncidentReportIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIncidentReportIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IncidentReports
	}
	return resp, err
}

func (inv *ActionIncidentReportIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Codename") {
			ret["incident_report[codename]"] = inv.Input.Codename
		}
		if inv.IsParameterSelected("FiledBy") {
			ret["incident_report[filed_by]"] = convertInt64ToString(inv.Input.FiledBy)
		}
		if inv.IsParameterSelected("FromId") {
			ret["incident_report[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("IpAddr") {
			ret["incident_report[ip_addr]"] = inv.Input.IpAddr
		}
		if inv.IsParameterSelected("IpAddressAssignment") {
			ret["incident_report[ip_address_assignment]"] = convertInt64ToString(inv.Input.IpAddressAssignment)
		}
		if inv.IsParameterSelected("Limit") {
			ret["incident_report[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Mailbox") {
			ret["incident_report[mailbox]"] = convertInt64ToString(inv.Input.Mailbox)
		}
		if inv.IsParameterSelected("User") {
			ret["incident_report[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["incident_report[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionIncidentReportIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
