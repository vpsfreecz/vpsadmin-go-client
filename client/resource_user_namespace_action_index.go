package client

import ()

// ActionUserNamespaceIndex is a type for action User_namespace#Index
type ActionUserNamespaceIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceIndex(client *Client) *ActionUserNamespaceIndex {
	return &ActionUserNamespaceIndex{
		Client: client,
	}
}

// ActionUserNamespaceIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserNamespaceIndexMetaGlobalInput) SetCount(value bool) *ActionUserNamespaceIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceIndexMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceIndexMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceIndexInput is a type for action input parameters
type ActionUserNamespaceIndexInput struct {
	BlockCount int64 `json:"block_count"`
	Limit      int64 `json:"limit"`
	Offset     int64 `json:"offset"`
	Size       int64 `json:"size"`
	User       int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetBlockCount sets parameter BlockCount to value and selects it for sending
func (in *ActionUserNamespaceIndexInput) SetBlockCount(value int64) *ActionUserNamespaceIndexInput {
	in.BlockCount = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["BlockCount"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserNamespaceIndexInput) SetLimit(value int64) *ActionUserNamespaceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserNamespaceIndexInput) SetOffset(value int64) *ActionUserNamespaceIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetSize sets parameter Size to value and selects it for sending
func (in *ActionUserNamespaceIndexInput) SetSize(value int64) *ActionUserNamespaceIndexInput {
	in.Size = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Size"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserNamespaceIndexInput) SetUser(value int64) *ActionUserNamespaceIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionUserNamespaceIndexInput) SetUserNil(set bool) *ActionUserNamespaceIndexInput {
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

// SelectParameters sets parameters from ActionUserNamespaceIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceIndexInput) SelectParameters(params ...string) *ActionUserNamespaceIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserNamespaceIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserNamespaceIndexInput) UnselectParameters(params ...string) *ActionUserNamespaceIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserNamespaceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceIndexOutput is a type for action output parameters
type ActionUserNamespaceIndexOutput struct {
	BlockCount int64                 `json:"block_count"`
	Id         int64                 `json:"id"`
	Offset     int64                 `json:"offset"`
	Size       int64                 `json:"size"`
	User       *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionUserNamespaceIndexResponse struct {
	Action *ActionUserNamespaceIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserNamespaces []*ActionUserNamespaceIndexOutput `json:"user_namespaces"`
	}

	// Action output without the namespace
	Output []*ActionUserNamespaceIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserNamespaceIndex) Prepare() *ActionUserNamespaceIndexInvocation {
	return &ActionUserNamespaceIndexInvocation{
		Action: action,
		Path:   "/v6.0/user_namespaces",
	}
}

// ActionUserNamespaceIndexInvocation is used to configure action for invocation
type ActionUserNamespaceIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNamespaceIndexInput
	// Global meta input parameters
	MetaInput *ActionUserNamespaceIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNamespaceIndexInvocation) NewInput() *ActionUserNamespaceIndexInput {
	inv.Input = &ActionUserNamespaceIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNamespaceIndexInvocation) SetInput(input *ActionUserNamespaceIndexInput) *ActionUserNamespaceIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNamespaceIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserNamespaceIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceIndexInvocation) NewMetaInput() *ActionUserNamespaceIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceIndexInvocation) SetMetaInput(input *ActionUserNamespaceIndexMetaGlobalInput) *ActionUserNamespaceIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNamespaceIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceIndexInvocation) Call() (*ActionUserNamespaceIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserNamespaceIndexInvocation) callAsQuery() (*ActionUserNamespaceIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNamespaceIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserNamespaces
	}
	return resp, err
}

func (inv *ActionUserNamespaceIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("BlockCount") {
			ret["user_namespace[block_count]"] = convertInt64ToString(inv.Input.BlockCount)
		}
		if inv.IsParameterSelected("Limit") {
			ret["user_namespace[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["user_namespace[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Size") {
			ret["user_namespace[size]"] = convertInt64ToString(inv.Input.Size)
		}
		if inv.IsParameterSelected("User") {
			ret["user_namespace[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionUserNamespaceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
