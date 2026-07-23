package client

import ()

// ActionNodeSoftwareVersionIndex is a type for action Node_software_version#Index
type ActionNodeSoftwareVersionIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeSoftwareVersionIndex(client *Client) *ActionNodeSoftwareVersionIndex {
	return &ActionNodeSoftwareVersionIndex{
		Client: client,
	}
}

// ActionNodeSoftwareVersionIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeSoftwareVersionIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexMetaGlobalInput) SetCount(value bool) *ActionNodeSoftwareVersionIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeSoftwareVersionIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexMetaGlobalInput) SetNo(value bool) *ActionNodeSoftwareVersionIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSoftwareVersionIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSoftwareVersionIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeSoftwareVersionIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeSoftwareVersionIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSoftwareVersionIndexInput is a type for action input parameters
type ActionNodeSoftwareVersionIndexInput struct {
	Component          string "json:\"component\""
	From               string "json:\"from\""
	FromId             int64  "json:\"from_id\""
	Generation         string "json:\"generation\""
	Limit              int64  "json:\"limit\""
	Node               int64  "json:\"node\""
	NodeActive         bool   "json:\"node_active\""
	NodeKernelEvidence int64  "json:\"node_kernel_evidence\""
	Revision           string "json:\"revision\""
	RevisionDirty      bool   "json:\"revision_dirty\""
	RevisionSource     string "json:\"revision_source\""
	Source             string "json:\"source\""
	To                 string "json:\"to\""
	Version            string "json:\"version\""
	VersionSource      string "json:\"version_source\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetComponent sets parameter Component to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetComponent(value string) *ActionNodeSoftwareVersionIndexInput {
	in.Component = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Component"] = nil
	return in
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetFrom(value string) *ActionNodeSoftwareVersionIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetFromId(value int64) *ActionNodeSoftwareVersionIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetGeneration sets parameter Generation to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetGeneration(value string) *ActionNodeSoftwareVersionIndexInput {
	in.Generation = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Generation"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetLimit(value int64) *ActionNodeSoftwareVersionIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetNode(value int64) *ActionNodeSoftwareVersionIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetNodeActive(value bool) *ActionNodeSoftwareVersionIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SetNodeKernelEvidence sets parameter NodeKernelEvidence to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetNodeKernelEvidence(value int64) *ActionNodeSoftwareVersionIndexInput {
	in.NodeKernelEvidence = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeKernelEvidence"] = nil
	return in
}

// SetRevision sets parameter Revision to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetRevision(value string) *ActionNodeSoftwareVersionIndexInput {
	in.Revision = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Revision"] = nil
	return in
}

// SetRevisionDirty sets parameter RevisionDirty to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetRevisionDirty(value bool) *ActionNodeSoftwareVersionIndexInput {
	in.RevisionDirty = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RevisionDirty"] = nil
	return in
}

// SetRevisionSource sets parameter RevisionSource to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetRevisionSource(value string) *ActionNodeSoftwareVersionIndexInput {
	in.RevisionSource = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RevisionSource"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetSource(value string) *ActionNodeSoftwareVersionIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetTo(value string) *ActionNodeSoftwareVersionIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SetVersion sets parameter Version to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetVersion(value string) *ActionNodeSoftwareVersionIndexInput {
	in.Version = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Version"] = nil
	return in
}

