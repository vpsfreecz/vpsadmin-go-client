package client

import (
	"net/url"
	"strings"
)

// ActionIpReleaseRequestShow is a type for action Ip_release_request#Show
type ActionIpReleaseRequestShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIpReleaseRequestShow(client *Client) *ActionIpReleaseRequestShow {
	return &ActionIpReleaseRequestShow{
		Client: client,
	}
}

// ActionIpReleaseRequestShowMetaGlobalInput is a type for action global meta input parameters
type ActionIpReleaseRequestShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpReleaseRequestShowMetaGlobalInput) SetIncludes(value string) *ActionIpReleaseRequestShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpReleaseRequestShowMetaGlobalInput) SetNo(value bool) *ActionIpReleaseRequestShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpReleaseRequestShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpReleaseRequestShowMetaGlobalInput) SelectParameters(params ...string) *ActionIpReleaseRequestShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpReleaseRequestShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpReleaseRequestShowOutput is a type for action output parameters
type ActionIpReleaseRequestShowOutput struct {
	AllowKeep         bool                               "json:\"allow_keep\""
	CanAssign         bool                               "json:\"can_assign\""
	CanKeep           bool                               "json:\"can_keep\""
	ClosedAt          string                             "json:\"closed_at\""
	Deadline          string                             "json:\"deadline\""
	Id                int64                              "json:\"id\""
	IpReleaseCampaign *ActionIpReleaseCampaignShowOutput "json:\"ip_release_campaign\""
	MailLog           *ActionMailLogShowOutput           "json:\"mail_log\""
	NotifiedAt        string                             "json:\"notified_at\""
	OriginalUserId    int64                              "json:\"original_user_id\""
	User              *ActionUserShowOutput              "json:\"user\""
	UserLogin         string                             "json:\"user_login\""
}

// Type for action response, including envelope
type ActionIpReleaseRequestShowResponse struct {
	Action *ActionIpReleaseRequestShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpReleaseRequest *ActionIpReleaseRequestShowOutput "json:\"ip_release_request\""
	}

	// Action output without the namespace
	Output *ActionIpReleaseRequestShowOutput
}

// Prepare the action for invocation
func (action *ActionIpReleaseRequestShow) Prepare() *ActionIpReleaseRequestShowInvocation {
	return &ActionIpReleaseRequestShowInvocation{
		Action: action,
		Path:   "/v7.0/ip_release_requests/{ip_release_request_id}",
	}
}

// ActionIpReleaseRequestShowInvocation is used to configure action for invocation
type ActionIpReleaseRequestShowInvocation struct {
	// Pointer to the action
	Action *ActionIpReleaseRequestShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpReleaseRequestShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpReleaseRequestShowInvocation) SetPathParamInt(param string, value int64) *ActionIpReleaseRequestShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpReleaseRequestShowInvocation) SetPathParamString(param string, value string) *ActionIpReleaseRequestShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpReleaseRequestShowInvocation) NewMetaInput() *ActionIpReleaseRequestShowMetaGlobalInput {
	inv.MetaInput = &ActionIpReleaseRequestShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpReleaseRequestShowInvocation) SetMetaInput(input *ActionIpReleaseRequestShowMetaGlobalInput) *ActionIpReleaseRequestShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpReleaseRequestShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpReleaseRequestShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpReleaseRequestShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIpReleaseRequestShowInvocation) Call() (*ActionIpReleaseRequestShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionIpReleaseRequestShowInvocation) callAsQuery() (*ActionIpReleaseRequestShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIpReleaseRequestShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IpReleaseRequest
	}
	return resp, err
}

func (inv *ActionIpReleaseRequestShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
