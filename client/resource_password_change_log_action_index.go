package client

import (
	"context"
	"net/http"
)

// ActionPasswordChangeLogIndex is a type for action Password_change_log#Index
type ActionPasswordChangeLogIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionPasswordChangeLogIndex(client *Client) *ActionPasswordChangeLogIndex {
	return &ActionPasswordChangeLogIndex{
		Client: client,
	}
}

// ActionPasswordChangeLogIndexMetaGlobalInput is a type for action global meta input parameters
type ActionPasswordChangeLogIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionPasswordChangeLogIndexMetaGlobalInput) SetCount(value bool) *ActionPasswordChangeLogIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionPasswordChangeLogIndexMetaGlobalInput) SetIncludes(value string) *ActionPasswordChangeLogIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionPasswordChangeLogIndexMetaGlobalInput) SetNo(value bool) *ActionPasswordChangeLogIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionPasswordChangeLogIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPasswordChangeLogIndexMetaGlobalInput) SelectParameters(params ...string) *ActionPasswordChangeLogIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPasswordChangeLogIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPasswordChangeLogIndexInput is a type for action input parameters
type ActionPasswordChangeLogIndexInput struct {
	FromId        int64  "json:\"from_id\""
	Limit         int64  "json:\"limit\""
	Source        string "json:\"source\""
	User          int64  "json:\"user\""
	UserSessionId int64  "json:\"user_session_id\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionPasswordChangeLogIndexInput) SetFromId(value int64) *ActionPasswordChangeLogIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionPasswordChangeLogIndexInput) SetLimit(value int64) *ActionPasswordChangeLogIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetSource sets parameter Source to value and selects it for sending
func (in *ActionPasswordChangeLogIndexInput) SetSource(value string) *ActionPasswordChangeLogIndexInput {
	in.Source = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Source"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionPasswordChangeLogIndexInput) SetUser(value int64) *ActionPasswordChangeLogIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SetUserSessionId sets parameter UserSessionId to value and selects it for sending
func (in *ActionPasswordChangeLogIndexInput) SetUserSessionId(value int64) *ActionPasswordChangeLogIndexInput {
	in.UserSessionId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserSessionIdNil(false)
	in._selectedParameters["UserSessionId"] = nil
	return in
}

// SetUserSessionIdNil sets parameter UserSessionId to nil and selects it for sending
func (in *ActionPasswordChangeLogIndexInput) SetUserSessionIdNil(set bool) *ActionPasswordChangeLogIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["UserSessionId"] = nil
		in.SelectParameters("UserSessionId")
	} else {
		delete(in._nilParameters, "UserSessionId")
	}
	return in
}

// SelectParameters sets parameters from ActionPasswordChangeLogIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPasswordChangeLogIndexInput) SelectParameters(params ...string) *ActionPasswordChangeLogIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionPasswordChangeLogIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionPasswordChangeLogIndexInput) UnselectParameters(params ...string) *ActionPasswordChangeLogIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionPasswordChangeLogIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPasswordChangeLogIndexOutput is a type for action output parameters
type ActionPasswordChangeLogIndexOutput struct {
	ClientIpAddr           string                       "json:\"client_ip_addr\""
	ClientIpPtr            string                       "json:\"client_ip_ptr\""
	CreatedAt              string                       "json:\"created_at\""
	Id                     int64                        "json:\"id\""
	Source                 string                       "json:\"source\""
	User                   *ActionUserShowOutput        "json:\"user\""
	UserAgent              string                       "json:\"user_agent\""
	UserSession            *ActionUserSessionShowOutput "json:\"user_session\""
	UserSessionId          int64                        "json:\"user_session_id\""
	UserSessionOwnedByUser bool                         "json:\"user_session_owned_by_user\""
}

// Type for action response, including envelope
type ActionPasswordChangeLogIndexResponse struct {
	Action *ActionPasswordChangeLogIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PasswordChangeLogs []*ActionPasswordChangeLogIndexOutput "json:\"password_change_logs\""
	}

	// Action output without the namespace
	Output []*ActionPasswordChangeLogIndexOutput
}

// Prepare the action for invocation
func (action *ActionPasswordChangeLogIndex) Prepare() *ActionPasswordChangeLogIndexInvocation {
	return &ActionPasswordChangeLogIndexInvocation{
		Action: action,
		Path:   "/v7.0/password_change_logs",
	}
}

// ActionPasswordChangeLogIndexInvocation is used to configure action for invocation
type ActionPasswordChangeLogIndexInvocation struct {
	// Pointer to the action
	Action *ActionPasswordChangeLogIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionPasswordChangeLogIndexInput
	// Global meta input parameters
	MetaInput *ActionPasswordChangeLogIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionPasswordChangeLogIndexInvocation) NewInput() *ActionPasswordChangeLogIndexInput {
	inv.Input = &ActionPasswordChangeLogIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionPasswordChangeLogIndexInvocation) SetInput(input *ActionPasswordChangeLogIndexInput) *ActionPasswordChangeLogIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionPasswordChangeLogIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionPasswordChangeLogIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionPasswordChangeLogIndexInvocation) NewMetaInput() *ActionPasswordChangeLogIndexMetaGlobalInput {
	inv.MetaInput = &ActionPasswordChangeLogIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionPasswordChangeLogIndexInvocation) SetMetaInput(input *ActionPasswordChangeLogIndexMetaGlobalInput) *ActionPasswordChangeLogIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionPasswordChangeLogIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionPasswordChangeLogIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionPasswordChangeLogIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", t(inv.Action.Client, "validation.invalid_resource_id", nil))
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
func (inv *ActionPasswordChangeLogIndexInvocation) Call() (*ActionPasswordChangeLogIndexResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionPasswordChangeLogIndexInvocation) CallContext(ctx context.Context) (*ActionPasswordChangeLogIndexResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionPasswordChangeLogIndexResponse{Action: inv.Action}
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
		if resp.Response.PasswordChangeLogs == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.PasswordChangeLogs
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionPasswordChangeLogIndexInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionPasswordChangeLogIndexInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionPasswordChangeLogIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["password_change_log[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["password_change_log[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Source") {
			ret["password_change_log[source]"] = inv.Input.Source
		}
		if inv.IsParameterSelected("User") {
			ret["password_change_log[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("UserSessionId") {
			if inv.IsParameterNil("UserSessionId") {
				ret["password_change_log[user_session_id]"] = ""
			} else {
				ret["password_change_log[user_session_id]"] = convertInt64ToString(inv.Input.UserSessionId)
			}
		}
	}
}

func (inv *ActionPasswordChangeLogIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
