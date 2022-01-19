package client

import ()

// ActionLocationIndex is a type for action Location#Index
type ActionLocationIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionLocationIndex(client *Client) *ActionLocationIndex {
	return &ActionLocationIndex{
		Client: client,
	}
}

// ActionLocationIndexMetaGlobalInput is a type for action global meta input parameters
type ActionLocationIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionLocationIndexMetaGlobalInput) SetCount(value bool) *ActionLocationIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLocationIndexMetaGlobalInput) SetIncludes(value string) *ActionLocationIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLocationIndexMetaGlobalInput) SetNo(value bool) *ActionLocationIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationIndexMetaGlobalInput) SelectParameters(params ...string) *ActionLocationIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationIndexInput is a type for action input parameters
type ActionLocationIndexInput struct {
	Environment           int64  `json:"environment"`
	HasHypervisor         bool   `json:"has_hypervisor"`
	HasStorage            bool   `json:"has_storage"`
	HypervisorType        string `json:"hypervisor_type"`
	Limit                 int64  `json:"limit"`
	Offset                int64  `json:"offset"`
	SharesAnyNetworksWith int64  `json:"shares_any_networks_with"`
	SharesNetworksPrimary bool   `json:"shares_networks_primary"`
	SharesV4NetworksWith  int64  `json:"shares_v4_networks_with"`
	SharesV6NetworksWith  int64  `json:"shares_v6_networks_with"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionLocationIndexInput) SetEnvironment(value int64) *ActionLocationIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}

// SetHasHypervisor sets parameter HasHypervisor to value and selects it for sending
func (in *ActionLocationIndexInput) SetHasHypervisor(value bool) *ActionLocationIndexInput {
	in.HasHypervisor = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HasHypervisor"] = nil
	return in
}

// SetHasStorage sets parameter HasStorage to value and selects it for sending
func (in *ActionLocationIndexInput) SetHasStorage(value bool) *ActionLocationIndexInput {
	in.HasStorage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HasStorage"] = nil
	return in
}

// SetHypervisorType sets parameter HypervisorType to value and selects it for sending
func (in *ActionLocationIndexInput) SetHypervisorType(value string) *ActionLocationIndexInput {
	in.HypervisorType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HypervisorType"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionLocationIndexInput) SetLimit(value int64) *ActionLocationIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionLocationIndexInput) SetOffset(value int64) *ActionLocationIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetSharesAnyNetworksWith sets parameter SharesAnyNetworksWith to value and selects it for sending
func (in *ActionLocationIndexInput) SetSharesAnyNetworksWith(value int64) *ActionLocationIndexInput {
	in.SharesAnyNetworksWith = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SharesAnyNetworksWith"] = nil
	return in
}

// SetSharesNetworksPrimary sets parameter SharesNetworksPrimary to value and selects it for sending
func (in *ActionLocationIndexInput) SetSharesNetworksPrimary(value bool) *ActionLocationIndexInput {
	in.SharesNetworksPrimary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SharesNetworksPrimary"] = nil
	return in
}

// SetSharesV4NetworksWith sets parameter SharesV4NetworksWith to value and selects it for sending
func (in *ActionLocationIndexInput) SetSharesV4NetworksWith(value int64) *ActionLocationIndexInput {
	in.SharesV4NetworksWith = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SharesV4NetworksWith"] = nil
	return in
}

// SetSharesV6NetworksWith sets parameter SharesV6NetworksWith to value and selects it for sending
func (in *ActionLocationIndexInput) SetSharesV6NetworksWith(value int64) *ActionLocationIndexInput {
	in.SharesV6NetworksWith = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["SharesV6NetworksWith"] = nil
	return in
}

// SelectParameters sets parameters from ActionLocationIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLocationIndexInput) SelectParameters(params ...string) *ActionLocationIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLocationIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLocationIndexOutput is a type for action output parameters
type ActionLocationIndexOutput struct {
	Description           string                       `json:"description"`
	Domain                string                       `json:"domain"`
	Environment           *ActionEnvironmentShowOutput `json:"environment"`
	HasIpv6               bool                         `json:"has_ipv6"`
	Id                    int64                        `json:"id"`
	Label                 string                       `json:"label"`
	MaintenanceLock       string                       `json:"maintenance_lock"`
	MaintenanceLockReason string                       `json:"maintenance_lock_reason"`
	RemoteConsoleServer   string                       `json:"remote_console_server"`
	VpsOnboot             bool                         `json:"vps_onboot"`
}

// Type for action response, including envelope
type ActionLocationIndexResponse struct {
	Action *ActionLocationIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Locations []*ActionLocationIndexOutput `json:"locations"`
	}

	// Action output without the namespace
	Output []*ActionLocationIndexOutput
}

// Prepare the action for invocation
func (action *ActionLocationIndex) Prepare() *ActionLocationIndexInvocation {
	return &ActionLocationIndexInvocation{
		Action: action,
		Path:   "/v6.0/locations",
	}
}

// ActionLocationIndexInvocation is used to configure action for invocation
type ActionLocationIndexInvocation struct {
	// Pointer to the action
	Action *ActionLocationIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionLocationIndexInput
	// Global meta input parameters
	MetaInput *ActionLocationIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionLocationIndexInvocation) NewInput() *ActionLocationIndexInput {
	inv.Input = &ActionLocationIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionLocationIndexInvocation) SetInput(input *ActionLocationIndexInput) *ActionLocationIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionLocationIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLocationIndexInvocation) NewMetaInput() *ActionLocationIndexMetaGlobalInput {
	inv.MetaInput = &ActionLocationIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLocationIndexInvocation) SetMetaInput(input *ActionLocationIndexMetaGlobalInput) *ActionLocationIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLocationIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLocationIndexInvocation) Call() (*ActionLocationIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionLocationIndexInvocation) callAsQuery() (*ActionLocationIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionLocationIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Locations
	}
	return resp, err
}

func (inv *ActionLocationIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["location[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("HasHypervisor") {
			ret["location[has_hypervisor]"] = convertBoolToString(inv.Input.HasHypervisor)
		}
		if inv.IsParameterSelected("HasStorage") {
			ret["location[has_storage]"] = convertBoolToString(inv.Input.HasStorage)
		}
		if inv.IsParameterSelected("HypervisorType") {
			ret["location[hypervisor_type]"] = inv.Input.HypervisorType
		}
		if inv.IsParameterSelected("Limit") {
			ret["location[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["location[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("SharesAnyNetworksWith") {
			ret["location[shares_any_networks_with]"] = convertInt64ToString(inv.Input.SharesAnyNetworksWith)
		}
		if inv.IsParameterSelected("SharesNetworksPrimary") {
			ret["location[shares_networks_primary]"] = convertBoolToString(inv.Input.SharesNetworksPrimary)
		}
		if inv.IsParameterSelected("SharesV4NetworksWith") {
			ret["location[shares_v4_networks_with]"] = convertInt64ToString(inv.Input.SharesV4NetworksWith)
		}
		if inv.IsParameterSelected("SharesV6NetworksWith") {
			ret["location[shares_v6_networks_with]"] = convertInt64ToString(inv.Input.SharesV6NetworksWith)
		}
	}
}

func (inv *ActionLocationIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
