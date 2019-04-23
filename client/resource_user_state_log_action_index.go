package client

import (
	"strings"
)

// ActionUserStateLogIndex is a type for action User.State_log#Index
type ActionUserStateLogIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserStateLogIndex(client *Client) *ActionUserStateLogIndex {
	return &ActionUserStateLogIndex{
		Client: client,
	}
}

// ActionUserStateLogIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserStateLogIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserStateLogIndexMetaGlobalInput) SetNo(value bool) *ActionUserStateLogIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserStateLogIndexMetaGlobalInput) SetCount(value bool) *ActionUserStateLogIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserStateLogIndexMetaGlobalInput) SetIncludes(value string) *ActionUserStateLogIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserStateLogIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserStateLogIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserStateLogIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserStateLogIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserStateLogIndexInput is a type for action input parameters
type ActionUserStateLogIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserStateLogIndexInput) SetOffset(value int64) *ActionUserStateLogIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserStateLogIndexInput) SetLimit(value int64) *ActionUserStateLogIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserStateLogIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserStateLogIndexInput) SelectParameters(params ...string) *ActionUserStateLogIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserStateLogIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserStateLogIndexOutput is a type for action output parameters
type ActionUserStateLogIndexOutput struct {
	Id int64 `json:"id"`
	State string `json:"state"`
	ChangedAt string `json:"changed_at"`
	Expiration string `json:"expiration"`
	User *ActionUserShowOutput `json:"user"`
	Reason string `json:"reason"`
}


// Type for action response, including envelope
type ActionUserStateLogIndexResponse struct {
	Action *ActionUserStateLogIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		StateLogs []*ActionUserStateLogIndexOutput `json:"state_logs"`
	}

	// Action output without the namespace
	Output []*ActionUserStateLogIndexOutput
}


// Prepare the action for invocation
func (action *ActionUserStateLogIndex) Prepare() *ActionUserStateLogIndexInvocation {
	return &ActionUserStateLogIndexInvocation{
		Action: action,
		Path: "/v5.0/users/:user_id/state_logs",
	}
}

// ActionUserStateLogIndexInvocation is used to configure action for invocation
type ActionUserStateLogIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserStateLogIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserStateLogIndexInput
	// Global meta input parameters
	MetaInput *ActionUserStateLogIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserStateLogIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserStateLogIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserStateLogIndexInvocation) SetPathParamString(param string, value string) *ActionUserStateLogIndexInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserStateLogIndexInvocation) SetInput(input *ActionUserStateLogIndexInput) *ActionUserStateLogIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserStateLogIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserStateLogIndexInvocation) SetMetaInput(input *ActionUserStateLogIndexMetaGlobalInput) *ActionUserStateLogIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserStateLogIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserStateLogIndexInvocation) Call() (*ActionUserStateLogIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserStateLogIndexInvocation) callAsQuery() (*ActionUserStateLogIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserStateLogIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.StateLogs
	}
	return resp, err
}



func (inv *ActionUserStateLogIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["state_log[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["state_log[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserStateLogIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

