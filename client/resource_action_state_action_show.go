package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionActionStateShow is a type for action Action_state#Show
type ActionActionStateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionActionStateShow(client *Client) *ActionActionStateShow {
	return &ActionActionStateShow{
		Client: client,
	}
}

// ActionActionStateShowMetaGlobalInput is a type for action global meta input parameters
type ActionActionStateShowMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionActionStateShowMetaGlobalInput) SetNo(value bool) *ActionActionStateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionActionStateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionActionStateShowMetaGlobalInput) SelectParameters(params ...string) *ActionActionStateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionActionStateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionActionStateShowOutput is a type for action output parameters
type ActionActionStateShowOutput struct {
	CanCancel bool   "json:\"can_cancel\""
	CreatedAt string "json:\"created_at\""
	Current   int64  "json:\"current\""
	Finished  bool   "json:\"finished\""
	Id        int64  "json:\"id\""
	Label     string "json:\"label\""
	Status    bool   "json:\"status\""
	Total     int64  "json:\"total\""
	Unit      string "json:\"unit\""
	UpdatedAt string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionActionStateShowResponse struct {
	Action *ActionActionStateShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ActionState *ActionActionStateShowOutput "json:\"action_state\""
	}

	// Action output without the namespace
	Output *ActionActionStateShowOutput
}

// Prepare the action for invocation
func (action *ActionActionStateShow) Prepare() *ActionActionStateShowInvocation {
	return &ActionActionStateShowInvocation{
		Action: action,
		Path:   "/v7.0/action_states/{action_state_id}",
	}
}

// ActionActionStateShowInvocation is used to configure action for invocation
type ActionActionStateShowInvocation struct {
	// Pointer to the action
	Action *ActionActionStateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionActionStateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionActionStateShowInvocation) SetPathParamInt(param string, value int64) *ActionActionStateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionActionStateShowInvocation) SetPathParamString(param string, value string) *ActionActionStateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionActionStateShowInvocation) NewMetaInput() *ActionActionStateShowMetaGlobalInput {
	inv.MetaInput = &ActionActionStateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionActionStateShowInvocation) SetMetaInput(input *ActionActionStateShowMetaGlobalInput) *ActionActionStateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionActionStateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionActionStateShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionActionStateShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionActionStateShowInvocation) Call() (*ActionActionStateShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionActionStateShowInvocation) CallContext(ctx context.Context) (*ActionActionStateShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionActionStateShowResponse{Action: inv.Action}
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
		if resp.Response.ActionState == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.ActionState
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionActionStateShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionActionStateShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionActionStateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
