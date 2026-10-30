package client

import (
	"net/url"
	"strings"
)

// ActionIpReleaseCampaignShow is a type for action Ip_release_campaign#Show
type ActionIpReleaseCampaignShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseCampaignShow(client *Client) *ActionIpReleaseCampaignShow {
	return &ActionIpReleaseCampaignShow{
		Client: client,
	}
}

// ActionIpReleaseCampaignShowMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseCampaignShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseCampaignShowMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseCampaignShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseCampaignShowMetaGlobalInput) SetNo(value bool) *ActionIpReleaseCampaignShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseCampaignShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseCampaignShowMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseCampaignShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseCampaignShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseCampaignShowOutput is a type for action output parameters
type ActionIpReleaseCampaignShowOutput struct {
	AllowKeep             bool                                      "json:\"allow_keep\""
	CanRelease            bool                                      "json:\"can_release\""
	CanSendInitialNotices bool                                      "json:\"can_send_initial_notices\""
	CanSendReminders      bool                                      "json:\"can_send_reminders\""
	ClosedAt              string                                    "json:\"closed_at\""
	ClosedBy              *ActionUserShowOutput                     "json:\"closed_by\""
	CreatedAt             string                                    "json:\"created_at\""
	CreatedBy             *ActionUserShowOutput                     "json:\"created_by\""
	Deadline              string                                    "json:\"deadline\""
	Id                    int64                                     "json:\"id\""
	KeptIpCount           int64                                     "json:\"kept_ip_count\""
	LatestReleaseAttempt  *ActionIpReleaseCampaignAttemptShowOutput "json:\"latest_release_attempt\""
	ToReleaseIpCount      int64                                     "json:\"to_release_ip_count\""
	TotalIpCount          int64                                     "json:\"total_ip_count\""
}

// Type for action response, including envelope
type ActionIpReleaseCampaignShowResponse struct {
	Action *ActionIpReleaseCampaignShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseCampaign *ActionIpReleaseCampaignShowOutput "json:\"ip_release_campaign\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseCampaignShowOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseCampaignShow) Prepare() *ActionIpReleaseCampaignShowInvocation {
	return &ActionIpReleaseCampaignShowInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_campaigns/{ip_release_campaign_id}",
	}
}

// ActionIpReleaseCampaignShowInvocation is used to configure action for invocation
type ActionIpReleaseCampaignShowInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseCampaignShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpReleaseCampaignShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseCampaignShowInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseCampaignShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseCampaignShowInvocation) SetPathParamString(param string, value string) *ActionIpReleaseCampaignShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseCampaignShowInvocation) NewMetaInput() *ActionIpReleaseCampaignShowMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseCampaignShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseCampaignShowInvocation) SetMetaInput(input *ActionIpReleaseCampaignShowMetaGlobalInput) *ActionIpReleaseCampaignShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseCampaignShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseCampaignShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseCampaignShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpReleaseCampaignShowInvocation) Call() (*ActionIpReleaseCampaignShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIpReleaseCampaignShowInvocation) callAsQuery() (*ActionIpReleaseCampaignShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpReleaseCampaignShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpReleaseCampaign
	}
	return resp, err
}

func (inv *ActionIpReleaseCampaignShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
