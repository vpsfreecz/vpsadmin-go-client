package client

import (
	"strings"
)

// ActionVpsOutageShow is a type for action Vps_outage#Show
type ActionVpsOutageShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsOutageShow(client *Client) *ActionVpsOutageShow {
	return &ActionVpsOutageShow{
		Client: client,
	}
}

// ActionVpsOutageShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsOutageShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsOutageShowMetaGlobalInput) SetNo(value bool) *ActionVpsOutageShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsOutageShowMetaGlobalInput) SetIncludes(value string) *ActionVpsOutageShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsOutageShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsOutageShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsOutageShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsOutageShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionVpsOutageShowOutput is a type for action output parameters
type ActionVpsOutageShowOutput struct {
	Id int64 `json:"id"`
	Outage *ActionOutageShowOutput `json:"outage"`
	Vps *ActionVpsShowOutput `json:"vps"`
	User *ActionUserShowOutput `json:"user"`
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	Location *ActionLocationShowOutput `json:"location"`
	Node *ActionNodeShowOutput `json:"node"`
	Direct bool `json:"direct"`
}


// Type for action response, including envelope
type ActionVpsOutageShowResponse struct {
	Action *ActionVpsOutageShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsOutage *ActionVpsOutageShowOutput `json:"vps_outage"`
	}

	// Action output without the namespace
	Output *ActionVpsOutageShowOutput
}


// Prepare the action for invocation
func (action *ActionVpsOutageShow) Prepare() *ActionVpsOutageShowInvocation {
	return &ActionVpsOutageShowInvocation{
		Action: action,
		Path: "/v5.0/vps_outages/:vps_outage_id",
	}
}

// ActionVpsOutageShowInvocation is used to configure action for invocation
type ActionVpsOutageShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsOutageShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsOutageShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsOutageShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsOutageShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsOutageShowInvocation) SetPathParamString(param string, value string) *ActionVpsOutageShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsOutageShowInvocation) SetMetaInput(input *ActionVpsOutageShowMetaGlobalInput) *ActionVpsOutageShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsOutageShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsOutageShowInvocation) Call() (*ActionVpsOutageShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsOutageShowInvocation) callAsQuery() (*ActionVpsOutageShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsOutageShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsOutage
	}
	return resp, err
}




func (inv *ActionVpsOutageShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

