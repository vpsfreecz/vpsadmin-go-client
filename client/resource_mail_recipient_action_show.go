package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionMailRecipientShow is a type for action Mail_recipient#Show
type ActionMailRecipientShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMailRecipientShow(client *Client) *ActionMailRecipientShow {
	return &ActionMailRecipientShow{
		Client: client,
	}
}

// ActionMailRecipientShowMetaGlobalInput is a type for action global meta input parameters
type ActionMailRecipientShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailRecipientShowMetaGlobalInput) SetIncludes(value string) *ActionMailRecipientShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailRecipientShowMetaGlobalInput) SetNo(value bool) *ActionMailRecipientShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailRecipientShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailRecipientShowMetaGlobalInput) SelectParameters(params ...string) *ActionMailRecipientShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailRecipientShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailRecipientShowOutput is a type for action output parameters
type ActionMailRecipientShowOutput struct {
	Bcc   string "json:\"bcc\""
	Cc    string "json:\"cc\""
	Id    int64  "json:\"id\""
	Label string "json:\"label\""
	To    string "json:\"to\""
}

// Type for action response, including envelope
type ActionMailRecipientShowResponse struct {
	Action *ActionMailRecipientShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailRecipient *ActionMailRecipientShowOutput "json:\"mail_recipient\""
	}

	// Action output without the namespace
	Output *ActionMailRecipientShowOutput
}

// Prepare the action for invocation
func (action *ActionMailRecipientShow) Prepare() *ActionMailRecipientShowInvocation {
	return &ActionMailRecipientShowInvocation{
		Action: action,
		Path:   "/v7.0/mail_recipients/{mail_recipient_id}",
	}
}

// ActionMailRecipientShowInvocation is used to configure action for invocation
type ActionMailRecipientShowInvocation struct {
	// Pointer to the action
	Action *ActionMailRecipientShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailRecipientShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailRecipientShowInvocation) SetPathParamInt(param string, value int64) *ActionMailRecipientShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailRecipientShowInvocation) SetPathParamString(param string, value string) *ActionMailRecipientShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailRecipientShowInvocation) NewMetaInput() *ActionMailRecipientShowMetaGlobalInput {
	inv.MetaInput = &ActionMailRecipientShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailRecipientShowInvocation) SetMetaInput(input *ActionMailRecipientShowMetaGlobalInput) *ActionMailRecipientShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailRecipientShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailRecipientShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionMailRecipientShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionMailRecipientShowInvocation) Call() (*ActionMailRecipientShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionMailRecipientShowInvocation) CallContext(ctx context.Context) (*ActionMailRecipientShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionMailRecipientShowResponse{Action: inv.Action}
	raw, err := inv.Action.Client.doRaw(req)
	if err != nil {
		return resp, err
	}
	if err := decodeTypedResponse(raw.Body, resp); err != nil {
		return resp, err
	}
	if resp.Status {
		if resp.Response == nil {
			return resp, errMalformedResponse
		}
		if resp.Response.MailRecipient == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.MailRecipient
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionMailRecipientShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionMailRecipientShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionMailRecipientShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
