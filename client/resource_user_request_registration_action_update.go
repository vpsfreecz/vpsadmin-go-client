package client

import (
	"strings"
)

// ActionUserRequestRegistrationUpdate is a type for action User_request.Registration#Update
type ActionUserRequestRegistrationUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestRegistrationUpdate(client *Client) *ActionUserRequestRegistrationUpdate {
	return &ActionUserRequestRegistrationUpdate{
		Client: client,
	}
}

// ActionUserRequestRegistrationUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestRegistrationUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserRequestRegistrationUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateMetaGlobalInput) SetNo(value bool) *ActionUserRequestRegistrationUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestRegistrationUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationUpdateInput is a type for action input parameters
type ActionUserRequestRegistrationUpdateInput struct {
	Address     string `json:"address"`
	Currency    string `json:"currency"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	How         string `json:"how"`
	Language    int64  `json:"language"`
	Location    int64  `json:"location"`
	Login       string `json:"login"`
	Note        string `json:"note"`
	OrgId       string `json:"org_id"`
	OrgName     string `json:"org_name"`
	OsTemplate  int64  `json:"os_template"`
	YearOfBirth int64  `json:"year_of_birth"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetAddress(value string) *ActionUserRequestRegistrationUpdateInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}

// SetCurrency sets parameter Currency to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetCurrency(value string) *ActionUserRequestRegistrationUpdateInput {
	in.Currency = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Currency"] = nil
	return in
}

// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetEmail(value string) *ActionUserRequestRegistrationUpdateInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}

// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetFullName(value string) *ActionUserRequestRegistrationUpdateInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
	return in
}

// SetHow sets parameter How to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetHow(value string) *ActionUserRequestRegistrationUpdateInput {
	in.How = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["How"] = nil
	return in
}

// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetLanguage(value int64) *ActionUserRequestRegistrationUpdateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLanguageNil(false)
	in._selectedParameters["Language"] = nil
	return in
}

// SetLanguageNil sets parameter Language to nil and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetLanguageNil(set bool) *ActionUserRequestRegistrationUpdateInput {
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

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetLocation(value int64) *ActionUserRequestRegistrationUpdateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetLocationNil(set bool) *ActionUserRequestRegistrationUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Location"] = nil
		in.SelectParameters("Location")
	} else {
		delete(in._nilParameters, "Location")
	}
	return in
}

// SetLogin sets parameter Login to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetLogin(value string) *ActionUserRequestRegistrationUpdateInput {
	in.Login = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Login"] = nil
	return in
}

// SetNote sets parameter Note to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetNote(value string) *ActionUserRequestRegistrationUpdateInput {
	in.Note = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Note"] = nil
	return in
}

// SetOrgId sets parameter OrgId to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetOrgId(value string) *ActionUserRequestRegistrationUpdateInput {
	in.OrgId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OrgId"] = nil
	return in
}

// SetOrgName sets parameter OrgName to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetOrgName(value string) *ActionUserRequestRegistrationUpdateInput {
	in.OrgName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OrgName"] = nil
	return in
}

// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetOsTemplate(value int64) *ActionUserRequestRegistrationUpdateInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetOsTemplateNil(false)
	in._selectedParameters["OsTemplate"] = nil
	return in
}

// SetOsTemplateNil sets parameter OsTemplate to nil and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetOsTemplateNil(set bool) *ActionUserRequestRegistrationUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["OsTemplate"] = nil
		in.SelectParameters("OsTemplate")
	} else {
		delete(in._nilParameters, "OsTemplate")
	}
	return in
}

// SetYearOfBirth sets parameter YearOfBirth to value and selects it for sending
func (in *ActionUserRequestRegistrationUpdateInput) SetYearOfBirth(value int64) *ActionUserRequestRegistrationUpdateInput {
	in.YearOfBirth = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["YearOfBirth"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationUpdateInput) SelectParameters(params ...string) *ActionUserRequestRegistrationUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserRequestRegistrationUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationUpdateInput) UnselectParameters(params ...string) *ActionUserRequestRegistrationUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserRequestRegistrationUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationUpdateRequest is a type for the entire action request
type ActionUserRequestRegistrationUpdateRequest struct {
	Registration map[string]interface{} `json:"registration"`
	Meta         map[string]interface{} `json:"_meta"`
}

