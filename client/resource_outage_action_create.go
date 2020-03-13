package client

import (
)

// ActionOutageCreate is a type for action Outage#Create
type ActionOutageCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageCreate(client *Client) *ActionOutageCreate {
	return &ActionOutageCreate{
		Client: client,
	}
}

// ActionOutageCreateMetaGlobalInput is a type for action global meta input parameters
type ActionOutageCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageCreateMetaGlobalInput) SetNo(value bool) *ActionOutageCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageCreateMetaGlobalInput) SetIncludes(value string) *ActionOutageCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageCreateMetaGlobalInput) SelectParameters(params ...string) *ActionOutageCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageCreateInput is a type for action input parameters
type ActionOutageCreateInput struct {
	BeginsAt string `json:"begins_at"`
	FinishedAt string `json:"finished_at"`
	Duration int64 `json:"duration"`
	Planned bool `json:"planned"`
	Type string `json:"type"`
	EnSummary string `json:"en_summary"`
	EnDescription string `json:"en_description"`
	CsSummary string `json:"cs_summary"`
	CsDescription string `json:"cs_description"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetBeginsAt sets parameter BeginsAt to value and selects it for sending
func (in *ActionOutageCreateInput) SetBeginsAt(value string) *ActionOutageCreateInput {
	in.BeginsAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["BeginsAt"] = nil
	return in
}
// SetFinishedAt sets parameter FinishedAt to value and selects it for sending
func (in *ActionOutageCreateInput) SetFinishedAt(value string) *ActionOutageCreateInput {
	in.FinishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FinishedAt"] = nil
	return in
}
// SetDuration sets parameter Duration to value and selects it for sending
func (in *ActionOutageCreateInput) SetDuration(value int64) *ActionOutageCreateInput {
	in.Duration = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Duration"] = nil
	return in
}
// SetPlanned sets parameter Planned to value and selects it for sending
func (in *ActionOutageCreateInput) SetPlanned(value bool) *ActionOutageCreateInput {
	in.Planned = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Planned"] = nil
	return in
}
// SetType sets parameter Type to value and selects it for sending
func (in *ActionOutageCreateInput) SetType(value string) *ActionOutageCreateInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}
// SetEnSummary sets parameter EnSummary to value and selects it for sending
func (in *ActionOutageCreateInput) SetEnSummary(value string) *ActionOutageCreateInput {
	in.EnSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnSummary"] = nil
	return in
}
// SetEnDescription sets parameter EnDescription to value and selects it for sending
func (in *ActionOutageCreateInput) SetEnDescription(value string) *ActionOutageCreateInput {
	in.EnDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnDescription"] = nil
	return in
}
// SetCsSummary sets parameter CsSummary to value and selects it for sending
func (in *ActionOutageCreateInput) SetCsSummary(value string) *ActionOutageCreateInput {
	in.CsSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsSummary"] = nil
	return in
}
// SetCsDescription sets parameter CsDescription to value and selects it for sending
func (in *ActionOutageCreateInput) SetCsDescription(value string) *ActionOutageCreateInput {
	in.CsDescription = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsDescription"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageCreateInput) SelectParameters(params ...string) *ActionOutageCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageCreateRequest is a type for the entire action request
type ActionOutageCreateRequest struct {
	Outage map[string]interface{} `json:"outage"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionOutageCreateOutput is a type for action output parameters
type ActionOutageCreateOutput struct {
	Id int64 `json:"id"`
	BeginsAt string `json:"begins_at"`
	FinishedAt string `json:"finished_at"`
	Duration int64 `json:"duration"`
	Planned bool `json:"planned"`
	State string `json:"state"`
	Type string `json:"type"`
	EnSummary string `json:"en_summary"`
	EnDescription string `json:"en_description"`
	CsSummary string `json:"cs_summary"`
	CsDescription string `json:"cs_description"`
	Affected bool `json:"affected"`
	AffectedUserCount int64 `json:"affected_user_count"`
	AffectedDirectVpsCount int64 `json:"affected_direct_vps_count"`
	AffectedIndirectVpsCount int64 `json:"affected_indirect_vps_count"`
}


// Type for action response, including envelope
type ActionOutageCreateResponse struct {
	Action *ActionOutageCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Outage *ActionOutageCreateOutput `json:"outage"`
	}

	// Action output without the namespace
	Output *ActionOutageCreateOutput
}


// Prepare the action for invocation
func (action *ActionOutageCreate) Prepare() *ActionOutageCreateInvocation {
	return &ActionOutageCreateInvocation{
		Action: action,
		Path: "/v6.0/outages",
	}
}

// ActionOutageCreateInvocation is used to configure action for invocation
type ActionOutageCreateInvocation struct {
	// Pointer to the action
	Action *ActionOutageCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageCreateInput
	// Global meta input parameters
	MetaInput *ActionOutageCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageCreateInvocation) NewInput() *ActionOutageCreateInput {
	inv.Input = &ActionOutageCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageCreateInvocation) SetInput(input *ActionOutageCreateInput) *ActionOutageCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageCreateInvocation) NewMetaInput() *ActionOutageCreateMetaGlobalInput {
	inv.MetaInput = &ActionOutageCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageCreateInvocation) SetMetaInput(input *ActionOutageCreateMetaGlobalInput) *ActionOutageCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageCreateInvocation) Call() (*ActionOutageCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionOutageCreateInvocation) callAsBody() (*ActionOutageCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Outage
	}
	return resp, err
}




func (inv *ActionOutageCreateInvocation) makeAllInputParams() *ActionOutageCreateRequest {
	return &ActionOutageCreateRequest{
		Outage: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOutageCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("BeginsAt") {
			ret["begins_at"] = inv.Input.BeginsAt
		}
		if inv.IsParameterSelected("FinishedAt") {
			ret["finished_at"] = inv.Input.FinishedAt
		}
		if inv.IsParameterSelected("Duration") {
			ret["duration"] = inv.Input.Duration
		}
		if inv.IsParameterSelected("Planned") {
			ret["planned"] = inv.Input.Planned
		}
		if inv.IsParameterSelected("Type") {
			ret["type"] = inv.Input.Type
		}
		if inv.IsParameterSelected("EnSummary") {
			ret["en_summary"] = inv.Input.EnSummary
		}
		if inv.IsParameterSelected("EnDescription") {
			ret["en_description"] = inv.Input.EnDescription
		}
		if inv.IsParameterSelected("CsSummary") {
			ret["cs_summary"] = inv.Input.CsSummary
		}
		if inv.IsParameterSelected("CsDescription") {
			ret["cs_description"] = inv.Input.CsDescription
		}
	}

	return ret
}

func (inv *ActionOutageCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
