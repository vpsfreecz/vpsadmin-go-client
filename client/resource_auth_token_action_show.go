package client

import (
	"strings"
)

// ActionAuthTokenShow is a type for action Auth_token#Show
type ActionAuthTokenShow struct {
	// Pointer to client
	Client *Client
}

func NewActionAuthTokenShow(client *Client) *ActionAuthTokenShow {
	return &ActionAuthTokenShow{
		Client: client,
	}
}

// ActionAuthTokenShowMetaGlobalInput is a type for action global meta input parameters
type ActionAuthTokenShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionAuthTokenShowMetaGlobalInput) SetNo(value bool) *ActionAuthTokenShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionAuthTokenShowMetaGlobalInput) SetIncludes(value string) *ActionAuthTokenShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionAuthTokenShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionAuthTokenShowMetaGlobalInput) SelectParameters(params ...string) *ActionAuthTokenShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionAuthTokenShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionAuthTokenShowOutput is a type for action output parameters
type ActionAuthTokenShowOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	Token string `json:"token"`
	ValidTo string `json:"valid_to"`
	Label string `json:"label"`
	Lifetime string `json:"lifetime"`
	Interval int64 `json:"interval"`
	UseCount int64 `json:"use_count"`
	CreatedAt string `json:"created_at"`
}


// Type for action response, including envelope
type ActionAuthTokenShowResponse struct {
	Action *ActionAuthTokenShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		AuthToken *ActionAuthTokenShowOutput `json:"auth_token"`
	}

	// Action output without the namespace
	Output *ActionAuthTokenShowOutput
}


// Prepare the action for invocation
func (action *ActionAuthTokenShow) Prepare() *ActionAuthTokenShowInvocation {
	return &ActionAuthTokenShowInvocation{
		Action: action,
		Path: "/v5.0/auth_tokens/:auth_token_id",
	}
}

// ActionAuthTokenShowInvocation is used to configure action for invocation
type ActionAuthTokenShowInvocation struct {
	// Pointer to the action
	Action *ActionAuthTokenShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionAuthTokenShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionAuthTokenShowInvocation) SetPathParamInt(param string, value int64) *ActionAuthTokenShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionAuthTokenShowInvocation) SetPathParamString(param string, value string) *ActionAuthTokenShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionAuthTokenShowInvocation) NewMetaInput() *ActionAuthTokenShowMetaGlobalInput {
	inv.MetaInput = &ActionAuthTokenShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionAuthTokenShowInvocation) SetMetaInput(input *ActionAuthTokenShowMetaGlobalInput) *ActionAuthTokenShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionAuthTokenShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionAuthTokenShowInvocation) Call() (*ActionAuthTokenShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionAuthTokenShowInvocation) callAsQuery() (*ActionAuthTokenShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionAuthTokenShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.AuthToken
	}
	return resp, err
}




func (inv *ActionAuthTokenShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

