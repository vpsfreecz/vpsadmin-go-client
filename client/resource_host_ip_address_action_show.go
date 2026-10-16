package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionHostIpAddressShow is a type for action Host_ip_address#Show
type ActionHostIpAddressShow struct {
	// Pointer to client
	Client *Client
}

func NewActionHostIpAddressShow(client *Client) *ActionHostIpAddressShow {
	return &ActionHostIpAddressShow{
		Client: client,
	}
}

// ActionHostIpAddressShowMetaGlobalInput is a type for action global meta input parameters
type ActionHostIpAddressShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHostIpAddressShowMetaGlobalInput) SetIncludes(value string) *ActionHostIpAddressShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHostIpAddressShowMetaGlobalInput) SetNo(value bool) *ActionHostIpAddressShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionHostIpAddressShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHostIpAddressShowMetaGlobalInput) SelectParameters(params ...string) *ActionHostIpAddressShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHostIpAddressShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionHostIpAddressShowOutput is a type for action output parameters
type ActionHostIpAddressShowOutput struct {
	Addr               string                     "json:\"addr\""
	Assigned           bool                       "json:\"assigned\""
	Id                 int64                      "json:\"id\""
	IpAddress          *ActionIpAddressShowOutput "json:\"ip_address\""
	ReverseRecordValue string                     "json:\"reverse_record_value\""
	UserCreated        bool                       "json:\"user_created\""
}

// Type for action response, including envelope
type ActionHostIpAddressShowResponse struct {
	Action *ActionHostIpAddressShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HostIpAddress *ActionHostIpAddressShowOutput "json:\"host_ip_address\""
	}

	// Action output without the namespace
	Output *ActionHostIpAddressShowOutput
}

// Prepare the action for invocation
func (action *ActionHostIpAddressShow) Prepare() *ActionHostIpAddressShowInvocation {
	return &ActionHostIpAddressShowInvocation{
		Action: action,
		Path:   "/v7.0/host_ip_addresses/{host_ip_address_id}",
	}
}

// ActionHostIpAddressShowInvocation is used to configure action for invocation
type ActionHostIpAddressShowInvocation struct {
	// Pointer to the action
	Action *ActionHostIpAddressShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionHostIpAddressShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionHostIpAddressShowInvocation) SetPathParamInt(param string, value int64) *ActionHostIpAddressShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionHostIpAddressShowInvocation) SetPathParamString(param string, value string) *ActionHostIpAddressShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHostIpAddressShowInvocation) NewMetaInput() *ActionHostIpAddressShowMetaGlobalInput {
	inv.MetaInput = &ActionHostIpAddressShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHostIpAddressShowInvocation) SetMetaInput(input *ActionHostIpAddressShowMetaGlobalInput) *ActionHostIpAddressShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHostIpAddressShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionHostIpAddressShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionHostIpAddressShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionHostIpAddressShowInvocation) Call() (*ActionHostIpAddressShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionHostIpAddressShowInvocation) CallContext(ctx context.Context) (*ActionHostIpAddressShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionHostIpAddressShowResponse{Action: inv.Action}
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
		if resp.Response.HostIpAddress == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.HostIpAddress
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionHostIpAddressShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionHostIpAddressShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionHostIpAddressShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
