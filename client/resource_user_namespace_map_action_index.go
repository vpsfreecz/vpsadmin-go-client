package client

import (
)

// ActionUserNamespaceMapIndex is a type for action User_namespace_map#Index
type ActionUserNamespaceMapIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapIndex(client *Client) *ActionUserNamespaceMapIndex {
	return &ActionUserNamespaceMapIndex{
		Client: client,
	}
}

// ActionUserNamespaceMapIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapIndexMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserNamespaceMapIndexMetaGlobalInput) SetCount(value bool) *ActionUserNamespaceMapIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapIndexMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapIndexInput is a type for action input parameters
type ActionUserNamespaceMapIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	User int64 `json:"user"`
	UserNamespace int64 `json:"user_namespace"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserNamespaceMapIndexInput) SetOffset(value int64) *ActionUserNamespaceMapIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserNamespaceMapIndexInput) SetLimit(value int64) *ActionUserNamespaceMapIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserNamespaceMapIndexInput) SetUser(value int64) *ActionUserNamespaceMapIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetUserNamespace sets parameter UserNamespace to value and selects it for sending
func (in *ActionUserNamespaceMapIndexInput) SetUserNamespace(value int64) *ActionUserNamespaceMapIndexInput {
	in.UserNamespace = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserNamespace"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapIndexInput) SelectParameters(params ...string) *ActionUserNamespaceMapIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserNamespaceMapIndexOutput is a type for action output parameters
type ActionUserNamespaceMapIndexOutput struct {
	Id int64 `json:"id"`
	UserNamespace *ActionUserNamespaceShowOutput `json:"user_namespace"`
	Label string `json:"label"`
}


// Type for action response, including envelope
type ActionUserNamespaceMapIndexResponse struct {
	Action *ActionUserNamespaceMapIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserNamespaceMaps []*ActionUserNamespaceMapIndexOutput `json:"user_namespace_maps"`
	}

	// Action output without the namespace
	Output []*ActionUserNamespaceMapIndexOutput
}


// Prepare the action for invocation
func (action *ActionUserNamespaceMapIndex) Prepare() *ActionUserNamespaceMapIndexInvocation {
	return &ActionUserNamespaceMapIndexInvocation{
		Action: action,
		Path: "/v6.0/user_namespace_maps",
	}
}

// ActionUserNamespaceMapIndexInvocation is used to configure action for invocation
type ActionUserNamespaceMapIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNamespaceMapIndexInput
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNamespaceMapIndexInvocation) NewInput() *ActionUserNamespaceMapIndexInput {
	inv.Input = &ActionUserNamespaceMapIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNamespaceMapIndexInvocation) SetInput(input *ActionUserNamespaceMapIndexInput) *ActionUserNamespaceMapIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNamespaceMapIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapIndexInvocation) NewMetaInput() *ActionUserNamespaceMapIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapIndexInvocation) SetMetaInput(input *ActionUserNamespaceMapIndexMetaGlobalInput) *ActionUserNamespaceMapIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserNamespaceMapIndexInvocation) Call() (*ActionUserNamespaceMapIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserNamespaceMapIndexInvocation) callAsQuery() (*ActionUserNamespaceMapIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserNamespaceMapIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserNamespaceMaps
	}
	return resp, err
}



func (inv *ActionUserNamespaceMapIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["user_namespace_map[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["user_namespace_map[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["user_namespace_map[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("UserNamespace") {
			ret["user_namespace_map[user_namespace]"] = convertInt64ToString(inv.Input.UserNamespace)
		}
	}
}

func (inv *ActionUserNamespaceMapIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

