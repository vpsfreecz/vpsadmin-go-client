package client

import (
	"strings"
)

// ActionDatasetPlanShow is a type for action Dataset_plan#Show
type ActionDatasetPlanShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetPlanShow(client *Client) *ActionDatasetPlanShow {
	return &ActionDatasetPlanShow{
		Client: client,
	}
}

// ActionDatasetPlanShowMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetPlanShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetPlanShowMetaGlobalInput) SetNo(value bool) *ActionDatasetPlanShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetPlanShowMetaGlobalInput) SetIncludes(value string) *ActionDatasetPlanShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPlanShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPlanShowMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetPlanShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetPlanShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionDatasetPlanShowOutput is a type for action output parameters
type ActionDatasetPlanShowOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Description string `json:"description"`
}


// Type for action response, including envelope
type ActionDatasetPlanShowResponse struct {
	Action *ActionDatasetPlanShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DatasetPlan *ActionDatasetPlanShowOutput `json:"dataset_plan"`
	}

	// Action output without the namespace
	Output *ActionDatasetPlanShowOutput
}


// Prepare the action for invocation
func (action *ActionDatasetPlanShow) Prepare() *ActionDatasetPlanShowInvocation {
	return &ActionDatasetPlanShowInvocation{
		Action: action,
		Path: "/v5.0/dataset_plans/{dataset_plan_id}",
	}
}

// ActionDatasetPlanShowInvocation is used to configure action for invocation
type ActionDatasetPlanShowInvocation struct {
	// Pointer to the action
	Action *ActionDatasetPlanShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDatasetPlanShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetPlanShowInvocation) SetPathParamInt(param string, value int64) *ActionDatasetPlanShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetPlanShowInvocation) SetPathParamString(param string, value string) *ActionDatasetPlanShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetPlanShowInvocation) NewMetaInput() *ActionDatasetPlanShowMetaGlobalInput {
	inv.MetaInput = &ActionDatasetPlanShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetPlanShowInvocation) SetMetaInput(input *ActionDatasetPlanShowMetaGlobalInput) *ActionDatasetPlanShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetPlanShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetPlanShowInvocation) Call() (*ActionDatasetPlanShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetPlanShowInvocation) callAsQuery() (*ActionDatasetPlanShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetPlanShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DatasetPlan
	}
	return resp, err
}




func (inv *ActionDatasetPlanShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

