package client

import (
	"strings"
)

// ActionEnvironmentShow is a type for action Environment#Show
type ActionEnvironmentShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEnvironmentShow(client *Client) *ActionEnvironmentShow {
	return &ActionEnvironmentShow{
		Client: client,
	}
}

// ActionEnvironmentShowMetaGlobalInput is a type for action global meta input parameters
type ActionEnvironmentShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEnvironmentShowMetaGlobalInput) SetNo(value bool) *ActionEnvironmentShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEnvironmentShowMetaGlobalInput) SetIncludes(value string) *ActionEnvironmentShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentShowMetaGlobalInput) SelectParameters(params ...string) *ActionEnvironmentShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionEnvironmentShowOutput is a type for action output parameters
type ActionEnvironmentShowOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Domain string `json:"domain"`
	CanCreateVps bool `json:"can_create_vps"`
	CanDestroyVps bool `json:"can_destroy_vps"`
	VpsLifetime int64 `json:"vps_lifetime"`
	MaxVpsCount int64 `json:"max_vps_count"`
	UserIpOwnership bool `json:"user_ip_ownership"`
	MaintenanceLock string `json:"maintenance_lock"`
	MaintenanceLockReason string `json:"maintenance_lock_reason"`
}


// Type for action response, including envelope
type ActionEnvironmentShowResponse struct {
	Action *ActionEnvironmentShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Environment *ActionEnvironmentShowOutput `json:"environment"`
	}

	// Action output without the namespace
	Output *ActionEnvironmentShowOutput
}


// Prepare the action for invocation
func (action *ActionEnvironmentShow) Prepare() *ActionEnvironmentShowInvocation {
	return &ActionEnvironmentShowInvocation{
		Action: action,
		Path: "/v5.0/environments/{environment_id}",
	}
}

// ActionEnvironmentShowInvocation is used to configure action for invocation
type ActionEnvironmentShowInvocation struct {
	// Pointer to the action
	Action *ActionEnvironmentShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEnvironmentShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEnvironmentShowInvocation) SetPathParamInt(param string, value int64) *ActionEnvironmentShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEnvironmentShowInvocation) SetPathParamString(param string, value string) *ActionEnvironmentShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEnvironmentShowInvocation) NewMetaInput() *ActionEnvironmentShowMetaGlobalInput {
	inv.MetaInput = &ActionEnvironmentShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEnvironmentShowInvocation) SetMetaInput(input *ActionEnvironmentShowMetaGlobalInput) *ActionEnvironmentShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEnvironmentShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEnvironmentShowInvocation) Call() (*ActionEnvironmentShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionEnvironmentShowInvocation) callAsQuery() (*ActionEnvironmentShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEnvironmentShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Environment
	}
	return resp, err
}




func (inv *ActionEnvironmentShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

