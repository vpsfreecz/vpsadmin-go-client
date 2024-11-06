package client

import (
	"strings"
)

// ActionIpAddressAssignmentShow is a type for action Ip_address_assignment#Show
type ActionIpAddressAssignmentShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressAssignmentShow(client *Client) *ActionIpAddressAssignmentShow {
	return &ActionIpAddressAssignmentShow{
		Client: client,
	}
}

// ActionIpAddressAssignmentShowMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressAssignmentShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressAssignmentShowMetaGlobalInput) SetIncludes(value string) *ActionIpAddressAssignmentShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressAssignmentShowMetaGlobalInput) SetNo(value bool) *ActionIpAddressAssignmentShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressAssignmentShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressAssignmentShowMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressAssignmentShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressAssignmentShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressAssignmentShowOutput is a type for action output parameters
type ActionIpAddressAssignmentShowOutput struct {
	AssignedByChain   *ActionTransactionChainShowOutput `json:"assigned_by_chain"`
	CreatedAt         string                            `json:"created_at"`
	FromDate          string                            `json:"from_date"`
	Id                int64                             `json:"id"`
	IpAddr            string                            `json:"ip_addr"`
	IpAddress         *ActionIpAddressShowOutput        `json:"ip_address"`
	IpPrefix          int64                             `json:"ip_prefix"`
	RawUserId         int64                             `json:"raw_user_id"`
	RawVpsId          int64                             `json:"raw_vps_id"`
	Reconstructed     bool                              `json:"reconstructed"`
	ToDate            string                            `json:"to_date"`
	UnassignedByChain *ActionTransactionChainShowOutput `json:"unassigned_by_chain"`
	UpdatedAt         string                            `json:"updated_at"`
	User              *ActionUserShowOutput             `json:"user"`
	Vps               *ActionVpsShowOutput              `json:"vps"`
}

// Type for action response, including envelope
type ActionIpAddressAssignmentShowResponse struct {
	Action *ActionIpAddressAssignmentShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddressAssignment *ActionIpAddressAssignmentShowOutput `json:"ip_address_assignment"`
	}

	// Action output without the namespace
	Output *ActionIpAddressAssignmentShowOutput
}

// Prepare the action for invocation
func (action *ActionIpAddressAssignmentShow) Prepare() *ActionIpAddressAssignmentShowInvocation {
	return &ActionIpAddressAssignmentShowInvocation{
		Action: action,
		Path:   "/v7.0/ip_address_assignments/{ip_address_assignment_id}",
	}
}

// ActionIpAddressAssignmentShowInvocation is used to configure action for invocation
type ActionIpAddressAssignmentShowInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressAssignmentShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpAddressAssignmentShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpAddressAssignmentShowInvocation) SetPathParamInt(param string, value int64) *ActionIpAddressAssignmentShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpAddressAssignmentShowInvocation) SetPathParamString(param string, value string) *ActionIpAddressAssignmentShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpAddressAssignmentShowInvocation) NewMetaInput() *ActionIpAddressAssignmentShowMetaGlobalInput {
	inv.MetaInput = &ActionIpAddressAssignmentShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressAssignmentShowInvocation) SetMetaInput(input *ActionIpAddressAssignmentShowMetaGlobalInput) *ActionIpAddressAssignmentShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressAssignmentShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpAddressAssignmentShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressAssignmentShowInvocation) Call() (*ActionIpAddressAssignmentShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIpAddressAssignmentShowInvocation) callAsQuery() (*ActionIpAddressAssignmentShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpAddressAssignmentShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpAddressAssignment
	}
	return resp, err
}

func (inv *ActionIpAddressAssignmentShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
