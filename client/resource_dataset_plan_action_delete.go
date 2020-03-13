package client

import (
	"strings"
)

// ActionDatasetPlanDelete is a type for action Dataset.Plan#Delete
type ActionDatasetPlanDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetPlanDelete(client *Client) *ActionDatasetPlanDelete {
	return &ActionDatasetPlanDelete{
		Client: client,
	}
}

// ActionDatasetPlanDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetPlanDeleteMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetPlanDeleteMetaGlobalInput) SetNo(value bool) *ActionDatasetPlanDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetPlanDeleteMetaGlobalInput) SetIncludes(value string) *ActionDatasetPlanDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetPlanDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetPlanDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetPlanDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetPlanDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionDatasetPlanDeleteRequest is a type for the entire action request
type ActionDatasetPlanDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionDatasetPlanDeleteResponse struct {
	Action *ActionDatasetPlanDelete `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionDatasetPlanDelete) Prepare() *ActionDatasetPlanDeleteInvocation {
	return &ActionDatasetPlanDeleteInvocation{
		Action: action,
		Path: "/v6.0/datasets/{dataset_id}/plans/{plan_id}",
	}
}

// ActionDatasetPlanDeleteInvocation is used to configure action for invocation
type ActionDatasetPlanDeleteInvocation struct {
	// Pointer to the action
	Action *ActionDatasetPlanDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDatasetPlanDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetPlanDeleteInvocation) SetPathParamInt(param string, value int64) *ActionDatasetPlanDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetPlanDeleteInvocation) SetPathParamString(param string, value string) *ActionDatasetPlanDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetPlanDeleteInvocation) NewMetaInput() *ActionDatasetPlanDeleteMetaGlobalInput {
	inv.MetaInput = &ActionDatasetPlanDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetPlanDeleteInvocation) SetMetaInput(input *ActionDatasetPlanDeleteMetaGlobalInput) *ActionDatasetPlanDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetPlanDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetPlanDeleteInvocation) Call() (*ActionDatasetPlanDeleteResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionDatasetPlanDeleteInvocation) callAsBody() (*ActionDatasetPlanDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetPlanDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionDatasetPlanDeleteInvocation) makeAllInputParams() *ActionDatasetPlanDeleteRequest {
	return &ActionDatasetPlanDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}


func (inv *ActionDatasetPlanDeleteInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
