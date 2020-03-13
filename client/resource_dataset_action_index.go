package client

import (
)

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
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	User int64 `json:"user"`
	Dataset int64 `json:"dataset"`
	UserNamespaceMap int64 `json:"user_namespace_map"`
	Role string `json:"role"`
	ToDepth int64 `json:"to_depth"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDatasetIndexInput) SetLimit(value int64) *ActionDatasetIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionDatasetIndexInput) SetUser(value int64) *ActionDatasetIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetDataset sets parameter Dataset to value and selects it for sending
func (in *ActionDatasetIndexInput) SetDataset(value int64) *ActionDatasetIndexInput {
	in.Dataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Dataset"] = nil
	return in
}
// SetUserNamespaceMap sets parameter UserNamespaceMap to value and selects it for sending
func (in *ActionDatasetIndexInput) SetUserNamespaceMap(value int64) *ActionDatasetIndexInput {
	in.UserNamespaceMap = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserNamespaceMap"] = nil
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

func (in *ActionDatasetIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionDatasetIndexOutput is a type for action output parameters
type ActionDatasetIndexOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Parent *ActionDatasetShowOutput `json:"parent"`
	User *ActionUserShowOutput `json:"user"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	UserNamespaceMap *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
	CurrentHistoryId int64 `json:"current_history_id"`
	Atime bool `json:"atime"`
	Compression bool `json:"compression"`
	Recordsize int64 `json:"recordsize"`
	Quota int64 `json:"quota"`
	Refquota int64 `json:"refquota"`
	Relatime bool `json:"relatime"`
	Sync string `json:"sync"`
	Sharenfs string `json:"sharenfs"`
	Used int64 `json:"used"`
	Referenced int64 `json:"referenced"`
	Avail int64 `json:"avail"`
	Export *ActionExportShowOutput `json:"export"`
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
		Path: "/v6.0/datasets",
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
		if inv.IsParameterSelected("Offset") {
			ret["dataset[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["dataset[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["dataset[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Dataset") {
			ret["dataset[dataset]"] = convertInt64ToString(inv.Input.Dataset)
		}
		if inv.IsParameterSelected("UserNamespaceMap") {
			ret["dataset[user_namespace_map]"] = convertInt64ToString(inv.Input.UserNamespaceMap)
		}
		if inv.IsParameterSelected("Role") {
			ret["dataset[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("ToDepth") {
			ret["dataset[to_depth]"] = convertInt64ToString(inv.Input.ToDepth)
		}
	}
}

func (inv *ActionDatasetIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

