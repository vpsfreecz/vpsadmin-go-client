package client

import (
)

// ActionSnapshotDownloadCreate is a type for action Snapshot_download#Create
type ActionSnapshotDownloadCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionSnapshotDownloadCreate(client *Client) *ActionSnapshotDownloadCreate {
	return &ActionSnapshotDownloadCreate{
		Client: client,
	}
}

// ActionSnapshotDownloadCreateMetaGlobalInput is a type for action global meta input parameters
type ActionSnapshotDownloadCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSnapshotDownloadCreateMetaGlobalInput) SetNo(value bool) *ActionSnapshotDownloadCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSnapshotDownloadCreateMetaGlobalInput) SetIncludes(value string) *ActionSnapshotDownloadCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionSnapshotDownloadCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSnapshotDownloadCreateMetaGlobalInput) SelectParameters(params ...string) *ActionSnapshotDownloadCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSnapshotDownloadCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSnapshotDownloadCreateInput is a type for action input parameters
type ActionSnapshotDownloadCreateInput struct {
	Snapshot int64 `json:"snapshot"`
	FromSnapshot int64 `json:"from_snapshot"`
	Format string `json:"format"`
	SendMail bool `json:"send_mail"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetSnapshot sets parameter Snapshot to value and selects it for sending
func (in *ActionSnapshotDownloadCreateInput) SetSnapshot(value int64) *ActionSnapshotDownloadCreateInput {
	in.Snapshot = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Snapshot"] = nil
	return in
}
// SetFromSnapshot sets parameter FromSnapshot to value and selects it for sending
func (in *ActionSnapshotDownloadCreateInput) SetFromSnapshot(value int64) *ActionSnapshotDownloadCreateInput {
	in.FromSnapshot = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromSnapshot"] = nil
	return in
}
// SetFormat sets parameter Format to value and selects it for sending
func (in *ActionSnapshotDownloadCreateInput) SetFormat(value string) *ActionSnapshotDownloadCreateInput {
	in.Format = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Format"] = nil
	return in
}
// SetSendMail sets parameter SendMail to value and selects it for sending
func (in *ActionSnapshotDownloadCreateInput) SetSendMail(value bool) *ActionSnapshotDownloadCreateInput {
	in.SendMail = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SendMail"] = nil
	return in
}

// SelectParameters sets parameters from ActionSnapshotDownloadCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSnapshotDownloadCreateInput) SelectParameters(params ...string) *ActionSnapshotDownloadCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSnapshotDownloadCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSnapshotDownloadCreateRequest is a type for the entire action request
type ActionSnapshotDownloadCreateRequest struct {
	SnapshotDownload map[string]interface{} `json:"snapshot_download"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionSnapshotDownloadCreateOutput is a type for action output parameters
type ActionSnapshotDownloadCreateOutput struct {
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

// ActionSnapshotDownloadCreateMetaGlobalOutput is a type for global output metadata parameters
type ActionSnapshotDownloadCreateMetaGlobalOutput struct {
	ActionStateId int64 `json:"action_state_id"`
}

// Type for action response, including envelope
type ActionSnapshotDownloadCreateResponse struct {
	Action *ActionSnapshotDownloadCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SnapshotDownload *ActionSnapshotDownloadCreateOutput `json:"snapshot_download"`
		// Global output metadata
		Meta *ActionSnapshotDownloadCreateMetaGlobalOutput `json:"_meta"`
	}

	// Action output without the namespace
	Output *ActionSnapshotDownloadCreateOutput
}


// Prepare the action for invocation
func (action *ActionSnapshotDownloadCreate) Prepare() *ActionSnapshotDownloadCreateInvocation {
	return &ActionSnapshotDownloadCreateInvocation{
		Action: action,
		Path: "/v5.0/snapshot_downloads",
	}
}

// ActionSnapshotDownloadCreateInvocation is used to configure action for invocation
type ActionSnapshotDownloadCreateInvocation struct {
	// Pointer to the action
	Action *ActionSnapshotDownloadCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSnapshotDownloadCreateInput
	// Global meta input parameters
	MetaInput *ActionSnapshotDownloadCreateMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionSnapshotDownloadCreateInvocation) SetInput(input *ActionSnapshotDownloadCreateInput) *ActionSnapshotDownloadCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSnapshotDownloadCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSnapshotDownloadCreateInvocation) SetMetaInput(input *ActionSnapshotDownloadCreateMetaGlobalInput) *ActionSnapshotDownloadCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSnapshotDownloadCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSnapshotDownloadCreateInvocation) Call() (*ActionSnapshotDownloadCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionSnapshotDownloadCreateInvocation) callAsBody() (*ActionSnapshotDownloadCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSnapshotDownloadCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SnapshotDownload
	}
	return resp, err
}

