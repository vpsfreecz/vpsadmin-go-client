package client

import ()

// ActionVpsPressureHistoryIndex is a type for action Vps_pressure_history#Index
type ActionVpsPressureHistoryIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsPressureHistoryIndex(client *Client) *ActionVpsPressureHistoryIndex {
	return &ActionVpsPressureHistoryIndex{
		Client: client,
	}
}

// ActionVpsPressureHistoryIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsPressureHistoryIndexMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsPressureHistoryIndexMetaGlobalInput) SetNo(value bool) *ActionVpsPressureHistoryIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsPressureHistoryIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsPressureHistoryIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsPressureHistoryIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsPressureHistoryIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsPressureHistoryIndexInput is a type for action input parameters
type ActionVpsPressureHistoryIndexInput struct {
	Days int64 "json:\"days\""
	Vps  int64 "json:\"vps\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDays sets parameter Days to value and selects it for sending
func (in *ActionVpsPressureHistoryIndexInput) SetDays(value int64) *ActionVpsPressureHistoryIndexInput {
	in.Days = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Days"] = nil
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionVpsPressureHistoryIndexInput) SetVps(value int64) *ActionVpsPressureHistoryIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsPressureHistoryIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsPressureHistoryIndexInput) SelectParameters(params ...string) *ActionVpsPressureHistoryIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsPressureHistoryIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsPressureHistoryIndexInput) UnselectParameters(params ...string) *ActionVpsPressureHistoryIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsPressureHistoryIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsPressureHistoryIndexOutput is a type for action output parameters
type ActionVpsPressureHistoryIndexOutput struct {
	FromAt     string "json:\"from_at\""
	Id         string "json:\"id\""
	PartialDay bool   "json:\"partial_day\""
	RecordedUs int64  "json:\"recorded_us\""
	ToAt       string "json:\"to_at\""
}

// Type for action response, including envelope
type ActionVpsPressureHistoryIndexResponse struct {
	Action *ActionVpsPressureHistoryIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsPressureHistory []*ActionVpsPressureHistoryIndexOutput "json:\"vps_pressure_history\""
	}

	// Action output without the namespace
	Output []*ActionVpsPressureHistoryIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsPressureHistoryIndex) Prepare() *ActionVpsPressureHistoryIndexInvocation {
	return &ActionVpsPressureHistoryIndexInvocation{
		Action: action,
		Path:   "/v7.0/vps_pressure_history",
	}
}

// ActionVpsPressureHistoryIndexInvocation is used to configure action for invocation
type ActionVpsPressureHistoryIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsPressureHistoryIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsPressureHistoryIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsPressureHistoryIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsPressureHistoryIndexInvocation) NewInput() *ActionVpsPressureHistoryIndexInput {
	inv.Input = &ActionVpsPressureHistoryIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsPressureHistoryIndexInvocation) SetInput(input *ActionVpsPressureHistoryIndexInput) *ActionVpsPressureHistoryIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsPressureHistoryIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsPressureHistoryIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsPressureHistoryIndexInvocation) NewMetaInput() *ActionVpsPressureHistoryIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsPressureHistoryIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsPressureHistoryIndexInvocation) SetMetaInput(input *ActionVpsPressureHistoryIndexMetaGlobalInput) *ActionVpsPressureHistoryIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsPressureHistoryIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsPressureHistoryIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionVpsPressureHistoryIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("Vps") {
			if !inv.IsParameterNil("Vps") {
				if inv.Input.Vps < 0 {
					verr.Add("vps", t(inv.Action.Client, "validation.invalid_resource_id", nil))
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsPressureHistoryIndexInvocation) Call() (*ActionVpsPressureHistoryIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionVpsPressureHistoryIndexInvocation) callAsQuery() (*ActionVpsPressureHistoryIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsPressureHistoryIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsPressureHistory
	}
	return resp, err
}

func (inv *ActionVpsPressureHistoryIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Days") {
			ret["vps_pressure_history[days]"] = convertInt64ToString(inv.Input.Days)
		}
		if inv.IsParameterSelected("Vps") {
			ret["vps_pressure_history[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionVpsPressureHistoryIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
