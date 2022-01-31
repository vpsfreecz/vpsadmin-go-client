package client

import (
	"strings"
)

// ActionUserTotpDeviceIndex is a type for action User.Totp_device#Index
type ActionUserTotpDeviceIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTotpDeviceIndex(client *Client) *ActionUserTotpDeviceIndex {
	return &ActionUserTotpDeviceIndex{
		Client: client,
	}
}

// ActionUserTotpDeviceIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserTotpDeviceIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserTotpDeviceIndexMetaGlobalInput) SetCount(value bool) *ActionUserTotpDeviceIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserTotpDeviceIndexMetaGlobalInput) SetIncludes(value string) *ActionUserTotpDeviceIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTotpDeviceIndexMetaGlobalInput) SetNo(value bool) *ActionUserTotpDeviceIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserTotpDeviceIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpDeviceIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceIndexInput is a type for action input parameters
type ActionUserTotpDeviceIndexInput struct {
	Confirmed bool  `json:"confirmed"`
	Enabled   bool  `json:"enabled"`
	Limit     int64 `json:"limit"`
	Offset    int64 `json:"offset"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetConfirmed sets parameter Confirmed to value and selects it for sending
func (in *ActionUserTotpDeviceIndexInput) SetConfirmed(value bool) *ActionUserTotpDeviceIndexInput {
	in.Confirmed = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Confirmed"] = nil
	return in
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionUserTotpDeviceIndexInput) SetEnabled(value bool) *ActionUserTotpDeviceIndexInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserTotpDeviceIndexInput) SetLimit(value int64) *ActionUserTotpDeviceIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserTotpDeviceIndexInput) SetOffset(value int64) *ActionUserTotpDeviceIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceIndexInput) SelectParameters(params ...string) *ActionUserTotpDeviceIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserTotpDeviceIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceIndexInput) UnselectParameters(params ...string) *ActionUserTotpDeviceIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserTotpDeviceIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceIndexOutput is a type for action output parameters
type ActionUserTotpDeviceIndexOutput struct {
	Confirmed bool   `json:"confirmed"`
	CreatedAt string `json:"created_at"`
	Enabled   bool   `json:"enabled"`
	Id        int64  `json:"id"`
	Label     string `json:"label"`
	LastUseAt string `json:"last_use_at"`
	UpdatedAt string `json:"updated_at"`
	UseCount  int64  `json:"use_count"`
}

// Type for action response, including envelope
type ActionUserTotpDeviceIndexResponse struct {
	Action *ActionUserTotpDeviceIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TotpDevices []*ActionUserTotpDeviceIndexOutput `json:"totp_devices"`
	}

	// Action output without the namespace
	Output []*ActionUserTotpDeviceIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserTotpDeviceIndex) Prepare() *ActionUserTotpDeviceIndexInvocation {
	return &ActionUserTotpDeviceIndexInvocation{
		Action: action,
		Path:   "/v6.0/users/{user_id}/totp_devices",
	}
}

// ActionUserTotpDeviceIndexInvocation is used to configure action for invocation
type ActionUserTotpDeviceIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserTotpDeviceIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserTotpDeviceIndexInput
	// Global meta input parameters
	MetaInput *ActionUserTotpDeviceIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTotpDeviceIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserTotpDeviceIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTotpDeviceIndexInvocation) SetPathParamString(param string, value string) *ActionUserTotpDeviceIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserTotpDeviceIndexInvocation) NewInput() *ActionUserTotpDeviceIndexInput {
	inv.Input = &ActionUserTotpDeviceIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserTotpDeviceIndexInvocation) SetInput(input *ActionUserTotpDeviceIndexInput) *ActionUserTotpDeviceIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserTotpDeviceIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserTotpDeviceIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTotpDeviceIndexInvocation) NewMetaInput() *ActionUserTotpDeviceIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserTotpDeviceIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTotpDeviceIndexInvocation) SetMetaInput(input *ActionUserTotpDeviceIndexMetaGlobalInput) *ActionUserTotpDeviceIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTotpDeviceIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserTotpDeviceIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserTotpDeviceIndexInvocation) Call() (*ActionUserTotpDeviceIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserTotpDeviceIndexInvocation) callAsQuery() (*ActionUserTotpDeviceIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserTotpDeviceIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TotpDevices
	}
	return resp, err
}

func (inv *ActionUserTotpDeviceIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Confirmed") {
			ret["totp_device[confirmed]"] = convertBoolToString(inv.Input.Confirmed)
		}
		if inv.IsParameterSelected("Enabled") {
			ret["totp_device[enabled]"] = convertBoolToString(inv.Input.Enabled)
		}
		if inv.IsParameterSelected("Limit") {
			ret["totp_device[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["totp_device[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
	}
}

func (inv *ActionUserTotpDeviceIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
