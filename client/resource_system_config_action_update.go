package client

import (
	"strings"
)

// ActionSystemConfigUpdate is a type for action System_config#Update
type ActionSystemConfigUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionSystemConfigUpdate(client *Client) *ActionSystemConfigUpdate {
	return &ActionSystemConfigUpdate{
		Client: client,
	}
}

// ActionSystemConfigUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionSystemConfigUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSystemConfigUpdateMetaGlobalInput) SetNo(value bool) *ActionSystemConfigUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSystemConfigUpdateMetaGlobalInput) SetIncludes(value string) *ActionSystemConfigUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionSystemConfigUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSystemConfigUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionSystemConfigUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSystemConfigUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionSystemConfigUpdateRequest is a type for the entire action request
type ActionSystemConfigUpdateRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// ActionSystemConfigUpdateOutput is a type for action output parameters
type ActionSystemConfigUpdateOutput struct {
	Category string `json:"category"`
	Name string `json:"name"`
	Type string `json:"type"`
	Label string `json:"label"`
	Description string `json:"description"`
	MinUserLevel int64 `json:"min_user_level"`
}


// Type for action response, including envelope
type ActionSystemConfigUpdateResponse struct {
	Action *ActionSystemConfigUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SystemConfig *ActionSystemConfigUpdateOutput `json:"system_config"`
	}

	// Action output without the namespace
	Output *ActionSystemConfigUpdateOutput
}


// Prepare the action for invocation
func (action *ActionSystemConfigUpdate) Prepare() *ActionSystemConfigUpdateInvocation {
	return &ActionSystemConfigUpdateInvocation{
		Action: action,
		Path: "/v5.0/system_configs/:category/:name",
	}
}

// ActionSystemConfigUpdateInvocation is used to configure action for invocation
type ActionSystemConfigUpdateInvocation struct {
	// Pointer to the action
	Action *ActionSystemConfigUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSystemConfigUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSystemConfigUpdateInvocation) SetPathParamInt(param string, value int64) *ActionSystemConfigUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSystemConfigUpdateInvocation) SetPathParamString(param string, value string) *ActionSystemConfigUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSystemConfigUpdateInvocation) NewMetaInput() *ActionSystemConfigUpdateMetaGlobalInput {
	inv.MetaInput = &ActionSystemConfigUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSystemConfigUpdateInvocation) SetMetaInput(input *ActionSystemConfigUpdateMetaGlobalInput) *ActionSystemConfigUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSystemConfigUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSystemConfigUpdateInvocation) Call() (*ActionSystemConfigUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionSystemConfigUpdateInvocation) callAsBody() (*ActionSystemConfigUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSystemConfigUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SystemConfig
	}
	return resp, err
}




func (inv *ActionSystemConfigUpdateInvocation) makeAllInputParams() *ActionSystemConfigUpdateRequest {
	return &ActionSystemConfigUpdateRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionSystemConfigUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