// ActionUserRequestRegistrationUpdateOutput is a type for action output parameters
type ActionUserRequestRegistrationUpdateOutput struct {
	Address     string                      `json:"address"`
	Currency    string                      `json:"currency"`
	Email       string                      `json:"email"`
	FullName    string                      `json:"full_name"`
	How         string                      `json:"how"`
	Id          int64                       `json:"id"`
	Language    *ActionLanguageShowOutput   `json:"language"`
	Location    *ActionLocationShowOutput   `json:"location"`
	Login       string                      `json:"login"`
	Note        string                      `json:"note"`
	OrgId       string                      `json:"org_id"`
	OrgName     string                      `json:"org_name"`
	OsTemplate  *ActionOsTemplateShowOutput `json:"os_template"`
	YearOfBirth int64                       `json:"year_of_birth"`
}

// Type for action response, including envelope
type ActionUserRequestRegistrationUpdateResponse struct {
	Action *ActionUserRequestRegistrationUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Registration *ActionUserRequestRegistrationUpdateOutput `json:"registration"`
	}

	// Action output without the namespace
	Output *ActionUserRequestRegistrationUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserRequestRegistrationUpdate) Prepare() *ActionUserRequestRegistrationUpdateInvocation {
	return &ActionUserRequestRegistrationUpdateInvocation{
		Action: action,
		Path:   "/v7.0/user_request/registrations/{registration_id}/{token}",
	}
}

// ActionUserRequestRegistrationUpdateInvocation is used to configure action for invocation
type ActionUserRequestRegistrationUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestRegistrationUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserRequestRegistrationUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserRequestRegistrationUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserRequestRegistrationUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserRequestRegistrationUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserRequestRegistrationUpdateInvocation) SetPathParamString(param string, value string) *ActionUserRequestRegistrationUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserRequestRegistrationUpdateInvocation) NewInput() *ActionUserRequestRegistrationUpdateInput {
	inv.Input = &ActionUserRequestRegistrationUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserRequestRegistrationUpdateInvocation) SetInput(input *ActionUserRequestRegistrationUpdateInput) *ActionUserRequestRegistrationUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserRequestRegistrationUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserRequestRegistrationUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestRegistrationUpdateInvocation) NewMetaInput() *ActionUserRequestRegistrationUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestRegistrationUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestRegistrationUpdateInvocation) SetMetaInput(input *ActionUserRequestRegistrationUpdateMetaGlobalInput) *ActionUserRequestRegistrationUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestRegistrationUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserRequestRegistrationUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestRegistrationUpdateInvocation) Call() (*ActionUserRequestRegistrationUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionUserRequestRegistrationUpdateInvocation) callAsBody() (*ActionUserRequestRegistrationUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserRequestRegistrationUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Registration
	}
	return resp, err
}

func (inv *ActionUserRequestRegistrationUpdateInvocation) makeAllInputParams() *ActionUserRequestRegistrationUpdateRequest {
	return &ActionUserRequestRegistrationUpdateRequest{
		Registration: inv.makeInputParams(),
		Meta:         inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserRequestRegistrationUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Address") {
			ret["address"] = inv.Input.Address
		}
		if inv.IsParameterSelected("Currency") {
			ret["currency"] = inv.Input.Currency
		}
		if inv.IsParameterSelected("Email") {
			ret["email"] = inv.Input.Email
		}
		if inv.IsParameterSelected("FullName") {
			ret["full_name"] = inv.Input.FullName
		}
		if inv.IsParameterSelected("How") {
			ret["how"] = inv.Input.How
		}
		if inv.IsParameterSelected("Language") {
			if inv.IsParameterNil("Language") {
				ret["language"] = nil
			} else {
				ret["language"] = inv.Input.Language
			}
		}
		if inv.IsParameterSelected("Location") {
			if inv.IsParameterNil("Location") {
				ret["location"] = nil
			} else {
				ret["location"] = inv.Input.Location
			}
		}
		if inv.IsParameterSelected("Login") {
			ret["login"] = inv.Input.Login
		}
		if inv.IsParameterSelected("Note") {
			ret["note"] = inv.Input.Note
		}
		if inv.IsParameterSelected("OrgId") {
			ret["org_id"] = inv.Input.OrgId
		}
		if inv.IsParameterSelected("OrgName") {
			ret["org_name"] = inv.Input.OrgName
		}
		if inv.IsParameterSelected("OsTemplate") {
			if inv.IsParameterNil("OsTemplate") {
				ret["os_template"] = nil
			} else {
				ret["os_template"] = inv.Input.OsTemplate
			}
		}
		if inv.IsParameterSelected("YearOfBirth") {
			ret["year_of_birth"] = inv.Input.YearOfBirth
		}
	}

	return ret
}

func (inv *ActionUserRequestRegistrationUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
