package client

import (
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
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
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
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNewsLogCreateMetaGlobalInput) SetIncludes(value string) *ActionNewsLogCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
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
	Message string `json:"message"`
	PublishedAt string `json:"published_at"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetMessage sets parameter Message to value and selects it for sending
func (in *ActionNewsLogCreateInput) SetMessage(value string) *ActionNewsLogCreateInput {
	in.Message = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Message"] = nil
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

func (in *ActionNewsLogCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogCreateRequest is a type for the entire action request
type ActionNewsLogCreateRequest struct {
	NewsLog map[string]interface{} `json:"news_log"`
	Meta map[string]interface{} `json:"_meta"`
}

// ActionNewsLogCreateOutput is a type for action output parameters
type ActionNewsLogCreateOutput struct {
	Id int64 `json:"id"`
	Message string `json:"message"`
	PublishedAt string `json:"published_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionNewsLogCreateResponse struct {
	Action *ActionNewsLogCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NewsLog *ActionNewsLogCreateOutput `json:"news_log"`
	}

	// Action output without the namespace
	Output *ActionNewsLogCreateOutput
}


// Prepare the action for invocation
func (action *ActionNewsLogCreate) Prepare() *ActionNewsLogCreateInvocation {
	return &ActionNewsLogCreateInvocation{
		Action: action,
		Path: "/v5.0/news_logs",
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

// Call() invokes the action and returns a response from the API server
func (inv *ActionNewsLogCreateInvocation) Call() (*ActionNewsLogCreateResponse, error) {
	return inv.callAsBody()
}


func (inv *ActionNewsLogCreateInvocation) callAsBody() (*ActionNewsLogCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionNewsLogCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NewsLog
	}
	return resp, err
}




func (inv *ActionNewsLogCreateInvocation) makeAllInputParams() *ActionNewsLogCreateRequest {
	return &ActionNewsLogCreateRequest{
		NewsLog: inv.makeInputParams(),
		Meta: inv.makeMetaInputParams(),
	}
}

func (inv *ActionNewsLogCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionNewsLogCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
