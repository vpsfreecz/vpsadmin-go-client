package client

import (
	"strings"
)

// ActionEnvironmentConfigChainReplace is a type for action Environment.Config_chain#Replace
type ActionEnvironmentConfigChainReplace struct {
	// Pointer to client
	Client *Client
}

func NewActionEnvironmentConfigChainReplace(client *Client) *ActionEnvironmentConfigChainReplace {
	return &ActionEnvironmentConfigChainReplace{
		Client: client,
	}
}

// ActionEnvironmentConfigChainReplaceMetaGlobalInput is a type for action global meta input parameters
type ActionEnvironmentConfigChainReplaceMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEnvironmentConfigChainReplaceMetaGlobalInput) SetNo(value bool) *ActionEnvironmentConfigChainReplaceMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEnvironmentConfigChainReplaceMetaGlobalInput) SetIncludes(value string) *ActionEnvironmentConfigChainReplaceMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentConfigChainReplaceMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentConfigChainReplaceMetaGlobalInput) SelectParameters(params ...string) *ActionEnvironmentConfigChainReplaceMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentConfigChainReplaceMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentConfigChainReplaceInput is a type for action input parameters
type ActionEnvironmentConfigChainReplaceInput struct {
	VpsConfig int64 `json:"vps_config"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetVpsConfig sets parameter VpsConfig to value and selects it for sending
func (in *ActionEnvironmentConfigChainReplaceInput) SetVpsConfig(value int64) *ActionEnvironmentConfigChainReplaceInput {
	in.VpsConfig = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsConfig"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentConfigChainReplaceInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentConfigChainReplaceInput) SelectParameters(params ...string) *ActionEnvironmentConfigChainReplaceInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentConfigChainReplaceInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentConfigChainReplaceRequest is a type for the entire action request
type ActionEnvironmentConfigChainReplaceRequest struct {
	ConfigChains map[string]interface{} `json:"config_chains"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionEnvironmentConfigChainReplaceResponse struct {
	Action *ActionEnvironmentConfigChainReplace `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionEnvironmentConfigChainReplace) Prepare() *ActionEnvironmentConfigChainReplaceInvocation {
	return &ActionEnvironmentConfigChainReplaceInvocation{
		Action: action,
		Path: "/v6.0/environments/{environment_id}/config_chains/replace",
	}
}

// ActionEnvironmentConfigChainReplaceInvocation is used to configure action for invocation
type ActionEnvironmentConfigChainReplaceInvocation struct {
	// Pointer to the action
	Action *ActionEnvironmentConfigChainReplace

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEnvironmentConfigChainReplaceInput
	// Global meta input parameters
	MetaInput *ActionEnvironmentConfigChainReplaceMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEnvironmentConfigChainReplaceInvocation) SetPathParamInt(param string, value int64) *ActionEnvironmentConfigChainReplaceInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEnvironmentConfigChainReplaceInvocation) SetPathParamString(param string, value string) *ActionEnvironmentConfigChainReplaceInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEnvironmentConfigChainReplaceInvocation) NewInput() *ActionEnvironmentConfigChainReplaceInput {
	inv.Input = &ActionEnvironmentConfigChainReplaceInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEnvironmentConfigChainReplaceInvocation) SetInput(input *ActionEnvironmentConfigChainReplaceInput) *ActionEnvironmentConfigChainReplaceInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEnvironmentConfigChainReplaceInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEnvironmentConfigChainReplaceInvocation) NewMetaInput() *ActionEnvironmentConfigChainReplaceMetaGlobalInput {
	inv.MetaInput = &ActionEnvironmentConfigChainReplaceMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEnvironmentConfigChainReplaceInvocation) SetMetaInput(input *ActionEnvironmentConfigChainReplaceMetaGlobalInput) *ActionEnvironmentConfigChainReplaceInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEnvironmentConfigChainReplaceInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEnvironmentConfigChainReplaceInvocation) Call() (*ActionEnvironmentConfigChainReplaceResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionEnvironmentConfigChainReplaceInvocation) callAsBody() (*ActionEnvironmentConfigChainReplaceResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEnvironmentConfigChainReplaceResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionEnvironmentConfigChainReplaceInvocation) makeAllInputParams() *ActionEnvironmentConfigChainReplaceRequest {
	return &ActionEnvironmentConfigChainReplaceRequest{
		ConfigChains: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionEnvironmentConfigChainReplaceInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("VpsConfig") {
			ret["vps_config"] = inv.Input.VpsConfig
		}
	}

	return ret
}

func (inv *ActionEnvironmentConfigChainReplaceInvocation) makeMetaInputParams() map[string]interface{} {
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
