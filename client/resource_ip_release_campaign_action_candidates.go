package client

import ()

// ActionIpReleaseCampaignCandidates is a type for action Ip_release_campaign#Candidates
type ActionIpReleaseCampaignCandidates struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignCandidates(client *Client) *ActionIpReleaseCampaignCandidates {
	return &ActionIpReleaseCampaignCandidates{
		Client: client,
	}
}

// ActionIpReleaseCampaignCandidatesMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignCandidatesMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesMetaGlobalInput) SetCount(value bool) *ActionIpReleaseCampaignCandidatesMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignCandidatesMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignCandidatesMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignCandidatesMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignCandidatesMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignCandidatesMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignCandidatesMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignCandidatesInput is a type for action input parameters
type ActionIpReleaseCampaignCandidatesInput struct {
	Access    string "json:\"access\""
	FromId    int64  "json:\"from_id\""
	Limit     int64  "json:\"limit\""
	Locations string "json:\"locations\""
	Networks  string "json:\"networks\""
	User      int64  "json:\"user\""
	Versions  string "json:\"versions\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAccess sets parameter Access to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesInput) SetAccess(value string) *ActionIpReleaseCampaignCandidatesInput {
	in.Access = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Access"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesInput) SetFromId(value int64) *ActionIpReleaseCampaignCandidatesInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesInput) SetLimit(value int64) *ActionIpReleaseCampaignCandidatesInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocations sets parameter Locations to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesInput) SetLocations(value string) *ActionIpReleaseCampaignCandidatesInput {
	in.Locations = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Locations"] = nil
	return in
}

// SetNetworks sets parameter Networks to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesInput) SetNetworks(value string) *ActionIpReleaseCampaignCandidatesInput {
	in.Networks = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Networks"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesInput) SetUser(value int64) *ActionIpReleaseCampaignCandidatesInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SetVersions sets parameter Versions to value and selects it for sending
func (in *ActionIpReleaseCampaignCandidatesInput) SetVersions(value string) *ActionIpReleaseCampaignCandidatesInput {
	in.Versions = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Versions"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignCandidatesInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignCandidatesInput) SelectParameters(params ...string) *ActionIpReleaseCampaignCandidatesInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseCampaignCandidatesInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignCandidatesInput) UnselectParameters(params ...string) *ActionIpReleaseCampaignCandidatesInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseCampaignCandidatesInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignCandidatesOutput is a type for action output parameters
type ActionIpReleaseCampaignCandidatesOutput struct {
	Addr    string                   "json:\"addr\""
	Id      int64                    "json:\"id\""
	Network *ActionNetworkShowOutput "json:\"network\""
	Prefix  int64                    "json:\"prefix\""
	Size    int64                    "json:\"size\""
	User    *ActionUserShowOutput    "json:\"user\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignCandidatesResponse struct {
	Action *ActionIpReleaseCampaignCandidates "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseCampaigns []*ActionIpReleaseCampaignCandidatesOutput "json:\"ip_release_campaigns\""
	}

	// Action output without the namespace
	Output []*ActionIpReleaseCampaignCandidatesOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignCandidates) Prepare() *ActionIpReleaseCampaignCandidatesInvocation {
	return &ActionIpReleaseCampaignCandidatesInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/candidates",
	}
}

// ActionIpReleaseCampaignCandidatesInvocation is used to configure action for invocation
type ActionIpReleaseCampaignCandidatesInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignCandidates

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseCampaignCandidatesInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignCandidatesMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseCampaignCandidatesInvocation) NewInput() *ActionIpReleaseCampaignCandidatesInput {
	inv.Input = &ActionIpReleaseCampaignCandidatesInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseCampaignCandidatesInvocation) SetInput(input *ActionIpReleaseCampaignCandidatesInput) *ActionIpReleaseCampaignCandidatesInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseCampaignCandidatesInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignCandidatesInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignCandidatesInvocation) NewMetaInput() *ActionIpReleaseCampaignCandidatesMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignCandidatesMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignCandidatesInvocation) SetMetaInput(input *ActionIpReleaseCampaignCandidatesMetaGlobalInput) *ActionIpReleaseCampaignCandidatesInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignCandidatesInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignCandidatesInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignCandidatesInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", t(inv.Action.Client, "validation.invalid_resource_id", nil))
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
func (inv *ActionIpReleaseCampaignCandidatesInvocation) Call() (*ActionIpReleaseCampaignCandidatesResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIpReleaseCampaignCandidatesInvocation) callAsQuery() (*ActionIpReleaseCampaignCandidatesResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpReleaseCampaignCandidatesResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpReleaseCampaigns
	}
	return resp, err
}

func (inv *ActionIpReleaseCampaignCandidatesInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Access") {
			ret["ip_release_campaign[access]"] = inv.Input.Access
		}
		if inv.IsParameterSelected("FromId") {
			ret["ip_release_campaign[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["ip_release_campaign[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Locations") {
			ret["ip_release_campaign[locations]"] = inv.Input.Locations
		}
		if inv.IsParameterSelected("Networks") {
			ret["ip_release_campaign[networks]"] = inv.Input.Networks
		}
		if inv.IsParameterSelected("User") {
			ret["ip_release_campaign[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Versions") {
			ret["ip_release_campaign[versions]"] = inv.Input.Versions
		}
	}
}

func (inv *ActionIpReleaseCampaignCandidatesInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
