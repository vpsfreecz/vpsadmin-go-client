package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionMailTemplateShow is a type for action Mail_template#Show
type ActionMailTemplateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMailTemplateShow(client *Client) *ActionMailTemplateShow {
	return &ActionMailTemplateShow{
		Client: client,
	}
}

// ActionMailTemplateShowMetaGlobalInput is a type for action global meta input parameters
type ActionMailTemplateShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMailTemplateShowMetaGlobalInput) SetIncludes(value string) *ActionMailTemplateShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMailTemplateShowMetaGlobalInput) SetNo(value bool) *ActionMailTemplateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMailTemplateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMailTemplateShowMetaGlobalInput) SelectParameters(params ...string) *ActionMailTemplateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMailTemplateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMailTemplateShowOutput is a type for action output parameters
type ActionMailTemplateShowOutput struct {
	CreatedAt      string "json:\"created_at\""
	Id             int64  "json:\"id\""
	Label          string "json:\"label\""
	Name           string "json:\"name\""
	TemplateId     string "json:\"template_id\""
	UpdatedAt      string "json:\"updated_at\""
	UserVisibility string "json:\"user_visibility\""
}

// Type for action response, including envelope
type ActionMailTemplateShowResponse struct {
	Action *ActionMailTemplateShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MailTemplate *ActionMailTemplateShowOutput "json:\"mail_template\""
	}

	// Action output without the namespace
	Output *ActionMailTemplateShowOutput
}

// Prepare the action for invocation
func (action *ActionMailTemplateShow) Prepare() *ActionMailTemplateShowInvocation {
	return &ActionMailTemplateShowInvocation{
		Action: action,
		Path:   "/v7.0/mail_templates/{mail_template_id}",
	}
}

// ActionMailTemplateShowInvocation is used to configure action for invocation
type ActionMailTemplateShowInvocation struct {
	// Pointer to the action
	Action *ActionMailTemplateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMailTemplateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMailTemplateShowInvocation) SetPathParamInt(param string, value int64) *ActionMailTemplateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMailTemplateShowInvocation) SetPathParamString(param string, value string) *ActionMailTemplateShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMailTemplateShowInvocation) NewMetaInput() *ActionMailTemplateShowMetaGlobalInput {
	inv.MetaInput = &ActionMailTemplateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMailTemplateShowInvocation) SetMetaInput(input *ActionMailTemplateShowMetaGlobalInput) *ActionMailTemplateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMailTemplateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMailTemplateShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionMailTemplateShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionMailTemplateShowInvocation) Call() (*ActionMailTemplateShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionMailTemplateShowInvocation) CallContext(ctx context.Context) (*ActionMailTemplateShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionMailTemplateShowResponse{Action: inv.Action}
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
		if resp.Response.MailTemplate == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.MailTemplate
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionMailTemplateShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionMailTemplateShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionMailTemplateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
