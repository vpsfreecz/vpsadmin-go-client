package client

import ()

// ActionUserIndex is a type for action User#Index
type ActionUserIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserIndex(client *Client) *ActionUserIndex {
	return &ActionUserIndex{
		Client: client,
	}
}

// ActionUserIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserIndexMetaGlobalInput) SetCount(value bool) *ActionUserIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserIndexMetaGlobalInput) SetIncludes(value string) *ActionUserIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserIndexMetaGlobalInput) SetNo(value bool) *ActionUserIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserIndexInput is a type for action input parameters
type ActionUserIndexInput struct {
	Address       string `json:"address"`
	Admin         bool   `json:"admin"`
	Email         string `json:"email"`
	FullName      string `json:"full_name"`
	Info          string `json:"info"`
	Language      int64  `json:"language"`
	Level         int64  `json:"level"`
	Limit         int64  `json:"limit"`
	Lockout       bool   `json:"lockout"`
	Login         string `json:"login"`
	MailerEnabled bool   `json:"mailer_enabled"`
	ObjectState   string `json:"object_state"`
	Offset        int64  `json:"offset"`
	PasswordReset bool   `json:"password_reset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserIndexInput) SetAddress(value string) *ActionUserIndexInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}

// SetAdmin sets parameter Admin to value and selects it for sending
func (in *ActionUserIndexInput) SetAdmin(value bool) *ActionUserIndexInput {
	in.Admin = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Admin"] = nil
	return in
}

// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionUserIndexInput) SetEmail(value string) *ActionUserIndexInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}

// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserIndexInput) SetFullName(value string) *ActionUserIndexInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
	return in
}

// SetInfo sets parameter Info to value and selects it for sending
func (in *ActionUserIndexInput) SetInfo(value string) *ActionUserIndexInput {
	in.Info = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Info"] = nil
	return in
}

// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionUserIndexInput) SetLanguage(value int64) *ActionUserIndexInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLanguageNil(false)
	in._selectedParameters["Language"] = nil
	return in
}

// SetLanguageNil sets parameter Language to nil and selects it for sending
func (in *ActionUserIndexInput) SetLanguageNil(set bool) *ActionUserIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Language"] = nil
		in.SelectParameters("Language")
	} else {
		delete(in._nilParameters, "Language")
	}
	return in
}

// SetLevel sets parameter Level to value and selects it for sending
func (in *ActionUserIndexInput) SetLevel(value int64) *ActionUserIndexInput {
	in.Level = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Level"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserIndexInput) SetLimit(value int64) *ActionUserIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLockout sets parameter Lockout to value and selects it for sending
func (in *ActionUserIndexInput) SetLockout(value bool) *ActionUserIndexInput {
	in.Lockout = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Lockout"] = nil
	return in
}

// SetLogin sets parameter Login to value and selects it for sending
func (in *ActionUserIndexInput) SetLogin(value string) *ActionUserIndexInput {
	in.Login = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Login"] = nil
	return in
}

// SetMailerEnabled sets parameter MailerEnabled to value and selects it for sending
func (in *ActionUserIndexInput) SetMailerEnabled(value bool) *ActionUserIndexInput {
	in.MailerEnabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MailerEnabled"] = nil
	return in
}

// SetObjectState sets parameter ObjectState to value and selects it for sending
func (in *ActionUserIndexInput) SetObjectState(value string) *ActionUserIndexInput {
	in.ObjectState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectState"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserIndexInput) SetOffset(value int64) *ActionUserIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetPasswordReset sets parameter PasswordReset to value and selects it for sending
func (in *ActionUserIndexInput) SetPasswordReset(value bool) *ActionUserIndexInput {
	in.PasswordReset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PasswordReset"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserIndexInput) SelectParameters(params ...string) *ActionUserIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserIndexInput) UnselectParameters(params ...string) *ActionUserIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserIndexOutput is a type for action output parameters
type ActionUserIndexOutput struct {
	Address         string                    `json:"address"`
	CreatedAt       string                    `json:"created_at"`
	Email           string                    `json:"email"`
	ExpirationDate  string                    `json:"expiration_date"`
	FullName        string                    `json:"full_name"`
	Id              int64                     `json:"id"`
	Info            string                    `json:"info"`
	Language        *ActionLanguageShowOutput `json:"language"`
	LastActivityAt  string                    `json:"last_activity_at"`
	Level           int64                     `json:"level"`
	Lockout         bool                      `json:"lockout"`
	Login           string                    `json:"login"`
	MailerEnabled   bool                      `json:"mailer_enabled"`
	MonthlyPayment  int64                     `json:"monthly_payment"`
	ObjectState     string                    `json:"object_state"`
	PaidUntil       string                    `json:"paid_until"`
	PasswordReset   bool                      `json:"password_reset"`
	RemindAfterDate string                    `json:"remind_after_date"`
}

// Type for action response, including envelope
type ActionUserIndexResponse struct {
	Action *ActionUserIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Users []*ActionUserIndexOutput `json:"users"`
	}

	// Action output without the namespace
	Output []*ActionUserIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserIndex) Prepare() *ActionUserIndexInvocation {
	return &ActionUserIndexInvocation{
		Action: action,
		Path:   "/v6.0/users",
	}
}

// ActionUserIndexInvocation is used to configure action for invocation
type ActionUserIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserIndexInput
	// Global meta input parameters
	MetaInput *ActionUserIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserIndexInvocation) NewInput() *ActionUserIndexInput {
	inv.Input = &ActionUserIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserIndexInvocation) SetInput(input *ActionUserIndexInput) *ActionUserIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserIndexInvocation) NewMetaInput() *ActionUserIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserIndexInvocation) SetMetaInput(input *ActionUserIndexMetaGlobalInput) *ActionUserIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserIndexInvocation) Call() (*ActionUserIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserIndexInvocation) callAsQuery() (*ActionUserIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Users
	}
	return resp, err
}

func (inv *ActionUserIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Address") {
			ret["user[address]"] = inv.Input.Address
		}
		if inv.IsParameterSelected("Admin") {
			ret["user[admin]"] = convertBoolToString(inv.Input.Admin)
		}
		if inv.IsParameterSelected("Email") {
			ret["user[email]"] = inv.Input.Email
		}
		if inv.IsParameterSelected("FullName") {
			ret["user[full_name]"] = inv.Input.FullName
		}
		if inv.IsParameterSelected("Info") {
			ret["user[info]"] = inv.Input.Info
		}
		if inv.IsParameterSelected("Language") {
			ret["user[language]"] = convertInt64ToString(inv.Input.Language)
		}
		if inv.IsParameterSelected("Level") {
			ret["user[level]"] = convertInt64ToString(inv.Input.Level)
		}
		if inv.IsParameterSelected("Limit") {
			ret["user[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Lockout") {
			ret["user[lockout]"] = convertBoolToString(inv.Input.Lockout)
		}
		if inv.IsParameterSelected("Login") {
			ret["user[login]"] = inv.Input.Login
		}
		if inv.IsParameterSelected("MailerEnabled") {
			ret["user[mailer_enabled]"] = convertBoolToString(inv.Input.MailerEnabled)
		}
		if inv.IsParameterSelected("ObjectState") {
			ret["user[object_state]"] = inv.Input.ObjectState
		}
		if inv.IsParameterSelected("Offset") {
			ret["user[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("PasswordReset") {
			ret["user[password_reset]"] = convertBoolToString(inv.Input.PasswordReset)
		}
	}
}

func (inv *ActionUserIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
