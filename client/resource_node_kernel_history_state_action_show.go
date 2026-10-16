package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionNodeKernelHistoryStateShow is a type for action Node_kernel_history_state#Show
type ActionNodeKernelHistoryStateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNodeKernelHistoryStateShow(client *Client) *ActionNodeKernelHistoryStateShow {
	return &ActionNodeKernelHistoryStateShow{
		Client: client,
	}
}

// ActionNodeKernelHistoryStateShowMetaGlobalInput is a type for action global meta input parameters
type ActionNodeKernelHistoryStateShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodeKernelHistoryStateShowMetaGlobalInput) SetIncludes(value string) *ActionNodeKernelHistoryStateShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodeKernelHistoryStateShowMetaGlobalInput) SetNo(value bool) *ActionNodeKernelHistoryStateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodeKernelHistoryStateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodeKernelHistoryStateShowMetaGlobalInput) SelectParameters(params ...string) *ActionNodeKernelHistoryStateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodeKernelHistoryStateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodeKernelHistoryStateShowOutput is a type for action output parameters
type ActionNodeKernelHistoryStateShowOutput struct {
	CompletedAt     string                "json:\"completed_at\""
	FromStatusId    int64                 "json:\"from_status_id\""
	Id              int64                 "json:\"id\""
	Node            *ActionNodeShowOutput "json:\"node\""
	ObservedThrough string                "json:\"observed_through\""
	StartedAt       string                "json:\"started_at\""
	ThroughStatusId int64                 "json:\"through_status_id\""
}

// Type for action response, including envelope
type ActionNodeKernelHistoryStateShowResponse struct {
	Action *ActionNodeKernelHistoryStateShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeKernelHistoryState *ActionNodeKernelHistoryStateShowOutput "json:\"node_kernel_history_state\""
	}

	// Action output without the namespace
	Output *ActionNodeKernelHistoryStateShowOutput
}

// Prepare the action for invocation
func (action *ActionNodeKernelHistoryStateShow) Prepare() *ActionNodeKernelHistoryStateShowInvocation {
	return &ActionNodeKernelHistoryStateShowInvocation{
		Action: action,
		Path:   "/v7.0/node_kernel_history_states/{node_kernel_history_state_id}",
	}
}

// ActionNodeKernelHistoryStateShowInvocation is used to configure action for invocation
type ActionNodeKernelHistoryStateShowInvocation struct {
	// Pointer to the action
	Action *ActionNodeKernelHistoryStateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodeKernelHistoryStateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNodeKernelHistoryStateShowInvocation) SetPathParamInt(param string, value int64) *ActionNodeKernelHistoryStateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNodeKernelHistoryStateShowInvocation) SetPathParamString(param string, value string) *ActionNodeKernelHistoryStateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodeKernelHistoryStateShowInvocation) NewMetaInput() *ActionNodeKernelHistoryStateShowMetaGlobalInput {
	inv.MetaInput = &ActionNodeKernelHistoryStateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodeKernelHistoryStateShowInvocation) SetMetaInput(input *ActionNodeKernelHistoryStateShowMetaGlobalInput) *ActionNodeKernelHistoryStateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodeKernelHistoryStateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodeKernelHistoryStateShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNodeKernelHistoryStateShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionNodeKernelHistoryStateShowInvocation) Call() (*ActionNodeKernelHistoryStateShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionNodeKernelHistoryStateShowInvocation) CallContext(ctx context.Context) (*ActionNodeKernelHistoryStateShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionNodeKernelHistoryStateShowResponse{Action: inv.Action}
	raw, err := inv.Action.Client.doRaw(req)
	if err != nil {
		return resp, err
	}
	if err := decodeTypedResponse(raw.Body, resp); err != nil {
		return resp, err
	}
	if resp.Status {
		if resp.Response == nil {
			return resp, errMalformedResponse
		}
		if resp.Response.NodeKernelHistoryState == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.NodeKernelHistoryState
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionNodeKernelHistoryStateShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionNodeKernelHistoryStateShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionNodeKernelHistoryStateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
