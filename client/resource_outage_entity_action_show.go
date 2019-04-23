package client

import (
	"strings"
)

// ActionOutageEntityShow is a type for action Outage.Entity#Show
type ActionOutageEntityShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageEntityShow(client *Client) *ActionOutageEntityShow {
	return &ActionOutageEntityShow{
		Client: client,
	}
}

// ActionOutageEntityShowMetaGlobalInput is a type for action global meta input parameters
type ActionOutageEntityShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageEntityShowMetaGlobalInput) SetNo(value bool) *ActionOutageEntityShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageEntityShowMetaGlobalInput) SetIncludes(value string) *ActionOutageEntityShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageEntityShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageEntityShowMetaGlobalInput) SelectParameters(params ...string) *ActionOutageEntityShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageEntityShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionOutageEntityShowOutput is a type for action output parameters
type ActionOutageEntityShowOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	EntityId int64 `json:"entity_id"`
	Label string `json:"label"`
}


// Type for action response, including envelope
type ActionOutageEntityShowResponse struct {
	Action *ActionOutageEntityShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Entity *ActionOutageEntityShowOutput `json:"entity"`
	}

	// Action output without the namespace
	Output *ActionOutageEntityShowOutput
}


// Prepare the action for invocation
func (action *ActionOutageEntityShow) Prepare() *ActionOutageEntityShowInvocation {
	return &ActionOutageEntityShowInvocation{
		Action: action,
		Path: "/v5.0/outages/:outage_id/entities/:entity_id",
	}
}

// ActionOutageEntityShowInvocation is used to configure action for invocation
type ActionOutageEntityShowInvocation struct {
	// Pointer to the action
	Action *ActionOutageEntityShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOutageEntityShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageEntityShowInvocation) SetPathParamInt(param string, value int64) *ActionOutageEntityShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageEntityShowInvocation) SetPathParamString(param string, value string) *ActionOutageEntityShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageEntityShowInvocation) SetMetaInput(input *ActionOutageEntityShowMetaGlobalInput) *ActionOutageEntityShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageEntityShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageEntityShowInvocation) Call() (*ActionOutageEntityShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOutageEntityShowInvocation) callAsQuery() (*ActionOutageEntityShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOutageEntityShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Entity
	}
	return resp, err
}




func (inv *ActionOutageEntityShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

