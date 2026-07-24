package client

import ()

// ActionVpsSecurityAdvisoryIndex is a type for action Vps_security_advisory#Index
type ActionVpsSecurityAdvisoryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsSecurityAdvisoryIndex(client *Client) *ActionVpsSecurityAdvisoryIndex {
	return &ActionVpsSecurityAdvisoryIndex{
		Client: client,
	}
}

// ActionVpsSecurityAdvisoryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsSecurityAdvisoryIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexMetaGlobalInput) SetCount(value bool) *ActionVpsSecurityAdvisoryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsSecurityAdvisoryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexMetaGlobalInput) SetNo(value bool) *ActionVpsSecurityAdvisoryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsSecurityAdvisoryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsSecurityAdvisoryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsSecurityAdvisoryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsSecurityAdvisoryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsSecurityAdvisoryIndexInput is a type for action input parameters
type ActionVpsSecurityAdvisoryIndexInput struct {
	Environment      int64 "json:\"environment\""
	FromId           int64 "json:\"from_id\""
	Limit            int64 "json:\"limit\""
	Location         int64 "json:\"location\""
	Node             int64 "json:\"node\""
	SecurityAdvisory int64 "json:\"security_advisory\""
	User             int64 "json:\"user\""
	Vps              int64 "json:\"vps\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexInput) SetEnvironment(value int64) *ActionVpsSecurityAdvisoryIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexInput) SetFromId(value int64) *ActionVpsSecurityAdvisoryIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexInput) SetLimit(value int64) *ActionVpsSecurityAdvisoryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexInput) SetLocation(value int64) *ActionVpsSecurityAdvisoryIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexInput) SetNode(value int64) *ActionVpsSecurityAdvisoryIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetSecurityAdvisory sets parameter SecurityAdvisory to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexInput) SetSecurityAdvisory(value int64) *ActionVpsSecurityAdvisoryIndexInput {
	in.SecurityAdvisory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SecurityAdvisory"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexInput) SetUser(value int64) *ActionVpsSecurityAdvisoryIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionVpsSecurityAdvisoryIndexInput) SetVps(value int64) *ActionVpsSecurityAdvisoryIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsSecurityAdvisoryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsSecurityAdvisoryIndexInput) SelectParameters(params ...string) *ActionVpsSecurityAdvisoryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsSecurityAdvisoryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsSecurityAdvisoryIndexInput) UnselectParameters(params ...string) *ActionVpsSecurityAdvisoryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsSecurityAdvisoryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsSecurityAdvisoryIndexOutput is a type for action output parameters
type ActionVpsSecurityAdvisoryIndexOutput struct {
	Environment      *ActionEnvironmentShowOutput      "json:\"environment\""
	Id               int64                             "json:\"id\""
	Location         *ActionLocationShowOutput         "json:\"location\""
	MitigatedSince   string                            "json:\"mitigated_since\""
	Node             *ActionNodeShowOutput             "json:\"node\""
	NodeState        string                            "json:\"node_state\""
	SecurityAdvisory *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	User             *ActionUserShowOutput             "json:\"user\""
	Vps              *ActionVpsShowOutput              "json:\"vps\""
	VulnerableUntil  string                            "json:\"vulnerable_until\""
}

// Type for action response, including envelope
type ActionVpsSecurityAdvisoryIndexResponse struct {
	Action *ActionVpsSecurityAdvisoryIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsSecurityAdvisories []*ActionVpsSecurityAdvisoryIndexOutput "json:\"vps_security_advisories\""
	}

	// Action output without the namespace
	Output []*ActionVpsSecurityAdvisoryIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsSecurityAdvisoryIndex) Prepare() *ActionVpsSecurityAdvisoryIndexInvocation {
	return &ActionVpsSecurityAdvisoryIndexInvocation{
		Action: action,
		Path:   "/v7.0/vps_security_advisories",
	}
}

// ActionVpsSecurityAdvisoryIndexInvocation is used to configure action for invocation
type ActionVpsSecurityAdvisoryIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsSecurityAdvisoryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsSecurityAdvisoryIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsSecurityAdvisoryIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsSecurityAdvisoryIndexInvocation) NewInput() *ActionVpsSecurityAdvisoryIndexInput {
	inv.Input = &ActionVpsSecurityAdvisoryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsSecurityAdvisoryIndexInvocation) SetInput(input *ActionVpsSecurityAdvisoryIndexInput) *ActionVpsSecurityAdvisoryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsSecurityAdvisoryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsSecurityAdvisoryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsSecurityAdvisoryIndexInvocation) NewMetaInput() *ActionVpsSecurityAdvisoryIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsSecurityAdvisoryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsSecurityAdvisoryIndexInvocation) SetMetaInput(input *ActionVpsSecurityAdvisoryIndexMetaGlobalInput) *ActionVpsSecurityAdvisoryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsSecurityAdvisoryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsSecurityAdvisoryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionVpsSecurityAdvisoryIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			if !inv.IsParameterNil("Environment") {
				if inv.Input.Environment < 0 {
					verr.Add("environment", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Location") {
			if !inv.IsParameterNil("Location") {
				if inv.Input.Location < 0 {
					verr.Add("location", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			if !inv.IsParameterNil("SecurityAdvisory") {
				if inv.Input.SecurityAdvisory < 0 {
					verr.Add("security_advisory", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Vps") {
			if !inv.IsParameterNil("Vps") {
				if inv.Input.Vps < 0 {
					verr.Add("vps", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsSecurityAdvisoryIndexInvocation) Call() (*ActionVpsSecurityAdvisoryIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionVpsSecurityAdvisoryIndexInvocation) callAsQuery() (*ActionVpsSecurityAdvisoryIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionVpsSecurityAdvisoryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsSecurityAdvisories
	}
	return resp, err
}

func (inv *ActionVpsSecurityAdvisoryIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["vps_security_advisory[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("FromId") {
			ret["vps_security_advisory[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["vps_security_advisory[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["vps_security_advisory[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Node") {
			ret["vps_security_advisory[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			ret["vps_security_advisory[security_advisory]"] = convertInt64ToString(inv.Input.SecurityAdvisory)
		}
		if inv.IsParameterSelected("User") {
			ret["vps_security_advisory[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["vps_security_advisory[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}

	return nil
}

func (inv *ActionVpsSecurityAdvisoryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
