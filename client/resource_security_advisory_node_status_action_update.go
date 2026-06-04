package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryNodeStatusUpdate is a type for action Security_advisory.Node_status#Update
type ActionSecurityAdvisoryNodeStatusUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryNodeStatusUpdate(client *Client) *ActionSecurityAdvisoryNodeStatusUpdate {
	return &ActionSecurityAdvisoryNodeStatusUpdate{
		Client: client,
	}
}

// ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryNodeStatusUpdateInput is a type for action input parameters
type ActionSecurityAdvisoryNodeStatusUpdateInput struct {
	MitigatedSince  string "json:\"mitigated_since\""
	Note            string "json:\"note\""
	State           string "json:\"state\""
	VulnerableUntil string "json:\"vulnerable_until\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetMitigatedSince sets parameter MitigatedSince to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) SetMitigatedSince(value string) *ActionSecurityAdvisoryNodeStatusUpdateInput {
	in.MitigatedSince = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetMitigatedSinceNil(false)
	in._selectedParameters["MitigatedSince"] = nil
	return in
}

// SetMitigatedSinceNil sets parameter MitigatedSince to nil and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) SetMitigatedSinceNil(set bool) *ActionSecurityAdvisoryNodeStatusUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["MitigatedSince"] = nil
		in.SelectParameters("MitigatedSince")
	} else {
		delete(in._nilParameters, "MitigatedSince")
	}
	return in
}

// SetNote sets parameter Note to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) SetNote(value string) *ActionSecurityAdvisoryNodeStatusUpdateInput {
	in.Note = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNoteNil(false)
	in._selectedParameters["Note"] = nil
	return in
}

// SetNoteNil sets parameter Note to nil and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) SetNoteNil(set bool) *ActionSecurityAdvisoryNodeStatusUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Note"] = nil
		in.SelectParameters("Note")
	} else {
		delete(in._nilParameters, "Note")
	}
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) SetState(value string) *ActionSecurityAdvisoryNodeStatusUpdateInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetVulnerableUntil sets parameter VulnerableUntil to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) SetVulnerableUntil(value string) *ActionSecurityAdvisoryNodeStatusUpdateInput {
	in.VulnerableUntil = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVulnerableUntilNil(false)
	in._selectedParameters["VulnerableUntil"] = nil
	return in
}

// SetVulnerableUntilNil sets parameter VulnerableUntil to nil and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) SetVulnerableUntilNil(set bool) *ActionSecurityAdvisoryNodeStatusUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["VulnerableUntil"] = nil
		in.SelectParameters("VulnerableUntil")
	} else {
		delete(in._nilParameters, "VulnerableUntil")
	}
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryNodeStatusUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) SelectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryNodeStatusUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryNodeStatusUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryNodeStatusUpdateRequest is a type for the entire action request
type ActionSecurityAdvisoryNodeStatusUpdateRequest struct {
	NodeStatus map[string]interface{} "json:\"node_status\""
	Meta       map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryNodeStatusUpdateOutput is a type for action output parameters
type ActionSecurityAdvisoryNodeStatusUpdateOutput struct {
	Id               int64                             "json:\"id\""
	MitigatedSince   string                            "json:\"mitigated_since\""
	NodeId           int64                             "json:\"node_id\""
	NodeName         string                            "json:\"node_name\""
	Note             string                            "json:\"note\""
	SecurityAdvisory *ActionSecurityAdvisoryShowOutput "json:\"security_advisory\""
	State            string                            "json:\"state\""
	VulnerableUntil  string                            "json:\"vulnerable_until\""
}

// Type for action response, including envelope
type ActionSecurityAdvisoryNodeStatusUpdateResponse struct {
	Action *ActionSecurityAdvisoryNodeStatusUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeStatus *ActionSecurityAdvisoryNodeStatusUpdateOutput "json:\"node_status\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryNodeStatusUpdateOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryNodeStatusUpdate) Prepare() *ActionSecurityAdvisoryNodeStatusUpdateInvocation {
	return &ActionSecurityAdvisoryNodeStatusUpdateInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}/node_statuses/{node_status_id}",
	}
}

// ActionSecurityAdvisoryNodeStatusUpdateInvocation is used to configure action for invocation
type ActionSecurityAdvisoryNodeStatusUpdateInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryNodeStatusUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryNodeStatusUpdateInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryNodeStatusUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryNodeStatusUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) NewInput() *ActionSecurityAdvisoryNodeStatusUpdateInput {
	inv.Input = &ActionSecurityAdvisoryNodeStatusUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) SetInput(input *ActionSecurityAdvisoryNodeStatusUpdateInput) *ActionSecurityAdvisoryNodeStatusUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) NewMetaInput() *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) SetMetaInput(input *ActionSecurityAdvisoryNodeStatusUpdateMetaGlobalInput) *ActionSecurityAdvisoryNodeStatusUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("MitigatedSince") {
			if !inv.IsParameterNil("MitigatedSince") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.MitigatedSince)
				if !ok {
					verr.Add("mitigated_since", "not a valid datetime")
				} else {
					inv.Input.MitigatedSince = normalized
				}
			}
		}
		if inv.IsParameterSelected("VulnerableUntil") {
			if !inv.IsParameterNil("VulnerableUntil") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.VulnerableUntil)
				if !ok {
					verr.Add("vulnerable_until", "not a valid datetime")
				} else {
					inv.Input.VulnerableUntil = normalized
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) Call() (*ActionSecurityAdvisoryNodeStatusUpdateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) callAsBody() (*ActionSecurityAdvisoryNodeStatusUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryNodeStatusUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeStatus
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) makeAllInputParams() *ActionSecurityAdvisoryNodeStatusUpdateRequest {
	return &ActionSecurityAdvisoryNodeStatusUpdateRequest{
		NodeStatus: inv.makeInputParams(),
		Meta:       inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("MitigatedSince") {
			if inv.IsParameterNil("MitigatedSince") {
				ret["mitigated_since"] = nil
			} else {
				ret["mitigated_since"] = inv.Input.MitigatedSince
			}
		}
		if inv.IsParameterSelected("Note") {
			if inv.IsParameterNil("Note") {
				ret["note"] = nil
			} else {
				ret["note"] = inv.Input.Note
			}
		}
		if inv.IsParameterSelected("State") {
			ret["state"] = inv.Input.State
		}
		if inv.IsParameterSelected("VulnerableUntil") {
			if inv.IsParameterNil("VulnerableUntil") {
				ret["vulnerable_until"] = nil
			} else {
				ret["vulnerable_until"] = inv.Input.VulnerableUntil
			}
		}
	}

	return ret
}

func (inv *ActionSecurityAdvisoryNodeStatusUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
