package client

import ()

// ActionVpsUserDataIndex is a type for action Vps_user_data#Index
type ActionVpsUserDataIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionVpsUserDataIndex(client *Client) *ActionVpsUserDataIndex {
	return &ActionVpsUserDataIndex{
		Client: client,
	}
}

// ActionVpsUserDataIndexMetaGlobalInput is a type for action global meta input parameters
type ActionVpsUserDataIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionVpsUserDataIndexMetaGlobalInput) SetCount(value bool) *ActionVpsUserDataIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionVpsUserDataIndexMetaGlobalInput) SetIncludes(value string) *ActionVpsUserDataIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionVpsUserDataIndexMetaGlobalInput) SetNo(value bool) *ActionVpsUserDataIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionVpsUserDataIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataIndexMetaGlobalInput) SelectParameters(params ...string) *ActionVpsUserDataIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionVpsUserDataIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataIndexInput is a type for action input parameters
type ActionVpsUserDataIndexInput struct {
	Format string `json:"format"`
	FromId int64  `json:"from_id"`
	Limit  int64  `json:"limit"`
	User   int64  `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetFormat sets parameter Format to value and selects it for sending
func (in *ActionVpsUserDataIndexInput) SetFormat(value string) *ActionVpsUserDataIndexInput {
	in.Format = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Format"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionVpsUserDataIndexInput) SetFromId(value int64) *ActionVpsUserDataIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionVpsUserDataIndexInput) SetLimit(value int64) *ActionVpsUserDataIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionVpsUserDataIndexInput) SetUser(value int64) *ActionVpsUserDataIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionVpsUserDataIndexInput) SetUserNil(set bool) *ActionVpsUserDataIndexInput {
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

// SelectParameters sets parameters from ActionVpsUserDataIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionVpsUserDataIndexInput) SelectParameters(params ...string) *ActionVpsUserDataIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionVpsUserDataIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionVpsUserDataIndexInput) UnselectParameters(params ...string) *ActionVpsUserDataIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionVpsUserDataIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionVpsUserDataIndexOutput is a type for action output parameters
type ActionVpsUserDataIndexOutput struct {
	Content   string                `json:"content"`
	CreatedAt string                `json:"created_at"`
	Format    string                `json:"format"`
	Id        int64                 `json:"id"`
	Label     string                `json:"label"`
	UpdatedAt string                `json:"updated_at"`
	User      *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionVpsUserDataIndexResponse struct {
	Action *ActionVpsUserDataIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		VpsUserData []*ActionVpsUserDataIndexOutput `json:"vps_user_data"`
	}

	// Action output without the namespace
	Output []*ActionVpsUserDataIndexOutput
}

// Prepare the action for invocation
func (action *ActionVpsUserDataIndex) Prepare() *ActionVpsUserDataIndexInvocation {
	return &ActionVpsUserDataIndexInvocation{
		Action: action,
		Path:   "/v7.0/vps_user_data",
	}
}

// ActionVpsUserDataIndexInvocation is used to configure action for invocation
type ActionVpsUserDataIndexInvocation struct {
	// Pointer to the action
	Action *ActionVpsUserDataIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionVpsUserDataIndexInput
	// Global meta input parameters
	MetaInput *ActionVpsUserDataIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionVpsUserDataIndexInvocation) NewInput() *ActionVpsUserDataIndexInput {
	inv.Input = &ActionVpsUserDataIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionVpsUserDataIndexInvocation) SetInput(input *ActionVpsUserDataIndexInput) *ActionVpsUserDataIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionVpsUserDataIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionVpsUserDataIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionVpsUserDataIndexInvocation) NewMetaInput() *ActionVpsUserDataIndexMetaGlobalInput {
	inv.MetaInput = &ActionVpsUserDataIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionVpsUserDataIndexInvocation) SetMetaInput(input *ActionVpsUserDataIndexMetaGlobalInput) *ActionVpsUserDataIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionVpsUserDataIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionVpsUserDataIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionVpsUserDataIndexInvocation) Call() (*ActionVpsUserDataIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionVpsUserDataIndexInvocation) callAsQuery() (*ActionVpsUserDataIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionVpsUserDataIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.VpsUserData
	}
	return resp, err
}

func (inv *ActionVpsUserDataIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Format") {
			ret["vps_user_data[format]"] = inv.Input.Format
		}
		if inv.IsParameterSelected("FromId") {
			ret["vps_user_data[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["vps_user_data[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("User") {
			ret["vps_user_data[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionVpsUserDataIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
