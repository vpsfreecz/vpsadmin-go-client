package client

import (
	"strings"
)

// ActionClusterResourceUpdate is a type for action Cluster_resource#Update
type ActionClusterResourceUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourceUpdate(client *Client) *ActionClusterResourceUpdate {
	return &ActionClusterResourceUpdate{
		Client: client,
	}
}

// ActionClusterResourceUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourceUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourceUpdateMetaGlobalInput) SetNo(value bool) *ActionClusterResourceUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourceUpdateMetaGlobalInput) SetIncludes(value string) *ActionClusterResourceUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourceUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourceUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourceUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourceUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourceUpdateInput is a type for action input parameters
type ActionClusterResourceUpdateInput struct {
	Name string `json:"name"`
	Label string `json:"label"`
	Min int64 `json:"min"`
	Max int64 `json:"max"`
	Stepsize int64 `json:"stepsize"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionClusterResourceUpdateInput) SetName(value string) *ActionClusterResourceUpdateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionClusterResourceUpdateInput) SetLabel(value string) *ActionClusterResourceUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetMin sets parameter Min to value and selects it for sending
func (in *ActionClusterResourceUpdateInput) SetMin(value int64) *ActionClusterResourceUpdateInput {
	in.Min = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Min"] = nil
	return in
}
// SetMax sets parameter Max to value and selects it for sending
func (in *ActionClusterResourceUpdateInput) SetMax(value int64) *ActionClusterResourceUpdateInput {
	in.Max = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Max"] = nil
	return in
}
// SetStepsize sets parameter Stepsize to value and selects it for sending
func (in *ActionClusterResourceUpdateInput) SetStepsize(value int64) *ActionClusterResourceUpdateInput {
	in.Stepsize = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Stepsize"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourceUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourceUpdateInput) SelectParameters(params ...string) *ActionClusterResourceUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourceUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourceUpdateRequest is a type for the entire action request
type ActionClusterResourceUpdateRequest struct {
	ClusterResource map[string]interface{} `json:"cluster_resource"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionClusterResourceUpdateOutput is a type for action output parameters
type ActionClusterResourceUpdateOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	Min int64 `json:"min"`
	Max int64 `json:"max"`
	Stepsize int64 `json:"stepsize"`
}


// Type for action response, including envelope
type ActionClusterResourceUpdateResponse struct {
	Action *ActionClusterResourceUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ClusterResource *ActionClusterResourceUpdateOutput `json:"cluster_resource"`
	}

	// Action output without the namespace
	Output *ActionClusterResourceUpdateOutput
}


// Prepare the action for invocation
func (action *ActionClusterResourceUpdate) Prepare() *ActionClusterResourceUpdateInvocation {
	return &ActionClusterResourceUpdateInvocation{
		Action: action,
		Path: "/v5.0/cluster_resources/{cluster_resource_id}",
	}
}

// ActionClusterResourceUpdateInvocation is used to configure action for invocation
type ActionClusterResourceUpdateInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourceUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterResourceUpdateInput
	// Global meta input parameters
	MetaInput *ActionClusterResourceUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourceUpdateInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourceUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourceUpdateInvocation) SetPathParamString(param string, value string) *ActionClusterResourceUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionClusterResourceUpdateInvocation) NewInput() *ActionClusterResourceUpdateInput {
	inv.Input = &ActionClusterResourceUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionClusterResourceUpdateInvocation) SetInput(input *ActionClusterResourceUpdateInput) *ActionClusterResourceUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterResourceUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourceUpdateInvocation) NewMetaInput() *ActionClusterResourceUpdateMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourceUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourceUpdateInvocation) SetMetaInput(input *ActionClusterResourceUpdateMetaGlobalInput) *ActionClusterResourceUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourceUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourceUpdateInvocation) Call() (*ActionClusterResourceUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterResourceUpdateInvocation) callAsBody() (*ActionClusterResourceUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterResourceUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ClusterResource
	}
	return resp, err
}




func (inv *ActionClusterResourceUpdateInvocation) makeAllInputParams() *ActionClusterResourceUpdateRequest {
	return &ActionClusterResourceUpdateRequest{
		ClusterResource: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionClusterResourceUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Name") {
			ret["name"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Min") {
			ret["min"] = inv.Input.Min
		}
		if inv.IsParameterSelected("Max") {
			ret["max"] = inv.Input.Max
		}
		if inv.IsParameterSelected("Stepsize") {
			ret["stepsize"] = inv.Input.Stepsize
		}
	}

	return ret
}

func (inv *ActionClusterResourceUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
