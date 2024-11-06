package client

import (
	"strings"
)

// ActionUserKnownDeviceShow is a type for action User.Known_device#Show
type ActionUserKnownDeviceShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserKnownDeviceShow(client *Client) *ActionUserKnownDeviceShow {
	return &ActionUserKnownDeviceShow{
		Client: client,
	}
}

// ActionUserKnownDeviceShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserKnownDeviceShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserKnownDeviceShowMetaGlobalInput) SetIncludes(value string) *ActionUserKnownDeviceShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserKnownDeviceShowMetaGlobalInput) SetNo(value bool) *ActionUserKnownDeviceShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserKnownDeviceShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserKnownDeviceShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserKnownDeviceShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserKnownDeviceShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserKnownDeviceShowOutput is a type for action output parameters
type ActionUserKnownDeviceShowOutput struct {
	ApiIpAddr                string `json:"api_ip_addr"`
	ApiIpPtr                 string `json:"api_ip_ptr"`
	ClientIpAddr             string `json:"client_ip_addr"`
	ClientIpPtr              string `json:"client_ip_ptr"`
	CreatedAt                string `json:"created_at"`
	Id                       int64  `json:"id"`
	LastSeenAt               string `json:"last_seen_at"`
	SkipMultiFactorAuthUntil string `json:"skip_multi_factor_auth_until"`
	UpdatedAt                string `json:"updated_at"`
	UserAgent                string `json:"user_agent"`
}

// Type for action response, including envelope
type ActionUserKnownDeviceShowResponse struct {
	Action *ActionUserKnownDeviceShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		KnownDevice *ActionUserKnownDeviceShowOutput `json:"known_device"`
	}

	// Action output without the namespace
	Output *ActionUserKnownDeviceShowOutput
}

// Prepare the action for invocation
func (action *ActionUserKnownDeviceShow) Prepare() *ActionUserKnownDeviceShowInvocation {
	return &ActionUserKnownDeviceShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/known_devices/{known_device_id}",
	}
}

// ActionUserKnownDeviceShowInvocation is used to configure action for invocation
type ActionUserKnownDeviceShowInvocation struct {
	// Pointer to the action
	Action *ActionUserKnownDeviceShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserKnownDeviceShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserKnownDeviceShowInvocation) SetPathParamInt(param string, value int64) *ActionUserKnownDeviceShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserKnownDeviceShowInvocation) SetPathParamString(param string, value string) *ActionUserKnownDeviceShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserKnownDeviceShowInvocation) NewMetaInput() *ActionUserKnownDeviceShowMetaGlobalInput {
	inv.MetaInput = &ActionUserKnownDeviceShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserKnownDeviceShowInvocation) SetMetaInput(input *ActionUserKnownDeviceShowMetaGlobalInput) *ActionUserKnownDeviceShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserKnownDeviceShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserKnownDeviceShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserKnownDeviceShowInvocation) Call() (*ActionUserKnownDeviceShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserKnownDeviceShowInvocation) callAsQuery() (*ActionUserKnownDeviceShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserKnownDeviceShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.KnownDevice
	}
	return resp, err
}

func (inv *ActionUserKnownDeviceShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
