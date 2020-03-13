package client

import (
)

// ActionUserRequestRegistrationIndex is a type for action User_request.Registration#Index
type ActionUserRequestRegistrationIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestRegistrationIndex(client *Client) *ActionUserRequestRegistrationIndex {
	return &ActionUserRequestRegistrationIndex{
		Client: client,
	}
}

// ActionUserRequestRegistrationIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestRegistrationIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) SetNo(value bool) *ActionUserRequestRegistrationIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) SetCount(value bool) *ActionUserRequestRegistrationIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) SetIncludes(value string) *ActionUserRequestRegistrationIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestRegistrationIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationIndexInput is a type for action input parameters
type ActionUserRequestRegistrationIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	User int64 `json:"user"`
	State string `json:"state"`
	ApiIpAddr string `json:"api_ip_addr"`
	ClientIpAddr string `json:"client_ip_addr"`
	ClientIpPtr string `json:"client_ip_ptr"`
	Admin int64 `json:"admin"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetOffset(value int64) *ActionUserRequestRegistrationIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetLimit(value int64) *ActionUserRequestRegistrationIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetUser(value int64) *ActionUserRequestRegistrationIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetState sets parameter State to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetState(value string) *ActionUserRequestRegistrationIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}
// SetApiIpAddr sets parameter ApiIpAddr to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetApiIpAddr(value string) *ActionUserRequestRegistrationIndexInput {
	in.ApiIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ApiIpAddr"] = nil
	return in
}
// SetClientIpAddr sets parameter ClientIpAddr to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetClientIpAddr(value string) *ActionUserRequestRegistrationIndexInput {
	in.ClientIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientIpAddr"] = nil
	return in
}
// SetClientIpPtr sets parameter ClientIpPtr to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetClientIpPtr(value string) *ActionUserRequestRegistrationIndexInput {
	in.ClientIpPtr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientIpPtr"] = nil
	return in
}
// SetAdmin sets parameter Admin to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetAdmin(value int64) *ActionUserRequestRegistrationIndexInput {
	in.Admin = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Admin"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationIndexInput) SelectParameters(params ...string) *ActionUserRequestRegistrationIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserRequestRegistrationIndexOutput is a type for action output parameters
type ActionUserRequestRegistrationIndexOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	State string `json:"state"`
	ApiIpAddr string `json:"api_ip_addr"`
	ApiIpPtr string `json:"api_ip_ptr"`
	ClientIpAddr string `json:"client_ip_addr"`
	ClientIpPtr string `json:"client_ip_ptr"`
	Admin *ActionUserShowOutput `json:"admin"`
	AdminResponse string `json:"admin_response"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Label string `json:"label"`
	Login string `json:"login"`
	FullName string `json:"full_name"`
	OrgName string `json:"org_name"`
	OrgId string `json:"org_id"`
	Email string `json:"email"`
	Address string `json:"address"`
	YearOfBirth int64 `json:"year_of_birth"`
	How string `json:"how"`
	Note string `json:"note"`
	OsTemplate *ActionOsTemplateShowOutput `json:"os_template"`
	Location *ActionLocationShowOutput `json:"location"`
	Currency string `json:"currency"`
	Language *ActionLanguageShowOutput `json:"language"`
}


// Type for action response, including envelope
type ActionUserRequestRegistrationIndexResponse struct {
	Action *ActionUserRequestRegistrationIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Registrations []*ActionUserRequestRegistrationIndexOutput `json:"registrations"`
	}

	// Action output without the namespace
	Output []*ActionUserRequestRegistrationIndexOutput
}


// Prepare the action for invocation
func (action *ActionUserRequestRegistrationIndex) Prepare() *ActionUserRequestRegistrationIndexInvocation {
	return &ActionUserRequestRegistrationIndexInvocation{
		Action: action,
		Path: "/v6.0/user_request/registrations",
	}
}

// ActionUserRequestRegistrationIndexInvocation is used to configure action for invocation
type ActionUserRequestRegistrationIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestRegistrationIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserRequestRegistrationIndexInput
	// Global meta input parameters
	MetaInput *ActionUserRequestRegistrationIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserRequestRegistrationIndexInvocation) NewInput() *ActionUserRequestRegistrationIndexInput {
	inv.Input = &ActionUserRequestRegistrationIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserRequestRegistrationIndexInvocation) SetInput(input *ActionUserRequestRegistrationIndexInput) *ActionUserRequestRegistrationIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserRequestRegistrationIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestRegistrationIndexInvocation) NewMetaInput() *ActionUserRequestRegistrationIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestRegistrationIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestRegistrationIndexInvocation) SetMetaInput(input *ActionUserRequestRegistrationIndexMetaGlobalInput) *ActionUserRequestRegistrationIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestRegistrationIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestRegistrationIndexInvocation) Call() (*ActionUserRequestRegistrationIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserRequestRegistrationIndexInvocation) callAsQuery() (*ActionUserRequestRegistrationIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserRequestRegistrationIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Registrations
	}
	return resp, err
}



func (inv *ActionUserRequestRegistrationIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["registration[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["registration[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["registration[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("State") {
			ret["registration[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("ApiIpAddr") {
			ret["registration[api_ip_addr]"] = inv.Input.ApiIpAddr
		}
		if inv.IsParameterSelected("ClientIpAddr") {
			ret["registration[client_ip_addr]"] = inv.Input.ClientIpAddr
		}
		if inv.IsParameterSelected("ClientIpPtr") {
			ret["registration[client_ip_ptr]"] = inv.Input.ClientIpPtr
		}
		if inv.IsParameterSelected("Admin") {
			ret["registration[admin]"] = convertInt64ToString(inv.Input.Admin)
		}
	}
}

func (inv *ActionUserRequestRegistrationIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

