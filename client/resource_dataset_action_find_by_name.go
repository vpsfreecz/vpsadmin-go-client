package client

import ()

// ActionDatasetFindByName is a type for action Dataset#Find_by_name
type ActionDatasetFindByName struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetFindByName(client *Client) *ActionDatasetFindByName {
	return &ActionDatasetFindByName{
		Client: client,
	}
}

// ActionDatasetFindByNameMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetFindByNameMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetFindByNameMetaGlobalInput) SetIncludes(value string) *ActionDatasetFindByNameMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetFindByNameMetaGlobalInput) SetNo(value bool) *ActionDatasetFindByNameMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetFindByNameMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetFindByNameMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetFindByNameMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetFindByNameMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetFindByNameInput is a type for action input parameters
type ActionDatasetFindByNameInput struct {
	Name string `json:"name"`
	User int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionDatasetFindByNameInput) SetName(value string) *ActionDatasetFindByNameInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionDatasetFindByNameInput) SetUser(value int64) *ActionDatasetFindByNameInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetFindByNameInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetFindByNameInput) SelectParameters(params ...string) *ActionDatasetFindByNameInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetFindByNameInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetFindByNameOutput is a type for action output parameters
type ActionDatasetFindByNameOutput struct {
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
type ActionDatasetFindByNameResponse struct {
	Action *ActionDatasetFindByName `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Dataset *ActionDatasetFindByNameOutput `json:"dataset"`
	}

	// Action output without the namespace
	Output *ActionDatasetFindByNameOutput
}

// Prepare the action for invocation
func (action *ActionDatasetFindByName) Prepare() *ActionDatasetFindByNameInvocation {
	return &ActionDatasetFindByNameInvocation{
		Action: action,
		Path:   "/v6.0/datasets/find_by_name",
	}
}

// ActionDatasetFindByNameInvocation is used to configure action for invocation
type ActionDatasetFindByNameInvocation struct {
	// Pointer to the action
	Action *ActionDatasetFindByName

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetFindByNameInput
	// Global meta input parameters
	MetaInput *ActionDatasetFindByNameMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetFindByNameInvocation) NewInput() *ActionDatasetFindByNameInput {
	inv.Input = &ActionDatasetFindByNameInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetFindByNameInvocation) SetInput(input *ActionDatasetFindByNameInput) *ActionDatasetFindByNameInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetFindByNameInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetFindByNameInvocation) NewMetaInput() *ActionDatasetFindByNameMetaGlobalInput {
	inv.MetaInput = &ActionDatasetFindByNameMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetFindByNameInvocation) SetMetaInput(input *ActionDatasetFindByNameMetaGlobalInput) *ActionDatasetFindByNameInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetFindByNameInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetFindByNameInvocation) Call() (*ActionDatasetFindByNameResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetFindByNameInvocation) callAsQuery() (*ActionDatasetFindByNameResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetFindByNameResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Dataset
	}
	return resp, err
}

func (inv *ActionDatasetFindByNameInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Name") {
			ret["dataset[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("User") {
			ret["dataset[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionDatasetFindByNameInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
