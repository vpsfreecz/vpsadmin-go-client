package client

import (
	"strings"
)

// ActionVpsFeatureIndex is a type for action Vps.Feature#Index
type ActionVpsFeatureIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsFeatureIndex(client *Client) *ActionVpsFeatureIndex {
	return &ActionVpsFeatureIndex{
		Client: client,
	}
}

// ActionVpsFeatureIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsFeatureIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsFeatureIndexMetaGlobalInput) SetCount(value bool) *ActionVpsFeatureIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsFeatureIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsFeatureIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsFeatureIndexMetaGlobalInput) SetNo(value bool) *ActionVpsFeatureIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsFeatureIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsFeatureIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsFeatureIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsFeatureIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsFeatureIndexInput is a type for action input parameters
type ActionVpsFeatureIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionVpsFeatureIndexInput) SetFromId(value int64) *ActionVpsFeatureIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsFeatureIndexInput) SetLimit(value int64) *ActionVpsFeatureIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsFeatureIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsFeatureIndexInput) SelectParameters(params ...string) *ActionVpsFeatureIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsFeatureIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsFeatureIndexInput) UnselectParameters(params ...string) *ActionVpsFeatureIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsFeatureIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsFeatureIndexOutput is a type for action output parameters
type ActionVpsFeatureIndexOutput struct {
	Enabled bool   `json:"enabled"`
	Id      int64  `json:"id"`
	Label   string `json:"label"`
	Name    string `json:"name"`
}

// Type for action response, including envelope
type ActionVpsFeatureIndexResponse struct {
	Action *ActionVpsFeatureIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Features []*ActionVpsFeatureIndexOutput `json:"features"`
	}

	// Action output without the namespace
	Output []*ActionVpsFeatureIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsFeatureIndex) Prepare() *ActionVpsFeatureIndexInvocation {
	return &ActionVpsFeatureIndexInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/features",
	}
}

// ActionVpsFeatureIndexInvocation is used to configure action for invocation
type ActionVpsFeatureIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsFeatureIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsFeatureIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsFeatureIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsFeatureIndexInvocation) SetPathParamInt(param string, value int64) *ActionVpsFeatureIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsFeatureIndexInvocation) SetPathParamString(param string, value string) *ActionVpsFeatureIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsFeatureIndexInvocation) NewInput() *ActionVpsFeatureIndexInput {
	inv.Input = &ActionVpsFeatureIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsFeatureIndexInvocation) SetInput(input *ActionVpsFeatureIndexInput) *ActionVpsFeatureIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsFeatureIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsFeatureIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsFeatureIndexInvocation) NewMetaInput() *ActionVpsFeatureIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsFeatureIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsFeatureIndexInvocation) SetMetaInput(input *ActionVpsFeatureIndexMetaGlobalInput) *ActionVpsFeatureIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsFeatureIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsFeatureIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsFeatureIndexInvocation) Call() (*ActionVpsFeatureIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsFeatureIndexInvocation) callAsQuery() (*ActionVpsFeatureIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsFeatureIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Features
	}
	return resp, err
}

func (inv *ActionVpsFeatureIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["feature[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["feature[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionVpsFeatureIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
