package client

import (
	"strings"
)

// ActionDatasetShow is a type for action Dataset#Show
type ActionDatasetShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetShow(client *Client) *ActionDatasetShow {
	return &ActionDatasetShow{
		Client: client,
	}
}

// ActionDatasetShowMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetShowMetaGlobalInput) SetNo(value bool) *ActionDatasetShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetShowMetaGlobalInput) SetIncludes(value string) *ActionDatasetShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetShowMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionDatasetShowOutput is a type for action output parameters
type ActionDatasetShowOutput struct {
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
type ActionDatasetShowResponse struct {
	Action *ActionDatasetShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Dataset *ActionDatasetShowOutput `json:"dataset"`
	}

	// Action output without the namespace
	Output *ActionDatasetShowOutput
}


// Prepare the action for invocation
func (action *ActionDatasetShow) Prepare() *ActionDatasetShowInvocation {
	return &ActionDatasetShowInvocation{
		Action: action,
		Path: "/v5.0/datasets/{dataset_id}",
	}
}

// ActionDatasetShowInvocation is used to configure action for invocation
type ActionDatasetShowInvocation struct {
	// Pointer to the action
	Action *ActionDatasetShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDatasetShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetShowInvocation) SetPathParamInt(param string, value int64) *ActionDatasetShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetShowInvocation) SetPathParamString(param string, value string) *ActionDatasetShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetShowInvocation) NewMetaInput() *ActionDatasetShowMetaGlobalInput {
	inv.MetaInput = &ActionDatasetShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetShowInvocation) SetMetaInput(input *ActionDatasetShowMetaGlobalInput) *ActionDatasetShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetShowInvocation) Call() (*ActionDatasetShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetShowInvocation) callAsQuery() (*ActionDatasetShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Dataset
	}
	return resp, err
}




func (inv *ActionDatasetShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

