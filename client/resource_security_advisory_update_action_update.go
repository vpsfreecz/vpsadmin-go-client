package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryUpdateUpdate is a type for action Security_advisory_update#Update
type ActionSecurityAdvisoryUpdateUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryUpdateUpdate(client *Client) *ActionSecurityAdvisoryUpdateUpdate {
	return &ActionSecurityAdvisoryUpdateUpdate{
		Client: client,
	}
}

// ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateUpdateInput is a type for action input parameters
type ActionSecurityAdvisoryUpdateUpdateInput struct {
	CsMessage string "json:\"cs_message\""
	CsSummary string "json:\"cs_summary\""
	EnMessage string "json:\"en_message\""
	EnSummary string "json:\"en_summary\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCsMessage sets parameter CsMessage to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateUpdateInput) SetCsMessage(value string) *ActionSecurityAdvisoryUpdateUpdateInput {
	in.CsMessage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetCsMessageNil(false)
	in._selectedParameters["CsMessage"] = nil
	return in
}

// SetCsMessageNil sets parameter CsMessage to nil and selects it for sending
func (in *ActionSecurityAdvisoryUpdateUpdateInput) SetCsMessageNil(set bool) *ActionSecurityAdvisoryUpdateUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["CsMessage"] = nil
		in.SelectParameters("CsMessage")
	} else {
		delete(in._nilParameters, "CsMessage")
	}
	return in
}

// SetCsSummary sets parameter CsSummary to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateUpdateInput) SetCsSummary(value string) *ActionSecurityAdvisoryUpdateUpdateInput {
	in.CsSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsSummary"] = nil
	return in
}

// SetEnMessage sets parameter EnMessage to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateUpdateInput) SetEnMessage(value string) *ActionSecurityAdvisoryUpdateUpdateInput {
	in.EnMessage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEnMessageNil(false)
	in._selectedParameters["EnMessage"] = nil
	return in
}

// SetEnMessageNil sets parameter EnMessage to nil and selects it for sending
func (in *ActionSecurityAdvisoryUpdateUpdateInput) SetEnMessageNil(set bool) *ActionSecurityAdvisoryUpdateUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["EnMessage"] = nil
		in.SelectParameters("EnMessage")
	} else {
		delete(in._nilParameters, "EnMessage")
	}
	return in
}

// SetEnSummary sets parameter EnSummary to value and selects it for sending
func (in *ActionSecurityAdvisoryUpdateUpdateInput) SetEnSummary(value string) *ActionSecurityAdvisoryUpdateUpdateInput {
	in.EnSummary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnSummary"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryUpdateUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateUpdateInput) SelectParameters(params ...string) *ActionSecurityAdvisoryUpdateUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryUpdateUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryUpdateUpdateInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryUpdateUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryUpdateUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryUpdateUpdateRequest is a type for the entire action request
type ActionSecurityAdvisoryUpdateUpdateRequest struct {
	SecurityAdvisoryUpdate map[string]interface{} "json:\"security_advisory_update\""
	Meta                   map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryUpdateUpdateOutput is a type for action output parameters
type ActionSecurityAdvisoryUpdateUpdateOutput struct {
	CreatedAt        string                            "json:\"created_at\""
	CsMessage        string                            "json:\"cs_message\""
	CsSummary        string                            "json:\"cs_summary\""
	EnMessage        string                            "json:\"en_message\""
	EnSummary        string                            "json:\"en_summary\""
	Id               int64                             "json:\"id\""
	Name             string                            "json:\"name\""
	ReportedBy       *ActionUserShowOutput             "json:\"reported_by\""
	ReporterName     string                            "json:\"reporter_name\""
	SecurityAdvisory *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	State            string                            "json:\"state\""
	UpdatedAt        string                            "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryUpdateUpdateResponse struct {
	Action *ActionSecurityAdvisoryUpdateUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SecurityAdvisoryUpdate *ActionSecurityAdvisoryUpdateUpdateOutput "json:\"security_advisory_update\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryUpdateUpdateOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryUpdateUpdate) Prepare() *ActionSecurityAdvisoryUpdateUpdateInvocation {
	return &ActionSecurityAdvisoryUpdateUpdateInvocation{
		Action: action,
		Path:   "/v7.0/security_advisory_updates/{security_advisory_update_id}",
	}
}

// ActionSecurityAdvisoryUpdateUpdateInvocation is used to configure action for invocation
type ActionSecurityAdvisoryUpdateUpdateInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryUpdateUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryUpdateUpdateInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryUpdateUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryUpdateUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) NewInput() *ActionSecurityAdvisoryUpdateUpdateInput {
	inv.Input = &ActionSecurityAdvisoryUpdateUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) SetInput(input *ActionSecurityAdvisoryUpdateUpdateInput) *ActionSecurityAdvisoryUpdateUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) NewMetaInput() *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) SetMetaInput(input *ActionSecurityAdvisoryUpdateUpdateMetaGlobalInput) *ActionSecurityAdvisoryUpdateUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) Call() (*ActionSecurityAdvisoryUpdateUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) callAsBody() (*ActionSecurityAdvisoryUpdateUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryUpdateUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.SecurityAdvisoryUpdate
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) makeAllInputParams() *ActionSecurityAdvisoryUpdateUpdateRequest {
	return &ActionSecurityAdvisoryUpdateUpdateRequest{
		SecurityAdvisoryUpdate: inv.makeInputParams(),
		Meta:                   inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("CsMessage") {
			if inv.IsParameterNil("CsMessage") {
				ret["cs_message"] = nil
			} else {
				ret["cs_message"] = inv.Input.CsMessage
			}
		}
		if inv.IsParameterSelected("CsSummary") {
			ret["cs_summary"] = inv.Input.CsSummary
		}
		if inv.IsParameterSelected("EnMessage") {
			if inv.IsParameterNil("EnMessage") {
				ret["en_message"] = nil
			} else {
				ret["en_message"] = inv.Input.EnMessage
			}
		}
		if inv.IsParameterSelected("EnSummary") {
			ret["en_summary"] = inv.Input.EnSummary
		}
	}

	return ret
}

func (inv *ActionSecurityAdvisoryUpdateUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
