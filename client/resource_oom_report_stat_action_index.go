package client

import (
	"strings"
)

// ActionOomReportStatIndex is a type for action Oom_report.Stat#Index
type ActionOomReportStatIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportStatIndex(client *Client) *ActionOomReportStatIndex {
	return &ActionOomReportStatIndex{
		Client: client,
	}
}

// ActionOomReportStatIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportStatIndexMetaGlobalInput struct {
	Count bool `json:"count"`
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOomReportStatIndexMetaGlobalInput) SetCount(value bool) *ActionOomReportStatIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportStatIndexMetaGlobalInput) SetIncludes(value string) *ActionOomReportStatIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportStatIndexMetaGlobalInput) SetNo(value bool) *ActionOomReportStatIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportStatIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportStatIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportStatIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportStatIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportStatIndexInput is a type for action input parameters
type ActionOomReportStatIndexInput struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOomReportStatIndexInput) SetLimit(value int64) *ActionOomReportStatIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionOomReportStatIndexInput) SetOffset(value int64) *ActionOomReportStatIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportStatIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportStatIndexInput) SelectParameters(params ...string) *ActionOomReportStatIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportStatIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionOomReportStatIndexOutput is a type for action output parameters
type ActionOomReportStatIndexOutput struct {
	Id int64 `json:"id"`
	Parameter string `json:"parameter"`
	Value int64 `json:"value"`
}


// Type for action response, including envelope
type ActionOomReportStatIndexResponse struct {
	Action *ActionOomReportStatIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Stats []*ActionOomReportStatIndexOutput `json:"stats"`
	}

	// Action output without the namespace
	Output []*ActionOomReportStatIndexOutput
}


// Prepare the action for invocation
func (action *ActionOomReportStatIndex) Prepare() *ActionOomReportStatIndexInvocation {
	return &ActionOomReportStatIndexInvocation{
		Action: action,
		Path: "/v6.0/oom_reports/{oom_report_id}/stats",
	}
}

// ActionOomReportStatIndexInvocation is used to configure action for invocation
type ActionOomReportStatIndexInvocation struct {
	// Pointer to the action
	Action *ActionOomReportStatIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOomReportStatIndexInput
	// Global meta input parameters
	MetaInput *ActionOomReportStatIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportStatIndexInvocation) SetPathParamInt(param string, value int64) *ActionOomReportStatIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportStatIndexInvocation) SetPathParamString(param string, value string) *ActionOomReportStatIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOomReportStatIndexInvocation) NewInput() *ActionOomReportStatIndexInput {
	inv.Input = &ActionOomReportStatIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOomReportStatIndexInvocation) SetInput(input *ActionOomReportStatIndexInput) *ActionOomReportStatIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOomReportStatIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportStatIndexInvocation) NewMetaInput() *ActionOomReportStatIndexMetaGlobalInput {
	inv.MetaInput = &ActionOomReportStatIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportStatIndexInvocation) SetMetaInput(input *ActionOomReportStatIndexMetaGlobalInput) *ActionOomReportStatIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportStatIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportStatIndexInvocation) Call() (*ActionOomReportStatIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOomReportStatIndexInvocation) callAsQuery() (*ActionOomReportStatIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOomReportStatIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Stats
	}
	return resp, err
}



func (inv *ActionOomReportStatIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["stat[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["stat[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionOomReportStatIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

