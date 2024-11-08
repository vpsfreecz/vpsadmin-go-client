package client

import ()

// ActionVpsOutageIndex is a type for action Vps_outage#Index
type ActionVpsOutageIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsOutageIndex(client *Client) *ActionVpsOutageIndex {
	return &ActionVpsOutageIndex{
		Client: client,
	}
}

// ActionVpsOutageIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsOutageIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsOutageIndexMetaGlobalInput) SetCount(value bool) *ActionVpsOutageIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsOutageIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsOutageIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsOutageIndexMetaGlobalInput) SetNo(value bool) *ActionVpsOutageIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsOutageIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsOutageIndexInput is a type for action input parameters
type ActionVpsOutageIndexInput struct {
	Direct      bool  `json:"direct"`
	Environment int64 `json:"environment"`
	FromId      int64 `json:"from_id"`
	Limit       int64 `json:"limit"`
	Location    int64 `json:"location"`
	Node        int64 `json:"node"`
	Outage      int64 `json:"outage"`
	User        int64 `json:"user"`
	Vps         int64 `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDirect sets parameter Direct to value and selects it for sending
func (in *ActionVpsOutageIndexInput) SetDirect(value bool) *ActionVpsOutageIndexInput {
	in.Direct = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Direct"] = nil
	return in
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionVpsOutageIndexInput) SetEnvironment(value int64) *ActionVpsOutageIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnvironmentNil(false)
	in._selectedParameters["Environment"] = nil
	return in
}

// SetEnvironmentNil sets parameter Environment to nil and selects it for sending
func (in *ActionVpsOutageIndexInput) SetEnvironmentNil(set bool) *ActionVpsOutageIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Environment"] = nil
		in.SelectParameters("Environment")
	} else {
		delete(in._nilParameters, "Environment")
	}
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionVpsOutageIndexInput) SetFromId(value int64) *ActionVpsOutageIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsOutageIndexInput) SetLimit(value int64) *ActionVpsOutageIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionVpsOutageIndexInput) SetLocation(value int64) *ActionVpsOutageIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionVpsOutageIndexInput) SetLocationNil(set bool) *ActionVpsOutageIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Location"] = nil
		in.SelectParameters("Location")
	} else {
		delete(in._nilParameters, "Location")
	}
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsOutageIndexInput) SetNode(value int64) *ActionVpsOutageIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionVpsOutageIndexInput) SetNodeNil(set bool) *ActionVpsOutageIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Node"] = nil
		in.SelectParameters("Node")
	} else {
		delete(in._nilParameters, "Node")
	}
	return in
}

// SetOutage sets parameter Outage to value and selects it for sending
func (in *ActionVpsOutageIndexInput) SetOutage(value int64) *ActionVpsOutageIndexInput {
	in.Outage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetOutageNil(false)
	in._selectedParameters["Outage"] = nil
	return in
}

// SetOutageNil sets parameter Outage to nil and selects it for sending
func (in *ActionVpsOutageIndexInput) SetOutageNil(set bool) *ActionVpsOutageIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Outage"] = nil
		in.SelectParameters("Outage")
	} else {
		delete(in._nilParameters, "Outage")
	}
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsOutageIndexInput) SetUser(value int64) *ActionVpsOutageIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionVpsOutageIndexInput) SetUserNil(set bool) *ActionVpsOutageIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionVpsOutageIndexInput) SetVps(value int64) *ActionVpsOutageIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionVpsOutageIndexInput) SetVpsNil(set bool) *ActionVpsOutageIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Vps"] = nil
		in.SelectParameters("Vps")
	} else {
		delete(in._nilParameters, "Vps")
	}
	return in
}

// SelectParameters sets parameters from ActionVpsOutageIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageIndexInput) SelectParameters(params ...string) *ActionVpsOutageIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsOutageIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsOutageIndexInput) UnselectParameters(params ...string) *ActionVpsOutageIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsOutageIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsOutageIndexOutput is a type for action output parameters
type ActionVpsOutageIndexOutput struct {
	Direct      bool                         `json:"direct"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	Id          int64                        `json:"id"`
	Location    *ActionLocationShowOutput    `json:"location"`
	Node        *ActionNodeShowOutput        `json:"node"`
	Outage      *ActionOutageShowOutput      `json:"outage"`
	User        *ActionUserShowOutput        `json:"user"`
	Vps         *ActionVpsShowOutput         `json:"vps"`
}

// Type for action response, including envelope
type ActionVpsOutageIndexResponse struct {
	Action *ActionVpsOutageIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsOutages []*ActionVpsOutageIndexOutput `json:"vps_outages"`
	}

	// Action output without the namespace
	Output []*ActionVpsOutageIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsOutageIndex) Prepare() *ActionVpsOutageIndexInvocation {
	return &ActionVpsOutageIndexInvocation{
		Action: action,
		Path:   "/v7.0/vps_outages",
	}
}

// ActionVpsOutageIndexInvocation is used to configure action for invocation
type ActionVpsOutageIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsOutageIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsOutageIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsOutageIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsOutageIndexInvocation) NewInput() *ActionVpsOutageIndexInput {
	inv.Input = &ActionVpsOutageIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsOutageIndexInvocation) SetInput(input *ActionVpsOutageIndexInput) *ActionVpsOutageIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsOutageIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsOutageIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsOutageIndexInvocation) NewMetaInput() *ActionVpsOutageIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsOutageIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsOutageIndexInvocation) SetMetaInput(input *ActionVpsOutageIndexMetaGlobalInput) *ActionVpsOutageIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsOutageIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsOutageIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsOutageIndexInvocation) Call() (*ActionVpsOutageIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsOutageIndexInvocation) callAsQuery() (*ActionVpsOutageIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsOutageIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsOutages
	}
	return resp, err
}

func (inv *ActionVpsOutageIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Direct") {
			ret["vps_outage[direct]"] = convertBoolToString(inv.Input.Direct)
		}
		if inv.IsParameterSelected("Environment") {
			ret["vps_outage[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("FromId") {
			ret["vps_outage[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["vps_outage[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["vps_outage[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Node") {
			ret["vps_outage[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Outage") {
			ret["vps_outage[outage]"] = convertInt64ToString(inv.Input.Outage)
		}
		if inv.IsParameterSelected("User") {
			ret["vps_outage[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["vps_outage[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionVpsOutageIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
