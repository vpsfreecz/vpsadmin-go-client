package client

import (
	"strings"
)

// ActionVpsUserDataShow is a type for action Vps_user_data#Show
type ActionVpsUserDataShow struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsUserDataShow(client *Client) *ActionVpsUserDataShow {
	return &ActionVpsUserDataShow{
		Client: client,
	}
}

// ActionVpsUserDataShowMetaGlobalInput is a type for action global meta input parameters
type ActionVpsUserDataShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsUserDataShowMetaGlobalInput) SetIncludes(value string) *ActionVpsUserDataShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsUserDataShowMetaGlobalInput) SetNo(value bool) *ActionVpsUserDataShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUserDataShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataShowMetaGlobalInput) SelectParameters(params ...string) *ActionVpsUserDataShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsUserDataShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataShowOutput is a type for action output parameters
type ActionVpsUserDataShowOutput struct {
	Content   string                `json:"content"`
	CreatedAt string                `json:"created_at"`
	Format    string                `json:"format"`
	Id        int64                 `json:"id"`
	Label     string                `json:"label"`
	UpdatedAt string                `json:"updated_at"`
	User      *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionVpsUserDataShowResponse struct {
	Action *ActionVpsUserDataShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsUserData *ActionVpsUserDataShowOutput `json:"vps_user_data"`
	}

	// Action output without the namespace
	Output *ActionVpsUserDataShowOutput
}

// Prepare the action for invocation
func (action *ActionVpsUserDataShow) Prepare() *ActionVpsUserDataShowInvocation {
	return &ActionVpsUserDataShowInvocation{
		Action: action,
		Path:   "/v7.0/vps_user_data/{vps_user_data_id}",
	}
}

// ActionVpsUserDataShowInvocation is used to configure action for invocation
type ActionVpsUserDataShowInvocation struct {
	// Pointer to the action
	Action *ActionVpsUserDataShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionVpsUserDataShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionVpsUserDataShowInvocation) SetPathParamInt(param string, value int64) *ActionVpsUserDataShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionVpsUserDataShowInvocation) SetPathParamString(param string, value string) *ActionVpsUserDataShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsUserDataShowInvocation) NewMetaInput() *ActionVpsUserDataShowMetaGlobalInput {
	inv.MetaInput = &ActionVpsUserDataShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsUserDataShowInvocation) SetMetaInput(input *ActionVpsUserDataShowMetaGlobalInput) *ActionVpsUserDataShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsUserDataShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsUserDataShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsUserDataShowInvocation) Call() (*ActionVpsUserDataShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsUserDataShowInvocation) callAsQuery() (*ActionVpsUserDataShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsUserDataShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsUserData
	}
	return resp, err
}

func (inv *ActionVpsUserDataShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
