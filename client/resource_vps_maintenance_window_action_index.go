package client

import (
	"strings"
)

// ActionVpsMaintenanceWindowIndex is a type for action Vps.Maintenance_window#Index
type ActionVpsMaintenanceWindowIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMaintenanceWindowIndex(client *Client) *ActionVpsMaintenanceWindowIndex {
	return &ActionVpsMaintenanceWindowIndex{
		Client: client,
	}
}

// ActionVpsMaintenanceWindowIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMaintenanceWindowIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsMaintenanceWindowIndexMetaGlobalInput) SetCount(value bool) *ActionVpsMaintenanceWindowIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMaintenanceWindowIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsMaintenanceWindowIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMaintenanceWindowIndexMetaGlobalInput) SetNo(value bool) *ActionVpsMaintenanceWindowIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMaintenanceWindowIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMaintenanceWindowIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMaintenanceWindowIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMaintenanceWindowIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMaintenanceWindowIndexInput is a type for action input parameters
type ActionVpsMaintenanceWindowIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionVpsMaintenanceWindowIndexInput) SetFromId(value int64) *ActionVpsMaintenanceWindowIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsMaintenanceWindowIndexInput) SetLimit(value int64) *ActionVpsMaintenanceWindowIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMaintenanceWindowIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMaintenanceWindowIndexInput) SelectParameters(params ...string) *ActionVpsMaintenanceWindowIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsMaintenanceWindowIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsMaintenanceWindowIndexInput) UnselectParameters(params ...string) *ActionVpsMaintenanceWindowIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsMaintenanceWindowIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMaintenanceWindowIndexOutput is a type for action output parameters
type ActionVpsMaintenanceWindowIndexOutput struct {
	ClosesAt int64 `json:"closes_at"`
	IsOpen   bool  `json:"is_open"`
	OpensAt  int64 `json:"opens_at"`
	Weekday  int64 `json:"weekday"`
}

// Type for action response, including envelope
type ActionVpsMaintenanceWindowIndexResponse struct {
	Action *ActionVpsMaintenanceWindowIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MaintenanceWindows []*ActionVpsMaintenanceWindowIndexOutput `json:"maintenance_windows"`
	}

	// Action output without the namespace
	Output []*ActionVpsMaintenanceWindowIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsMaintenanceWindowIndex) Prepare() *ActionVpsMaintenanceWindowIndexInvocation {
	return &ActionVpsMaintenanceWindowIndexInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/maintenance_windows",
	}
}

// ActionVpsMaintenanceWindowIndexInvocation is used to configure action for invocation
type ActionVpsMaintenanceWindowIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsMaintenanceWindowIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsMaintenanceWindowIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsMaintenanceWindowIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMaintenanceWindowIndexInvocation) SetPathParamInt(param string, value int64) *ActionVpsMaintenanceWindowIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMaintenanceWindowIndexInvocation) SetPathParamString(param string, value string) *ActionVpsMaintenanceWindowIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsMaintenanceWindowIndexInvocation) NewInput() *ActionVpsMaintenanceWindowIndexInput {
	inv.Input = &ActionVpsMaintenanceWindowIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsMaintenanceWindowIndexInvocation) SetInput(input *ActionVpsMaintenanceWindowIndexInput) *ActionVpsMaintenanceWindowIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsMaintenanceWindowIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsMaintenanceWindowIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMaintenanceWindowIndexInvocation) NewMetaInput() *ActionVpsMaintenanceWindowIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsMaintenanceWindowIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMaintenanceWindowIndexInvocation) SetMetaInput(input *ActionVpsMaintenanceWindowIndexMetaGlobalInput) *ActionVpsMaintenanceWindowIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMaintenanceWindowIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsMaintenanceWindowIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMaintenanceWindowIndexInvocation) Call() (*ActionVpsMaintenanceWindowIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsMaintenanceWindowIndexInvocation) callAsQuery() (*ActionVpsMaintenanceWindowIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsMaintenanceWindowIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MaintenanceWindows
	}
	return resp, err
}

func (inv *ActionVpsMaintenanceWindowIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["maintenance_window[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["maintenance_window[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionVpsMaintenanceWindowIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
