package client

import (
	"strings"
)

// ActionTransactionChainShow is a type for action Transaction_chain#Show
type ActionTransactionChainShow struct {
	// Pointer to client
	Client *Client
}

func NewActionTransactionChainShow(client *Client) *ActionTransactionChainShow {
	return &ActionTransactionChainShow{
		Client: client,
	}
}

// ActionTransactionChainShowMetaGlobalInput is a type for action global meta input parameters
type ActionTransactionChainShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionTransactionChainShowMetaGlobalInput) SetNo(value bool) *ActionTransactionChainShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionTransactionChainShowMetaGlobalInput) SetIncludes(value string) *ActionTransactionChainShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionTransactionChainShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionTransactionChainShowMetaGlobalInput) SelectParameters(params ...string) *ActionTransactionChainShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionTransactionChainShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionTransactionChainShowOutput is a type for action output parameters
type ActionTransactionChainShowOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	State string `json:"state"`
	Size int64 `json:"size"`
	Progress int64 `json:"progress"`
	User *ActionUserShowOutput `json:"user"`
	CreatedAt string `json:"created_at"`
}


// Type for action response, including envelope
type ActionTransactionChainShowResponse struct {
	Action *ActionTransactionChainShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TransactionChain *ActionTransactionChainShowOutput `json:"transaction_chain"`
	}

	// Action output without the namespace
	Output *ActionTransactionChainShowOutput
}


// Prepare the action for invocation
func (action *ActionTransactionChainShow) Prepare() *ActionTransactionChainShowInvocation {
	return &ActionTransactionChainShowInvocation{
		Action: action,
		Path: "/v5.0/transaction_chains/{transaction_chain_id}",
	}
}

// ActionTransactionChainShowInvocation is used to configure action for invocation
type ActionTransactionChainShowInvocation struct {
	// Pointer to the action
	Action *ActionTransactionChainShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionTransactionChainShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionTransactionChainShowInvocation) SetPathParamInt(param string, value int64) *ActionTransactionChainShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionTransactionChainShowInvocation) SetPathParamString(param string, value string) *ActionTransactionChainShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionTransactionChainShowInvocation) NewMetaInput() *ActionTransactionChainShowMetaGlobalInput {
	inv.MetaInput = &ActionTransactionChainShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionTransactionChainShowInvocation) SetMetaInput(input *ActionTransactionChainShowMetaGlobalInput) *ActionTransactionChainShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionTransactionChainShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionTransactionChainShowInvocation) Call() (*ActionTransactionChainShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionTransactionChainShowInvocation) callAsQuery() (*ActionTransactionChainShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionTransactionChainShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TransactionChain
	}
	return resp, err
}




func (inv *ActionTransactionChainShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

