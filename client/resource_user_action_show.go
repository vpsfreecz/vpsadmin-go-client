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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserShowMetaGlobalInput) SetIncludes(value string) *ActionUserShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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
	Id int64 `json:"id"`
	Login string `json:"login"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
	Level int64 `json:"level"`
	Info string `json:"info"`
	MailerEnabled bool `json:"mailer_enabled"`
	PasswordReset bool `json:"password_reset"`
	Lockout bool `json:"lockout"`
	Language *ActionLanguageShowOutput `json:"language"`
	TotpEnabled bool `json:"totp_enabled"`
	LastActivityAt string `json:"last_activity_at"`
	CreatedAt string `json:"created_at"`
	ObjectState string `json:"object_state"`
	ExpirationDate string `json:"expiration_date"`
	MonthlyPayment int64 `json:"monthly_payment"`
	PaidUntil string `json:"paid_until"`
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
		Path: "/v5.0/users/{user_id}",
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
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

