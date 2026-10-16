package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionDnsTsigKeyShow is a type for action Dns_tsig_key#Show
type ActionDnsTsigKeyShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsTsigKeyShow(client *Client) *ActionDnsTsigKeyShow {
	return &ActionDnsTsigKeyShow{
		Client: client,
	}
}

// ActionDnsTsigKeyShowMetaGlobalInput is a type for action global meta input parameters
type ActionDnsTsigKeyShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsTsigKeyShowMetaGlobalInput) SetIncludes(value string) *ActionDnsTsigKeyShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsTsigKeyShowMetaGlobalInput) SetNo(value bool) *ActionDnsTsigKeyShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsTsigKeyShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsTsigKeyShowMetaGlobalInput) SelectParameters(params ...string) *ActionDnsTsigKeyShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsTsigKeyShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsTsigKeyShowOutput is a type for action output parameters
type ActionDnsTsigKeyShowOutput struct {
	Algorithm string                "json:\"algorithm\""
	CreatedAt string                "json:\"created_at\""
	Id        int64                 "json:\"id\""
	Name      string                "json:\"name\""
	Secret    string                "json:\"secret\""
	UpdatedAt string                "json:\"updated_at\""
	User      *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionDnsTsigKeyShowResponse struct {
	Action *ActionDnsTsigKeyShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsTsigKey *ActionDnsTsigKeyShowOutput "json:\"dns_tsig_key\""
	}

	// Action output without the namespace
	Output *ActionDnsTsigKeyShowOutput
}

// Prepare the action for invocation
func (action *ActionDnsTsigKeyShow) Prepare() *ActionDnsTsigKeyShowInvocation {
	return &ActionDnsTsigKeyShowInvocation{
		Action: action,
		Path:   "/v7.0/dns_tsig_keys/{dns_tsig_key_id}",
	}
}

// ActionDnsTsigKeyShowInvocation is used to configure action for invocation
type ActionDnsTsigKeyShowInvocation struct {
	// Pointer to the action
	Action *ActionDnsTsigKeyShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDnsTsigKeyShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsTsigKeyShowInvocation) SetPathParamInt(param string, value int64) *ActionDnsTsigKeyShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsTsigKeyShowInvocation) SetPathParamString(param string, value string) *ActionDnsTsigKeyShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsTsigKeyShowInvocation) NewMetaInput() *ActionDnsTsigKeyShowMetaGlobalInput {
	inv.MetaInput = &ActionDnsTsigKeyShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsTsigKeyShowInvocation) SetMetaInput(input *ActionDnsTsigKeyShowMetaGlobalInput) *ActionDnsTsigKeyShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsTsigKeyShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsTsigKeyShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionDnsTsigKeyShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionDnsTsigKeyShowInvocation) Call() (*ActionDnsTsigKeyShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionDnsTsigKeyShowInvocation) CallContext(ctx context.Context) (*ActionDnsTsigKeyShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionDnsTsigKeyShowResponse{Action: inv.Action}
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
		if resp.Response.DnsTsigKey == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.DnsTsigKey
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionDnsTsigKeyShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionDnsTsigKeyShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionDnsTsigKeyShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
