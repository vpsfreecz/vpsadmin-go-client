package client

import (
	"strings"
)

// ActionNewsLogShow is a type for action News_log#Show
type ActionNewsLogShow struct {
	// Pointer to client
	Client *Client
}

func NewActionNewsLogShow(client *Client) *ActionNewsLogShow {
	return &ActionNewsLogShow{
		Client: client,
	}
}

// ActionNewsLogShowMetaGlobalInput is a type for action global meta input parameters
type ActionNewsLogShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNewsLogShowMetaGlobalInput) SetIncludes(value string) *ActionNewsLogShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNewsLogShowMetaGlobalInput) SetNo(value bool) *ActionNewsLogShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNewsLogShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNewsLogShowMetaGlobalInput) SelectParameters(params ...string) *ActionNewsLogShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNewsLogShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNewsLogShowOutput is a type for action output parameters
type ActionNewsLogShowOutput struct {
	CreatedAt   string `json:"created_at"`
	Id          int64  `json:"id"`
	Message     string `json:"message"`
	PublishedAt string `json:"published_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Type for action response, including envelope
type ActionNewsLogShowResponse struct {
	Action *ActionNewsLogShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NewsLog *ActionNewsLogShowOutput `json:"news_log"`
	}

	// Action output without the namespace
	Output *ActionNewsLogShowOutput
}

// Prepare the action for invocation
func (action *ActionNewsLogShow) Prepare() *ActionNewsLogShowInvocation {
	return &ActionNewsLogShowInvocation{
		Action: action,
		Path:   "/v6.0/news_logs/{news_log_id}",
	}
}

// ActionNewsLogShowInvocation is used to configure action for invocation
type ActionNewsLogShowInvocation struct {
	// Pointer to the action
	Action *ActionNewsLogShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionNewsLogShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionNewsLogShowInvocation) SetPathParamInt(param string, value int64) *ActionNewsLogShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionNewsLogShowInvocation) SetPathParamString(param string, value string) *ActionNewsLogShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNewsLogShowInvocation) NewMetaInput() *ActionNewsLogShowMetaGlobalInput {
	inv.MetaInput = &ActionNewsLogShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNewsLogShowInvocation) SetMetaInput(input *ActionNewsLogShowMetaGlobalInput) *ActionNewsLogShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNewsLogShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNewsLogShowInvocation) Call() (*ActionNewsLogShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionNewsLogShowInvocation) callAsQuery() (*ActionNewsLogShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionNewsLogShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NewsLog
	}
	return resp, err
}

func (inv *ActionNewsLogShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
