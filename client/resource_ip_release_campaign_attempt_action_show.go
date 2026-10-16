package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpReleaseCampaignAttemptShow is a type for action Ip_release_campaign.Attempt#Show
type ActionIpReleaseCampaignAttemptShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignAttemptShow(client *Client) *ActionIpReleaseCampaignAttemptShow {
	return &ActionIpReleaseCampaignAttemptShow{
		Client: client,
	}
}

// ActionIpReleaseCampaignAttemptShowMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignAttemptShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignAttemptShowMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignAttemptShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignAttemptShowMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignAttemptShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignAttemptShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignAttemptShowMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignAttemptShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignAttemptShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignAttemptShowOutput is a type for action output parameters
type ActionIpReleaseCampaignAttemptShowOutput struct {
	BlockedChainId    int64                             "json:\"blocked_chain_id\""
	BlockedResource   string                            "json:\"blocked_resource\""
	BlockedResourceId int64                             "json:\"blocked_resource_id\""
	CreatedAt         string                            "json:\"created_at\""
	CreatedById       int64                             "json:\"created_by_id\""
	CreatedByLogin    string                            "json:\"created_by_login\""
	Error             string                            "json:\"error\""
	Id                int64                             "json:\"id\""
	IpCount           int64                             "json:\"ip_count\""
	State             string                            "json:\"state\""
	TransactionChain  *ActionTransactionChainShowOutput "json:\"transaction_chain\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignAttemptShowResponse struct {
	Action *ActionIpReleaseCampaignAttemptShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Attempt *ActionIpReleaseCampaignAttemptShowOutput "json:\"attempt\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseCampaignAttemptShowOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignAttemptShow) Prepare() *ActionIpReleaseCampaignAttemptShowInvocation {
	return &ActionIpReleaseCampaignAttemptShowInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/{ip_release_campaign_id}/attempts/{attempt_id}",
	}
}

// ActionIpReleaseCampaignAttemptShowInvocation is used to configure action for invocation
type ActionIpReleaseCampaignAttemptShowInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignAttemptShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignAttemptShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseCampaignAttemptShowInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseCampaignAttemptShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseCampaignAttemptShowInvocation) SetPathParamString(param string, value string) *ActionIpReleaseCampaignAttemptShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignAttemptShowInvocation) NewMetaInput() *ActionIpReleaseCampaignAttemptShowMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignAttemptShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignAttemptShowInvocation) SetMetaInput(input *ActionIpReleaseCampaignAttemptShowMetaGlobalInput) *ActionIpReleaseCampaignAttemptShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignAttemptShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignAttemptShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignAttemptShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionIpReleaseCampaignAttemptShowInvocation) Call() (*ActionIpReleaseCampaignAttemptShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseCampaignAttemptShowInvocation) CallContext(ctx context.Context) (*ActionIpReleaseCampaignAttemptShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseCampaignAttemptShowResponse{Action: inv.Action}
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
		if resp.Response.Attempt == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Attempt
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpReleaseCampaignAttemptShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseCampaignAttemptShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionIpReleaseCampaignAttemptShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
