package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionVpsSshHostKeyShow is a type for action Vps.Ssh_host_key#Show
type ActionVpsSshHostKeyShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsSshHostKeyShow(client *Client) *ActionVpsSshHostKeyShow {
	return &ActionVpsSshHostKeyShow{
		Client: client,
	}
}

// ActionVpsSshHostKeyShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsSshHostKeyShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsSshHostKeyShowMetaGlobalInput) SetIncludes(value string) *ActionVpsSshHostKeyShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsSshHostKeyShowMetaGlobalInput) SetNo(value bool) *ActionVpsSshHostKeyShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsSshHostKeyShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsSshHostKeyShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsSshHostKeyShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsSshHostKeyShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsSshHostKeyShowOutput is a type for action output parameters
type ActionVpsSshHostKeyShowOutput struct {
	Algorithm   string "json:\"algorithm\""
	Bits        int64  "json:\"bits\""
	CreatedAt   string "json:\"created_at\""
	Fingerprint string "json:\"fingerprint\""
	Id          int64  "json:\"id\""
	UpdatedAt   string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionVpsSshHostKeyShowResponse struct {
	Action *ActionVpsSshHostKeyShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		SshHostKey *ActionVpsSshHostKeyShowOutput "json:\"ssh_host_key\""
	}

	// Action output without the namespace
	Output *ActionVpsSshHostKeyShowOutput
}

// Prepare the action for invocation
func (action *ActionVpsSshHostKeyShow) Prepare() *ActionVpsSshHostKeyShowInvocation {
	return &ActionVpsSshHostKeyShowInvocation{
		Action: action,
		Path:   "/v7.0/vpses/{vps_id}/ssh_host_keys/{ssh_host_key_id}",
	}
}

// ActionVpsSshHostKeyShowInvocation is used to configure action for invocation
type ActionVpsSshHostKeyShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsSshHostKeyShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsSshHostKeyShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsSshHostKeyShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsSshHostKeyShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsSshHostKeyShowInvocation) SetPathParamString(param string, value string) *ActionVpsSshHostKeyShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsSshHostKeyShowInvocation) NewMetaInput() *ActionVpsSshHostKeyShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsSshHostKeyShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsSshHostKeyShowInvocation) SetMetaInput(input *ActionVpsSshHostKeyShowMetaGlobalInput) *ActionVpsSshHostKeyShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsSshHostKeyShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsSshHostKeyShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionVpsSshHostKeyShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionVpsSshHostKeyShowInvocation) Call() (*ActionVpsSshHostKeyShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionVpsSshHostKeyShowInvocation) CallContext(ctx context.Context) (*ActionVpsSshHostKeyShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionVpsSshHostKeyShowResponse{Action: inv.Action}
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
		if resp.Response.SshHostKey == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.SshHostKey
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionVpsSshHostKeyShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionVpsSshHostKeyShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionVpsSshHostKeyShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
