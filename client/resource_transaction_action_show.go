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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionTransactionShowMetaGlobalInput) SetIncludes(value string) *ActionTransactionShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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
	Id int64 `json:"id"`
	TransactionChain *ActionTransactionChainShowOutput `json:"transaction_chain"`
	Node *ActionNodeShowOutput `json:"node"`
	User *ActionUserShowOutput `json:"user"`
	Type int64 `json:"type"`
	Name string `json:"name"`
	Vps *ActionVpsShowOutput `json:"vps"`
	DependsOn *ActionTransactionShowOutput `json:"depends_on"`
	Urgent bool `json:"urgent"`
	Priority int64 `json:"priority"`
	Success int64 `json:"success"`
	Done string `json:"done"`
	Input string `json:"input"`
	Output string `json:"output"`
	CreatedAt string `json:"created_at"`
	StartedAt string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
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
		Path: "/v5.0/transactions/:transaction_id",
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
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
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
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

