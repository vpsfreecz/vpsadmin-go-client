package client

import (
)

// ActionUserOutageIndex is a type for action User_outage#Index
type ActionUserOutageIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionUserOutageIndex(client *Client) *ActionUserOutageIndex {
	return &ActionUserOutageIndex{
		Client: client,
	}
}

// ActionUserOutageIndexMetaGlobalInput is a type for action global meta input parameters
type ActionUserOutageIndexMetaGlobalInput struct {
	Count bool `json:"count"`
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionUserOutageIndexMetaGlobalInput) SetCount(value bool) *ActionUserOutageIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserOutageIndexMetaGlobalInput) SetIncludes(value string) *ActionUserOutageIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserOutageIndexMetaGlobalInput) SetNo(value bool) *ActionUserOutageIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserOutageIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserOutageIndexMetaGlobalInput) SelectParameters(params ...string) *ActionUserOutageIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserOutageIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserOutageIndexInput is a type for action input parameters
type ActionUserOutageIndexInput struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Outage int64 `json:"outage"`
	User int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionUserOutageIndexInput) SetLimit(value int64) *ActionUserOutageIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionUserOutageIndexInput) SetOffset(value int64) *ActionUserOutageIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetOutage sets parameter Outage to value and selects it for sending
func (in *ActionUserOutageIndexInput) SetOutage(value int64) *ActionUserOutageIndexInput {
	in.Outage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Outage"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionUserOutageIndexInput) SetUser(value int64) *ActionUserOutageIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserOutageIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserOutageIndexInput) SelectParameters(params ...string) *ActionUserOutageIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserOutageIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionUserOutageIndexOutput is a type for action output parameters
type ActionUserOutageIndexOutput struct {
	ExportCount int64 `json:"export_count"`
	Id int64 `json:"id"`
	Outage *ActionOutageShowOutput `json:"outage"`
	User *ActionUserShowOutput `json:"user"`
	VpsCount int64 `json:"vps_count"`
}


// Type for action response, including envelope
type ActionUserOutageIndexResponse struct {
	Action *ActionUserOutageIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		UserOutages []*ActionUserOutageIndexOutput `json:"user_outages"`
	}

	// Action output without the namespace
	Output []*ActionUserOutageIndexOutput
}


// Prepare the action for invocation
func (action *ActionUserOutageIndex) Prepare() *ActionUserOutageIndexInvocation {
	return &ActionUserOutageIndexInvocation{
		Action: action,
		Path: "/v6.0/user_outages",
	}
}

// ActionUserOutageIndexInvocation is used to configure action for invocation
type ActionUserOutageIndexInvocation struct {
	// Pointer to the action
	Action *ActionUserOutageIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserOutageIndexInput
	// Global meta input parameters
	MetaInput *ActionUserOutageIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserOutageIndexInvocation) NewInput() *ActionUserOutageIndexInput {
	inv.Input = &ActionUserOutageIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserOutageIndexInvocation) SetInput(input *ActionUserOutageIndexInput) *ActionUserOutageIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserOutageIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserOutageIndexInvocation) NewMetaInput() *ActionUserOutageIndexMetaGlobalInput {
	inv.MetaInput = &ActionUserOutageIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserOutageIndexInvocation) SetMetaInput(input *ActionUserOutageIndexMetaGlobalInput) *ActionUserOutageIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserOutageIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserOutageIndexInvocation) Call() (*ActionUserOutageIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserOutageIndexInvocation) callAsQuery() (*ActionUserOutageIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserOutageIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.UserOutages
	}
	return resp, err
}



func (inv *ActionUserOutageIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["user_outage[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["user_outage[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Outage") {
			ret["user_outage[outage]"] = convertInt64ToString(inv.Input.Outage)
		}
		if inv.IsParameterSelected("User") {
			ret["user_outage[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionUserOutageIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

