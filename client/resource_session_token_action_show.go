package client

import (
	"strings"
)

// ActionSessionTokenShow is a type for action Session_token#Show
type ActionSessionTokenShow struct {
	// Pointer to client
	Client *Client
}

func NewActionSessionTokenShow(client *Client) *ActionSessionTokenShow {
	return &ActionSessionTokenShow{
		Client: client,
	}
}

// ActionSessionTokenShowMetaGlobalInput is a type for action global meta input parameters
type ActionSessionTokenShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSessionTokenShowMetaGlobalInput) SetNo(value bool) *ActionSessionTokenShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSessionTokenShowMetaGlobalInput) SetIncludes(value string) *ActionSessionTokenShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionSessionTokenShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenShowMetaGlobalInput) SelectParameters(params ...string) *ActionSessionTokenShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSessionTokenShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionSessionTokenShowOutput is a type for action output parameters
type ActionSessionTokenShowOutput struct {
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
type ActionSessionTokenShowResponse struct {
	Action *ActionSessionTokenShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SessionToken *ActionSessionTokenShowOutput `json:"session_token"`
	}

	// Action output without the namespace
	Output *ActionSessionTokenShowOutput
}


// Prepare the action for invocation
func (action *ActionSessionTokenShow) Prepare() *ActionSessionTokenShowInvocation {
	return &ActionSessionTokenShowInvocation{
		Action: action,
		Path: "/v6.0/session_tokens/{session_token_id}",
	}
}

// ActionSessionTokenShowInvocation is used to configure action for invocation
type ActionSessionTokenShowInvocation struct {
	// Pointer to the action
	Action *ActionSessionTokenShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSessionTokenShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSessionTokenShowInvocation) SetPathParamInt(param string, value int64) *ActionSessionTokenShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSessionTokenShowInvocation) SetPathParamString(param string, value string) *ActionSessionTokenShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSessionTokenShowInvocation) NewMetaInput() *ActionSessionTokenShowMetaGlobalInput {
	inv.MetaInput = &ActionSessionTokenShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSessionTokenShowInvocation) SetMetaInput(input *ActionSessionTokenShowMetaGlobalInput) *ActionSessionTokenShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSessionTokenShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSessionTokenShowInvocation) Call() (*ActionSessionTokenShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionSessionTokenShowInvocation) callAsQuery() (*ActionSessionTokenShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSessionTokenShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SessionToken
	}
	return resp, err
}




func (inv *ActionSessionTokenShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

