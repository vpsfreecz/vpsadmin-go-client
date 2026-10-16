package client

import (
	"context"
	"net/http"
)

// ActionNewsLogCreate is a type for action News_log#Create
type ActionNewsLogCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionNewsLogCreate(client *Client) *ActionNewsLogCreate {
	return &ActionNewsLogCreate{
		Client: client,
	}
}

// ActionNewsLogCreateMetaGlobalInput is a type for action global meta input parameters
type ActionNewsLogCreateMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNewsLogCreateMetaGlobalInput) SetIncludes(value string) *ActionNewsLogCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNewsLogCreateMetaGlobalInput) SetNo(value bool) *ActionNewsLogCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogCreateMetaGlobalInput) SelectParameters(params ...string) *ActionNewsLogCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNewsLogCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogCreateInput is a type for action input parameters
type ActionNewsLogCreateInput struct {
	CsMessage   string "json:\"cs_message\""
	EnMessage   string "json:\"en_message\""
	PublishedAt string "json:\"published_at\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCsMessage sets parameter CsMessage to value and selects it for sending
func (in *ActionNewsLogCreateInput) SetCsMessage(value string) *ActionNewsLogCreateInput {
	in.CsMessage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CsMessage"] = nil
	return in
}

// SetEnMessage sets parameter EnMessage to value and selects it for sending
func (in *ActionNewsLogCreateInput) SetEnMessage(value string) *ActionNewsLogCreateInput {
	in.EnMessage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EnMessage"] = nil
	return in
}

// SetPublishedAt sets parameter PublishedAt to value and selects it for sending
func (in *ActionNewsLogCreateInput) SetPublishedAt(value string) *ActionNewsLogCreateInput {
	in.PublishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PublishedAt"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogCreateInput) SelectParameters(params ...string) *ActionNewsLogCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNewsLogCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNewsLogCreateInput) UnselectParameters(params ...string) *ActionNewsLogCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNewsLogCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogCreateRequest is a type for the entire action request
type ActionNewsLogCreateRequest struct {
	NewsLog map[string]interface{} "json:\"news_log\""
	Meta    map[string]interface{} "json:\"_meta\""
}

// ActionNewsLogCreateOutput is a type for action output parameters
type ActionNewsLogCreateOutput struct {
	CreatedAt   string "json:\"created_at\""
	CsMessage   string "json:\"cs_message\""
	EnMessage   string "json:\"en_message\""
	Id          int64  "json:\"id\""
	Message     string "json:\"message\""
	PublishedAt string "json:\"published_at\""
	UpdatedAt   string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionNewsLogCreateResponse struct {
	Action *ActionNewsLogCreate "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NewsLog *ActionNewsLogCreateOutput "json:\"news_log\""
	}

	// Action output without the namespace
	Output *ActionNewsLogCreateOutput
}

// Prepare the action for invocation
func (action *ActionNewsLogCreate) Prepare() *ActionNewsLogCreateInvocation {
	return &ActionNewsLogCreateInvocation{
		Action: action,
		Path:   "/v7.0/news_logs",
	}
}

// ActionNewsLogCreateInvocation is used to configure action for invocation
type ActionNewsLogCreateInvocation struct {
	// Pointer to the action
	Action *ActionNewsLogCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNewsLogCreateInput
	// Global meta input parameters
	MetaInput *ActionNewsLogCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNewsLogCreateInvocation) NewInput() *ActionNewsLogCreateInput {
	inv.Input = &ActionNewsLogCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNewsLogCreateInvocation) SetInput(input *ActionNewsLogCreateInput) *ActionNewsLogCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNewsLogCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNewsLogCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNewsLogCreateInvocation) NewMetaInput() *ActionNewsLogCreateMetaGlobalInput {
	inv.MetaInput = &ActionNewsLogCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNewsLogCreateInvocation) SetMetaInput(input *ActionNewsLogCreateMetaGlobalInput) *ActionNewsLogCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNewsLogCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNewsLogCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNewsLogCreateInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.Input != nil {
		if inv.IsParameterSelected("PublishedAt") {
			if !inv.IsParameterNil("PublishedAt") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.PublishedAt)
				if !ok {
					verr.Add("published_at", t(inv.Action.Client, "validation.invalid_datetime", nil))
				} else {
					inv.Input.PublishedAt = normalized
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionNewsLogCreateInvocation) Call() (*ActionNewsLogCreateResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionNewsLogCreateInvocation) CallContext(ctx context.Context) (*ActionNewsLogCreateResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionNewsLogCreateResponse{Action: inv.Action}
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
		if resp.Response.NewsLog == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.NewsLog
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionNewsLogCreateInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionNewsLogCreateInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	input := inv.makeAllInputParams()
	return inv.Action.Client.newBodyRequest(ctx, "POST", inv.Path, input)
}

func (inv *ActionNewsLogCreateInvocation) makeAllInputParams() *ActionNewsLogCreateRequest {
	return &ActionNewsLogCreateRequest{
		NewsLog: inv.makeInputParams(),
		Meta:    inv.makeMetaInputParams(),
	}
}

func (inv *ActionNewsLogCreateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("CsMessage") {
			ret["cs_message"] = inv.Input.CsMessage
		}
		if inv.IsParameterSelected("EnMessage") {
			ret["en_message"] = inv.Input.EnMessage
		}
		if inv.IsParameterSelected("PublishedAt") {
			ret["published_at"] = inv.Input.PublishedAt
		}
	}

	return ret
}

func (inv *ActionNewsLogCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
