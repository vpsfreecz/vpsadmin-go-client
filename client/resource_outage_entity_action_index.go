package client

import (
	"strings"
)

// ActionOutageEntityIndex is a type for action Outage.Entity#Index
type ActionOutageEntityIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageEntityIndex(client *Client) *ActionOutageEntityIndex {
	return &ActionOutageEntityIndex{
		Client: client,
	}
}

// ActionOutageEntityIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOutageEntityIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageEntityIndexMetaGlobalInput) SetNo(value bool) *ActionOutageEntityIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOutageEntityIndexMetaGlobalInput) SetCount(value bool) *ActionOutageEntityIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageEntityIndexMetaGlobalInput) SetIncludes(value string) *ActionOutageEntityIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageEntityIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageEntityIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOutageEntityIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageEntityIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageEntityIndexInput is a type for action input parameters
type ActionOutageEntityIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionOutageEntityIndexInput) SetOffset(value int64) *ActionOutageEntityIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOutageEntityIndexInput) SetLimit(value int64) *ActionOutageEntityIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageEntityIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageEntityIndexInput) SelectParameters(params ...string) *ActionOutageEntityIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageEntityIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionOutageEntityIndexOutput is a type for action output parameters
type ActionOutageEntityIndexOutput struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	EntityId int64 `json:"entity_id"`
	Label string `json:"label"`
}


// Type for action response, including envelope
type ActionOutageEntityIndexResponse struct {
	Action *ActionOutageEntityIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Entities []*ActionOutageEntityIndexOutput `json:"entities"`
	}

	// Action output without the namespace
	Output []*ActionOutageEntityIndexOutput
}


// Prepare the action for invocation
func (action *ActionOutageEntityIndex) Prepare() *ActionOutageEntityIndexInvocation {
	return &ActionOutageEntityIndexInvocation{
		Action: action,
		Path: "/v6.0/outages/{outage_id}/entities",
	}
}

// ActionOutageEntityIndexInvocation is used to configure action for invocation
type ActionOutageEntityIndexInvocation struct {
	// Pointer to the action
	Action *ActionOutageEntityIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageEntityIndexInput
	// Global meta input parameters
	MetaInput *ActionOutageEntityIndexMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOutageEntityIndexInvocation) SetPathParamInt(param string, value int64) *ActionOutageEntityIndexInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOutageEntityIndexInvocation) SetPathParamString(param string, value string) *ActionOutageEntityIndexInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageEntityIndexInvocation) NewInput() *ActionOutageEntityIndexInput {
	inv.Input = &ActionOutageEntityIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageEntityIndexInvocation) SetInput(input *ActionOutageEntityIndexInput) *ActionOutageEntityIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageEntityIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageEntityIndexInvocation) NewMetaInput() *ActionOutageEntityIndexMetaGlobalInput {
	inv.MetaInput = &ActionOutageEntityIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageEntityIndexInvocation) SetMetaInput(input *ActionOutageEntityIndexMetaGlobalInput) *ActionOutageEntityIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageEntityIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageEntityIndexInvocation) Call() (*ActionOutageEntityIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOutageEntityIndexInvocation) callAsQuery() (*ActionOutageEntityIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOutageEntityIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Entities
	}
	return resp, err
}



func (inv *ActionOutageEntityIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["entity[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["entity[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
	}
}

func (inv *ActionOutageEntityIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

