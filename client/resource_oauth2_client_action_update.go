package client

import (
	"strings"
)

// ActionOauth2ClientUpdate is a type for action Oauth2_client#Update
type ActionOauth2ClientUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionOauth2ClientUpdate(client *Client) *ActionOauth2ClientUpdate {
	return &ActionOauth2ClientUpdate{
		Client: client,
	}
}

// ActionOauth2ClientUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionOauth2ClientUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOauth2ClientUpdateMetaGlobalInput) SetIncludes(value string) *ActionOauth2ClientUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOauth2ClientUpdateMetaGlobalInput) SetNo(value bool) *ActionOauth2ClientUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOauth2ClientUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOauth2ClientUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionOauth2ClientUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOauth2ClientUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOauth2ClientUpdateInput is a type for action input parameters
type ActionOauth2ClientUpdateInput struct {
	AccessTokenLifetime string `json:"access_token_lifetime"`
	AccessTokenSeconds  int64  `json:"access_token_seconds"`
	AllowSingleSignOn   bool   `json:"allow_single_sign_on"`
	ClientId            string `json:"client_id"`
	ClientSecret        string `json:"client_secret"`
	IssueRefreshToken   bool   `json:"issue_refresh_token"`
	Name                string `json:"name"`
	RedirectUri         string `json:"redirect_uri"`
	RefreshTokenSeconds int64  `json:"refresh_token_seconds"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAccessTokenLifetime sets parameter AccessTokenLifetime to value and selects it for sending
func (in *ActionOauth2ClientUpdateInput) SetAccessTokenLifetime(value string) *ActionOauth2ClientUpdateInput {
	in.AccessTokenLifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AccessTokenLifetime"] = nil
	return in
}

// SetAccessTokenSeconds sets parameter AccessTokenSeconds to value and selects it for sending
func (in *ActionOauth2ClientUpdateInput) SetAccessTokenSeconds(value int64) *ActionOauth2ClientUpdateInput {
	in.AccessTokenSeconds = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AccessTokenSeconds"] = nil
	return in
}

// SetAllowSingleSignOn sets parameter AllowSingleSignOn to value and selects it for sending
func (in *ActionOauth2ClientUpdateInput) SetAllowSingleSignOn(value bool) *ActionOauth2ClientUpdateInput {
	in.AllowSingleSignOn = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AllowSingleSignOn"] = nil
	return in
}

// SetClientId sets parameter ClientId to value and selects it for sending
func (in *ActionOauth2ClientUpdateInput) SetClientId(value string) *ActionOauth2ClientUpdateInput {
	in.ClientId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientId"] = nil
	return in
}

// SetClientSecret sets parameter ClientSecret to value and selects it for sending
func (in *ActionOauth2ClientUpdateInput) SetClientSecret(value string) *ActionOauth2ClientUpdateInput {
	in.ClientSecret = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientSecret"] = nil
	return in
}

// SetIssueRefreshToken sets parameter IssueRefreshToken to value and selects it for sending
func (in *ActionOauth2ClientUpdateInput) SetIssueRefreshToken(value bool) *ActionOauth2ClientUpdateInput {
	in.IssueRefreshToken = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IssueRefreshToken"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionOauth2ClientUpdateInput) SetName(value string) *ActionOauth2ClientUpdateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetRedirectUri sets parameter RedirectUri to value and selects it for sending
func (in *ActionOauth2ClientUpdateInput) SetRedirectUri(value string) *ActionOauth2ClientUpdateInput {
	in.RedirectUri = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RedirectUri"] = nil
	return in
}

// SetRefreshTokenSeconds sets parameter RefreshTokenSeconds to value and selects it for sending
func (in *ActionOauth2ClientUpdateInput) SetRefreshTokenSeconds(value int64) *ActionOauth2ClientUpdateInput {
	in.RefreshTokenSeconds = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RefreshTokenSeconds"] = nil
	return in
}

// SelectParameters sets parameters from ActionOauth2ClientUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOauth2ClientUpdateInput) SelectParameters(params ...string) *ActionOauth2ClientUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOauth2ClientUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOauth2ClientUpdateInput) UnselectParameters(params ...string) *ActionOauth2ClientUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOauth2ClientUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOauth2ClientUpdateRequest is a type for the entire action request
type ActionOauth2ClientUpdateRequest struct {
	Oauth2Client map[string]interface{} `json:"oauth2_client"`
	Meta         map[string]interface{} `json:"_meta"`
}

// ActionOauth2ClientUpdateOutput is a type for action output parameters
type ActionOauth2ClientUpdateOutput struct {
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
type ActionOauth2ClientUpdateResponse struct {
	Action *ActionOauth2ClientUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Oauth2Client *ActionOauth2ClientUpdateOutput `json:"oauth2_client"`
	}

	// Action output without the namespace
	Output *ActionOauth2ClientUpdateOutput
}

// Prepare the action for invocation
func (action *ActionOauth2ClientUpdate) Prepare() *ActionOauth2ClientUpdateInvocation {
	return &ActionOauth2ClientUpdateInvocation{
		Action: action,
		Path:   "/v7.0/oauth2_clients/{oauth2_client_id}",
	}
}

// ActionOauth2ClientUpdateInvocation is used to configure action for invocation
type ActionOauth2ClientUpdateInvocation struct {
	// Pointer to the action
	Action *ActionOauth2ClientUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOauth2ClientUpdateInput
	// Global meta input parameters
	MetaInput *ActionOauth2ClientUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOauth2ClientUpdateInvocation) SetPathParamInt(param string, value int64) *ActionOauth2ClientUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOauth2ClientUpdateInvocation) SetPathParamString(param string, value string) *ActionOauth2ClientUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOauth2ClientUpdateInvocation) NewInput() *ActionOauth2ClientUpdateInput {
	inv.Input = &ActionOauth2ClientUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOauth2ClientUpdateInvocation) SetInput(input *ActionOauth2ClientUpdateInput) *ActionOauth2ClientUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOauth2ClientUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOauth2ClientUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOauth2ClientUpdateInvocation) NewMetaInput() *ActionOauth2ClientUpdateMetaGlobalInput {
	inv.MetaInput = &ActionOauth2ClientUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOauth2ClientUpdateInvocation) SetMetaInput(input *ActionOauth2ClientUpdateMetaGlobalInput) *ActionOauth2ClientUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOauth2ClientUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOauth2ClientUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOauth2ClientUpdateInvocation) Call() (*ActionOauth2ClientUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOauth2ClientUpdateInvocation) callAsBody() (*ActionOauth2ClientUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOauth2ClientUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Oauth2Client
	}
	return resp, err
}

func (inv *ActionOauth2ClientUpdateInvocation) makeAllInputParams() *ActionOauth2ClientUpdateRequest {
	return &ActionOauth2ClientUpdateRequest{
		Oauth2Client: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionOauth2ClientUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("AccessTokenLifetime") {
			ret["access_token_lifetime"] = inv.Input.AccessTokenLifetime
		}
		if inv.IsParameterSelected("AccessTokenSeconds") {
			ret["access_token_seconds"] = inv.Input.AccessTokenSeconds
		}
		if inv.IsParameterSelected("AllowSingleSignOn") {
			ret["allow_single_sign_on"] = inv.Input.AllowSingleSignOn
		}
		if inv.IsParameterSelected("ClientId") {
			ret["client_id"] = inv.Input.ClientId
		}
		if inv.IsParameterSelected("ClientSecret") {
			ret["client_secret"] = inv.Input.ClientSecret
		}
		if inv.IsParameterSelected("IssueRefreshToken") {
			ret["issue_refresh_token"] = inv.Input.IssueRefreshToken
		}
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("RedirectUri") {
			ret["redirect_uri"] = inv.Input.RedirectUri
		}
		if inv.IsParameterSelected("RefreshTokenSeconds") {
			ret["refresh_token_seconds"] = inv.Input.RefreshTokenSeconds
		}
	}

	return ret
}

func (inv *ActionOauth2ClientUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
