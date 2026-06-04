package client

import ()

// ActionSecurityAdvisoryUpdateIndex is a type for action Security_advisory_update#Index
type ActionSecurityAdvisoryUpdateIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryUpdateIndex(client *Client) *ActionSecurityAdvisoryUpdateIndex {
	return &ActionSecurityAdvisoryUpdateIndex{
		Client: client,
	}
}

// ActionSecurityAdvisoryUpdateIndexMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryUpdateIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput) SetCount(value bool) *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateIndexInput is a type for action input parameters
type ActionSecurityAdvisoryUpdateIndexInput struct {
	FromId           int64  "json:\"from_id\""
	Limit            int64  "json:\"limit\""
	SecurityAdvisory int64  "json:\"security_advisory\""
	Since            string "json:\"since\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateIndexInput) SetFromId(value int64) *ActionSecurityAdvisoryUpdateIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateIndexInput) SetLimit(value int64) *ActionSecurityAdvisoryUpdateIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetSecurityAdvisory sets parameter SecurityAdvisory to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateIndexInput) SetSecurityAdvisory(value int64) *ActionSecurityAdvisoryUpdateIndexInput {
	in.SecurityAdvisory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SecurityAdvisory"] = nil
	return in
}

// SetSince sets parameter Since to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateIndexInput) SetSince(value string) *ActionSecurityAdvisoryUpdateIndexInput {
	in.Since = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Since"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateIndexInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryUpdateIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateIndexInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryUpdateIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateIndexOutput is a type for action output parameters
type ActionSecurityAdvisoryUpdateIndexOutput struct {
	CreatedAt        string                            "json:\"created_at\""
	CsMessage        string                            "json:\"cs_message\""
	CsSummary        string                            "json:\"cs_summary\""
	EnMessage        string                            "json:\"en_message\""
	EnSummary        string                            "json:\"en_summary\""
	Id               int64                             "json:\"id\""
	Name             string                            "json:\"name\""
	ReportedBy       *ActionUserShowOutput             "json:\"reported_by\""
	ReporterName     string                            "json:\"reporter_name\""
	SecurityAdvisory *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	State            string                            "json:\"state\""
	UpdatedAt        string                            "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryUpdateIndexResponse struct {
	Action *ActionSecurityAdvisoryUpdateIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisoryUpdates []*ActionSecurityAdvisoryUpdateIndexOutput "json:\"security_advisory_updates\""
	}

	// Action output without the namespace
	Output []*ActionSecurityAdvisoryUpdateIndexOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryUpdateIndex) Prepare() *ActionSecurityAdvisoryUpdateIndexInvocation {
	return &ActionSecurityAdvisoryUpdateIndexInvocation{
		Action: action,
		Path:   "/v7.0/security_advisory_updates",
	}
}

// ActionSecurityAdvisoryUpdateIndexInvocation is used to configure action for invocation
type ActionSecurityAdvisoryUpdateIndexInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryUpdateIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryUpdateIndexInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) NewInput() *ActionSecurityAdvisoryUpdateIndexInput {
	inv.Input = &ActionSecurityAdvisoryUpdateIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) SetInput(input *ActionSecurityAdvisoryUpdateIndexInput) *ActionSecurityAdvisoryUpdateIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) NewMetaInput() *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryUpdateIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) SetMetaInput(input *ActionSecurityAdvisoryUpdateIndexMetaGlobalInput) *ActionSecurityAdvisoryUpdateIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("SecurityAdvisory") {
			if !inv.IsParameterNil("SecurityAdvisory") {
				if inv.Input.SecurityAdvisory < 0 {
					verr.Add("security_advisory", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Since") {
			if !inv.IsParameterNil("Since") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.Since)
				if !ok {
					verr.Add("since", "not a valid datetime")
				} else {
					inv.Input.Since = normalized
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
func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) Call() (*ActionSecurityAdvisoryUpdateIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) callAsQuery() (*ActionSecurityAdvisoryUpdateIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSecurityAdvisoryUpdateIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisoryUpdates
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["security_advisory_update[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["security_advisory_update[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			ret["security_advisory_update[security_advisory]"] = convertInt64ToString(inv.Input.SecurityAdvisory)
		}
		if inv.IsParameterSelected("Since") {
			ret["security_advisory_update[since]"] = inv.Input.Since
		}
	}
}

func (inv *ActionSecurityAdvisoryUpdateIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
