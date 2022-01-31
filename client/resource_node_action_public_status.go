package client

import ()

// ActionNodePublicStatus is a type for action Node#Public_status
type ActionNodePublicStatus struct {
	// Pointer to client
	Client *Client
}

func NewActionNodePublicStatus(client *Client) *ActionNodePublicStatus {
	return &ActionNodePublicStatus{
		Client: client,
	}
}

// ActionNodePublicStatusMetaGlobalInput is a type for action global meta input parameters
type ActionNodePublicStatusMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNodePublicStatusMetaGlobalInput) SetIncludes(value string) *ActionNodePublicStatusMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNodePublicStatusMetaGlobalInput) SetNo(value bool) *ActionNodePublicStatusMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNodePublicStatusMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNodePublicStatusMetaGlobalInput) SelectParameters(params ...string) *ActionNodePublicStatusMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNodePublicStatusMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNodePublicStatusOutput is a type for action output parameters
type ActionNodePublicStatusOutput struct {
	CpuIdle               float64                   `json:"cpu_idle"`
	HypervisorType        string                    `json:"hypervisor_type"`
	Kernel                string                    `json:"kernel"`
	LastReport            string                    `json:"last_report"`
	Location              *ActionLocationShowOutput `json:"location"`
	MaintenanceLock       string                    `json:"maintenance_lock"`
	MaintenanceLockReason string                    `json:"maintenance_lock_reason"`
	Name                  string                    `json:"name"`
	Status                bool                      `json:"status"`
	VpsCount              int64                     `json:"vps_count"`
	VpsFree               int64                     `json:"vps_free"`
}

// Type for action response, including envelope
type ActionNodePublicStatusResponse struct {
	Action *ActionNodePublicStatus `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Nodes []*ActionNodePublicStatusOutput `json:"nodes"`
	}

	// Action output without the namespace
	Output []*ActionNodePublicStatusOutput
}

// Call the action directly without any path or input parameters
func (action *ActionNodePublicStatus) Call() (*ActionNodePublicStatusResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *ActionNodePublicStatus) Prepare() *ActionNodePublicStatusInvocation {
	return &ActionNodePublicStatusInvocation{
		Action: action,
		Path:   "/v6.0/nodes/public_status",
	}
}

// ActionNodePublicStatusInvocation is used to configure action for invocation
type ActionNodePublicStatusInvocation struct {
	// Pointer to the action
	Action *ActionNodePublicStatus

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNodePublicStatusMetaGlobalInput
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNodePublicStatusInvocation) NewMetaInput() *ActionNodePublicStatusMetaGlobalInput {
	inv.MetaInput = &ActionNodePublicStatusMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNodePublicStatusInvocation) SetMetaInput(input *ActionNodePublicStatusMetaGlobalInput) *ActionNodePublicStatusInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNodePublicStatusInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNodePublicStatusInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNodePublicStatusInvocation) Call() (*ActionNodePublicStatusResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNodePublicStatusInvocation) callAsQuery() (*ActionNodePublicStatusResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNodePublicStatusResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Nodes
	}
	return resp, err
}

func (inv *ActionNodePublicStatusInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
