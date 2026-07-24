package client

import (
	"net/url"
	"strings"
)

// ActionNodeEbpfProgramShow is a type for action Node_ebpf_program#Show
type ActionNodeEbpfProgramShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeEbpfProgramShow(client *Client) *ActionNodeEbpfProgramShow {
	return &ActionNodeEbpfProgramShow{
		Client: client,
	}
}

// ActionNodeEbpfProgramShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeEbpfProgramShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeEbpfProgramShowMetaGlobalInput) SetIncludes(value string) *ActionNodeEbpfProgramShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeEbpfProgramShowMetaGlobalInput) SetNo(value bool) *ActionNodeEbpfProgramShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeEbpfProgramShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeEbpfProgramShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeEbpfProgramShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeEbpfProgramShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeEbpfProgramShowOutput is a type for action output parameters
type ActionNodeEbpfProgramShowOutput struct {
	Active             bool                                "json:\"active\""
	AttachedAt         string                              "json:\"attached_at\""
	Description        string                              "json:\"description\""
	Digest             string                              "json:\"digest\""
	Id                 int64                               "json:\"id\""
	Name               string                              "json:\"name\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAt         string                              "json:\"observed_at\""
	Revision           string                              "json:\"revision\""
	SinceKernel        string                              "json:\"since_kernel\""
	Source             string                              "json:\"source\""
	SourceRevision     string                              "json:\"source_revision\""
	UntilKernel        string                              "json:\"until_kernel\""
	VerifiedAt         string                              "json:\"verified_at\""
}

// Type for action response, including envelope
type ActionNodeEbpfProgramShowResponse struct {
	Action *ActionNodeEbpfProgramShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeEbpfProgram *ActionNodeEbpfProgramShowOutput "json:\"node_ebpf_program\""
	}

	// Action output without the namespace
	Output *ActionNodeEbpfProgramShowOutput
}

// Prepare the action for invocation
func (action *ActionNodeEbpfProgramShow) Prepare() *ActionNodeEbpfProgramShowInvocation {
	return &ActionNodeEbpfProgramShowInvocation{
		Action: action,
		Path:   "/v7.0/node_ebpf_programs/{node_ebpf_program_id}",
	}
}

// ActionNodeEbpfProgramShowInvocation is used to configure action for invocation
type ActionNodeEbpfProgramShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeEbpfProgramShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeEbpfProgramShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeEbpfProgramShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeEbpfProgramShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeEbpfProgramShowInvocation) SetPathParamString(param string, value string) *ActionNodeEbpfProgramShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeEbpfProgramShowInvocation) NewMetaInput() *ActionNodeEbpfProgramShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeEbpfProgramShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeEbpfProgramShowInvocation) SetMetaInput(input *ActionNodeEbpfProgramShowMetaGlobalInput) *ActionNodeEbpfProgramShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeEbpfProgramShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeEbpfProgramShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeEbpfProgramShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeEbpfProgramShowInvocation) Call() (*ActionNodeEbpfProgramShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeEbpfProgramShowInvocation) callAsQuery() (*ActionNodeEbpfProgramShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeEbpfProgramShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeEbpfProgram
	}
	return resp, err
}

func (inv *ActionNodeEbpfProgramShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
