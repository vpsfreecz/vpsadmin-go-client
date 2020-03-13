package client

import (
	"strings"
)

// ActionEnvironmentDatasetPlanShow is a type for action Environment.Dataset_plan#Show
type ActionEnvironmentDatasetPlanShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEnvironmentDatasetPlanShow(client *Client) *ActionEnvironmentDatasetPlanShow {
	return &ActionEnvironmentDatasetPlanShow{
		Client: client,
	}
}

// ActionEnvironmentDatasetPlanShowMetaGlobalInput is a type for action global meta input parameters
type ActionEnvironmentDatasetPlanShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEnvironmentDatasetPlanShowMetaGlobalInput) SetNo(value bool) *ActionEnvironmentDatasetPlanShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEnvironmentDatasetPlanShowMetaGlobalInput) SetIncludes(value string) *ActionEnvironmentDatasetPlanShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionEnvironmentDatasetPlanShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEnvironmentDatasetPlanShowMetaGlobalInput) SelectParameters(params ...string) *ActionEnvironmentDatasetPlanShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEnvironmentDatasetPlanShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionEnvironmentDatasetPlanShowOutput is a type for action output parameters
type ActionEnvironmentDatasetPlanShowOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	DatasetPlan *ActionDatasetPlanShowOutput `json:"dataset_plan"`
	UserAdd bool `json:"user_add"`
	UserRemove bool `json:"user_remove"`
}


// Type for action response, including envelope
type ActionEnvironmentDatasetPlanShowResponse struct {
	Action *ActionEnvironmentDatasetPlanShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DatasetPlan *ActionEnvironmentDatasetPlanShowOutput `json:"dataset_plan"`
	}

	// Action output without the namespace
	Output *ActionEnvironmentDatasetPlanShowOutput
}


// Prepare the action for invocation
func (action *ActionEnvironmentDatasetPlanShow) Prepare() *ActionEnvironmentDatasetPlanShowInvocation {
	return &ActionEnvironmentDatasetPlanShowInvocation{
		Action: action,
		Path: "/v6.0/environments/{environment_id}/dataset_plans/{dataset_plan_id}",
	}
}

// ActionEnvironmentDatasetPlanShowInvocation is used to configure action for invocation
type ActionEnvironmentDatasetPlanShowInvocation struct {
	// Pointer to the action
	Action *ActionEnvironmentDatasetPlanShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEnvironmentDatasetPlanShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEnvironmentDatasetPlanShowInvocation) SetPathParamInt(param string, value int64) *ActionEnvironmentDatasetPlanShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEnvironmentDatasetPlanShowInvocation) SetPathParamString(param string, value string) *ActionEnvironmentDatasetPlanShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEnvironmentDatasetPlanShowInvocation) NewMetaInput() *ActionEnvironmentDatasetPlanShowMetaGlobalInput {
	inv.MetaInput = &ActionEnvironmentDatasetPlanShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEnvironmentDatasetPlanShowInvocation) SetMetaInput(input *ActionEnvironmentDatasetPlanShowMetaGlobalInput) *ActionEnvironmentDatasetPlanShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEnvironmentDatasetPlanShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEnvironmentDatasetPlanShowInvocation) Call() (*ActionEnvironmentDatasetPlanShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionEnvironmentDatasetPlanShowInvocation) callAsQuery() (*ActionEnvironmentDatasetPlanShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionEnvironmentDatasetPlanShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DatasetPlan
	}
	return resp, err
}




func (inv *ActionEnvironmentDatasetPlanShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

