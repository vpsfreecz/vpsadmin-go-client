package client

import (
	"strings"
)

// ActionOomReportRuleDelete is a type for action Oom_report_rule#Delete
type ActionOomReportRuleDelete struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportRuleDelete(client *Client) *ActionOomReportRuleDelete {
	return &ActionOomReportRuleDelete{
		Client: client,
	}
}

// ActionOomReportRuleDeleteMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportRuleDeleteMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportRuleDeleteMetaGlobalInput) SetIncludes(value string) *ActionOomReportRuleDeleteMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportRuleDeleteMetaGlobalInput) SetNo(value bool) *ActionOomReportRuleDeleteMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportRuleDeleteMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportRuleDeleteMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportRuleDeleteMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportRuleDeleteMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportRuleDeleteRequest is a type for the entire action request
type ActionOomReportRuleDeleteRequest struct {
	Meta map[string]interface{} `json:"_meta"`
}

// Type for action response, including envelope
type ActionOomReportRuleDeleteResponse struct {
	Action *ActionOomReportRuleDelete `json:"-"`
	*Envelope
}

// Prepare the action for invocation
func (action *ActionOomReportRuleDelete) Prepare() *ActionOomReportRuleDeleteInvocation {
	return &ActionOomReportRuleDeleteInvocation{
		Action: action,
		Path:   "/v7.0/oom_report_rules/{oom_report_rule_id}",
	}
}

// ActionOomReportRuleDeleteInvocation is used to configure action for invocation
type ActionOomReportRuleDeleteInvocation struct {
	// Pointer to the action
	Action *ActionOomReportRuleDelete

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOomReportRuleDeleteMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportRuleDeleteInvocation) SetPathParamInt(param string, value int64) *ActionOomReportRuleDeleteInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportRuleDeleteInvocation) SetPathParamString(param string, value string) *ActionOomReportRuleDeleteInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportRuleDeleteInvocation) NewMetaInput() *ActionOomReportRuleDeleteMetaGlobalInput {
	inv.MetaInput = &ActionOomReportRuleDeleteMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportRuleDeleteInvocation) SetMetaInput(input *ActionOomReportRuleDeleteMetaGlobalInput) *ActionOomReportRuleDeleteInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportRuleDeleteInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportRuleDeleteInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOomReportRuleDeleteInvocation) Call() (*ActionOomReportRuleDeleteResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionOomReportRuleDeleteInvocation) callAsBody() (*ActionOomReportRuleDeleteResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionOomReportRuleDeleteResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("DELETE", inv.Path, input, resp)
	return resp, err
}

func (inv *ActionOomReportRuleDeleteInvocation) makeAllInputParams() *ActionOomReportRuleDeleteRequest {
	return &ActionOomReportRuleDeleteRequest{
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionOomReportRuleDeleteInvocation) makeMetaInputParams() map[string]interface{} {
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
