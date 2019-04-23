package client

import (
	"strings"
)

// ActionNewsLogUpdate is a type for action News_log#Update
type ActionNewsLogUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionNewsLogUpdate(client *Client) *ActionNewsLogUpdate {
	return &ActionNewsLogUpdate{
		Client: client,
	}
}

// ActionNewsLogUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionNewsLogUpdateMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNewsLogUpdateMetaGlobalInput) SetNo(value bool) *ActionNewsLogUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNewsLogUpdateMetaGlobalInput) SetIncludes(value string) *ActionNewsLogUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionNewsLogUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNewsLogUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogUpdateInput is a type for action input parameters
type ActionNewsLogUpdateInput struct {
	Message string `json:"message"`
	PublishedAt string `json:"published_at"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetMessage sets parameter Message to value and selects it for sending
func (in *ActionNewsLogUpdateInput) SetMessage(value string) *ActionNewsLogUpdateInput {
	in.Message = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Message"] = nil
	return in
}
// SetPublishedAt sets parameter PublishedAt to value and selects it for sending
func (in *ActionNewsLogUpdateInput) SetPublishedAt(value string) *ActionNewsLogUpdateInput {
	in.PublishedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["PublishedAt"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogUpdateInput) SelectParameters(params ...string) *ActionNewsLogUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNewsLogUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogUpdateRequest is a type for the entire action request
type ActionNewsLogUpdateRequest struct {
	NewsLog map[string]interface{} `json:"news_log"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionNewsLogUpdateOutput is a type for action output parameters
type ActionNewsLogUpdateOutput struct {
	Id int64 `json:"id"`
	Message string `json:"message"`
	PublishedAt string `json:"published_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionNewsLogUpdateResponse struct {
	Action *ActionNewsLogUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NewsLog *ActionNewsLogUpdateOutput `json:"news_log"`
	}

	// Action output without the namespace
	Output *ActionNewsLogUpdateOutput
}


// Prepare the action for invocation
func (action *ActionNewsLogUpdate) Prepare() *ActionNewsLogUpdateInvocation {
	return &ActionNewsLogUpdateInvocation{
		Action: action,
		Path: "/v5.0/news_logs/:news_log_id",
	}
}

// ActionNewsLogUpdateInvocation is used to configure action for invocation
type ActionNewsLogUpdateInvocation struct {
	// Pointer to the action
	Action *ActionNewsLogUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNewsLogUpdateInput
	// Global meta input parameters
	MetaInput *ActionNewsLogUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNewsLogUpdateInvocation) SetPathParamInt(param string, value int64) *ActionNewsLogUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNewsLogUpdateInvocation) SetPathParamString(param string, value string) *ActionNewsLogUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetInput provides input parameters to send to the API
func (inv *ActionNewsLogUpdateInvocation) SetInput(input *ActionNewsLogUpdateInput) *ActionNewsLogUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNewsLogUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNewsLogUpdateInvocation) SetMetaInput(input *ActionNewsLogUpdateMetaGlobalInput) *ActionNewsLogUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNewsLogUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNewsLogUpdateInvocation) Call() (*ActionNewsLogUpdateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionNewsLogUpdateInvocation) callAsBody() (*ActionNewsLogUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNewsLogUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NewsLog
	}
	return resp, err
}




func (inv *ActionNewsLogUpdateInvocation) makeAllInputParams() *ActionNewsLogUpdateRequest {
	return &ActionNewsLogUpdateRequest{
		NewsLog: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNewsLogUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("Message") {
			ret["message"] = inv.Input.Message
		}
		if inv.IsParameterSelected("PublishedAt") {
			ret["published_at"] = inv.Input.PublishedAt
		}
	}

	return ret
}

func (inv *ActionNewsLogUpdateInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
	}

	return ret
}
