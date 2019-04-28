package client

import (
	"strings"
)

// ActionActionStateShow is a type for action Action_state#Show
type ActionActionStateShow struct {
	// Pointer to client
	Client *Client
}

func NewActionActionStateShow(client *Client) *ActionActionStateShow {
	return &ActionActionStateShow{
		Client: client,
	}
}

// ActionActionStateShowMetaGlobalInput is a type for action global meta input parameters
type ActionActionStateShowMetaGlobalInput struct {
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionActionStateShowMetaGlobalInput) SetNo(value bool) *ActionActionStateShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionActionStateShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionActionStateShowMetaGlobalInput) SelectParameters(params ...string) *ActionActionStateShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionActionStateShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionActionStateShowOutput is a type for action output parameters
type ActionActionStateShowOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Finished bool `json:"finished"`
	Status bool `json:"status"`
	Current int64 `json:"current"`
	Total int64 `json:"total"`
	Unit string `json:"unit"`
	CanCancel bool `json:"can_cancel"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionActionStateShowResponse struct {
	Action *ActionActionStateShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ActionState *ActionActionStateShowOutput `json:"action_state"`
	}

	// Action output without the namespace
	Output *ActionActionStateShowOutput
}


// Prepare the action for invocation
func (action *ActionActionStateShow) Prepare() *ActionActionStateShowInvocation {
	return &ActionActionStateShowInvocation{
		Action: action,
		Path: "/v5.0/action_states/:action_state_id",
	}
}

// ActionActionStateShowInvocation is used to configure action for invocation
type ActionActionStateShowInvocation struct {
	// Pointer to the action
	Action *ActionActionStateShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionActionStateShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionActionStateShowInvocation) SetPathParamInt(param string, value int64) *ActionActionStateShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionActionStateShowInvocation) SetPathParamString(param string, value string) *ActionActionStateShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionActionStateShowInvocation) NewMetaInput() *ActionActionStateShowMetaGlobalInput {
	inv.MetaInput = &ActionActionStateShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionActionStateShowInvocation) SetMetaInput(input *ActionActionStateShowMetaGlobalInput) *ActionActionStateShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionActionStateShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionActionStateShowInvocation) Call() (*ActionActionStateShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionActionStateShowInvocation) callAsQuery() (*ActionActionStateShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionActionStateShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ActionState
	}
	return resp, err
}




func (inv *ActionActionStateShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

