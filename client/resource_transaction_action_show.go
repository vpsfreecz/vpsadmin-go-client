package client

import (
	"strings"
)

// ActionTransactionShow is a type for action Transaction#Show
type ActionTransactionShow struct {
	// Pointer to client
	Client *Client
}

func NewActionTransactionShow(client *Client) *ActionTransactionShow {
	return &ActionTransactionShow{
		Client: client,
	}
}

// ActionTransactionShowMetaGlobalInput is a type for action global meta input parameters
type ActionTransactionShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionTransactionShowMetaGlobalInput) SetIncludes(value string) *ActionTransactionShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionTransactionShowMetaGlobalInput) SetNo(value bool) *ActionTransactionShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionTransactionShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionTransactionShowMetaGlobalInput) SelectParameters(params ...string) *ActionTransactionShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionTransactionShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionTransactionShowOutput is a type for action output parameters
type ActionTransactionShowOutput struct {
	CreatedAt        string                            `json:"created_at"`
	DependsOn        *ActionTransactionShowOutput      `json:"depends_on"`
	Done             string                            `json:"done"`
	FinishedAt       string                            `json:"finished_at"`
	Id               int64                             `json:"id"`
	Input            string                            `json:"input"`
	Name             string                            `json:"name"`
	Node             *ActionNodeShowOutput             `json:"node"`
	Output           string                            `json:"output"`
	Priority         int64                             `json:"priority"`
	StartedAt        string                            `json:"started_at"`
	Success          int64                             `json:"success"`
	TransactionChain *ActionTransactionChainShowOutput `json:"transaction_chain"`
	Type             int64                             `json:"type"`
	Urgent           bool                              `json:"urgent"`
	User             *ActionUserShowOutput             `json:"user"`
	Vps              *ActionVpsShowOutput              `json:"vps"`
}

// Type for action response, including envelope
type ActionTransactionShowResponse struct {
	Action *ActionTransactionShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Transaction *ActionTransactionShowOutput `json:"transaction"`
	}

	// Action output without the namespace
	Output *ActionTransactionShowOutput
}

// Prepare the action for invocation
func (action *ActionTransactionShow) Prepare() *ActionTransactionShowInvocation {
	return &ActionTransactionShowInvocation{
		Action: action,
		Path:   "/v6.0/transactions/{transaction_id}",
	}
}

// ActionTransactionShowInvocation is used to configure action for invocation
type ActionTransactionShowInvocation struct {
	// Pointer to the action
	Action *ActionTransactionShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionTransactionShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionTransactionShowInvocation) SetPathParamInt(param string, value int64) *ActionTransactionShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionTransactionShowInvocation) SetPathParamString(param string, value string) *ActionTransactionShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionTransactionShowInvocation) NewMetaInput() *ActionTransactionShowMetaGlobalInput {
	inv.MetaInput = &ActionTransactionShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionTransactionShowInvocation) SetMetaInput(input *ActionTransactionShowMetaGlobalInput) *ActionTransactionShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionTransactionShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionTransactionShowInvocation) Call() (*ActionTransactionShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionTransactionShowInvocation) callAsQuery() (*ActionTransactionShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionTransactionShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Transaction
	}
	return resp, err
}

func (inv *ActionTransactionShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
