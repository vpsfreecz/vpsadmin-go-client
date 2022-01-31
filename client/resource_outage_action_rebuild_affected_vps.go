package client

import (
	"strings"
)

// ActionOutageRebuildAffectedVps is a type for action Outage#Rebuild_affected_vps
type ActionOutageRebuildAffectedVps struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageRebuildAffectedVps(client *Client) *ActionOutageRebuildAffectedVps {
	return &ActionOutageRebuildAffectedVps{
		Client: client,
	}
}

// ActionOutageRebuildAffectedVpsMetaGlobalInput is a type for action global meta input parameters
type ActionOutageRebuildAffectedVpsMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageRebuildAffectedVpsMetaGlobalInput) SetIncludes(value string) *ActionOutageRebuildAffectedVpsMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageRebuildAffectedVpsMetaGlobalInput) SetNo(value bool) *ActionOutageRebuildAffectedVpsMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageRebuildAffectedVpsMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageRebuildAffectedVpsMetaGlobalInput) SelectParameters(params ...string) *ActionOutageRebuildAffectedVpsMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageRebuildAffectedVpsMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageRebuildAffectedVpsRequest is a type for the entire action request
type ActionOutageRebuildAffectedVpsRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionOutageRebuildAffectedVpsResponse struct {
	Action *ActionOutageRebuildAffectedVps `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionOutageRebuildAffectedVps) Prepare() *ActionOutageRebuildAffectedVpsInvocation {
	return &ActionOutageRebuildAffectedVpsInvocation{
		Action: action,
		Path:   "/v6.0/outages/{outage_id}/rebuild_affected_vps",
	}
}

// ActionOutageRebuildAffectedVpsInvocation is used to configure action for invocation
type ActionOutageRebuildAffectedVpsInvocation struct {
	// Pointer to the action
	Action *ActionOutageRebuildAffectedVps

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOutageRebuildAffectedVpsMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageRebuildAffectedVpsInvocation) SetPathParamInt(param string, value int64) *ActionOutageRebuildAffectedVpsInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageRebuildAffectedVpsInvocation) SetPathParamString(param string, value string) *ActionOutageRebuildAffectedVpsInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageRebuildAffectedVpsInvocation) NewMetaInput() *ActionOutageRebuildAffectedVpsMetaGlobalInput {
	inv.MetaInput = &ActionOutageRebuildAffectedVpsMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageRebuildAffectedVpsInvocation) SetMetaInput(input *ActionOutageRebuildAffectedVpsMetaGlobalInput) *ActionOutageRebuildAffectedVpsInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageRebuildAffectedVpsInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOutageRebuildAffectedVpsInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageRebuildAffectedVpsInvocation) Call() (*ActionOutageRebuildAffectedVpsResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOutageRebuildAffectedVpsInvocation) callAsBody() (*ActionOutageRebuildAffectedVpsResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOutageRebuildAffectedVpsResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionOutageRebuildAffectedVpsInvocation) makeAllInputParams() *ActionOutageRebuildAffectedVpsRequest {
	return &ActionOutageRebuildAffectedVpsRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOutageRebuildAffectedVpsInvocation) makeMetaInputParams() map[string]interface{} {
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
