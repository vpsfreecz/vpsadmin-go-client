package client

import (
)

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
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Login string `json:"login"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
	Level int64 `json:"level"`
	Info string `json:"info"`
	MailerEnabled bool `json:"mailer_enabled"`
	Language int64 `json:"language"`
	Admin bool `json:"admin"`
	ObjectState string `json:"object_state"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserIndexInput) SetLimit(value int64) *ActionUserIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
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
// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserIndexInput) SetFullName(value string) *ActionUserIndexInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
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
// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserIndexInput) SetAddress(value string) *ActionUserIndexInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
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
// SetInfo sets parameter Info to value and selects it for sending
func (in *ActionUserIndexInput) SetInfo(value string) *ActionUserIndexInput {
	in.Info = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Info"] = nil
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
// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionUserIndexInput) SetLanguage(value int64) *ActionUserIndexInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Language"] = nil
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
// SetObjectState sets parameter ObjectState to value and selects it for sending
func (in *ActionUserIndexInput) SetObjectState(value string) *ActionUserIndexInput {
	in.ObjectState = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectState"] = nil
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

func (in *ActionUserIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserIndexOutput is a type for action output parameters
type ActionUserIndexOutput struct {
	Id int64 `json:"id"`
	Login string `json:"login"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
	Level int64 `json:"level"`
	Info string `json:"info"`
	MailerEnabled bool `json:"mailer_enabled"`
	Language *ActionLanguageShowOutput `json:"language"`
	LastActivityAt string `json:"last_activity_at"`
	CreatedAt string `json:"created_at"`
	ObjectState string `json:"object_state"`
	ExpirationDate string `json:"expiration_date"`
	MonthlyPayment int64 `json:"monthly_payment"`
	PaidUntil string `json:"paid_until"`
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
		Path: "/v5.0/users",
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
		if inv.IsParameterSelected("Offset") {
			ret["user[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["user[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Login") {
			ret["user[login]"] = inv.Input.Login
		}
		if inv.IsParameterSelected("FullName") {
			ret["user[full_name]"] = inv.Input.FullName
		}
		if inv.IsParameterSelected("Email") {
			ret["user[email]"] = inv.Input.Email
		}
		if inv.IsParameterSelected("Address") {
			ret["user[address]"] = inv.Input.Address
		}
		if inv.IsParameterSelected("Level") {
			ret["user[level]"] = convertInt64ToString(inv.Input.Level)
		}
		if inv.IsParameterSelected("Info") {
			ret["user[info]"] = inv.Input.Info
		}
		if inv.IsParameterSelected("MailerEnabled") {
			ret["user[mailer_enabled]"] = convertBoolToString(inv.Input.MailerEnabled)
		}
		if inv.IsParameterSelected("Language") {
			ret["user[language]"] = convertInt64ToString(inv.Input.Language)
		}
		if inv.IsParameterSelected("Admin") {
			ret["user[admin]"] = convertBoolToString(inv.Input.Admin)
		}
		if inv.IsParameterSelected("ObjectState") {
			ret["user[object_state]"] = inv.Input.ObjectState
		}
	}
}

func (inv *ActionUserIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

