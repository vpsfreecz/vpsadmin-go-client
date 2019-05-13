package client

import (
	"strings"
)

// ActionUserPublicKeyShow is a type for action User.Public_key#Show
type ActionUserPublicKeyShow struct {
	// Pointer to client
	Client *Client
}

func NewActionUserPublicKeyShow(client *Client) *ActionUserPublicKeyShow {
	return &ActionUserPublicKeyShow{
		Client: client,
	}
}

// ActionUserPublicKeyShowMetaGlobalInput is a type for action global meta input parameters
type ActionUserPublicKeyShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserPublicKeyShowMetaGlobalInput) SetNo(value bool) *ActionUserPublicKeyShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserPublicKeyShowMetaGlobalInput) SetIncludes(value string) *ActionUserPublicKeyShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPublicKeyShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPublicKeyShowMetaGlobalInput) SelectParameters(params ...string) *ActionUserPublicKeyShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPublicKeyShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserPublicKeyShowOutput is a type for action output parameters
type ActionUserPublicKeyShowOutput struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	Key string `json:"key"`
	AutoAdd bool `json:"auto_add"`
	Fingerprint string `json:"fingerprint"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionUserPublicKeyShowResponse struct {
	Action *ActionUserPublicKeyShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PublicKey *ActionUserPublicKeyShowOutput `json:"public_key"`
	}

	// Action output without the namespace
	Output *ActionUserPublicKeyShowOutput
}


// Prepare the action for invocation
func (action *ActionUserPublicKeyShow) Prepare() *ActionUserPublicKeyShowInvocation {
	return &ActionUserPublicKeyShowInvocation{
		Action: action,
		Path: "/v5.0/users/{user_id}/public_keys/{public_key_id}",
	}
}

// ActionUserPublicKeyShowInvocation is used to configure action for invocation
type ActionUserPublicKeyShowInvocation struct {
	// Pointer to the action
	Action *ActionUserPublicKeyShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserPublicKeyShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserPublicKeyShowInvocation) SetPathParamInt(param string, value int64) *ActionUserPublicKeyShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserPublicKeyShowInvocation) SetPathParamString(param string, value string) *ActionUserPublicKeyShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserPublicKeyShowInvocation) NewMetaInput() *ActionUserPublicKeyShowMetaGlobalInput {
	inv.MetaInput = &ActionUserPublicKeyShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserPublicKeyShowInvocation) SetMetaInput(input *ActionUserPublicKeyShowMetaGlobalInput) *ActionUserPublicKeyShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserPublicKeyShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserPublicKeyShowInvocation) Call() (*ActionUserPublicKeyShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserPublicKeyShowInvocation) callAsQuery() (*ActionUserPublicKeyShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserPublicKeyShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.PublicKey
	}
	return resp, err
}




func (inv *ActionUserPublicKeyShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

