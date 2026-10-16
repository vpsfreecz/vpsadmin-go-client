package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpReleaseCampaignExempt is a type for action Ip_release_campaign#Exempt
type ActionIpReleaseCampaignExempt struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignExempt(client *Client) *ActionIpReleaseCampaignExempt {
	return &ActionIpReleaseCampaignExempt{
		Client: client,
	}
}

// ActionIpReleaseCampaignExemptMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignExemptMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignExemptMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignExemptMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignExemptMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignExemptMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignExemptMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignExemptMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignExemptMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignExemptMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignExemptInput is a type for action input parameters
type ActionIpReleaseCampaignExemptInput struct {
	Addresses json.RawMessage "json:\"addresses\""
	Reason    string          "json:\"reason\""
	Remove    bool            "json:\"remove\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddresses sets parameter Addresses to value and selects it for sending
func (in *ActionIpReleaseCampaignExemptInput) SetAddresses(value json.RawMessage) *ActionIpReleaseCampaignExemptInput {
	in.Addresses = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Addresses"] = nil
	return in
}

// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionIpReleaseCampaignExemptInput) SetReason(value string) *ActionIpReleaseCampaignExemptInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}

// SetRemove sets parameter Remove to value and selects it for sending
func (in *ActionIpReleaseCampaignExemptInput) SetRemove(value bool) *ActionIpReleaseCampaignExemptInput {
	in.Remove = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Remove"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignExemptInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignExemptInput) SelectParameters(params ...string) *ActionIpReleaseCampaignExemptInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseCampaignExemptInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignExemptInput) UnselectParameters(params ...string) *ActionIpReleaseCampaignExemptInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseCampaignExemptInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignExemptRequest is a type for the entire action request
type ActionIpReleaseCampaignExemptRequest struct {
	IpReleaseCampaign map[string]interface{} "json:\"ip_release_campaign\""
	Meta              map[string]interface{} "json:\"_meta\""
}

// ActionIpReleaseCampaignExemptOutput is a type for action output parameters
type ActionIpReleaseCampaignExemptOutput struct {
	AllowKeep bool                  "json:\"allow_keep\""
	ClosedAt  string                "json:\"closed_at\""
	ClosedBy  *ActionUserShowOutput "json:\"closed_by\""
	CreatedAt string                "json:\"created_at\""
	CreatedBy *ActionUserShowOutput "json:\"created_by\""
	Deadline  string                "json:\"deadline\""
	Id        int64                 "json:\"id\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignExemptResponse struct {
	Action *ActionIpReleaseCampaignExempt "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseCampaign *ActionIpReleaseCampaignExemptOutput "json:\"ip_release_campaign\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseCampaignExemptOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignExempt) Prepare() *ActionIpReleaseCampaignExemptInvocation {
	return &ActionIpReleaseCampaignExemptInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/{ip_release_campaign_id}/exempt",
	}
}

// ActionIpReleaseCampaignExemptInvocation is used to configure action for invocation
type ActionIpReleaseCampaignExemptInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignExempt

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseCampaignExemptInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignExemptMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseCampaignExemptInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseCampaignExemptInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseCampaignExemptInvocation) SetPathParamString(param string, value string) *ActionIpReleaseCampaignExemptInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseCampaignExemptInvocation) NewInput() *ActionIpReleaseCampaignExemptInput {
	inv.Input = &ActionIpReleaseCampaignExemptInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseCampaignExemptInvocation) SetInput(input *ActionIpReleaseCampaignExemptInput) *ActionIpReleaseCampaignExemptInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseCampaignExemptInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignExemptInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignExemptInvocation) NewMetaInput() *ActionIpReleaseCampaignExemptMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignExemptMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignExemptInvocation) SetMetaInput(input *ActionIpReleaseCampaignExemptMetaGlobalInput) *ActionIpReleaseCampaignExemptInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignExemptInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignExemptInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignExemptInvocation) validate() error {
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
func (inv *ActionIpReleaseCampaignExemptInvocation) Call() (*ActionIpReleaseCampaignExemptResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseCampaignExemptInvocation) CallContext(ctx context.Context) (*ActionIpReleaseCampaignExemptResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseCampaignExemptResponse{Action: inv.Action}
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
		if resp.Response.IpReleaseCampaign == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.IpReleaseCampaign
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpReleaseCampaignExemptInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseCampaignExemptInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "POST", inv.Path, input)
}

func (inv *ActionIpReleaseCampaignExemptInvocation) makeAllInputParams() *ActionIpReleaseCampaignExemptRequest {
	return &ActionIpReleaseCampaignExemptRequest{
		IpReleaseCampaign: inv.makeInputParams(),
		Meta:              inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpReleaseCampaignExemptInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Addresses") {
			ret["addresses"] = inv.Input.Addresses
		}
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
		if inv.IsParameterSelected("Remove") {
			ret["remove"] = inv.Input.Remove
		}
	}

	return ret
}

func (inv *ActionIpReleaseCampaignExemptInvocation) makeMetaInputParams() map[string]interface{} {
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
