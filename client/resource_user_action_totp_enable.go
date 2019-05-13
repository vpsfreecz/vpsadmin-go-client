package client

import (
	"strings"
)

// ActionUserTotpEnable is a type for action User#Totp_enable
type ActionUserTotpEnable struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTotpEnable(client *Client) *ActionUserTotpEnable {
	return &ActionUserTotpEnable{
		Client: client,
	}
}

// ActionUserTotpEnableMetaGlobalInput is a type for action global meta input parameters
type ActionUserTotpEnableMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTotpEnableMetaGlobalInput) SetNo(value bool) *ActionUserTotpEnableMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpEnableMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpEnableMetaGlobalInput) SelectParameters(params ...string) *ActionUserTotpEnableMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpEnableMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserTotpEnableRequest is a type for the entire action request
type ActionUserTotpEnableRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionUserTotpEnableOutput is a type for action output parameters
type ActionUserTotpEnableOutput struct {
	Secret string `json:"secret"`
	ProvisioningUri string `json:"provisioning_uri"`
}


// Type for action response, including envelope
type ActionUserTotpEnableResponse struct {
	Action *ActionUserTotpEnable `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		User *ActionUserTotpEnableOutput `json:"user"`
	}

	// Action output without the namespace
	Output *ActionUserTotpEnableOutput
}


// Prepare the action for invocation
func (action *ActionUserTotpEnable) Prepare() *ActionUserTotpEnableInvocation {
	return &ActionUserTotpEnableInvocation{
		Action: action,
		Path: "/v5.0/users/totp_enable/{user_id}",
	}
}

// ActionUserTotpEnableInvocation is used to configure action for invocation
type ActionUserTotpEnableInvocation struct {
	// Pointer to the action
	Action *ActionUserTotpEnable

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserTotpEnableMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTotpEnableInvocation) SetPathParamInt(param string, value int64) *ActionUserTotpEnableInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTotpEnableInvocation) SetPathParamString(param string, value string) *ActionUserTotpEnableInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTotpEnableInvocation) NewMetaInput() *ActionUserTotpEnableMetaGlobalInput {
	inv.MetaInput = &ActionUserTotpEnableMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTotpEnableInvocation) SetMetaInput(input *ActionUserTotpEnableMetaGlobalInput) *ActionUserTotpEnableInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTotpEnableInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserTotpEnableInvocation) Call() (*ActionUserTotpEnableResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionUserTotpEnableInvocation) callAsBody() (*ActionUserTotpEnableResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionUserTotpEnableResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.User
	}
	return resp, err
}




func (inv *ActionUserTotpEnableInvocation) makeAllInputParams() *ActionUserTotpEnableRequest {
	return &ActionUserTotpEnableRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionUserTotpEnableInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
