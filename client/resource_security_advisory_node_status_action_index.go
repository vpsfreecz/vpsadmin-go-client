package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryNodeStatusIndex is a type for action Security_advisory.Node_status#Index
type ActionSecurityAdvisoryNodeStatusIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryNodeStatusIndex(client *Client) *ActionSecurityAdvisoryNodeStatusIndex {
	return &ActionSecurityAdvisoryNodeStatusIndex{
		Client: client,
	}
}

// ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput) SetCount(value bool) *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryNodeStatusIndexInput is a type for action input parameters
type ActionSecurityAdvisoryNodeStatusIndexInput struct {
	FromId int64  "json:\"from_id\""
	Limit  int64  "json:\"limit\""
	Node   int64  "json:\"node\""
	State  string "json:\"state\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusIndexInput) SetFromId(value int64) *ActionSecurityAdvisoryNodeStatusIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusIndexInput) SetLimit(value int64) *ActionSecurityAdvisoryNodeStatusIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusIndexInput) SetNode(value int64) *ActionSecurityAdvisoryNodeStatusIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusIndexInput) SetState(value string) *ActionSecurityAdvisoryNodeStatusIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryNodeStatusIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusIndexInput) SelectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryNodeStatusIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusIndexInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryNodeStatusIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryNodeStatusIndexOutput is a type for action output parameters
type ActionSecurityAdvisoryNodeStatusIndexOutput struct {
	Id               int64                             "json:\"id\""
	MitigatedSince   string                            "json:\"mitigated_since\""
	NodeId           int64                             "json:\"node_id\""
	NodeName         string                            "json:\"node_name\""
	Note             string                            "json:\"note\""
	SecurityAdvisory *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	State            string                            "json:\"state\""
	VulnerableUntil  string                            "json:\"vulnerable_until\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryNodeStatusIndexResponse struct {
	Action *ActionSecurityAdvisoryNodeStatusIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeStatuses []*ActionSecurityAdvisoryNodeStatusIndexOutput "json:\"node_statuses\""
	}

	// Action output without the namespace
	Output []*ActionSecurityAdvisoryNodeStatusIndexOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryNodeStatusIndex) Prepare() *ActionSecurityAdvisoryNodeStatusIndexInvocation {
	return &ActionSecurityAdvisoryNodeStatusIndexInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}/node_statuses",
	}
}

// ActionSecurityAdvisoryNodeStatusIndexInvocation is used to configure action for invocation
type ActionSecurityAdvisoryNodeStatusIndexInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryNodeStatusIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryNodeStatusIndexInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryNodeStatusIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryNodeStatusIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) NewInput() *ActionSecurityAdvisoryNodeStatusIndexInput {
	inv.Input = &ActionSecurityAdvisoryNodeStatusIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) SetInput(input *ActionSecurityAdvisoryNodeStatusIndexInput) *ActionSecurityAdvisoryNodeStatusIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) NewMetaInput() *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) SetMetaInput(input *ActionSecurityAdvisoryNodeStatusIndexMetaGlobalInput) *ActionSecurityAdvisoryNodeStatusIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) validate() error {
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
func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) Call() (*ActionSecurityAdvisoryNodeStatusIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) callAsQuery() (*ActionSecurityAdvisoryNodeStatusIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSecurityAdvisoryNodeStatusIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeStatuses
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["node_status[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["node_status[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["node_status[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("State") {
			ret["node_status[state]"] = inv.Input.State
		}
	}
}

func (inv *ActionSecurityAdvisoryNodeStatusIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
