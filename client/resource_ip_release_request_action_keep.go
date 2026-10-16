package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpReleaseRequestKeep is a type for action Ip_release_request#Keep
type ActionIpReleaseRequestKeep struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseRequestKeep(client *Client) *ActionIpReleaseRequestKeep {
	return &ActionIpReleaseRequestKeep{
		Client: client,
	}
}

// ActionIpReleaseRequestKeepMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseRequestKeepMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseRequestKeepMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseRequestKeepMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseRequestKeepMetaGlobalInput) SetNo(value bool) *ActionIpReleaseRequestKeepMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestKeepMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestKeepMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseRequestKeepMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseRequestKeepMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestKeepInput is a type for action input parameters
type ActionIpReleaseRequestKeepInput struct {
	Addresses json.RawMessage "json:\"addresses\""
	Reason    string          "json:\"reason\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddresses sets parameter Addresses to value and selects it for sending
func (in *ActionIpReleaseRequestKeepInput) SetAddresses(value json.RawMessage) *ActionIpReleaseRequestKeepInput {
	in.Addresses = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Addresses"] = nil
	return in
}

// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionIpReleaseRequestKeepInput) SetReason(value string) *ActionIpReleaseRequestKeepInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestKeepInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestKeepInput) SelectParameters(params ...string) *ActionIpReleaseRequestKeepInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseRequestKeepInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestKeepInput) UnselectParameters(params ...string) *ActionIpReleaseRequestKeepInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseRequestKeepInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestKeepRequest is a type for the entire action request
type ActionIpReleaseRequestKeepRequest struct {
	IpReleaseRequest map[string]interface{} "json:\"ip_release_request\""
	Meta             map[string]interface{} "json:\"_meta\""
}

// ActionIpReleaseRequestKeepOutput is a type for action output parameters
type ActionIpReleaseRequestKeepOutput struct {
	CanAssign bool   "json:\"can_assign\""
	CanKeep   bool   "json:\"can_keep\""
	Deadline  string "json:\"deadline\""
	Id        int64  "json:\"id\""
}

// Type for action response, including envelope
type ActionIpReleaseRequestKeepResponse struct {
	Action *ActionIpReleaseRequestKeep "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseRequest *ActionIpReleaseRequestKeepOutput "json:\"ip_release_request\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseRequestKeepOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseRequestKeep) Prepare() *ActionIpReleaseRequestKeepInvocation {
	return &ActionIpReleaseRequestKeepInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_requests/{ip_release_request_id}/keep",
	}
}

// ActionIpReleaseRequestKeepInvocation is used to configure action for invocation
type ActionIpReleaseRequestKeepInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseRequestKeep

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseRequestKeepInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseRequestKeepMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseRequestKeepInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseRequestKeepInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseRequestKeepInvocation) SetPathParamString(param string, value string) *ActionIpReleaseRequestKeepInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseRequestKeepInvocation) NewInput() *ActionIpReleaseRequestKeepInput {
	inv.Input = &ActionIpReleaseRequestKeepInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseRequestKeepInvocation) SetInput(input *ActionIpReleaseRequestKeepInput) *ActionIpReleaseRequestKeepInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseRequestKeepInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestKeepInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseRequestKeepInvocation) NewMetaInput() *ActionIpReleaseRequestKeepMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseRequestKeepMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseRequestKeepInvocation) SetMetaInput(input *ActionIpReleaseRequestKeepMetaGlobalInput) *ActionIpReleaseRequestKeepInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseRequestKeepInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestKeepInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseRequestKeepInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionIpReleaseRequestKeepInvocation) Call() (*ActionIpReleaseRequestKeepResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseRequestKeepInvocation) CallContext(ctx context.Context) (*ActionIpReleaseRequestKeepResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseRequestKeepResponse{Action: inv.Action}
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
		if resp.Response.IpReleaseRequest == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.IpReleaseRequest
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpReleaseRequestKeepInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseRequestKeepInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "POST", inv.Path, input)
}

func (inv *ActionIpReleaseRequestKeepInvocation) makeAllInputParams() *ActionIpReleaseRequestKeepRequest {
	return &ActionIpReleaseRequestKeepRequest{
		IpReleaseRequest: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpReleaseRequestKeepInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Addresses") {
			ret["addresses"] = inv.Input.Addresses
		}
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
	}

	return ret
}

func (inv *ActionIpReleaseRequestKeepInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
