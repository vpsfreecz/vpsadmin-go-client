package client

import ()

// ActionIpAddressAssignmentIndex is a type for action Ip_address_assignment#Index
type ActionIpAddressAssignmentIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressAssignmentIndex(client *Client) *ActionIpAddressAssignmentIndex {
	return &ActionIpAddressAssignmentIndex{
		Client: client,
	}
}

// ActionIpAddressAssignmentIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressAssignmentIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexMetaGlobalInput) SetCount(value bool) *ActionIpAddressAssignmentIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexMetaGlobalInput) SetIncludes(value string) *ActionIpAddressAssignmentIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexMetaGlobalInput) SetNo(value bool) *ActionIpAddressAssignmentIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressAssignmentIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressAssignmentIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressAssignmentIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressAssignmentIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressAssignmentIndexInput is a type for action input parameters
type ActionIpAddressAssignmentIndexInput struct {
	Active            bool   `json:"active"`
	AssignedByChain   int64  `json:"assigned_by_chain"`
	FromId            int64  `json:"from_id"`
	IpAddr            string `json:"ip_addr"`
	IpAddress         int64  `json:"ip_address"`
	IpPrefix          int64  `json:"ip_prefix"`
	IpVersion         int64  `json:"ip_version"`
	Limit             int64  `json:"limit"`
	Location          int64  `json:"location"`
	Network           int64  `json:"network"`
	Order             string `json:"order"`
	Reconstructed     bool   `json:"reconstructed"`
	UnassignedByChain int64  `json:"unassigned_by_chain"`
	User              int64  `json:"user"`
	Vps               int64  `json:"vps"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetActive sets parameter Active to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetActive(value bool) *ActionIpAddressAssignmentIndexInput {
	in.Active = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Active"] = nil
	return in
}

// SetAssignedByChain sets parameter AssignedByChain to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetAssignedByChain(value int64) *ActionIpAddressAssignmentIndexInput {
	in.AssignedByChain = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetAssignedByChainNil(false)
	in._selectedParameters["AssignedByChain"] = nil
	return in
}

// SetAssignedByChainNil sets parameter AssignedByChain to nil and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetAssignedByChainNil(set bool) *ActionIpAddressAssignmentIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["AssignedByChain"] = nil
		in.SelectParameters("AssignedByChain")
	} else {
		delete(in._nilParameters, "AssignedByChain")
	}
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetFromId(value int64) *ActionIpAddressAssignmentIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetIpAddr sets parameter IpAddr to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetIpAddr(value string) *ActionIpAddressAssignmentIndexInput {
	in.IpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpAddr"] = nil
	return in
}

// SetIpAddress sets parameter IpAddress to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetIpAddress(value int64) *ActionIpAddressAssignmentIndexInput {
	in.IpAddress = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetIpAddressNil(false)
	in._selectedParameters["IpAddress"] = nil
	return in
}

// SetIpAddressNil sets parameter IpAddress to nil and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetIpAddressNil(set bool) *ActionIpAddressAssignmentIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["IpAddress"] = nil
		in.SelectParameters("IpAddress")
	} else {
		delete(in._nilParameters, "IpAddress")
	}
	return in
}

// SetIpPrefix sets parameter IpPrefix to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetIpPrefix(value int64) *ActionIpAddressAssignmentIndexInput {
	in.IpPrefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpPrefix"] = nil
	return in
}

// SetIpVersion sets parameter IpVersion to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetIpVersion(value int64) *ActionIpAddressAssignmentIndexInput {
	in.IpVersion = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpVersion"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetLimit(value int64) *ActionIpAddressAssignmentIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetLocation(value int64) *ActionIpAddressAssignmentIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetLocationNil(set bool) *ActionIpAddressAssignmentIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Location"] = nil
		in.SelectParameters("Location")
	} else {
		delete(in._nilParameters, "Location")
	}
	return in
}

// SetNetwork sets parameter Network to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetNetwork(value int64) *ActionIpAddressAssignmentIndexInput {
	in.Network = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNetworkNil(false)
	in._selectedParameters["Network"] = nil
	return in
}

// SetNetworkNil sets parameter Network to nil and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetNetworkNil(set bool) *ActionIpAddressAssignmentIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Network"] = nil
		in.SelectParameters("Network")
	} else {
		delete(in._nilParameters, "Network")
	}
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetOrder(value string) *ActionIpAddressAssignmentIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetReconstructed sets parameter Reconstructed to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetReconstructed(value bool) *ActionIpAddressAssignmentIndexInput {
	in.Reconstructed = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Reconstructed"] = nil
	return in
}

// SetUnassignedByChain sets parameter UnassignedByChain to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetUnassignedByChain(value int64) *ActionIpAddressAssignmentIndexInput {
	in.UnassignedByChain = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUnassignedByChainNil(false)
	in._selectedParameters["UnassignedByChain"] = nil
	return in
}

// SetUnassignedByChainNil sets parameter UnassignedByChain to nil and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetUnassignedByChainNil(set bool) *ActionIpAddressAssignmentIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["UnassignedByChain"] = nil
		in.SelectParameters("UnassignedByChain")
	} else {
		delete(in._nilParameters, "UnassignedByChain")
	}
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetUser(value int64) *ActionIpAddressAssignmentIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetUserNil(set bool) *ActionIpAddressAssignmentIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetVps(value int64) *ActionIpAddressAssignmentIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetVpsNil(false)
	in._selectedParameters["Vps"] = nil
	return in
}

// SetVpsNil sets parameter Vps to nil and selects it for sending
func (in *ActionIpAddressAssignmentIndexInput) SetVpsNil(set bool) *ActionIpAddressAssignmentIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Vps"] = nil
		in.SelectParameters("Vps")
	} else {
		delete(in._nilParameters, "Vps")
	}
	return in
}

// SelectParameters sets parameters from ActionIpAddressAssignmentIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressAssignmentIndexInput) SelectParameters(params ...string) *ActionIpAddressAssignmentIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpAddressAssignmentIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpAddressAssignmentIndexInput) UnselectParameters(params ...string) *ActionIpAddressAssignmentIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpAddressAssignmentIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressAssignmentIndexOutput is a type for action output parameters
type ActionIpAddressAssignmentIndexOutput struct {
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
type ActionIpAddressAssignmentIndexResponse struct {
	Action *ActionIpAddressAssignmentIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddressAssignments []*ActionIpAddressAssignmentIndexOutput `json:"ip_address_assignments"`
	}

	// Action output without the namespace
	Output []*ActionIpAddressAssignmentIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpAddressAssignmentIndex) Prepare() *ActionIpAddressAssignmentIndexInvocation {
	return &ActionIpAddressAssignmentIndexInvocation{
		Action: action,
		Path:   "/v7.0/ip_address_assignments",
	}
}

// ActionIpAddressAssignmentIndexInvocation is used to configure action for invocation
type ActionIpAddressAssignmentIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressAssignmentIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpAddressAssignmentIndexInput
	// Global meta input parameters
	MetaInput *ActionIpAddressAssignmentIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpAddressAssignmentIndexInvocation) NewInput() *ActionIpAddressAssignmentIndexInput {
	inv.Input = &ActionIpAddressAssignmentIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpAddressAssignmentIndexInvocation) SetInput(input *ActionIpAddressAssignmentIndexInput) *ActionIpAddressAssignmentIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpAddressAssignmentIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpAddressAssignmentIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpAddressAssignmentIndexInvocation) NewMetaInput() *ActionIpAddressAssignmentIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpAddressAssignmentIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressAssignmentIndexInvocation) SetMetaInput(input *ActionIpAddressAssignmentIndexMetaGlobalInput) *ActionIpAddressAssignmentIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressAssignmentIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpAddressAssignmentIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpAddressAssignmentIndexInvocation) Call() (*ActionIpAddressAssignmentIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIpAddressAssignmentIndexInvocation) callAsQuery() (*ActionIpAddressAssignmentIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpAddressAssignmentIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpAddressAssignments
	}
	return resp, err
}

func (inv *ActionIpAddressAssignmentIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Active") {
			ret["ip_address_assignment[active]"] = convertBoolToString(inv.Input.Active)
		}
		if inv.IsParameterSelected("AssignedByChain") {
			ret["ip_address_assignment[assigned_by_chain]"] = convertInt64ToString(inv.Input.AssignedByChain)
		}
		if inv.IsParameterSelected("FromId") {
			ret["ip_address_assignment[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("IpAddr") {
			ret["ip_address_assignment[ip_addr]"] = inv.Input.IpAddr
		}
		if inv.IsParameterSelected("IpAddress") {
			ret["ip_address_assignment[ip_address]"] = convertInt64ToString(inv.Input.IpAddress)
		}
		if inv.IsParameterSelected("IpPrefix") {
			ret["ip_address_assignment[ip_prefix]"] = convertInt64ToString(inv.Input.IpPrefix)
		}
		if inv.IsParameterSelected("IpVersion") {
			ret["ip_address_assignment[ip_version]"] = convertInt64ToString(inv.Input.IpVersion)
		}
		if inv.IsParameterSelected("Limit") {
			ret["ip_address_assignment[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["ip_address_assignment[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Network") {
			ret["ip_address_assignment[network]"] = convertInt64ToString(inv.Input.Network)
		}
		if inv.IsParameterSelected("Order") {
			ret["ip_address_assignment[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("Reconstructed") {
			ret["ip_address_assignment[reconstructed]"] = convertBoolToString(inv.Input.Reconstructed)
		}
		if inv.IsParameterSelected("UnassignedByChain") {
			ret["ip_address_assignment[unassigned_by_chain]"] = convertInt64ToString(inv.Input.UnassignedByChain)
		}
		if inv.IsParameterSelected("User") {
			ret["ip_address_assignment[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["ip_address_assignment[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
	}
}

func (inv *ActionIpAddressAssignmentIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
