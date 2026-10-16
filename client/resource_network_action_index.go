package client

import (
	"context"
	"net/http"
)

// ActionNetworkIndex is a type for action Network#Index
type ActionNetworkIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNetworkIndex(client *Client) *ActionNetworkIndex {
	return &ActionNetworkIndex{
		Client: client,
	}
}

// ActionNetworkIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNetworkIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNetworkIndexMetaGlobalInput) SetCount(value bool) *ActionNetworkIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNetworkIndexMetaGlobalInput) SetIncludes(value string) *ActionNetworkIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNetworkIndexMetaGlobalInput) SetNo(value bool) *ActionNetworkIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNetworkIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNetworkIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkIndexInput is a type for action input parameters
type ActionNetworkIndexInput struct {
	Enabled   bool   "json:\"enabled\""
	FromId    int64  "json:\"from_id\""
	Limit     int64  "json:\"limit\""
	Location  int64  "json:\"location\""
	Purpose   string "json:\"purpose\""
	UsableFor string "json:\"usable_for\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionNetworkIndexInput) SetEnabled(value bool) *ActionNetworkIndexInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNetworkIndexInput) SetFromId(value int64) *ActionNetworkIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNetworkIndexInput) SetLimit(value int64) *ActionNetworkIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionNetworkIndexInput) SetLocation(value int64) *ActionNetworkIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}

// SetPurpose sets parameter Purpose to value and selects it for sending
func (in *ActionNetworkIndexInput) SetPurpose(value string) *ActionNetworkIndexInput {
	in.Purpose = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Purpose"] = nil
	return in
}

// SetUsableFor sets parameter UsableFor to value and selects it for sending
func (in *ActionNetworkIndexInput) SetUsableFor(value string) *ActionNetworkIndexInput {
	in.UsableFor = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UsableFor"] = nil
	return in
}

// SelectParameters sets parameters from ActionNetworkIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNetworkIndexInput) SelectParameters(params ...string) *ActionNetworkIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNetworkIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNetworkIndexInput) UnselectParameters(params ...string) *ActionNetworkIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNetworkIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNetworkIndexOutput is a type for action output parameters
type ActionNetworkIndexOutput struct {
	Address          string                    "json:\"address\""
	Assigned         int64                     "json:\"assigned\""
	AvailableToUsers int64                     "json:\"available_to_users\""
	Enabled          bool                      "json:\"enabled\""
	Id               int64                     "json:\"id\""
	IpVersion        int64                     "json:\"ip_version\""
	Label            string                    "json:\"label\""
	Managed          bool                      "json:\"managed\""
	Owned            int64                     "json:\"owned\""
	OwnedUnassigned  int64                     "json:\"owned_unassigned\""
	Prefix           int64                     "json:\"prefix\""
	PrimaryLocation  *ActionLocationShowOutput "json:\"primary_location\""
	Purpose          string                    "json:\"purpose\""
	Role             string                    "json:\"role\""
	Size             int64                     "json:\"size\""
	SplitAccess      string                    "json:\"split_access\""
	SplitPrefix      int64                     "json:\"split_prefix\""
	Taken            int64                     "json:\"taken\""
	Used             int64                     "json:\"used\""
}

// Type for action response, including envelope
type ActionNetworkIndexResponse struct {
	Action *ActionNetworkIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Networks []*ActionNetworkIndexOutput "json:\"networks\""
	}

	// Action output without the namespace
	Output []*ActionNetworkIndexOutput
}

// Prepare the action for invocation
func (action *ActionNetworkIndex) Prepare() *ActionNetworkIndexInvocation {
	return &ActionNetworkIndexInvocation{
		Action: action,
		Path:   "/v7.0/networks",
	}
}

// ActionNetworkIndexInvocation is used to configure action for invocation
type ActionNetworkIndexInvocation struct {
	// Pointer to the action
	Action *ActionNetworkIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNetworkIndexInput
	// Global meta input parameters
	MetaInput *ActionNetworkIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNetworkIndexInvocation) NewInput() *ActionNetworkIndexInput {
	inv.Input = &ActionNetworkIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNetworkIndexInvocation) SetInput(input *ActionNetworkIndexInput) *ActionNetworkIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNetworkIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNetworkIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNetworkIndexInvocation) NewMetaInput() *ActionNetworkIndexMetaGlobalInput {
	inv.MetaInput = &ActionNetworkIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNetworkIndexInvocation) SetMetaInput(input *ActionNetworkIndexMetaGlobalInput) *ActionNetworkIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNetworkIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNetworkIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNetworkIndexInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("Location") {
			if !inv.IsParameterNil("Location") {
				if inv.Input.Location < 0 {
					verr.Add("location", t(inv.Action.Client, "validation.invalid_resource_id", nil))
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionNetworkIndexInvocation) Call() (*ActionNetworkIndexResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionNetworkIndexInvocation) CallContext(ctx context.Context) (*ActionNetworkIndexResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionNetworkIndexResponse{Action: inv.Action}
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
		if resp.Response.Networks == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Networks
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionNetworkIndexInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionNetworkIndexInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionNetworkIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["network[enabled]"] = convertBoolToString(inv.Input.Enabled)
		}
		if inv.IsParameterSelected("FromId") {
			ret["network[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["network[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["network[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Purpose") {
			ret["network[purpose]"] = inv.Input.Purpose
		}
		if inv.IsParameterSelected("UsableFor") {
			ret["network[usable_for]"] = inv.Input.UsableFor
		}
	}
}

func (inv *ActionNetworkIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
