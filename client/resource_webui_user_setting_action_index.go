package client

import ()

// ActionWebuiUserSettingIndex is a type for action Webui_user_setting#Index
type ActionWebuiUserSettingIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionWebuiUserSettingIndex(client *Client) *ActionWebuiUserSettingIndex {
	return &ActionWebuiUserSettingIndex{
		Client: client,
	}
}

// ActionWebuiUserSettingIndexMetaGlobalInput is a type for action global meta input parameters
type ActionWebuiUserSettingIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionWebuiUserSettingIndexMetaGlobalInput) SetCount(value bool) *ActionWebuiUserSettingIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionWebuiUserSettingIndexMetaGlobalInput) SetIncludes(value string) *ActionWebuiUserSettingIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionWebuiUserSettingIndexMetaGlobalInput) SetNo(value bool) *ActionWebuiUserSettingIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebuiUserSettingIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebuiUserSettingIndexMetaGlobalInput) SelectParameters(params ...string) *ActionWebuiUserSettingIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionWebuiUserSettingIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebuiUserSettingIndexInput is a type for action input parameters
type ActionWebuiUserSettingIndexInput struct {
	FromId    int64  "json:\"from_id\""
	Key       string "json:\"key\""
	Limit     int64  "json:\"limit\""
	Namespace string "json:\"namespace\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionWebuiUserSettingIndexInput) SetFromId(value int64) *ActionWebuiUserSettingIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetKey sets parameter Key to value and selects it for sending
func (in *ActionWebuiUserSettingIndexInput) SetKey(value string) *ActionWebuiUserSettingIndexInput {
	in.Key = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Key"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionWebuiUserSettingIndexInput) SetLimit(value int64) *ActionWebuiUserSettingIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNamespace sets parameter Namespace to value and selects it for sending
func (in *ActionWebuiUserSettingIndexInput) SetNamespace(value string) *ActionWebuiUserSettingIndexInput {
	in.Namespace = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Namespace"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebuiUserSettingIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebuiUserSettingIndexInput) SelectParameters(params ...string) *ActionWebuiUserSettingIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionWebuiUserSettingIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionWebuiUserSettingIndexInput) UnselectParameters(params ...string) *ActionWebuiUserSettingIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionWebuiUserSettingIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebuiUserSettingIndexOutput is a type for action output parameters
type ActionWebuiUserSettingIndexOutput struct {
	CreatedAt string                "json:\"created_at\""
	Id        int64                 "json:\"id\""
	Key       string                "json:\"key\""
	Namespace string                "json:\"namespace\""
	UpdatedAt string                "json:\"updated_at\""
	User      *ActionUserShowOutput "json:\"user\""
	Value     interface{}           "json:\"value\""
}

// Type for action response, including envelope
type ActionWebuiUserSettingIndexResponse struct {
	Action *ActionWebuiUserSettingIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		WebuiUserSettings []*ActionWebuiUserSettingIndexOutput "json:\"webui_user_settings\""
	}

	// Action output without the namespace
	Output []*ActionWebuiUserSettingIndexOutput
}

// Prepare the action for invocation
func (action *ActionWebuiUserSettingIndex) Prepare() *ActionWebuiUserSettingIndexInvocation {
	return &ActionWebuiUserSettingIndexInvocation{
		Action: action,
		Path:   "/v7.0/webui_user_settings",
	}
}

// ActionWebuiUserSettingIndexInvocation is used to configure action for invocation
type ActionWebuiUserSettingIndexInvocation struct {
	// Pointer to the action
	Action *ActionWebuiUserSettingIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionWebuiUserSettingIndexInput
	// Global meta input parameters
	MetaInput *ActionWebuiUserSettingIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionWebuiUserSettingIndexInvocation) NewInput() *ActionWebuiUserSettingIndexInput {
	inv.Input = &ActionWebuiUserSettingIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionWebuiUserSettingIndexInvocation) SetInput(input *ActionWebuiUserSettingIndexInput) *ActionWebuiUserSettingIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionWebuiUserSettingIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionWebuiUserSettingIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionWebuiUserSettingIndexInvocation) NewMetaInput() *ActionWebuiUserSettingIndexMetaGlobalInput {
	inv.MetaInput = &ActionWebuiUserSettingIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionWebuiUserSettingIndexInvocation) SetMetaInput(input *ActionWebuiUserSettingIndexMetaGlobalInput) *ActionWebuiUserSettingIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionWebuiUserSettingIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionWebuiUserSettingIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionWebuiUserSettingIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
	}
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionWebuiUserSettingIndexInvocation) Call() (*ActionWebuiUserSettingIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionWebuiUserSettingIndexInvocation) callAsQuery() (*ActionWebuiUserSettingIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionWebuiUserSettingIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.WebuiUserSettings
	}
	return resp, err
}

func (inv *ActionWebuiUserSettingIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("FromId") {
			ret["webui_user_setting[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Key") {
			ret["webui_user_setting[key]"] = inv.Input.Key
		}
		if inv.IsParameterSelected("Limit") {
			ret["webui_user_setting[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Namespace") {
			ret["webui_user_setting[namespace]"] = inv.Input.Namespace
		}
	}

	return nil
}

func (inv *ActionWebuiUserSettingIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
