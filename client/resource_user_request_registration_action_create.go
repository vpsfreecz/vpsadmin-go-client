package client

import (
)

// ActionUserRequestRegistrationCreate is a type for action User_request.Registration#Create
type ActionUserRequestRegistrationCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestRegistrationCreate(client *Client) *ActionUserRequestRegistrationCreate {
	return &ActionUserRequestRegistrationCreate{
		Client: client,
	}
}

// ActionUserRequestRegistrationCreateMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestRegistrationCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateMetaGlobalInput) SetNo(value bool) *ActionUserRequestRegistrationCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateMetaGlobalInput) SetIncludes(value string) *ActionUserRequestRegistrationCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationCreateMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestRegistrationCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationCreateInput is a type for action input parameters
type ActionUserRequestRegistrationCreateInput struct {
	Login string `json:"login"`
	FullName string `json:"full_name"`
	OrgName string `json:"org_name"`
	OrgId string `json:"org_id"`
	Email string `json:"email"`
	Address string `json:"address"`
	YearOfBirth int64 `json:"year_of_birth"`
	How string `json:"how"`
	Note string `json:"note"`
	OsTemplate int64 `json:"os_template"`
	Location int64 `json:"location"`
	Currency string `json:"currency"`
	Language int64 `json:"language"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLogin sets parameter Login to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetLogin(value string) *ActionUserRequestRegistrationCreateInput {
	in.Login = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Login"] = nil
	return in
}
// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetFullName(value string) *ActionUserRequestRegistrationCreateInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
	return in
}
// SetOrgName sets parameter OrgName to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetOrgName(value string) *ActionUserRequestRegistrationCreateInput {
	in.OrgName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OrgName"] = nil
	return in
}
// SetOrgId sets parameter OrgId to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetOrgId(value string) *ActionUserRequestRegistrationCreateInput {
	in.OrgId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OrgId"] = nil
	return in
}
// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetEmail(value string) *ActionUserRequestRegistrationCreateInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}
// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetAddress(value string) *ActionUserRequestRegistrationCreateInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}
// SetYearOfBirth sets parameter YearOfBirth to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetYearOfBirth(value int64) *ActionUserRequestRegistrationCreateInput {
	in.YearOfBirth = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["YearOfBirth"] = nil
	return in
}
// SetHow sets parameter How to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetHow(value string) *ActionUserRequestRegistrationCreateInput {
	in.How = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["How"] = nil
	return in
}
// SetNote sets parameter Note to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetNote(value string) *ActionUserRequestRegistrationCreateInput {
	in.Note = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Note"] = nil
	return in
}
// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetOsTemplate(value int64) *ActionUserRequestRegistrationCreateInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OsTemplate"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetLocation(value int64) *ActionUserRequestRegistrationCreateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetCurrency sets parameter Currency to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetCurrency(value string) *ActionUserRequestRegistrationCreateInput {
	in.Currency = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Currency"] = nil
	return in
}
// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionUserRequestRegistrationCreateInput) SetLanguage(value int64) *ActionUserRequestRegistrationCreateInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Language"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationCreateInput) SelectParameters(params ...string) *ActionUserRequestRegistrationCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationCreateRequest is a type for the entire action request
type ActionUserRequestRegistrationCreateRequest struct {
	Registration map[string]interface{} `json:"registration"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserRequestRegistrationCreateOutput is a type for action output parameters
type ActionUserRequestRegistrationCreateOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	State string `json:"state"`
	ApiIpAddr string `json:"api_ip_addr"`
	ApiIpPtr string `json:"api_ip_ptr"`
	ClientIpAddr string `json:"client_ip_addr"`
	ClientIpPtr string `json:"client_ip_ptr"`
	Admin *ActionUserShowOutput `json:"admin"`
	AdminResponse string `json:"admin_response"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Label string `json:"label"`
	Login string `json:"login"`
	FullName string `json:"full_name"`
	OrgName string `json:"org_name"`
	OrgId string `json:"org_id"`
	Email string `json:"email"`
	Address string `json:"address"`
	YearOfBirth int64 `json:"year_of_birth"`
	How string `json:"how"`
	Note string `json:"note"`
	OsTemplate *ActionOsTemplateShowOutput `json:"os_template"`
	Location *ActionLocationShowOutput `json:"location"`
	Currency string `json:"currency"`
	Language *ActionLanguageShowOutput `json:"language"`
}


// Type for action response, including envelope
type ActionUserRequestRegistrationCreateResponse struct {
	Action *ActionUserRequestRegistrationCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Registration *ActionUserRequestRegistrationCreateOutput `json:"registration"`
	}

	// Action output without the namespace
	Output *ActionUserRequestRegistrationCreateOutput
}


// Prepare the action for invocation
func (action *ActionUserRequestRegistrationCreate) Prepare() *ActionUserRequestRegistrationCreateInvocation {
	return &ActionUserRequestRegistrationCreateInvocation{
		Action: action,
		Path: "/v5.0/user_request/registrations",
	}
}

// ActionUserRequestRegistrationCreateInvocation is used to configure action for invocation
type ActionUserRequestRegistrationCreateInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestRegistrationCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserRequestRegistrationCreateInput
	// Global meta input parameters
	MetaInput *ActionUserRequestRegistrationCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserRequestRegistrationCreateInvocation) NewInput() *ActionUserRequestRegistrationCreateInput {
	inv.Input = &ActionUserRequestRegistrationCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserRequestRegistrationCreateInvocation) SetInput(input *ActionUserRequestRegistrationCreateInput) *ActionUserRequestRegistrationCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserRequestRegistrationCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestRegistrationCreateInvocation) NewMetaInput() *ActionUserRequestRegistrationCreateMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestRegistrationCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestRegistrationCreateInvocation) SetMetaInput(input *ActionUserRequestRegistrationCreateMetaGlobalInput) *ActionUserRequestRegistrationCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestRegistrationCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestRegistrationCreateInvocation) Call() (*ActionUserRequestRegistrationCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserRequestRegistrationCreateInvocation) callAsBody() (*ActionUserRequestRegistrationCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserRequestRegistrationCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Registration
	}
	return resp, err
}




func (inv *ActionUserRequestRegistrationCreateInvocation) makeAllInputParams() *ActionUserRequestRegistrationCreateRequest {
	return &ActionUserRequestRegistrationCreateRequest{
		Registration: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserRequestRegistrationCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Login") {
			ret["login"] = inv.Input.Login
		}
		if inv.IsParameterSelected("FullName") {
			ret["full_name"] = inv.Input.FullName
		}
		if inv.IsParameterSelected("OrgName") {
			ret["org_name"] = inv.Input.OrgName
		}
		if inv.IsParameterSelected("OrgId") {
			ret["org_id"] = inv.Input.OrgId
		}
		if inv.IsParameterSelected("Email") {
			ret["email"] = inv.Input.Email
		}
		if inv.IsParameterSelected("Address") {
			ret["address"] = inv.Input.Address
		}
		if inv.IsParameterSelected("YearOfBirth") {
			ret["year_of_birth"] = inv.Input.YearOfBirth
		}
		if inv.IsParameterSelected("How") {
			ret["how"] = inv.Input.How
		}
		if inv.IsParameterSelected("Note") {
			ret["note"] = inv.Input.Note
		}
		if inv.IsParameterSelected("OsTemplate") {
			ret["os_template"] = inv.Input.OsTemplate
		}
		if inv.IsParameterSelected("Location") {
			ret["location"] = inv.Input.Location
		}
		if inv.IsParameterSelected("Currency") {
			ret["currency"] = inv.Input.Currency
		}
		if inv.IsParameterSelected("Language") {
			ret["language"] = inv.Input.Language
		}
	}

	return ret
}

func (inv *ActionUserRequestRegistrationCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
