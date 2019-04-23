package client

import (
)

// ActionUserSessionIndex is a type for action User_session#Index
type ActionUserSessionIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserSessionIndex(client *Client) *ActionUserSessionIndex {
	return &ActionUserSessionIndex{
		Client: client,
	}
}

// ActionUserSessionIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserSessionIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserSessionIndexMetaGlobalInput) SetNo(value bool) *ActionUserSessionIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserSessionIndexMetaGlobalInput) SetCount(value bool) *ActionUserSessionIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserSessionIndexMetaGlobalInput) SetIncludes(value string) *ActionUserSessionIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserSessionIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSessionIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserSessionIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserSessionIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserSessionIndexInput is a type for action input parameters
type ActionUserSessionIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	User int64 `json:"user"`
	AuthType string `json:"auth_type"`
	ApiIpAddr string `json:"api_ip_addr"`
	ClientIpAddr string `json:"client_ip_addr"`
	UserAgent string `json:"user_agent"`
	ClientVersion string `json:"client_version"`
	AuthTokenStr string `json:"auth_token_str"`
	Admin int64 `json:"admin"`
	IpAddr string `json:"ip_addr"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetOffset(value int64) *ActionUserSessionIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetLimit(value int64) *ActionUserSessionIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetUser(value int64) *ActionUserSessionIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetAuthType sets parameter AuthType to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetAuthType(value string) *ActionUserSessionIndexInput {
	in.AuthType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AuthType"] = nil
	return in
}
// SetApiIpAddr sets parameter ApiIpAddr to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetApiIpAddr(value string) *ActionUserSessionIndexInput {
	in.ApiIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ApiIpAddr"] = nil
	return in
}
// SetClientIpAddr sets parameter ClientIpAddr to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetClientIpAddr(value string) *ActionUserSessionIndexInput {
	in.ClientIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientIpAddr"] = nil
	return in
}
// SetUserAgent sets parameter UserAgent to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetUserAgent(value string) *ActionUserSessionIndexInput {
	in.UserAgent = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserAgent"] = nil
	return in
}
// SetClientVersion sets parameter ClientVersion to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetClientVersion(value string) *ActionUserSessionIndexInput {
	in.ClientVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientVersion"] = nil
	return in
}
// SetAuthTokenStr sets parameter AuthTokenStr to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetAuthTokenStr(value string) *ActionUserSessionIndexInput {
	in.AuthTokenStr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AuthTokenStr"] = nil
	return in
}
// SetAdmin sets parameter Admin to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetAdmin(value int64) *ActionUserSessionIndexInput {
	in.Admin = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Admin"] = nil
	return in
}
// SetIpAddr sets parameter IpAddr to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetIpAddr(value string) *ActionUserSessionIndexInput {
	in.IpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpAddr"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserSessionIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSessionIndexInput) SelectParameters(params ...string) *ActionUserSessionIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserSessionIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserSessionIndexOutput is a type for action output parameters
type ActionUserSessionIndexOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	AuthType string `json:"auth_type"`
	ApiIpAddr string `json:"api_ip_addr"`
	ApiIpPtr string `json:"api_ip_ptr"`
	ClientIpAddr string `json:"client_ip_addr"`
	ClientIpPtr string `json:"client_ip_ptr"`
	UserAgent string `json:"user_agent"`
	ClientVersion string `json:"client_version"`
	AuthToken *ActionAuthTokenShowOutput `json:"auth_token"`
	AuthTokenStr string `json:"auth_token_str"`
	CreatedAt string `json:"created_at"`
	LastRequestAt string `json:"last_request_at"`
	ClosedAt string `json:"closed_at"`
	Admin *ActionUserShowOutput `json:"admin"`
}


// Type for action response, including envelope
type ActionUserSessionIndexResponse struct {
	Action *ActionUserSessionIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserSessions []*ActionUserSessionIndexOutput `json:"user_sessions"`
	}

	// Action output without the namespace
	Output []*ActionUserSessionIndexOutput
}


// Prepare the action for invocation
func (action *ActionUserSessionIndex) Prepare() *ActionUserSessionIndexInvocation {
	return &ActionUserSessionIndexInvocation{
		Action: action,
		Path: "/v5.0/user_sessions",
	}
}

// ActionUserSessionIndexInvocation is used to configure action for invocation
type ActionUserSessionIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserSessionIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserSessionIndexInput
	// Global meta input parameters
	MetaInput *ActionUserSessionIndexMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionUserSessionIndexInvocation) SetInput(input *ActionUserSessionIndexInput) *ActionUserSessionIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserSessionIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserSessionIndexInvocation) SetMetaInput(input *ActionUserSessionIndexMetaGlobalInput) *ActionUserSessionIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserSessionIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserSessionIndexInvocation) Call() (*ActionUserSessionIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserSessionIndexInvocation) callAsQuery() (*ActionUserSessionIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserSessionIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserSessions
	}
	return resp, err
}



func (inv *ActionUserSessionIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["user_session[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["user_session[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["user_session[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("AuthType") {
			ret["user_session[auth_type]"] = inv.Input.AuthType
		}
		if inv.IsParameterSelected("ApiIpAddr") {
			ret["user_session[api_ip_addr]"] = inv.Input.ApiIpAddr
		}
		if inv.IsParameterSelected("ClientIpAddr") {
			ret["user_session[client_ip_addr]"] = inv.Input.ClientIpAddr
		}
		if inv.IsParameterSelected("UserAgent") {
			ret["user_session[user_agent]"] = inv.Input.UserAgent
		}
		if inv.IsParameterSelected("ClientVersion") {
			ret["user_session[client_version]"] = inv.Input.ClientVersion
		}
		if inv.IsParameterSelected("AuthTokenStr") {
			ret["user_session[auth_token_str]"] = inv.Input.AuthTokenStr
		}
		if inv.IsParameterSelected("Admin") {
			ret["user_session[admin]"] = convertInt64ToString(inv.Input.Admin)
		}
		if inv.IsParameterSelected("IpAddr") {
			ret["user_session[ip_addr]"] = inv.Input.IpAddr
		}
	}
}

func (inv *ActionUserSessionIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

