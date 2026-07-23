package client

import (
	"net/url"
	"strings"
)

// ActionNodeKernelEvidenceShow is a type for action Node_kernel_evidence#Show
type ActionNodeKernelEvidenceShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelEvidenceShow(client *Client) *ActionNodeKernelEvidenceShow {
	return &ActionNodeKernelEvidenceShow{
		Client: client,
	}
}

// ActionNodeKernelEvidenceShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelEvidenceShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelEvidenceShowMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelEvidenceShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelEvidenceShowMetaGlobalInput) SetNo(value bool) *ActionNodeKernelEvidenceShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelEvidenceShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelEvidenceShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelEvidenceShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelEvidenceShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelEvidenceShowOutput is a type for action output parameters
type ActionNodeKernelEvidenceShowOutput struct {
	BootId                string                "json:\"boot_id\""
	BootedAt              string                "json:\"booted_at\""
	BootedRelease         string                "json:\"booted_release\""
	BootedSystem          string                "json:\"booted_system\""
	CurrentSystem         string                "json:\"current_system\""
	EvidenceRevision      string                "json:\"evidence_revision\""
	Id                    int64                 "json:\"id\""
	KernelCommandLine     string                "json:\"kernel_command_line\""
	KernelConfigAvailable bool                  "json:\"kernel_config_available\""
	KernelConfigDigest    string                "json:\"kernel_config_digest\""
	KernelSourceRevision  string                "json:\"kernel_source_revision\""
	Node                  *ActionNodeShowOutput "json:\"node\""
	ObservedAt            string                "json:\"observed_at\""
	ReceivedAt            string                "json:\"received_at\""
	ReportSchemaVersion   int64                 "json:\"report_schema_version\""
	ReportedRelease       string                "json:\"reported_release\""
	SnapshotRevision      string                "json:\"snapshot_revision\""
}

// Type for action response, including envelope
type ActionNodeKernelEvidenceShowResponse struct {
	Action *ActionNodeKernelEvidenceShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	}

	// Action output without the namespace
	Output *ActionNodeKernelEvidenceShowOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelEvidenceShow) Prepare() *ActionNodeKernelEvidenceShowInvocation {
	return &ActionNodeKernelEvidenceShowInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_evidences/{node_kernel_evidence_id}",
	}
}

// ActionNodeKernelEvidenceShowInvocation is used to configure action for invocation
type ActionNodeKernelEvidenceShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelEvidenceShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeKernelEvidenceShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeKernelEvidenceShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeKernelEvidenceShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeKernelEvidenceShowInvocation) SetPathParamString(param string, value string) *ActionNodeKernelEvidenceShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelEvidenceShowInvocation) NewMetaInput() *ActionNodeKernelEvidenceShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelEvidenceShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelEvidenceShowInvocation) SetMetaInput(input *ActionNodeKernelEvidenceShowMetaGlobalInput) *ActionNodeKernelEvidenceShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelEvidenceShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelEvidenceShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelEvidenceShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeKernelEvidenceShowInvocation) Call() (*ActionNodeKernelEvidenceShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelEvidenceShowInvocation) callAsQuery() (*ActionNodeKernelEvidenceShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeKernelEvidenceShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelEvidence
	}
	return resp, err
}

func (inv *ActionNodeKernelEvidenceShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
