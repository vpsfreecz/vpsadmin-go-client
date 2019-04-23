package client

import (
)

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
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	User int64 `json:"user"`
	State string `json:"state"`
	ApiIpAddr string `json:"api_ip_addr"`
	ClientIpAddr string `json:"client_ip_addr"`
	Admin int64 `json:"admin"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetOffset(value int64) *ActionUserRequestChangeIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
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
// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetUser(value int64) *ActionUserRequestChangeIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
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
// SetAdmin sets parameter Admin to value and selects it for sending
func (in *ActionUserRequestChangeIndexInput) SetAdmin(value int64) *ActionUserRequestChangeIndexInput {
	in.Admin = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Admin"] = nil
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

func (in *ActionUserRequestChangeIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserRequestChangeIndexOutput is a type for action output parameters
type ActionUserRequestChangeIndexOutput struct {
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
	ChangeReason string `json:"change_reason"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
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
		Path: "/v5.0/user_request/changes",
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
		if inv.IsParameterSelected("Offset") {
			ret["change[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["change[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["change[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("State") {
			ret["change[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("ApiIpAddr") {
			ret["change[api_ip_addr]"] = inv.Input.ApiIpAddr
		}
		if inv.IsParameterSelected("ClientIpAddr") {
			ret["change[client_ip_addr]"] = inv.Input.ClientIpAddr
		}
		if inv.IsParameterSelected("Admin") {
			ret["change[admin]"] = convertInt64ToString(inv.Input.Admin)
		}
	}
}

func (inv *ActionUserRequestChangeIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

