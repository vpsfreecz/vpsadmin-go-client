package client

import (
	"strings"
)

// ActionUserEnvironmentConfigShow is a type for action User.Environment_config#Show
type ActionUserEnvironmentConfigShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserEnvironmentConfigShow(client *Client) *ActionUserEnvironmentConfigShow {
	return &ActionUserEnvironmentConfigShow{
		Client: client,
	}
}

// ActionUserEnvironmentConfigShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserEnvironmentConfigShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserEnvironmentConfigShowMetaGlobalInput) SetIncludes(value string) *ActionUserEnvironmentConfigShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserEnvironmentConfigShowMetaGlobalInput) SetNo(value bool) *ActionUserEnvironmentConfigShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserEnvironmentConfigShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserEnvironmentConfigShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserEnvironmentConfigShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserEnvironmentConfigShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserEnvironmentConfigShowOutput is a type for action output parameters
type ActionUserEnvironmentConfigShowOutput struct {
	CanCreateVps  bool                         `json:"can_create_vps"`
	CanDestroyVps bool                         `json:"can_destroy_vps"`
	Default       bool                         `json:"default"`
	Environment   *ActionEnvironmentShowOutput `json:"environment"`
	Id            int64                        `json:"id"`
	MaxVpsCount   int64                        `json:"max_vps_count"`
	VpsLifetime   int64                        `json:"vps_lifetime"`
}

// Type for action response, including envelope
type ActionUserEnvironmentConfigShowResponse struct {
	Action *ActionUserEnvironmentConfigShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EnvironmentConfig *ActionUserEnvironmentConfigShowOutput `json:"environment_config"`
	}

	// Action output without the namespace
	Output *ActionUserEnvironmentConfigShowOutput
}

// Prepare the action for invocation
func (action *ActionUserEnvironmentConfigShow) Prepare() *ActionUserEnvironmentConfigShowInvocation {
	return &ActionUserEnvironmentConfigShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/environment_configs/{environment_config_id}",
	}
}

// ActionUserEnvironmentConfigShowInvocation is used to configure action for invocation
type ActionUserEnvironmentConfigShowInvocation struct {
	// Pointer to the action
	Action *ActionUserEnvironmentConfigShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserEnvironmentConfigShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserEnvironmentConfigShowInvocation) SetPathParamInt(param string, value int64) *ActionUserEnvironmentConfigShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserEnvironmentConfigShowInvocation) SetPathParamString(param string, value string) *ActionUserEnvironmentConfigShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserEnvironmentConfigShowInvocation) NewMetaInput() *ActionUserEnvironmentConfigShowMetaGlobalInput {
	inv.MetaInput = &ActionUserEnvironmentConfigShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserEnvironmentConfigShowInvocation) SetMetaInput(input *ActionUserEnvironmentConfigShowMetaGlobalInput) *ActionUserEnvironmentConfigShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserEnvironmentConfigShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserEnvironmentConfigShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserEnvironmentConfigShowInvocation) Call() (*ActionUserEnvironmentConfigShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserEnvironmentConfigShowInvocation) callAsQuery() (*ActionUserEnvironmentConfigShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserEnvironmentConfigShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EnvironmentConfig
	}
	return resp, err
}

func (inv *ActionUserEnvironmentConfigShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
