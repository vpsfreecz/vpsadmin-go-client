package client

import (
	"strings"
)

// ActionVpsFeatureShow is a type for action Vps.Feature#Show
type ActionVpsFeatureShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsFeatureShow(client *Client) *ActionVpsFeatureShow {
	return &ActionVpsFeatureShow{
		Client: client,
	}
}

// ActionVpsFeatureShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsFeatureShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsFeatureShowMetaGlobalInput) SetNo(value bool) *ActionVpsFeatureShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsFeatureShowMetaGlobalInput) SetIncludes(value string) *ActionVpsFeatureShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsFeatureShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsFeatureShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsFeatureShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsFeatureShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionVpsFeatureShowOutput is a type for action output parameters
type ActionVpsFeatureShowOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	Enabled bool `json:"enabled"`
}


// Type for action response, including envelope
type ActionVpsFeatureShowResponse struct {
	Action *ActionVpsFeatureShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Feature *ActionVpsFeatureShowOutput `json:"feature"`
	}

	// Action output without the namespace
	Output *ActionVpsFeatureShowOutput
}


// Prepare the action for invocation
func (action *ActionVpsFeatureShow) Prepare() *ActionVpsFeatureShowInvocation {
	return &ActionVpsFeatureShowInvocation{
		Action: action,
		Path: "/v6.0/vpses/{vps_id}/features/{feature_id}",
	}
}

// ActionVpsFeatureShowInvocation is used to configure action for invocation
type ActionVpsFeatureShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsFeatureShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsFeatureShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsFeatureShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsFeatureShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsFeatureShowInvocation) SetPathParamString(param string, value string) *ActionVpsFeatureShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsFeatureShowInvocation) NewMetaInput() *ActionVpsFeatureShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsFeatureShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsFeatureShowInvocation) SetMetaInput(input *ActionVpsFeatureShowMetaGlobalInput) *ActionVpsFeatureShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsFeatureShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsFeatureShowInvocation) Call() (*ActionVpsFeatureShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsFeatureShowInvocation) callAsQuery() (*ActionVpsFeatureShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsFeatureShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Feature
	}
	return resp, err
}




func (inv *ActionVpsFeatureShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

