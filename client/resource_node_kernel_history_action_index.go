package client

import (
	"net/url"
	"strings"
)

// ActionNodeKernelHistoryIndex is a type for action Node.Kernel_history#Index
type ActionNodeKernelHistoryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelHistoryIndex(client *Client) *ActionNodeKernelHistoryIndex {
	return &ActionNodeKernelHistoryIndex{
		Client: client,
	}
}

// ActionNodeKernelHistoryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelHistoryIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNodeKernelHistoryIndexMetaGlobalInput) SetCount(value bool) *ActionNodeKernelHistoryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelHistoryIndexMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelHistoryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelHistoryIndexMetaGlobalInput) SetNo(value bool) *ActionNodeKernelHistoryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelHistoryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelHistoryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelHistoryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelHistoryIndexInput is a type for action input parameters
type ActionNodeKernelHistoryIndexInput struct {
	From   string "json:\"from\""
	FromId int64  "json:\"from_id\""
	Limit  int64  "json:\"limit\""
	To     string "json:\"to\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFrom sets parameter From to value and selects it for sending
func (in *ActionNodeKernelHistoryIndexInput) SetFrom(value string) *ActionNodeKernelHistoryIndexInput {
	in.From = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["From"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNodeKernelHistoryIndexInput) SetFromId(value int64) *ActionNodeKernelHistoryIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNodeKernelHistoryIndexInput) SetLimit(value int64) *ActionNodeKernelHistoryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetTo sets parameter To to value and selects it for sending
func (in *ActionNodeKernelHistoryIndexInput) SetTo(value string) *ActionNodeKernelHistoryIndexInput {
	in.To = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["To"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelHistoryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryIndexInput) SelectParameters(params ...string) *ActionNodeKernelHistoryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNodeKernelHistoryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryIndexInput) UnselectParameters(params ...string) *ActionNodeKernelHistoryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNodeKernelHistoryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelHistoryIndexOutput is a type for action output parameters
type ActionNodeKernelHistoryIndexOutput struct {
	BootedAt        string "json:\"booted_at\""
	BootedRelease   string "json:\"booted_release\""
	Confidence      string "json:\"confidence\""
	Current         bool   "json:\"current\""
	EffectiveAt     string "json:\"effective_at\""
	EventType       string "json:\"event_type\""
	Id              int64  "json:\"id\""
	ObservedAfter   string "json:\"observed_after\""
	ObservedBefore  string "json:\"observed_before\""
	ReportedRelease string "json:\"reported_release\""
	Source          string "json:\"source\""
}

// Type for action response, including envelope
type ActionNodeKernelHistoryIndexResponse struct {
	Action *ActionNodeKernelHistoryIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		KernelHistories []*ActionNodeKernelHistoryIndexOutput "json:\"kernel_histories\""
	}

	// Action output without the namespace
	Output []*ActionNodeKernelHistoryIndexOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelHistoryIndex) Prepare() *ActionNodeKernelHistoryIndexInvocation {
	return &ActionNodeKernelHistoryIndexInvocation{
		Action: action,
		Path:   "/v7.0/nodes/{node_id}/kernel_history",
	}
}

// ActionNodeKernelHistoryIndexInvocation is used to configure action for invocation
type ActionNodeKernelHistoryIndexInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelHistoryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNodeKernelHistoryIndexInput
	// Global meta input parameters
	MetaInput *ActionNodeKernelHistoryIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeKernelHistoryIndexInvocation) SetPathParamInt(param string, value int64) *ActionNodeKernelHistoryIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeKernelHistoryIndexInvocation) SetPathParamString(param string, value string) *ActionNodeKernelHistoryIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNodeKernelHistoryIndexInvocation) NewInput() *ActionNodeKernelHistoryIndexInput {
	inv.Input = &ActionNodeKernelHistoryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNodeKernelHistoryIndexInvocation) SetInput(input *ActionNodeKernelHistoryIndexInput) *ActionNodeKernelHistoryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNodeKernelHistoryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNodeKernelHistoryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelHistoryIndexInvocation) NewMetaInput() *ActionNodeKernelHistoryIndexMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelHistoryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelHistoryIndexInvocation) SetMetaInput(input *ActionNodeKernelHistoryIndexMetaGlobalInput) *ActionNodeKernelHistoryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelHistoryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelHistoryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelHistoryIndexInvocation) validate() error {
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
func (inv *ActionNodeKernelHistoryIndexInvocation) Call() (*ActionNodeKernelHistoryIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNodeKernelHistoryIndexInvocation) callAsQuery() (*ActionNodeKernelHistoryIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodeKernelHistoryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.KernelHistories
	}
	return resp, err
}

func (inv *ActionNodeKernelHistoryIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("From") {
			ret["kernel_history[from]"] = inv.Input.From
		}
		if inv.IsParameterSelected("FromId") {
			ret["kernel_history[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["kernel_history[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("To") {
			ret["kernel_history[to]"] = inv.Input.To
		}
	}
}

func (inv *ActionNodeKernelHistoryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