// SetVersionSource sets parameter VersionSource to value and selects it for sending
func (in *ActionNodeSoftwareVersionIndexInput) SetVersionSource(value string) *ActionNodeSoftwareVersionIndexInput {
	in.VersionSource = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VersionSource"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSoftwareVersionIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSoftwareVersionIndexInput) SelectParameters(params ...string) *ActionNodeSoftwareVersionIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeSoftwareVersionIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeSoftwareVersionIndexInput) UnselectParameters(params ...string) *ActionNodeSoftwareVersionIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeSoftwareVersionIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSoftwareVersionIndexOutput is a type for action output parameters
type ActionNodeSoftwareVersionIndexOutput struct {
	Component          string                              "json:\"component\""
	Generation         string                              "json:\"generation\""
	Id                 int64                               "json:\"id\""
	Node               *ActionNodeShowOutput               "json:\"node\""
	NodeKernelEvidence *ActionNodeKernelEvidenceShowOutput "json:\"node_kernel_evidence\""
	ObservedAt         string                              "json:\"observed_at\""
	Revision           string                              "json:\"revision\""
	RevisionDirty      bool                                "json:\"revision_dirty\""
	RevisionSource     string                              "json:\"revision_source\""
	Source             string                              "json:\"source\""
	SourceRevision     string                              "json:\"source_revision\""
	Version            string                              "json:\"version\""
	VersionSource      string                              "json:\"version_source\""
}

// Type for action response, including envelope
type ActionNodeSoftwareVersionIndexResponse struct {
	Action *ActionNodeSoftwareVersionIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeSoftwareVersions []*ActionNodeSoftwareVersionIndexOutput "json:\"node_software_versions\""
	}

	// Action output without the namespace
	Output []*ActionNodeSoftwareVersionIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeSoftwareVersionIndex) Prepare() *ActionNodeSoftwareVersionIndexInvocation {
	return &ActionNodeSoftwareVersionIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_software_versions",
	}
}

// ActionNodeSoftwareVersionIndexInvocation is used to configure action for invocation
type ActionNodeSoftwareVersionIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeSoftwareVersionIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeSoftwareVersionIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeSoftwareVersionIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeSoftwareVersionIndexInvocation) NewInput() *ActionNodeSoftwareVersionIndexInput {
	inv.Input = &ActionNodeSoftwareVersionIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeSoftwareVersionIndexInvocation) SetInput(input *ActionNodeSoftwareVersionIndexInput) *ActionNodeSoftwareVersionIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeSoftwareVersionIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeSoftwareVersionIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeSoftwareVersionIndexInvocation) NewMetaInput() *ActionNodeSoftwareVersionIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeSoftwareVersionIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeSoftwareVersionIndexInvocation) SetMetaInput(input *ActionNodeSoftwareVersionIndexMetaGlobalInput) *ActionNodeSoftwareVersionIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeSoftwareVersionIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeSoftwareVersionIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeSoftwareVersionIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			if !inv.IsParameterNil("From") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.From)
				if !ok {
					verr.Add("from", "not a valid datetime")
				} else {
					inv.Input.From = normalized
				}
			}
		}
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			if !inv.IsParameterNil("NodeKernelEvidence") {
				if inv.Input.NodeKernelEvidence < 0 {
					verr.Add("node_kernel_evidence", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("To") {
			if !inv.IsParameterNil("To") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.To)
				if !ok {
					verr.Add("to", "not a valid datetime")
				} else {
					inv.Input.To = normalized
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeSoftwareVersionIndexInvocation) Call() (*ActionNodeSoftwareVersionIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeSoftwareVersionIndexInvocation) callAsQuery() (*ActionNodeSoftwareVersionIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeSoftwareVersionIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeSoftwareVersions
	}
	return resp, err
}

func (inv *ActionNodeSoftwareVersionIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Component") {
			ret["node_software_version[component]"] = inv.Input.Component
		}
		if inv.IsParameterSelected("From") {
			ret["node_software_version[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_software_version[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Generation") {
			ret["node_software_version[generation]"] = inv.Input.Generation
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_software_version[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_software_version[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_software_version[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
		if inv.IsParameterSelected("NodeKernelEvidence") {
			ret["node_software_version[node_kernel_evidence]"] = convertInt64ToString(inv.Input.NodeKernelEvidence)
		}
		if inv.IsParameterSelected("Revision") {
			ret["node_software_version[revision]"] = inv.Input.Revision
		}
		if inv.IsParameterSelected("RevisionDirty") {
			ret["node_software_version[revision_dirty]"] = convertBoolToString(inv.Input.RevisionDirty)
		}
		if inv.IsParameterSelected("RevisionSource") {
			ret["node_software_version[revision_source]"] = inv.Input.RevisionSource
		}
		if inv.IsParameterSelected("Source") {
			ret["node_software_version[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("To") {
			ret["node_software_version[to]"] = inv.Input.To
		}
		if inv.IsParameterSelected("Version") {
			ret["node_software_version[version]"] = inv.Input.Version
		}
		if inv.IsParameterSelected("VersionSource") {
			ret["node_software_version[version_source]"] = inv.Input.VersionSource
		}
	}
}

func (inv *ActionNodeSoftwareVersionIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
