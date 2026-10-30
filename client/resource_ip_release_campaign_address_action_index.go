package client

import (
	"net/url"
	"strings"
)

// ActionIpReleaseCampaignAddressIndex is a type for action Ip_release_campaign.Address#Index
type ActionIpReleaseCampaignAddressIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignAddressIndex(client *Client) *ActionIpReleaseCampaignAddressIndex {
	return &ActionIpReleaseCampaignAddressIndex{
		Client: client,
	}
}

// ActionIpReleaseCampaignAddressIndexMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignAddressIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionIpReleaseCampaignAddressIndexMetaGlobalInput) SetCount(value bool) *ActionIpReleaseCampaignAddressIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignAddressIndexMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignAddressIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignAddressIndexMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignAddressIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignAddressIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignAddressIndexMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignAddressIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignAddressIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignAddressIndexInput is a type for action input parameters
type ActionIpReleaseCampaignAddressIndexInput struct {
	FromId int64 "json:\"from_id\""
	Limit  int64 "json:\"limit\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionIpReleaseCampaignAddressIndexInput) SetFromId(value int64) *ActionIpReleaseCampaignAddressIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionIpReleaseCampaignAddressIndexInput) SetLimit(value int64) *ActionIpReleaseCampaignAddressIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignAddressIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignAddressIndexInput) SelectParameters(params ...string) *ActionIpReleaseCampaignAddressIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionIpReleaseCampaignAddressIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignAddressIndexInput) UnselectParameters(params ...string) *ActionIpReleaseCampaignAddressIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionIpReleaseCampaignAddressIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignAddressIndexOutput is a type for action output parameters
type ActionIpReleaseCampaignAddressIndexOutput struct {
	Address           string                            "json:\"address\""
	AssignIpAddressId int64                             "json:\"assign_ip_address_id\""
	CleanupState      string                            "json:\"cleanup_state\""
	ExcludedAt        string                            "json:\"excluded_at\""
	ExclusionReason   string                            "json:\"exclusion_reason\""
	ExemptedAt        string                            "json:\"exempted_at\""
	ExemptedById      int64                             "json:\"exempted_by_id\""
	ExemptedByLogin   string                            "json:\"exempted_by_login\""
	ExemptionReason   string                            "json:\"exemption_reason\""
	Id                int64                             "json:\"id\""
	IpAddress         *ActionIpAddressShowOutput        "json:\"ip_address\""
	KeepReason        string                            "json:\"keep_reason\""
	KeptAt            string                            "json:\"kept_at\""
	KeptById          int64                             "json:\"kept_by_id\""
	KeptByLogin       string                            "json:\"kept_by_login\""
	LastResult        string                            "json:\"last_result\""
	LocationLabel     string                            "json:\"location_label\""
	OriginalUserId    int64                             "json:\"original_user_id\""
	Prefix            int64                             "json:\"prefix\""
	Protection        string                            "json:\"protection\""
	ReleaseChain      *ActionTransactionChainShowOutput "json:\"release_chain\""
	ReleasedAt        string                            "json:\"released_at\""
	Size              int64                             "json:\"size\""
	UserLogin         string                            "json:\"user_login\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignAddressIndexResponse struct {
	Action *ActionIpReleaseCampaignAddressIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Addresses []*ActionIpReleaseCampaignAddressIndexOutput "json:\"addresses\""
	}

	// Action output without the namespace
	Output []*ActionIpReleaseCampaignAddressIndexOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignAddressIndex) Prepare() *ActionIpReleaseCampaignAddressIndexInvocation {
	return &ActionIpReleaseCampaignAddressIndexInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/{ip_release_campaign_id}/addresses",
	}
}

// ActionIpReleaseCampaignAddressIndexInvocation is used to configure action for invocation
type ActionIpReleaseCampaignAddressIndexInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignAddressIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionIpReleaseCampaignAddressIndexInput
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignAddressIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseCampaignAddressIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) SetPathParamString(param string, value string) *ActionIpReleaseCampaignAddressIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) NewInput() *ActionIpReleaseCampaignAddressIndexInput {
	inv.Input = &ActionIpReleaseCampaignAddressIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) SetInput(input *ActionIpReleaseCampaignAddressIndexInput) *ActionIpReleaseCampaignAddressIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) NewMetaInput() *ActionIpReleaseCampaignAddressIndexMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignAddressIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) SetMetaInput(input *ActionIpReleaseCampaignAddressIndexMetaGlobalInput) *ActionIpReleaseCampaignAddressIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignAddressIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpReleaseCampaignAddressIndexInvocation) Call() (*ActionIpReleaseCampaignAddressIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIpReleaseCampaignAddressIndexInvocation) callAsQuery() (*ActionIpReleaseCampaignAddressIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpReleaseCampaignAddressIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Addresses
	}
	return resp, err
}

func (inv *ActionIpReleaseCampaignAddressIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["address[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["address[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionIpReleaseCampaignAddressIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
