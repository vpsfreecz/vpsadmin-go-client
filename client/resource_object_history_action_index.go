package client

import ()

// ActionObjectHistoryIndex is a type for action Object_history#Index
type ActionObjectHistoryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionObjectHistoryIndex(client *Client) *ActionObjectHistoryIndex {
	return &ActionObjectHistoryIndex{
		Client: client,
	}
}

// ActionObjectHistoryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionObjectHistoryIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionObjectHistoryIndexMetaGlobalInput) SetCount(value bool) *ActionObjectHistoryIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionObjectHistoryIndexMetaGlobalInput) SetIncludes(value string) *ActionObjectHistoryIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionObjectHistoryIndexMetaGlobalInput) SetNo(value bool) *ActionObjectHistoryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionObjectHistoryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionObjectHistoryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionObjectHistoryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionObjectHistoryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionObjectHistoryIndexInput is a type for action input parameters
type ActionObjectHistoryIndexInput struct {
	EventType   string `json:"event_type"`
	Limit       int64  `json:"limit"`
	Object      string `json:"object"`
	ObjectId    int64  `json:"object_id"`
	Offset      int64  `json:"offset"`
	User        int64  `json:"user"`
	UserSession int64  `json:"user_session"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEventType sets parameter EventType to value and selects it for sending
func (in *ActionObjectHistoryIndexInput) SetEventType(value string) *ActionObjectHistoryIndexInput {
	in.EventType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EventType"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionObjectHistoryIndexInput) SetLimit(value int64) *ActionObjectHistoryIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetObject sets parameter Object to value and selects it for sending
func (in *ActionObjectHistoryIndexInput) SetObject(value string) *ActionObjectHistoryIndexInput {
	in.Object = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Object"] = nil
	return in
}

// SetObjectId sets parameter ObjectId to value and selects it for sending
func (in *ActionObjectHistoryIndexInput) SetObjectId(value int64) *ActionObjectHistoryIndexInput {
	in.ObjectId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectId"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionObjectHistoryIndexInput) SetOffset(value int64) *ActionObjectHistoryIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionObjectHistoryIndexInput) SetUser(value int64) *ActionObjectHistoryIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionObjectHistoryIndexInput) SetUserNil(set bool) *ActionObjectHistoryIndexInput {
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

// SetUserSession sets parameter UserSession to value and selects it for sending
func (in *ActionObjectHistoryIndexInput) SetUserSession(value int64) *ActionObjectHistoryIndexInput {
	in.UserSession = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserSessionNil(false)
	in._selectedParameters["UserSession"] = nil
	return in
}

// SetUserSessionNil sets parameter UserSession to nil and selects it for sending
func (in *ActionObjectHistoryIndexInput) SetUserSessionNil(set bool) *ActionObjectHistoryIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["UserSession"] = nil
		in.SelectParameters("UserSession")
	} else {
		delete(in._nilParameters, "UserSession")
	}
	return in
}

// SelectParameters sets parameters from ActionObjectHistoryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionObjectHistoryIndexInput) SelectParameters(params ...string) *ActionObjectHistoryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionObjectHistoryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionObjectHistoryIndexInput) UnselectParameters(params ...string) *ActionObjectHistoryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionObjectHistoryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionObjectHistoryIndexOutput is a type for action output parameters
type ActionObjectHistoryIndexOutput struct {
	CreatedAt   string                       `json:"created_at"`
	EventType   string                       `json:"event_type"`
	Id          int64                        `json:"id"`
	Object      string                       `json:"object"`
	ObjectId    int64                        `json:"object_id"`
	User        *ActionUserShowOutput        `json:"user"`
	UserSession *ActionUserSessionShowOutput `json:"user_session"`
}

// Type for action response, including envelope
type ActionObjectHistoryIndexResponse struct {
	Action *ActionObjectHistoryIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ObjectHistories []*ActionObjectHistoryIndexOutput `json:"object_histories"`
	}

	// Action output without the namespace
	Output []*ActionObjectHistoryIndexOutput
}

// Prepare the action for invocation
func (action *ActionObjectHistoryIndex) Prepare() *ActionObjectHistoryIndexInvocation {
	return &ActionObjectHistoryIndexInvocation{
		Action: action,
		Path:   "/v6.0/object_histories",
	}
}

// ActionObjectHistoryIndexInvocation is used to configure action for invocation
type ActionObjectHistoryIndexInvocation struct {
	// Pointer to the action
	Action *ActionObjectHistoryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionObjectHistoryIndexInput
	// Global meta input parameters
	MetaInput *ActionObjectHistoryIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionObjectHistoryIndexInvocation) NewInput() *ActionObjectHistoryIndexInput {
	inv.Input = &ActionObjectHistoryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionObjectHistoryIndexInvocation) SetInput(input *ActionObjectHistoryIndexInput) *ActionObjectHistoryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionObjectHistoryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionObjectHistoryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionObjectHistoryIndexInvocation) NewMetaInput() *ActionObjectHistoryIndexMetaGlobalInput {
	inv.MetaInput = &ActionObjectHistoryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionObjectHistoryIndexInvocation) SetMetaInput(input *ActionObjectHistoryIndexMetaGlobalInput) *ActionObjectHistoryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionObjectHistoryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionObjectHistoryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionObjectHistoryIndexInvocation) Call() (*ActionObjectHistoryIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionObjectHistoryIndexInvocation) callAsQuery() (*ActionObjectHistoryIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionObjectHistoryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ObjectHistories
	}
	return resp, err
}

func (inv *ActionObjectHistoryIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("EventType") {
			ret["object_history[event_type]"] = inv.Input.EventType
		}
		if inv.IsParameterSelected("Limit") {
			ret["object_history[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Object") {
			ret["object_history[object]"] = inv.Input.Object
		}
		if inv.IsParameterSelected("ObjectId") {
			ret["object_history[object_id]"] = convertInt64ToString(inv.Input.ObjectId)
		}
		if inv.IsParameterSelected("Offset") {
			ret["object_history[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("User") {
			ret["object_history[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("UserSession") {
			ret["object_history[user_session]"] = convertInt64ToString(inv.Input.UserSession)
		}
	}
}

func (inv *ActionObjectHistoryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
