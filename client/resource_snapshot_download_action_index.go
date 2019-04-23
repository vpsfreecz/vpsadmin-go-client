package client

import (
)

// ActionSnapshotDownloadIndex is a type for action Snapshot_download#Index
type ActionSnapshotDownloadIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionSnapshotDownloadIndex(client *Client) *ActionSnapshotDownloadIndex {
	return &ActionSnapshotDownloadIndex{
		Client: client,
	}
}

// ActionSnapshotDownloadIndexMetaGlobalInput is a type for action global meta input parameters
type ActionSnapshotDownloadIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSnapshotDownloadIndexMetaGlobalInput) SetNo(value bool) *ActionSnapshotDownloadIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionSnapshotDownloadIndexMetaGlobalInput) SetCount(value bool) *ActionSnapshotDownloadIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSnapshotDownloadIndexMetaGlobalInput) SetIncludes(value string) *ActionSnapshotDownloadIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionSnapshotDownloadIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSnapshotDownloadIndexMetaGlobalInput) SelectParameters(params ...string) *ActionSnapshotDownloadIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSnapshotDownloadIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSnapshotDownloadIndexInput is a type for action input parameters
type ActionSnapshotDownloadIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Dataset int64 `json:"dataset"`
	Snapshot int64 `json:"snapshot"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionSnapshotDownloadIndexInput) SetOffset(value int64) *ActionSnapshotDownloadIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionSnapshotDownloadIndexInput) SetLimit(value int64) *ActionSnapshotDownloadIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetDataset sets parameter Dataset to value and selects it for sending
func (in *ActionSnapshotDownloadIndexInput) SetDataset(value int64) *ActionSnapshotDownloadIndexInput {
	in.Dataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Dataset"] = nil
	return in
}
// SetSnapshot sets parameter Snapshot to value and selects it for sending
func (in *ActionSnapshotDownloadIndexInput) SetSnapshot(value int64) *ActionSnapshotDownloadIndexInput {
	in.Snapshot = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Snapshot"] = nil
	return in
}

// SelectParameters sets parameters from ActionSnapshotDownloadIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSnapshotDownloadIndexInput) SelectParameters(params ...string) *ActionSnapshotDownloadIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSnapshotDownloadIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionSnapshotDownloadIndexOutput is a type for action output parameters
type ActionSnapshotDownloadIndexOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	Snapshot *ActionDatasetSnapshotShowOutput `json:"snapshot"`
	FromSnapshot *ActionDatasetSnapshotShowOutput `json:"from_snapshot"`
	Format string `json:"format"`
	FileName string `json:"file_name"`
	Url string `json:"url"`
	Size int64 `json:"size"`
	Sha256sum string `json:"sha256sum"`
	Ready bool `json:"ready"`
	ExpirationDate string `json:"expiration_date"`
}


// Type for action response, including envelope
type ActionSnapshotDownloadIndexResponse struct {
	Action *ActionSnapshotDownloadIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SnapshotDownloads []*ActionSnapshotDownloadIndexOutput `json:"snapshot_downloads"`
	}

	// Action output without the namespace
	Output []*ActionSnapshotDownloadIndexOutput
}


// Prepare the action for invocation
func (action *ActionSnapshotDownloadIndex) Prepare() *ActionSnapshotDownloadIndexInvocation {
	return &ActionSnapshotDownloadIndexInvocation{
		Action: action,
		Path: "/v5.0/snapshot_downloads",
	}
}

// ActionSnapshotDownloadIndexInvocation is used to configure action for invocation
type ActionSnapshotDownloadIndexInvocation struct {
	// Pointer to the action
	Action *ActionSnapshotDownloadIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSnapshotDownloadIndexInput
	// Global meta input parameters
	MetaInput *ActionSnapshotDownloadIndexMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionSnapshotDownloadIndexInvocation) SetInput(input *ActionSnapshotDownloadIndexInput) *ActionSnapshotDownloadIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSnapshotDownloadIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSnapshotDownloadIndexInvocation) SetMetaInput(input *ActionSnapshotDownloadIndexMetaGlobalInput) *ActionSnapshotDownloadIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSnapshotDownloadIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSnapshotDownloadIndexInvocation) Call() (*ActionSnapshotDownloadIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionSnapshotDownloadIndexInvocation) callAsQuery() (*ActionSnapshotDownloadIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSnapshotDownloadIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SnapshotDownloads
	}
	return resp, err
}



func (inv *ActionSnapshotDownloadIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["snapshot_download[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["snapshot_download[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Dataset") {
			ret["snapshot_download[dataset]"] = convertInt64ToString(inv.Input.Dataset)
		}
		if inv.IsParameterSelected("Snapshot") {
			ret["snapshot_download[snapshot]"] = convertInt64ToString(inv.Input.Snapshot)
		}
	}
}

func (inv *ActionSnapshotDownloadIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

