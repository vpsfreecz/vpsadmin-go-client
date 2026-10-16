package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionClusterResourcePackageItemUpdate is a type for action Cluster_resource_package.Item#Update
type ActionClusterResourcePackageItemUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterResourcePackageItemUpdate(client *Client) *ActionClusterResourcePackageItemUpdate {
	return &ActionClusterResourcePackageItemUpdate{
		Client: client,
	}
}

// ActionClusterResourcePackageItemUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionClusterResourcePackageItemUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionClusterResourcePackageItemUpdateMetaGlobalInput) SetIncludes(value string) *ActionClusterResourcePackageItemUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterResourcePackageItemUpdateMetaGlobalInput) SetNo(value bool) *ActionClusterResourcePackageItemUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageItemUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageItemUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionClusterResourcePackageItemUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterResourcePackageItemUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageItemUpdateInput is a type for action input parameters
type ActionClusterResourcePackageItemUpdateInput struct {
	Value int64 "json:\"value\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetValue sets parameter Value to value and selects it for sending
func (in *ActionClusterResourcePackageItemUpdateInput) SetValue(value int64) *ActionClusterResourcePackageItemUpdateInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterResourcePackageItemUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageItemUpdateInput) SelectParameters(params ...string) *ActionClusterResourcePackageItemUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionClusterResourcePackageItemUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionClusterResourcePackageItemUpdateInput) UnselectParameters(params ...string) *ActionClusterResourcePackageItemUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionClusterResourcePackageItemUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterResourcePackageItemUpdateRequest is a type for the entire action request
type ActionClusterResourcePackageItemUpdateRequest struct {
	Item map[string]interface{} "json:\"item\""
	Meta map[string]interface{} "json:\"_meta\""
}

// ActionClusterResourcePackageItemUpdateOutput is a type for action output parameters
type ActionClusterResourcePackageItemUpdateOutput struct {
	ClusterResource *ActionClusterResourceShowOutput "json:\"cluster_resource\""
	Id              int64                            "json:\"id\""
	Value           int64                            "json:\"value\""
}

// Type for action response, including envelope
type ActionClusterResourcePackageItemUpdateResponse struct {
	Action *ActionClusterResourcePackageItemUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Item *ActionClusterResourcePackageItemUpdateOutput "json:\"item\""
	}

	// Action output without the namespace
	Output *ActionClusterResourcePackageItemUpdateOutput
}

// Prepare the action for invocation
func (action *ActionClusterResourcePackageItemUpdate) Prepare() *ActionClusterResourcePackageItemUpdateInvocation {
	return &ActionClusterResourcePackageItemUpdateInvocation{
		Action: action,
		Path:   "/v7.0/cluster_resource_packages/{cluster_resource_package_id}/items/{item_id}",
	}
}

// ActionClusterResourcePackageItemUpdateInvocation is used to configure action for invocation
type ActionClusterResourcePackageItemUpdateInvocation struct {
	// Pointer to the action
	Action *ActionClusterResourcePackageItemUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterResourcePackageItemUpdateInput
	// Global meta input parameters
	MetaInput *ActionClusterResourcePackageItemUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionClusterResourcePackageItemUpdateInvocation) SetPathParamInt(param string, value int64) *ActionClusterResourcePackageItemUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionClusterResourcePackageItemUpdateInvocation) SetPathParamString(param string, value string) *ActionClusterResourcePackageItemUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionClusterResourcePackageItemUpdateInvocation) NewInput() *ActionClusterResourcePackageItemUpdateInput {
	inv.Input = &ActionClusterResourcePackageItemUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionClusterResourcePackageItemUpdateInvocation) SetInput(input *ActionClusterResourcePackageItemUpdateInput) *ActionClusterResourcePackageItemUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterResourcePackageItemUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionClusterResourcePackageItemUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterResourcePackageItemUpdateInvocation) NewMetaInput() *ActionClusterResourcePackageItemUpdateMetaGlobalInput {
	inv.MetaInput = &ActionClusterResourcePackageItemUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterResourcePackageItemUpdateInvocation) SetMetaInput(input *ActionClusterResourcePackageItemUpdateMetaGlobalInput) *ActionClusterResourcePackageItemUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterResourcePackageItemUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionClusterResourcePackageItemUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionClusterResourcePackageItemUpdateInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionClusterResourcePackageItemUpdateInvocation) Call() (*ActionClusterResourcePackageItemUpdateResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionClusterResourcePackageItemUpdateInvocation) CallContext(ctx context.Context) (*ActionClusterResourcePackageItemUpdateResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionClusterResourcePackageItemUpdateResponse{Action: inv.Action}
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
		if resp.Response.Item == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Item
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionClusterResourcePackageItemUpdateInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionClusterResourcePackageItemUpdateInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "PUT", inv.Path, input)
}

func (inv *ActionClusterResourcePackageItemUpdateInvocation) makeAllInputParams() *ActionClusterResourcePackageItemUpdateRequest {
	return &ActionClusterResourcePackageItemUpdateRequest{
		Item: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionClusterResourcePackageItemUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Value") {
			ret["value"] = inv.Input.Value
		}
	}

	return ret
}

func (inv *ActionClusterResourcePackageItemUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
