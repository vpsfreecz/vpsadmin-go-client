package client

import (
	"strings"
)

// ActionDatasetExpansionUpdate is a type for action Dataset_expansion#Update
type ActionDatasetExpansionUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetExpansionUpdate(client *Client) *ActionDatasetExpansionUpdate {
	return &ActionDatasetExpansionUpdate{
		Client: client,
	}
}

// ActionDatasetExpansionUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetExpansionUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetExpansionUpdateMetaGlobalInput) SetIncludes(value string) *ActionDatasetExpansionUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetExpansionUpdateMetaGlobalInput) SetNo(value bool) *ActionDatasetExpansionUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetExpansionUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetExpansionUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionUpdateInput is a type for action input parameters
type ActionDatasetExpansionUpdateInput struct {
	EnableNotifications    bool  `json:"enable_notifications"`
	EnableShrink           bool  `json:"enable_shrink"`
	MaxOverRefquotaSeconds int64 `json:"max_over_refquota_seconds"`
	StopVps                bool  `json:"stop_vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnableNotifications sets parameter EnableNotifications to value and selects it for sending
func (in *ActionDatasetExpansionUpdateInput) SetEnableNotifications(value bool) *ActionDatasetExpansionUpdateInput {
	in.EnableNotifications = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableNotifications"] = nil
	return in
}

// SetEnableShrink sets parameter EnableShrink to value and selects it for sending
func (in *ActionDatasetExpansionUpdateInput) SetEnableShrink(value bool) *ActionDatasetExpansionUpdateInput {
	in.EnableShrink = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnableShrink"] = nil
	return in
}

// SetMaxOverRefquotaSeconds sets parameter MaxOverRefquotaSeconds to value and selects it for sending
func (in *ActionDatasetExpansionUpdateInput) SetMaxOverRefquotaSeconds(value int64) *ActionDatasetExpansionUpdateInput {
	in.MaxOverRefquotaSeconds = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxOverRefquotaSeconds"] = nil
	return in
}

// SetStopVps sets parameter StopVps to value and selects it for sending
func (in *ActionDatasetExpansionUpdateInput) SetStopVps(value bool) *ActionDatasetExpansionUpdateInput {
	in.StopVps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["StopVps"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionUpdateInput) SelectParameters(params ...string) *ActionDatasetExpansionUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDatasetExpansionUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDatasetExpansionUpdateInput) UnselectParameters(params ...string) *ActionDatasetExpansionUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDatasetExpansionUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionUpdateRequest is a type for the entire action request
type ActionDatasetExpansionUpdateRequest struct {
	DatasetExpansion map[string]interface{} `json:"dataset_expansion"`
	Meta             map[string]interface{} `json:"_meta"`
}

// ActionDatasetExpansionUpdateOutput is a type for action output parameters
type ActionDatasetExpansionUpdateOutput struct {
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
type ActionDatasetExpansionUpdateResponse struct {
	Action *ActionDatasetExpansionUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DatasetExpansion *ActionDatasetExpansionUpdateOutput `json:"dataset_expansion"`
	}

	// Action output without the namespace
	Output *ActionDatasetExpansionUpdateOutput
}

// Prepare the action for invocation
func (action *ActionDatasetExpansionUpdate) Prepare() *ActionDatasetExpansionUpdateInvocation {
	return &ActionDatasetExpansionUpdateInvocation{
		Action: action,
		Path:   "/v7.0/dataset_expansions/{dataset_expansion_id}",
	}
}

// ActionDatasetExpansionUpdateInvocation is used to configure action for invocation
type ActionDatasetExpansionUpdateInvocation struct {
	// Pointer to the action
	Action *ActionDatasetExpansionUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDatasetExpansionUpdateInput
	// Global meta input parameters
	MetaInput *ActionDatasetExpansionUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetExpansionUpdateInvocation) SetPathParamInt(param string, value int64) *ActionDatasetExpansionUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetExpansionUpdateInvocation) SetPathParamString(param string, value string) *ActionDatasetExpansionUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDatasetExpansionUpdateInvocation) NewInput() *ActionDatasetExpansionUpdateInput {
	inv.Input = &ActionDatasetExpansionUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDatasetExpansionUpdateInvocation) SetInput(input *ActionDatasetExpansionUpdateInput) *ActionDatasetExpansionUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDatasetExpansionUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDatasetExpansionUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetExpansionUpdateInvocation) NewMetaInput() *ActionDatasetExpansionUpdateMetaGlobalInput {
	inv.MetaInput = &ActionDatasetExpansionUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetExpansionUpdateInvocation) SetMetaInput(input *ActionDatasetExpansionUpdateMetaGlobalInput) *ActionDatasetExpansionUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetExpansionUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetExpansionUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetExpansionUpdateInvocation) Call() (*ActionDatasetExpansionUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDatasetExpansionUpdateInvocation) callAsBody() (*ActionDatasetExpansionUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDatasetExpansionUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DatasetExpansion
	}
	return resp, err
}

func (inv *ActionDatasetExpansionUpdateInvocation) makeAllInputParams() *ActionDatasetExpansionUpdateRequest {
	return &ActionDatasetExpansionUpdateRequest{
		DatasetExpansion: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionDatasetExpansionUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("EnableNotifications") {
			ret["enable_notifications"] = inv.Input.EnableNotifications
		}
		if inv.IsParameterSelected("EnableShrink") {
			ret["enable_shrink"] = inv.Input.EnableShrink
		}
		if inv.IsParameterSelected("MaxOverRefquotaSeconds") {
			ret["max_over_refquota_seconds"] = inv.Input.MaxOverRefquotaSeconds
		}
		if inv.IsParameterSelected("StopVps") {
			ret["stop_vps"] = inv.Input.StopVps
		}
	}

	return ret
}

func (inv *ActionDatasetExpansionUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
