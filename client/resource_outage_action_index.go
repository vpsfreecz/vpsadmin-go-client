package client

import (
)

// ActionOutageIndex is a type for action Outage#Index
type ActionOutageIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionOutageIndex(client *Client) *ActionOutageIndex {
	return &ActionOutageIndex{
		Client: client,
	}
}

// ActionOutageIndexMetaGlobalInput is a type for action global meta input parameters
type ActionOutageIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOutageIndexMetaGlobalInput) SetNo(value bool) *ActionOutageIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionOutageIndexMetaGlobalInput) SetCount(value bool) *ActionOutageIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOutageIndexMetaGlobalInput) SetIncludes(value string) *ActionOutageIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageIndexMetaGlobalInput) SelectParameters(params ...string) *ActionOutageIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOutageIndexInput is a type for action input parameters
type ActionOutageIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Planned bool `json:"planned"`
	State string `json:"state"`
	Type string `json:"type"`
	Affected bool `json:"affected"`
	User int64 `json:"user"`
	Vps int64 `json:"vps"`
	HandledBy int64 `json:"handled_by"`
	Environment int64 `json:"environment"`
	Location int64 `json:"location"`
	Node int64 `json:"node"`
	EntityName string `json:"entity_name"`
	EntityId int64 `json:"entity_id"`
	Order string `json:"order"`
	Since string `json:"since"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionOutageIndexInput) SetOffset(value int64) *ActionOutageIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionOutageIndexInput) SetLimit(value int64) *ActionOutageIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetPlanned sets parameter Planned to value and selects it for sending
func (in *ActionOutageIndexInput) SetPlanned(value bool) *ActionOutageIndexInput {
	in.Planned = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Planned"] = nil
	return in
}
// SetState sets parameter State to value and selects it for sending
func (in *ActionOutageIndexInput) SetState(value string) *ActionOutageIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}
// SetType sets parameter Type to value and selects it for sending
func (in *ActionOutageIndexInput) SetType(value string) *ActionOutageIndexInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}
// SetAffected sets parameter Affected to value and selects it for sending
func (in *ActionOutageIndexInput) SetAffected(value bool) *ActionOutageIndexInput {
	in.Affected = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Affected"] = nil
	return in
}
// SetUser sets parameter User to value and selects it for sending
func (in *ActionOutageIndexInput) SetUser(value int64) *ActionOutageIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}
// SetVps sets parameter Vps to value and selects it for sending
func (in *ActionOutageIndexInput) SetVps(value int64) *ActionOutageIndexInput {
	in.Vps = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Vps"] = nil
	return in
}
// SetHandledBy sets parameter HandledBy to value and selects it for sending
func (in *ActionOutageIndexInput) SetHandledBy(value int64) *ActionOutageIndexInput {
	in.HandledBy = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["HandledBy"] = nil
	return in
}
// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionOutageIndexInput) SetEnvironment(value int64) *ActionOutageIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}
// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionOutageIndexInput) SetLocation(value int64) *ActionOutageIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionOutageIndexInput) SetNode(value int64) *ActionOutageIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetEntityName sets parameter EntityName to value and selects it for sending
func (in *ActionOutageIndexInput) SetEntityName(value string) *ActionOutageIndexInput {
	in.EntityName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EntityName"] = nil
	return in
}
// SetEntityId sets parameter EntityId to value and selects it for sending
func (in *ActionOutageIndexInput) SetEntityId(value int64) *ActionOutageIndexInput {
	in.EntityId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["EntityId"] = nil
	return in
}
// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionOutageIndexInput) SetOrder(value string) *ActionOutageIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}
// SetSince sets parameter Since to value and selects it for sending
func (in *ActionOutageIndexInput) SetSince(value string) *ActionOutageIndexInput {
	in.Since = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Since"] = nil
	return in
}

// SelectParameters sets parameters from ActionOutageIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOutageIndexInput) SelectParameters(params ...string) *ActionOutageIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOutageIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionOutageIndexOutput is a type for action output parameters
type ActionOutageIndexOutput struct {
	Id int64 `json:"id"`
	BeginsAt string `json:"begins_at"`
	FinishedAt string `json:"finished_at"`
	Duration int64 `json:"duration"`
	Planned bool `json:"planned"`
	State string `json:"state"`
	Type string `json:"type"`
	EnSummary string `json:"en_summary"`
	EnDescription string `json:"en_description"`
	CsSummary string `json:"cs_summary"`
	CsDescription string `json:"cs_description"`
	Affected bool `json:"affected"`
	AffectedUserCount int64 `json:"affected_user_count"`
	AffectedDirectVpsCount int64 `json:"affected_direct_vps_count"`
	AffectedIndirectVpsCount int64 `json:"affected_indirect_vps_count"`
}


// Type for action response, including envelope
type ActionOutageIndexResponse struct {
	Action *ActionOutageIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Outages []*ActionOutageIndexOutput `json:"outages"`
	}

	// Action output without the namespace
	Output []*ActionOutageIndexOutput
}


// Prepare the action for invocation
func (action *ActionOutageIndex) Prepare() *ActionOutageIndexInvocation {
	return &ActionOutageIndexInvocation{
		Action: action,
		Path: "/v6.0/outages",
	}
}

// ActionOutageIndexInvocation is used to configure action for invocation
type ActionOutageIndexInvocation struct {
	// Pointer to the action
	Action *ActionOutageIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionOutageIndexInput
	// Global meta input parameters
	MetaInput *ActionOutageIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionOutageIndexInvocation) NewInput() *ActionOutageIndexInput {
	inv.Input = &ActionOutageIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionOutageIndexInvocation) SetInput(input *ActionOutageIndexInput) *ActionOutageIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionOutageIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOutageIndexInvocation) NewMetaInput() *ActionOutageIndexMetaGlobalInput {
	inv.MetaInput = &ActionOutageIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOutageIndexInvocation) SetMetaInput(input *ActionOutageIndexMetaGlobalInput) *ActionOutageIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOutageIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionOutageIndexInvocation) Call() (*ActionOutageIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionOutageIndexInvocation) callAsQuery() (*ActionOutageIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionOutageIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Outages
	}
	return resp, err
}



func (inv *ActionOutageIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["outage[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["outage[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Planned") {
			ret["outage[planned]"] = convertBoolToString(inv.Input.Planned)
		}
		if inv.IsParameterSelected("State") {
			ret["outage[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("Type") {
			ret["outage[type]"] = inv.Input.Type
		}
		if inv.IsParameterSelected("Affected") {
			ret["outage[affected]"] = convertBoolToString(inv.Input.Affected)
		}
		if inv.IsParameterSelected("User") {
			ret["outage[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("Vps") {
			ret["outage[vps]"] = convertInt64ToString(inv.Input.Vps)
		}
		if inv.IsParameterSelected("HandledBy") {
			ret["outage[handled_by]"] = convertInt64ToString(inv.Input.HandledBy)
		}
		if inv.IsParameterSelected("Environment") {
			ret["outage[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("Location") {
			ret["outage[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Node") {
			ret["outage[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("EntityName") {
			ret["outage[entity_name]"] = inv.Input.EntityName
		}
		if inv.IsParameterSelected("EntityId") {
			ret["outage[entity_id]"] = convertInt64ToString(inv.Input.EntityId)
		}
		if inv.IsParameterSelected("Order") {
			ret["outage[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("Since") {
			ret["outage[since]"] = inv.Input.Since
		}
	}
}

func (inv *ActionOutageIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

