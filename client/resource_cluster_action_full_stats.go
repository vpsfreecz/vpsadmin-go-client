package client

import (
	"context"
	"net/http"
)

// ActionClusterFullStats is a type for action Cluster#Full_stats
type ActionClusterFullStats struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterFullStats(client *Client) *ActionClusterFullStats {
	return &ActionClusterFullStats{
		Client: client,
	}
}

// ActionClusterFullStatsMetaGlobalInput is a type for action global meta input parameters
type ActionClusterFullStatsMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterFullStatsMetaGlobalInput) SetNo(value bool) *ActionClusterFullStatsMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterFullStatsMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterFullStatsMetaGlobalInput) SelectParameters(params ...string) *ActionClusterFullStatsMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterFullStatsMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterFullStatsOutput is a type for action output parameters
type ActionClusterFullStatsOutput struct {
	Ipv4Count     int64 "json:\"ipv4_count\""
	Ipv4Used      int64 "json:\"ipv4_used\""
	NodeCount     int64 "json:\"node_count\""
	NodesOnline   int64 "json:\"nodes_online\""
	UserActive    int64 "json:\"user_active\""
	UserCount     int64 "json:\"user_count\""
	UserDeleted   int64 "json:\"user_deleted\""
	UserSuspended int64 "json:\"user_suspended\""
	VpsCount      int64 "json:\"vps_count\""
	VpsDeleted    int64 "json:\"vps_deleted\""
	VpsRunning    int64 "json:\"vps_running\""
	VpsStopped    int64 "json:\"vps_stopped\""
	VpsSuspended  int64 "json:\"vps_suspended\""
}

// Type for action response, including envelope
type ActionClusterFullStatsResponse struct {
	Action *ActionClusterFullStats "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Cluster *ActionClusterFullStatsOutput "json:\"cluster\""
	}

	// Action output without the namespace
	Output *ActionClusterFullStatsOutput
}

// Call the action directly without any path or input parameters
func (action *ActionClusterFullStats) Call() (*ActionClusterFullStatsResponse, error) {
	return action.CallContext(context.Background())
}

// CallContext calls the action with cancellation and deadline support.
func (action *ActionClusterFullStats) CallContext(ctx context.Context) (*ActionClusterFullStatsResponse, error) {
	return action.Prepare().CallContext(ctx)
}

// CallRaw calls the action without decoding or classifying its response.
func (action *ActionClusterFullStats) CallRaw(ctx context.Context) (*RawResponse, error) {
	return action.Prepare().CallRaw(ctx)
}

// Prepare the action for invocation
func (action *ActionClusterFullStats) Prepare() *ActionClusterFullStatsInvocation {
	return &ActionClusterFullStatsInvocation{
		Action: action,
		Path:   "/v7.0/cluster/full_stats",
	}
}

// ActionClusterFullStatsInvocation is used to configure action for invocation
type ActionClusterFullStatsInvocation struct {
	// Pointer to the action
	Action *ActionClusterFullStats

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterFullStatsMetaGlobalInput
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterFullStatsInvocation) NewMetaInput() *ActionClusterFullStatsMetaGlobalInput {
	inv.MetaInput = &ActionClusterFullStatsMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterFullStatsInvocation) SetMetaInput(input *ActionClusterFullStatsMetaGlobalInput) *ActionClusterFullStatsInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterFullStatsInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionClusterFullStatsInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionClusterFullStatsInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionClusterFullStatsInvocation) Call() (*ActionClusterFullStatsResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionClusterFullStatsInvocation) CallContext(ctx context.Context) (*ActionClusterFullStatsResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionClusterFullStatsResponse{Action: inv.Action}
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
func (inv *ActionClusterFullStatsInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionClusterFullStatsInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionClusterFullStatsInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
