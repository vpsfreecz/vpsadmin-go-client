package client

import ()

// ActionIntegrityCheckIndex is a type for action Integrity_check#Index
type ActionIntegrityCheckIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIntegrityCheckIndex(client *Client) *ActionIntegrityCheckIndex {
	return &ActionIntegrityCheckIndex{
		Client: client,
	}
}

// ActionIntegrityCheckIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIntegrityCheckIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIntegrityCheckIndexMetaGlobalInput) SetCount(value bool) *ActionIntegrityCheckIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIntegrityCheckIndexMetaGlobalInput) SetIncludes(value string) *ActionIntegrityCheckIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIntegrityCheckIndexMetaGlobalInput) SetNo(value bool) *ActionIntegrityCheckIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityCheckIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityCheckIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIntegrityCheckIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityCheckIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIntegrityCheckIndexInput is a type for action input parameters
type ActionIntegrityCheckIndexInput struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Status string `json:"status"`
	User   int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIntegrityCheckIndexInput) SetLimit(value int64) *ActionIntegrityCheckIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionIntegrityCheckIndexInput) SetOffset(value int64) *ActionIntegrityCheckIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetStatus sets parameter Status to value and selects it for sending
func (in *ActionIntegrityCheckIndexInput) SetStatus(value string) *ActionIntegrityCheckIndexInput {
	in.Status = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Status"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionIntegrityCheckIndexInput) SetUser(value int64) *ActionIntegrityCheckIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityCheckIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityCheckIndexInput) SelectParameters(params ...string) *ActionIntegrityCheckIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityCheckIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIntegrityCheckIndexOutput is a type for action output parameters
type ActionIntegrityCheckIndexOutput struct {
	BrokenObjects   int64                 `json:"broken_objects"`
	CheckedFacts    int64                 `json:"checked_facts"`
	CheckedObjects  int64                 `json:"checked_objects"`
	CreatedAt       string                `json:"created_at"`
	FalseFacts      int64                 `json:"false_facts"`
	FinishedAt      string                `json:"finished_at"`
	Id              int64                 `json:"id"`
	IntegralObjects int64                 `json:"integral_objects"`
	Status          string                `json:"status"`
	TrueFacts       int64                 `json:"true_facts"`
	UpdatedAt       string                `json:"updated_at"`
	User            *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionIntegrityCheckIndexResponse struct {
	Action *ActionIntegrityCheckIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IntegrityChecks []*ActionIntegrityCheckIndexOutput `json:"integrity_checks"`
	}

	// Action output without the namespace
	Output []*ActionIntegrityCheckIndexOutput
}

// Prepare the action for invocation
func (action *ActionIntegrityCheckIndex) Prepare() *ActionIntegrityCheckIndexInvocation {
	return &ActionIntegrityCheckIndexInvocation{
		Action: action,
		Path:   "/v6.0/integrity_checks",
	}
}

// ActionIntegrityCheckIndexInvocation is used to configure action for invocation
type ActionIntegrityCheckIndexInvocation struct {
	// Pointer to the action
	Action *ActionIntegrityCheckIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIntegrityCheckIndexInput
	// Global meta input parameters
	MetaInput *ActionIntegrityCheckIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIntegrityCheckIndexInvocation) NewInput() *ActionIntegrityCheckIndexInput {
	inv.Input = &ActionIntegrityCheckIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIntegrityCheckIndexInvocation) SetInput(input *ActionIntegrityCheckIndexInput) *ActionIntegrityCheckIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIntegrityCheckIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIntegrityCheckIndexInvocation) NewMetaInput() *ActionIntegrityCheckIndexMetaGlobalInput {
	inv.MetaInput = &ActionIntegrityCheckIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIntegrityCheckIndexInvocation) SetMetaInput(input *ActionIntegrityCheckIndexMetaGlobalInput) *ActionIntegrityCheckIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIntegrityCheckIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIntegrityCheckIndexInvocation) Call() (*ActionIntegrityCheckIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIntegrityCheckIndexInvocation) callAsQuery() (*ActionIntegrityCheckIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIntegrityCheckIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IntegrityChecks
	}
	return resp, err
}

func (inv *ActionIntegrityCheckIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["integrity_check[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["integrity_check[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Status") {
			ret["integrity_check[status]"] = inv.Input.Status
		}
		if inv.IsParameterSelected("User") {
			ret["integrity_check[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionIntegrityCheckIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
