package client

import (
	"strings"
)

// ActionUserRequestRegistrationShow is a type for action User_request.Registration#Show
type ActionUserRequestRegistrationShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestRegistrationShow(client *Client) *ActionUserRequestRegistrationShow {
	return &ActionUserRequestRegistrationShow{
		Client: client,
	}
}

// ActionUserRequestRegistrationShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestRegistrationShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestRegistrationShowMetaGlobalInput) SetIncludes(value string) *ActionUserRequestRegistrationShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestRegistrationShowMetaGlobalInput) SetNo(value bool) *ActionUserRequestRegistrationShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestRegistrationShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationShowOutput is a type for action output parameters
type ActionUserRequestRegistrationShowOutput struct {
	Address                string                      `json:"address"`
	Admin                  *ActionUserShowOutput       `json:"admin"`
	AdminResponse          string                      `json:"admin_response"`
	ApiIpAddr              string                      `json:"api_ip_addr"`
	ApiIpPtr               string                      `json:"api_ip_ptr"`
	ClientIpAddr           string                      `json:"client_ip_addr"`
	ClientIpPtr            string                      `json:"client_ip_ptr"`
	CreatedAt              string                      `json:"created_at"`
	Currency               string                      `json:"currency"`
	Email                  string                      `json:"email"`
	FullName               string                      `json:"full_name"`
	How                    string                      `json:"how"`
	Id                     int64                       `json:"id"`
	IpChecked              bool                        `json:"ip_checked"`
	IpCrawler              bool                        `json:"ip_crawler"`
	IpErrors               string                      `json:"ip_errors"`
	IpFraudScore           int64                       `json:"ip_fraud_score"`
	IpMessage              string                      `json:"ip_message"`
	IpProxy                bool                        `json:"ip_proxy"`
	IpRecentAbuse          bool                        `json:"ip_recent_abuse"`
	IpRequestId            string                      `json:"ip_request_id"`
	IpSuccess              bool                        `json:"ip_success"`
	IpTor                  bool                        `json:"ip_tor"`
	IpVpn                  bool                        `json:"ip_vpn"`
	Label                  string                      `json:"label"`
	Language               *ActionLanguageShowOutput   `json:"language"`
	Location               *ActionLocationShowOutput   `json:"location"`
	Login                  string                      `json:"login"`
	MailCatchAll           bool                        `json:"mail_catch_all"`
	MailChecked            bool                        `json:"mail_checked"`
	MailDeliverability     string                      `json:"mail_deliverability"`
	MailDisposable         bool                        `json:"mail_disposable"`
	MailDnsValid           bool                        `json:"mail_dns_valid"`
	MailErrors             string                      `json:"mail_errors"`
	MailFraudScore         int64                       `json:"mail_fraud_score"`
	MailFrequentComplainer bool                        `json:"mail_frequent_complainer"`
	MailHoneypot           bool                        `json:"mail_honeypot"`
	MailLeaked             bool                        `json:"mail_leaked"`
	MailMessage            string                      `json:"mail_message"`
	MailOverallScore       int64                       `json:"mail_overall_score"`
	MailRecentAbuse        bool                        `json:"mail_recent_abuse"`
	MailRequestId          string                      `json:"mail_request_id"`
	MailSmtpScore          int64                       `json:"mail_smtp_score"`
	MailSpamTrapScore      string                      `json:"mail_spam_trap_score"`
	MailSuccess            bool                        `json:"mail_success"`
	MailSuspect            bool                        `json:"mail_suspect"`
	MailTimedOut           bool                        `json:"mail_timed_out"`
	MailValid              bool                        `json:"mail_valid"`
	Note                   string                      `json:"note"`
	OrgId                  string                      `json:"org_id"`
	OrgName                string                      `json:"org_name"`
	OsTemplate             *ActionOsTemplateShowOutput `json:"os_template"`
	State                  string                      `json:"state"`
	UpdatedAt              string                      `json:"updated_at"`
	User                   *ActionUserShowOutput       `json:"user"`
	YearOfBirth            int64                       `json:"year_of_birth"`
}

// Type for action response, including envelope
type ActionUserRequestRegistrationShowResponse struct {
	Action *ActionUserRequestRegistrationShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Registration *ActionUserRequestRegistrationShowOutput `json:"registration"`
	}

	// Action output without the namespace
	Output *ActionUserRequestRegistrationShowOutput
}

// Prepare the action for invocation
func (action *ActionUserRequestRegistrationShow) Prepare() *ActionUserRequestRegistrationShowInvocation {
	return &ActionUserRequestRegistrationShowInvocation{
		Action: action,
		Path:   "/v7.0/user_request/registrations/{registration_id}",
	}
}

// ActionUserRequestRegistrationShowInvocation is used to configure action for invocation
type ActionUserRequestRegistrationShowInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestRegistrationShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserRequestRegistrationShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserRequestRegistrationShowInvocation) SetPathParamInt(param string, value int64) *ActionUserRequestRegistrationShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserRequestRegistrationShowInvocation) SetPathParamString(param string, value string) *ActionUserRequestRegistrationShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestRegistrationShowInvocation) NewMetaInput() *ActionUserRequestRegistrationShowMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestRegistrationShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestRegistrationShowInvocation) SetMetaInput(input *ActionUserRequestRegistrationShowMetaGlobalInput) *ActionUserRequestRegistrationShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestRegistrationShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserRequestRegistrationShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestRegistrationShowInvocation) Call() (*ActionUserRequestRegistrationShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserRequestRegistrationShowInvocation) callAsQuery() (*ActionUserRequestRegistrationShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserRequestRegistrationShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Registration
	}
	return resp, err
}

func (inv *ActionUserRequestRegistrationShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
