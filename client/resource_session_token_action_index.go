package client

import ()

// ActionSessionTokenIndex is a type for action Session_token#Index
type ActionSessionTokenIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionSessionTokenIndex(client *Client) *ActionSessionTokenIndex {
	return &ActionSessionTokenIndex{
		Client: client,
	}
}

// ActionSessionTokenIndexMetaGlobalInput is a type for action global meta input parameters
type ActionSessionTokenIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionSessionTokenIndexMetaGlobalInput) SetCount(value bool) *ActionSessionTokenIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSessionTokenIndexMetaGlobalInput) SetIncludes(value string) *ActionSessionTokenIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSessionTokenIndexMetaGlobalInput) SetNo(value bool) *ActionSessionTokenIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSessionTokenIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenIndexMetaGlobalInput) SelectParameters(params ...string) *ActionSessionTokenIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSessionTokenIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSessionTokenIndexInput is a type for action input parameters
type ActionSessionTokenIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	User   int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionSessionTokenIndexInput) SetLimit(value int64) *ActionSessionTokenIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionSessionTokenIndexInput) SetOffset(value int64) *ActionSessionTokenIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionSessionTokenIndexInput) SetUser(value int64) *ActionSessionTokenIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionSessionTokenIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSessionTokenIndexInput) SelectParameters(params ...string) *ActionSessionTokenIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSessionTokenIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSessionTokenIndexOutput is a type for action output parameters
type ActionSessionTokenIndexOutput struct {
	CreatedAt string                `json:"created_at"`
	Id        int64                 `json:"id"`
	Interval  int64                 `json:"interval"`
	Label     string                `json:"label"`
	Lifetime  string                `json:"lifetime"`
	Token     string                `json:"token"`
	UseCount  int64                 `json:"use_count"`
	User      *ActionUserShowOutput `json:"user"`
	ValidTo   string                `json:"valid_to"`
}

// Type for action response, including envelope
type ActionSessionTokenIndexResponse struct {
	Action *ActionSessionTokenIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SessionTokens []*ActionSessionTokenIndexOutput `json:"session_tokens"`
	}

	// Action output without the namespace
	Output []*ActionSessionTokenIndexOutput
}

// Prepare the action for invocation
func (action *ActionSessionTokenIndex) Prepare() *ActionSessionTokenIndexInvocation {
	return &ActionSessionTokenIndexInvocation{
		Action: action,
		Path:   "/v6.0/session_tokens",
	}
}

// ActionSessionTokenIndexInvocation is used to configure action for invocation
type ActionSessionTokenIndexInvocation struct {
	// Pointer to the action
	Action *ActionSessionTokenIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSessionTokenIndexInput
	// Global meta input parameters
	MetaInput *ActionSessionTokenIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSessionTokenIndexInvocation) NewInput() *ActionSessionTokenIndexInput {
	inv.Input = &ActionSessionTokenIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSessionTokenIndexInvocation) SetInput(input *ActionSessionTokenIndexInput) *ActionSessionTokenIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSessionTokenIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSessionTokenIndexInvocation) NewMetaInput() *ActionSessionTokenIndexMetaGlobalInput {
	inv.MetaInput = &ActionSessionTokenIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSessionTokenIndexInvocation) SetMetaInput(input *ActionSessionTokenIndexMetaGlobalInput) *ActionSessionTokenIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSessionTokenIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSessionTokenIndexInvocation) Call() (*ActionSessionTokenIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionSessionTokenIndexInvocation) callAsQuery() (*ActionSessionTokenIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionSessionTokenIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SessionTokens
	}
	return resp, err
}

func (inv *ActionSessionTokenIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["session_token[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["session_token[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("User") {
			ret["session_token[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionSessionTokenIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
