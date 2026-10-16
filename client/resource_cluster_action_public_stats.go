package client

import (
	"context"
	"net/http"
)

// ActionClusterPublicStats is a type for action Cluster#Public_stats
type ActionClusterPublicStats struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterPublicStats(client *Client) *ActionClusterPublicStats {
	return &ActionClusterPublicStats{
		Client: client,
	}
}

// ActionClusterPublicStatsMetaGlobalInput is a type for action global meta input parameters
type ActionClusterPublicStatsMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterPublicStatsMetaGlobalInput) SetNo(value bool) *ActionClusterPublicStatsMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterPublicStatsMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterPublicStatsMetaGlobalInput) SelectParameters(params ...string) *ActionClusterPublicStatsMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterPublicStatsMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterPublicStatsOutput is a type for action output parameters
type ActionClusterPublicStatsOutput struct {
	Ipv4Left  int64 "json:\"ipv4_left\""
	UserCount int64 "json:\"user_count\""
	VpsCount  int64 "json:\"vps_count\""
}

// Type for action response, including envelope
type ActionClusterPublicStatsResponse struct {
	Action *ActionClusterPublicStats "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Cluster *ActionClusterPublicStatsOutput "json:\"cluster\""
	}

	// Action output without the namespace
	Output *ActionClusterPublicStatsOutput
}

// Call the action directly without any path or input parameters
func (action *ActionClusterPublicStats) Call() (*ActionClusterPublicStatsResponse, error) {
	return action.CallContext(context.Background())
}

// CallContext calls the action with cancellation and deadline support.
func (action *ActionClusterPublicStats) CallContext(ctx context.Context) (*ActionClusterPublicStatsResponse, error) {
	return action.Prepare().CallContext(ctx)
}

// CallRaw calls the action without decoding or classifying its response.
func (action *ActionClusterPublicStats) CallRaw(ctx context.Context) (*RawResponse, error) {
	return action.Prepare().CallRaw(ctx)
}

// Prepare the action for invocation
func (action *ActionClusterPublicStats) Prepare() *ActionClusterPublicStatsInvocation {
	return &ActionClusterPublicStatsInvocation{
		Action: action,
		Path:   "/v7.0/cluster/public_stats",
	}
}

// ActionClusterPublicStatsInvocation is used to configure action for invocation
type ActionClusterPublicStatsInvocation struct {
	// Pointer to the action
	Action *ActionClusterPublicStats

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterPublicStatsMetaGlobalInput
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterPublicStatsInvocation) NewMetaInput() *ActionClusterPublicStatsMetaGlobalInput {
	inv.MetaInput = &ActionClusterPublicStatsMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterPublicStatsInvocation) SetMetaInput(input *ActionClusterPublicStatsMetaGlobalInput) *ActionClusterPublicStatsInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterPublicStatsInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionClusterPublicStatsInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionClusterPublicStatsInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionClusterPublicStatsInvocation) Call() (*ActionClusterPublicStatsResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionClusterPublicStatsInvocation) CallContext(ctx context.Context) (*ActionClusterPublicStatsResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionClusterPublicStatsResponse{Action: inv.Action}
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
		if resp.Response.Cluster == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Cluster
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionClusterPublicStatsInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionClusterPublicStatsInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionClusterPublicStatsInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
