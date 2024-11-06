package client

import ()

// ActionDatasetExpansionRegisterExpanded is a type for action Dataset_expansion#Register_expanded
type ActionDatasetExpansionRegisterExpanded struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetExpansionRegisterExpanded(client *Client) *ActionDatasetExpansionRegisterExpanded {
	return &ActionDatasetExpansionRegisterExpanded{
		Client: client,
	}
}

// ActionDatasetExpansionRegisterExpandedMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetExpansionRegisterExpandedMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetExpansionRegisterExpandedMetaGlobalInput) SetIncludes(value string) *ActionDatasetExpansionRegisterExpandedMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetExpansionRegisterExpandedMetaGlobalInput) SetNo(value bool) *ActionDatasetExpansionRegisterExpandedMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionRegisterExpandedMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionRegisterExpandedMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetExpansionRegisterExpandedMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetExpansionRegisterExpandedMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionRegisterExpandedInput is a type for action input parameters
type ActionDatasetExpansionRegisterExpandedInput struct {
	Dataset                int64 `json:"dataset"`
	EnableNotifications    bool  `json:"enable_notifications"`
	EnableShrink           bool  `json:"enable_shrink"`
	MaxOverRefquotaSeconds int64 `json:"max_over_refquota_seconds"`
	OriginalRefquota       int64 `json:"original_refquota"`
	StopVps                bool  `json:"stop_vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDataset sets parameter Dataset to value and selects it for sending
func (in *ActionDatasetExpansionRegisterExpandedInput) SetDataset(value int64) *ActionDatasetExpansionRegisterExpandedInput {
	in.Dataset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetDatasetNil(false)
	in._selectedParameters["Dataset"] = nil
	return in
}

// SetDatasetNil sets parameter Dataset to nil and selects it for sending
func (in *ActionDatasetExpansionRegisterExpandedInput) SetDatasetNil(set bool) *ActionDatasetExpansionRegisterExpandedInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Dataset"] = nil
		in.SelectParameters("Dataset")
	} else {
		delete(in._nilParameters, "Dataset")
	}
	return in
}

// SetEnableNotifications sets parameter EnableNotifications to value and selects it for sending
func (in *ActionDatasetExpansionRegisterExpandedInput) SetEnableNotifications(value bool) *ActionDatasetExpansionRegisterExpandedInput {
	in.EnableNotifications = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableNotifications"] = nil
	return in
}

// SetEnableShrink sets parameter EnableShrink to value and selects it for sending
func (in *ActionDatasetExpansionRegisterExpandedInput) SetEnableShrink(value bool) *ActionDatasetExpansionRegisterExpandedInput {
	in.EnableShrink = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableShrink"] = nil
	return in
}

// SetMaxOverRefquotaSeconds sets parameter MaxOverRefquotaSeconds to value and selects it for sending
func (in *ActionDatasetExpansionRegisterExpandedInput) SetMaxOverRefquotaSeconds(value int64) *ActionDatasetExpansionRegisterExpandedInput {
	in.MaxOverRefquotaSeconds = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxOverRefquotaSeconds"] = nil
	return in
}

// SetOriginalRefquota sets parameter OriginalRefquota to value and selects it for sending
func (in *ActionDatasetExpansionRegisterExpandedInput) SetOriginalRefquota(value int64) *ActionDatasetExpansionRegisterExpandedInput {
	in.OriginalRefquota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["OriginalRefquota"] = nil
	return in
}

// SetStopVps sets parameter StopVps to value and selects it for sending
func (in *ActionDatasetExpansionRegisterExpandedInput) SetStopVps(value bool) *ActionDatasetExpansionRegisterExpandedInput {
	in.StopVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["StopVps"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionRegisterExpandedInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionRegisterExpandedInput) SelectParameters(params ...string) *ActionDatasetExpansionRegisterExpandedInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetExpansionRegisterExpandedInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetExpansionRegisterExpandedInput) UnselectParameters(params ...string) *ActionDatasetExpansionRegisterExpandedInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetExpansionRegisterExpandedInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionRegisterExpandedRequest is a type for the entire action request
type ActionDatasetExpansionRegisterExpandedRequest struct {
	DatasetExpansion map[string]interface{} `json:"dataset_expansion"`
	Meta             map[string]interface{} `json:"_meta"`
}

// ActionDatasetExpansionRegisterExpandedOutput is a type for action output parameters
type ActionDatasetExpansionRegisterExpandedOutput struct {
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
type ActionDatasetExpansionRegisterExpandedResponse struct {
	Action *ActionDatasetExpansionRegisterExpanded `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DatasetExpansion *ActionDatasetExpansionRegisterExpandedOutput `json:"dataset_expansion"`
	}

	// Action output without the namespace
	Output *ActionDatasetExpansionRegisterExpandedOutput
}

