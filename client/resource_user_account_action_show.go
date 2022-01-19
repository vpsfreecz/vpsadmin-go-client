package client

import (
	"strings"
)

// ActionUserAccountShow is a type for action User_account#Show
type ActionUserAccountShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserAccountShow(client *Client) *ActionUserAccountShow {
	return &ActionUserAccountShow{
		Client: client,
	}
}

// ActionUserAccountShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserAccountShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserAccountShowMetaGlobalInput) SetIncludes(value string) *ActionUserAccountShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserAccountShowMetaGlobalInput) SetNo(value bool) *ActionUserAccountShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserAccountShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserAccountShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserAccountShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserAccountShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserAccountShowOutput is a type for action output parameters
type ActionUserAccountShowOutput struct {
	Id int64 `json:"id"`
	MonthlyPayment int64 `json:"monthly_payment"`
	PaidUntil string `json:"paid_until"`
}


// Type for action response, including envelope
type ActionUserAccountShowResponse struct {
	Action *ActionUserAccountShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserAccount *ActionUserAccountShowOutput `json:"user_account"`
	}

	// Action output without the namespace
	Output *ActionUserAccountShowOutput
}


// Prepare the action for invocation
func (action *ActionUserAccountShow) Prepare() *ActionUserAccountShowInvocation {
	return &ActionUserAccountShowInvocation{
		Action: action,
		Path: "/v6.0/user_accounts/{user_account_id}",
	}
}

// ActionUserAccountShowInvocation is used to configure action for invocation
type ActionUserAccountShowInvocation struct {
	// Pointer to the action
	Action *ActionUserAccountShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserAccountShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserAccountShowInvocation) SetPathParamInt(param string, value int64) *ActionUserAccountShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserAccountShowInvocation) SetPathParamString(param string, value string) *ActionUserAccountShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserAccountShowInvocation) NewMetaInput() *ActionUserAccountShowMetaGlobalInput {
	inv.MetaInput = &ActionUserAccountShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserAccountShowInvocation) SetMetaInput(input *ActionUserAccountShowMetaGlobalInput) *ActionUserAccountShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserAccountShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserAccountShowInvocation) Call() (*ActionUserAccountShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserAccountShowInvocation) callAsQuery() (*ActionUserAccountShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserAccountShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserAccount
	}
	return resp, err
}




func (inv *ActionUserAccountShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

