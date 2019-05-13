package client

import (
	"strings"
)

// ActionObjectHistoryShow is a type for action Object_history#Show
type ActionObjectHistoryShow struct {
	// Pointer to client
	Client *Client
}

func NewActionObjectHistoryShow(client *Client) *ActionObjectHistoryShow {
	return &ActionObjectHistoryShow{
		Client: client,
	}
}

// ActionObjectHistoryShowMetaGlobalInput is a type for action global meta input parameters
type ActionObjectHistoryShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionObjectHistoryShowMetaGlobalInput) SetNo(value bool) *ActionObjectHistoryShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionObjectHistoryShowMetaGlobalInput) SetIncludes(value string) *ActionObjectHistoryShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionObjectHistoryShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionObjectHistoryShowMetaGlobalInput) SelectParameters(params ...string) *ActionObjectHistoryShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionObjectHistoryShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionObjectHistoryShowOutput is a type for action output parameters
type ActionObjectHistoryShowOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	UserSession *ActionUserSessionShowOutput `json:"user_session"`
	Object string `json:"object"`
	ObjectId int64 `json:"object_id"`
	EventType string `json:"event_type"`
	CreatedAt string `json:"created_at"`
}


// Type for action response, including envelope
type ActionObjectHistoryShowResponse struct {
	Action *ActionObjectHistoryShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ObjectHistory *ActionObjectHistoryShowOutput `json:"object_history"`
	}

	// Action output without the namespace
	Output *ActionObjectHistoryShowOutput
}


// Prepare the action for invocation
func (action *ActionObjectHistoryShow) Prepare() *ActionObjectHistoryShowInvocation {
	return &ActionObjectHistoryShowInvocation{
		Action: action,
		Path: "/v5.0/object_histories/{object_history_id}",
	}
}

// ActionObjectHistoryShowInvocation is used to configure action for invocation
type ActionObjectHistoryShowInvocation struct {
	// Pointer to the action
	Action *ActionObjectHistoryShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionObjectHistoryShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionObjectHistoryShowInvocation) SetPathParamInt(param string, value int64) *ActionObjectHistoryShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionObjectHistoryShowInvocation) SetPathParamString(param string, value string) *ActionObjectHistoryShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionObjectHistoryShowInvocation) NewMetaInput() *ActionObjectHistoryShowMetaGlobalInput {
	inv.MetaInput = &ActionObjectHistoryShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionObjectHistoryShowInvocation) SetMetaInput(input *ActionObjectHistoryShowMetaGlobalInput) *ActionObjectHistoryShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionObjectHistoryShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionObjectHistoryShowInvocation) Call() (*ActionObjectHistoryShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionObjectHistoryShowInvocation) callAsQuery() (*ActionObjectHistoryShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionObjectHistoryShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ObjectHistory
	}
	return resp, err
}




func (inv *ActionObjectHistoryShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

