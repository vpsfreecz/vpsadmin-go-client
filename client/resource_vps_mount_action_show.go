package client

import (
	"strings"
)

// ActionVpsMountShow is a type for action Vps.Mount#Show
type ActionVpsMountShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMountShow(client *Client) *ActionVpsMountShow {
	return &ActionVpsMountShow{
		Client: client,
	}
}

// ActionVpsMountShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMountShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMountShowMetaGlobalInput) SetIncludes(value string) *ActionVpsMountShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMountShowMetaGlobalInput) SetNo(value bool) *ActionVpsMountShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMountShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMountShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMountShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMountShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMountShowOutput is a type for action output parameters
type ActionVpsMountShowOutput struct {
	CurrentState     string                            `json:"current_state"`
	Dataset          *ActionDatasetShowOutput          `json:"dataset"`
	Enabled          bool                              `json:"enabled"`
	ExpirationDate   string                            `json:"expiration_date"`
	Id               int64                             `json:"id"`
	MasterEnabled    bool                              `json:"master_enabled"`
	Mode             string                            `json:"mode"`
	Mountpoint       string                            `json:"mountpoint"`
	OnStartFail      string                            `json:"on_start_fail"`
	UserNamespaceMap *ActionUserNamespaceMapShowOutput `json:"user_namespace_map"`
	Vps              *ActionVpsShowOutput              `json:"vps"`
}

// Type for action response, including envelope
type ActionVpsMountShowResponse struct {
	Action *ActionVpsMountShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Mount *ActionVpsMountShowOutput `json:"mount"`
	}

	// Action output without the namespace
	Output *ActionVpsMountShowOutput
}

// Prepare the action for invocation
func (action *ActionVpsMountShow) Prepare() *ActionVpsMountShowInvocation {
	return &ActionVpsMountShowInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/mounts/{mount_id}",
	}
}

// ActionVpsMountShowInvocation is used to configure action for invocation
type ActionVpsMountShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsMountShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsMountShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMountShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsMountShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMountShowInvocation) SetPathParamString(param string, value string) *ActionVpsMountShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMountShowInvocation) NewMetaInput() *ActionVpsMountShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsMountShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMountShowInvocation) SetMetaInput(input *ActionVpsMountShowMetaGlobalInput) *ActionVpsMountShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMountShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsMountShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMountShowInvocation) Call() (*ActionVpsMountShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsMountShowInvocation) callAsQuery() (*ActionVpsMountShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsMountShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Mount
	}
	return resp, err
}

func (inv *ActionVpsMountShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
