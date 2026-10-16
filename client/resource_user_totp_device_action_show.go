package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionUserTotpDeviceShow is a type for action User.Totp_device#Show
type ActionUserTotpDeviceShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserTotpDeviceShow(client *Client) *ActionUserTotpDeviceShow {
	return &ActionUserTotpDeviceShow{
		Client: client,
	}
}

// ActionUserTotpDeviceShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserTotpDeviceShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserTotpDeviceShowMetaGlobalInput) SetIncludes(value string) *ActionUserTotpDeviceShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserTotpDeviceShowMetaGlobalInput) SetNo(value bool) *ActionUserTotpDeviceShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserTotpDeviceShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserTotpDeviceShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserTotpDeviceShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserTotpDeviceShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserTotpDeviceShowOutput is a type for action output parameters
type ActionUserTotpDeviceShowOutput struct {
	Confirmed bool   "json:\"confirmed\""
	CreatedAt string "json:\"created_at\""
	Enabled   bool   "json:\"enabled\""
	Id        int64  "json:\"id\""
	Label     string "json:\"label\""
	LastUseAt string "json:\"last_use_at\""
	UpdatedAt string "json:\"updated_at\""
	UseCount  int64  "json:\"use_count\""
}

// Type for action response, including envelope
type ActionUserTotpDeviceShowResponse struct {
	Action *ActionUserTotpDeviceShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TotpDevice *ActionUserTotpDeviceShowOutput "json:\"totp_device\""
	}

	// Action output without the namespace
	Output *ActionUserTotpDeviceShowOutput
}

// Prepare the action for invocation
func (action *ActionUserTotpDeviceShow) Prepare() *ActionUserTotpDeviceShowInvocation {
	return &ActionUserTotpDeviceShowInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/totp_devices/{totp_device_id}",
	}
}

// ActionUserTotpDeviceShowInvocation is used to configure action for invocation
type ActionUserTotpDeviceShowInvocation struct {
	// Pointer to the action
	Action *ActionUserTotpDeviceShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserTotpDeviceShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserTotpDeviceShowInvocation) SetPathParamInt(param string, value int64) *ActionUserTotpDeviceShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserTotpDeviceShowInvocation) SetPathParamString(param string, value string) *ActionUserTotpDeviceShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserTotpDeviceShowInvocation) NewMetaInput() *ActionUserTotpDeviceShowMetaGlobalInput {
	inv.MetaInput = &ActionUserTotpDeviceShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserTotpDeviceShowInvocation) SetMetaInput(input *ActionUserTotpDeviceShowMetaGlobalInput) *ActionUserTotpDeviceShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserTotpDeviceShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserTotpDeviceShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserTotpDeviceShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionUserTotpDeviceShowInvocation) Call() (*ActionUserTotpDeviceShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionUserTotpDeviceShowInvocation) CallContext(ctx context.Context) (*ActionUserTotpDeviceShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionUserTotpDeviceShowResponse{Action: inv.Action}
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
		if resp.Response.TotpDevice == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.TotpDevice
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionUserTotpDeviceShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionUserTotpDeviceShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionUserTotpDeviceShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
