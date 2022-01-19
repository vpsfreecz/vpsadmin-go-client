package client

import ()

// ActionAuthTokenIndex is a type for action Auth_token#Index
type ActionAuthTokenIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionAuthTokenIndex(client *Client) *ActionAuthTokenIndex {
	return &ActionAuthTokenIndex{
		Client: client,
	}
}

// ActionAuthTokenIndexMetaGlobalInput is a type for action global meta input parameters
type ActionAuthTokenIndexMetaGlobalInput struct {
	No       bool   `json:"no"`
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionAuthTokenIndexMetaGlobalInput) SetNo(value bool) *ActionAuthTokenIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionAuthTokenIndexMetaGlobalInput) SetCount(value bool) *ActionAuthTokenIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionAuthTokenIndexMetaGlobalInput) SetIncludes(value string) *ActionAuthTokenIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionAuthTokenIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionAuthTokenIndexMetaGlobalInput) SelectParameters(params ...string) *ActionAuthTokenIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionAuthTokenIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionAuthTokenIndexInput is a type for action input parameters
type ActionAuthTokenIndexInput struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
	User   int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionAuthTokenIndexInput) SetOffset(value int64) *ActionAuthTokenIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionAuthTokenIndexInput) SetLimit(value int64) *ActionAuthTokenIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionAuthTokenIndexInput) SetUser(value int64) *ActionAuthTokenIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionAuthTokenIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionAuthTokenIndexInput) SelectParameters(params ...string) *ActionAuthTokenIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionAuthTokenIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionAuthTokenIndexOutput is a type for action output parameters
type ActionAuthTokenIndexOutput struct {
	Id        int64                 `json:"id"`
	User      *ActionUserShowOutput `json:"user"`
	Token     string                `json:"token"`
	ValidTo   string                `json:"valid_to"`
	Label     string                `json:"label"`
	Lifetime  string                `json:"lifetime"`
	Interval  int64                 `json:"interval"`
	UseCount  int64                 `json:"use_count"`
	CreatedAt string                `json:"created_at"`
}

// Type for action response, including envelope
type ActionAuthTokenIndexResponse struct {
	Action *ActionAuthTokenIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		AuthTokens []*ActionAuthTokenIndexOutput `json:"auth_tokens"`
	}

	// Action output without the namespace
	Output []*ActionAuthTokenIndexOutput
}

// Prepare the action for invocation
func (action *ActionAuthTokenIndex) Prepare() *ActionAuthTokenIndexInvocation {
	return &ActionAuthTokenIndexInvocation{
		Action: action,
		Path:   "/v5.0/auth_tokens",
	}
}

// ActionAuthTokenIndexInvocation is used to configure action for invocation
type ActionAuthTokenIndexInvocation struct {
	// Pointer to the action
	Action *ActionAuthTokenIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionAuthTokenIndexInput
	// Global meta input parameters
	MetaInput *ActionAuthTokenIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionAuthTokenIndexInvocation) NewInput() *ActionAuthTokenIndexInput {
	inv.Input = &ActionAuthTokenIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionAuthTokenIndexInvocation) SetInput(input *ActionAuthTokenIndexInput) *ActionAuthTokenIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionAuthTokenIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionAuthTokenIndexInvocation) NewMetaInput() *ActionAuthTokenIndexMetaGlobalInput {
	inv.MetaInput = &ActionAuthTokenIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionAuthTokenIndexInvocation) SetMetaInput(input *ActionAuthTokenIndexMetaGlobalInput) *ActionAuthTokenIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionAuthTokenIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionAuthTokenIndexInvocation) Call() (*ActionAuthTokenIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionAuthTokenIndexInvocation) callAsQuery() (*ActionAuthTokenIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionAuthTokenIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.AuthTokens
	}
	return resp, err
}

func (inv *ActionAuthTokenIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["auth_token[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["auth_token[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["auth_token[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionAuthTokenIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
