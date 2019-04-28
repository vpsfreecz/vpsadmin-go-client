package client

import (
	"strings"
)

// ActionUserRequestRegistrationResolve is a type for action User_request.Registration#Resolve
type ActionUserRequestRegistrationResolve struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestRegistrationResolve(client *Client) *ActionUserRequestRegistrationResolve {
	return &ActionUserRequestRegistrationResolve{
		Client: client,
	}
}

// ActionUserRequestRegistrationResolveMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestRegistrationResolveMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveMetaGlobalInput) SetNo(value bool) *ActionUserRequestRegistrationResolveMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveMetaGlobalInput) SetIncludes(value string) *ActionUserRequestRegistrationResolveMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationResolveMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationResolveMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestRegistrationResolveMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationResolveMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationResolveInput is a type for action input parameters
type ActionUserRequestRegistrationResolveInput struct {
	Action string `json:"action"`
	Reason string `json:"reason"`
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
	Activate bool `json:"activate"`
	Node int64 `json:"node"`
	CreateVps bool `json:"create_vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetAction sets parameter Action to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetAction(value string) *ActionUserRequestRegistrationResolveInput {
	in.Action = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Action"] = nil
	return in
}
// SetReason sets parameter Reason to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetReason(value string) *ActionUserRequestRegistrationResolveInput {
	in.Reason = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reason"] = nil
	return in
}
// SetLogin sets parameter Login to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetLogin(value string) *ActionUserRequestRegistrationResolveInput {
	in.Login = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Login"] = nil
	return in
}
// SetFullName sets parameter FullName to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetFullName(value string) *ActionUserRequestRegistrationResolveInput {
	in.FullName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FullName"] = nil
	return in
}
// SetOrgName sets parameter OrgName to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetOrgName(value string) *ActionUserRequestRegistrationResolveInput {
	in.OrgName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OrgName"] = nil
	return in
}
// SetOrgId sets parameter OrgId to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetOrgId(value string) *ActionUserRequestRegistrationResolveInput {
	in.OrgId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OrgId"] = nil
	return in
}
// SetEmail sets parameter Email to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetEmail(value string) *ActionUserRequestRegistrationResolveInput {
	in.Email = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Email"] = nil
	return in
}
// SetAddress sets parameter Address to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetAddress(value string) *ActionUserRequestRegistrationResolveInput {
	in.Address = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Address"] = nil
	return in
}
// SetYearOfBirth sets parameter YearOfBirth to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetYearOfBirth(value int64) *ActionUserRequestRegistrationResolveInput {
	in.YearOfBirth = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["YearOfBirth"] = nil
	return in
}
// SetHow sets parameter How to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetHow(value string) *ActionUserRequestRegistrationResolveInput {
	in.How = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["How"] = nil
	return in
}
// SetNote sets parameter Note to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetNote(value string) *ActionUserRequestRegistrationResolveInput {
	in.Note = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Note"] = nil
	return in
}
// SetOsTemplate sets parameter OsTemplate to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetOsTemplate(value int64) *ActionUserRequestRegistrationResolveInput {
	in.OsTemplate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OsTemplate"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetLocation(value int64) *ActionUserRequestRegistrationResolveInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetCurrency sets parameter Currency to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetCurrency(value string) *ActionUserRequestRegistrationResolveInput {
	in.Currency = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Currency"] = nil
	return in
}
// SetLanguage sets parameter Language to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetLanguage(value int64) *ActionUserRequestRegistrationResolveInput {
	in.Language = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Language"] = nil
	return in
}
// SetActivate sets parameter Activate to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetActivate(value bool) *ActionUserRequestRegistrationResolveInput {
	in.Activate = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Activate"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetNode(value int64) *ActionUserRequestRegistrationResolveInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetCreateVps sets parameter CreateVps to value and selects it for sending
func (in *ActionUserRequestRegistrationResolveInput) SetCreateVps(value bool) *ActionUserRequestRegistrationResolveInput {
	in.CreateVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CreateVps"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationResolveInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationResolveInput) SelectParameters(params ...string) *ActionUserRequestRegistrationResolveInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationResolveInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationResolveRequest is a type for the entire action request
type ActionUserRequestRegistrationResolveRequest struct {
	Registration map[string]interface{} `json:"registration"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionUserRequestRegistrationResolveResponse struct {
	Action *ActionUserRequestRegistrationResolve `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionUserRequestRegistrationResolve) Prepare() *ActionUserRequestRegistrationResolveInvocation {
	return &ActionUserRequestRegistrationResolveInvocation{
		Action: action,
		Path: "/v5.0/user_request/registrations/:registration_id/resolve",
	}
}

// ActionUserRequestRegistrationResolveInvocation is used to configure action for invocation
type ActionUserRequestRegistrationResolveInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestRegistrationResolve

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserRequestRegistrationResolveInput
	// Global meta input parameters
	MetaInput *ActionUserRequestRegistrationResolveMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserRequestRegistrationResolveInvocation) SetPathParamInt(param string, value int64) *ActionUserRequestRegistrationResolveInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserRequestRegistrationResolveInvocation) SetPathParamString(param string, value string) *ActionUserRequestRegistrationResolveInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserRequestRegistrationResolveInvocation) NewInput() *ActionUserRequestRegistrationResolveInput {
	inv.Input = &ActionUserRequestRegistrationResolveInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserRequestRegistrationResolveInvocation) SetInput(input *ActionUserRequestRegistrationResolveInput) *ActionUserRequestRegistrationResolveInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserRequestRegistrationResolveInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestRegistrationResolveInvocation) NewMetaInput() *ActionUserRequestRegistrationResolveMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestRegistrationResolveMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestRegistrationResolveInvocation) SetMetaInput(input *ActionUserRequestRegistrationResolveMetaGlobalInput) *ActionUserRequestRegistrationResolveInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestRegistrationResolveInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestRegistrationResolveInvocation) Call() (*ActionUserRequestRegistrationResolveResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserRequestRegistrationResolveInvocation) callAsBody() (*ActionUserRequestRegistrationResolveResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserRequestRegistrationResolveResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionUserRequestRegistrationResolveInvocation) makeAllInputParams() *ActionUserRequestRegistrationResolveRequest {
	return &ActionUserRequestRegistrationResolveRequest{
		Registration: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserRequestRegistrationResolveInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Action") {
			ret["action"] = inv.Input.Action
		}
		if inv.IsParameterSelected("Reason") {
			ret["reason"] = inv.Input.Reason
		}
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
		if inv.IsParameterSelected("Activate") {
			ret["activate"] = inv.Input.Activate
		}
		if inv.IsParameterSelected("Node") {
			ret["node"] = inv.Input.Node
		}
		if inv.IsParameterSelected("CreateVps") {
			ret["create_vps"] = inv.Input.CreateVps
		}
	}

	return ret
}

func (inv *ActionUserRequestRegistrationResolveInvocation) makeMetaInputParams() map[string]interface{} {
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
