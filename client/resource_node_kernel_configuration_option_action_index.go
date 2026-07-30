package client

import ()

// ActionNodeKernelConfigurationOptionIndex is a type for action Node_kernel_configuration_option#Index
type ActionNodeKernelConfigurationOptionIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelConfigurationOptionIndex(client *Client) *ActionNodeKernelConfigurationOptionIndex {
	return &ActionNodeKernelConfigurationOptionIndex{
		Client: client,
	}
}

// ActionNodeKernelConfigurationOptionIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelConfigurationOptionIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelConfigurationOptionIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelConfigurationOptionIndexInput is a type for action input parameters
type ActionNodeKernelConfigurationOptionIndexInput struct {
	ConfigurationDigest string "json:\"configuration_digest\""
	FromId              int64  "json:\"from_id\""
	Limit               int64  "json:\"limit\""
	Name                string "json:\"name\""
	Node                int64  "json:\"node\""
	NodeActive          bool   "json:\"node_active\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetConfigurationDigest sets parameter ConfigurationDigest to value and selects it for sending
func (in *ActionNodeKernelConfigurationOptionIndexInput) SetConfigurationDigest(value string) *ActionNodeKernelConfigurationOptionIndexInput {
	in.ConfigurationDigest = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ConfigurationDigest"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeKernelConfigurationOptionIndexInput) SetFromId(value int64) *ActionNodeKernelConfigurationOptionIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelConfigurationOptionIndexInput) SetLimit(value int64) *ActionNodeKernelConfigurationOptionIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionNodeKernelConfigurationOptionIndexInput) SetName(value string) *ActionNodeKernelConfigurationOptionIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeKernelConfigurationOptionIndexInput) SetNode(value int64) *ActionNodeKernelConfigurationOptionIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeActive sets parameter NodeActive to value and selects it for sending
func (in *ActionNodeKernelConfigurationOptionIndexInput) SetNodeActive(value bool) *ActionNodeKernelConfigurationOptionIndexInput {
	in.NodeActive = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NodeActive"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelConfigurationOptionIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelConfigurationOptionIndexInput) SelectParameters(params ...string) *ActionNodeKernelConfigurationOptionIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelConfigurationOptionIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelConfigurationOptionIndexInput) UnselectParameters(params ...string) *ActionNodeKernelConfigurationOptionIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelConfigurationOptionIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelConfigurationOptionIndexOutput is a type for action output parameters
type ActionNodeKernelConfigurationOptionIndexOutput struct {
	ConfigurationDigest string "json:\"configuration_digest\""
	Id                  int64  "json:\"id\""
	Name                string "json:\"name\""
	Value               string "json:\"value\""
}

// Type for action response, including envelope
type ActionNodeKernelConfigurationOptionIndexResponse struct {
	Action *ActionNodeKernelConfigurationOptionIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelConfigurationOptions []*ActionNodeKernelConfigurationOptionIndexOutput "json:\"node_kernel_configuration_options\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelConfigurationOptionIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelConfigurationOptionIndex) Prepare() *ActionNodeKernelConfigurationOptionIndexInvocation {
	return &ActionNodeKernelConfigurationOptionIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_configuration_options",
	}
}

// ActionNodeKernelConfigurationOptionIndexInvocation is used to configure action for invocation
type ActionNodeKernelConfigurationOptionIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelConfigurationOptionIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelConfigurationOptionIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) NewInput() *ActionNodeKernelConfigurationOptionIndexInput {
	inv.Input = &ActionNodeKernelConfigurationOptionIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) SetInput(input *ActionNodeKernelConfigurationOptionIndexInput) *ActionNodeKernelConfigurationOptionIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) NewMetaInput() *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelConfigurationOptionIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) SetMetaInput(input *ActionNodeKernelConfigurationOptionIndexMetaGlobalInput) *ActionNodeKernelConfigurationOptionIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
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
func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) Call() (*ActionNodeKernelConfigurationOptionIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) callAsQuery() (*ActionNodeKernelConfigurationOptionIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeKernelConfigurationOptionIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeKernelConfigurationOptions
	}
	return resp, err
}

func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("ConfigurationDigest") {
			ret["node_kernel_configuration_option[configuration_digest]"] = inv.Input.ConfigurationDigest
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_kernel_configuration_option[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_kernel_configuration_option[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["node_kernel_configuration_option[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["node_kernel_configuration_option[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("NodeActive") {
			ret["node_kernel_configuration_option[node_active]"] = convertBoolToString(inv.Input.NodeActive)
		}
	}

	return nil
}

func (inv *ActionNodeKernelConfigurationOptionIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
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
