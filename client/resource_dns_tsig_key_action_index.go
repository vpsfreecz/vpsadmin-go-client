package client

import ()

// ActionDnsTsigKeyIndex is a type for action Dns_tsig_key#Index
type ActionDnsTsigKeyIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsTsigKeyIndex(client *Client) *ActionDnsTsigKeyIndex {
	return &ActionDnsTsigKeyIndex{
		Client: client,
	}
}

// ActionDnsTsigKeyIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDnsTsigKeyIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDnsTsigKeyIndexMetaGlobalInput) SetCount(value bool) *ActionDnsTsigKeyIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsTsigKeyIndexMetaGlobalInput) SetIncludes(value string) *ActionDnsTsigKeyIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsTsigKeyIndexMetaGlobalInput) SetNo(value bool) *ActionDnsTsigKeyIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsTsigKeyIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsTsigKeyIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDnsTsigKeyIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsTsigKeyIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsTsigKeyIndexInput is a type for action input parameters
type ActionDnsTsigKeyIndexInput struct {
	Algorithm string `json:"algorithm"`
	FromId    int64  `json:"from_id"`
	Limit     int64  `json:"limit"`
	User      int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAlgorithm sets parameter Algorithm to value and selects it for sending
func (in *ActionDnsTsigKeyIndexInput) SetAlgorithm(value string) *ActionDnsTsigKeyIndexInput {
	in.Algorithm = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Algorithm"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionDnsTsigKeyIndexInput) SetFromId(value int64) *ActionDnsTsigKeyIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDnsTsigKeyIndexInput) SetLimit(value int64) *ActionDnsTsigKeyIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionDnsTsigKeyIndexInput) SetUser(value int64) *ActionDnsTsigKeyIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionDnsTsigKeyIndexInput) SetUserNil(set bool) *ActionDnsTsigKeyIndexInput {
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

// SelectParameters sets parameters from ActionDnsTsigKeyIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsTsigKeyIndexInput) SelectParameters(params ...string) *ActionDnsTsigKeyIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsTsigKeyIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsTsigKeyIndexInput) UnselectParameters(params ...string) *ActionDnsTsigKeyIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsTsigKeyIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsTsigKeyIndexOutput is a type for action output parameters
type ActionDnsTsigKeyIndexOutput struct {
	Algorithm string                `json:"algorithm"`
	CreatedAt string                `json:"created_at"`
	Id        int64                 `json:"id"`
	Name      string                `json:"name"`
	Secret    string                `json:"secret"`
	UpdatedAt string                `json:"updated_at"`
	User      *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionDnsTsigKeyIndexResponse struct {
	Action *ActionDnsTsigKeyIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsTsigKeys []*ActionDnsTsigKeyIndexOutput `json:"dns_tsig_keys"`
	}

	// Action output without the namespace
	Output []*ActionDnsTsigKeyIndexOutput
}

// Prepare the action for invocation
func (action *ActionDnsTsigKeyIndex) Prepare() *ActionDnsTsigKeyIndexInvocation {
	return &ActionDnsTsigKeyIndexInvocation{
		Action: action,
		Path:   "/v7.0/dns_tsig_keys",
	}
}

// ActionDnsTsigKeyIndexInvocation is used to configure action for invocation
type ActionDnsTsigKeyIndexInvocation struct {
	// Pointer to the action
	Action *ActionDnsTsigKeyIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsTsigKeyIndexInput
	// Global meta input parameters
	MetaInput *ActionDnsTsigKeyIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsTsigKeyIndexInvocation) NewInput() *ActionDnsTsigKeyIndexInput {
	inv.Input = &ActionDnsTsigKeyIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsTsigKeyIndexInvocation) SetInput(input *ActionDnsTsigKeyIndexInput) *ActionDnsTsigKeyIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsTsigKeyIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsTsigKeyIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsTsigKeyIndexInvocation) NewMetaInput() *ActionDnsTsigKeyIndexMetaGlobalInput {
	inv.MetaInput = &ActionDnsTsigKeyIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsTsigKeyIndexInvocation) SetMetaInput(input *ActionDnsTsigKeyIndexMetaGlobalInput) *ActionDnsTsigKeyIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsTsigKeyIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsTsigKeyIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsTsigKeyIndexInvocation) Call() (*ActionDnsTsigKeyIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDnsTsigKeyIndexInvocation) callAsQuery() (*ActionDnsTsigKeyIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDnsTsigKeyIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsTsigKeys
	}
	return resp, err
}

func (inv *ActionDnsTsigKeyIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Algorithm") {
			ret["dns_tsig_key[algorithm]"] = inv.Input.Algorithm
		}
		if inv.IsParameterSelected("FromId") {
			ret["dns_tsig_key[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dns_tsig_key[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["dns_tsig_key[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionDnsTsigKeyIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
