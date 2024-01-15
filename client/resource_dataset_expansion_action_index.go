package client

import ()

// ActionDatasetExpansionIndex is a type for action Dataset_expansion#Index
type ActionDatasetExpansionIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetExpansionIndex(client *Client) *ActionDatasetExpansionIndex {
	return &ActionDatasetExpansionIndex{
		Client: client,
	}
}

// ActionDatasetExpansionIndexMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetExpansionIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionDatasetExpansionIndexMetaGlobalInput) SetCount(value bool) *ActionDatasetExpansionIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetExpansionIndexMetaGlobalInput) SetIncludes(value string) *ActionDatasetExpansionIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetExpansionIndexMetaGlobalInput) SetNo(value bool) *ActionDatasetExpansionIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionIndexMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetExpansionIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetExpansionIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionIndexInput is a type for action input parameters
type ActionDatasetExpansionIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionDatasetExpansionIndexInput) SetLimit(value int64) *ActionDatasetExpansionIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionDatasetExpansionIndexInput) SetOffset(value int64) *ActionDatasetExpansionIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionIndexInput) SelectParameters(params ...string) *ActionDatasetExpansionIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetExpansionIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetExpansionIndexInput) UnselectParameters(params ...string) *ActionDatasetExpansionIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetExpansionIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionIndexOutput is a type for action output parameters
type ActionDatasetExpansionIndexOutput struct {
	AddedSpace             int64                    `json:"added_space"`
	CreatedAt              string                   `json:"created_at"`
	Dataset                *ActionDatasetShowOutput `json:"dataset"`
	EnableNotifications    bool                     `json:"enable_notifications"`
	EnableShrink           bool                     `json:"enable_shrink"`
	Id                     int64                    `json:"id"`
	MaxOverRefquotaSeconds int64                    `json:"max_over_refquota_seconds"`
	OriginalRefquota       int64                    `json:"original_refquota"`
	OverRefquotaSeconds    int64                    `json:"over_refquota_seconds"`
	State                  string                   `json:"state"`
	StopVps                bool                     `json:"stop_vps"`
	Vps                    *ActionVpsShowOutput     `json:"vps"`
}

// Type for action response, including envelope
type ActionDatasetExpansionIndexResponse struct {
	Action *ActionDatasetExpansionIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DatasetExpansions []*ActionDatasetExpansionIndexOutput `json:"dataset_expansions"`
	}

	// Action output without the namespace
	Output []*ActionDatasetExpansionIndexOutput
}

// Prepare the action for invocation
func (action *ActionDatasetExpansionIndex) Prepare() *ActionDatasetExpansionIndexInvocation {
	return &ActionDatasetExpansionIndexInvocation{
		Action: action,
		Path:   "/v6.0/dataset_expansions",
	}
}

// ActionDatasetExpansionIndexInvocation is used to configure action for invocation
type ActionDatasetExpansionIndexInvocation struct {
	// Pointer to the action
	Action *ActionDatasetExpansionIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetExpansionIndexInput
	// Global meta input parameters
	MetaInput *ActionDatasetExpansionIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetExpansionIndexInvocation) NewInput() *ActionDatasetExpansionIndexInput {
	inv.Input = &ActionDatasetExpansionIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetExpansionIndexInvocation) SetInput(input *ActionDatasetExpansionIndexInput) *ActionDatasetExpansionIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetExpansionIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetExpansionIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetExpansionIndexInvocation) NewMetaInput() *ActionDatasetExpansionIndexMetaGlobalInput {
	inv.MetaInput = &ActionDatasetExpansionIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetExpansionIndexInvocation) SetMetaInput(input *ActionDatasetExpansionIndexMetaGlobalInput) *ActionDatasetExpansionIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetExpansionIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetExpansionIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetExpansionIndexInvocation) Call() (*ActionDatasetExpansionIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetExpansionIndexInvocation) callAsQuery() (*ActionDatasetExpansionIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetExpansionIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DatasetExpansions
	}
	return resp, err
}

func (inv *ActionDatasetExpansionIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["dataset_expansion[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["dataset_expansion[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionDatasetExpansionIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
