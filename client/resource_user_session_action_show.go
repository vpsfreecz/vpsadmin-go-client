package client

import (
	"strings"
)

// ActionUserSessionShow is a type for action User_session#Show
type ActionUserSessionShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserSessionShow(client *Client) *ActionUserSessionShow {
	return &ActionUserSessionShow{
		Client: client,
	}
}

// ActionUserSessionShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserSessionShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserSessionShowMetaGlobalInput) SetNo(value bool) *ActionUserSessionShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserSessionShowMetaGlobalInput) SetIncludes(value string) *ActionUserSessionShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserSessionShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSessionShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserSessionShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserSessionShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserSessionShowOutput is a type for action output parameters
type ActionUserSessionShowOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	AuthType string `json:"auth_type"`
	ApiIpAddr string `json:"api_ip_addr"`
	ApiIpPtr string `json:"api_ip_ptr"`
	ClientIpAddr string `json:"client_ip_addr"`
	ClientIpPtr string `json:"client_ip_ptr"`
	UserAgent string `json:"user_agent"`
	ClientVersion string `json:"client_version"`
	SessionToken *ActionSessionTokenShowOutput `json:"session_token"`
	SessionTokenStr string `json:"session_token_str"`
	CreatedAt string `json:"created_at"`
	LastRequestAt string `json:"last_request_at"`
	ClosedAt string `json:"closed_at"`
	Admin *ActionUserShowOutput `json:"admin"`
}


// Type for action response, including envelope
type ActionUserSessionShowResponse struct {
	Action *ActionUserSessionShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserSession *ActionUserSessionShowOutput `json:"user_session"`
	}

	// Action output without the namespace
	Output *ActionUserSessionShowOutput
}


// Prepare the action for invocation
func (action *ActionUserSessionShow) Prepare() *ActionUserSessionShowInvocation {
	return &ActionUserSessionShowInvocation{
		Action: action,
		Path: "/v6.0/user_sessions/{user_session_id}",
	}
}

// ActionUserSessionShowInvocation is used to configure action for invocation
type ActionUserSessionShowInvocation struct {
	// Pointer to the action
	Action *ActionUserSessionShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserSessionShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserSessionShowInvocation) SetPathParamInt(param string, value int64) *ActionUserSessionShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserSessionShowInvocation) SetPathParamString(param string, value string) *ActionUserSessionShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserSessionShowInvocation) NewMetaInput() *ActionUserSessionShowMetaGlobalInput {
	inv.MetaInput = &ActionUserSessionShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserSessionShowInvocation) SetMetaInput(input *ActionUserSessionShowMetaGlobalInput) *ActionUserSessionShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserSessionShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserSessionShowInvocation) Call() (*ActionUserSessionShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserSessionShowInvocation) callAsQuery() (*ActionUserSessionShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserSessionShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserSession
	}
	return resp, err
}




func (inv *ActionUserSessionShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

