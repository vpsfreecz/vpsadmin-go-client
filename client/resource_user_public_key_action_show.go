package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionUserPublicKeyShow is a type for action User.Public_key#Show
type ActionUserPublicKeyShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserPublicKeyShow(client *Client) *ActionUserPublicKeyShow {
	return &ActionUserPublicKeyShow{
		Client: client,
	}
}

// ActionUserPublicKeyShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserPublicKeyShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserPublicKeyShowMetaGlobalInput) SetIncludes(value string) *ActionUserPublicKeyShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserPublicKeyShowMetaGlobalInput) SetNo(value bool) *ActionUserPublicKeyShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPublicKeyShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPublicKeyShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserPublicKeyShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPublicKeyShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserPublicKeyShowOutput is a type for action output parameters
type ActionUserPublicKeyShowOutput struct {
	AutoAdd     bool   "json:\"auto_add\""
	Comment     string "json:\"comment\""
	CreatedAt   string "json:\"created_at\""
	Fingerprint string "json:\"fingerprint\""
	Id          int64  "json:\"id\""
	Key         string "json:\"key\""
	Label       string "json:\"label\""
	UpdatedAt   string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionUserPublicKeyShowResponse struct {
	Action *ActionUserPublicKeyShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PublicKey *ActionUserPublicKeyShowOutput "json:\"public_key\""
	}

	// Action output without the namespace
	Output *ActionUserPublicKeyShowOutput
}

// Prepare the action for invocation
func (action *ActionUserPublicKeyShow) Prepare() *ActionUserPublicKeyShowInvocation {
	return &ActionUserPublicKeyShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/public_keys/{public_key_id}",
	}
}

// ActionUserPublicKeyShowInvocation is used to configure action for invocation
type ActionUserPublicKeyShowInvocation struct {
	// Pointer to the action
	Action *ActionUserPublicKeyShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserPublicKeyShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserPublicKeyShowInvocation) SetPathParamInt(param string, value int64) *ActionUserPublicKeyShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserPublicKeyShowInvocation) SetPathParamString(param string, value string) *ActionUserPublicKeyShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserPublicKeyShowInvocation) NewMetaInput() *ActionUserPublicKeyShowMetaGlobalInput {
	inv.MetaInput = &ActionUserPublicKeyShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserPublicKeyShowInvocation) SetMetaInput(input *ActionUserPublicKeyShowMetaGlobalInput) *ActionUserPublicKeyShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserPublicKeyShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserPublicKeyShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserPublicKeyShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionUserPublicKeyShowInvocation) Call() (*ActionUserPublicKeyShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionUserPublicKeyShowInvocation) CallContext(ctx context.Context) (*ActionUserPublicKeyShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionUserPublicKeyShowResponse{Action: inv.Action}
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
		if resp.Response.PublicKey == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.PublicKey
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionUserPublicKeyShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionUserPublicKeyShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionUserPublicKeyShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
