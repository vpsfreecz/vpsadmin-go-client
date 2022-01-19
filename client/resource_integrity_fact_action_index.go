package client

import (
)

// ActionIntegrityFactIndex is a type for action Integrity_fact#Index
type ActionIntegrityFactIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIntegrityFactIndex(client *Client) *ActionIntegrityFactIndex {
	return &ActionIntegrityFactIndex{
		Client: client,
	}
}

// ActionIntegrityFactIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIntegrityFactIndexMetaGlobalInput struct {
	Count bool `json:"count"`
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIntegrityFactIndexMetaGlobalInput) SetCount(value bool) *ActionIntegrityFactIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIntegrityFactIndexMetaGlobalInput) SetIncludes(value string) *ActionIntegrityFactIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionIntegrityFactIndexMetaGlobalInput) SetNo(value bool) *ActionIntegrityFactIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityFactIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityFactIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIntegrityFactIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityFactIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIntegrityFactIndexInput is a type for action input parameters
type ActionIntegrityFactIndexInput struct {
	ClassName string `json:"class_name"`
	IntegrityCheck int64 `json:"integrity_check"`
	IntegrityObject int64 `json:"integrity_object"`
	Limit int64 `json:"limit"`
	Name string `json:"name"`
	Node int64 `json:"node"`
	Offset int64 `json:"offset"`
	Severity string `json:"severity"`
	Status string `json:"status"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetClassName sets parameter ClassName to value and selects it for sending
func (in *ActionIntegrityFactIndexInput) SetClassName(value string) *ActionIntegrityFactIndexInput {
	in.ClassName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClassName"] = nil
	return in
}
// SetIntegrityCheck sets parameter IntegrityCheck to value and selects it for sending
func (in *ActionIntegrityFactIndexInput) SetIntegrityCheck(value int64) *ActionIntegrityFactIndexInput {
	in.IntegrityCheck = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IntegrityCheck"] = nil
	return in
}
// SetIntegrityObject sets parameter IntegrityObject to value and selects it for sending
func (in *ActionIntegrityFactIndexInput) SetIntegrityObject(value int64) *ActionIntegrityFactIndexInput {
	in.IntegrityObject = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IntegrityObject"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIntegrityFactIndexInput) SetLimit(value int64) *ActionIntegrityFactIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetName sets parameter Name to value and selects it for sending
func (in *ActionIntegrityFactIndexInput) SetName(value string) *ActionIntegrityFactIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionIntegrityFactIndexInput) SetNode(value int64) *ActionIntegrityFactIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIntegrityFactIndexInput) SetOffset(value int64) *ActionIntegrityFactIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetSeverity sets parameter Severity to value and selects it for sending
func (in *ActionIntegrityFactIndexInput) SetSeverity(value string) *ActionIntegrityFactIndexInput {
	in.Severity = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Severity"] = nil
	return in
}
// SetStatus sets parameter Status to value and selects it for sending
func (in *ActionIntegrityFactIndexInput) SetStatus(value string) *ActionIntegrityFactIndexInput {
	in.Status = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Status"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityFactIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityFactIndexInput) SelectParameters(params ...string) *ActionIntegrityFactIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityFactIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionIntegrityFactIndexOutput is a type for action output parameters
type ActionIntegrityFactIndexOutput struct {
	CreatedAt string `json:"created_at"`
	Id int64 `json:"id"`
	IntegrityObject *ActionIntegrityObjectShowOutput `json:"integrity_object"`
	Message string `json:"message"`
	Name string `json:"name"`
	Severity string `json:"severity"`
	Status string `json:"status"`
}


// Type for action response, including envelope
type ActionIntegrityFactIndexResponse struct {
	Action *ActionIntegrityFactIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IntegrityFacts []*ActionIntegrityFactIndexOutput `json:"integrity_facts"`
	}

	// Action output without the namespace
	Output []*ActionIntegrityFactIndexOutput
}


// Prepare the action for invocation
func (action *ActionIntegrityFactIndex) Prepare() *ActionIntegrityFactIndexInvocation {
	return &ActionIntegrityFactIndexInvocation{
		Action: action,
		Path: "/v6.0/integrity_facts",
	}
}

// ActionIntegrityFactIndexInvocation is used to configure action for invocation
type ActionIntegrityFactIndexInvocation struct {
	// Pointer to the action
	Action *ActionIntegrityFactIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIntegrityFactIndexInput
	// Global meta input parameters
	MetaInput *ActionIntegrityFactIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIntegrityFactIndexInvocation) NewInput() *ActionIntegrityFactIndexInput {
	inv.Input = &ActionIntegrityFactIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIntegrityFactIndexInvocation) SetInput(input *ActionIntegrityFactIndexInput) *ActionIntegrityFactIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIntegrityFactIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIntegrityFactIndexInvocation) NewMetaInput() *ActionIntegrityFactIndexMetaGlobalInput {
	inv.MetaInput = &ActionIntegrityFactIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIntegrityFactIndexInvocation) SetMetaInput(input *ActionIntegrityFactIndexMetaGlobalInput) *ActionIntegrityFactIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIntegrityFactIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIntegrityFactIndexInvocation) Call() (*ActionIntegrityFactIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIntegrityFactIndexInvocation) callAsQuery() (*ActionIntegrityFactIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIntegrityFactIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IntegrityFacts
	}
	return resp, err
}



func (inv *ActionIntegrityFactIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("ClassName") {
			ret["integrity_fact[class_name]"] = inv.Input.ClassName
		}
		if inv.IsParameterSelected("IntegrityCheck") {
			ret["integrity_fact[integrity_check]"] = convertInt64ToString(inv.Input.IntegrityCheck)
		}
		if inv.IsParameterSelected("IntegrityObject") {
			ret["integrity_fact[integrity_object]"] = convertInt64ToString(inv.Input.IntegrityObject)
		}
		if inv.IsParameterSelected("Limit") {
			ret["integrity_fact[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["integrity_fact[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["integrity_fact[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Offset") {
			ret["integrity_fact[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Severity") {
			ret["integrity_fact[severity]"] = inv.Input.Severity
		}
		if inv.IsParameterSelected("Status") {
			ret["integrity_fact[status]"] = inv.Input.Status
		}
	}
}

func (inv *ActionIntegrityFactIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

