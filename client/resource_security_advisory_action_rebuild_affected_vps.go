package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryRebuildAffectedVps is a type for action Security_advisory#Rebuild_affected_vps
type ActionSecurityAdvisoryRebuildAffectedVps struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryRebuildAffectedVps(client *Client) *ActionSecurityAdvisoryRebuildAffectedVps {
	return &ActionSecurityAdvisoryRebuildAffectedVps{
		Client: client,
	}
}

// ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryRebuildAffectedVpsRequest is a type for the entire action request
type ActionSecurityAdvisoryRebuildAffectedVpsRequest struct {
	Meta map[string]interface{} "json:\"_meta\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryRebuildAffectedVpsResponse struct {
	Action *ActionSecurityAdvisoryRebuildAffectedVps "json:\"-\""
	*Envelope
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryRebuildAffectedVps) Prepare() *ActionSecurityAdvisoryRebuildAffectedVpsInvocation {
	return &ActionSecurityAdvisoryRebuildAffectedVpsInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}/rebuild_affected_vps",
	}
}

// ActionSecurityAdvisoryRebuildAffectedVpsInvocation is used to configure action for invocation
type ActionSecurityAdvisoryRebuildAffectedVpsInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryRebuildAffectedVps

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryRebuildAffectedVpsInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryRebuildAffectedVpsInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) NewMetaInput() *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) SetMetaInput(input *ActionSecurityAdvisoryRebuildAffectedVpsMetaGlobalInput) *ActionSecurityAdvisoryRebuildAffectedVpsInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) Call() (*ActionSecurityAdvisoryRebuildAffectedVpsResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) callAsBody() (*ActionSecurityAdvisoryRebuildAffectedVpsResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryRebuildAffectedVpsResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) makeAllInputParams() *ActionSecurityAdvisoryRebuildAffectedVpsRequest {
	return &ActionSecurityAdvisoryRebuildAffectedVpsRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryRebuildAffectedVpsInvocation) makeMetaInputParams() map[string]interface{} {
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
