package client

import (
	"net/url"
	"strings"
)

// ActionSecurityAdvisoryNodeStatusCreate is a type for action Security_advisory.Node_status#Create
type ActionSecurityAdvisoryNodeStatusCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionSecurityAdvisoryNodeStatusCreate(client *Client) *ActionSecurityAdvisoryNodeStatusCreate {
	return &ActionSecurityAdvisoryNodeStatusCreate{
		Client: client,
	}
}

// ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput is a type for action global meta input parameters
type ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput) SetIncludes(value string) *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput) SetNo(value bool) *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput) SelectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryNodeStatusCreateInput is a type for action input parameters
type ActionSecurityAdvisoryNodeStatusCreateInput struct {
	MitigatedSince  string "json:\"mitigated_since\""
	Node            int64  "json:\"node\""
	Note            string "json:\"note\""
	State           string "json:\"state\""
	VulnerableUntil string "json:\"vulnerable_until\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetMitigatedSince sets parameter MitigatedSince to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) SetMitigatedSince(value string) *ActionSecurityAdvisoryNodeStatusCreateInput {
	in.MitigatedSince = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetMitigatedSinceNil(false)
	in._selectedParameters["MitigatedSince"] = nil
	return in
}

// SetMitigatedSinceNil sets parameter MitigatedSince to nil and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) SetMitigatedSinceNil(set bool) *ActionSecurityAdvisoryNodeStatusCreateInput {
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

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) SetNode(value int64) *ActionSecurityAdvisoryNodeStatusCreateInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetNote sets parameter Note to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) SetNote(value string) *ActionSecurityAdvisoryNodeStatusCreateInput {
	in.Note = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNoteNil(false)
	in._selectedParameters["Note"] = nil
	return in
}

// SetNoteNil sets parameter Note to nil and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) SetNoteNil(set bool) *ActionSecurityAdvisoryNodeStatusCreateInput {
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
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) SetState(value string) *ActionSecurityAdvisoryNodeStatusCreateInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetVulnerableUntil sets parameter VulnerableUntil to value and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) SetVulnerableUntil(value string) *ActionSecurityAdvisoryNodeStatusCreateInput {
	in.VulnerableUntil = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVulnerableUntilNil(false)
	in._selectedParameters["VulnerableUntil"] = nil
	return in
}

// SetVulnerableUntilNil sets parameter VulnerableUntil to nil and selects it for sending
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) SetVulnerableUntilNil(set bool) *ActionSecurityAdvisoryNodeStatusCreateInput {
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

// SelectParameters sets parameters from ActionSecurityAdvisoryNodeStatusCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) SelectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionSecurityAdvisoryNodeStatusCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionSecurityAdvisoryNodeStatusCreateInput) UnselectParameters(params ...string) *ActionSecurityAdvisoryNodeStatusCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionSecurityAdvisoryNodeStatusCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionSecurityAdvisoryNodeStatusCreateRequest is a type for the entire action request
type ActionSecurityAdvisoryNodeStatusCreateRequest struct {
	NodeStatus map[string]interface{} "json:\"node_status\""
	Meta       map[string]interface{} "json:\"_meta\""
}

// ActionSecurityAdvisoryNodeStatusCreateOutput is a type for action output parameters
type ActionSecurityAdvisoryNodeStatusCreateOutput struct {
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
type ActionSecurityAdvisoryNodeStatusCreateResponse struct {
	Action *ActionSecurityAdvisoryNodeStatusCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NodeStatus *ActionSecurityAdvisoryNodeStatusCreateOutput "json:\"node_status\""
	}

	// Action output without the namespace
	Output *ActionSecurityAdvisoryNodeStatusCreateOutput
}

// Prepare the action for invocation
func (action *ActionSecurityAdvisoryNodeStatusCreate) Prepare() *ActionSecurityAdvisoryNodeStatusCreateInvocation {
	return &ActionSecurityAdvisoryNodeStatusCreateInvocation{
		Action: action,
		Path:   "/v7.0/security_advisories/{security_advisory_id}/node_statuses",
	}
}

// ActionSecurityAdvisoryNodeStatusCreateInvocation is used to configure action for invocation
type ActionSecurityAdvisoryNodeStatusCreateInvocation struct {
	// Pointer to the action
	Action *ActionSecurityAdvisoryNodeStatusCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionSecurityAdvisoryNodeStatusCreateInput
	// Global meta input parameters
	MetaInput *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) SetPathParamInt(param string, value int64) *ActionSecurityAdvisoryNodeStatusCreateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) SetPathParamString(param string, value string) *ActionSecurityAdvisoryNodeStatusCreateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) NewInput() *ActionSecurityAdvisoryNodeStatusCreateInput {
	inv.Input = &ActionSecurityAdvisoryNodeStatusCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) SetInput(input *ActionSecurityAdvisoryNodeStatusCreateInput) *ActionSecurityAdvisoryNodeStatusCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) NewMetaInput() *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput {
	inv.MetaInput = &ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) SetMetaInput(input *ActionSecurityAdvisoryNodeStatusCreateMetaGlobalInput) *ActionSecurityAdvisoryNodeStatusCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) validate() error {
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
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
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
func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) Call() (*ActionSecurityAdvisoryNodeStatusCreateResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) callAsBody() (*ActionSecurityAdvisoryNodeStatusCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionSecurityAdvisoryNodeStatusCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NodeStatus
	}
	return resp, err
}

func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) makeAllInputParams() *ActionSecurityAdvisoryNodeStatusCreateRequest {
	return &ActionSecurityAdvisoryNodeStatusCreateRequest{
		NodeStatus: inv.makeInputParams(),
		Meta:       inv.makeMetaInputParams(),
	}
}

func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("MitigatedSince") {
			if inv.IsParameterNil("MitigatedSince") {
				ret["mitigated_since"] = nil
			} else {
				ret["mitigated_since"] = inv.Input.MitigatedSince
			}
		}
		if inv.IsParameterSelected("Node") {
			ret["node"] = inv.Input.Node
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

func (inv *ActionSecurityAdvisoryNodeStatusCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
