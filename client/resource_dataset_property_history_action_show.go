package client

import (
	"strings"
)

// ActionDatasetPropertyHistoryShow is a type for action Dataset.Property_history#Show
type ActionDatasetPropertyHistoryShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetPropertyHistoryShow(client *Client) *ActionDatasetPropertyHistoryShow {
	return &ActionDatasetPropertyHistoryShow{
		Client: client,
	}
}

// ActionDatasetPropertyHistoryShowMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetPropertyHistoryShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetPropertyHistoryShowMetaGlobalInput) SetNo(value bool) *ActionDatasetPropertyHistoryShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetPropertyHistoryShowMetaGlobalInput) SetIncludes(value string) *ActionDatasetPropertyHistoryShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPropertyHistoryShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPropertyHistoryShowMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetPropertyHistoryShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetPropertyHistoryShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionDatasetPropertyHistoryShowOutput is a type for action output parameters
type ActionDatasetPropertyHistoryShowOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Value int64 `json:"value"`
	CreatedAt string `json:"created_at"`
}


// Type for action response, including envelope
type ActionDatasetPropertyHistoryShowResponse struct {
	Action *ActionDatasetPropertyHistoryShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PropertyHistory *ActionDatasetPropertyHistoryShowOutput `json:"property_history"`
	}

	// Action output without the namespace
	Output *ActionDatasetPropertyHistoryShowOutput
}


// Prepare the action for invocation
func (action *ActionDatasetPropertyHistoryShow) Prepare() *ActionDatasetPropertyHistoryShowInvocation {
	return &ActionDatasetPropertyHistoryShowInvocation{
		Action: action,
		Path: "/v5.0/datasets/:dataset_id/property_history/:property_history_id",
	}
}

// ActionDatasetPropertyHistoryShowInvocation is used to configure action for invocation
type ActionDatasetPropertyHistoryShowInvocation struct {
	// Pointer to the action
	Action *ActionDatasetPropertyHistoryShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDatasetPropertyHistoryShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetPropertyHistoryShowInvocation) SetPathParamInt(param string, value int64) *ActionDatasetPropertyHistoryShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetPropertyHistoryShowInvocation) SetPathParamString(param string, value string) *ActionDatasetPropertyHistoryShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetPropertyHistoryShowInvocation) NewMetaInput() *ActionDatasetPropertyHistoryShowMetaGlobalInput {
	inv.MetaInput = &ActionDatasetPropertyHistoryShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetPropertyHistoryShowInvocation) SetMetaInput(input *ActionDatasetPropertyHistoryShowMetaGlobalInput) *ActionDatasetPropertyHistoryShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetPropertyHistoryShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetPropertyHistoryShowInvocation) Call() (*ActionDatasetPropertyHistoryShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetPropertyHistoryShowInvocation) callAsQuery() (*ActionDatasetPropertyHistoryShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetPropertyHistoryShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.PropertyHistory
	}
	return resp, err
}




func (inv *ActionDatasetPropertyHistoryShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

