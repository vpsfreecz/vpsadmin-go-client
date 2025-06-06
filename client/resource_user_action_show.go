package client

import (
	"strings"
)

// ActionUserShow is a type for action User#Show
type ActionUserShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserShow(client *Client) *ActionUserShow {
	return &ActionUserShow{
		Client: client,
	}
}

// ActionUserShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserShowMetaGlobalInput) SetIncludes(value string) *ActionUserShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserShowMetaGlobalInput) SetNo(value bool) *ActionUserShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserShowOutput is a type for action output parameters
type ActionUserShowOutput struct {
	Address                    string                    `json:"address"`
	CreatedAt                  string                    `json:"created_at"`
	DokuwikiGroups             string                    `json:"dokuwiki_groups"`
	Email                      string                    `json:"email"`
	EnableBasicAuth            bool                      `json:"enable_basic_auth"`
	EnableMultiFactorAuth      bool                      `json:"enable_multi_factor_auth"`
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
type ActionUserShowResponse struct {
	Action *ActionUserShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		User *ActionUserShowOutput `json:"user"`
	}

	// Action output without the namespace
	Output *ActionUserShowOutput
}

// Prepare the action for invocation
func (action *ActionUserShow) Prepare() *ActionUserShowInvocation {
	return &ActionUserShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}",
	}
}

// ActionUserShowInvocation is used to configure action for invocation
type ActionUserShowInvocation struct {
	// Pointer to the action
	Action *ActionUserShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserShowInvocation) SetPathParamInt(param string, value int64) *ActionUserShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserShowInvocation) SetPathParamString(param string, value string) *ActionUserShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserShowInvocation) NewMetaInput() *ActionUserShowMetaGlobalInput {
	inv.MetaInput = &ActionUserShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserShowInvocation) SetMetaInput(input *ActionUserShowMetaGlobalInput) *ActionUserShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserShowInvocation) Call() (*ActionUserShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserShowInvocation) callAsQuery() (*ActionUserShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.User
	}
	return resp, err
}

func (inv *ActionUserShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
