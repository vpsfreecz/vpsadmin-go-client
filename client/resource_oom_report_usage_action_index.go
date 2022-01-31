package client

import (
	"strings"
)

// ActionOomReportUsageIndex is a type for action Oom_report.Usage#Index
type ActionOomReportUsageIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportUsageIndex(client *Client) *ActionOomReportUsageIndex {
	return &ActionOomReportUsageIndex{
		Client: client,
	}
}

// ActionOomReportUsageIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportUsageIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOomReportUsageIndexMetaGlobalInput) SetCount(value bool) *ActionOomReportUsageIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportUsageIndexMetaGlobalInput) SetIncludes(value string) *ActionOomReportUsageIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportUsageIndexMetaGlobalInput) SetNo(value bool) *ActionOomReportUsageIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportUsageIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportUsageIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportUsageIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportUsageIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportUsageIndexInput is a type for action input parameters
type ActionOomReportUsageIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOomReportUsageIndexInput) SetLimit(value int64) *ActionOomReportUsageIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionOomReportUsageIndexInput) SetOffset(value int64) *ActionOomReportUsageIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportUsageIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportUsageIndexInput) SelectParameters(params ...string) *ActionOomReportUsageIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOomReportUsageIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOomReportUsageIndexInput) UnselectParameters(params ...string) *ActionOomReportUsageIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOomReportUsageIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportUsageIndexOutput is a type for action output parameters
type ActionOomReportUsageIndexOutput struct {
	Failcnt int64  `json:"failcnt"`
	Id      int64  `json:"id"`
	Limit   int64  `json:"limit"`
	Memtype string `json:"memtype"`
	Usage   int64  `json:"usage"`
}

// Type for action response, including envelope
type ActionOomReportUsageIndexResponse struct {
	Action *ActionOomReportUsageIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Usages []*ActionOomReportUsageIndexOutput `json:"usages"`
	}

	// Action output without the namespace
	Output []*ActionOomReportUsageIndexOutput
}

// Prepare the action for invocation
func (action *ActionOomReportUsageIndex) Prepare() *ActionOomReportUsageIndexInvocation {
	return &ActionOomReportUsageIndexInvocation{
		Action: action,
		Path:   "/v6.0/oom_reports/{oom_report_id}/usages",
	}
}

// ActionOomReportUsageIndexInvocation is used to configure action for invocation
type ActionOomReportUsageIndexInvocation struct {
	// Pointer to the action
	Action *ActionOomReportUsageIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOomReportUsageIndexInput
	// Global meta input parameters
	MetaInput *ActionOomReportUsageIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportUsageIndexInvocation) SetPathParamInt(param string, value int64) *ActionOomReportUsageIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportUsageIndexInvocation) SetPathParamString(param string, value string) *ActionOomReportUsageIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOomReportUsageIndexInvocation) NewInput() *ActionOomReportUsageIndexInput {
	inv.Input = &ActionOomReportUsageIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOomReportUsageIndexInvocation) SetInput(input *ActionOomReportUsageIndexInput) *ActionOomReportUsageIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOomReportUsageIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOomReportUsageIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportUsageIndexInvocation) NewMetaInput() *ActionOomReportUsageIndexMetaGlobalInput {
	inv.MetaInput = &ActionOomReportUsageIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportUsageIndexInvocation) SetMetaInput(input *ActionOomReportUsageIndexMetaGlobalInput) *ActionOomReportUsageIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportUsageIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportUsageIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportUsageIndexInvocation) Call() (*ActionOomReportUsageIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOomReportUsageIndexInvocation) callAsQuery() (*ActionOomReportUsageIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOomReportUsageIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Usages
	}
	return resp, err
}

func (inv *ActionOomReportUsageIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["usage[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["usage[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionOomReportUsageIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
