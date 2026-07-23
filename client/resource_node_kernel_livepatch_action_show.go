package client

import (
	"net/url"
	"strings"
)

// ActionNodeKernelLivepatchShow is a type for action Node_kernel_livepatch#Show
type ActionNodeKernelLivepatchShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelLivepatchShow(client *Client) *ActionNodeKernelLivepatchShow {
	return &ActionNodeKernelLivepatchShow{
		Client: client,
	}
}

// ActionNodeKernelLivepatchShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelLivepatchShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelLivepatchShowMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelLivepatchShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelLivepatchShowMetaGlobalInput) SetNo(value bool) *ActionNodeKernelLivepatchShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelLivepatchShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelLivepatchShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelLivepatchShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelLivepatchShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelLivepatchShowOutput is a type for action output parameters
type ActionNodeKernelLivepatchShowOutput struct {
	AppliedAt          string                              "json:\"applied_at\""
	Enabled            bool                                "json:\"enabled\""
	Id                 int64                               "json:\"id\""
	KernelVersion      string                              "json:\"kernel_version\""
	LivepatchId        string                              "json:\"livepatch_id\""
	Loaded             bool                                "json:\"loaded\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAt         string                              "json:\"observed_at\""
	PatchVersion       string                              "json:\"patch_version\""
	Source             string                              "json:\"source\""
	SourceRevision     string                              "json:\"source_revision\""
	Transition         bool                                "json:\"transition\""
	VerifiedAt         string                              "json:\"verified_at\""
}

// Type for action response, including envelope
type ActionNodeKernelLivepatchShowResponse struct {
	Action *ActionNodeKernelLivepatchShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelLivepatch *ActionNodeKernelLivepatchShowOutput "json:\"node_kernel_livepatch\""
	}

	// Action output without the namespace
	Output *ActionNodeKernelLivepatchShowOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelLivepatchShow) Prepare() *ActionNodeKernelLivepatchShowInvocation {
	return &ActionNodeKernelLivepatchShowInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_livepatches/{node_kernel_livepatch_id}",
	}
}

// ActionNodeKernelLivepatchShowInvocation is used to configure action for invocation
type ActionNodeKernelLivepatchShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelLivepatchShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeKernelLivepatchShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeKernelLivepatchShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeKernelLivepatchShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeKernelLivepatchShowInvocation) SetPathParamString(param string, value string) *ActionNodeKernelLivepatchShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelLivepatchShowInvocation) NewMetaInput() *ActionNodeKernelLivepatchShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelLivepatchShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelLivepatchShowInvocation) SetMetaInput(input *ActionNodeKernelLivepatchShowMetaGlobalInput) *ActionNodeKernelLivepatchShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelLivepatchShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelLivepatchShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelLivepatchShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeKernelLivepatchShowInvocation) Call() (*ActionNodeKernelLivepatchShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelLivepatchShowInvocation) callAsQuery() (*ActionNodeKernelLivepatchShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeKernelLivepatchShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelLivepatch
	}
	return resp, err
}

func (inv *ActionNodeKernelLivepatchShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
