package client

import (
	"strings"
)

// ActionDatasetSnapshotShow is a type for action Dataset.Snapshot#Show
type ActionDatasetSnapshotShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetSnapshotShow(client *Client) *ActionDatasetSnapshotShow {
	return &ActionDatasetSnapshotShow{
		Client: client,
	}
}

// ActionDatasetSnapshotShowMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetSnapshotShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetSnapshotShowMetaGlobalInput) SetNo(value bool) *ActionDatasetSnapshotShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetSnapshotShowMetaGlobalInput) SetIncludes(value string) *ActionDatasetSnapshotShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetSnapshotShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetSnapshotShowMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetSnapshotShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetSnapshotShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionDatasetSnapshotShowOutput is a type for action output parameters
type ActionDatasetSnapshotShowOutput struct {
	Id int64 `json:"id"`
	Dataset *ActionDatasetShowOutput `json:"dataset"`
	Name string `json:"name"`
	Label string `json:"label"`
	CreatedAt string `json:"created_at"`
	HistoryId int64 `json:"history_id"`
	Mount *ActionVpsMountShowOutput `json:"mount"`
}


// Type for action response, including envelope
type ActionDatasetSnapshotShowResponse struct {
	Action *ActionDatasetSnapshotShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Snapshot *ActionDatasetSnapshotShowOutput `json:"snapshot"`
	}

	// Action output without the namespace
	Output *ActionDatasetSnapshotShowOutput
}


// Prepare the action for invocation
func (action *ActionDatasetSnapshotShow) Prepare() *ActionDatasetSnapshotShowInvocation {
	return &ActionDatasetSnapshotShowInvocation{
		Action: action,
		Path: "/v5.0/datasets/:dataset_id/snapshots/:snapshot_id",
	}
}

// ActionDatasetSnapshotShowInvocation is used to configure action for invocation
type ActionDatasetSnapshotShowInvocation struct {
	// Pointer to the action
	Action *ActionDatasetSnapshotShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDatasetSnapshotShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetSnapshotShowInvocation) SetPathParamInt(param string, value int64) *ActionDatasetSnapshotShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetSnapshotShowInvocation) SetPathParamString(param string, value string) *ActionDatasetSnapshotShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetSnapshotShowInvocation) SetMetaInput(input *ActionDatasetSnapshotShowMetaGlobalInput) *ActionDatasetSnapshotShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetSnapshotShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetSnapshotShowInvocation) Call() (*ActionDatasetSnapshotShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetSnapshotShowInvocation) callAsQuery() (*ActionDatasetSnapshotShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetSnapshotShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Snapshot
	}
	return resp, err
}




func (inv *ActionDatasetSnapshotShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

