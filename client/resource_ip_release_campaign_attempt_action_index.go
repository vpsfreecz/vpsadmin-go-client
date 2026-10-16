package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpReleaseCampaignAttemptIndex is a type for action Ip_release_campaign.Attempt#Index
type ActionIpReleaseCampaignAttemptIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignAttemptIndex(client *Client) *ActionIpReleaseCampaignAttemptIndex {
	return &ActionIpReleaseCampaignAttemptIndex{
		Client: client,
	}
}

// ActionIpReleaseCampaignAttemptIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignAttemptIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput) SetCount(value bool) *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignAttemptIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignAttemptIndexInput is a type for action input parameters
type ActionIpReleaseCampaignAttemptIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpReleaseCampaignAttemptIndexInput) SetFromId(value int64) *ActionIpReleaseCampaignAttemptIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpReleaseCampaignAttemptIndexInput) SetLimit(value int64) *ActionIpReleaseCampaignAttemptIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignAttemptIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignAttemptIndexInput) SelectParameters(params ...string) *ActionIpReleaseCampaignAttemptIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseCampaignAttemptIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignAttemptIndexInput) UnselectParameters(params ...string) *ActionIpReleaseCampaignAttemptIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseCampaignAttemptIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignAttemptIndexOutput is a type for action output parameters
type ActionIpReleaseCampaignAttemptIndexOutput struct {
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
type ActionIpReleaseCampaignAttemptIndexResponse struct {
	Action *ActionIpReleaseCampaignAttemptIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Attempts []*ActionIpReleaseCampaignAttemptIndexOutput "json:\"attempts\""
	}

	// Action output without the namespace
	Output []*ActionIpReleaseCampaignAttemptIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignAttemptIndex) Prepare() *ActionIpReleaseCampaignAttemptIndexInvocation {
	return &ActionIpReleaseCampaignAttemptIndexInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/{ip_release_campaign_id}/attempts",
	}
}

// ActionIpReleaseCampaignAttemptIndexInvocation is used to configure action for invocation
type ActionIpReleaseCampaignAttemptIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignAttemptIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseCampaignAttemptIndexInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseCampaignAttemptIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) SetPathParamString(param string, value string) *ActionIpReleaseCampaignAttemptIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) NewInput() *ActionIpReleaseCampaignAttemptIndexInput {
	inv.Input = &ActionIpReleaseCampaignAttemptIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) SetInput(input *ActionIpReleaseCampaignAttemptIndexInput) *ActionIpReleaseCampaignAttemptIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) NewMetaInput() *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignAttemptIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) SetMetaInput(input *ActionIpReleaseCampaignAttemptIndexMetaGlobalInput) *ActionIpReleaseCampaignAttemptIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) validate() error {
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
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) Call() (*ActionIpReleaseCampaignAttemptIndexResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) CallContext(ctx context.Context) (*ActionIpReleaseCampaignAttemptIndexResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseCampaignAttemptIndexResponse{Action: inv.Action}
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
		if resp.Response.Attempts == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Attempts
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["attempt[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["attempt[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionIpReleaseCampaignAttemptIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
