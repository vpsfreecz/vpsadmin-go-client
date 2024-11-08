package client

import (
	"strings"
)

// ActionDatasetSnapshotIndex is a type for action Dataset.Snapshot#Index
type ActionDatasetSnapshotIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetSnapshotIndex(client *Client) *ActionDatasetSnapshotIndex {
	return &ActionDatasetSnapshotIndex{
		Client: client,
	}
}

// ActionDatasetSnapshotIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetSnapshotIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDatasetSnapshotIndexMetaGlobalInput) SetCount(value bool) *ActionDatasetSnapshotIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetSnapshotIndexMetaGlobalInput) SetIncludes(value string) *ActionDatasetSnapshotIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetSnapshotIndexMetaGlobalInput) SetNo(value bool) *ActionDatasetSnapshotIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetSnapshotIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetSnapshotIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetSnapshotIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetSnapshotIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetSnapshotIndexInput is a type for action input parameters
type ActionDatasetSnapshotIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDatasetSnapshotIndexInput) SetFromId(value int64) *ActionDatasetSnapshotIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDatasetSnapshotIndexInput) SetLimit(value int64) *ActionDatasetSnapshotIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetSnapshotIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetSnapshotIndexInput) SelectParameters(params ...string) *ActionDatasetSnapshotIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetSnapshotIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetSnapshotIndexInput) UnselectParameters(params ...string) *ActionDatasetSnapshotIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetSnapshotIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetSnapshotIndexOutput is a type for action output parameters
type ActionDatasetSnapshotIndexOutput struct {
	CreatedAt string                    `json:"created_at"`
	Dataset   *ActionDatasetShowOutput  `json:"dataset"`
	Export    *ActionExportShowOutput   `json:"export"`
	HistoryId int64                     `json:"history_id"`
	Id        int64                     `json:"id"`
	Label     string                    `json:"label"`
	Mount     *ActionVpsMountShowOutput `json:"mount"`
	Name      string                    `json:"name"`
}

// Type for action response, including envelope
type ActionDatasetSnapshotIndexResponse struct {
	Action *ActionDatasetSnapshotIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Snapshots []*ActionDatasetSnapshotIndexOutput `json:"snapshots"`
	}

	// Action output without the namespace
	Output []*ActionDatasetSnapshotIndexOutput
}

// Prepare the action for invocation
func (action *ActionDatasetSnapshotIndex) Prepare() *ActionDatasetSnapshotIndexInvocation {
	return &ActionDatasetSnapshotIndexInvocation{
		Action: action,
		Path:   "/v7.0/datasets/{dataset_id}/snapshots",
	}
}

// ActionDatasetSnapshotIndexInvocation is used to configure action for invocation
type ActionDatasetSnapshotIndexInvocation struct {
	// Pointer to the action
	Action *ActionDatasetSnapshotIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetSnapshotIndexInput
	// Global meta input parameters
	MetaInput *ActionDatasetSnapshotIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetSnapshotIndexInvocation) SetPathParamInt(param string, value int64) *ActionDatasetSnapshotIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetSnapshotIndexInvocation) SetPathParamString(param string, value string) *ActionDatasetSnapshotIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetSnapshotIndexInvocation) NewInput() *ActionDatasetSnapshotIndexInput {
	inv.Input = &ActionDatasetSnapshotIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetSnapshotIndexInvocation) SetInput(input *ActionDatasetSnapshotIndexInput) *ActionDatasetSnapshotIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetSnapshotIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetSnapshotIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetSnapshotIndexInvocation) NewMetaInput() *ActionDatasetSnapshotIndexMetaGlobalInput {
	inv.MetaInput = &ActionDatasetSnapshotIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetSnapshotIndexInvocation) SetMetaInput(input *ActionDatasetSnapshotIndexMetaGlobalInput) *ActionDatasetSnapshotIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetSnapshotIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetSnapshotIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetSnapshotIndexInvocation) Call() (*ActionDatasetSnapshotIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetSnapshotIndexInvocation) callAsQuery() (*ActionDatasetSnapshotIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetSnapshotIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Snapshots
	}
	return resp, err
}

func (inv *ActionDatasetSnapshotIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["snapshot[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["snapshot[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionDatasetSnapshotIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
