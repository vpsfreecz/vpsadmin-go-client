package client

import (
	"strings"
)

// ActionUserRequestChangeShow is a type for action User_request.Change#Show
type ActionUserRequestChangeShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestChangeShow(client *Client) *ActionUserRequestChangeShow {
	return &ActionUserRequestChangeShow{
		Client: client,
	}
}

// ActionUserRequestChangeShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestChangeShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestChangeShowMetaGlobalInput) SetNo(value bool) *ActionUserRequestChangeShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestChangeShowMetaGlobalInput) SetIncludes(value string) *ActionUserRequestChangeShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestChangeShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestChangeShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestChangeShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestChangeShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserRequestChangeShowOutput is a type for action output parameters
type ActionUserRequestChangeShowOutput struct {
	Id int64 `json:"id"`
	User *ActionUserShowOutput `json:"user"`
	State string `json:"state"`
	ApiIpAddr string `json:"api_ip_addr"`
	ApiIpPtr string `json:"api_ip_ptr"`
	ClientIpAddr string `json:"client_ip_addr"`
	ClientIpPtr string `json:"client_ip_ptr"`
	Admin *ActionUserShowOutput `json:"admin"`
	AdminResponse string `json:"admin_response"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Label string `json:"label"`
	ChangeReason string `json:"change_reason"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Address string `json:"address"`
}


// Type for action response, including envelope
type ActionUserRequestChangeShowResponse struct {
	Action *ActionUserRequestChangeShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Change *ActionUserRequestChangeShowOutput `json:"change"`
	}

	// Action output without the namespace
	Output *ActionUserRequestChangeShowOutput
}


// Prepare the action for invocation
func (action *ActionUserRequestChangeShow) Prepare() *ActionUserRequestChangeShowInvocation {
	return &ActionUserRequestChangeShowInvocation{
		Action: action,
		Path: "/v5.0/user_request/changes/{change_id}",
	}
}

// ActionUserRequestChangeShowInvocation is used to configure action for invocation
type ActionUserRequestChangeShowInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestChangeShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserRequestChangeShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserRequestChangeShowInvocation) SetPathParamInt(param string, value int64) *ActionUserRequestChangeShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserRequestChangeShowInvocation) SetPathParamString(param string, value string) *ActionUserRequestChangeShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestChangeShowInvocation) NewMetaInput() *ActionUserRequestChangeShowMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestChangeShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestChangeShowInvocation) SetMetaInput(input *ActionUserRequestChangeShowMetaGlobalInput) *ActionUserRequestChangeShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestChangeShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestChangeShowInvocation) Call() (*ActionUserRequestChangeShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserRequestChangeShowInvocation) callAsQuery() (*ActionUserRequestChangeShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserRequestChangeShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Change
	}
	return resp, err
}




func (inv *ActionUserRequestChangeShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

