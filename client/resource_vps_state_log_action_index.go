package client

import (
	"strings"
)

// ActionVpsStateLogIndex is a type for action Vps.State_log#Index
type ActionVpsStateLogIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsStateLogIndex(client *Client) *ActionVpsStateLogIndex {
	return &ActionVpsStateLogIndex{
		Client: client,
	}
}

// ActionVpsStateLogIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsStateLogIndexMetaGlobalInput struct {
	Count bool `json:"count"`
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsStateLogIndexMetaGlobalInput) SetCount(value bool) *ActionVpsStateLogIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsStateLogIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsStateLogIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsStateLogIndexMetaGlobalInput) SetNo(value bool) *ActionVpsStateLogIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsStateLogIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsStateLogIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsStateLogIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsStateLogIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsStateLogIndexInput is a type for action input parameters
type ActionVpsStateLogIndexInput struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsStateLogIndexInput) SetLimit(value int64) *ActionVpsStateLogIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionVpsStateLogIndexInput) SetOffset(value int64) *ActionVpsStateLogIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsStateLogIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsStateLogIndexInput) SelectParameters(params ...string) *ActionVpsStateLogIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsStateLogIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionVpsStateLogIndexOutput is a type for action output parameters
type ActionVpsStateLogIndexOutput struct {
	ChangedAt string `json:"changed_at"`
	Expiration string `json:"expiration"`
	Id int64 `json:"id"`
	Reason string `json:"reason"`
	State string `json:"state"`
	User *ActionUserShowOutput `json:"user"`
}


// Type for action response, including envelope
type ActionVpsStateLogIndexResponse struct {
	Action *ActionVpsStateLogIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		StateLogs []*ActionVpsStateLogIndexOutput `json:"state_logs"`
	}

	// Action output without the namespace
	Output []*ActionVpsStateLogIndexOutput
}


// Prepare the action for invocation
func (action *ActionVpsStateLogIndex) Prepare() *ActionVpsStateLogIndexInvocation {
	return &ActionVpsStateLogIndexInvocation{
		Action: action,
		Path: "/v6.0/vpses/{vps_id}/state_logs",
	}
}

// ActionVpsStateLogIndexInvocation is used to configure action for invocation
type ActionVpsStateLogIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsStateLogIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsStateLogIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsStateLogIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsStateLogIndexInvocation) SetPathParamInt(param string, value int64) *ActionVpsStateLogIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsStateLogIndexInvocation) SetPathParamString(param string, value string) *ActionVpsStateLogIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsStateLogIndexInvocation) NewInput() *ActionVpsStateLogIndexInput {
	inv.Input = &ActionVpsStateLogIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsStateLogIndexInvocation) SetInput(input *ActionVpsStateLogIndexInput) *ActionVpsStateLogIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsStateLogIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsStateLogIndexInvocation) NewMetaInput() *ActionVpsStateLogIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsStateLogIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsStateLogIndexInvocation) SetMetaInput(input *ActionVpsStateLogIndexMetaGlobalInput) *ActionVpsStateLogIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsStateLogIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsStateLogIndexInvocation) Call() (*ActionVpsStateLogIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsStateLogIndexInvocation) callAsQuery() (*ActionVpsStateLogIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsStateLogIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.StateLogs
	}
	return resp, err
}



func (inv *ActionVpsStateLogIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["state_log[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["state_log[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionVpsStateLogIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

