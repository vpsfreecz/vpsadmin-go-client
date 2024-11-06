package client

import (
	"strings"
)

// ActionExportHostIndex is a type for action Export.Host#Index
type ActionExportHostIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionExportHostIndex(client *Client) *ActionExportHostIndex {
	return &ActionExportHostIndex{
		Client: client,
	}
}

// ActionExportHostIndexMetaGlobalInput is a type for action global meta input parameters
type ActionExportHostIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionExportHostIndexMetaGlobalInput) SetCount(value bool) *ActionExportHostIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportHostIndexMetaGlobalInput) SetIncludes(value string) *ActionExportHostIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportHostIndexMetaGlobalInput) SetNo(value bool) *ActionExportHostIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportHostIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportHostIndexMetaGlobalInput) SelectParameters(params ...string) *ActionExportHostIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportHostIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportHostIndexInput is a type for action input parameters
type ActionExportHostIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionExportHostIndexInput) SetFromId(value int64) *ActionExportHostIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionExportHostIndexInput) SetLimit(value int64) *ActionExportHostIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportHostIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportHostIndexInput) SelectParameters(params ...string) *ActionExportHostIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionExportHostIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionExportHostIndexInput) UnselectParameters(params ...string) *ActionExportHostIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionExportHostIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportHostIndexOutput is a type for action output parameters
type ActionExportHostIndexOutput struct {
	Id           int64                      `json:"id"`
	IpAddress    *ActionIpAddressShowOutput `json:"ip_address"`
	RootSquash   bool                       `json:"root_squash"`
	Rw           bool                       `json:"rw"`
	SubtreeCheck bool                       `json:"subtree_check"`
	Sync         bool                       `json:"sync"`
}

// Type for action response, including envelope
type ActionExportHostIndexResponse struct {
	Action *ActionExportHostIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Hosts []*ActionExportHostIndexOutput `json:"hosts"`
	}

	// Action output without the namespace
	Output []*ActionExportHostIndexOutput
}

// Prepare the action for invocation
func (action *ActionExportHostIndex) Prepare() *ActionExportHostIndexInvocation {
	return &ActionExportHostIndexInvocation{
		Action: action,
		Path:   "/v7.0/exports/{export_id}/hosts",
	}
}

// ActionExportHostIndexInvocation is used to configure action for invocation
type ActionExportHostIndexInvocation struct {
	// Pointer to the action
	Action *ActionExportHostIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionExportHostIndexInput
	// Global meta input parameters
	MetaInput *ActionExportHostIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionExportHostIndexInvocation) SetPathParamInt(param string, value int64) *ActionExportHostIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionExportHostIndexInvocation) SetPathParamString(param string, value string) *ActionExportHostIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionExportHostIndexInvocation) NewInput() *ActionExportHostIndexInput {
	inv.Input = &ActionExportHostIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionExportHostIndexInvocation) SetInput(input *ActionExportHostIndexInput) *ActionExportHostIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionExportHostIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionExportHostIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportHostIndexInvocation) NewMetaInput() *ActionExportHostIndexMetaGlobalInput {
	inv.MetaInput = &ActionExportHostIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportHostIndexInvocation) SetMetaInput(input *ActionExportHostIndexMetaGlobalInput) *ActionExportHostIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportHostIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionExportHostIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionExportHostIndexInvocation) Call() (*ActionExportHostIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionExportHostIndexInvocation) callAsQuery() (*ActionExportHostIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionExportHostIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Hosts
	}
	return resp, err
}

func (inv *ActionExportHostIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["host[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["host[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionExportHostIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
