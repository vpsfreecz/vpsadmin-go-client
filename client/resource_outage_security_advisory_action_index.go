package client

import ()

// ActionOutageSecurityAdvisoryIndex is a type for action Outage_security_advisory#Index
type ActionOutageSecurityAdvisoryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageSecurityAdvisoryIndex(client *Client) *ActionOutageSecurityAdvisoryIndex {
	return &ActionOutageSecurityAdvisoryIndex{
		Client: client,
	}
}

// ActionOutageSecurityAdvisoryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOutageSecurityAdvisoryIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryIndexMetaGlobalInput) SetCount(value bool) *ActionOutageSecurityAdvisoryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryIndexMetaGlobalInput) SetIncludes(value string) *ActionOutageSecurityAdvisoryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryIndexMetaGlobalInput) SetNo(value bool) *ActionOutageSecurityAdvisoryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageSecurityAdvisoryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageSecurityAdvisoryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOutageSecurityAdvisoryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageSecurityAdvisoryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageSecurityAdvisoryIndexInput is a type for action input parameters
type ActionOutageSecurityAdvisoryIndexInput struct {
	FromId           int64 "json:\"from_id\""
	Limit            int64 "json:\"limit\""
	Outage           int64 "json:\"outage\""
	SecurityAdvisory int64 "json:\"security_advisory\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryIndexInput) SetFromId(value int64) *ActionOutageSecurityAdvisoryIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryIndexInput) SetLimit(value int64) *ActionOutageSecurityAdvisoryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOutage sets parameter Outage to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryIndexInput) SetOutage(value int64) *ActionOutageSecurityAdvisoryIndexInput {
	in.Outage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Outage"] = nil
	return in
}

// SetSecurityAdvisory sets parameter SecurityAdvisory to value and selects it for sending
func (in *ActionOutageSecurityAdvisoryIndexInput) SetSecurityAdvisory(value int64) *ActionOutageSecurityAdvisoryIndexInput {
	in.SecurityAdvisory = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SecurityAdvisory"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageSecurityAdvisoryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageSecurityAdvisoryIndexInput) SelectParameters(params ...string) *ActionOutageSecurityAdvisoryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionOutageSecurityAdvisoryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionOutageSecurityAdvisoryIndexInput) UnselectParameters(params ...string) *ActionOutageSecurityAdvisoryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionOutageSecurityAdvisoryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageSecurityAdvisoryIndexOutput is a type for action output parameters
type ActionOutageSecurityAdvisoryIndexOutput struct {
	Id                 int64                             "json:\"id\""
	Outage             *ActionOutageShowOutput           "json:\"outage\""
	OutageId           int64                             "json:\"outage_id\""
	SecurityAdvisory   *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	SecurityAdvisoryId int64                             "json:\"security_advisory_id\""
}

// Type for action response, including envelope
type ActionOutageSecurityAdvisoryIndexResponse struct {
	Action *ActionOutageSecurityAdvisoryIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		OutageSecurityAdvisories []*ActionOutageSecurityAdvisoryIndexOutput "json:\"outage_security_advisories\""
	}

	// Action output without the namespace
	Output []*ActionOutageSecurityAdvisoryIndexOutput
}

// Prepare the action for invocation
func (action *ActionOutageSecurityAdvisoryIndex) Prepare() *ActionOutageSecurityAdvisoryIndexInvocation {
	return &ActionOutageSecurityAdvisoryIndexInvocation{
		Action: action,
		Path:   "/v7.0/outage_security_advisories",
	}
}

// ActionOutageSecurityAdvisoryIndexInvocation is used to configure action for invocation
type ActionOutageSecurityAdvisoryIndexInvocation struct {
	// Pointer to the action
	Action *ActionOutageSecurityAdvisoryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageSecurityAdvisoryIndexInput
	// Global meta input parameters
	MetaInput *ActionOutageSecurityAdvisoryIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageSecurityAdvisoryIndexInvocation) NewInput() *ActionOutageSecurityAdvisoryIndexInput {
	inv.Input = &ActionOutageSecurityAdvisoryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageSecurityAdvisoryIndexInvocation) SetInput(input *ActionOutageSecurityAdvisoryIndexInput) *ActionOutageSecurityAdvisoryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageSecurityAdvisoryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionOutageSecurityAdvisoryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageSecurityAdvisoryIndexInvocation) NewMetaInput() *ActionOutageSecurityAdvisoryIndexMetaGlobalInput {
	inv.MetaInput = &ActionOutageSecurityAdvisoryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageSecurityAdvisoryIndexInvocation) SetMetaInput(input *ActionOutageSecurityAdvisoryIndexMetaGlobalInput) *ActionOutageSecurityAdvisoryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageSecurityAdvisoryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageSecurityAdvisoryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionOutageSecurityAdvisoryIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Outage") {
			if !inv.IsParameterNil("Outage") {
				if inv.Input.Outage < 0 {
					verr.Add("outage", "not a valid resource id")
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
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageSecurityAdvisoryIndexInvocation) Call() (*ActionOutageSecurityAdvisoryIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionOutageSecurityAdvisoryIndexInvocation) callAsQuery() (*ActionOutageSecurityAdvisoryIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionOutageSecurityAdvisoryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.OutageSecurityAdvisories
	}
	return resp, err
}

func (inv *ActionOutageSecurityAdvisoryIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["outage_security_advisory[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["outage_security_advisory[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Outage") {
			ret["outage_security_advisory[outage]"] = convertInt64ToString(inv.Input.Outage)
		}
		if inv.IsParameterSelected("SecurityAdvisory") {
			ret["outage_security_advisory[security_advisory]"] = convertInt64ToString(inv.Input.SecurityAdvisory)
		}
	}

	return nil
}

func (inv *ActionOutageSecurityAdvisoryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
