package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpReleaseRequestAddressIndex is a type for action Ip_release_request.Address#Index
type ActionIpReleaseRequestAddressIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseRequestAddressIndex(client *Client) *ActionIpReleaseRequestAddressIndex {
	return &ActionIpReleaseRequestAddressIndex{
		Client: client,
	}
}

// ActionIpReleaseRequestAddressIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseRequestAddressIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpReleaseRequestAddressIndexMetaGlobalInput) SetCount(value bool) *ActionIpReleaseRequestAddressIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseRequestAddressIndexMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseRequestAddressIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseRequestAddressIndexMetaGlobalInput) SetNo(value bool) *ActionIpReleaseRequestAddressIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestAddressIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestAddressIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseRequestAddressIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseRequestAddressIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestAddressIndexInput is a type for action input parameters
type ActionIpReleaseRequestAddressIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpReleaseRequestAddressIndexInput) SetFromId(value int64) *ActionIpReleaseRequestAddressIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpReleaseRequestAddressIndexInput) SetLimit(value int64) *ActionIpReleaseRequestAddressIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestAddressIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestAddressIndexInput) SelectParameters(params ...string) *ActionIpReleaseRequestAddressIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseRequestAddressIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestAddressIndexInput) UnselectParameters(params ...string) *ActionIpReleaseRequestAddressIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseRequestAddressIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestAddressIndexOutput is a type for action output parameters
type ActionIpReleaseRequestAddressIndexOutput struct {
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
type ActionIpReleaseRequestAddressIndexResponse struct {
	Action *ActionIpReleaseRequestAddressIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Addresses []*ActionIpReleaseRequestAddressIndexOutput "json:\"addresses\""
	}

	// Action output without the namespace
	Output []*ActionIpReleaseRequestAddressIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseRequestAddressIndex) Prepare() *ActionIpReleaseRequestAddressIndexInvocation {
	return &ActionIpReleaseRequestAddressIndexInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_requests/{ip_release_request_id}/addresses",
	}
}

// ActionIpReleaseRequestAddressIndexInvocation is used to configure action for invocation
type ActionIpReleaseRequestAddressIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseRequestAddressIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseRequestAddressIndexInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseRequestAddressIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseRequestAddressIndexInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseRequestAddressIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseRequestAddressIndexInvocation) SetPathParamString(param string, value string) *ActionIpReleaseRequestAddressIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseRequestAddressIndexInvocation) NewInput() *ActionIpReleaseRequestAddressIndexInput {
	inv.Input = &ActionIpReleaseRequestAddressIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseRequestAddressIndexInvocation) SetInput(input *ActionIpReleaseRequestAddressIndexInput) *ActionIpReleaseRequestAddressIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseRequestAddressIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestAddressIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseRequestAddressIndexInvocation) NewMetaInput() *ActionIpReleaseRequestAddressIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseRequestAddressIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseRequestAddressIndexInvocation) SetMetaInput(input *ActionIpReleaseRequestAddressIndexMetaGlobalInput) *ActionIpReleaseRequestAddressIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseRequestAddressIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestAddressIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseRequestAddressIndexInvocation) validate() error {
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
func (inv *ActionIpReleaseRequestAddressIndexInvocation) Call() (*ActionIpReleaseRequestAddressIndexResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseRequestAddressIndexInvocation) CallContext(ctx context.Context) (*ActionIpReleaseRequestAddressIndexResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseRequestAddressIndexResponse{Action: inv.Action}
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
		if resp.Response.Addresses == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Addresses
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpReleaseRequestAddressIndexInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseRequestAddressIndexInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionIpReleaseRequestAddressIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["address[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["address[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionIpReleaseRequestAddressIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
