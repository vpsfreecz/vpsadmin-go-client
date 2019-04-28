package client

import (
	"strings"
)

// ActionEnvironmentConfigChainIndex is a type for action Environment.Config_chain#Index
type ActionEnvironmentConfigChainIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionEnvironmentConfigChainIndex(client *Client) *ActionEnvironmentConfigChainIndex {
	return &ActionEnvironmentConfigChainIndex{
		Client: client,
	}
}

// ActionEnvironmentConfigChainIndexMetaGlobalInput is a type for action global meta input parameters
type ActionEnvironmentConfigChainIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEnvironmentConfigChainIndexMetaGlobalInput) SetNo(value bool) *ActionEnvironmentConfigChainIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionEnvironmentConfigChainIndexMetaGlobalInput) SetCount(value bool) *ActionEnvironmentConfigChainIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEnvironmentConfigChainIndexMetaGlobalInput) SetIncludes(value string) *ActionEnvironmentConfigChainIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentConfigChainIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentConfigChainIndexMetaGlobalInput) SelectParameters(params ...string) *ActionEnvironmentConfigChainIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentConfigChainIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEnvironmentConfigChainIndexInput is a type for action input parameters
type ActionEnvironmentConfigChainIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionEnvironmentConfigChainIndexInput) SetOffset(value int64) *ActionEnvironmentConfigChainIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionEnvironmentConfigChainIndexInput) SetLimit(value int64) *ActionEnvironmentConfigChainIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentConfigChainIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentConfigChainIndexInput) SelectParameters(params ...string) *ActionEnvironmentConfigChainIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentConfigChainIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionEnvironmentConfigChainIndexOutput is a type for action output parameters
type ActionEnvironmentConfigChainIndexOutput struct {
	VpsConfig *ActionVpsConfigShowOutput `json:"vps_config"`
}


// Type for action response, including envelope
type ActionEnvironmentConfigChainIndexResponse struct {
	Action *ActionEnvironmentConfigChainIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ConfigChains []*ActionEnvironmentConfigChainIndexOutput `json:"config_chains"`
	}

	// Action output without the namespace
	Output []*ActionEnvironmentConfigChainIndexOutput
}


// Prepare the action for invocation
func (action *ActionEnvironmentConfigChainIndex) Prepare() *ActionEnvironmentConfigChainIndexInvocation {
	return &ActionEnvironmentConfigChainIndexInvocation{
		Action: action,
		Path: "/v5.0/environments/:environment_id/config_chains",
	}
}

// ActionEnvironmentConfigChainIndexInvocation is used to configure action for invocation
type ActionEnvironmentConfigChainIndexInvocation struct {
	// Pointer to the action
	Action *ActionEnvironmentConfigChainIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEnvironmentConfigChainIndexInput
	// Global meta input parameters
	MetaInput *ActionEnvironmentConfigChainIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEnvironmentConfigChainIndexInvocation) SetPathParamInt(param string, value int64) *ActionEnvironmentConfigChainIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEnvironmentConfigChainIndexInvocation) SetPathParamString(param string, value string) *ActionEnvironmentConfigChainIndexInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEnvironmentConfigChainIndexInvocation) NewInput() *ActionEnvironmentConfigChainIndexInput {
	inv.Input = &ActionEnvironmentConfigChainIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEnvironmentConfigChainIndexInvocation) SetInput(input *ActionEnvironmentConfigChainIndexInput) *ActionEnvironmentConfigChainIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEnvironmentConfigChainIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEnvironmentConfigChainIndexInvocation) NewMetaInput() *ActionEnvironmentConfigChainIndexMetaGlobalInput {
	inv.MetaInput = &ActionEnvironmentConfigChainIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEnvironmentConfigChainIndexInvocation) SetMetaInput(input *ActionEnvironmentConfigChainIndexMetaGlobalInput) *ActionEnvironmentConfigChainIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEnvironmentConfigChainIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEnvironmentConfigChainIndexInvocation) Call() (*ActionEnvironmentConfigChainIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionEnvironmentConfigChainIndexInvocation) callAsQuery() (*ActionEnvironmentConfigChainIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEnvironmentConfigChainIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ConfigChains
	}
	return resp, err
}



func (inv *ActionEnvironmentConfigChainIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["config_chain[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["config_chain[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionEnvironmentConfigChainIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

