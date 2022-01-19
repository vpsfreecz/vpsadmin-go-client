package client

import (
	"strings"
)

// ActionIntegrityCheckShow is a type for action Integrity_check#Show
type ActionIntegrityCheckShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIntegrityCheckShow(client *Client) *ActionIntegrityCheckShow {
	return &ActionIntegrityCheckShow{
		Client: client,
	}
}

// ActionIntegrityCheckShowMetaGlobalInput is a type for action global meta input parameters
type ActionIntegrityCheckShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIntegrityCheckShowMetaGlobalInput) SetIncludes(value string) *ActionIntegrityCheckShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIntegrityCheckShowMetaGlobalInput) SetNo(value bool) *ActionIntegrityCheckShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityCheckShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityCheckShowMetaGlobalInput) SelectParameters(params ...string) *ActionIntegrityCheckShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityCheckShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionIntegrityCheckShowOutput is a type for action output parameters
type ActionIntegrityCheckShowOutput struct {
	BrokenObjects   int64                 `json:"broken_objects"`
	CheckedFacts    int64                 `json:"checked_facts"`
	CheckedObjects  int64                 `json:"checked_objects"`
	CreatedAt       string                `json:"created_at"`
	FalseFacts      int64                 `json:"false_facts"`
	FinishedAt      string                `json:"finished_at"`
	Id              int64                 `json:"id"`
	IntegralObjects int64                 `json:"integral_objects"`
	Status          string                `json:"status"`
	TrueFacts       int64                 `json:"true_facts"`
	UpdatedAt       string                `json:"updated_at"`
	User            *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionIntegrityCheckShowResponse struct {
	Action *ActionIntegrityCheckShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IntegrityCheck *ActionIntegrityCheckShowOutput `json:"integrity_check"`
	}

	// Action output without the namespace
	Output *ActionIntegrityCheckShowOutput
}

// Prepare the action for invocation
func (action *ActionIntegrityCheckShow) Prepare() *ActionIntegrityCheckShowInvocation {
	return &ActionIntegrityCheckShowInvocation{
		Action: action,
		Path:   "/v6.0/integrity_checks/{integrity_check_id}",
	}
}

// ActionIntegrityCheckShowInvocation is used to configure action for invocation
type ActionIntegrityCheckShowInvocation struct {
	// Pointer to the action
	Action *ActionIntegrityCheckShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIntegrityCheckShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIntegrityCheckShowInvocation) SetPathParamInt(param string, value int64) *ActionIntegrityCheckShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIntegrityCheckShowInvocation) SetPathParamString(param string, value string) *ActionIntegrityCheckShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIntegrityCheckShowInvocation) NewMetaInput() *ActionIntegrityCheckShowMetaGlobalInput {
	inv.MetaInput = &ActionIntegrityCheckShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIntegrityCheckShowInvocation) SetMetaInput(input *ActionIntegrityCheckShowMetaGlobalInput) *ActionIntegrityCheckShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIntegrityCheckShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIntegrityCheckShowInvocation) Call() (*ActionIntegrityCheckShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIntegrityCheckShowInvocation) callAsQuery() (*ActionIntegrityCheckShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIntegrityCheckShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IntegrityCheck
	}
	return resp, err
}

func (inv *ActionIntegrityCheckShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
