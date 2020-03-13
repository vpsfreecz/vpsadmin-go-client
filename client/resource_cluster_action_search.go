package client

import (
)

// ActionClusterSearch is a type for action Cluster#Search
type ActionClusterSearch struct {
	// Pointer to client
	Client *Client
}

func NewActionClusterSearch(client *Client) *ActionClusterSearch {
	return &ActionClusterSearch{
		Client: client,
	}
}

// ActionClusterSearchMetaGlobalInput is a type for action global meta input parameters
type ActionClusterSearchMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionClusterSearchMetaGlobalInput) SetNo(value bool) *ActionClusterSearchMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterSearchMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterSearchMetaGlobalInput) SelectParameters(params ...string) *ActionClusterSearchMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterSearchMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterSearchInput is a type for action input parameters
type ActionClusterSearchInput struct {
	Value string `json:"value"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetValue sets parameter Value to value and selects it for sending
func (in *ActionClusterSearchInput) SetValue(value string) *ActionClusterSearchInput {
	in.Value = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Value"] = nil
	return in
}

// SelectParameters sets parameters from ActionClusterSearchInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionClusterSearchInput) SelectParameters(params ...string) *ActionClusterSearchInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionClusterSearchInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionClusterSearchRequest is a type for the entire action request
type ActionClusterSearchRequest struct {
	Cluster map[string]interface{} `json:"cluster"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionClusterSearchOutput is a type for action output parameters
type ActionClusterSearchOutput struct {
	Resource string `json:"resource"`
	Id int64 `json:"id"`
	Attribute string `json:"attribute"`
	Value string `json:"value"`
}


// Type for action response, including envelope
type ActionClusterSearchResponse struct {
	Action *ActionClusterSearch `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Clusters []*ActionClusterSearchOutput `json:"clusters"`
	}

	// Action output without the namespace
	Output []*ActionClusterSearchOutput
}


// Prepare the action for invocation
func (action *ActionClusterSearch) Prepare() *ActionClusterSearchInvocation {
	return &ActionClusterSearchInvocation{
		Action: action,
		Path: "/v6.0/cluster/search",
	}
}

// ActionClusterSearchInvocation is used to configure action for invocation
type ActionClusterSearchInvocation struct {
	// Pointer to the action
	Action *ActionClusterSearch

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionClusterSearchInput
	// Global meta input parameters
	MetaInput *ActionClusterSearchMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionClusterSearchInvocation) NewInput() *ActionClusterSearchInput {
	inv.Input = &ActionClusterSearchInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionClusterSearchInvocation) SetInput(input *ActionClusterSearchInput) *ActionClusterSearchInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionClusterSearchInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionClusterSearchInvocation) NewMetaInput() *ActionClusterSearchMetaGlobalInput {
	inv.MetaInput = &ActionClusterSearchMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionClusterSearchInvocation) SetMetaInput(input *ActionClusterSearchMetaGlobalInput) *ActionClusterSearchInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionClusterSearchInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionClusterSearchInvocation) Call() (*ActionClusterSearchResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionClusterSearchInvocation) callAsBody() (*ActionClusterSearchResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionClusterSearchResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Clusters
	}
	return resp, err
}




func (inv *ActionClusterSearchInvocation) makeAllInputParams() *ActionClusterSearchRequest {
	return &ActionClusterSearchRequest{
		Cluster: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionClusterSearchInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Value") {
			ret["value"] = inv.Input.Value
		}
	}

	return ret
}

func (inv *ActionClusterSearchInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
