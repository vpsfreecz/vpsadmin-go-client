package client

import (
	"strings"
)

// ActionOutageUpdate is a type for action Outage#Update
type ActionOutageUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageUpdate(client *Client) *ActionOutageUpdate {
	return &ActionOutageUpdate{
		Client: client,
	}
}

// ActionOutageUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionOutageUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageUpdateMetaGlobalInput) SetIncludes(value string) *ActionOutageUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageUpdateMetaGlobalInput) SetNo(value bool) *ActionOutageUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionOutageUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageUpdateInput is a type for action input parameters
type ActionOutageUpdateInput struct {
	AutoResolve   bool   `json:"auto_resolve"`
	BeginsAt      string `json:"begins_at"`
	CsDescription string `json:"cs_description"`
	CsSummary     string `json:"cs_summary"`
	Duration      int64  `json:"duration"`
	EnDescription string `json:"en_description"`
	EnSummary     string `json:"en_summary"`
	FinishedAt    string `json:"finished_at"`
	Impact        string `json:"impact"`
	Type          string `json:"type"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAutoResolve sets parameter AutoResolve to value and selects it for sending
func (in *ActionOutageUpdateInput) SetAutoResolve(value bool) *ActionOutageUpdateInput {
	in.AutoResolve = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AutoResolve"] = nil
	return in
}

// SetBeginsAt sets parameter BeginsAt to value and selects it for sending
func (in *ActionOutageUpdateInput) SetBeginsAt(value string) *ActionOutageUpdateInput {
	in.BeginsAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["BeginsAt"] = nil
	return in
}

// SetCsDescription sets parameter CsDescription to value and selects it for sending
func (in *ActionOutageUpdateInput) SetCsDescription(value string) *ActionOutageUpdateInput {
	in.CsDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsDescription"] = nil
	return in
}

// SetCsSummary sets parameter CsSummary to value and selects it for sending
func (in *ActionOutageUpdateInput) SetCsSummary(value string) *ActionOutageUpdateInput {
	in.CsSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsSummary"] = nil
	return in
}

// SetDuration sets parameter Duration to value and selects it for sending
func (in *ActionOutageUpdateInput) SetDuration(value int64) *ActionOutageUpdateInput {
	in.Duration = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Duration"] = nil
	return in
}

// SetEnDescription sets parameter EnDescription to value and selects it for sending
func (in *ActionOutageUpdateInput) SetEnDescription(value string) *ActionOutageUpdateInput {
	in.EnDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnDescription"] = nil
	return in
}

// SetEnSummary sets parameter EnSummary to value and selects it for sending
func (in *ActionOutageUpdateInput) SetEnSummary(value string) *ActionOutageUpdateInput {
	in.EnSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnSummary"] = nil
	return in
}

// SetFinishedAt sets parameter FinishedAt to value and selects it for sending
func (in *ActionOutageUpdateInput) SetFinishedAt(value string) *ActionOutageUpdateInput {
	in.FinishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FinishedAt"] = nil
	return in
}

// SetImpact sets parameter Impact to value and selects it for sending
func (in *ActionOutageUpdateInput) SetImpact(value string) *ActionOutageUpdateInput {
	in.Impact = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Impact"] = nil
	return in
}

// SetType sets parameter Type to value and selects it for sending
func (in *ActionOutageUpdateInput) SetType(value string) *ActionOutageUpdateInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageUpdateInput) SelectParameters(params ...string) *ActionOutageUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOutageUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOutageUpdateInput) UnselectParameters(params ...string) *ActionOutageUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOutageUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageUpdateRequest is a type for the entire action request
type ActionOutageUpdateRequest struct {
	Outage map[string]interface{} `json:"outage"`
	Meta   map[string]interface{} `json:"_meta"`
}

// ActionOutageUpdateOutput is a type for action output parameters
type ActionOutageUpdateOutput struct {
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
type ActionOutageUpdateResponse struct {
	Action *ActionOutageUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Outage *ActionOutageUpdateOutput `json:"outage"`
	}

	// Action output without the namespace
	Output *ActionOutageUpdateOutput
}

// Prepare the action for invocation
func (action *ActionOutageUpdate) Prepare() *ActionOutageUpdateInvocation {
	return &ActionOutageUpdateInvocation{
		Action: action,
		Path:   "/v7.0/outages/{outage_id}",
	}
}

// ActionOutageUpdateInvocation is used to configure action for invocation
type ActionOutageUpdateInvocation struct {
	// Pointer to the action
	Action *ActionOutageUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageUpdateInput
	// Global meta input parameters
	MetaInput *ActionOutageUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageUpdateInvocation) SetPathParamInt(param string, value int64) *ActionOutageUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageUpdateInvocation) SetPathParamString(param string, value string) *ActionOutageUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageUpdateInvocation) NewInput() *ActionOutageUpdateInput {
	inv.Input = &ActionOutageUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageUpdateInvocation) SetInput(input *ActionOutageUpdateInput) *ActionOutageUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOutageUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageUpdateInvocation) NewMetaInput() *ActionOutageUpdateMetaGlobalInput {
	inv.MetaInput = &ActionOutageUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageUpdateInvocation) SetMetaInput(input *ActionOutageUpdateMetaGlobalInput) *ActionOutageUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageUpdateInvocation) Call() (*ActionOutageUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOutageUpdateInvocation) callAsBody() (*ActionOutageUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Outage
	}
	return resp, err
}

func (inv *ActionOutageUpdateInvocation) makeAllInputParams() *ActionOutageUpdateRequest {
	return &ActionOutageUpdateRequest{
		Outage: inv.makeInputParams(),
		Meta:   inv.makeMetaInputParams(),
	}
}

func (inv *ActionOutageUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AutoResolve") {
			ret["auto_resolve"] = inv.Input.AutoResolve
		}
		if inv.IsParameterSelected("BeginsAt") {
			ret["begins_at"] = inv.Input.BeginsAt
		}
		if inv.IsParameterSelected("CsDescription") {
			ret["cs_description"] = inv.Input.CsDescription
		}
		if inv.IsParameterSelected("CsSummary") {
			ret["cs_summary"] = inv.Input.CsSummary
		}
		if inv.IsParameterSelected("Duration") {
			ret["duration"] = inv.Input.Duration
		}
		if inv.IsParameterSelected("EnDescription") {
			ret["en_description"] = inv.Input.EnDescription
		}
		if inv.IsParameterSelected("EnSummary") {
			ret["en_summary"] = inv.Input.EnSummary
		}
		if inv.IsParameterSelected("FinishedAt") {
			ret["finished_at"] = inv.Input.FinishedAt
		}
		if inv.IsParameterSelected("Impact") {
			ret["impact"] = inv.Input.Impact
		}
		if inv.IsParameterSelected("Type") {
			ret["type"] = inv.Input.Type
		}
	}

	return ret
}

func (inv *ActionOutageUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
