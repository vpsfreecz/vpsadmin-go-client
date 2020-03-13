package client

import (
)

// ActionNewsLogIndex is a type for action News_log#Index
type ActionNewsLogIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNewsLogIndex(client *Client) *ActionNewsLogIndex {
	return &ActionNewsLogIndex{
		Client: client,
	}
}

// ActionNewsLogIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNewsLogIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNewsLogIndexMetaGlobalInput) SetNo(value bool) *ActionNewsLogIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNewsLogIndexMetaGlobalInput) SetCount(value bool) *ActionNewsLogIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNewsLogIndexMetaGlobalInput) SetIncludes(value string) *ActionNewsLogIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNewsLogIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNewsLogIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogIndexInput is a type for action input parameters
type ActionNewsLogIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Since string `json:"since"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionNewsLogIndexInput) SetOffset(value int64) *ActionNewsLogIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNewsLogIndexInput) SetLimit(value int64) *ActionNewsLogIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetSince sets parameter Since to value and selects it for sending
func (in *ActionNewsLogIndexInput) SetSince(value string) *ActionNewsLogIndexInput {
	in.Since = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Since"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogIndexInput) SelectParameters(params ...string) *ActionNewsLogIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNewsLogIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionNewsLogIndexOutput is a type for action output parameters
type ActionNewsLogIndexOutput struct {
	Id int64 `json:"id"`
	Message string `json:"message"`
	PublishedAt string `json:"published_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionNewsLogIndexResponse struct {
	Action *ActionNewsLogIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NewsLogs []*ActionNewsLogIndexOutput `json:"news_logs"`
	}

	// Action output without the namespace
	Output []*ActionNewsLogIndexOutput
}


// Prepare the action for invocation
func (action *ActionNewsLogIndex) Prepare() *ActionNewsLogIndexInvocation {
	return &ActionNewsLogIndexInvocation{
		Action: action,
		Path: "/v6.0/news_logs",
	}
}

// ActionNewsLogIndexInvocation is used to configure action for invocation
type ActionNewsLogIndexInvocation struct {
	// Pointer to the action
	Action *ActionNewsLogIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNewsLogIndexInput
	// Global meta input parameters
	MetaInput *ActionNewsLogIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNewsLogIndexInvocation) NewInput() *ActionNewsLogIndexInput {
	inv.Input = &ActionNewsLogIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNewsLogIndexInvocation) SetInput(input *ActionNewsLogIndexInput) *ActionNewsLogIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNewsLogIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNewsLogIndexInvocation) NewMetaInput() *ActionNewsLogIndexMetaGlobalInput {
	inv.MetaInput = &ActionNewsLogIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNewsLogIndexInvocation) SetMetaInput(input *ActionNewsLogIndexMetaGlobalInput) *ActionNewsLogIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNewsLogIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNewsLogIndexInvocation) Call() (*ActionNewsLogIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNewsLogIndexInvocation) callAsQuery() (*ActionNewsLogIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNewsLogIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NewsLogs
	}
	return resp, err
}



func (inv *ActionNewsLogIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["news_log[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["news_log[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Since") {
			ret["news_log[since]"] = inv.Input.Since
		}
	}
}

func (inv *ActionNewsLogIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

