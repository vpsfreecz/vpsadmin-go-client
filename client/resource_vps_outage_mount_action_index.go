package client

import ()

// ActionVpsOutageMountIndex is a type for action Vps_outage_mount#Index
type ActionVpsOutageMountIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsOutageMountIndex(client *Client) *ActionVpsOutageMountIndex {
	return &ActionVpsOutageMountIndex{
		Client: client,
	}
}

// ActionVpsOutageMountIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsOutageMountIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsOutageMountIndexMetaGlobalInput) SetCount(value bool) *ActionVpsOutageMountIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsOutageMountIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsOutageMountIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsOutageMountIndexMetaGlobalInput) SetNo(value bool) *ActionVpsOutageMountIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageMountIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageMountIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsOutageMountIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageMountIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsOutageMountIndexInput is a type for action input parameters
type ActionVpsOutageMountIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Outage int64 `json:"outage"`
	User   int64 `json:"user"`
	Vps    int64 `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsOutageMountIndexInput) SetLimit(value int64) *ActionVpsOutageMountIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionVpsOutageMountIndexInput) SetOffset(value int64) *ActionVpsOutageMountIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetOutage sets parameter Outage to value and selects it for sending
func (in *ActionVpsOutageMountIndexInput) SetOutage(value int64) *ActionVpsOutageMountIndexInput {
	in.Outage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Outage"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsOutageMountIndexInput) SetUser(value int64) *ActionVpsOutageMountIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionVpsOutageMountIndexInput) SetVps(value int64) *ActionVpsOutageMountIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageMountIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageMountIndexInput) SelectParameters(params ...string) *ActionVpsOutageMountIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageMountIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsOutageMountIndexOutput is a type for action output parameters
type ActionVpsOutageMountIndexOutput struct {
	Id        int64                      `json:"id"`
	Mount     *ActionVpsMountShowOutput  `json:"mount"`
	VpsOutage *ActionVpsOutageShowOutput `json:"vps_outage"`
}

// Type for action response, including envelope
type ActionVpsOutageMountIndexResponse struct {
	Action *ActionVpsOutageMountIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsOutageMounts []*ActionVpsOutageMountIndexOutput `json:"vps_outage_mounts"`
	}

	// Action output without the namespace
	Output []*ActionVpsOutageMountIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsOutageMountIndex) Prepare() *ActionVpsOutageMountIndexInvocation {
	return &ActionVpsOutageMountIndexInvocation{
		Action: action,
		Path:   "/v6.0/vps_outage_mounts",
	}
}

// ActionVpsOutageMountIndexInvocation is used to configure action for invocation
type ActionVpsOutageMountIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsOutageMountIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsOutageMountIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsOutageMountIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsOutageMountIndexInvocation) NewInput() *ActionVpsOutageMountIndexInput {
	inv.Input = &ActionVpsOutageMountIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsOutageMountIndexInvocation) SetInput(input *ActionVpsOutageMountIndexInput) *ActionVpsOutageMountIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsOutageMountIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsOutageMountIndexInvocation) NewMetaInput() *ActionVpsOutageMountIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsOutageMountIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsOutageMountIndexInvocation) SetMetaInput(input *ActionVpsOutageMountIndexMetaGlobalInput) *ActionVpsOutageMountIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsOutageMountIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsOutageMountIndexInvocation) Call() (*ActionVpsOutageMountIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsOutageMountIndexInvocation) callAsQuery() (*ActionVpsOutageMountIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsOutageMountIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsOutageMounts
	}
	return resp, err
}

func (inv *ActionVpsOutageMountIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["vps_outage_mount[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["vps_outage_mount[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Outage") {
			ret["vps_outage_mount[outage]"] = convertInt64ToString(inv.Input.Outage)
		}
		if inv.IsParameterSelected("User") {
			ret["vps_outage_mount[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["vps_outage_mount[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionVpsOutageMountIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
