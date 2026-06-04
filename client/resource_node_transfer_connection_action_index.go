package client

import ()

// ActionNodeTransferConnectionIndex is a type for action Node_transfer_connection#Index
type ActionNodeTransferConnectionIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeTransferConnectionIndex(client *Client) *ActionNodeTransferConnectionIndex {
	return &ActionNodeTransferConnectionIndex{
		Client: client,
	}
}

// ActionNodeTransferConnectionIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeTransferConnectionIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeTransferConnectionIndexMetaGlobalInput) SetCount(value bool) *ActionNodeTransferConnectionIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeTransferConnectionIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeTransferConnectionIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeTransferConnectionIndexMetaGlobalInput) SetNo(value bool) *ActionNodeTransferConnectionIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeTransferConnectionIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeTransferConnectionIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeTransferConnectionIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeTransferConnectionIndexInput is a type for action input parameters
type ActionNodeTransferConnectionIndexInput struct {
	Enabled bool  "json:\"enabled\""
	FromId  int64 "json:\"from_id\""
	Limit   int64 "json:\"limit\""
	Node    int64 "json:\"node\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionNodeTransferConnectionIndexInput) SetEnabled(value bool) *ActionNodeTransferConnectionIndexInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeTransferConnectionIndexInput) SetFromId(value int64) *ActionNodeTransferConnectionIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeTransferConnectionIndexInput) SetLimit(value int64) *ActionNodeTransferConnectionIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionNodeTransferConnectionIndexInput) SetNode(value int64) *ActionNodeTransferConnectionIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeTransferConnectionIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionIndexInput) SelectParameters(params ...string) *ActionNodeTransferConnectionIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeTransferConnectionIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeTransferConnectionIndexInput) UnselectParameters(params ...string) *ActionNodeTransferConnectionIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeTransferConnectionIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeTransferConnectionIndexOutput is a type for action output parameters
type ActionNodeTransferConnectionIndexOutput struct {
	CreatedAt   string                "json:\"created_at\""
	Enabled     bool                  "json:\"enabled\""
	Id          int64                 "json:\"id\""
	NodeA       *ActionNodeShowOutput "json:\"node_a\""
	NodeAIpAddr string                "json:\"node_a_ip_addr\""
	NodeB       *ActionNodeShowOutput "json:\"node_b\""
	NodeBIpAddr string                "json:\"node_b_ip_addr\""
	UpdatedAt   string                "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionNodeTransferConnectionIndexResponse struct {
	Action *ActionNodeTransferConnectionIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeTransferConnections []*ActionNodeTransferConnectionIndexOutput "json:\"node_transfer_connections\""
	}

	// Action output without the namespace
	Output []*ActionNodeTransferConnectionIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeTransferConnectionIndex) Prepare() *ActionNodeTransferConnectionIndexInvocation {
	return &ActionNodeTransferConnectionIndexInvocation{
		Action: action,
		Path:   "/v7.0/node_transfer_connections",
	}
}

// ActionNodeTransferConnectionIndexInvocation is used to configure action for invocation
type ActionNodeTransferConnectionIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeTransferConnectionIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeTransferConnectionIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeTransferConnectionIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeTransferConnectionIndexInvocation) NewInput() *ActionNodeTransferConnectionIndexInput {
	inv.Input = &ActionNodeTransferConnectionIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeTransferConnectionIndexInvocation) SetInput(input *ActionNodeTransferConnectionIndexInput) *ActionNodeTransferConnectionIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeTransferConnectionIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeTransferConnectionIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeTransferConnectionIndexInvocation) NewMetaInput() *ActionNodeTransferConnectionIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeTransferConnectionIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeTransferConnectionIndexInvocation) SetMetaInput(input *ActionNodeTransferConnectionIndexMetaGlobalInput) *ActionNodeTransferConnectionIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeTransferConnectionIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeTransferConnectionIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeTransferConnectionIndexInvocation) validate() error {
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
func (inv *ActionNodeTransferConnectionIndexInvocation) Call() (*ActionNodeTransferConnectionIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeTransferConnectionIndexInvocation) callAsQuery() (*ActionNodeTransferConnectionIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeTransferConnectionIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeTransferConnections
	}
	return resp, err
}

func (inv *ActionNodeTransferConnectionIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["node_transfer_connection[enabled]"] = convertBoolToString(inv.Input.Enabled)
		}
		if inv.IsParameterSelected("FromId") {
			ret["node_transfer_connection[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_transfer_connection[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_transfer_connection[node]"] = convertInt64ToString(inv.Input.Node)
		}
	}
}

func (inv *ActionNodeTransferConnectionIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
