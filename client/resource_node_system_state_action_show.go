package client

import (
	"net/url"
	"strings"
)

// ActionNodeSystemStateShow is a type for action Node_system_state#Show
type ActionNodeSystemStateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeSystemStateShow(client *Client) *ActionNodeSystemStateShow {
	return &ActionNodeSystemStateShow{
		Client: client,
	}
}

// ActionNodeSystemStateShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeSystemStateShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeSystemStateShowMetaGlobalInput) SetIncludes(value string) *ActionNodeSystemStateShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeSystemStateShowMetaGlobalInput) SetNo(value bool) *ActionNodeSystemStateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeSystemStateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeSystemStateShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeSystemStateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeSystemStateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeSystemStateShowOutput is a type for action output parameters
type ActionNodeSystemStateShowOutput struct {
	CgroupVersion   string                "json:\"cgroup_version\""
	Cpus            int64                 "json:\"cpus\""
	Current         bool                  "json:\"current\""
	FirstObservedAt string                "json:\"first_observed_at\""
	Id              int64                 "json:\"id\""
	LastObservedAt  string                "json:\"last_observed_at\""
	Node            *ActionNodeShowOutput "json:\"node\""
	TotalMemory     int64                 "json:\"total_memory\""
	TotalSwap       int64                 "json:\"total_swap\""
}

// Type for action response, including envelope
type ActionNodeSystemStateShowResponse struct {
	Action *ActionNodeSystemStateShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeSystemState *ActionNodeSystemStateShowOutput "json:\"node_system_state\""
	}

	// Action output without the namespace
	Output *ActionNodeSystemStateShowOutput
}

// Prepare the action for invocation
func (action *ActionNodeSystemStateShow) Prepare() *ActionNodeSystemStateShowInvocation {
	return &ActionNodeSystemStateShowInvocation{
		Action: action,
		Path:   "/v7.0/node_system_states/{node_system_state_id}",
	}
}

// ActionNodeSystemStateShowInvocation is used to configure action for invocation
type ActionNodeSystemStateShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeSystemStateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeSystemStateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeSystemStateShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeSystemStateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeSystemStateShowInvocation) SetPathParamString(param string, value string) *ActionNodeSystemStateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeSystemStateShowInvocation) NewMetaInput() *ActionNodeSystemStateShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeSystemStateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeSystemStateShowInvocation) SetMetaInput(input *ActionNodeSystemStateShowMetaGlobalInput) *ActionNodeSystemStateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeSystemStateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeSystemStateShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeSystemStateShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodeSystemStateShowInvocation) Call() (*ActionNodeSystemStateShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeSystemStateShowInvocation) callAsQuery() (*ActionNodeSystemStateShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNodeSystemStateShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeSystemState
	}
	return resp, err
}

func (inv *ActionNodeSystemStateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
