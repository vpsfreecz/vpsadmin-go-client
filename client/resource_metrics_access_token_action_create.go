package client

import ()

// ActionMetricsAccessTokenCreate is a type for action Metrics_access_token#Create
type ActionMetricsAccessTokenCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionMetricsAccessTokenCreate(client *Client) *ActionMetricsAccessTokenCreate {
	return &ActionMetricsAccessTokenCreate{
		Client: client,
	}
}

// ActionMetricsAccessTokenCreateMetaGlobalInput is a type for action global meta input parameters
type ActionMetricsAccessTokenCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMetricsAccessTokenCreateMetaGlobalInput) SetIncludes(value string) *ActionMetricsAccessTokenCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMetricsAccessTokenCreateMetaGlobalInput) SetNo(value bool) *ActionMetricsAccessTokenCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMetricsAccessTokenCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMetricsAccessTokenCreateMetaGlobalInput) SelectParameters(params ...string) *ActionMetricsAccessTokenCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMetricsAccessTokenCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMetricsAccessTokenCreateInput is a type for action input parameters
type ActionMetricsAccessTokenCreateInput struct {
	MetricPrefix string `json:"metric_prefix"`
	User         int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetMetricPrefix sets parameter MetricPrefix to value and selects it for sending
func (in *ActionMetricsAccessTokenCreateInput) SetMetricPrefix(value string) *ActionMetricsAccessTokenCreateInput {
	in.MetricPrefix = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MetricPrefix"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionMetricsAccessTokenCreateInput) SetUser(value int64) *ActionMetricsAccessTokenCreateInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionMetricsAccessTokenCreateInput) SetUserNil(set bool) *ActionMetricsAccessTokenCreateInput {
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

// SelectParameters sets parameters from ActionMetricsAccessTokenCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMetricsAccessTokenCreateInput) SelectParameters(params ...string) *ActionMetricsAccessTokenCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMetricsAccessTokenCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMetricsAccessTokenCreateInput) UnselectParameters(params ...string) *ActionMetricsAccessTokenCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMetricsAccessTokenCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMetricsAccessTokenCreateRequest is a type for the entire action request
type ActionMetricsAccessTokenCreateRequest struct {
	MetricsAccessToken map[string]interface{} `json:"metrics_access_token"`
	Meta               map[string]interface{} `json:"_meta"`
}

// ActionMetricsAccessTokenCreateOutput is a type for action output parameters
type ActionMetricsAccessTokenCreateOutput struct {
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
type ActionMetricsAccessTokenCreateResponse struct {
	Action *ActionMetricsAccessTokenCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MetricsAccessToken *ActionMetricsAccessTokenCreateOutput `json:"metrics_access_token"`
	}

	// Action output without the namespace
	Output *ActionMetricsAccessTokenCreateOutput
}

// Prepare the action for invocation
func (action *ActionMetricsAccessTokenCreate) Prepare() *ActionMetricsAccessTokenCreateInvocation {
	return &ActionMetricsAccessTokenCreateInvocation{
		Action: action,
		Path:   "/v7.0/metrics_access_tokens",
	}
}

// ActionMetricsAccessTokenCreateInvocation is used to configure action for invocation
type ActionMetricsAccessTokenCreateInvocation struct {
	// Pointer to the action
	Action *ActionMetricsAccessTokenCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMetricsAccessTokenCreateInput
	// Global meta input parameters
	MetaInput *ActionMetricsAccessTokenCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMetricsAccessTokenCreateInvocation) NewInput() *ActionMetricsAccessTokenCreateInput {
	inv.Input = &ActionMetricsAccessTokenCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMetricsAccessTokenCreateInvocation) SetInput(input *ActionMetricsAccessTokenCreateInput) *ActionMetricsAccessTokenCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMetricsAccessTokenCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMetricsAccessTokenCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMetricsAccessTokenCreateInvocation) NewMetaInput() *ActionMetricsAccessTokenCreateMetaGlobalInput {
	inv.MetaInput = &ActionMetricsAccessTokenCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMetricsAccessTokenCreateInvocation) SetMetaInput(input *ActionMetricsAccessTokenCreateMetaGlobalInput) *ActionMetricsAccessTokenCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMetricsAccessTokenCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMetricsAccessTokenCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMetricsAccessTokenCreateInvocation) Call() (*ActionMetricsAccessTokenCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionMetricsAccessTokenCreateInvocation) callAsBody() (*ActionMetricsAccessTokenCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionMetricsAccessTokenCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MetricsAccessToken
	}
	return resp, err
}

func (inv *ActionMetricsAccessTokenCreateInvocation) makeAllInputParams() *ActionMetricsAccessTokenCreateRequest {
	return &ActionMetricsAccessTokenCreateRequest{
		MetricsAccessToken: inv.makeInputParams(),
		Meta:               inv.makeMetaInputParams(),
	}
}

func (inv *ActionMetricsAccessTokenCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("MetricPrefix") {
			ret["metric_prefix"] = inv.Input.MetricPrefix
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["user"] = nil
			} else {
				ret["user"] = inv.Input.User
			}
		}
	}

	return ret
}

func (inv *ActionMetricsAccessTokenCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
