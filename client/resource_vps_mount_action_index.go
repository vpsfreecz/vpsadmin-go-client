package client

import (
	"strings"
)

// ActionVpsMountIndex is a type for action Vps.Mount#Index
type ActionVpsMountIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMountIndex(client *Client) *ActionVpsMountIndex {
	return &ActionVpsMountIndex{
		Client: client,
	}
}

// ActionVpsMountIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMountIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsMountIndexMetaGlobalInput) SetCount(value bool) *ActionVpsMountIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMountIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsMountIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMountIndexMetaGlobalInput) SetNo(value bool) *ActionVpsMountIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMountIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMountIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMountIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMountIndexInput is a type for action input parameters
type ActionVpsMountIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsMountIndexInput) SetLimit(value int64) *ActionVpsMountIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionVpsMountIndexInput) SetOffset(value int64) *ActionVpsMountIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMountIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountIndexInput) SelectParameters(params ...string) *ActionVpsMountIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsMountIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsMountIndexInput) UnselectParameters(params ...string) *ActionVpsMountIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsMountIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMountIndexOutput is a type for action output parameters
type ActionVpsMountIndexOutput struct {
	CurrentState     string                            `json:"current_state"`
	Dataset          *ActionDatasetShowOutput          `json:"dataset"`
	Enabled          bool                              `json:"enabled"`
	ExpirationDate   string                            `json:"expiration_date"`
	Id               int64                             `json:"id"`
	MasterEnabled    bool                              `json:"master_enabled"`
	Mode             string                            `json:"mode"`
	Mountpoint       string                            `json:"mountpoint"`
	OnStartFail      string                            `json:"on_start_fail"`
	Snapshot         *ActionDatasetSnapshotShowOutput  `json:"snapshot"`
	UserNamespaceMap *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
	Vps              *ActionVpsShowOutput              `json:"vps"`
}

// Type for action response, including envelope
type ActionVpsMountIndexResponse struct {
	Action *ActionVpsMountIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Mounts []*ActionVpsMountIndexOutput `json:"mounts"`
	}

	// Action output without the namespace
	Output []*ActionVpsMountIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsMountIndex) Prepare() *ActionVpsMountIndexInvocation {
	return &ActionVpsMountIndexInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/mounts",
	}
}

// ActionVpsMountIndexInvocation is used to configure action for invocation
type ActionVpsMountIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsMountIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsMountIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsMountIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMountIndexInvocation) SetPathParamInt(param string, value int64) *ActionVpsMountIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMountIndexInvocation) SetPathParamString(param string, value string) *ActionVpsMountIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsMountIndexInvocation) NewInput() *ActionVpsMountIndexInput {
	inv.Input = &ActionVpsMountIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsMountIndexInvocation) SetInput(input *ActionVpsMountIndexInput) *ActionVpsMountIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsMountIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsMountIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMountIndexInvocation) NewMetaInput() *ActionVpsMountIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsMountIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMountIndexInvocation) SetMetaInput(input *ActionVpsMountIndexMetaGlobalInput) *ActionVpsMountIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMountIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsMountIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMountIndexInvocation) Call() (*ActionVpsMountIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsMountIndexInvocation) callAsQuery() (*ActionVpsMountIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsMountIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Mounts
	}
	return resp, err
}

func (inv *ActionVpsMountIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["mount[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["mount[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionVpsMountIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
