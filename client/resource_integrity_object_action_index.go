package client

import ()

// ActionIntegrityObjectIndex is a type for action Integrity_object#Index
type ActionIntegrityObjectIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIntegrityObjectIndex(client *Client) *ActionIntegrityObjectIndex {
	return &ActionIntegrityObjectIndex{
		Client: client,
	}
}

// ActionIntegrityObjectIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIntegrityObjectIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIntegrityObjectIndexMetaGlobalInput) SetCount(value bool) *ActionIntegrityObjectIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIntegrityObjectIndexMetaGlobalInput) SetIncludes(value string) *ActionIntegrityObjectIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIntegrityObjectIndexMetaGlobalInput) SetNo(value bool) *ActionIntegrityObjectIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityObjectIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityObjectIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIntegrityObjectIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityObjectIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIntegrityObjectIndexInput is a type for action input parameters
type ActionIntegrityObjectIndexInput struct {
	ClassName      string `json:"class_name"`
	IntegrityCheck int64  `json:"integrity_check"`
	Limit          int64  `json:"limit"`
	Node           int64  `json:"node"`
	Offset         int64  `json:"offset"`
	Parent         int64  `json:"parent"`
	RowId          int64  `json:"row_id"`
	Status         string `json:"status"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetClassName sets parameter ClassName to value and selects it for sending
func (in *ActionIntegrityObjectIndexInput) SetClassName(value string) *ActionIntegrityObjectIndexInput {
	in.ClassName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClassName"] = nil
	return in
}

// SetIntegrityCheck sets parameter IntegrityCheck to value and selects it for sending
func (in *ActionIntegrityObjectIndexInput) SetIntegrityCheck(value int64) *ActionIntegrityObjectIndexInput {
	in.IntegrityCheck = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IntegrityCheck"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIntegrityObjectIndexInput) SetLimit(value int64) *ActionIntegrityObjectIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionIntegrityObjectIndexInput) SetNode(value int64) *ActionIntegrityObjectIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIntegrityObjectIndexInput) SetOffset(value int64) *ActionIntegrityObjectIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetParent sets parameter Parent to value and selects it for sending
func (in *ActionIntegrityObjectIndexInput) SetParent(value int64) *ActionIntegrityObjectIndexInput {
	in.Parent = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Parent"] = nil
	return in
}

// SetRowId sets parameter RowId to value and selects it for sending
func (in *ActionIntegrityObjectIndexInput) SetRowId(value int64) *ActionIntegrityObjectIndexInput {
	in.RowId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RowId"] = nil
	return in
}

// SetStatus sets parameter Status to value and selects it for sending
func (in *ActionIntegrityObjectIndexInput) SetStatus(value string) *ActionIntegrityObjectIndexInput {
	in.Status = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Status"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityObjectIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityObjectIndexInput) SelectParameters(params ...string) *ActionIntegrityObjectIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityObjectIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIntegrityObjectIndexOutput is a type for action output parameters
type ActionIntegrityObjectIndexOutput struct {
	CheckedFacts   int64                            `json:"checked_facts"`
	ClassName      string                           `json:"class_name"`
	CreatedAt      string                           `json:"created_at"`
	FalseFacts     int64                            `json:"false_facts"`
	Id             int64                            `json:"id"`
	IntegrityCheck *ActionIntegrityCheckShowOutput  `json:"integrity_check"`
	Node           *ActionNodeShowOutput            `json:"node"`
	Parent         *ActionIntegrityObjectShowOutput `json:"parent"`
	RowId          int64                            `json:"row_id"`
	Status         string                           `json:"status"`
	TrueFacts      int64                            `json:"true_facts"`
	UpdatedAt      string                           `json:"updated_at"`
}

// Type for action response, including envelope
type ActionIntegrityObjectIndexResponse struct {
	Action *ActionIntegrityObjectIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IntegrityObjects []*ActionIntegrityObjectIndexOutput `json:"integrity_objects"`
	}

	// Action output without the namespace
	Output []*ActionIntegrityObjectIndexOutput
}

// Prepare the action for invocation
func (action *ActionIntegrityObjectIndex) Prepare() *ActionIntegrityObjectIndexInvocation {
	return &ActionIntegrityObjectIndexInvocation{
		Action: action,
		Path:   "/v6.0/integrity_objects",
	}
}

// ActionIntegrityObjectIndexInvocation is used to configure action for invocation
type ActionIntegrityObjectIndexInvocation struct {
	// Pointer to the action
	Action *ActionIntegrityObjectIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIntegrityObjectIndexInput
	// Global meta input parameters
	MetaInput *ActionIntegrityObjectIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIntegrityObjectIndexInvocation) NewInput() *ActionIntegrityObjectIndexInput {
	inv.Input = &ActionIntegrityObjectIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIntegrityObjectIndexInvocation) SetInput(input *ActionIntegrityObjectIndexInput) *ActionIntegrityObjectIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIntegrityObjectIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIntegrityObjectIndexInvocation) NewMetaInput() *ActionIntegrityObjectIndexMetaGlobalInput {
	inv.MetaInput = &ActionIntegrityObjectIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIntegrityObjectIndexInvocation) SetMetaInput(input *ActionIntegrityObjectIndexMetaGlobalInput) *ActionIntegrityObjectIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIntegrityObjectIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIntegrityObjectIndexInvocation) Call() (*ActionIntegrityObjectIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIntegrityObjectIndexInvocation) callAsQuery() (*ActionIntegrityObjectIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIntegrityObjectIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IntegrityObjects
	}
	return resp, err
}

func (inv *ActionIntegrityObjectIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("ClassName") {
			ret["integrity_object[class_name]"] = inv.Input.ClassName
		}
		if inv.IsParameterSelected("IntegrityCheck") {
			ret["integrity_object[integrity_check]"] = convertInt64ToString(inv.Input.IntegrityCheck)
		}
		if inv.IsParameterSelected("Limit") {
			ret["integrity_object[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["integrity_object[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Offset") {
			ret["integrity_object[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Parent") {
			ret["integrity_object[parent]"] = convertInt64ToString(inv.Input.Parent)
		}
		if inv.IsParameterSelected("RowId") {
			ret["integrity_object[row_id]"] = convertInt64ToString(inv.Input.RowId)
		}
		if inv.IsParameterSelected("Status") {
			ret["integrity_object[status]"] = inv.Input.Status
		}
	}
}

func (inv *ActionIntegrityObjectIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
