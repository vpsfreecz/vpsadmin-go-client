package client

import (
)

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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserCurrentMetaGlobalInput) SetIncludes(value string) *ActionUserCurrentMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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
		Path: "/v5.0/users/current",
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
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

