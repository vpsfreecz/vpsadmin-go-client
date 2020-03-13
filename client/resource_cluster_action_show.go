package client

import (
)

// ActionClusterShow is a type for action Cluster#Show
type ActionClusterShow struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterShow(client *Client) *ActionClusterShow {
	return &ActionClusterShow{
		Client: client,
	}
}

// ActionClusterShowMetaGlobalInput is a type for action global meta input parameters
type ActionClusterShowMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterShowMetaGlobalInput) SetNo(value bool) *ActionClusterShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterShowMetaGlobalInput) SelectParameters(params ...string) *ActionClusterShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionClusterShowOutput is a type for action output parameters
type ActionClusterShowOutput struct {
	MaintenanceLock bool `json:"maintenance_lock"`
	MaintenanceLockReason string `json:"maintenance_lock_reason"`
}


// Type for action response, including envelope
type ActionClusterShowResponse struct {
	Action *ActionClusterShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Cluster *ActionClusterShowOutput `json:"cluster"`
	}

	// Action output without the namespace
	Output *ActionClusterShowOutput
}

// Call the action directly without any path or input parameters
func (action *ActionClusterShow) Call() (*ActionClusterShowResponse, error) {
	return action.Prepare().Call()
}

// Prepare the action for invocation
func (action *ActionClusterShow) Prepare() *ActionClusterShowInvocation {
	return &ActionClusterShowInvocation{
		Action: action,
		Path: "/v6.0/cluster",
	}
}

// ActionClusterShowInvocation is used to configure action for invocation
type ActionClusterShowInvocation struct {
	// Pointer to the action
	Action *ActionClusterShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionClusterShowMetaGlobalInput
}


// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterShowInvocation) NewMetaInput() *ActionClusterShowMetaGlobalInput {
	inv.MetaInput = &ActionClusterShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterShowInvocation) SetMetaInput(input *ActionClusterShowMetaGlobalInput) *ActionClusterShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterShowInvocation) Call() (*ActionClusterShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionClusterShowInvocation) callAsQuery() (*ActionClusterShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionClusterShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Cluster
	}
	return resp, err
}




func (inv *ActionClusterShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

