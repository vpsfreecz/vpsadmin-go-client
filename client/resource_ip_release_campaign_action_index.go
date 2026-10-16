package client

import (
	"context"
	"net/http"
)

// ActionIpReleaseCampaignIndex is a type for action Ip_release_campaign#Index
type ActionIpReleaseCampaignIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignIndex(client *Client) *ActionIpReleaseCampaignIndex {
	return &ActionIpReleaseCampaignIndex{
		Client: client,
	}
}

// ActionIpReleaseCampaignIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpReleaseCampaignIndexMetaGlobalInput) SetCount(value bool) *ActionIpReleaseCampaignIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignIndexMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignIndexMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignIndexInput is a type for action input parameters
type ActionIpReleaseCampaignIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	Open   bool  "json:\"open\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpReleaseCampaignIndexInput) SetFromId(value int64) *ActionIpReleaseCampaignIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpReleaseCampaignIndexInput) SetLimit(value int64) *ActionIpReleaseCampaignIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOpen sets parameter Open to value and selects it for sending
func (in *ActionIpReleaseCampaignIndexInput) SetOpen(value bool) *ActionIpReleaseCampaignIndexInput {
	in.Open = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Open"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignIndexInput) SelectParameters(params ...string) *ActionIpReleaseCampaignIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseCampaignIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignIndexInput) UnselectParameters(params ...string) *ActionIpReleaseCampaignIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseCampaignIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignIndexOutput is a type for action output parameters
type ActionIpReleaseCampaignIndexOutput struct {
	AllowKeep        bool                  "json:\"allow_keep\""
	ClosedAt         string                "json:\"closed_at\""
	ClosedBy         *ActionUserShowOutput "json:\"closed_by\""
	CreatedAt        string                "json:\"created_at\""
	CreatedBy        *ActionUserShowOutput "json:\"created_by\""
	Deadline         string                "json:\"deadline\""
	Id               int64                 "json:\"id\""
	KeptIpCount      int64                 "json:\"kept_ip_count\""
	ToReleaseIpCount int64                 "json:\"to_release_ip_count\""
	TotalIpCount     int64                 "json:\"total_ip_count\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignIndexResponse struct {
	Action *ActionIpReleaseCampaignIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseCampaigns []*ActionIpReleaseCampaignIndexOutput "json:\"ip_release_campaigns\""
	}

	// Action output without the namespace
	Output []*ActionIpReleaseCampaignIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignIndex) Prepare() *ActionIpReleaseCampaignIndexInvocation {
	return &ActionIpReleaseCampaignIndexInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns",
	}
}

// ActionIpReleaseCampaignIndexInvocation is used to configure action for invocation
type ActionIpReleaseCampaignIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseCampaignIndexInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseCampaignIndexInvocation) NewInput() *ActionIpReleaseCampaignIndexInput {
	inv.Input = &ActionIpReleaseCampaignIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseCampaignIndexInvocation) SetInput(input *ActionIpReleaseCampaignIndexInput) *ActionIpReleaseCampaignIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseCampaignIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignIndexInvocation) NewMetaInput() *ActionIpReleaseCampaignIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignIndexInvocation) SetMetaInput(input *ActionIpReleaseCampaignIndexMetaGlobalInput) *ActionIpReleaseCampaignIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignIndexInvocation) validate() error {
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
func (inv *ActionIpReleaseCampaignIndexInvocation) Call() (*ActionIpReleaseCampaignIndexResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseCampaignIndexInvocation) CallContext(ctx context.Context) (*ActionIpReleaseCampaignIndexResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseCampaignIndexResponse{Action: inv.Action}
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
		if resp.Response.IpReleaseCampaigns == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.IpReleaseCampaigns
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpReleaseCampaignIndexInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseCampaignIndexInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionIpReleaseCampaignIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["ip_release_campaign[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["ip_release_campaign[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Open") {
			ret["ip_release_campaign[open]"] = convertBoolToString(inv.Input.Open)
		}
	}
}

func (inv *ActionIpReleaseCampaignIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
