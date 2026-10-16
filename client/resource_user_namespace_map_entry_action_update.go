package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionUserNamespaceMapEntryUpdate is a type for action User_namespace_map.Entry#Update
type ActionUserNamespaceMapEntryUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionUserNamespaceMapEntryUpdate(client *Client) *ActionUserNamespaceMapEntryUpdate {
	return &ActionUserNamespaceMapEntryUpdate{
		Client: client,
	}
}

// ActionUserNamespaceMapEntryUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionUserNamespaceMapEntryUpdateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) SetIncludes(value string) *ActionUserNamespaceMapEntryUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) SetNo(value bool) *ActionUserNamespaceMapEntryUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapEntryUpdateInput is a type for action input parameters
type ActionUserNamespaceMapEntryUpdateInput struct {
	Count int64 "json:\"count\""
	NsId  int64 "json:\"ns_id\""
	VpsId int64 "json:\"vps_id\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateInput) SetCount(value int64) *ActionUserNamespaceMapEntryUpdateInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetNsId sets parameter NsId to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateInput) SetNsId(value int64) *ActionUserNamespaceMapEntryUpdateInput {
	in.NsId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["NsId"] = nil
	return in
}

// SetVpsId sets parameter VpsId to value and selects it for sending
func (in *ActionUserNamespaceMapEntryUpdateInput) SetVpsId(value int64) *ActionUserNamespaceMapEntryUpdateInput {
	in.VpsId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["VpsId"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserNamespaceMapEntryUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryUpdateInput) SelectParameters(params ...string) *ActionUserNamespaceMapEntryUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserNamespaceMapEntryUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserNamespaceMapEntryUpdateInput) UnselectParameters(params ...string) *ActionUserNamespaceMapEntryUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserNamespaceMapEntryUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserNamespaceMapEntryUpdateRequest is a type for the entire action request
type ActionUserNamespaceMapEntryUpdateRequest struct {
	Entry map[string]interface{} "json:\"entry\""
	Meta  map[string]interface{} "json:\"_meta\""
}

// ActionUserNamespaceMapEntryUpdateOutput is a type for action output parameters
type ActionUserNamespaceMapEntryUpdateOutput struct {
	Count int64  "json:\"count\""
	Id    int64  "json:\"id\""
	Kind  string "json:\"kind\""
	NsId  int64  "json:\"ns_id\""
	VpsId int64  "json:\"vps_id\""
}

// Type for action response, including envelope
type ActionUserNamespaceMapEntryUpdateResponse struct {
	Action *ActionUserNamespaceMapEntryUpdate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Entry *ActionUserNamespaceMapEntryUpdateOutput "json:\"entry\""
	}

	// Action output without the namespace
	Output *ActionUserNamespaceMapEntryUpdateOutput
}

// Prepare the action for invocation
func (action *ActionUserNamespaceMapEntryUpdate) Prepare() *ActionUserNamespaceMapEntryUpdateInvocation {
	return &ActionUserNamespaceMapEntryUpdateInvocation{
		Action: action,
		Path:   "/v7.0/user_namespace_maps/{user_namespace_map_id}/entries/{entry_id}",
	}
}

// ActionUserNamespaceMapEntryUpdateInvocation is used to configure action for invocation
type ActionUserNamespaceMapEntryUpdateInvocation struct {
	// Pointer to the action
	Action *ActionUserNamespaceMapEntryUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserNamespaceMapEntryUpdateInput
	// Global meta input parameters
	MetaInput *ActionUserNamespaceMapEntryUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) SetPathParamInt(param string, value int64) *ActionUserNamespaceMapEntryUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) SetPathParamString(param string, value string) *ActionUserNamespaceMapEntryUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) NewInput() *ActionUserNamespaceMapEntryUpdateInput {
	inv.Input = &ActionUserNamespaceMapEntryUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) SetInput(input *ActionUserNamespaceMapEntryUpdateInput) *ActionUserNamespaceMapEntryUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) NewMetaInput() *ActionUserNamespaceMapEntryUpdateMetaGlobalInput {
	inv.MetaInput = &ActionUserNamespaceMapEntryUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) SetMetaInput(input *ActionUserNamespaceMapEntryUpdateMetaGlobalInput) *ActionUserNamespaceMapEntryUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserNamespaceMapEntryUpdateInvocation) validate() error {
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
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) Call() (*ActionUserNamespaceMapEntryUpdateResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) CallContext(ctx context.Context) (*ActionUserNamespaceMapEntryUpdateResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionUserNamespaceMapEntryUpdateResponse{Action: inv.Action}
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
		if resp.Response.Entry == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Entry
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionUserNamespaceMapEntryUpdateInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionUserNamespaceMapEntryUpdateInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "PUT", inv.Path, input)
}

func (inv *ActionUserNamespaceMapEntryUpdateInvocation) makeAllInputParams() *ActionUserNamespaceMapEntryUpdateRequest {
	return &ActionUserNamespaceMapEntryUpdateRequest{
		Entry: inv.makeInputParams(),
		Meta:  inv.makeMetaInputParams(),
	}
}

func (inv *ActionUserNamespaceMapEntryUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Count") {
			ret["count"] = inv.Input.Count
		}
		if inv.IsParameterSelected("NsId") {
			ret["ns_id"] = inv.Input.NsId
		}
		if inv.IsParameterSelected("VpsId") {
			ret["vps_id"] = inv.Input.VpsId
		}
	}

	return ret
}

func (inv *ActionUserNamespaceMapEntryUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
