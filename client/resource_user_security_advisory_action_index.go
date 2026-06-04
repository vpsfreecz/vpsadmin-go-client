package client

import ()

// ActionUserSecurityAdvisoryIndex is a type for action User_security_advisory#Index
type ActionUserSecurityAdvisoryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserSecurityAdvisoryIndex(client *Client) *ActionUserSecurityAdvisoryIndex {
	return &ActionUserSecurityAdvisoryIndex{
		Client: client,
	}
}

// ActionUserSecurityAdvisoryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserSecurityAdvisoryIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserSecurityAdvisoryIndexMetaGlobalInput) SetCount(value bool) *ActionUserSecurityAdvisoryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserSecurityAdvisoryIndexMetaGlobalInput) SetIncludes(value string) *ActionUserSecurityAdvisoryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserSecurityAdvisoryIndexMetaGlobalInput) SetNo(value bool) *ActionUserSecurityAdvisoryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserSecurityAdvisoryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSecurityAdvisoryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserSecurityAdvisoryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserSecurityAdvisoryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserSecurityAdvisoryIndexInput is a type for action input parameters
type ActionUserSecurityAdvisoryIndexInput struct {
	FromId           int64 "json:\"from_id\""
	Limit            int64 "json:\"limit\""
	SecurityAdvisory int64 "json:\"security_advisory\""
	User             int64 "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserSecurityAdvisoryIndexInput) SetFromId(value int64) *ActionUserSecurityAdvisoryIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserSecurityAdvisoryIndexInput) SetLimit(value int64) *ActionUserSecurityAdvisoryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetSecurityAdvisory sets parameter SecurityAdvisory to value and selects it for sending
func (in *ActionUserSecurityAdvisoryIndexInput) SetSecurityAdvisory(value int64) *ActionUserSecurityAdvisoryIndexInput {
	in.SecurityAdvisory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SecurityAdvisory"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserSecurityAdvisoryIndexInput) SetUser(value int64) *ActionUserSecurityAdvisoryIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserSecurityAdvisoryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserSecurityAdvisoryIndexInput) SelectParameters(params ...string) *ActionUserSecurityAdvisoryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserSecurityAdvisoryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserSecurityAdvisoryIndexInput) UnselectParameters(params ...string) *ActionUserSecurityAdvisoryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserSecurityAdvisoryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserSecurityAdvisoryIndexOutput is a type for action output parameters
type ActionUserSecurityAdvisoryIndexOutput struct {
	Id               int64                             "json:\"id\""
	SecurityAdvisory *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	User             *ActionUserShowOutput             "json:\"user\""
	VpsCount         int64                             "json:\"vps_count\""
}

// Type for action response, including envelope
type ActionUserSecurityAdvisoryIndexResponse struct {
	Action *ActionUserSecurityAdvisoryIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserSecurityAdvisories []*ActionUserSecurityAdvisoryIndexOutput "json:\"user_security_advisories\""
	}

	// Action output without the namespace
	Output []*ActionUserSecurityAdvisoryIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserSecurityAdvisoryIndex) Prepare() *ActionUserSecurityAdvisoryIndexInvocation {
	return &ActionUserSecurityAdvisoryIndexInvocation{
		Action: action,
		Path:   "/v7.0/user_security_advisories",
	}
}

// ActionUserSecurityAdvisoryIndexInvocation is used to configure action for invocation
type ActionUserSecurityAdvisoryIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserSecurityAdvisoryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserSecurityAdvisoryIndexInput
	// Global meta input parameters
	MetaInput *ActionUserSecurityAdvisoryIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserSecurityAdvisoryIndexInvocation) NewInput() *ActionUserSecurityAdvisoryIndexInput {
	inv.Input = &ActionUserSecurityAdvisoryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserSecurityAdvisoryIndexInvocation) SetInput(input *ActionUserSecurityAdvisoryIndexInput) *ActionUserSecurityAdvisoryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserSecurityAdvisoryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserSecurityAdvisoryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserSecurityAdvisoryIndexInvocation) NewMetaInput() *ActionUserSecurityAdvisoryIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserSecurityAdvisoryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserSecurityAdvisoryIndexInvocation) SetMetaInput(input *ActionUserSecurityAdvisoryIndexMetaGlobalInput) *ActionUserSecurityAdvisoryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserSecurityAdvisoryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserSecurityAdvisoryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserSecurityAdvisoryIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
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
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserSecurityAdvisoryIndexInvocation) Call() (*ActionUserSecurityAdvisoryIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionUserSecurityAdvisoryIndexInvocation) callAsQuery() (*ActionUserSecurityAdvisoryIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserSecurityAdvisoryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserSecurityAdvisories
	}
	return resp, err
}

func (inv *ActionUserSecurityAdvisoryIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["user_security_advisory[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["user_security_advisory[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			ret["user_security_advisory[security_advisory]"] = convertInt64ToString(inv.Input.SecurityAdvisory)
		}
		if inv.IsParameterSelected("User") {
			ret["user_security_advisory[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionUserSecurityAdvisoryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
