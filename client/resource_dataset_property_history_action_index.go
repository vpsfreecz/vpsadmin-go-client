package client

import (
	"strings"
)

// ActionDatasetPropertyHistoryIndex is a type for action Dataset.Property_history#Index
type ActionDatasetPropertyHistoryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetPropertyHistoryIndex(client *Client) *ActionDatasetPropertyHistoryIndex {
	return &ActionDatasetPropertyHistoryIndex{
		Client: client,
	}
}

// ActionDatasetPropertyHistoryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetPropertyHistoryIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDatasetPropertyHistoryIndexMetaGlobalInput) SetCount(value bool) *ActionDatasetPropertyHistoryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetPropertyHistoryIndexMetaGlobalInput) SetIncludes(value string) *ActionDatasetPropertyHistoryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetPropertyHistoryIndexMetaGlobalInput) SetNo(value bool) *ActionDatasetPropertyHistoryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPropertyHistoryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPropertyHistoryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetPropertyHistoryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetPropertyHistoryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetPropertyHistoryIndexInput is a type for action input parameters
type ActionDatasetPropertyHistoryIndexInput struct {
	From   string `json:"from"`
	Limit  int64  `json:"limit"`
	Name   string `json:"name"`
	Offset int64  `json:"offset"`
	To     string `json:"to"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionDatasetPropertyHistoryIndexInput) SetFrom(value string) *ActionDatasetPropertyHistoryIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDatasetPropertyHistoryIndexInput) SetLimit(value int64) *ActionDatasetPropertyHistoryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionDatasetPropertyHistoryIndexInput) SetName(value string) *ActionDatasetPropertyHistoryIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionDatasetPropertyHistoryIndexInput) SetOffset(value int64) *ActionDatasetPropertyHistoryIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionDatasetPropertyHistoryIndexInput) SetTo(value string) *ActionDatasetPropertyHistoryIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPropertyHistoryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPropertyHistoryIndexInput) SelectParameters(params ...string) *ActionDatasetPropertyHistoryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetPropertyHistoryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetPropertyHistoryIndexInput) UnselectParameters(params ...string) *ActionDatasetPropertyHistoryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetPropertyHistoryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetPropertyHistoryIndexOutput is a type for action output parameters
type ActionDatasetPropertyHistoryIndexOutput struct {
	CreatedAt string `json:"created_at"`
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Value     int64  `json:"value"`
}

// Type for action response, including envelope
type ActionDatasetPropertyHistoryIndexResponse struct {
	Action *ActionDatasetPropertyHistoryIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PropertyHistories []*ActionDatasetPropertyHistoryIndexOutput `json:"property_histories"`
	}

	// Action output without the namespace
	Output []*ActionDatasetPropertyHistoryIndexOutput
}

// Prepare the action for invocation
func (action *ActionDatasetPropertyHistoryIndex) Prepare() *ActionDatasetPropertyHistoryIndexInvocation {
	return &ActionDatasetPropertyHistoryIndexInvocation{
		Action: action,
		Path:   "/v6.0/datasets/{dataset_id}/property_history",
	}
}

// ActionDatasetPropertyHistoryIndexInvocation is used to configure action for invocation
type ActionDatasetPropertyHistoryIndexInvocation struct {
	// Pointer to the action
	Action *ActionDatasetPropertyHistoryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetPropertyHistoryIndexInput
	// Global meta input parameters
	MetaInput *ActionDatasetPropertyHistoryIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetPropertyHistoryIndexInvocation) SetPathParamInt(param string, value int64) *ActionDatasetPropertyHistoryIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetPropertyHistoryIndexInvocation) SetPathParamString(param string, value string) *ActionDatasetPropertyHistoryIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetPropertyHistoryIndexInvocation) NewInput() *ActionDatasetPropertyHistoryIndexInput {
	inv.Input = &ActionDatasetPropertyHistoryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetPropertyHistoryIndexInvocation) SetInput(input *ActionDatasetPropertyHistoryIndexInput) *ActionDatasetPropertyHistoryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetPropertyHistoryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetPropertyHistoryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetPropertyHistoryIndexInvocation) NewMetaInput() *ActionDatasetPropertyHistoryIndexMetaGlobalInput {
	inv.MetaInput = &ActionDatasetPropertyHistoryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetPropertyHistoryIndexInvocation) SetMetaInput(input *ActionDatasetPropertyHistoryIndexMetaGlobalInput) *ActionDatasetPropertyHistoryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetPropertyHistoryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetPropertyHistoryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetPropertyHistoryIndexInvocation) Call() (*ActionDatasetPropertyHistoryIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetPropertyHistoryIndexInvocation) callAsQuery() (*ActionDatasetPropertyHistoryIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetPropertyHistoryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.PropertyHistories
	}
	return resp, err
}

func (inv *ActionDatasetPropertyHistoryIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["property_history[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("Limit") {
			ret["property_history[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["property_history[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Offset") {
			ret["property_history[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("To") {
			ret["property_history[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionDatasetPropertyHistoryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
