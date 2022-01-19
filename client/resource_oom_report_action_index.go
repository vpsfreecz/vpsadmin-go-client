package client

import ()

// ActionOomReportIndex is a type for action Oom_report#Index
type ActionOomReportIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportIndex(client *Client) *ActionOomReportIndex {
	return &ActionOomReportIndex{
		Client: client,
	}
}

// ActionOomReportIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOomReportIndexMetaGlobalInput) SetCount(value bool) *ActionOomReportIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportIndexMetaGlobalInput) SetIncludes(value string) *ActionOomReportIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportIndexMetaGlobalInput) SetNo(value bool) *ActionOomReportIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportIndexInput is a type for action input parameters
type ActionOomReportIndexInput struct {
	Environment int64  `json:"environment"`
	Limit       int64  `json:"limit"`
	Location    int64  `json:"location"`
	Node        int64  `json:"node"`
	Offset      int64  `json:"offset"`
	Since       string `json:"since"`
	Until       string `json:"until"`
	User        int64  `json:"user"`
	Vps         int64  `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionOomReportIndexInput) SetEnvironment(value int64) *ActionOomReportIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOomReportIndexInput) SetLimit(value int64) *ActionOomReportIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionOomReportIndexInput) SetLocation(value int64) *ActionOomReportIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionOomReportIndexInput) SetNode(value int64) *ActionOomReportIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionOomReportIndexInput) SetOffset(value int64) *ActionOomReportIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetSince sets parameter Since to value and selects it for sending
func (in *ActionOomReportIndexInput) SetSince(value string) *ActionOomReportIndexInput {
	in.Since = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Since"] = nil
	return in
}

// SetUntil sets parameter Until to value and selects it for sending
func (in *ActionOomReportIndexInput) SetUntil(value string) *ActionOomReportIndexInput {
	in.Until = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Until"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionOomReportIndexInput) SetUser(value int64) *ActionOomReportIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionOomReportIndexInput) SetVps(value int64) *ActionOomReportIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportIndexInput) SelectParameters(params ...string) *ActionOomReportIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportIndexOutput is a type for action output parameters
type ActionOomReportIndexOutput struct {
	CreatedAt     string               `json:"created_at"`
	Id            int64                `json:"id"`
	InvokedByName string               `json:"invoked_by_name"`
	InvokedByPid  int64                `json:"invoked_by_pid"`
	KilledName    string               `json:"killed_name"`
	KilledPid     int64                `json:"killed_pid"`
	ReportedAt    string               `json:"reported_at"`
	Vps           *ActionVpsShowOutput `json:"vps"`
}

// Type for action response, including envelope
type ActionOomReportIndexResponse struct {
	Action *ActionOomReportIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OomReports []*ActionOomReportIndexOutput `json:"oom_reports"`
	}

	// Action output without the namespace
	Output []*ActionOomReportIndexOutput
}

// Prepare the action for invocation
func (action *ActionOomReportIndex) Prepare() *ActionOomReportIndexInvocation {
	return &ActionOomReportIndexInvocation{
		Action: action,
		Path:   "/v6.0/oom_reports",
	}
}

// ActionOomReportIndexInvocation is used to configure action for invocation
type ActionOomReportIndexInvocation struct {
	// Pointer to the action
	Action *ActionOomReportIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOomReportIndexInput
	// Global meta input parameters
	MetaInput *ActionOomReportIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOomReportIndexInvocation) NewInput() *ActionOomReportIndexInput {
	inv.Input = &ActionOomReportIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOomReportIndexInvocation) SetInput(input *ActionOomReportIndexInput) *ActionOomReportIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOomReportIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportIndexInvocation) NewMetaInput() *ActionOomReportIndexMetaGlobalInput {
	inv.MetaInput = &ActionOomReportIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportIndexInvocation) SetMetaInput(input *ActionOomReportIndexMetaGlobalInput) *ActionOomReportIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportIndexInvocation) Call() (*ActionOomReportIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOomReportIndexInvocation) callAsQuery() (*ActionOomReportIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOomReportIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OomReports
	}
	return resp, err
}

func (inv *ActionOomReportIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["oom_report[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("Limit") {
			ret["oom_report[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["oom_report[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Node") {
			ret["oom_report[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Offset") {
			ret["oom_report[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Since") {
			ret["oom_report[since]"] = inv.Input.Since
		}
		if inv.IsParameterSelected("Until") {
			ret["oom_report[until]"] = inv.Input.Until
		}
		if inv.IsParameterSelected("User") {
			ret["oom_report[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["oom_report[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionOomReportIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
