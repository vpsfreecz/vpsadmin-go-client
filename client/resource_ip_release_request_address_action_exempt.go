package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpReleaseRequestAddressExempt is a type for action Ip_release_request.Address#Exempt
type ActionIpReleaseRequestAddressExempt struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseRequestAddressExempt(client *Client) *ActionIpReleaseRequestAddressExempt {
	return &ActionIpReleaseRequestAddressExempt{
		Client: client,
	}
}

// ActionIpReleaseRequestAddressExemptMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseRequestAddressExemptMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseRequestAddressExemptMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseRequestAddressExemptMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseRequestAddressExemptMetaGlobalInput) SetNo(value bool) *ActionIpReleaseRequestAddressExemptMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestAddressExemptMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestAddressExemptMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseRequestAddressExemptMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseRequestAddressExemptMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestAddressExemptInput is a type for action input parameters
type ActionIpReleaseRequestAddressExemptInput struct {
	Reason string "json:\"reason\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionIpReleaseRequestAddressExemptInput) SetReason(value string) *ActionIpReleaseRequestAddressExemptInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestAddressExemptInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestAddressExemptInput) SelectParameters(params ...string) *ActionIpReleaseRequestAddressExemptInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseRequestAddressExemptInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestAddressExemptInput) UnselectParameters(params ...string) *ActionIpReleaseRequestAddressExemptInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseRequestAddressExemptInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestAddressExemptRequest is a type for the entire action request
type ActionIpReleaseRequestAddressExemptRequest struct {
	Address map[string]interface{} "json:\"address\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionIpReleaseRequestAddressExemptOutput is a type for action output parameters
type ActionIpReleaseRequestAddressExemptOutput struct {
	Address           string                            "json:\"address\""
	AssignIpAddressId int64                             "json:\"assign_ip_address_id\""
	CleanupState      string                            "json:\"cleanup_state\""
	ExcludedAt        string                            "json:\"excluded_at\""
	ExclusionReason   string                            "json:\"exclusion_reason\""
	ExemptedAt        string                            "json:\"exempted_at\""
	ExemptedById      int64                             "json:\"exempted_by_id\""
	ExemptedByLogin   string                            "json:\"exempted_by_login\""
	ExemptionReason   string                            "json:\"exemption_reason\""
	Id                int64                             "json:\"id\""
	IpAddress         *ActionIpAddressShowOutput        "json:\"ip_address\""
	KeepReason        string                            "json:\"keep_reason\""
	KeptAt            string                            "json:\"kept_at\""
	KeptById          int64                             "json:\"kept_by_id\""
	KeptByLogin       string                            "json:\"kept_by_login\""
	LastResult        string                            "json:\"last_result\""
	LocationLabel     string                            "json:\"location_label\""
	OriginalUserId    int64                             "json:\"original_user_id\""
	Prefix            int64                             "json:\"prefix\""
	Protection        string                            "json:\"protection\""
	ReleaseChain      *ActionTransactionChainShowOutput "json:\"release_chain\""
	ReleasedAt        string                            "json:\"released_at\""
	Size              int64                             "json:\"size\""
	UserLogin         string                            "json:\"user_login\""
}

// Type for action response, including envelope
type ActionIpReleaseRequestAddressExemptResponse struct {
	Action *ActionIpReleaseRequestAddressExempt "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Address *ActionIpReleaseRequestAddressExemptOutput "json:\"address\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseRequestAddressExemptOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseRequestAddressExempt) Prepare() *ActionIpReleaseRequestAddressExemptInvocation {
	return &ActionIpReleaseRequestAddressExemptInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_requests/{ip_release_request_id}/addresses/{address_id}/exempt",
	}
}

// ActionIpReleaseRequestAddressExemptInvocation is used to configure action for invocation
type ActionIpReleaseRequestAddressExemptInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseRequestAddressExempt

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseRequestAddressExemptInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseRequestAddressExemptMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseRequestAddressExemptInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseRequestAddressExemptInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseRequestAddressExemptInvocation) SetPathParamString(param string, value string) *ActionIpReleaseRequestAddressExemptInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseRequestAddressExemptInvocation) NewInput() *ActionIpReleaseRequestAddressExemptInput {
	inv.Input = &ActionIpReleaseRequestAddressExemptInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseRequestAddressExemptInvocation) SetInput(input *ActionIpReleaseRequestAddressExemptInput) *ActionIpReleaseRequestAddressExemptInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseRequestAddressExemptInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestAddressExemptInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseRequestAddressExemptInvocation) NewMetaInput() *ActionIpReleaseRequestAddressExemptMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseRequestAddressExemptMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseRequestAddressExemptInvocation) SetMetaInput(input *ActionIpReleaseRequestAddressExemptMetaGlobalInput) *ActionIpReleaseRequestAddressExemptInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseRequestAddressExemptInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestAddressExemptInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseRequestAddressExemptInvocation) validate() error {
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
func (inv *ActionIpReleaseRequestAddressExemptInvocation) Call() (*ActionIpReleaseRequestAddressExemptResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseRequestAddressExemptInvocation) CallContext(ctx context.Context) (*ActionIpReleaseRequestAddressExemptResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseRequestAddressExemptResponse{Action: inv.Action}
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
		if resp.Response.Address == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Address
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpReleaseRequestAddressExemptInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseRequestAddressExemptInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "POST", inv.Path, input)
}

func (inv *ActionIpReleaseRequestAddressExemptInvocation) makeAllInputParams() *ActionIpReleaseRequestAddressExemptRequest {
	return &ActionIpReleaseRequestAddressExemptRequest{
		Address: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpReleaseRequestAddressExemptInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
	}

	return ret
}

func (inv *ActionIpReleaseRequestAddressExemptInvocation) makeMetaInputParams() map[string]interface{} {
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
