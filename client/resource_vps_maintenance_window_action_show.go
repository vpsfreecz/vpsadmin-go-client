package client

import (
	"strings"
)

// ActionVpsMaintenanceWindowShow is a type for action Vps.Maintenance_window#Show
type ActionVpsMaintenanceWindowShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsMaintenanceWindowShow(client *Client) *ActionVpsMaintenanceWindowShow {
	return &ActionVpsMaintenanceWindowShow{
		Client: client,
	}
}

// ActionVpsMaintenanceWindowShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsMaintenanceWindowShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsMaintenanceWindowShowMetaGlobalInput) SetIncludes(value string) *ActionVpsMaintenanceWindowShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsMaintenanceWindowShowMetaGlobalInput) SetNo(value bool) *ActionVpsMaintenanceWindowShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsMaintenanceWindowShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsMaintenanceWindowShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsMaintenanceWindowShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsMaintenanceWindowShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsMaintenanceWindowShowOutput is a type for action output parameters
type ActionVpsMaintenanceWindowShowOutput struct {
	ClosesAt int64 `json:"closes_at"`
	IsOpen   bool  `json:"is_open"`
	OpensAt  int64 `json:"opens_at"`
	Weekday  int64 `json:"weekday"`
}

// Type for action response, including envelope
type ActionVpsMaintenanceWindowShowResponse struct {
	Action *ActionVpsMaintenanceWindowShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MaintenanceWindow *ActionVpsMaintenanceWindowShowOutput `json:"maintenance_window"`
	}

	// Action output without the namespace
	Output *ActionVpsMaintenanceWindowShowOutput
}

// Prepare the action for invocation
func (action *ActionVpsMaintenanceWindowShow) Prepare() *ActionVpsMaintenanceWindowShowInvocation {
	return &ActionVpsMaintenanceWindowShowInvocation{
		Action: action,
		Path:   "/v6.0/vpses/{vps_id}/maintenance_windows/{maintenance_window_id}",
	}
}

// ActionVpsMaintenanceWindowShowInvocation is used to configure action for invocation
type ActionVpsMaintenanceWindowShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsMaintenanceWindowShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsMaintenanceWindowShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsMaintenanceWindowShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsMaintenanceWindowShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsMaintenanceWindowShowInvocation) SetPathParamString(param string, value string) *ActionVpsMaintenanceWindowShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsMaintenanceWindowShowInvocation) NewMetaInput() *ActionVpsMaintenanceWindowShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsMaintenanceWindowShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsMaintenanceWindowShowInvocation) SetMetaInput(input *ActionVpsMaintenanceWindowShowMetaGlobalInput) *ActionVpsMaintenanceWindowShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsMaintenanceWindowShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsMaintenanceWindowShowInvocation) Call() (*ActionVpsMaintenanceWindowShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsMaintenanceWindowShowInvocation) callAsQuery() (*ActionVpsMaintenanceWindowShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsMaintenanceWindowShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MaintenanceWindow
	}
	return resp, err
}

func (inv *ActionVpsMaintenanceWindowShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
