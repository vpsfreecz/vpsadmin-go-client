package client

import (
	"strings"
)

// ActionVpsConsoleTokenCreate is a type for action Vps.Console_token#Create
type ActionVpsConsoleTokenCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsConsoleTokenCreate(client *Client) *ActionVpsConsoleTokenCreate {
	return &ActionVpsConsoleTokenCreate{
		Client: client,
	}
}

// ActionVpsConsoleTokenCreateMetaGlobalInput is a type for action global meta input parameters
type ActionVpsConsoleTokenCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsConsoleTokenCreateMetaGlobalInput) SetNo(value bool) *ActionVpsConsoleTokenCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsConsoleTokenCreateMetaGlobalInput) SetIncludes(value string) *ActionVpsConsoleTokenCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConsoleTokenCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConsoleTokenCreateMetaGlobalInput) SelectParameters(params ...string) *ActionVpsConsoleTokenCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConsoleTokenCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionVpsConsoleTokenCreateRequest is a type for the entire action request
type ActionVpsConsoleTokenCreateRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionVpsConsoleTokenCreateOutput is a type for action output parameters
type ActionVpsConsoleTokenCreateOutput struct {
	Token string `json:"token"`
	Expiration string `json:"expiration"`
}


// Type for action response, including envelope
type ActionVpsConsoleTokenCreateResponse struct {
	Action *ActionVpsConsoleTokenCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ConsoleToken *ActionVpsConsoleTokenCreateOutput `json:"console_token"`
	}

	// Action output without the namespace
	Output *ActionVpsConsoleTokenCreateOutput
}


// Prepare the action for invocation
func (action *ActionVpsConsoleTokenCreate) Prepare() *ActionVpsConsoleTokenCreateInvocation {
	return &ActionVpsConsoleTokenCreateInvocation{
		Action: action,
		Path: "/v5.0/vpses/{vps_id}/console_token",
	}
}

// ActionVpsConsoleTokenCreateInvocation is used to configure action for invocation
type ActionVpsConsoleTokenCreateInvocation struct {
	// Pointer to the action
	Action *ActionVpsConsoleTokenCreate

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsConsoleTokenCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsConsoleTokenCreateInvocation) SetPathParamInt(param string, value int64) *ActionVpsConsoleTokenCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsConsoleTokenCreateInvocation) SetPathParamString(param string, value string) *ActionVpsConsoleTokenCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsConsoleTokenCreateInvocation) NewMetaInput() *ActionVpsConsoleTokenCreateMetaGlobalInput {
	inv.MetaInput = &ActionVpsConsoleTokenCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsConsoleTokenCreateInvocation) SetMetaInput(input *ActionVpsConsoleTokenCreateMetaGlobalInput) *ActionVpsConsoleTokenCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsConsoleTokenCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsConsoleTokenCreateInvocation) Call() (*ActionVpsConsoleTokenCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionVpsConsoleTokenCreateInvocation) callAsBody() (*ActionVpsConsoleTokenCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionVpsConsoleTokenCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ConsoleToken
	}
	return resp, err
}




func (inv *ActionVpsConsoleTokenCreateInvocation) makeAllInputParams() *ActionVpsConsoleTokenCreateRequest {
	return &ActionVpsConsoleTokenCreateRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionVpsConsoleTokenCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
