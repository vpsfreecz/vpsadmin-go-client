package client

import (
	"strings"
)

// ActionUserGetPaymentInstructions is a type for action User#Get_payment_instructions
type ActionUserGetPaymentInstructions struct {
	// Pointer to client
	Client *Client
}

func NewActionUserGetPaymentInstructions(client *Client) *ActionUserGetPaymentInstructions {
	return &ActionUserGetPaymentInstructions{
		Client: client,
	}
}

// ActionUserGetPaymentInstructionsMetaGlobalInput is a type for action global meta input parameters
type ActionUserGetPaymentInstructionsMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserGetPaymentInstructionsMetaGlobalInput) SetNo(value bool) *ActionUserGetPaymentInstructionsMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserGetPaymentInstructionsMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserGetPaymentInstructionsMetaGlobalInput) SelectParameters(params ...string) *ActionUserGetPaymentInstructionsMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserGetPaymentInstructionsMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserGetPaymentInstructionsOutput is a type for action output parameters
type ActionUserGetPaymentInstructionsOutput struct {
	Instructions string `json:"instructions"`
}

// Type for action response, including envelope
type ActionUserGetPaymentInstructionsResponse struct {
	Action *ActionUserGetPaymentInstructions `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		User *ActionUserGetPaymentInstructionsOutput `json:"user"`
	}

	// Action output without the namespace
	Output *ActionUserGetPaymentInstructionsOutput
}

// Prepare the action for invocation
func (action *ActionUserGetPaymentInstructions) Prepare() *ActionUserGetPaymentInstructionsInvocation {
	return &ActionUserGetPaymentInstructionsInvocation{
		Action: action,
		Path:   "/v6.0/users/{user_id}/get_payment_instructions",
	}
}

// ActionUserGetPaymentInstructionsInvocation is used to configure action for invocation
type ActionUserGetPaymentInstructionsInvocation struct {
	// Pointer to the action
	Action *ActionUserGetPaymentInstructions

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserGetPaymentInstructionsMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserGetPaymentInstructionsInvocation) SetPathParamInt(param string, value int64) *ActionUserGetPaymentInstructionsInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserGetPaymentInstructionsInvocation) SetPathParamString(param string, value string) *ActionUserGetPaymentInstructionsInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserGetPaymentInstructionsInvocation) NewMetaInput() *ActionUserGetPaymentInstructionsMetaGlobalInput {
	inv.MetaInput = &ActionUserGetPaymentInstructionsMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserGetPaymentInstructionsInvocation) SetMetaInput(input *ActionUserGetPaymentInstructionsMetaGlobalInput) *ActionUserGetPaymentInstructionsInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserGetPaymentInstructionsInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserGetPaymentInstructionsInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserGetPaymentInstructionsInvocation) Call() (*ActionUserGetPaymentInstructionsResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserGetPaymentInstructionsInvocation) callAsQuery() (*ActionUserGetPaymentInstructionsResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserGetPaymentInstructionsResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.User
	}
	return resp, err
}

func (inv *ActionUserGetPaymentInstructionsInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
