package client

import (
	"strings"
)

// ActionMailboxShow is a type for action Mailbox#Show
type ActionMailboxShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxShow(client *Client) *ActionMailboxShow {
	return &ActionMailboxShow{
		Client: client,
	}
}

// ActionMailboxShowMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxShowMetaGlobalInput) SetIncludes(value string) *ActionMailboxShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxShowMetaGlobalInput) SetNo(value bool) *ActionMailboxShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxShowMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxShowOutput is a type for action output parameters
type ActionMailboxShowOutput struct {
	CreatedAt string `json:"created_at"`
	EnableSsl bool   `json:"enable_ssl"`
	Id        int64  `json:"id"`
	Label     string `json:"label"`
	Port      int64  `json:"port"`
	Server    string `json:"server"`
	UpdatedAt string `json:"updated_at"`
	User      string `json:"user"`
}

// Type for action response, including envelope
type ActionMailboxShowResponse struct {
	Action *ActionMailboxShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Mailbox *ActionMailboxShowOutput `json:"mailbox"`
	}

	// Action output without the namespace
	Output *ActionMailboxShowOutput
}

// Prepare the action for invocation
func (action *ActionMailboxShow) Prepare() *ActionMailboxShowInvocation {
	return &ActionMailboxShowInvocation{
		Action: action,
		Path:   "/v6.0/mailboxes/{mailbox_id}",
	}
}

// ActionMailboxShowInvocation is used to configure action for invocation
type ActionMailboxShowInvocation struct {
	// Pointer to the action
	Action *ActionMailboxShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailboxShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailboxShowInvocation) SetPathParamInt(param string, value int64) *ActionMailboxShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailboxShowInvocation) SetPathParamString(param string, value string) *ActionMailboxShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxShowInvocation) NewMetaInput() *ActionMailboxShowMetaGlobalInput {
	inv.MetaInput = &ActionMailboxShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxShowInvocation) SetMetaInput(input *ActionMailboxShowMetaGlobalInput) *ActionMailboxShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxShowInvocation) Call() (*ActionMailboxShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailboxShowInvocation) callAsQuery() (*ActionMailboxShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailboxShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Mailbox
	}
	return resp, err
}

func (inv *ActionMailboxShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
