package client

import ()

// ActionUserCurrent is a type for action User#Current
type ActionUserCurrent struct {
	// Pointer to client
	Client *Client
}

func NewActionUserCurrent(client *Client) *ActionUserCurrent {
	return &ActionUserCurrent{
		Client: client,
	}
}

// ActionUserCurrentMetaGlobalInput is a type for action global meta input parameters
type ActionUserCurrentMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserCurrentMetaGlobalInput) SetIncludes(value string) *ActionUserCurrentMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserCurrentMetaGlobalInput) SetNo(value bool) *ActionUserCurrentMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserCurrentMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserCurrentMetaGlobalInput) SelectParameters(params ...string) *ActionUserCurrentMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserCurrentMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserCurrentOutput is a type for action output parameters
type ActionUserCurrentOutput struct {
	Address                    string                    `json:"address"`
	CreatedAt                  string                    `json:"created_at"`
	DokuwikiGroups             string                    `json:"dokuwiki_groups"`
	Email                      string                    `json:"email"`
	EnableBasicAuth            bool                      `json:"enable_basic_auth"`
	EnableNewLoginNotification bool                      `json:"enable_new_login_notification"`
	EnableOauth2Auth           bool                      `json:"enable_oauth2_auth"`
	EnableSingleSignOn         bool                      `json:"enable_single_sign_on"`
	EnableTokenAuth            bool                      `json:"enable_token_auth"`
	ExpirationDate             string                    `json:"expiration_date"`
	FullName                   string                    `json:"full_name"`
	Id                         int64                     `json:"id"`
	Info                       string                    `json:"info"`
	Language                   *ActionLanguageShowOutput `json:"language"`
	LastActivityAt             string                    `json:"last_activity_at"`
	Level                      int64                     `json:"level"`
	Lockout                    bool                      `json:"lockout"`
	Login                      string                    `json:"login"`
	MailerEnabled              bool                      `json:"mailer_enabled"`
	MonthlyPayment             int64                     `json:"monthly_payment"`
	ObjectState                string                    `json:"object_state"`
	PaidUntil                  string                    `json:"paid_until"`
	PasswordReset              bool                      `json:"password_reset"`
	PreferredLogoutAll         bool                      `json:"preferred_logout_all"`
	PreferredSessionLength     int64                     `json:"preferred_session_length"`
	RemindAfterDate            string                    `json:"remind_after_date"`
}

// Type for action response, including envelope
type ActionUserCurrentResponse struct {
	Action *ActionUserCurrent `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		User *ActionUserCurrentOutput `json:"user"`
	}

	// Action output without the namespace
	Output *ActionUserCurrentOutput
}

// Call the action directly without any path or input parameters
func (action *ActionUserCurrent) Call() (*ActionUserCurrentResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *ActionUserCurrent) Prepare() *ActionUserCurrentInvocation {
	return &ActionUserCurrentInvocation{
		Action: action,
		Path:   "/v7.0/users/current",
	}
}

// ActionUserCurrentInvocation is used to configure action for invocation
type ActionUserCurrentInvocation struct {
	// Pointer to the action
	Action *ActionUserCurrent

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserCurrentMetaGlobalInput
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserCurrentInvocation) NewMetaInput() *ActionUserCurrentMetaGlobalInput {
	inv.MetaInput = &ActionUserCurrentMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserCurrentInvocation) SetMetaInput(input *ActionUserCurrentMetaGlobalInput) *ActionUserCurrentInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserCurrentInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserCurrentInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserCurrentInvocation) Call() (*ActionUserCurrentResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserCurrentInvocation) callAsQuery() (*ActionUserCurrentResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserCurrentResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.User
	}
	return resp, err
}

func (inv *ActionUserCurrentInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
