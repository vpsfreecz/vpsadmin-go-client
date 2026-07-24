package client

import ()

// ActionNotificationReceiverIndex is a type for action Notification_receiver#Index
type ActionNotificationReceiverIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionNotificationReceiverIndex(client *Client) *ActionNotificationReceiverIndex {
	return &ActionNotificationReceiverIndex{
		Client: client,
	}
}

// ActionNotificationReceiverIndexMetaGlobalInput is a type for action global meta input parameters
type ActionNotificationReceiverIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionNotificationReceiverIndexMetaGlobalInput) SetCount(value bool) *ActionNotificationReceiverIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionNotificationReceiverIndexMetaGlobalInput) SetIncludes(value string) *ActionNotificationReceiverIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionNotificationReceiverIndexMetaGlobalInput) SetNo(value bool) *ActionNotificationReceiverIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverIndexMetaGlobalInput) SelectParameters(params ...string) *ActionNotificationReceiverIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionNotificationReceiverIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverIndexInput is a type for action input parameters
type ActionNotificationReceiverIndexInput struct {
	Enabled bool  "json:\"enabled\""
	FromId  int64 "json:\"from_id\""
	Limit   int64 "json:\"limit\""
	Mute    bool  "json:\"mute\""
	User    int64 "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnabled sets parameter Enabled to value and selects it for sending
func (in *ActionNotificationReceiverIndexInput) SetEnabled(value bool) *ActionNotificationReceiverIndexInput {
	in.Enabled = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Enabled"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionNotificationReceiverIndexInput) SetFromId(value int64) *ActionNotificationReceiverIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionNotificationReceiverIndexInput) SetLimit(value int64) *ActionNotificationReceiverIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetMute sets parameter Mute to value and selects it for sending
func (in *ActionNotificationReceiverIndexInput) SetMute(value bool) *ActionNotificationReceiverIndexInput {
	in.Mute = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Mute"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionNotificationReceiverIndexInput) SetUser(value int64) *ActionNotificationReceiverIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionNotificationReceiverIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionNotificationReceiverIndexInput) SelectParameters(params ...string) *ActionNotificationReceiverIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionNotificationReceiverIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionNotificationReceiverIndexInput) UnselectParameters(params ...string) *ActionNotificationReceiverIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionNotificationReceiverIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionNotificationReceiverIndexOutput is a type for action output parameters
type ActionNotificationReceiverIndexOutput struct {
	CreatedAt            string                "json:\"created_at\""
	Description          string                "json:\"description\""
	DisplayActionSummary string                "json:\"display_action_summary\""
	Enabled              bool                  "json:\"enabled\""
	Id                   int64                 "json:\"id\""
	Label                string                "json:\"label\""
	Mute                 bool                  "json:\"mute\""
	UpdatedAt            string                "json:\"updated_at\""
	User                 *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionNotificationReceiverIndexResponse struct {
	Action *ActionNotificationReceiverIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		NotificationReceivers []*ActionNotificationReceiverIndexOutput "json:\"notification_receivers\""
	}

	// Action output without the namespace
	Output []*ActionNotificationReceiverIndexOutput
}

// Prepare the action for invocation
func (action *ActionNotificationReceiverIndex) Prepare() *ActionNotificationReceiverIndexInvocation {
	return &ActionNotificationReceiverIndexInvocation{
		Action: action,
		Path:   "/v7.0/notification_receivers",
	}
}

// ActionNotificationReceiverIndexInvocation is used to configure action for invocation
type ActionNotificationReceiverIndexInvocation struct {
	// Pointer to the action
	Action *ActionNotificationReceiverIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionNotificationReceiverIndexInput
	// Global meta input parameters
	MetaInput *ActionNotificationReceiverIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionNotificationReceiverIndexInvocation) NewInput() *ActionNotificationReceiverIndexInput {
	inv.Input = &ActionNotificationReceiverIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionNotificationReceiverIndexInvocation) SetInput(input *ActionNotificationReceiverIndexInput) *ActionNotificationReceiverIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionNotificationReceiverIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionNotificationReceiverIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionNotificationReceiverIndexInvocation) NewMetaInput() *ActionNotificationReceiverIndexMetaGlobalInput {
	inv.MetaInput = &ActionNotificationReceiverIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionNotificationReceiverIndexInvocation) SetMetaInput(input *ActionNotificationReceiverIndexMetaGlobalInput) *ActionNotificationReceiverIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionNotificationReceiverIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionNotificationReceiverIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionNotificationReceiverIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("User") {
			if !inv.IsParameterNil("User") {
				if inv.Input.User < 0 {
					verr.Add("user", "not a valid resource id")
				}
			}
		}
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionNotificationReceiverIndexInvocation) Call() (*ActionNotificationReceiverIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionNotificationReceiverIndexInvocation) callAsQuery() (*ActionNotificationReceiverIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionNotificationReceiverIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.NotificationReceivers
	}
	return resp, err
}

func (inv *ActionNotificationReceiverIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("Enabled") {
			ret["notification_receiver[enabled]"] = convertBoolToString(inv.Input.Enabled)
		}
		if inv.IsParameterSelected("FromId") {
			ret["notification_receiver[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["notification_receiver[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Mute") {
			ret["notification_receiver[mute]"] = convertBoolToString(inv.Input.Mute)
		}
		if inv.IsParameterSelected("User") {
			ret["notification_receiver[user]"] = convertInt64ToString(inv.Input.User)
		}
	}

	return nil
}

func (inv *ActionNotificationReceiverIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Count") {
			ret["_meta[count]"] = convertBoolToString(inv.MetaInput.Count)
		}
		if inv.IsMetaParameterSelected("Includes") {
			queryValue, err := convertCustomToString(inv.MetaInput.Includes)
			if err != nil {
				return err
			}
			ret["_meta[includes]"] = queryValue
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
