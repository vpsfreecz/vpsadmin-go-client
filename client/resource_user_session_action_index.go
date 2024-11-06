package client

import ()

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
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserSessionIndexMetaGlobalInput) SetNo(value bool) *ActionUserSessionIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	Admin         int64  `json:"admin"`
	ApiIpAddr     string `json:"api_ip_addr"`
	AuthType      string `json:"auth_type"`
	ClientIpAddr  string `json:"client_ip_addr"`
	ClientVersion string `json:"client_version"`
	FromId        int64  `json:"from_id"`
	IpAddr        string `json:"ip_addr"`
	Limit         int64  `json:"limit"`
	State         string `json:"state"`
	TokenFragment string `json:"token_fragment"`
	User          int64  `json:"user"`
	UserAgent     string `json:"user_agent"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAdmin sets parameter Admin to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetAdmin(value int64) *ActionUserSessionIndexInput {
	in.Admin = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetAdminNil(false)
	in._selectedParameters["Admin"] = nil
	return in
}

// SetAdminNil sets parameter Admin to nil and selects it for sending
func (in *ActionUserSessionIndexInput) SetAdminNil(set bool) *ActionUserSessionIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Admin"] = nil
		in.SelectParameters("Admin")
	} else {
		delete(in._nilParameters, "Admin")
	}
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

// SetAuthType sets parameter AuthType to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetAuthType(value string) *ActionUserSessionIndexInput {
	in.AuthType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AuthType"] = nil
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

// SetClientVersion sets parameter ClientVersion to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetClientVersion(value string) *ActionUserSessionIndexInput {
	in.ClientVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientVersion"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetFromId(value int64) *ActionUserSessionIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
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

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetLimit(value int64) *ActionUserSessionIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetState(value string) *ActionUserSessionIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetTokenFragment sets parameter TokenFragment to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetTokenFragment(value string) *ActionUserSessionIndexInput {
	in.TokenFragment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TokenFragment"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserSessionIndexInput) SetUser(value int64) *ActionUserSessionIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionUserSessionIndexInput) SetUserNil(set bool) *ActionUserSessionIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
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

// UnselectParameters unsets parameters from ActionUserSessionIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserSessionIndexInput) UnselectParameters(params ...string) *ActionUserSessionIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Admin         *ActionUserShowOutput `json:"admin"`
	ApiIpAddr     string                `json:"api_ip_addr"`
	ApiIpPtr      string                `json:"api_ip_ptr"`
	AuthType      string                `json:"auth_type"`
	ClientIpAddr  string                `json:"client_ip_addr"`
	ClientIpPtr   string                `json:"client_ip_ptr"`
	ClientVersion string                `json:"client_version"`
	ClosedAt      string                `json:"closed_at"`
	CreatedAt     string                `json:"created_at"`
	Id            int64                 `json:"id"`
	Label         string                `json:"label"`
	LastRequestAt string                `json:"last_request_at"`
	RequestCount  int64                 `json:"request_count"`
	Scope         string                `json:"scope"`
	TokenFragment string                `json:"token_fragment"`
	TokenInterval int64                 `json:"token_interval"`
	TokenLifetime string                `json:"token_lifetime"`
	User          *ActionUserShowOutput `json:"user"`
	UserAgent     string                `json:"user_agent"`
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
		Path:   "/v7.0/user_sessions",
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

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserSessionIndexInvocation) NewInput() *ActionUserSessionIndexInput {
	inv.Input = &ActionUserSessionIndexInput{}
	return inv.Input
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserSessionIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserSessionIndexInvocation) NewMetaInput() *ActionUserSessionIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserSessionIndexMetaGlobalInput{}
	return inv.MetaInput
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserSessionIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		if inv.IsParameterSelected("Admin") {
			ret["user_session[admin]"] = convertInt64ToString(inv.Input.Admin)
		}
		if inv.IsParameterSelected("ApiIpAddr") {
			ret["user_session[api_ip_addr]"] = inv.Input.ApiIpAddr
		}
		if inv.IsParameterSelected("AuthType") {
			ret["user_session[auth_type]"] = inv.Input.AuthType
		}
		if inv.IsParameterSelected("ClientIpAddr") {
			ret["user_session[client_ip_addr]"] = inv.Input.ClientIpAddr
		}
		if inv.IsParameterSelected("ClientVersion") {
			ret["user_session[client_version]"] = inv.Input.ClientVersion
		}
		if inv.IsParameterSelected("FromId") {
			ret["user_session[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("IpAddr") {
			ret["user_session[ip_addr]"] = inv.Input.IpAddr
		}
		if inv.IsParameterSelected("Limit") {
			ret["user_session[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("State") {
			ret["user_session[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("TokenFragment") {
			ret["user_session[token_fragment]"] = inv.Input.TokenFragment
		}
		if inv.IsParameterSelected("User") {
			ret["user_session[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("UserAgent") {
			ret["user_session[user_agent]"] = inv.Input.UserAgent
		}
	}
}

func (inv *ActionUserSessionIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
