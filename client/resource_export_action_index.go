package client

import ()

// ActionExportIndex is a type for action Export#Index
type ActionExportIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionExportIndex(client *Client) *ActionExportIndex {
	return &ActionExportIndex{
		Client: client,
	}
}

// ActionExportIndexMetaGlobalInput is a type for action global meta input parameters
type ActionExportIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionExportIndexMetaGlobalInput) SetCount(value bool) *ActionExportIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportIndexMetaGlobalInput) SetIncludes(value string) *ActionExportIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportIndexMetaGlobalInput) SetNo(value bool) *ActionExportIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportIndexMetaGlobalInput) SelectParameters(params ...string) *ActionExportIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportIndexInput is a type for action input parameters
type ActionExportIndexInput struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	User   int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionExportIndexInput) SetLimit(value int64) *ActionExportIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionExportIndexInput) SetOffset(value int64) *ActionExportIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionExportIndexInput) SetUser(value int64) *ActionExportIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionExportIndexInput) SetUserNil(set bool) *ActionExportIndexInput {
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

// SelectParameters sets parameters from ActionExportIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportIndexInput) SelectParameters(params ...string) *ActionExportIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionExportIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionExportIndexInput) UnselectParameters(params ...string) *ActionExportIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionExportIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportIndexOutput is a type for action output parameters
type ActionExportIndexOutput struct {
	AllVps         bool                             `json:"all_vps"`
	CreatedAt      string                           `json:"created_at"`
	Dataset        *ActionDatasetShowOutput         `json:"dataset"`
	Enabled        bool                             `json:"enabled"`
	ExpirationDate string                           `json:"expiration_date"`
	HostIpAddress  *ActionHostIpAddressShowOutput   `json:"host_ip_address"`
	Id             int64                            `json:"id"`
	IpAddress      *ActionIpAddressShowOutput       `json:"ip_address"`
	Path           string                           `json:"path"`
	RootSquash     bool                             `json:"root_squash"`
	Rw             bool                             `json:"rw"`
	Snapshot       *ActionDatasetSnapshotShowOutput `json:"snapshot"`
	SubtreeCheck   bool                             `json:"subtree_check"`
	Sync           bool                             `json:"sync"`
	Threads        int64                            `json:"threads"`
	UpdatedAt      string                           `json:"updated_at"`
	User           *ActionUserShowOutput            `json:"user"`
}

// Type for action response, including envelope
type ActionExportIndexResponse struct {
	Action *ActionExportIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Exports []*ActionExportIndexOutput `json:"exports"`
	}

	// Action output without the namespace
	Output []*ActionExportIndexOutput
}

// Prepare the action for invocation
func (action *ActionExportIndex) Prepare() *ActionExportIndexInvocation {
	return &ActionExportIndexInvocation{
		Action: action,
		Path:   "/v6.0/exports",
	}
}

// ActionExportIndexInvocation is used to configure action for invocation
type ActionExportIndexInvocation struct {
	// Pointer to the action
	Action *ActionExportIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionExportIndexInput
	// Global meta input parameters
	MetaInput *ActionExportIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionExportIndexInvocation) NewInput() *ActionExportIndexInput {
	inv.Input = &ActionExportIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionExportIndexInvocation) SetInput(input *ActionExportIndexInput) *ActionExportIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionExportIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionExportIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportIndexInvocation) NewMetaInput() *ActionExportIndexMetaGlobalInput {
	inv.MetaInput = &ActionExportIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportIndexInvocation) SetMetaInput(input *ActionExportIndexMetaGlobalInput) *ActionExportIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionExportIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionExportIndexInvocation) Call() (*ActionExportIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionExportIndexInvocation) callAsQuery() (*ActionExportIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionExportIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Exports
	}
	return resp, err
}

func (inv *ActionExportIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Limit") {
			ret["export[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Offset") {
			ret["export[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("User") {
			ret["export[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionExportIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
