package client

import (
	"strings"
)

// ActionIntegrityFactShow is a type for action Integrity_fact#Show
type ActionIntegrityFactShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIntegrityFactShow(client *Client) *ActionIntegrityFactShow {
	return &ActionIntegrityFactShow{
		Client: client,
	}
}

// ActionIntegrityFactShowMetaGlobalInput is a type for action global meta input parameters
type ActionIntegrityFactShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIntegrityFactShowMetaGlobalInput) SetIncludes(value string) *ActionIntegrityFactShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIntegrityFactShowMetaGlobalInput) SetNo(value bool) *ActionIntegrityFactShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityFactShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityFactShowMetaGlobalInput) SelectParameters(params ...string) *ActionIntegrityFactShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityFactShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIntegrityFactShowOutput is a type for action output parameters
type ActionIntegrityFactShowOutput struct {
	CreatedAt       string                           `json:"created_at"`
	Id              int64                            `json:"id"`
	IntegrityObject *ActionIntegrityObjectShowOutput `json:"integrity_object"`
	Message         string                           `json:"message"`
	Name            string                           `json:"name"`
	Severity        string                           `json:"severity"`
	Status          string                           `json:"status"`
}

// Type for action response, including envelope
type ActionIntegrityFactShowResponse struct {
	Action *ActionIntegrityFactShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IntegrityFact *ActionIntegrityFactShowOutput `json:"integrity_fact"`
	}

	// Action output without the namespace
	Output *ActionIntegrityFactShowOutput
}

// Prepare the action for invocation
func (action *ActionIntegrityFactShow) Prepare() *ActionIntegrityFactShowInvocation {
	return &ActionIntegrityFactShowInvocation{
		Action: action,
		Path:   "/v6.0/integrity_facts/{integrity_fact_id}",
	}
}

// ActionIntegrityFactShowInvocation is used to configure action for invocation
type ActionIntegrityFactShowInvocation struct {
	// Pointer to the action
	Action *ActionIntegrityFactShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIntegrityFactShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIntegrityFactShowInvocation) SetPathParamInt(param string, value int64) *ActionIntegrityFactShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIntegrityFactShowInvocation) SetPathParamString(param string, value string) *ActionIntegrityFactShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIntegrityFactShowInvocation) NewMetaInput() *ActionIntegrityFactShowMetaGlobalInput {
	inv.MetaInput = &ActionIntegrityFactShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIntegrityFactShowInvocation) SetMetaInput(input *ActionIntegrityFactShowMetaGlobalInput) *ActionIntegrityFactShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIntegrityFactShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionIntegrityFactShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIntegrityFactShowInvocation) Call() (*ActionIntegrityFactShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIntegrityFactShowInvocation) callAsQuery() (*ActionIntegrityFactShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIntegrityFactShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IntegrityFact
	}
	return resp, err
}

func (inv *ActionIntegrityFactShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
