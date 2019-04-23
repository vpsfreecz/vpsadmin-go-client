package client

import (
	"strings"
)

// ActionUserRequestRegistrationShow is a type for action User_request.Registration#Show
type ActionUserRequestRegistrationShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestRegistrationShow(client *Client) *ActionUserRequestRegistrationShow {
	return &ActionUserRequestRegistrationShow{
		Client: client,
	}
}

// ActionUserRequestRegistrationShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestRegistrationShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestRegistrationShowMetaGlobalInput) SetNo(value bool) *ActionUserRequestRegistrationShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestRegistrationShowMetaGlobalInput) SetIncludes(value string) *ActionUserRequestRegistrationShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestRegistrationShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserRequestRegistrationShowOutput is a type for action output parameters
type ActionUserRequestRegistrationShowOutput struct {
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
type ActionUserRequestRegistrationShowResponse struct {
	Action *ActionUserRequestRegistrationShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Registration *ActionUserRequestRegistrationShowOutput `json:"registration"`
	}

	// Action output without the namespace
	Output *ActionUserRequestRegistrationShowOutput
}


// Prepare the action for invocation
func (action *ActionUserRequestRegistrationShow) Prepare() *ActionUserRequestRegistrationShowInvocation {
	return &ActionUserRequestRegistrationShowInvocation{
		Action: action,
		Path: "/v5.0/user_request/registrations/:registration_id",
	}
}

// ActionUserRequestRegistrationShowInvocation is used to configure action for invocation
type ActionUserRequestRegistrationShowInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestRegistrationShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserRequestRegistrationShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserRequestRegistrationShowInvocation) SetPathParamInt(param string, value int64) *ActionUserRequestRegistrationShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserRequestRegistrationShowInvocation) SetPathParamString(param string, value string) *ActionUserRequestRegistrationShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestRegistrationShowInvocation) SetMetaInput(input *ActionUserRequestRegistrationShowMetaGlobalInput) *ActionUserRequestRegistrationShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestRegistrationShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestRegistrationShowInvocation) Call() (*ActionUserRequestRegistrationShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserRequestRegistrationShowInvocation) callAsQuery() (*ActionUserRequestRegistrationShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserRequestRegistrationShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Registration
	}
	return resp, err
}




func (inv *ActionUserRequestRegistrationShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

