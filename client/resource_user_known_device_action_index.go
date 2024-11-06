package client

import (
	"strings"
)

// ActionUserKnownDeviceIndex is a type for action User.Known_device#Index
type ActionUserKnownDeviceIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserKnownDeviceIndex(client *Client) *ActionUserKnownDeviceIndex {
	return &ActionUserKnownDeviceIndex{
		Client: client,
	}
}

// ActionUserKnownDeviceIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserKnownDeviceIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserKnownDeviceIndexMetaGlobalInput) SetCount(value bool) *ActionUserKnownDeviceIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserKnownDeviceIndexMetaGlobalInput) SetIncludes(value string) *ActionUserKnownDeviceIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserKnownDeviceIndexMetaGlobalInput) SetNo(value bool) *ActionUserKnownDeviceIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserKnownDeviceIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserKnownDeviceIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserKnownDeviceIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserKnownDeviceIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserKnownDeviceIndexInput is a type for action input parameters
type ActionUserKnownDeviceIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserKnownDeviceIndexInput) SetFromId(value int64) *ActionUserKnownDeviceIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserKnownDeviceIndexInput) SetLimit(value int64) *ActionUserKnownDeviceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserKnownDeviceIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserKnownDeviceIndexInput) SelectParameters(params ...string) *ActionUserKnownDeviceIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserKnownDeviceIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserKnownDeviceIndexInput) UnselectParameters(params ...string) *ActionUserKnownDeviceIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserKnownDeviceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserKnownDeviceIndexOutput is a type for action output parameters
type ActionUserKnownDeviceIndexOutput struct {
	ApiIpAddr                string `json:"api_ip_addr"`
	ApiIpPtr                 string `json:"api_ip_ptr"`
	ClientIpAddr             string `json:"client_ip_addr"`
	ClientIpPtr              string `json:"client_ip_ptr"`
	CreatedAt                string `json:"created_at"`
	Id                       int64  `json:"id"`
	LastSeenAt               string `json:"last_seen_at"`
	SkipMultiFactorAuthUntil string `json:"skip_multi_factor_auth_until"`
	UpdatedAt                string `json:"updated_at"`
	UserAgent                string `json:"user_agent"`
}

// Type for action response, including envelope
type ActionUserKnownDeviceIndexResponse struct {
	Action *ActionUserKnownDeviceIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		KnownDevices []*ActionUserKnownDeviceIndexOutput `json:"known_devices"`
	}

	// Action output without the namespace
	Output []*ActionUserKnownDeviceIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserKnownDeviceIndex) Prepare() *ActionUserKnownDeviceIndexInvocation {
	return &ActionUserKnownDeviceIndexInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/known_devices",
	}
}

// ActionUserKnownDeviceIndexInvocation is used to configure action for invocation
type ActionUserKnownDeviceIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserKnownDeviceIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserKnownDeviceIndexInput
	// Global meta input parameters
	MetaInput *ActionUserKnownDeviceIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserKnownDeviceIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserKnownDeviceIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserKnownDeviceIndexInvocation) SetPathParamString(param string, value string) *ActionUserKnownDeviceIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserKnownDeviceIndexInvocation) NewInput() *ActionUserKnownDeviceIndexInput {
	inv.Input = &ActionUserKnownDeviceIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserKnownDeviceIndexInvocation) SetInput(input *ActionUserKnownDeviceIndexInput) *ActionUserKnownDeviceIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserKnownDeviceIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserKnownDeviceIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserKnownDeviceIndexInvocation) NewMetaInput() *ActionUserKnownDeviceIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserKnownDeviceIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserKnownDeviceIndexInvocation) SetMetaInput(input *ActionUserKnownDeviceIndexMetaGlobalInput) *ActionUserKnownDeviceIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserKnownDeviceIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserKnownDeviceIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserKnownDeviceIndexInvocation) Call() (*ActionUserKnownDeviceIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserKnownDeviceIndexInvocation) callAsQuery() (*ActionUserKnownDeviceIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserKnownDeviceIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.KnownDevices
	}
	return resp, err
}

func (inv *ActionUserKnownDeviceIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["known_device[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["known_device[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserKnownDeviceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
