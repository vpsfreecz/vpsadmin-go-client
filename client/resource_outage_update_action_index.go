package client

import ()

// ActionOutageUpdateIndex is a type for action Outage_update#Index
type ActionOutageUpdateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageUpdateIndex(client *Client) *ActionOutageUpdateIndex {
	return &ActionOutageUpdateIndex{
		Client: client,
	}
}

// ActionOutageUpdateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOutageUpdateIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOutageUpdateIndexMetaGlobalInput) SetCount(value bool) *ActionOutageUpdateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageUpdateIndexMetaGlobalInput) SetIncludes(value string) *ActionOutageUpdateIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageUpdateIndexMetaGlobalInput) SetNo(value bool) *ActionOutageUpdateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageUpdateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageUpdateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOutageUpdateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageUpdateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageUpdateIndexInput is a type for action input parameters
type ActionOutageUpdateIndexInput struct {
	Limit      int64  `json:"limit"`
	Offset     int64  `json:"offset"`
	Outage     int64  `json:"outage"`
	ReportedBy int64  `json:"reported_by"`
	Since      string `json:"since"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOutageUpdateIndexInput) SetLimit(value int64) *ActionOutageUpdateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionOutageUpdateIndexInput) SetOffset(value int64) *ActionOutageUpdateIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetOutage sets parameter Outage to value and selects it for sending
func (in *ActionOutageUpdateIndexInput) SetOutage(value int64) *ActionOutageUpdateIndexInput {
	in.Outage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Outage"] = nil
	return in
}

// SetReportedBy sets parameter ReportedBy to value and selects it for sending
func (in *ActionOutageUpdateIndexInput) SetReportedBy(value int64) *ActionOutageUpdateIndexInput {
	in.ReportedBy = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ReportedBy"] = nil
	return in
}

// SetSince sets parameter Since to value and selects it for sending
func (in *ActionOutageUpdateIndexInput) SetSince(value string) *ActionOutageUpdateIndexInput {
	in.Since = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Since"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageUpdateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageUpdateIndexInput) SelectParameters(params ...string) *ActionOutageUpdateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageUpdateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageUpdateIndexOutput is a type for action output parameters
type ActionOutageUpdateIndexOutput struct {
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
type ActionOutageUpdateIndexResponse struct {
	Action *ActionOutageUpdateIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageUpdates []*ActionOutageUpdateIndexOutput `json:"outage_updates"`
	}

	// Action output without the namespace
	Output []*ActionOutageUpdateIndexOutput
}

// Prepare the action for invocation
func (action *ActionOutageUpdateIndex) Prepare() *ActionOutageUpdateIndexInvocation {
	return &ActionOutageUpdateIndexInvocation{
		Action: action,
		Path:   "/v6.0/outage_updates",
	}
}

// ActionOutageUpdateIndexInvocation is used to configure action for invocation
type ActionOutageUpdateIndexInvocation struct {
	// Pointer to the action
	Action *ActionOutageUpdateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageUpdateIndexInput
	// Global meta input parameters
	MetaInput *ActionOutageUpdateIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageUpdateIndexInvocation) NewInput() *ActionOutageUpdateIndexInput {
	inv.Input = &ActionOutageUpdateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageUpdateIndexInvocation) SetInput(input *ActionOutageUpdateIndexInput) *ActionOutageUpdateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageUpdateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageUpdateIndexInvocation) NewMetaInput() *ActionOutageUpdateIndexMetaGlobalInput {
	inv.MetaInput = &ActionOutageUpdateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageUpdateIndexInvocation) SetMetaInput(input *ActionOutageUpdateIndexMetaGlobalInput) *ActionOutageUpdateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageUpdateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageUpdateIndexInvocation) Call() (*ActionOutageUpdateIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOutageUpdateIndexInvocation) callAsQuery() (*ActionOutageUpdateIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOutageUpdateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageUpdates
	}
	return resp, err
}

func (inv *ActionOutageUpdateIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["outage_update[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["outage_update[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Outage") {
			ret["outage_update[outage]"] = convertInt64ToString(inv.Input.Outage)
		}
		if inv.IsParameterSelected("ReportedBy") {
			ret["outage_update[reported_by]"] = convertInt64ToString(inv.Input.ReportedBy)
		}
		if inv.IsParameterSelected("Since") {
			ret["outage_update[since]"] = inv.Input.Since
		}
	}
}

func (inv *ActionOutageUpdateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
