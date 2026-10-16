package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpReleaseCampaignClose is a type for action Ip_release_campaign#Close
type ActionIpReleaseCampaignClose struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignClose(client *Client) *ActionIpReleaseCampaignClose {
	return &ActionIpReleaseCampaignClose{
		Client: client,
	}
}

// ActionIpReleaseCampaignCloseMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignCloseMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignCloseMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignCloseMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignCloseMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignCloseMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignCloseMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignCloseMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignCloseMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignCloseMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignCloseRequest is a type for the entire action request
type ActionIpReleaseCampaignCloseRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// ActionIpReleaseCampaignCloseOutput is a type for action output parameters
type ActionIpReleaseCampaignCloseOutput struct {
	AllowKeep bool                  "json:\"allow_keep\""
	ClosedAt  string                "json:\"closed_at\""
	ClosedBy  *ActionUserShowOutput "json:\"closed_by\""
	CreatedAt string                "json:\"created_at\""
	CreatedBy *ActionUserShowOutput "json:\"created_by\""
	Deadline  string                "json:\"deadline\""
	Id        int64                 "json:\"id\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignCloseResponse struct {
	Action *ActionIpReleaseCampaignClose "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseCampaign *ActionIpReleaseCampaignCloseOutput "json:\"ip_release_campaign\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseCampaignCloseOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignClose) Prepare() *ActionIpReleaseCampaignCloseInvocation {
	return &ActionIpReleaseCampaignCloseInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/{ip_release_campaign_id}/close",
	}
}

// ActionIpReleaseCampaignCloseInvocation is used to configure action for invocation
type ActionIpReleaseCampaignCloseInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignClose

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignCloseMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseCampaignCloseInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseCampaignCloseInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseCampaignCloseInvocation) SetPathParamString(param string, value string) *ActionIpReleaseCampaignCloseInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignCloseInvocation) NewMetaInput() *ActionIpReleaseCampaignCloseMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignCloseMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignCloseInvocation) SetMetaInput(input *ActionIpReleaseCampaignCloseMetaGlobalInput) *ActionIpReleaseCampaignCloseInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignCloseInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignCloseInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignCloseInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionIpReleaseCampaignCloseInvocation) Call() (*ActionIpReleaseCampaignCloseResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseCampaignCloseInvocation) CallContext(ctx context.Context) (*ActionIpReleaseCampaignCloseResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseCampaignCloseResponse{Action: inv.Action}
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
func (inv *ActionIpReleaseCampaignCloseInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseCampaignCloseInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "POST", inv.Path, input)
}

func (inv *ActionIpReleaseCampaignCloseInvocation) makeAllInputParams() *ActionIpReleaseCampaignCloseRequest {
	return &ActionIpReleaseCampaignCloseRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionIpReleaseCampaignCloseInvocation) makeMetaInputParams() map[string]interface{} {
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
