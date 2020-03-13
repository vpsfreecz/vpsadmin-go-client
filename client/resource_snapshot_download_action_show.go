package client

import (
	"strings"
)

// ActionSnapshotDownloadShow is a type for action Snapshot_download#Show
type ActionSnapshotDownloadShow struct {
	// Pointer to client
	Client *Client
}

func NewActionSnapshotDownloadShow(client *Client) *ActionSnapshotDownloadShow {
	return &ActionSnapshotDownloadShow{
		Client: client,
	}
}

// ActionSnapshotDownloadShowMetaGlobalInput is a type for action global meta input parameters
type ActionSnapshotDownloadShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSnapshotDownloadShowMetaGlobalInput) SetNo(value bool) *ActionSnapshotDownloadShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSnapshotDownloadShowMetaGlobalInput) SetIncludes(value string) *ActionSnapshotDownloadShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionSnapshotDownloadShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSnapshotDownloadShowMetaGlobalInput) SelectParameters(params ...string) *ActionSnapshotDownloadShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSnapshotDownloadShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionSnapshotDownloadShowOutput is a type for action output parameters
type ActionSnapshotDownloadShowOutput struct {
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
type ActionSnapshotDownloadShowResponse struct {
	Action *ActionSnapshotDownloadShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SnapshotDownload *ActionSnapshotDownloadShowOutput `json:"snapshot_download"`
	}

	// Action output without the namespace
	Output *ActionSnapshotDownloadShowOutput
}


// Prepare the action for invocation
func (action *ActionSnapshotDownloadShow) Prepare() *ActionSnapshotDownloadShowInvocation {
	return &ActionSnapshotDownloadShowInvocation{
		Action: action,
		Path: "/v6.0/snapshot_downloads/{snapshot_download_id}",
	}
}

// ActionSnapshotDownloadShowInvocation is used to configure action for invocation
type ActionSnapshotDownloadShowInvocation struct {
	// Pointer to the action
	Action *ActionSnapshotDownloadShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSnapshotDownloadShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSnapshotDownloadShowInvocation) SetPathParamInt(param string, value int64) *ActionSnapshotDownloadShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSnapshotDownloadShowInvocation) SetPathParamString(param string, value string) *ActionSnapshotDownloadShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSnapshotDownloadShowInvocation) NewMetaInput() *ActionSnapshotDownloadShowMetaGlobalInput {
	inv.MetaInput = &ActionSnapshotDownloadShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSnapshotDownloadShowInvocation) SetMetaInput(input *ActionSnapshotDownloadShowMetaGlobalInput) *ActionSnapshotDownloadShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSnapshotDownloadShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSnapshotDownloadShowInvocation) Call() (*ActionSnapshotDownloadShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionSnapshotDownloadShowInvocation) callAsQuery() (*ActionSnapshotDownloadShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSnapshotDownloadShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SnapshotDownload
	}
	return resp, err
}




func (inv *ActionSnapshotDownloadShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

