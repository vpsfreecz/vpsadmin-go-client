package client

import ()

// ActionMonitoredEventIndex is a type for action Monitored_event#Index
type ActionMonitoredEventIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionMonitoredEventIndex(client *Client) *ActionMonitoredEventIndex {
	return &ActionMonitoredEventIndex{
		Client: client,
	}
}

// ActionMonitoredEventIndexMetaGlobalInput is a type for action global meta input parameters
type ActionMonitoredEventIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionMonitoredEventIndexMetaGlobalInput) SetCount(value bool) *ActionMonitoredEventIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMonitoredEventIndexMetaGlobalInput) SetIncludes(value string) *ActionMonitoredEventIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMonitoredEventIndexMetaGlobalInput) SetNo(value bool) *ActionMonitoredEventIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventIndexMetaGlobalInput) SelectParameters(params ...string) *ActionMonitoredEventIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMonitoredEventIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMonitoredEventIndexInput is a type for action input parameters
type ActionMonitoredEventIndexInput struct {
	FromDuration float64 `json:"from_duration"`
	FromId       int64   `json:"from_id"`
	Limit        int64   `json:"limit"`
	Monitor      string  `json:"monitor"`
	ObjectId     int64   `json:"object_id"`
	ObjectName   string  `json:"object_name"`
	Order        string  `json:"order"`
	State        string  `json:"state"`
	User         int64   `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromDuration sets parameter FromDuration to value and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetFromDuration(value float64) *ActionMonitoredEventIndexInput {
	in.FromDuration = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromDuration"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetFromId(value int64) *ActionMonitoredEventIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetLimit(value int64) *ActionMonitoredEventIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetMonitor sets parameter Monitor to value and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetMonitor(value string) *ActionMonitoredEventIndexInput {
	in.Monitor = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Monitor"] = nil
	return in
}

// SetObjectId sets parameter ObjectId to value and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetObjectId(value int64) *ActionMonitoredEventIndexInput {
	in.ObjectId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectId"] = nil
	return in
}

// SetObjectName sets parameter ObjectName to value and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetObjectName(value string) *ActionMonitoredEventIndexInput {
	in.ObjectName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ObjectName"] = nil
	return in
}

// SetOrder sets parameter Order to value and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetOrder(value string) *ActionMonitoredEventIndexInput {
	in.Order = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Order"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetState(value string) *ActionMonitoredEventIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetUser(value int64) *ActionMonitoredEventIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionMonitoredEventIndexInput) SetUserNil(set bool) *ActionMonitoredEventIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["User"] = nil
		in.SelectParameters("User")
	} else {
		delete(in._nilParameters, "User")
	}
	return in
}

// SelectParameters sets parameters from ActionMonitoredEventIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMonitoredEventIndexInput) SelectParameters(params ...string) *ActionMonitoredEventIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionMonitoredEventIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionMonitoredEventIndexInput) UnselectParameters(params ...string) *ActionMonitoredEventIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionMonitoredEventIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMonitoredEventIndexOutput is a type for action output parameters
type ActionMonitoredEventIndexOutput struct {
	CreatedAt  string                `json:"created_at"`
	Duration   float64               `json:"duration"`
	Id         int64                 `json:"id"`
	Issue      string                `json:"issue"`
	Label      string                `json:"label"`
	Monitor    string                `json:"monitor"`
	ObjectId   int64                 `json:"object_id"`
	ObjectName string                `json:"object_name"`
	SavedUntil string                `json:"saved_until"`
	State      string                `json:"state"`
	UpdatedAt  string                `json:"updated_at"`
	User       *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionMonitoredEventIndexResponse struct {
	Action *ActionMonitoredEventIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MonitoredEvents []*ActionMonitoredEventIndexOutput `json:"monitored_events"`
	}

	// Action output without the namespace
	Output []*ActionMonitoredEventIndexOutput
}

// Prepare the action for invocation
func (action *ActionMonitoredEventIndex) Prepare() *ActionMonitoredEventIndexInvocation {
	return &ActionMonitoredEventIndexInvocation{
		Action: action,
		Path:   "/v7.0/monitored_events",
	}
}

// ActionMonitoredEventIndexInvocation is used to configure action for invocation
type ActionMonitoredEventIndexInvocation struct {
	// Pointer to the action
	Action *ActionMonitoredEventIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionMonitoredEventIndexInput
	// Global meta input parameters
	MetaInput *ActionMonitoredEventIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionMonitoredEventIndexInvocation) NewInput() *ActionMonitoredEventIndexInput {
	inv.Input = &ActionMonitoredEventIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionMonitoredEventIndexInvocation) SetInput(input *ActionMonitoredEventIndexInput) *ActionMonitoredEventIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionMonitoredEventIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionMonitoredEventIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMonitoredEventIndexInvocation) NewMetaInput() *ActionMonitoredEventIndexMetaGlobalInput {
	inv.MetaInput = &ActionMonitoredEventIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMonitoredEventIndexInvocation) SetMetaInput(input *ActionMonitoredEventIndexMetaGlobalInput) *ActionMonitoredEventIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMonitoredEventIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMonitoredEventIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionMonitoredEventIndexInvocation) Call() (*ActionMonitoredEventIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionMonitoredEventIndexInvocation) callAsQuery() (*ActionMonitoredEventIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionMonitoredEventIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.MonitoredEvents
	}
	return resp, err
}

func (inv *ActionMonitoredEventIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromDuration") {
			ret["monitored_event[from_duration]"] = convertFloat64ToString(inv.Input.FromDuration)
		}
		if inv.IsParameterSelected("FromId") {
			ret["monitored_event[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["monitored_event[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Monitor") {
			ret["monitored_event[monitor]"] = inv.Input.Monitor
		}
		if inv.IsParameterSelected("ObjectId") {
			ret["monitored_event[object_id]"] = convertInt64ToString(inv.Input.ObjectId)
		}
		if inv.IsParameterSelected("ObjectName") {
			ret["monitored_event[object_name]"] = inv.Input.ObjectName
		}
		if inv.IsParameterSelected("Order") {
			ret["monitored_event[order]"] = inv.Input.Order
		}
		if inv.IsParameterSelected("State") {
			ret["monitored_event[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("User") {
			ret["monitored_event[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionMonitoredEventIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
