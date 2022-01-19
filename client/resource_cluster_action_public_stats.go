package client

import (
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
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
	Ipv4Left int64 `json:"ipv4_left"`
	UserCount int64 `json:"user_count"`
	VpsCount int64 `json:"vps_count"`
}


// Type for action response, including envelope
type ActionClusterPublicStatsResponse struct {
	Action *ActionClusterPublicStats `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Cluster *ActionClusterPublicStatsOutput `json:"cluster"`
	}

	// Action output without the namespace
	Output *ActionClusterPublicStatsOutput
}

// Call the action directly without any path or input parameters
func (action *ActionClusterPublicStats) Call() (*ActionClusterPublicStatsResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *ActionClusterPublicStats) Prepare() *ActionClusterPublicStatsInvocation {
	return &ActionClusterPublicStatsInvocation{
		Action: action,
		Path: "/v6.0/cluster/public_stats",
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

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterPublicStatsInvocation) Call() (*ActionClusterPublicStatsResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionClusterPublicStatsInvocation) callAsQuery() (*ActionClusterPublicStatsResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionClusterPublicStatsResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Cluster
	}
	return resp, err
}




func (inv *ActionClusterPublicStatsInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

