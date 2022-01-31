package client

import ()

// ActionDatasetIndex is a type for action Dataset#Index
type ActionDatasetIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetIndex(client *Client) *ActionDatasetIndex {
	return &ActionDatasetIndex{
		Client: client,
	}
}

// ActionDatasetIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDatasetIndexMetaGlobalInput) SetCount(value bool) *ActionDatasetIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetIndexMetaGlobalInput) SetIncludes(value string) *ActionDatasetIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetIndexMetaGlobalInput) SetNo(value bool) *ActionDatasetIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetIndexInput is a type for action input parameters
type ActionDatasetIndexInput struct {
	Dataset          int64  `json:"dataset"`
	Limit            int64  `json:"limit"`
	Offset           int64  `json:"offset"`
	Role             string `json:"role"`
	ToDepth          int64  `json:"to_depth"`
	User             int64  `json:"user"`
	UserNamespaceMap int64  `json:"user_namespace_map"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDataset sets parameter Dataset to value and selects it for sending
func (in *ActionDatasetIndexInput) SetDataset(value int64) *ActionDatasetIndexInput {
	in.Dataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDatasetNil(false)
	in._selectedParameters["Dataset"] = nil
	return in
}

// SetDatasetNil sets parameter Dataset to nil and selects it for sending
func (in *ActionDatasetIndexInput) SetDatasetNil(set bool) *ActionDatasetIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Dataset"] = nil
		in.SelectParameters("Dataset")
	} else {
		delete(in._nilParameters, "Dataset")
	}
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDatasetIndexInput) SetLimit(value int64) *ActionDatasetIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionDatasetIndexInput) SetOffset(value int64) *ActionDatasetIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetRole sets parameter Role to value and selects it for sending
func (in *ActionDatasetIndexInput) SetRole(value string) *ActionDatasetIndexInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}

// SetToDepth sets parameter ToDepth to value and selects it for sending
func (in *ActionDatasetIndexInput) SetToDepth(value int64) *ActionDatasetIndexInput {
	in.ToDepth = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ToDepth"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionDatasetIndexInput) SetUser(value int64) *ActionDatasetIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionDatasetIndexInput) SetUserNil(set bool) *ActionDatasetIndexInput {
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

// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionDatasetIndexInput) SetUserNamespaceMap(value int64) *ActionDatasetIndexInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNamespaceMapNil(false)
	in._selectedParameters["UserNamespaceMap"] = nil
	return in
}

// SetUserNamespaceMapNil sets parameter UserNamespaceMap to nil and selects it for sending
func (in *ActionDatasetIndexInput) SetUserNamespaceMapNil(set bool) *ActionDatasetIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["UserNamespaceMap"] = nil
		in.SelectParameters("UserNamespaceMap")
	} else {
		delete(in._nilParameters, "UserNamespaceMap")
	}
	return in
}

// SelectParameters sets parameters from ActionDatasetIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetIndexInput) SelectParameters(params ...string) *ActionDatasetIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetIndexInput) UnselectParameters(params ...string) *ActionDatasetIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetIndexOutput is a type for action output parameters
type ActionDatasetIndexOutput struct {
	Atime            bool                              `json:"atime"`
	Avail            int64                             `json:"avail"`
	Compression      bool                              `json:"compression"`
	CurrentHistoryId int64                             `json:"current_history_id"`
	Environment      *ActionEnvironmentShowOutput      `json:"environment"`
	Export           *ActionExportShowOutput           `json:"export"`
	Id               int64                             `json:"id"`
	Name             string                            `json:"name"`
	Parent           *ActionDatasetShowOutput          `json:"parent"`
	Quota            int64                             `json:"quota"`
	Recordsize       int64                             `json:"recordsize"`
	Referenced       int64                             `json:"referenced"`
	Refquota         int64                             `json:"refquota"`
	Relatime         bool                              `json:"relatime"`
	Sharenfs         string                            `json:"sharenfs"`
	Sync             string                            `json:"sync"`
	Used             int64                             `json:"used"`
	User             *ActionUserShowOutput             `json:"user"`
	UserNamespaceMap *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
}

// Type for action response, including envelope
type ActionDatasetIndexResponse struct {
	Action *ActionDatasetIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Datasets []*ActionDatasetIndexOutput `json:"datasets"`
	}

	// Action output without the namespace
	Output []*ActionDatasetIndexOutput
}

// Prepare the action for invocation
func (action *ActionDatasetIndex) Prepare() *ActionDatasetIndexInvocation {
	return &ActionDatasetIndexInvocation{
		Action: action,
		Path:   "/v6.0/datasets",
	}
}

// ActionDatasetIndexInvocation is used to configure action for invocation
type ActionDatasetIndexInvocation struct {
	// Pointer to the action
	Action *ActionDatasetIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetIndexInput
	// Global meta input parameters
	MetaInput *ActionDatasetIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetIndexInvocation) NewInput() *ActionDatasetIndexInput {
	inv.Input = &ActionDatasetIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetIndexInvocation) SetInput(input *ActionDatasetIndexInput) *ActionDatasetIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetIndexInvocation) NewMetaInput() *ActionDatasetIndexMetaGlobalInput {
	inv.MetaInput = &ActionDatasetIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetIndexInvocation) SetMetaInput(input *ActionDatasetIndexMetaGlobalInput) *ActionDatasetIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetIndexInvocation) Call() (*ActionDatasetIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetIndexInvocation) callAsQuery() (*ActionDatasetIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Datasets
	}
	return resp, err
}

func (inv *ActionDatasetIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Dataset") {
			ret["dataset[dataset]"] = convertInt64ToString(inv.Input.Dataset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dataset[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["dataset[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Role") {
			ret["dataset[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("ToDepth") {
			ret["dataset[to_depth]"] = convertInt64ToString(inv.Input.ToDepth)
		}
		if inv.IsParameterSelected("User") {
			ret["dataset[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("UserNamespaceMap") {
			ret["dataset[user_namespace_map]"] = convertInt64ToString(inv.Input.UserNamespaceMap)
		}
	}
}

func (inv *ActionDatasetIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
