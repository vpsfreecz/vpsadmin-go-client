package client

import (
	"strings"
)

// ActionMailboxHandlerShow is a type for action Mailbox.Handler#Show
type ActionMailboxHandlerShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMailboxHandlerShow(client *Client) *ActionMailboxHandlerShow {
	return &ActionMailboxHandlerShow{
		Client: client,
	}
}

// ActionMailboxHandlerShowMetaGlobalInput is a type for action global meta input parameters
type ActionMailboxHandlerShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailboxHandlerShowMetaGlobalInput) SetIncludes(value string) *ActionMailboxHandlerShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailboxHandlerShowMetaGlobalInput) SetNo(value bool) *ActionMailboxHandlerShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailboxHandlerShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailboxHandlerShowMetaGlobalInput) SelectParameters(params ...string) *ActionMailboxHandlerShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailboxHandlerShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailboxHandlerShowOutput is a type for action output parameters
type ActionMailboxHandlerShowOutput struct {
	ClassName string `json:"class_name"`
	Continue  bool   `json:"continue"`
	CreatedAt string `json:"created_at"`
	Id        int64  `json:"id"`
	Order     int64  `json:"order"`
	UpdatedAt string `json:"updated_at"`
}

// Type for action response, including envelope
type ActionMailboxHandlerShowResponse struct {
	Action *ActionMailboxHandlerShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Handler *ActionMailboxHandlerShowOutput `json:"handler"`
	}

	// Action output without the namespace
	Output *ActionMailboxHandlerShowOutput
}

// Prepare the action for invocation
func (action *ActionMailboxHandlerShow) Prepare() *ActionMailboxHandlerShowInvocation {
	return &ActionMailboxHandlerShowInvocation{
		Action: action,
		Path:   "/v7.0/mailboxes/{mailbox_id}/handler/{handler_id}",
	}
}

// ActionMailboxHandlerShowInvocation is used to configure action for invocation
type ActionMailboxHandlerShowInvocation struct {
	// Pointer to the action
	Action *ActionMailboxHandlerShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailboxHandlerShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailboxHandlerShowInvocation) SetPathParamInt(param string, value int64) *ActionMailboxHandlerShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailboxHandlerShowInvocation) SetPathParamString(param string, value string) *ActionMailboxHandlerShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailboxHandlerShowInvocation) NewMetaInput() *ActionMailboxHandlerShowMetaGlobalInput {
	inv.MetaInput = &ActionMailboxHandlerShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailboxHandlerShowInvocation) SetMetaInput(input *ActionMailboxHandlerShowMetaGlobalInput) *ActionMailboxHandlerShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailboxHandlerShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailboxHandlerShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMailboxHandlerShowInvocation) Call() (*ActionMailboxHandlerShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMailboxHandlerShowInvocation) callAsQuery() (*ActionMailboxHandlerShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMailboxHandlerShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Handler
	}
	return resp, err
}

func (inv *ActionMailboxHandlerShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
