package client

import ()

// ActionPaymentStatsEstimateIncome is a type for action Payment_stats#Estimate_income
type ActionPaymentStatsEstimateIncome struct {
	// Pointer to client
	Client *Client
}

func NewActionPaymentStatsEstimateIncome(client *Client) *ActionPaymentStatsEstimateIncome {
	return &ActionPaymentStatsEstimateIncome{
		Client: client,
	}
}

// ActionPaymentStatsEstimateIncomeMetaGlobalInput is a type for action global meta input parameters
type ActionPaymentStatsEstimateIncomeMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionPaymentStatsEstimateIncomeMetaGlobalInput) SetNo(value bool) *ActionPaymentStatsEstimateIncomeMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionPaymentStatsEstimateIncomeMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPaymentStatsEstimateIncomeMetaGlobalInput) SelectParameters(params ...string) *ActionPaymentStatsEstimateIncomeMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPaymentStatsEstimateIncomeMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPaymentStatsEstimateIncomeInput is a type for action input parameters
type ActionPaymentStatsEstimateIncomeInput struct {
	Duration int64  `json:"duration"`
	Month    int64  `json:"month"`
	Select   string `json:"select"`
	Year     int64  `json:"year"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDuration sets parameter Duration to value and selects it for sending
func (in *ActionPaymentStatsEstimateIncomeInput) SetDuration(value int64) *ActionPaymentStatsEstimateIncomeInput {
	in.Duration = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Duration"] = nil
	return in
}

// SetMonth sets parameter Month to value and selects it for sending
func (in *ActionPaymentStatsEstimateIncomeInput) SetMonth(value int64) *ActionPaymentStatsEstimateIncomeInput {
	in.Month = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Month"] = nil
	return in
}

// SetSelect sets parameter Select to value and selects it for sending
func (in *ActionPaymentStatsEstimateIncomeInput) SetSelect(value string) *ActionPaymentStatsEstimateIncomeInput {
	in.Select = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Select"] = nil
	return in
}

// SetYear sets parameter Year to value and selects it for sending
func (in *ActionPaymentStatsEstimateIncomeInput) SetYear(value int64) *ActionPaymentStatsEstimateIncomeInput {
	in.Year = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Year"] = nil
	return in
}

// SelectParameters sets parameters from ActionPaymentStatsEstimateIncomeInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPaymentStatsEstimateIncomeInput) SelectParameters(params ...string) *ActionPaymentStatsEstimateIncomeInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionPaymentStatsEstimateIncomeInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionPaymentStatsEstimateIncomeInput) UnselectParameters(params ...string) *ActionPaymentStatsEstimateIncomeInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionPaymentStatsEstimateIncomeInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPaymentStatsEstimateIncomeOutput is a type for action output parameters
type ActionPaymentStatsEstimateIncomeOutput struct {
	EstimatedIncome int64 `json:"estimated_income"`
	UserCount       int64 `json:"user_count"`
}

// Type for action response, including envelope
type ActionPaymentStatsEstimateIncomeResponse struct {
	Action *ActionPaymentStatsEstimateIncome `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PaymentStats *ActionPaymentStatsEstimateIncomeOutput `json:"payment_stats"`
	}

	// Action output without the namespace
	Output *ActionPaymentStatsEstimateIncomeOutput
}

// Prepare the action for invocation
func (action *ActionPaymentStatsEstimateIncome) Prepare() *ActionPaymentStatsEstimateIncomeInvocation {
	return &ActionPaymentStatsEstimateIncomeInvocation{
		Action: action,
		Path:   "/v6.0/payment_stat/estimate_income",
	}
}

// ActionPaymentStatsEstimateIncomeInvocation is used to configure action for invocation
type ActionPaymentStatsEstimateIncomeInvocation struct {
	// Pointer to the action
	Action *ActionPaymentStatsEstimateIncome

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionPaymentStatsEstimateIncomeInput
	// Global meta input parameters
	MetaInput *ActionPaymentStatsEstimateIncomeMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionPaymentStatsEstimateIncomeInvocation) NewInput() *ActionPaymentStatsEstimateIncomeInput {
	inv.Input = &ActionPaymentStatsEstimateIncomeInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionPaymentStatsEstimateIncomeInvocation) SetInput(input *ActionPaymentStatsEstimateIncomeInput) *ActionPaymentStatsEstimateIncomeInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionPaymentStatsEstimateIncomeInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionPaymentStatsEstimateIncomeInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionPaymentStatsEstimateIncomeInvocation) NewMetaInput() *ActionPaymentStatsEstimateIncomeMetaGlobalInput {
	inv.MetaInput = &ActionPaymentStatsEstimateIncomeMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionPaymentStatsEstimateIncomeInvocation) SetMetaInput(input *ActionPaymentStatsEstimateIncomeMetaGlobalInput) *ActionPaymentStatsEstimateIncomeInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionPaymentStatsEstimateIncomeInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionPaymentStatsEstimateIncomeInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionPaymentStatsEstimateIncomeInvocation) Call() (*ActionPaymentStatsEstimateIncomeResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionPaymentStatsEstimateIncomeInvocation) callAsQuery() (*ActionPaymentStatsEstimateIncomeResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionPaymentStatsEstimateIncomeResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.PaymentStats
	}
	return resp, err
}

func (inv *ActionPaymentStatsEstimateIncomeInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Duration") {
			ret["payment_stats[duration]"] = convertInt64ToString(inv.Input.Duration)
		}
		if inv.IsParameterSelected("Month") {
			ret["payment_stats[month]"] = convertInt64ToString(inv.Input.Month)
		}
		if inv.IsParameterSelected("Select") {
			ret["payment_stats[select]"] = inv.Input.Select
		}
		if inv.IsParameterSelected("Year") {
			ret["payment_stats[year]"] = convertInt64ToString(inv.Input.Year)
		}
	}
}

func (inv *ActionPaymentStatsEstimateIncomeInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
