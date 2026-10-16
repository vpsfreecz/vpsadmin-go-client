package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpReleaseCampaignNoticeIndex is a type for action Ip_release_campaign.Notice#Index
type ActionIpReleaseCampaignNoticeIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignNoticeIndex(client *Client) *ActionIpReleaseCampaignNoticeIndex {
	return &ActionIpReleaseCampaignNoticeIndex{
		Client: client,
	}
}

// ActionIpReleaseCampaignNoticeIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignNoticeIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput) SetCount(value bool) *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignNoticeIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignNoticeIndexInput is a type for action input parameters
type ActionIpReleaseCampaignNoticeIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpReleaseCampaignNoticeIndexInput) SetFromId(value int64) *ActionIpReleaseCampaignNoticeIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpReleaseCampaignNoticeIndexInput) SetLimit(value int64) *ActionIpReleaseCampaignNoticeIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignNoticeIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignNoticeIndexInput) SelectParameters(params ...string) *ActionIpReleaseCampaignNoticeIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseCampaignNoticeIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignNoticeIndexInput) UnselectParameters(params ...string) *ActionIpReleaseCampaignNoticeIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseCampaignNoticeIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignNoticeIndexOutput is a type for action output parameters
type ActionIpReleaseCampaignNoticeIndexOutput struct {
	CreatedAt      string                   "json:\"created_at\""
	CreatedBy      *ActionUserShowOutput    "json:\"created_by\""
	CreatedByLogin string                   "json:\"created_by_login\""
	Event          string                   "json:\"event\""
	Id             int64                    "json:\"id\""
	MailLog        *ActionMailLogShowOutput "json:\"mail_log\""
	OriginalUserId int64                    "json:\"original_user_id\""
	Subject        string                   "json:\"subject\""
	UserLogin      string                   "json:\"user_login\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignNoticeIndexResponse struct {
	Action *ActionIpReleaseCampaignNoticeIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Notices []*ActionIpReleaseCampaignNoticeIndexOutput "json:\"notices\""
	}

	// Action output without the namespace
	Output []*ActionIpReleaseCampaignNoticeIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignNoticeIndex) Prepare() *ActionIpReleaseCampaignNoticeIndexInvocation {
	return &ActionIpReleaseCampaignNoticeIndexInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/{ip_release_campaign_id}/notices",
	}
}

// ActionIpReleaseCampaignNoticeIndexInvocation is used to configure action for invocation
type ActionIpReleaseCampaignNoticeIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignNoticeIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseCampaignNoticeIndexInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseCampaignNoticeIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) SetPathParamString(param string, value string) *ActionIpReleaseCampaignNoticeIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) NewInput() *ActionIpReleaseCampaignNoticeIndexInput {
	inv.Input = &ActionIpReleaseCampaignNoticeIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) SetInput(input *ActionIpReleaseCampaignNoticeIndexInput) *ActionIpReleaseCampaignNoticeIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) NewMetaInput() *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignNoticeIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) SetMetaInput(input *ActionIpReleaseCampaignNoticeIndexMetaGlobalInput) *ActionIpReleaseCampaignNoticeIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) validate() error {
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
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) Call() (*ActionIpReleaseCampaignNoticeIndexResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) CallContext(ctx context.Context) (*ActionIpReleaseCampaignNoticeIndexResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseCampaignNoticeIndexResponse{Action: inv.Action}
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
		if resp.Response.Notices == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Notices
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["notice[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["notice[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionIpReleaseCampaignNoticeIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
