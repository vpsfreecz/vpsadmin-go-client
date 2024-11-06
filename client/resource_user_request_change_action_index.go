package client

import ()

// ActionUserRequestChangeIndex is a type for action User_request.Change#Index
type ActionUserRequestChangeIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestChangeIndex(client *Client) *ActionUserRequestChangeIndex {
	return &ActionUserRequestChangeIndex{
		Client: client,
	}
}

// ActionUserRequestChangeIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestChangeIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserRequestChangeIndexMetaGlobalInput) SetCount(value bool) *ActionUserRequestChangeIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestChangeIndexMetaGlobalInput) SetIncludes(value string) *ActionUserRequestChangeIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestChangeIndexMetaGlobalInput) SetNo(value bool) *ActionUserRequestChangeIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestChangeIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestChangeIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestChangeIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestChangeIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestChangeIndexInput is a type for action input parameters
type ActionUserRequestChangeIndexInput struct {
	Admin        int64  `json:"admin"`
	ApiIpAddr    string `json:"api_ip_addr"`
	ClientIpAddr string `json:"client_ip_addr"`
	ClientIpPtr  string `json:"client_ip_ptr"`
	FromId       int64  `json:"from_id"`
	Limit        int64  `json:"limit"`
	State        string `json:"state"`
	User         int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAdmin sets parameter Admin to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetAdmin(value int64) *ActionUserRequestChangeIndexInput {
	in.Admin = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetAdminNil(false)
	in._selectedParameters["Admin"] = nil
	return in
}

// SetAdminNil sets parameter Admin to nil and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetAdminNil(set bool) *ActionUserRequestChangeIndexInput {
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
func (in *ActionUserRequestChangeIndexInput) SetApiIpAddr(value string) *ActionUserRequestChangeIndexInput {
	in.ApiIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ApiIpAddr"] = nil
	return in
}

// SetClientIpAddr sets parameter ClientIpAddr to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetClientIpAddr(value string) *ActionUserRequestChangeIndexInput {
	in.ClientIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientIpAddr"] = nil
	return in
}

// SetClientIpPtr sets parameter ClientIpPtr to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetClientIpPtr(value string) *ActionUserRequestChangeIndexInput {
	in.ClientIpPtr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientIpPtr"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetFromId(value int64) *ActionUserRequestChangeIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetLimit(value int64) *ActionUserRequestChangeIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetState(value string) *ActionUserRequestChangeIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetUser(value int64) *ActionUserRequestChangeIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetUserNil(set bool) *ActionUserRequestChangeIndexInput {
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

// SelectParameters sets parameters from ActionUserRequestChangeIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestChangeIndexInput) SelectParameters(params ...string) *ActionUserRequestChangeIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserRequestChangeIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserRequestChangeIndexInput) UnselectParameters(params ...string) *ActionUserRequestChangeIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserRequestChangeIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestChangeIndexOutput is a type for action output parameters
type ActionUserRequestChangeIndexOutput struct {
	Address       string                `json:"address"`
	Admin         *ActionUserShowOutput `json:"admin"`
	AdminResponse string                `json:"admin_response"`
	ApiIpAddr     string                `json:"api_ip_addr"`
	ApiIpPtr      string                `json:"api_ip_ptr"`
	ChangeReason  string                `json:"change_reason"`
	ClientIpAddr  string                `json:"client_ip_addr"`
	ClientIpPtr   string                `json:"client_ip_ptr"`
	CreatedAt     string                `json:"created_at"`
	Email         string                `json:"email"`
	FullName      string                `json:"full_name"`
	Id            int64                 `json:"id"`
	Label         string                `json:"label"`
	State         string                `json:"state"`
	UpdatedAt     string                `json:"updated_at"`
	User          *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionUserRequestChangeIndexResponse struct {
	Action *ActionUserRequestChangeIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Changes []*ActionUserRequestChangeIndexOutput `json:"changes"`
	}

	// Action output without the namespace
	Output []*ActionUserRequestChangeIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserRequestChangeIndex) Prepare() *ActionUserRequestChangeIndexInvocation {
	return &ActionUserRequestChangeIndexInvocation{
		Action: action,
		Path:   "/v7.0/user_request/changes",
	}
}

// ActionUserRequestChangeIndexInvocation is used to configure action for invocation
type ActionUserRequestChangeIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestChangeIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserRequestChangeIndexInput
	// Global meta input parameters
	MetaInput *ActionUserRequestChangeIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserRequestChangeIndexInvocation) NewInput() *ActionUserRequestChangeIndexInput {
	inv.Input = &ActionUserRequestChangeIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserRequestChangeIndexInvocation) SetInput(input *ActionUserRequestChangeIndexInput) *ActionUserRequestChangeIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserRequestChangeIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserRequestChangeIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestChangeIndexInvocation) NewMetaInput() *ActionUserRequestChangeIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestChangeIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestChangeIndexInvocation) SetMetaInput(input *ActionUserRequestChangeIndexMetaGlobalInput) *ActionUserRequestChangeIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestChangeIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserRequestChangeIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestChangeIndexInvocation) Call() (*ActionUserRequestChangeIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserRequestChangeIndexInvocation) callAsQuery() (*ActionUserRequestChangeIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserRequestChangeIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Changes
	}
	return resp, err
}

func (inv *ActionUserRequestChangeIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Admin") {
			ret["change[admin]"] = convertInt64ToString(inv.Input.Admin)
		}
		if inv.IsParameterSelected("ApiIpAddr") {
			ret["change[api_ip_addr]"] = inv.Input.ApiIpAddr
		}
		if inv.IsParameterSelected("ClientIpAddr") {
			ret["change[client_ip_addr]"] = inv.Input.ClientIpAddr
		}
		if inv.IsParameterSelected("ClientIpPtr") {
			ret["change[client_ip_ptr]"] = inv.Input.ClientIpPtr
		}
		if inv.IsParameterSelected("FromId") {
			ret["change[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["change[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("State") {
			ret["change[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("User") {
			ret["change[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionUserRequestChangeIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
