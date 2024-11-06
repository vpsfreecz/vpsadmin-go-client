package client

import (
	"strings"
)

// ActionOauth2ClientShow is a type for action Oauth2_client#Show
type ActionOauth2ClientShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOauth2ClientShow(client *Client) *ActionOauth2ClientShow {
	return &ActionOauth2ClientShow{
		Client: client,
	}
}

// ActionOauth2ClientShowMetaGlobalInput is a type for action global meta input parameters
type ActionOauth2ClientShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOauth2ClientShowMetaGlobalInput) SetIncludes(value string) *ActionOauth2ClientShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOauth2ClientShowMetaGlobalInput) SetNo(value bool) *ActionOauth2ClientShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOauth2ClientShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOauth2ClientShowMetaGlobalInput) SelectParameters(params ...string) *ActionOauth2ClientShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOauth2ClientShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOauth2ClientShowOutput is a type for action output parameters
type ActionOauth2ClientShowOutput struct {
	AccessTokenLifetime string `json:"access_token_lifetime"`
	AccessTokenSeconds  int64  `json:"access_token_seconds"`
	AllowSingleSignOn   bool   `json:"allow_single_sign_on"`
	ClientId            string `json:"client_id"`
	CreatedAt           string `json:"created_at"`
	Id                  int64  `json:"id"`
	IssueRefreshToken   bool   `json:"issue_refresh_token"`
	Name                string `json:"name"`
	RedirectUri         string `json:"redirect_uri"`
	RefreshTokenSeconds int64  `json:"refresh_token_seconds"`
	UpdatedAt           string `json:"updated_at"`
}

// Type for action response, including envelope
type ActionOauth2ClientShowResponse struct {
	Action *ActionOauth2ClientShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Oauth2Client *ActionOauth2ClientShowOutput `json:"oauth2_client"`
	}

	// Action output without the namespace
	Output *ActionOauth2ClientShowOutput
}

// Prepare the action for invocation
func (action *ActionOauth2ClientShow) Prepare() *ActionOauth2ClientShowInvocation {
	return &ActionOauth2ClientShowInvocation{
		Action: action,
		Path:   "/v7.0/oauth2_clients/{oauth2_client_id}",
	}
}

// ActionOauth2ClientShowInvocation is used to configure action for invocation
type ActionOauth2ClientShowInvocation struct {
	// Pointer to the action
	Action *ActionOauth2ClientShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOauth2ClientShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOauth2ClientShowInvocation) SetPathParamInt(param string, value int64) *ActionOauth2ClientShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOauth2ClientShowInvocation) SetPathParamString(param string, value string) *ActionOauth2ClientShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOauth2ClientShowInvocation) NewMetaInput() *ActionOauth2ClientShowMetaGlobalInput {
	inv.MetaInput = &ActionOauth2ClientShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOauth2ClientShowInvocation) SetMetaInput(input *ActionOauth2ClientShowMetaGlobalInput) *ActionOauth2ClientShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOauth2ClientShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOauth2ClientShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOauth2ClientShowInvocation) Call() (*ActionOauth2ClientShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOauth2ClientShowInvocation) callAsQuery() (*ActionOauth2ClientShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOauth2ClientShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Oauth2Client
	}
	return resp, err
}

func (inv *ActionOauth2ClientShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
