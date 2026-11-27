package client

import (
	"net/url"
	"strings"
)

// ActionDatasetPreviewSnapshotRemoval is a type for action Dataset#Preview_snapshot_removal
type ActionDatasetPreviewSnapshotRemoval struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetPreviewSnapshotRemoval(client *Client) *ActionDatasetPreviewSnapshotRemoval {
	return &ActionDatasetPreviewSnapshotRemoval{
		Client: client,
	}
}

// ActionDatasetPreviewSnapshotRemovalMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetPreviewSnapshotRemovalMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetPreviewSnapshotRemovalMetaGlobalInput) SetNo(value bool) *ActionDatasetPreviewSnapshotRemovalMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPreviewSnapshotRemovalMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPreviewSnapshotRemovalMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetPreviewSnapshotRemovalMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetPreviewSnapshotRemovalMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetPreviewSnapshotRemovalInput is a type for action input parameters
type ActionDatasetPreviewSnapshotRemovalInput struct {
	DisablePlans bool "json:\"disable_plans\""
	Recursive    bool "json:\"recursive\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDisablePlans sets parameter DisablePlans to value and selects it for sending
func (in *ActionDatasetPreviewSnapshotRemovalInput) SetDisablePlans(value bool) *ActionDatasetPreviewSnapshotRemovalInput {
	in.DisablePlans = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["DisablePlans"] = nil
	return in
}

// SetRecursive sets parameter Recursive to value and selects it for sending
func (in *ActionDatasetPreviewSnapshotRemovalInput) SetRecursive(value bool) *ActionDatasetPreviewSnapshotRemovalInput {
	in.Recursive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Recursive"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPreviewSnapshotRemovalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPreviewSnapshotRemovalInput) SelectParameters(params ...string) *ActionDatasetPreviewSnapshotRemovalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetPreviewSnapshotRemovalInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetPreviewSnapshotRemovalInput) UnselectParameters(params ...string) *ActionDatasetPreviewSnapshotRemovalInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetPreviewSnapshotRemovalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetPreviewSnapshotRemovalRequest is a type for the entire action request
type ActionDatasetPreviewSnapshotRemovalRequest struct {
	Dataset map[string]interface{} "json:\"dataset\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionDatasetPreviewSnapshotRemovalOutput is a type for action output parameters
type ActionDatasetPreviewSnapshotRemovalOutput struct {
	ManifestDigest  string "json:\"manifest_digest\""
	ManifestVersion int64  "json:\"manifest_version\""
}

// Type for action response, including envelope
type ActionDatasetPreviewSnapshotRemovalResponse struct {
	Action *ActionDatasetPreviewSnapshotRemoval "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Dataset *ActionDatasetPreviewSnapshotRemovalOutput "json:\"dataset\""
	}

	// Action output without the namespace
	Output *ActionDatasetPreviewSnapshotRemovalOutput
}

// Prepare the action for invocation
func (action *ActionDatasetPreviewSnapshotRemoval) Prepare() *ActionDatasetPreviewSnapshotRemovalInvocation {
	return &ActionDatasetPreviewSnapshotRemovalInvocation{
		Action: action,
		Path:   "/v7.0/datasets/{dataset_id}/preview_snapshot_removal",
	}
}

// ActionDatasetPreviewSnapshotRemovalInvocation is used to configure action for invocation
type ActionDatasetPreviewSnapshotRemovalInvocation struct {
	// Pointer to the action
	Action *ActionDatasetPreviewSnapshotRemoval

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetPreviewSnapshotRemovalInput
	// Global meta input parameters
	MetaInput *ActionDatasetPreviewSnapshotRemovalMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) SetPathParamInt(param string, value int64) *ActionDatasetPreviewSnapshotRemovalInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) SetPathParamString(param string, value string) *ActionDatasetPreviewSnapshotRemovalInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) NewInput() *ActionDatasetPreviewSnapshotRemovalInput {
	inv.Input = &ActionDatasetPreviewSnapshotRemovalInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) SetInput(input *ActionDatasetPreviewSnapshotRemovalInput) *ActionDatasetPreviewSnapshotRemovalInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) NewMetaInput() *ActionDatasetPreviewSnapshotRemovalMetaGlobalInput {
	inv.MetaInput = &ActionDatasetPreviewSnapshotRemovalMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) SetMetaInput(input *ActionDatasetPreviewSnapshotRemovalMetaGlobalInput) *ActionDatasetPreviewSnapshotRemovalInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) Call() (*ActionDatasetPreviewSnapshotRemovalResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) callAsBody() (*ActionDatasetPreviewSnapshotRemovalResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetPreviewSnapshotRemovalResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Dataset
	}
	return resp, err
}

func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) makeAllInputParams() *ActionDatasetPreviewSnapshotRemovalRequest {
	return &ActionDatasetPreviewSnapshotRemovalRequest{
		Dataset: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("DisablePlans") {
			ret["disable_plans"] = inv.Input.DisablePlans
		}
		if inv.IsParameterSelected("Recursive") {
			ret["recursive"] = inv.Input.Recursive
		}
	}

	return ret
}

func (inv *ActionDatasetPreviewSnapshotRemovalInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
