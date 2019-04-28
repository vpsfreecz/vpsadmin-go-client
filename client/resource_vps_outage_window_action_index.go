package client

import (
	"strings"
)

// ActionVpsOutageWindowIndex is a type for action Vps.Outage_window#Index
type ActionVpsOutageWindowIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsOutageWindowIndex(client *Client) *ActionVpsOutageWindowIndex {
	return &ActionVpsOutageWindowIndex{
		Client: client,
	}
}

// ActionVpsOutageWindowIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsOutageWindowIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsOutageWindowIndexMetaGlobalInput) SetNo(value bool) *ActionVpsOutageWindowIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsOutageWindowIndexMetaGlobalInput) SetCount(value bool) *ActionVpsOutageWindowIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsOutageWindowIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsOutageWindowIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageWindowIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageWindowIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsOutageWindowIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageWindowIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsOutageWindowIndexInput is a type for action input parameters
type ActionVpsOutageWindowIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionVpsOutageWindowIndexInput) SetOffset(value int64) *ActionVpsOutageWindowIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsOutageWindowIndexInput) SetLimit(value int64) *ActionVpsOutageWindowIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageWindowIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageWindowIndexInput) SelectParameters(params ...string) *ActionVpsOutageWindowIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageWindowIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionVpsOutageWindowIndexOutput is a type for action output parameters
type ActionVpsOutageWindowIndexOutput struct {
	Weekday int64 `json:"weekday"`
	IsOpen bool `json:"is_open"`
	OpensAt int64 `json:"opens_at"`
	ClosesAt int64 `json:"closes_at"`
}


// Type for action response, including envelope
type ActionVpsOutageWindowIndexResponse struct {
	Action *ActionVpsOutageWindowIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageWindows []*ActionVpsOutageWindowIndexOutput `json:"outage_windows"`
	}

	// Action output without the namespace
	Output []*ActionVpsOutageWindowIndexOutput
}


// Prepare the action for invocation
func (action *ActionVpsOutageWindowIndex) Prepare() *ActionVpsOutageWindowIndexInvocation {
	return &ActionVpsOutageWindowIndexInvocation{
		Action: action,
		Path: "/v5.0/vpses/:vps_id/outage_windows",
	}
}

// ActionVpsOutageWindowIndexInvocation is used to configure action for invocation
type ActionVpsOutageWindowIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsOutageWindowIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsOutageWindowIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsOutageWindowIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsOutageWindowIndexInvocation) SetPathParamInt(param string, value int64) *ActionVpsOutageWindowIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsOutageWindowIndexInvocation) SetPathParamString(param string, value string) *ActionVpsOutageWindowIndexInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsOutageWindowIndexInvocation) NewInput() *ActionVpsOutageWindowIndexInput {
	inv.Input = &ActionVpsOutageWindowIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsOutageWindowIndexInvocation) SetInput(input *ActionVpsOutageWindowIndexInput) *ActionVpsOutageWindowIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsOutageWindowIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsOutageWindowIndexInvocation) NewMetaInput() *ActionVpsOutageWindowIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsOutageWindowIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsOutageWindowIndexInvocation) SetMetaInput(input *ActionVpsOutageWindowIndexMetaGlobalInput) *ActionVpsOutageWindowIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsOutageWindowIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsOutageWindowIndexInvocation) Call() (*ActionVpsOutageWindowIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsOutageWindowIndexInvocation) callAsQuery() (*ActionVpsOutageWindowIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsOutageWindowIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageWindows
	}
	return resp, err
}



func (inv *ActionVpsOutageWindowIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["outage_window[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["outage_window[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionVpsOutageWindowIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

