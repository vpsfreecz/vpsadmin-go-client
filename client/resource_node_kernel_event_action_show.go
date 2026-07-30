package client

import (
	"net/url"
	"strings"
)

// ActionNodeKernelEventShow is a type for action Node_kernel_event#Show
type ActionNodeKernelEventShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelEventShow(client *Client) *ActionNodeKernelEventShow {
	return &ActionNodeKernelEventShow{
		Client: client,
	}
}

// ActionNodeKernelEventShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelEventShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelEventShowMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelEventShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelEventShowMetaGlobalInput) SetNo(value bool) *ActionNodeKernelEventShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelEventShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelEventShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelEventShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelEventShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelEventShowOutput is a type for action output parameters
type ActionNodeKernelEventShowOutput struct {
	BootId                string                              "json:\"boot_id\""
	BootedAt              string                              "json:\"booted_at\""
	BootedRelease         string                              "json:\"booted_release\""
	BootedSystem          string                              "json:\"booted_system\""
	Confidence            string                              "json:\"confidence\""
	Current               bool                                "json:\"current\""
	CurrentSystem         string                              "json:\"current_system\""
	EffectiveAt           string                              "json:\"effective_at\""
	EventType             string                              "json:\"event_type\""
	EvidenceRevision      string                              "json:\"evidence_revision\""
	Id                    int64                               "json:\"id\""
	KernelCommandLine     string                              "json:\"kernel_command_line\""
	KernelConfigAvailable bool                                "json:\"kernel_config_available\""
	KernelConfigDigest    string                              "json:\"kernel_config_digest\""
	KernelSourceRevision  string                              "json:\"kernel_source_revision\""
	Node                  *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence    *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAfter         string                              "json:\"observed_after\""
	ObservedBefore        string                              "json:\"observed_before\""
	ReportSchemaVersion   int64                               "json:\"report_schema_version\""
	ReportedRelease       string                              "json:\"reported_release\""
	SnapshotRevision      string                              "json:\"snapshot_revision\""
	Source                string                              "json:\"source\""
	SourceStatusId        int64                               "json:\"source_status_id\""
}

// Type for action response, including envelope
type ActionNodeKernelEventShowResponse struct {
	Action *ActionNodeKernelEventShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelEvent *ActionNodeKernelEventShowOutput "json:\"node_kernel_event\""
	}

	// Action output without the namespace
	Output *ActionNodeKernelEventShowOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelEventShow) Prepare() *ActionNodeKernelEventShowInvocation {
	return &ActionNodeKernelEventShowInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_events/{node_kernel_event_id}",
	}
}

// ActionNodeKernelEventShowInvocation is used to configure action for invocation
type ActionNodeKernelEventShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelEventShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeKernelEventShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeKernelEventShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeKernelEventShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeKernelEventShowInvocation) SetPathParamString(param string, value string) *ActionNodeKernelEventShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelEventShowInvocation) NewMetaInput() *ActionNodeKernelEventShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelEventShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelEventShowInvocation) SetMetaInput(input *ActionNodeKernelEventShowMetaGlobalInput) *ActionNodeKernelEventShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelEventShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelEventShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelEventShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeKernelEventShowInvocation) Call() (*ActionNodeKernelEventShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelEventShowInvocation) callAsQuery() (*ActionNodeKernelEventShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeKernelEventShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelEvent
	}
	return resp, err
}

func (inv *ActionNodeKernelEventShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
