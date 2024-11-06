package client

import (
	"strings"
)

// ActionDatasetExpansionHistoryIndex is a type for action Dataset_expansion.History#Index
type ActionDatasetExpansionHistoryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetExpansionHistoryIndex(client *Client) *ActionDatasetExpansionHistoryIndex {
	return &ActionDatasetExpansionHistoryIndex{
		Client: client,
	}
}

// ActionDatasetExpansionHistoryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetExpansionHistoryIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDatasetExpansionHistoryIndexMetaGlobalInput) SetCount(value bool) *ActionDatasetExpansionHistoryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetExpansionHistoryIndexMetaGlobalInput) SetIncludes(value string) *ActionDatasetExpansionHistoryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetExpansionHistoryIndexMetaGlobalInput) SetNo(value bool) *ActionDatasetExpansionHistoryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionHistoryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionHistoryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetExpansionHistoryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetExpansionHistoryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionHistoryIndexInput is a type for action input parameters
type ActionDatasetExpansionHistoryIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDatasetExpansionHistoryIndexInput) SetFromId(value int64) *ActionDatasetExpansionHistoryIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDatasetExpansionHistoryIndexInput) SetLimit(value int64) *ActionDatasetExpansionHistoryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionHistoryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionHistoryIndexInput) SelectParameters(params ...string) *ActionDatasetExpansionHistoryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetExpansionHistoryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetExpansionHistoryIndexInput) UnselectParameters(params ...string) *ActionDatasetExpansionHistoryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetExpansionHistoryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionHistoryIndexOutput is a type for action output parameters
type ActionDatasetExpansionHistoryIndexOutput struct {
	AddedSpace       int64                 `json:"added_space"`
	Admin            *ActionUserShowOutput `json:"admin"`
	CreatedAt        string                `json:"created_at"`
	Id               int64                 `json:"id"`
	NewRefquota      int64                 `json:"new_refquota"`
	OriginalRefquota int64                 `json:"original_refquota"`
}

// Type for action response, including envelope
type ActionDatasetExpansionHistoryIndexResponse struct {
	Action *ActionDatasetExpansionHistoryIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Histories []*ActionDatasetExpansionHistoryIndexOutput `json:"histories"`
	}

	// Action output without the namespace
	Output []*ActionDatasetExpansionHistoryIndexOutput
}

// Prepare the action for invocation
func (action *ActionDatasetExpansionHistoryIndex) Prepare() *ActionDatasetExpansionHistoryIndexInvocation {
	return &ActionDatasetExpansionHistoryIndexInvocation{
		Action: action,
		Path:   "/v7.0/dataset_expansions/{dataset_expansion_id}/history",
	}
}

// ActionDatasetExpansionHistoryIndexInvocation is used to configure action for invocation
type ActionDatasetExpansionHistoryIndexInvocation struct {
	// Pointer to the action
	Action *ActionDatasetExpansionHistoryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetExpansionHistoryIndexInput
	// Global meta input parameters
	MetaInput *ActionDatasetExpansionHistoryIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetExpansionHistoryIndexInvocation) SetPathParamInt(param string, value int64) *ActionDatasetExpansionHistoryIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetExpansionHistoryIndexInvocation) SetPathParamString(param string, value string) *ActionDatasetExpansionHistoryIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetExpansionHistoryIndexInvocation) NewInput() *ActionDatasetExpansionHistoryIndexInput {
	inv.Input = &ActionDatasetExpansionHistoryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetExpansionHistoryIndexInvocation) SetInput(input *ActionDatasetExpansionHistoryIndexInput) *ActionDatasetExpansionHistoryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetExpansionHistoryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetExpansionHistoryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetExpansionHistoryIndexInvocation) NewMetaInput() *ActionDatasetExpansionHistoryIndexMetaGlobalInput {
	inv.MetaInput = &ActionDatasetExpansionHistoryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetExpansionHistoryIndexInvocation) SetMetaInput(input *ActionDatasetExpansionHistoryIndexMetaGlobalInput) *ActionDatasetExpansionHistoryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetExpansionHistoryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetExpansionHistoryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetExpansionHistoryIndexInvocation) Call() (*ActionDatasetExpansionHistoryIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetExpansionHistoryIndexInvocation) callAsQuery() (*ActionDatasetExpansionHistoryIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetExpansionHistoryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Histories
	}
	return resp, err
}

func (inv *ActionDatasetExpansionHistoryIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["history[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["history[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionDatasetExpansionHistoryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
