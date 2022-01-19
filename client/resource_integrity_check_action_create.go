package client

import (
)

// ActionIntegrityCheckCreate is a type for action Integrity_check#Create
type ActionIntegrityCheckCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionIntegrityCheckCreate(client *Client) *ActionIntegrityCheckCreate {
	return &ActionIntegrityCheckCreate{
		Client: client,
	}
}

// ActionIntegrityCheckCreateMetaGlobalInput is a type for action global meta input parameters
type ActionIntegrityCheckCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIntegrityCheckCreateMetaGlobalInput) SetIncludes(value string) *ActionIntegrityCheckCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionIntegrityCheckCreateMetaGlobalInput) SetNo(value bool) *ActionIntegrityCheckCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityCheckCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityCheckCreateMetaGlobalInput) SelectParameters(params ...string) *ActionIntegrityCheckCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityCheckCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIntegrityCheckCreateInput is a type for action input parameters
type ActionIntegrityCheckCreateInput struct {
	Node int64 `json:"node"`
	SkipMaintenance bool `json:"skip_maintenance"`
	Storage bool `json:"storage"`
	Vps bool `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionIntegrityCheckCreateInput) SetNode(value int64) *ActionIntegrityCheckCreateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetSkipMaintenance sets parameter SkipMaintenance to value and selects it for sending
func (in *ActionIntegrityCheckCreateInput) SetSkipMaintenance(value bool) *ActionIntegrityCheckCreateInput {
	in.SkipMaintenance = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SkipMaintenance"] = nil
	return in
}
// SetStorage sets parameter Storage to value and selects it for sending
func (in *ActionIntegrityCheckCreateInput) SetStorage(value bool) *ActionIntegrityCheckCreateInput {
	in.Storage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Storage"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionIntegrityCheckCreateInput) SetVps(value bool) *ActionIntegrityCheckCreateInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityCheckCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityCheckCreateInput) SelectParameters(params ...string) *ActionIntegrityCheckCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityCheckCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIntegrityCheckCreateRequest is a type for the entire action request
type ActionIntegrityCheckCreateRequest struct {
	IntegrityCheck map[string]interface{} `json:"integrity_check"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionIntegrityCheckCreateOutput is a type for action output parameters
type ActionIntegrityCheckCreateOutput struct {
	BrokenObjects int64 `json:"broken_objects"`
	CheckedFacts int64 `json:"checked_facts"`
	CheckedObjects int64 `json:"checked_objects"`
	CreatedAt string `json:"created_at"`
	FalseFacts int64 `json:"false_facts"`
	FinishedAt string `json:"finished_at"`
	Id int64 `json:"id"`
	IntegralObjects int64 `json:"integral_objects"`
	Status string `json:"status"`
	TrueFacts int64 `json:"true_facts"`
	UpdatedAt string `json:"updated_at"`
	User *ActionUserShowOutput `json:"user"`
}


// Type for action response, including envelope
type ActionIntegrityCheckCreateResponse struct {
	Action *ActionIntegrityCheckCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IntegrityCheck *ActionIntegrityCheckCreateOutput `json:"integrity_check"`
	}

	// Action output without the namespace
	Output *ActionIntegrityCheckCreateOutput
}


// Prepare the action for invocation
func (action *ActionIntegrityCheckCreate) Prepare() *ActionIntegrityCheckCreateInvocation {
	return &ActionIntegrityCheckCreateInvocation{
		Action: action,
		Path: "/v6.0/integrity_checks",
	}
}

// ActionIntegrityCheckCreateInvocation is used to configure action for invocation
type ActionIntegrityCheckCreateInvocation struct {
	// Pointer to the action
	Action *ActionIntegrityCheckCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIntegrityCheckCreateInput
	// Global meta input parameters
	MetaInput *ActionIntegrityCheckCreateMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIntegrityCheckCreateInvocation) NewInput() *ActionIntegrityCheckCreateInput {
	inv.Input = &ActionIntegrityCheckCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIntegrityCheckCreateInvocation) SetInput(input *ActionIntegrityCheckCreateInput) *ActionIntegrityCheckCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIntegrityCheckCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIntegrityCheckCreateInvocation) NewMetaInput() *ActionIntegrityCheckCreateMetaGlobalInput {
	inv.MetaInput = &ActionIntegrityCheckCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIntegrityCheckCreateInvocation) SetMetaInput(input *ActionIntegrityCheckCreateMetaGlobalInput) *ActionIntegrityCheckCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIntegrityCheckCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIntegrityCheckCreateInvocation) Call() (*ActionIntegrityCheckCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionIntegrityCheckCreateInvocation) callAsBody() (*ActionIntegrityCheckCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionIntegrityCheckCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IntegrityCheck
	}
	return resp, err
}




func (inv *ActionIntegrityCheckCreateInvocation) makeAllInputParams() *ActionIntegrityCheckCreateRequest {
	return &ActionIntegrityCheckCreateRequest{
		IntegrityCheck: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionIntegrityCheckCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Node") {
			ret["node"] = inv.Input.Node
		}
		if inv.IsParameterSelected("SkipMaintenance") {
			ret["skip_maintenance"] = inv.Input.SkipMaintenance
		}
		if inv.IsParameterSelected("Storage") {
			ret["storage"] = inv.Input.Storage
		}
		if inv.IsParameterSelected("Vps") {
			ret["vps"] = inv.Input.Vps
		}
	}

	return ret
}

func (inv *ActionIntegrityCheckCreateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
