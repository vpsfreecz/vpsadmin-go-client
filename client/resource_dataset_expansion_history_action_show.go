package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionDatasetExpansionHistoryShow is a type for action Dataset_expansion.History#Show
type ActionDatasetExpansionHistoryShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetExpansionHistoryShow(client *Client) *ActionDatasetExpansionHistoryShow {
	return &ActionDatasetExpansionHistoryShow{
		Client: client,
	}
}

// ActionDatasetExpansionHistoryShowMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetExpansionHistoryShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetExpansionHistoryShowMetaGlobalInput) SetIncludes(value string) *ActionDatasetExpansionHistoryShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetExpansionHistoryShowMetaGlobalInput) SetNo(value bool) *ActionDatasetExpansionHistoryShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionHistoryShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionHistoryShowMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetExpansionHistoryShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetExpansionHistoryShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionHistoryShowOutput is a type for action output parameters
type ActionDatasetExpansionHistoryShowOutput struct {
	AddedSpace       int64                 "json:\"added_space\""
	Admin            *ActionUserShowOutput "json:\"admin\""
	CreatedAt        string                "json:\"created_at\""
	Id               int64                 "json:\"id\""
	NewRefquota      int64                 "json:\"new_refquota\""
	OriginalRefquota int64                 "json:\"original_refquota\""
}

// Type for action response, including envelope
type ActionDatasetExpansionHistoryShowResponse struct {
	Action *ActionDatasetExpansionHistoryShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		History *ActionDatasetExpansionHistoryShowOutput "json:\"history\""
	}

	// Action output without the namespace
	Output *ActionDatasetExpansionHistoryShowOutput
}

// Prepare the action for invocation
func (action *ActionDatasetExpansionHistoryShow) Prepare() *ActionDatasetExpansionHistoryShowInvocation {
	return &ActionDatasetExpansionHistoryShowInvocation{
		Action: action,
		Path:   "/v7.0/dataset_expansions/{dataset_expansion_id}/history/{history_id}",
	}
}

// ActionDatasetExpansionHistoryShowInvocation is used to configure action for invocation
type ActionDatasetExpansionHistoryShowInvocation struct {
	// Pointer to the action
	Action *ActionDatasetExpansionHistoryShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDatasetExpansionHistoryShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetExpansionHistoryShowInvocation) SetPathParamInt(param string, value int64) *ActionDatasetExpansionHistoryShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetExpansionHistoryShowInvocation) SetPathParamString(param string, value string) *ActionDatasetExpansionHistoryShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetExpansionHistoryShowInvocation) NewMetaInput() *ActionDatasetExpansionHistoryShowMetaGlobalInput {
	inv.MetaInput = &ActionDatasetExpansionHistoryShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetExpansionHistoryShowInvocation) SetMetaInput(input *ActionDatasetExpansionHistoryShowMetaGlobalInput) *ActionDatasetExpansionHistoryShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetExpansionHistoryShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetExpansionHistoryShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDatasetExpansionHistoryShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionDatasetExpansionHistoryShowInvocation) Call() (*ActionDatasetExpansionHistoryShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionDatasetExpansionHistoryShowInvocation) CallContext(ctx context.Context) (*ActionDatasetExpansionHistoryShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionDatasetExpansionHistoryShowResponse{Action: inv.Action}
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
		if resp.Response.History == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.History
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionDatasetExpansionHistoryShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionDatasetExpansionHistoryShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionDatasetExpansionHistoryShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
