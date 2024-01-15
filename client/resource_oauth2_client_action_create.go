package client

import ()

// ActionOauth2ClientCreate is a type for action Oauth2_client#Create
type ActionOauth2ClientCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionOauth2ClientCreate(client *Client) *ActionOauth2ClientCreate {
	return &ActionOauth2ClientCreate{
		Client: client,
	}
}

// ActionOauth2ClientCreateMetaGlobalInput is a type for action global meta input parameters
type ActionOauth2ClientCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOauth2ClientCreateMetaGlobalInput) SetIncludes(value string) *ActionOauth2ClientCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOauth2ClientCreateMetaGlobalInput) SetNo(value bool) *ActionOauth2ClientCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOauth2ClientCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOauth2ClientCreateMetaGlobalInput) SelectParameters(params ...string) *ActionOauth2ClientCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOauth2ClientCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOauth2ClientCreateInput is a type for action input parameters
type ActionOauth2ClientCreateInput struct {
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
func (in *ActionOauth2ClientCreateInput) SetAccessTokenLifetime(value string) *ActionOauth2ClientCreateInput {
	in.AccessTokenLifetime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AccessTokenLifetime"] = nil
	return in
}

// SetAccessTokenSeconds sets parameter AccessTokenSeconds to value and selects it for sending
func (in *ActionOauth2ClientCreateInput) SetAccessTokenSeconds(value int64) *ActionOauth2ClientCreateInput {
	in.AccessTokenSeconds = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AccessTokenSeconds"] = nil
	return in
}

// SetAllowSingleSignOn sets parameter AllowSingleSignOn to value and selects it for sending
func (in *ActionOauth2ClientCreateInput) SetAllowSingleSignOn(value bool) *ActionOauth2ClientCreateInput {
	in.AllowSingleSignOn = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AllowSingleSignOn"] = nil
	return in
}

// SetClientId sets parameter ClientId to value and selects it for sending
func (in *ActionOauth2ClientCreateInput) SetClientId(value string) *ActionOauth2ClientCreateInput {
	in.ClientId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientId"] = nil
	return in
}

// SetClientSecret sets parameter ClientSecret to value and selects it for sending
func (in *ActionOauth2ClientCreateInput) SetClientSecret(value string) *ActionOauth2ClientCreateInput {
	in.ClientSecret = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientSecret"] = nil
	return in
}

// SetIssueRefreshToken sets parameter IssueRefreshToken to value and selects it for sending
func (in *ActionOauth2ClientCreateInput) SetIssueRefreshToken(value bool) *ActionOauth2ClientCreateInput {
	in.IssueRefreshToken = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IssueRefreshToken"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionOauth2ClientCreateInput) SetName(value string) *ActionOauth2ClientCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetRedirectUri sets parameter RedirectUri to value and selects it for sending
func (in *ActionOauth2ClientCreateInput) SetRedirectUri(value string) *ActionOauth2ClientCreateInput {
	in.RedirectUri = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RedirectUri"] = nil
	return in
}

// SetRefreshTokenSeconds sets parameter RefreshTokenSeconds to value and selects it for sending
func (in *ActionOauth2ClientCreateInput) SetRefreshTokenSeconds(value int64) *ActionOauth2ClientCreateInput {
	in.RefreshTokenSeconds = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RefreshTokenSeconds"] = nil
	return in
}

// SelectParameters sets parameters from ActionOauth2ClientCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOauth2ClientCreateInput) SelectParameters(params ...string) *ActionOauth2ClientCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOauth2ClientCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOauth2ClientCreateInput) UnselectParameters(params ...string) *ActionOauth2ClientCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOauth2ClientCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOauth2ClientCreateRequest is a type for the entire action request
type ActionOauth2ClientCreateRequest struct {
	Oauth2Client map[string]interface{} `json:"oauth2_client"`
	Meta         map[string]interface{} `json:"_meta"`
}

// ActionOauth2ClientCreateOutput is a type for action output parameters
type ActionOauth2ClientCreateOutput struct {
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
type ActionOauth2ClientCreateResponse struct {
	Action *ActionOauth2ClientCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Oauth2Client *ActionOauth2ClientCreateOutput `json:"oauth2_client"`
	}

	// Action output without the namespace
	Output *ActionOauth2ClientCreateOutput
}

// Prepare the action for invocation
func (action *ActionOauth2ClientCreate) Prepare() *ActionOauth2ClientCreateInvocation {
	return &ActionOauth2ClientCreateInvocation{
		Action: action,
		Path:   "/v6.0/oauth2_clients",
	}
}

// ActionOauth2ClientCreateInvocation is used to configure action for invocation
type ActionOauth2ClientCreateInvocation struct {
	// Pointer to the action
	Action *ActionOauth2ClientCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOauth2ClientCreateInput
	// Global meta input parameters
	MetaInput *ActionOauth2ClientCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOauth2ClientCreateInvocation) NewInput() *ActionOauth2ClientCreateInput {
	inv.Input = &ActionOauth2ClientCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOauth2ClientCreateInvocation) SetInput(input *ActionOauth2ClientCreateInput) *ActionOauth2ClientCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOauth2ClientCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOauth2ClientCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOauth2ClientCreateInvocation) NewMetaInput() *ActionOauth2ClientCreateMetaGlobalInput {
	inv.MetaInput = &ActionOauth2ClientCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOauth2ClientCreateInvocation) SetMetaInput(input *ActionOauth2ClientCreateMetaGlobalInput) *ActionOauth2ClientCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOauth2ClientCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOauth2ClientCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOauth2ClientCreateInvocation) Call() (*ActionOauth2ClientCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOauth2ClientCreateInvocation) callAsBody() (*ActionOauth2ClientCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOauth2ClientCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Oauth2Client
	}
	return resp, err
}

func (inv *ActionOauth2ClientCreateInvocation) makeAllInputParams() *ActionOauth2ClientCreateRequest {
	return &ActionOauth2ClientCreateRequest{
		Oauth2Client: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionOauth2ClientCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionOauth2ClientCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