// Prepare the action for invocation
func (action *ActionDatasetExpansionRegisterExpanded) Prepare() *ActionDatasetExpansionRegisterExpandedInvocation {
	return &ActionDatasetExpansionRegisterExpandedInvocation{
		Action: action,
		Path:   "/v7.0/dataset_expansions/register_expanded",
	}
}

// ActionDatasetExpansionRegisterExpandedInvocation is used to configure action for invocation
type ActionDatasetExpansionRegisterExpandedInvocation struct {
	// Pointer to the action
	Action *ActionDatasetExpansionRegisterExpanded

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetExpansionRegisterExpandedInput
	// Global meta input parameters
	MetaInput *ActionDatasetExpansionRegisterExpandedMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetExpansionRegisterExpandedInvocation) NewInput() *ActionDatasetExpansionRegisterExpandedInput {
	inv.Input = &ActionDatasetExpansionRegisterExpandedInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetExpansionRegisterExpandedInvocation) SetInput(input *ActionDatasetExpansionRegisterExpandedInput) *ActionDatasetExpansionRegisterExpandedInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetExpansionRegisterExpandedInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetExpansionRegisterExpandedInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetExpansionRegisterExpandedInvocation) NewMetaInput() *ActionDatasetExpansionRegisterExpandedMetaGlobalInput {
	inv.MetaInput = &ActionDatasetExpansionRegisterExpandedMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetExpansionRegisterExpandedInvocation) SetMetaInput(input *ActionDatasetExpansionRegisterExpandedMetaGlobalInput) *ActionDatasetExpansionRegisterExpandedInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetExpansionRegisterExpandedInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetExpansionRegisterExpandedInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetExpansionRegisterExpandedInvocation) Call() (*ActionDatasetExpansionRegisterExpandedResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetExpansionRegisterExpandedInvocation) callAsBody() (*ActionDatasetExpansionRegisterExpandedResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetExpansionRegisterExpandedResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DatasetExpansion
	}
	return resp, err
}

func (inv *ActionDatasetExpansionRegisterExpandedInvocation) makeAllInputParams() *ActionDatasetExpansionRegisterExpandedRequest {
	return &ActionDatasetExpansionRegisterExpandedRequest{
		DatasetExpansion: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetExpansionRegisterExpandedInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Dataset") {
			if inv.IsParameterNil("Dataset") {
				ret["dataset"] = nil
			} else {
				ret["dataset"] = inv.Input.Dataset
			}
		}
		if inv.IsParameterSelected("EnableNotifications") {
			ret["enable_notifications"] = inv.Input.EnableNotifications
		}
		if inv.IsParameterSelected("EnableShrink") {
			ret["enable_shrink"] = inv.Input.EnableShrink
		}
		if inv.IsParameterSelected("MaxOverRefquotaSeconds") {
			ret["max_over_refquota_seconds"] = inv.Input.MaxOverRefquotaSeconds
		}
		if inv.IsParameterSelected("OriginalRefquota") {
			ret["original_refquota"] = inv.Input.OriginalRefquota
		}
		if inv.IsParameterSelected("StopVps") {
			ret["stop_vps"] = inv.Input.StopVps
		}
	}

	return ret
}

func (inv *ActionDatasetExpansionRegisterExpandedInvocation) makeMetaInputParams() map[string]interface{} {
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
