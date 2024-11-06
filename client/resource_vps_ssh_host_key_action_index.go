package client

import (
	"strings"
)

// ActionVpsSshHostKeyIndex is a type for action Vps.Ssh_host_key#Index
type ActionVpsSshHostKeyIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsSshHostKeyIndex(client *Client) *ActionVpsSshHostKeyIndex {
	return &ActionVpsSshHostKeyIndex{
		Client: client,
	}
}

// ActionVpsSshHostKeyIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsSshHostKeyIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsSshHostKeyIndexMetaGlobalInput) SetCount(value bool) *ActionVpsSshHostKeyIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsSshHostKeyIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsSshHostKeyIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsSshHostKeyIndexMetaGlobalInput) SetNo(value bool) *ActionVpsSshHostKeyIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsSshHostKeyIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsSshHostKeyIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsSshHostKeyIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsSshHostKeyIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsSshHostKeyIndexInput is a type for action input parameters
type ActionVpsSshHostKeyIndexInput struct {
	Algorithm string `json:"algorithm"`
	FromId    int64  `json:"from_id"`
	Limit     int64  `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAlgorithm sets parameter Algorithm to value and selects it for sending
func (in *ActionVpsSshHostKeyIndexInput) SetAlgorithm(value string) *ActionVpsSshHostKeyIndexInput {
	in.Algorithm = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Algorithm"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionVpsSshHostKeyIndexInput) SetFromId(value int64) *ActionVpsSshHostKeyIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsSshHostKeyIndexInput) SetLimit(value int64) *ActionVpsSshHostKeyIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsSshHostKeyIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsSshHostKeyIndexInput) SelectParameters(params ...string) *ActionVpsSshHostKeyIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsSshHostKeyIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsSshHostKeyIndexInput) UnselectParameters(params ...string) *ActionVpsSshHostKeyIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsSshHostKeyIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsSshHostKeyIndexOutput is a type for action output parameters
type ActionVpsSshHostKeyIndexOutput struct {
	Algorithm   string `json:"algorithm"`
	Bits        int64  `json:"bits"`
	CreatedAt   string `json:"created_at"`
	Fingerprint string `json:"fingerprint"`
	Id          int64  `json:"id"`
	UpdatedAt   string `json:"updated_at"`
}

// Type for action response, including envelope
type ActionVpsSshHostKeyIndexResponse struct {
	Action *ActionVpsSshHostKeyIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SshHostKeys []*ActionVpsSshHostKeyIndexOutput `json:"ssh_host_keys"`
	}

	// Action output without the namespace
	Output []*ActionVpsSshHostKeyIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsSshHostKeyIndex) Prepare() *ActionVpsSshHostKeyIndexInvocation {
	return &ActionVpsSshHostKeyIndexInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/ssh_host_keys",
	}
}

// ActionVpsSshHostKeyIndexInvocation is used to configure action for invocation
type ActionVpsSshHostKeyIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsSshHostKeyIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsSshHostKeyIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsSshHostKeyIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsSshHostKeyIndexInvocation) SetPathParamInt(param string, value int64) *ActionVpsSshHostKeyIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsSshHostKeyIndexInvocation) SetPathParamString(param string, value string) *ActionVpsSshHostKeyIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsSshHostKeyIndexInvocation) NewInput() *ActionVpsSshHostKeyIndexInput {
	inv.Input = &ActionVpsSshHostKeyIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsSshHostKeyIndexInvocation) SetInput(input *ActionVpsSshHostKeyIndexInput) *ActionVpsSshHostKeyIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsSshHostKeyIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsSshHostKeyIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsSshHostKeyIndexInvocation) NewMetaInput() *ActionVpsSshHostKeyIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsSshHostKeyIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsSshHostKeyIndexInvocation) SetMetaInput(input *ActionVpsSshHostKeyIndexMetaGlobalInput) *ActionVpsSshHostKeyIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsSshHostKeyIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsSshHostKeyIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsSshHostKeyIndexInvocation) Call() (*ActionVpsSshHostKeyIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsSshHostKeyIndexInvocation) callAsQuery() (*ActionVpsSshHostKeyIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsSshHostKeyIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SshHostKeys
	}
	return resp, err
}

func (inv *ActionVpsSshHostKeyIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Algorithm") {
			ret["ssh_host_key[algorithm]"] = inv.Input.Algorithm
		}
		if inv.IsParameterSelected("FromId") {
			ret["ssh_host_key[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["ssh_host_key[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionVpsSshHostKeyIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
