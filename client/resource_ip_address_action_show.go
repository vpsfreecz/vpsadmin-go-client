package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionIpAddressShow is a type for action Ip_address#Show
type ActionIpAddressShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIpAddressShow(client *Client) *ActionIpAddressShow {
	return &ActionIpAddressShow{
		Client: client,
	}
}

// ActionIpAddressShowMetaGlobalInput is a type for action global meta input parameters
type ActionIpAddressShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIpAddressShowMetaGlobalInput) SetIncludes(value string) *ActionIpAddressShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIpAddressShowMetaGlobalInput) SetNo(value bool) *ActionIpAddressShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIpAddressShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIpAddressShowMetaGlobalInput) SelectParameters(params ...string) *ActionIpAddressShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIpAddressShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIpAddressShowOutput is a type for action output parameters
type ActionIpAddressShowOutput struct {
	Addr               string                            "json:\"addr\""
	ChargedEnvironment *ActionEnvironmentShowOutput      "json:\"charged_environment\""
	Id                 int64                             "json:\"id\""
	Network            *ActionNetworkShowOutput          "json:\"network\""
	NetworkInterface   *ActionNetworkInterfaceShowOutput "json:\"network_interface\""
	Prefix             int64                             "json:\"prefix\""
	RouteVia           *ActionHostIpAddressShowOutput    "json:\"route_via\""
	Size               int64                             "json:\"size\""
	User               *ActionUserShowOutput             "json:\"user\""
}

// Type for action response, including envelope
type ActionIpAddressShowResponse struct {
	Action *ActionIpAddressShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IpAddress *ActionIpAddressShowOutput "json:\"ip_address\""
	}

	// Action output without the namespace
	Output *ActionIpAddressShowOutput
}

// Prepare the action for invocation
func (action *ActionIpAddressShow) Prepare() *ActionIpAddressShowInvocation {
	return &ActionIpAddressShowInvocation{
		Action: action,
		Path:   "/v7.0/ip_addresses/{ip_address_id}",
	}
}

// ActionIpAddressShowInvocation is used to configure action for invocation
type ActionIpAddressShowInvocation struct {
	// Pointer to the action
	Action *ActionIpAddressShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIpAddressShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIpAddressShowInvocation) SetPathParamInt(param string, value int64) *ActionIpAddressShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIpAddressShowInvocation) SetPathParamString(param string, value string) *ActionIpAddressShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIpAddressShowInvocation) NewMetaInput() *ActionIpAddressShowMetaGlobalInput {
	inv.MetaInput = &ActionIpAddressShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIpAddressShowInvocation) SetMetaInput(input *ActionIpAddressShowMetaGlobalInput) *ActionIpAddressShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIpAddressShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIpAddressShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionIpAddressShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionIpAddressShowInvocation) Call() (*ActionIpAddressShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionIpAddressShowInvocation) CallContext(ctx context.Context) (*ActionIpAddressShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionIpAddressShowResponse{Action: inv.Action}
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
		if resp.Response.IpAddress == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.IpAddress
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionIpAddressShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionIpAddressShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionIpAddressShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
