package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpReleaseRequestNoticeIndex is a type for action Ip_release_request.Notice#Index
type ActionIpReleaseRequestNoticeIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseRequestNoticeIndex(client *Client) *ActionIpReleaseRequestNoticeIndex {
	return &ActionIpReleaseRequestNoticeIndex{
		Client: client,
	}
}

// ActionIpReleaseRequestNoticeIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseRequestNoticeIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpReleaseRequestNoticeIndexMetaGlobalInput) SetCount(value bool) *ActionIpReleaseRequestNoticeIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseRequestNoticeIndexMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseRequestNoticeIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseRequestNoticeIndexMetaGlobalInput) SetNo(value bool) *ActionIpReleaseRequestNoticeIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestNoticeIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestNoticeIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseRequestNoticeIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseRequestNoticeIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestNoticeIndexInput is a type for action input parameters
type ActionIpReleaseRequestNoticeIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpReleaseRequestNoticeIndexInput) SetFromId(value int64) *ActionIpReleaseRequestNoticeIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpReleaseRequestNoticeIndexInput) SetLimit(value int64) *ActionIpReleaseRequestNoticeIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestNoticeIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestNoticeIndexInput) SelectParameters(params ...string) *ActionIpReleaseRequestNoticeIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseRequestNoticeIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestNoticeIndexInput) UnselectParameters(params ...string) *ActionIpReleaseRequestNoticeIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseRequestNoticeIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestNoticeIndexOutput is a type for action output parameters
type ActionIpReleaseRequestNoticeIndexOutput struct {
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
type ActionIpReleaseRequestNoticeIndexResponse struct {
	Action *ActionIpReleaseRequestNoticeIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Notices []*ActionIpReleaseRequestNoticeIndexOutput "json:\"notices\""
	}

	// Action output without the namespace
	Output []*ActionIpReleaseRequestNoticeIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseRequestNoticeIndex) Prepare() *ActionIpReleaseRequestNoticeIndexInvocation {
	return &ActionIpReleaseRequestNoticeIndexInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_requests/{ip_release_request_id}/notices",
	}
}

// ActionIpReleaseRequestNoticeIndexInvocation is used to configure action for invocation
type ActionIpReleaseRequestNoticeIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseRequestNoticeIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseRequestNoticeIndexInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseRequestNoticeIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseRequestNoticeIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) SetPathParamString(param string, value string) *ActionIpReleaseRequestNoticeIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) NewInput() *ActionIpReleaseRequestNoticeIndexInput {
	inv.Input = &ActionIpReleaseRequestNoticeIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) SetInput(input *ActionIpReleaseRequestNoticeIndexInput) *ActionIpReleaseRequestNoticeIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) NewMetaInput() *ActionIpReleaseRequestNoticeIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseRequestNoticeIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) SetMetaInput(input *ActionIpReleaseRequestNoticeIndexMetaGlobalInput) *ActionIpReleaseRequestNoticeIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseRequestNoticeIndexInvocation) validate() error {
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
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) Call() (*ActionIpReleaseRequestNoticeIndexResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) CallContext(ctx context.Context) (*ActionIpReleaseRequestNoticeIndexResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpReleaseRequestNoticeIndexResponse{Action: inv.Action}
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
func (inv *ActionIpReleaseRequestNoticeIndexInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpReleaseRequestNoticeIndexInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionIpReleaseRequestNoticeIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["notice[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["notice[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionIpReleaseRequestNoticeIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
