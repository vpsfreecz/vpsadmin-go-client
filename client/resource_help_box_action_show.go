package client

import (
	"strings"
)

// ActionHelpBoxShow is a type for action Help_box#Show
type ActionHelpBoxShow struct {
	// Pointer to client
	Client *Client
}

func NewActionHelpBoxShow(client *Client) *ActionHelpBoxShow {
	return &ActionHelpBoxShow{
		Client: client,
	}
}

// ActionHelpBoxShowMetaGlobalInput is a type for action global meta input parameters
type ActionHelpBoxShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionHelpBoxShowMetaGlobalInput) SetNo(value bool) *ActionHelpBoxShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionHelpBoxShowMetaGlobalInput) SetIncludes(value string) *ActionHelpBoxShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionHelpBoxShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionHelpBoxShowMetaGlobalInput) SelectParameters(params ...string) *ActionHelpBoxShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionHelpBoxShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionHelpBoxShowOutput is a type for action output parameters
type ActionHelpBoxShowOutput struct {
	Id int64 `json:"id"`
	Page string `json:"page"`
	Action string `json:"action"`
	Language *ActionLanguageShowOutput `json:"language"`
	Content string `json:"content"`
	Order int64 `json:"order"`
}


// Type for action response, including envelope
type ActionHelpBoxShowResponse struct {
	Action *ActionHelpBoxShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		HelpBox *ActionHelpBoxShowOutput `json:"help_box"`
	}

	// Action output without the namespace
	Output *ActionHelpBoxShowOutput
}


// Prepare the action for invocation
func (action *ActionHelpBoxShow) Prepare() *ActionHelpBoxShowInvocation {
	return &ActionHelpBoxShowInvocation{
		Action: action,
		Path: "/v5.0/help_boxes/{help_box_id}",
	}
}

// ActionHelpBoxShowInvocation is used to configure action for invocation
type ActionHelpBoxShowInvocation struct {
	// Pointer to the action
	Action *ActionHelpBoxShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionHelpBoxShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionHelpBoxShowInvocation) SetPathParamInt(param string, value int64) *ActionHelpBoxShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionHelpBoxShowInvocation) SetPathParamString(param string, value string) *ActionHelpBoxShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionHelpBoxShowInvocation) NewMetaInput() *ActionHelpBoxShowMetaGlobalInput {
	inv.MetaInput = &ActionHelpBoxShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionHelpBoxShowInvocation) SetMetaInput(input *ActionHelpBoxShowMetaGlobalInput) *ActionHelpBoxShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionHelpBoxShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionHelpBoxShowInvocation) Call() (*ActionHelpBoxShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionHelpBoxShowInvocation) callAsQuery() (*ActionHelpBoxShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionHelpBoxShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.HelpBox
	}
	return resp, err
}




func (inv *ActionHelpBoxShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

