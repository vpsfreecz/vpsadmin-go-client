package client

import (
	"strings"
)

// ActionUserPublicKeyIndex is a type for action User.Public_key#Index
type ActionUserPublicKeyIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserPublicKeyIndex(client *Client) *ActionUserPublicKeyIndex {
	return &ActionUserPublicKeyIndex{
		Client: client,
	}
}

// ActionUserPublicKeyIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserPublicKeyIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserPublicKeyIndexMetaGlobalInput) SetNo(value bool) *ActionUserPublicKeyIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserPublicKeyIndexMetaGlobalInput) SetCount(value bool) *ActionUserPublicKeyIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserPublicKeyIndexMetaGlobalInput) SetIncludes(value string) *ActionUserPublicKeyIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPublicKeyIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPublicKeyIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserPublicKeyIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPublicKeyIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserPublicKeyIndexInput is a type for action input parameters
type ActionUserPublicKeyIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserPublicKeyIndexInput) SetOffset(value int64) *ActionUserPublicKeyIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserPublicKeyIndexInput) SetLimit(value int64) *ActionUserPublicKeyIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserPublicKeyIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserPublicKeyIndexInput) SelectParameters(params ...string) *ActionUserPublicKeyIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserPublicKeyIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserPublicKeyIndexOutput is a type for action output parameters
type ActionUserPublicKeyIndexOutput struct {
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
type ActionUserPublicKeyIndexResponse struct {
	Action *ActionUserPublicKeyIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		PublicKeys []*ActionUserPublicKeyIndexOutput `json:"public_keys"`
	}

	// Action output without the namespace
	Output []*ActionUserPublicKeyIndexOutput
}


// Prepare the action for invocation
func (action *ActionUserPublicKeyIndex) Prepare() *ActionUserPublicKeyIndexInvocation {
	return &ActionUserPublicKeyIndexInvocation{
		Action: action,
		Path: "/v5.0/users/:user_id/public_keys",
	}
}

// ActionUserPublicKeyIndexInvocation is used to configure action for invocation
type ActionUserPublicKeyIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserPublicKeyIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserPublicKeyIndexInput
	// Global meta input parameters
	MetaInput *ActionUserPublicKeyIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserPublicKeyIndexInvocation) SetPathParamInt(param string, value int64) *ActionUserPublicKeyIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserPublicKeyIndexInvocation) SetPathParamString(param string, value string) *ActionUserPublicKeyIndexInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserPublicKeyIndexInvocation) NewInput() *ActionUserPublicKeyIndexInput {
	inv.Input = &ActionUserPublicKeyIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserPublicKeyIndexInvocation) SetInput(input *ActionUserPublicKeyIndexInput) *ActionUserPublicKeyIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserPublicKeyIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserPublicKeyIndexInvocation) NewMetaInput() *ActionUserPublicKeyIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserPublicKeyIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserPublicKeyIndexInvocation) SetMetaInput(input *ActionUserPublicKeyIndexMetaGlobalInput) *ActionUserPublicKeyIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserPublicKeyIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserPublicKeyIndexInvocation) Call() (*ActionUserPublicKeyIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserPublicKeyIndexInvocation) callAsQuery() (*ActionUserPublicKeyIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserPublicKeyIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.PublicKeys
	}
	return resp, err
}



func (inv *ActionUserPublicKeyIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["public_key[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["public_key[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionUserPublicKeyIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

