package client

import (
	"strings"
)

// ActionVpsStateLogShow is a type for action Vps.State_log#Show
type ActionVpsStateLogShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsStateLogShow(client *Client) *ActionVpsStateLogShow {
	return &ActionVpsStateLogShow{
		Client: client,
	}
}

// ActionVpsStateLogShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsStateLogShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsStateLogShowMetaGlobalInput) SetIncludes(value string) *ActionVpsStateLogShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsStateLogShowMetaGlobalInput) SetNo(value bool) *ActionVpsStateLogShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsStateLogShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsStateLogShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsStateLogShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsStateLogShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsStateLogShowOutput is a type for action output parameters
type ActionVpsStateLogShowOutput struct {
	ChangedAt   string                `json:"changed_at"`
	Expiration  string                `json:"expiration"`
	Id          int64                 `json:"id"`
	Reason      string                `json:"reason"`
	RemindAfter string                `json:"remind_after"`
	State       string                `json:"state"`
	User        *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionVpsStateLogShowResponse struct {
	Action *ActionVpsStateLogShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		StateLog *ActionVpsStateLogShowOutput `json:"state_log"`
	}

	// Action output without the namespace
	Output *ActionVpsStateLogShowOutput
}

// Prepare the action for invocation
func (action *ActionVpsStateLogShow) Prepare() *ActionVpsStateLogShowInvocation {
	return &ActionVpsStateLogShowInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/state_logs/{state_log_id}",
	}
}

// ActionVpsStateLogShowInvocation is used to configure action for invocation
type ActionVpsStateLogShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsStateLogShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsStateLogShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsStateLogShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsStateLogShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsStateLogShowInvocation) SetPathParamString(param string, value string) *ActionVpsStateLogShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsStateLogShowInvocation) NewMetaInput() *ActionVpsStateLogShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsStateLogShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsStateLogShowInvocation) SetMetaInput(input *ActionVpsStateLogShowMetaGlobalInput) *ActionVpsStateLogShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsStateLogShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsStateLogShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsStateLogShowInvocation) Call() (*ActionVpsStateLogShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsStateLogShowInvocation) callAsQuery() (*ActionVpsStateLogShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsStateLogShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.StateLog
	}
	return resp, err
}

func (inv *ActionVpsStateLogShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
