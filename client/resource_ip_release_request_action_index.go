package client

import (
	"context"
	"net/http"
)

// ActionIpReleaseRequestIndex is a type for action Ip_release_request#Index
type ActionIpReleaseRequestIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseRequestIndex(client *Client) *ActionIpReleaseRequestIndex {
	return &ActionIpReleaseRequestIndex{
		Client: client,
	}
}

// ActionIpReleaseRequestIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseRequestIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpReleaseRequestIndexMetaGlobalInput) SetCount(value bool) *ActionIpReleaseRequestIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseRequestIndexMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseRequestIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseRequestIndexMetaGlobalInput) SetNo(value bool) *ActionIpReleaseRequestIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseRequestIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseRequestIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestIndexInput is a type for action input parameters
type ActionIpReleaseRequestIndexInput struct {
	FromId            int64 "json:\"from_id\""
	IpReleaseCampaign int64 "json:\"ip_release_campaign\""
	Limit             int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpReleaseRequestIndexInput) SetFromId(value int64) *ActionIpReleaseRequestIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetIpReleaseCampaign sets parameter IpReleaseCampaign to value and selects it for sending
func (in *ActionIpReleaseRequestIndexInput) SetIpReleaseCampaign(value int64) *ActionIpReleaseRequestIndexInput {
	in.IpReleaseCampaign = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpReleaseCampaign"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpReleaseRequestIndexInput) SetLimit(value int64) *ActionIpReleaseRequestIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestIndexInput) SelectParameters(params ...string) *ActionIpReleaseRequestIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseRequestIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestIndexInput) UnselectParameters(params ...string) *ActionIpReleaseRequestIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseRequestIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestIndexOutput is a type for action output parameters
type ActionIpReleaseRequestIndexOutput struct {
	AllowKeep         bool                               "json:\"allow_keep\""
	CanAssign         bool                               "json:\"can_assign\""
	CanKeep           bool                               "json:\"can_keep\""
	ClosedAt          string                             "json:\"closed_at\""
	Deadline          string                             "json:\"deadline\""
	Id                int64                              "json:\"id\""
	IpReleaseCampaign *ActionIpReleaseCampaignShowOutput "json:\"ip_release_campaign\""
	MailLog           *ActionMailLogShowOutput           "json:\"mail_log\""
	NotifiedAt        string                             "json:\"notified_at\""
	OriginalUserId    int64                              "json:\"original_user_id\""
	User              *ActionUserShowOutput              "json:\"user\""
	UserLogin         string                             "json:\"user_login\""
}

// Type for action response, including envelope
type ActionIpReleaseRequestIndexResponse struct {
	Action *ActionIpReleaseRequestIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseRequests []*ActionIpReleaseRequestIndexOutput "json:\"ip_release_requests\""
	}

	// Action output without the namespace
	Output []*ActionIpReleaseRequestIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseRequestIndex) Prepare() *ActionIpReleaseRequestIndexInvocation {
	return &ActionIpReleaseRequestIndexInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_requests",
	}
}

// ActionIpReleaseRequestIndexInvocation is used to configure action for invocation
type ActionIpReleaseRequestIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseRequestIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseRequestIndexInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseRequestIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseRequestIndexInvocation) NewInput() *ActionIpReleaseRequestIndexInput {
	inv.Input = &ActionIpReleaseRequestIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseRequestIndexInvocation) SetInput(input *ActionIpReleaseRequestIndexInput) *ActionIpReleaseRequestIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseRequestIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseRequestIndexInvocation) NewMetaInput() *ActionIpReleaseRequestIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseRequestIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseRequestIndexInvocation) SetMetaInput(input *ActionIpReleaseRequestIndexMetaGlobalInput) *ActionIpReleaseRequestIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseRequestIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseRequestIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("IpReleaseCampaign") {
			if !inv.IsParameterNil("IpReleaseCampaign") {
				if inv.Input.IpReleaseCampaign < 0 {
					verr.Add("ip_release_campaign", t(inv.Action.Client, "validation.invalid_resource_id", nil))
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionIpReleaseRequestIndexInvocation) Call() (*ActionIpReleaseRequestIndexResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseRequestIndexInvocation) CallContext(ctx context.Context) (*ActionIpReleaseRequestIndexResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseRequestIndexResponse{Action: inv.Action}
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
		if resp.Response.IpReleaseRequests == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.IpReleaseRequests
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpReleaseRequestIndexInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseRequestIndexInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionIpReleaseRequestIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["ip_release_request[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("IpReleaseCampaign") {
			ret["ip_release_request[ip_release_campaign]"] = convertInt64ToString(inv.Input.IpReleaseCampaign)
		}
		if inv.IsParameterSelected("Limit") {
			ret["ip_release_request[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionIpReleaseRequestIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
