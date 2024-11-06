package client

import ()

// ActionOauth2ClientIndex is a type for action Oauth2_client#Index
type ActionOauth2ClientIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOauth2ClientIndex(client *Client) *ActionOauth2ClientIndex {
	return &ActionOauth2ClientIndex{
		Client: client,
	}
}

// ActionOauth2ClientIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOauth2ClientIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOauth2ClientIndexMetaGlobalInput) SetCount(value bool) *ActionOauth2ClientIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOauth2ClientIndexMetaGlobalInput) SetIncludes(value string) *ActionOauth2ClientIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOauth2ClientIndexMetaGlobalInput) SetNo(value bool) *ActionOauth2ClientIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOauth2ClientIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOauth2ClientIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOauth2ClientIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOauth2ClientIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOauth2ClientIndexInput is a type for action input parameters
type ActionOauth2ClientIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionOauth2ClientIndexInput) SetFromId(value int64) *ActionOauth2ClientIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOauth2ClientIndexInput) SetLimit(value int64) *ActionOauth2ClientIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionOauth2ClientIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOauth2ClientIndexInput) SelectParameters(params ...string) *ActionOauth2ClientIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOauth2ClientIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOauth2ClientIndexInput) UnselectParameters(params ...string) *ActionOauth2ClientIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOauth2ClientIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOauth2ClientIndexOutput is a type for action output parameters
type ActionOauth2ClientIndexOutput struct {
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
type ActionOauth2ClientIndexResponse struct {
	Action *ActionOauth2ClientIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Oauth2Clients []*ActionOauth2ClientIndexOutput `json:"oauth2_clients"`
	}

	// Action output without the namespace
	Output []*ActionOauth2ClientIndexOutput
}

// Prepare the action for invocation
func (action *ActionOauth2ClientIndex) Prepare() *ActionOauth2ClientIndexInvocation {
	return &ActionOauth2ClientIndexInvocation{
		Action: action,
		Path:   "/v7.0/oauth2_clients",
	}
}

// ActionOauth2ClientIndexInvocation is used to configure action for invocation
type ActionOauth2ClientIndexInvocation struct {
	// Pointer to the action
	Action *ActionOauth2ClientIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOauth2ClientIndexInput
	// Global meta input parameters
	MetaInput *ActionOauth2ClientIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOauth2ClientIndexInvocation) NewInput() *ActionOauth2ClientIndexInput {
	inv.Input = &ActionOauth2ClientIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOauth2ClientIndexInvocation) SetInput(input *ActionOauth2ClientIndexInput) *ActionOauth2ClientIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOauth2ClientIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOauth2ClientIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOauth2ClientIndexInvocation) NewMetaInput() *ActionOauth2ClientIndexMetaGlobalInput {
	inv.MetaInput = &ActionOauth2ClientIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOauth2ClientIndexInvocation) SetMetaInput(input *ActionOauth2ClientIndexMetaGlobalInput) *ActionOauth2ClientIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOauth2ClientIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOauth2ClientIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOauth2ClientIndexInvocation) Call() (*ActionOauth2ClientIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOauth2ClientIndexInvocation) callAsQuery() (*ActionOauth2ClientIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOauth2ClientIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Oauth2Clients
	}
	return resp, err
}

func (inv *ActionOauth2ClientIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["oauth2_client[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["oauth2_client[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionOauth2ClientIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
