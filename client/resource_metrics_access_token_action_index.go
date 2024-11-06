package client

import ()

// ActionMetricsAccessTokenIndex is a type for action Metrics_access_token#Index
type ActionMetricsAccessTokenIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMetricsAccessTokenIndex(client *Client) *ActionMetricsAccessTokenIndex {
	return &ActionMetricsAccessTokenIndex{
		Client: client,
	}
}

// ActionMetricsAccessTokenIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMetricsAccessTokenIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMetricsAccessTokenIndexMetaGlobalInput) SetCount(value bool) *ActionMetricsAccessTokenIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMetricsAccessTokenIndexMetaGlobalInput) SetIncludes(value string) *ActionMetricsAccessTokenIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMetricsAccessTokenIndexMetaGlobalInput) SetNo(value bool) *ActionMetricsAccessTokenIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMetricsAccessTokenIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMetricsAccessTokenIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMetricsAccessTokenIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMetricsAccessTokenIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMetricsAccessTokenIndexInput is a type for action input parameters
type ActionMetricsAccessTokenIndexInput struct {
	FromId int64 `json:"from_id"`
	Limit  int64 `json:"limit"`
	User   int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionMetricsAccessTokenIndexInput) SetFromId(value int64) *ActionMetricsAccessTokenIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMetricsAccessTokenIndexInput) SetLimit(value int64) *ActionMetricsAccessTokenIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionMetricsAccessTokenIndexInput) SetUser(value int64) *ActionMetricsAccessTokenIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionMetricsAccessTokenIndexInput) SetUserNil(set bool) *ActionMetricsAccessTokenIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionMetricsAccessTokenIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMetricsAccessTokenIndexInput) SelectParameters(params ...string) *ActionMetricsAccessTokenIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMetricsAccessTokenIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMetricsAccessTokenIndexInput) UnselectParameters(params ...string) *ActionMetricsAccessTokenIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMetricsAccessTokenIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMetricsAccessTokenIndexOutput is a type for action output parameters
type ActionMetricsAccessTokenIndexOutput struct {
	AccessToken  string                `json:"access_token"`
	CreatedAt    string                `json:"created_at"`
	Id           int64                 `json:"id"`
	LastUse      string                `json:"last_use"`
	MetricPrefix string                `json:"metric_prefix"`
	UpdatedAt    string                `json:"updated_at"`
	UseCount     int64                 `json:"use_count"`
	User         *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionMetricsAccessTokenIndexResponse struct {
	Action *ActionMetricsAccessTokenIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MetricsAccessTokens []*ActionMetricsAccessTokenIndexOutput `json:"metrics_access_tokens"`
	}

	// Action output without the namespace
	Output []*ActionMetricsAccessTokenIndexOutput
}

// Prepare the action for invocation
func (action *ActionMetricsAccessTokenIndex) Prepare() *ActionMetricsAccessTokenIndexInvocation {
	return &ActionMetricsAccessTokenIndexInvocation{
		Action: action,
		Path:   "/v7.0/metrics_access_tokens",
	}
}

// ActionMetricsAccessTokenIndexInvocation is used to configure action for invocation
type ActionMetricsAccessTokenIndexInvocation struct {
	// Pointer to the action
	Action *ActionMetricsAccessTokenIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMetricsAccessTokenIndexInput
	// Global meta input parameters
	MetaInput *ActionMetricsAccessTokenIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMetricsAccessTokenIndexInvocation) NewInput() *ActionMetricsAccessTokenIndexInput {
	inv.Input = &ActionMetricsAccessTokenIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMetricsAccessTokenIndexInvocation) SetInput(input *ActionMetricsAccessTokenIndexInput) *ActionMetricsAccessTokenIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMetricsAccessTokenIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMetricsAccessTokenIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMetricsAccessTokenIndexInvocation) NewMetaInput() *ActionMetricsAccessTokenIndexMetaGlobalInput {
	inv.MetaInput = &ActionMetricsAccessTokenIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMetricsAccessTokenIndexInvocation) SetMetaInput(input *ActionMetricsAccessTokenIndexMetaGlobalInput) *ActionMetricsAccessTokenIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMetricsAccessTokenIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMetricsAccessTokenIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMetricsAccessTokenIndexInvocation) Call() (*ActionMetricsAccessTokenIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMetricsAccessTokenIndexInvocation) callAsQuery() (*ActionMetricsAccessTokenIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMetricsAccessTokenIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MetricsAccessTokens
	}
	return resp, err
}

func (inv *ActionMetricsAccessTokenIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["metrics_access_token[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["metrics_access_token[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["metrics_access_token[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionMetricsAccessTokenIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
