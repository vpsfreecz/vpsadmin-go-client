package client

import ()

// ActionVpsConfigIndex is a type for action Vps_config#Index
type ActionVpsConfigIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsConfigIndex(client *Client) *ActionVpsConfigIndex {
	return &ActionVpsConfigIndex{
		Client: client,
	}
}

// ActionVpsConfigIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsConfigIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsConfigIndexMetaGlobalInput) SetCount(value bool) *ActionVpsConfigIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsConfigIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsConfigIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsConfigIndexMetaGlobalInput) SetNo(value bool) *ActionVpsConfigIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConfigIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConfigIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsConfigIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConfigIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConfigIndexInput is a type for action input parameters
type ActionVpsConfigIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsConfigIndexInput) SetLimit(value int64) *ActionVpsConfigIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionVpsConfigIndexInput) SetOffset(value int64) *ActionVpsConfigIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsConfigIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsConfigIndexInput) SelectParameters(params ...string) *ActionVpsConfigIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsConfigIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsConfigIndexOutput is a type for action output parameters
type ActionVpsConfigIndexOutput struct {
	Config string `json:"config"`
	Id     int64  `json:"id"`
	Label  string `json:"label"`
	Name   string `json:"name"`
}

// Type for action response, including envelope
type ActionVpsConfigIndexResponse struct {
	Action *ActionVpsConfigIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsConfigs []*ActionVpsConfigIndexOutput `json:"vps_configs"`
	}

	// Action output without the namespace
	Output []*ActionVpsConfigIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsConfigIndex) Prepare() *ActionVpsConfigIndexInvocation {
	return &ActionVpsConfigIndexInvocation{
		Action: action,
		Path:   "/v6.0/vps_configs",
	}
}

// ActionVpsConfigIndexInvocation is used to configure action for invocation
type ActionVpsConfigIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsConfigIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsConfigIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsConfigIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsConfigIndexInvocation) NewInput() *ActionVpsConfigIndexInput {
	inv.Input = &ActionVpsConfigIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsConfigIndexInvocation) SetInput(input *ActionVpsConfigIndexInput) *ActionVpsConfigIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsConfigIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsConfigIndexInvocation) NewMetaInput() *ActionVpsConfigIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsConfigIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsConfigIndexInvocation) SetMetaInput(input *ActionVpsConfigIndexMetaGlobalInput) *ActionVpsConfigIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsConfigIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsConfigIndexInvocation) Call() (*ActionVpsConfigIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsConfigIndexInvocation) callAsQuery() (*ActionVpsConfigIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsConfigIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsConfigs
	}
	return resp, err
}

func (inv *ActionVpsConfigIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["vps_config[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["vps_config[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionVpsConfigIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
