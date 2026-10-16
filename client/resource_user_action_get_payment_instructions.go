package client

import (
	"context"
	"net/http"
	"net/url"
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
	No bool "json:\"no\""
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
	Instructions string "json:\"instructions\""
}

// Type for action response, including envelope
type ActionUserGetPaymentInstructionsResponse struct {
	Action *ActionUserGetPaymentInstructions "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		User *ActionUserGetPaymentInstructionsOutput "json:\"user\""
	}

	// Action output without the namespace
	Output *ActionUserGetPaymentInstructionsOutput
}

// Prepare the action for invocation
func (action *ActionUserGetPaymentInstructions) Prepare() *ActionUserGetPaymentInstructionsInvocation {
	return &ActionUserGetPaymentInstructionsInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/get_payment_instructions",
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
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
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

func (inv *ActionUserGetPaymentInstructionsInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionUserGetPaymentInstructionsInvocation) Call() (*ActionUserGetPaymentInstructionsResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionUserGetPaymentInstructionsInvocation) CallContext(ctx context.Context) (*ActionUserGetPaymentInstructionsResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionUserGetPaymentInstructionsResponse{Action: inv.Action}
	raw, err := inv.Action.Client.doRaw(req)
	if err != nil {
		return resp, err
	}
	if err := decodeTypedResponse(raw.Body, resp); err != nil {
		return resp, err
	}
	if resp.Status {
		if resp.Response == nil {
			return resp, errMalformedResponse
		}
		if resp.Response.User == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.User
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionUserGetPaymentInstructionsInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionUserGetPaymentInstructionsInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionUserGetPaymentInstructionsInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
