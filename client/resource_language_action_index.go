package client

import (
)

// ActionLanguageIndex is a type for action Language#Index
type ActionLanguageIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionLanguageIndex(client *Client) *ActionLanguageIndex {
	return &ActionLanguageIndex{
		Client: client,
	}
}

// ActionLanguageIndexMetaGlobalInput is a type for action global meta input parameters
type ActionLanguageIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionLanguageIndexMetaGlobalInput) SetNo(value bool) *ActionLanguageIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionLanguageIndexMetaGlobalInput) SetCount(value bool) *ActionLanguageIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionLanguageIndexMetaGlobalInput) SetIncludes(value string) *ActionLanguageIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionLanguageIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLanguageIndexMetaGlobalInput) SelectParameters(params ...string) *ActionLanguageIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLanguageIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionLanguageIndexInput is a type for action input parameters
type ActionLanguageIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionLanguageIndexInput) SetOffset(value int64) *ActionLanguageIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionLanguageIndexInput) SetLimit(value int64) *ActionLanguageIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionLanguageIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionLanguageIndexInput) SelectParameters(params ...string) *ActionLanguageIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionLanguageIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionLanguageIndexOutput is a type for action output parameters
type ActionLanguageIndexOutput struct {
	Id int64 `json:"id"`
	Code string `json:"code"`
	Label string `json:"label"`
}


// Type for action response, including envelope
type ActionLanguageIndexResponse struct {
	Action *ActionLanguageIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Languages []*ActionLanguageIndexOutput `json:"languages"`
	}

	// Action output without the namespace
	Output []*ActionLanguageIndexOutput
}


// Prepare the action for invocation
func (action *ActionLanguageIndex) Prepare() *ActionLanguageIndexInvocation {
	return &ActionLanguageIndexInvocation{
		Action: action,
		Path: "/v6.0/languages",
	}
}

// ActionLanguageIndexInvocation is used to configure action for invocation
type ActionLanguageIndexInvocation struct {
	// Pointer to the action
	Action *ActionLanguageIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionLanguageIndexInput
	// Global meta input parameters
	MetaInput *ActionLanguageIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionLanguageIndexInvocation) NewInput() *ActionLanguageIndexInput {
	inv.Input = &ActionLanguageIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionLanguageIndexInvocation) SetInput(input *ActionLanguageIndexInput) *ActionLanguageIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionLanguageIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionLanguageIndexInvocation) NewMetaInput() *ActionLanguageIndexMetaGlobalInput {
	inv.MetaInput = &ActionLanguageIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionLanguageIndexInvocation) SetMetaInput(input *ActionLanguageIndexMetaGlobalInput) *ActionLanguageIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionLanguageIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionLanguageIndexInvocation) Call() (*ActionLanguageIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionLanguageIndexInvocation) callAsQuery() (*ActionLanguageIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionLanguageIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Languages
	}
	return resp, err
}



func (inv *ActionLanguageIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["language[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["language[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionLanguageIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