// IsBlocking checks whether the current invocation resulted in a blocking operation
func (resp *ActionSnapshotDownloadCreateResponse) IsBlocking() bool {
	return resp.Response.Meta != nil && resp.Response.Meta.ActionStateId > 0
}

// OperationStatus queries the current state of the blocking operation
func (resp *ActionSnapshotDownloadCreateResponse) OperationStatus() (*ActionActionStateShowResponse, error) {
	req := resp.Action.Client.ActionState.Show.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}

// WaitForOperation waits for a blocking operation to finish
func (resp *ActionSnapshotDownloadCreateResponse) WaitForOperation(timeout float64) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	req.SetInput(&ActionActionStatePollInput{
		Timeout: timeout,
	})
	req.Input.SelectParameters("Timeout")
	return req.Call()
}

// WatchOperation waits for a blocking operation to finish and calls a callback
// function with progress updates
func (resp *ActionSnapshotDownloadCreateResponse) WatchOperation(timeout float64, updateIn float64, callback OperationProgressCallback) (*ActionActionStatePollResponse, error) {
	req := resp.Action.Client.ActionState.Poll.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	req.SetInput(&ActionActionStatePollInput{
		Timeout: timeout,
		UpdateIn: updateIn,
	})
	req.Input.SelectParameters("Timeout", "UpdateIn")
	pollResp, err := req.Call()

	if err != nil {
		return pollResp, err
	} else if pollResp.Output.Finished {
		return pollResp, nil
	}

	if callback(pollResp.Output) == StopWatching {
		return pollResp, nil
	}

	for {
		req = resp.Action.Client.ActionState.Poll.Prepare()
		req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
		req.SetInput(&ActionActionStatePollInput{
			Timeout: timeout,
			UpdateIn: updateIn,
			Status: pollResp.Output.Status,
			Current: pollResp.Output.Current,
			Total: pollResp.Output.Total,
		})
		pollResp, err = req.Call()

		if err != nil {
			return pollResp, err
		} else if pollResp.Output.Finished {
			return pollResp, nil
		}

		if callback(pollResp.Output) == StopWatching {
			return pollResp, nil
		}
	}
}

// CancelOperation cancels the current blocking operation
func (resp *ActionSnapshotDownloadCreateResponse) CancelOperation() (*ActionActionStateCancelResponse, error) {
	req := resp.Action.Client.ActionState.Cancel.Prepare()
	req.SetPathParamInt("action_state_id", resp.Response.Meta.ActionStateId)
	return req.Call()
}



func (inv *ActionSnapshotDownloadCreateInvocation) makeAllInputParams() *ActionSnapshotDownloadCreateRequest {
	return &ActionSnapshotDownloadCreateRequest{
		SnapshotDownload: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionSnapshotDownloadCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Snapshot") {
			ret["snapshot"] = inv.Input.Snapshot
		}
		if inv.IsParameterSelected("FromSnapshot") {
			ret["from_snapshot"] = inv.Input.FromSnapshot
		}
		if inv.IsParameterSelected("Format") {
			ret["format"] = inv.Input.Format
		}
		if inv.IsParameterSelected("SendMail") {
			ret["send_mail"] = inv.Input.SendMail
		}
	}

	return ret
}

func (inv *ActionSnapshotDownloadCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
