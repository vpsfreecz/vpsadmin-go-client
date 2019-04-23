package client

import (
)

// ActionClusterResourceCreate is a type for action Cluster_resource#Create
type ActionClusterResourceCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourceCreate(client *Client) *ActionClusterResourceCreate {
	return &ActionClusterResourceCreate{
		Client: client,
	}
}

// ActionClusterResourceCreateMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourceCreateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourceCreateMetaGlobalInput) SetNo(value bool) *ActionClusterResourceCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourceCreateMetaGlobalInput) SetIncludes(value string) *ActionClusterResourceCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourceCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourceCreateMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourceCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourceCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourceCreateInput is a type for action input parameters
type ActionClusterResourceCreateInput struct {
	Name string `json:"name"`
	Label string `json:"label"`
	Min int64 `json:"min"`
	Max int64 `json:"max"`
	Stepsize int64 `json:"stepsize"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionClusterResourceCreateInput) SetName(value string) *ActionClusterResourceCreateInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionClusterResourceCreateInput) SetLabel(value string) *ActionClusterResourceCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetMin sets parameter Min to value and selects it for sending
func (in *ActionClusterResourceCreateInput) SetMin(value int64) *ActionClusterResourceCreateInput {
	in.Min = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Min"] = nil
	return in
}
// SetMax sets parameter Max to value and selects it for sending
func (in *ActionClusterResourceCreateInput) SetMax(value int64) *ActionClusterResourceCreateInput {
	in.Max = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Max"] = nil
	return in
}
// SetStepsize sets parameter Stepsize to value and selects it for sending
func (in *ActionClusterResourceCreateInput) SetStepsize(value int64) *ActionClusterResourceCreateInput {
	in.Stepsize = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Stepsize"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourceCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourceCreateInput) SelectParameters(params ...string) *ActionClusterResourceCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourceCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourceCreateRequest is a type for the entire action request
type ActionClusterResourceCreateRequest struct {
	ClusterResource map[string]interface{} `json:"cluster_resource"`
	Meta map[string]interface{} `json:"_meta"`
}



// Type for action response, including envelope
type ActionClusterResourceCreateResponse struct {
	Action *ActionClusterResourceCreate `json:"-"`
	*Envelope
}


// Prepare the action for invocation
func (action *ActionClusterResourceCreate) Prepare() *ActionClusterResourceCreateInvocation {
	return &ActionClusterResourceCreateInvocation{
		Action: action,
		Path: "/v5.0/cluster_resources",
	}
}

// ActionClusterResourceCreateInvocation is used to configure action for invocation
type ActionClusterResourceCreateInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourceCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterResourceCreateInput
	// Global meta input parameters
	MetaInput *ActionClusterResourceCreateMetaGlobalInput
}


// SetInput provides input parameters to send to the API
func (inv *ActionClusterResourceCreateInvocation) SetInput(input *ActionClusterResourceCreateInput) *ActionClusterResourceCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterResourceCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourceCreateInvocation) SetMetaInput(input *ActionClusterResourceCreateMetaGlobalInput) *ActionClusterResourceCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourceCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterResourceCreateInvocation) Call() (*ActionClusterResourceCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterResourceCreateInvocation) callAsBody() (*ActionClusterResourceCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterResourceCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}




func (inv *ActionClusterResourceCreateInvocation) makeAllInputParams() *ActionClusterResourceCreateRequest {
	return &ActionClusterResourceCreateRequest{
		ClusterResource: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionClusterResourceCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionClusterResourceCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
