package client

import ()

// ActionUserRequestRegistrationIndex is a type for action User_request.Registration#Index
type ActionUserRequestRegistrationIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestRegistrationIndex(client *Client) *ActionUserRequestRegistrationIndex {
	return &ActionUserRequestRegistrationIndex{
		Client: client,
	}
}

// ActionUserRequestRegistrationIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestRegistrationIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) SetCount(value bool) *ActionUserRequestRegistrationIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) SetIncludes(value string) *ActionUserRequestRegistrationIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) SetNo(value bool) *ActionUserRequestRegistrationIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestRegistrationIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationIndexInput is a type for action input parameters
type ActionUserRequestRegistrationIndexInput struct {
	Admin        int64  `json:"admin"`
	ApiIpAddr    string `json:"api_ip_addr"`
	ClientIpAddr string `json:"client_ip_addr"`
	ClientIpPtr  string `json:"client_ip_ptr"`
	FromId       int64  `json:"from_id"`
	Limit        int64  `json:"limit"`
	State        string `json:"state"`
	User         int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAdmin sets parameter Admin to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetAdmin(value int64) *ActionUserRequestRegistrationIndexInput {
	in.Admin = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetAdminNil(false)
	in._selectedParameters["Admin"] = nil
	return in
}

// SetAdminNil sets parameter Admin to nil and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetAdminNil(set bool) *ActionUserRequestRegistrationIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Admin"] = nil
		in.SelectParameters("Admin")
	} else {
		delete(in._nilParameters, "Admin")
	}
	return in
}

// SetApiIpAddr sets parameter ApiIpAddr to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetApiIpAddr(value string) *ActionUserRequestRegistrationIndexInput {
	in.ApiIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ApiIpAddr"] = nil
	return in
}

// SetClientIpAddr sets parameter ClientIpAddr to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetClientIpAddr(value string) *ActionUserRequestRegistrationIndexInput {
	in.ClientIpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientIpAddr"] = nil
	return in
}

// SetClientIpPtr sets parameter ClientIpPtr to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetClientIpPtr(value string) *ActionUserRequestRegistrationIndexInput {
	in.ClientIpPtr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClientIpPtr"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetFromId(value int64) *ActionUserRequestRegistrationIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetLimit(value int64) *ActionUserRequestRegistrationIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetState(value string) *ActionUserRequestRegistrationIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetUser(value int64) *ActionUserRequestRegistrationIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionUserRequestRegistrationIndexInput) SetUserNil(set bool) *ActionUserRequestRegistrationIndexInput {
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

// SelectParameters sets parameters from ActionUserRequestRegistrationIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationIndexInput) SelectParameters(params ...string) *ActionUserRequestRegistrationIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserRequestRegistrationIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationIndexInput) UnselectParameters(params ...string) *ActionUserRequestRegistrationIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserRequestRegistrationIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserRequestRegistrationIndexOutput is a type for action output parameters
type ActionUserRequestRegistrationIndexOutput struct {
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
type ActionUserRequestRegistrationIndexResponse struct {
	Action *ActionUserRequestRegistrationIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Registrations []*ActionUserRequestRegistrationIndexOutput `json:"registrations"`
	}

	// Action output without the namespace
	Output []*ActionUserRequestRegistrationIndexOutput
}

// Prepare the action for invocation
func (action *ActionUserRequestRegistrationIndex) Prepare() *ActionUserRequestRegistrationIndexInvocation {
	return &ActionUserRequestRegistrationIndexInvocation{
		Action: action,
		Path:   "/v7.0/user_request/registrations",
	}
}

// ActionUserRequestRegistrationIndexInvocation is used to configure action for invocation
type ActionUserRequestRegistrationIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestRegistrationIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserRequestRegistrationIndexInput
	// Global meta input parameters
	MetaInput *ActionUserRequestRegistrationIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserRequestRegistrationIndexInvocation) NewInput() *ActionUserRequestRegistrationIndexInput {
	inv.Input = &ActionUserRequestRegistrationIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserRequestRegistrationIndexInvocation) SetInput(input *ActionUserRequestRegistrationIndexInput) *ActionUserRequestRegistrationIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserRequestRegistrationIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserRequestRegistrationIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestRegistrationIndexInvocation) NewMetaInput() *ActionUserRequestRegistrationIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestRegistrationIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestRegistrationIndexInvocation) SetMetaInput(input *ActionUserRequestRegistrationIndexMetaGlobalInput) *ActionUserRequestRegistrationIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestRegistrationIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserRequestRegistrationIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestRegistrationIndexInvocation) Call() (*ActionUserRequestRegistrationIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserRequestRegistrationIndexInvocation) callAsQuery() (*ActionUserRequestRegistrationIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserRequestRegistrationIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Registrations
	}
	return resp, err
}

func (inv *ActionUserRequestRegistrationIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Admin") {
			ret["registration[admin]"] = convertInt64ToString(inv.Input.Admin)
		}
		if inv.IsParameterSelected("ApiIpAddr") {
			ret["registration[api_ip_addr]"] = inv.Input.ApiIpAddr
		}
		if inv.IsParameterSelected("ClientIpAddr") {
			ret["registration[client_ip_addr]"] = inv.Input.ClientIpAddr
		}
		if inv.IsParameterSelected("ClientIpPtr") {
			ret["registration[client_ip_ptr]"] = inv.Input.ClientIpPtr
		}
		if inv.IsParameterSelected("FromId") {
			ret["registration[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["registration[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("State") {
			ret["registration[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("User") {
			ret["registration[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionUserRequestRegistrationIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
