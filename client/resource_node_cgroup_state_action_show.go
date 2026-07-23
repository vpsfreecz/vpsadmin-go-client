package client

import (
	"net/url"
	"strings"
)

// ActionNodeCgroupStateShow is a type for action Node_cgroup_state#Show
type ActionNodeCgroupStateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeCgroupStateShow(client *Client) *ActionNodeCgroupStateShow {
	return &ActionNodeCgroupStateShow{
		Client: client,
	}
}

// ActionNodeCgroupStateShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeCgroupStateShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeCgroupStateShowMetaGlobalInput) SetIncludes(value string) *ActionNodeCgroupStateShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeCgroupStateShowMetaGlobalInput) SetNo(value bool) *ActionNodeCgroupStateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeCgroupStateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeCgroupStateShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeCgroupStateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeCgroupStateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeCgroupStateShowOutput is a type for action output parameters
type ActionNodeCgroupStateShowOutput struct {
	CgroupVersion   string                "json:\"cgroup_version\""
	Current         bool                  "json:\"current\""
	FirstObservedAt string                "json:\"first_observed_at\""
	Id              int64                 "json:\"id\""
	LastObservedAt  string                "json:\"last_observed_at\""
	Node            *ActionNodeShowOutput "json:\"node\""
}

// Type for action response, including envelope
type ActionNodeCgroupStateShowResponse struct {
	Action *ActionNodeCgroupStateShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeCgroupState *ActionNodeCgroupStateShowOutput "json:\"node_cgroup_state\""
	}

	// Action output without the namespace
	Output *ActionNodeCgroupStateShowOutput
}

// Prepare the action for invocation
func (action *ActionNodeCgroupStateShow) Prepare() *ActionNodeCgroupStateShowInvocation {
	return &ActionNodeCgroupStateShowInvocation{
		Action: action,
		Path:   "/v7.0/node_cgroup_states/{node_cgroup_state_id}",
	}
}

// ActionNodeCgroupStateShowInvocation is used to configure action for invocation
type ActionNodeCgroupStateShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeCgroupStateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeCgroupStateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeCgroupStateShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeCgroupStateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeCgroupStateShowInvocation) SetPathParamString(param string, value string) *ActionNodeCgroupStateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeCgroupStateShowInvocation) NewMetaInput() *ActionNodeCgroupStateShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeCgroupStateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeCgroupStateShowInvocation) SetMetaInput(input *ActionNodeCgroupStateShowMetaGlobalInput) *ActionNodeCgroupStateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeCgroupStateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeCgroupStateShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeCgroupStateShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeCgroupStateShowInvocation) Call() (*ActionNodeCgroupStateShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeCgroupStateShowInvocation) callAsQuery() (*ActionNodeCgroupStateShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeCgroupStateShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeCgroupState
	}
	return resp, err
}

func (inv *ActionNodeCgroupStateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
