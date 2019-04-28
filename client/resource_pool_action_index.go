package client

import (
)

// ActionPoolIndex is a type for action Pool#Index
type ActionPoolIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionPoolIndex(client *Client) *ActionPoolIndex {
	return &ActionPoolIndex{
		Client: client,
	}
}

// ActionPoolIndexMetaGlobalInput is a type for action global meta input parameters
type ActionPoolIndexMetaGlobalInput struct {
	No bool `json:"no"`
	Count bool `json:"count"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionPoolIndexMetaGlobalInput) SetNo(value bool) *ActionPoolIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetCount sets parameter Count to value and selects it for sending
func (in *ActionPoolIndexMetaGlobalInput) SetCount(value bool) *ActionPoolIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionPoolIndexMetaGlobalInput) SetIncludes(value string) *ActionPoolIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionPoolIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPoolIndexMetaGlobalInput) SelectParameters(params ...string) *ActionPoolIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPoolIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionPoolIndexInput is a type for action input parameters
type ActionPoolIndexInput struct {
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	Node int64 `json:"node"`
	Label string `json:"label"`
	Filesystem string `json:"filesystem"`
	Role string `json:"role"`
	RefquotaCheck bool `json:"refquota_check"`
	Atime bool `json:"atime"`
	Compression bool `json:"compression"`
	Recordsize int64 `json:"recordsize"`
	Quota int64 `json:"quota"`
	Refquota int64 `json:"refquota"`
	Relatime bool `json:"relatime"`
	Sync string `json:"sync"`
	Sharenfs string `json:"sharenfs"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionPoolIndexInput) SetOffset(value int64) *ActionPoolIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}
// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionPoolIndexInput) SetLimit(value int64) *ActionPoolIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}
// SetNode sets parameter Node to value and selects it for sending
func (in *ActionPoolIndexInput) SetNode(value int64) *ActionPoolIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}
// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionPoolIndexInput) SetLabel(value string) *ActionPoolIndexInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}
// SetFilesystem sets parameter Filesystem to value and selects it for sending
func (in *ActionPoolIndexInput) SetFilesystem(value string) *ActionPoolIndexInput {
	in.Filesystem = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Filesystem"] = nil
	return in
}
// SetRole sets parameter Role to value and selects it for sending
func (in *ActionPoolIndexInput) SetRole(value string) *ActionPoolIndexInput {
	in.Role = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Role"] = nil
	return in
}
// SetRefquotaCheck sets parameter RefquotaCheck to value and selects it for sending
func (in *ActionPoolIndexInput) SetRefquotaCheck(value bool) *ActionPoolIndexInput {
	in.RefquotaCheck = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RefquotaCheck"] = nil
	return in
}
// SetAtime sets parameter Atime to value and selects it for sending
func (in *ActionPoolIndexInput) SetAtime(value bool) *ActionPoolIndexInput {
	in.Atime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Atime"] = nil
	return in
}
// SetCompression sets parameter Compression to value and selects it for sending
func (in *ActionPoolIndexInput) SetCompression(value bool) *ActionPoolIndexInput {
	in.Compression = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Compression"] = nil
	return in
}
// SetRecordsize sets parameter Recordsize to value and selects it for sending
func (in *ActionPoolIndexInput) SetRecordsize(value int64) *ActionPoolIndexInput {
	in.Recordsize = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Recordsize"] = nil
	return in
}
// SetQuota sets parameter Quota to value and selects it for sending
func (in *ActionPoolIndexInput) SetQuota(value int64) *ActionPoolIndexInput {
	in.Quota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Quota"] = nil
	return in
}
// SetRefquota sets parameter Refquota to value and selects it for sending
func (in *ActionPoolIndexInput) SetRefquota(value int64) *ActionPoolIndexInput {
	in.Refquota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Refquota"] = nil
	return in
}
// SetRelatime sets parameter Relatime to value and selects it for sending
func (in *ActionPoolIndexInput) SetRelatime(value bool) *ActionPoolIndexInput {
	in.Relatime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Relatime"] = nil
	return in
}
// SetSync sets parameter Sync to value and selects it for sending
func (in *ActionPoolIndexInput) SetSync(value string) *ActionPoolIndexInput {
	in.Sync = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sync"] = nil
	return in
}
// SetSharenfs sets parameter Sharenfs to value and selects it for sending
func (in *ActionPoolIndexInput) SetSharenfs(value string) *ActionPoolIndexInput {
	in.Sharenfs = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Sharenfs"] = nil
	return in
}

// SelectParameters sets parameters from ActionPoolIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPoolIndexInput) SelectParameters(params ...string) *ActionPoolIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPoolIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}


// ActionPoolIndexOutput is a type for action output parameters
type ActionPoolIndexOutput struct {
	Id int64 `json:"id"`
	Node *ActionNodeShowOutput `json:"node"`
	Label string `json:"label"`
	Filesystem string `json:"filesystem"`
	Role string `json:"role"`
	RefquotaCheck bool `json:"refquota_check"`
	Atime bool `json:"atime"`
	Compression bool `json:"compression"`
	Recordsize int64 `json:"recordsize"`
	Quota int64 `json:"quota"`
	Refquota int64 `json:"refquota"`
	Relatime bool `json:"relatime"`
	Sync string `json:"sync"`
	Sharenfs string `json:"sharenfs"`
	Used int64 `json:"used"`
	Referenced int64 `json:"referenced"`
	Avail int64 `json:"avail"`
	MaintenanceLock string `json:"maintenance_lock"`
	MaintenanceLockReason string `json:"maintenance_lock_reason"`
}


// Type for action response, including envelope
type ActionPoolIndexResponse struct {
	Action *ActionPoolIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Pools []*ActionPoolIndexOutput `json:"pools"`
	}

	// Action output without the namespace
	Output []*ActionPoolIndexOutput
}


// Prepare the action for invocation
func (action *ActionPoolIndex) Prepare() *ActionPoolIndexInvocation {
	return &ActionPoolIndexInvocation{
		Action: action,
		Path: "/v5.0/pools",
	}
}

// ActionPoolIndexInvocation is used to configure action for invocation
type ActionPoolIndexInvocation struct {
	// Pointer to the action
	Action *ActionPoolIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionPoolIndexInput
	// Global meta input parameters
	MetaInput *ActionPoolIndexMetaGlobalInput
}


// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionPoolIndexInvocation) NewInput() *ActionPoolIndexInput {
	inv.Input = &ActionPoolIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionPoolIndexInvocation) SetInput(input *ActionPoolIndexInput) *ActionPoolIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionPoolIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}
// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionPoolIndexInvocation) NewMetaInput() *ActionPoolIndexMetaGlobalInput {
	inv.MetaInput = &ActionPoolIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionPoolIndexInvocation) SetMetaInput(input *ActionPoolIndexMetaGlobalInput) *ActionPoolIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionPoolIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionPoolIndexInvocation) Call() (*ActionPoolIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionPoolIndexInvocation) callAsQuery() (*ActionPoolIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionPoolIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Pools
	}
	return resp, err
}



func (inv *ActionPoolIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Offset") {
			ret["pool[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Limit") {
			ret["pool[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["pool[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Label") {
			ret["pool[label]"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Filesystem") {
			ret["pool[filesystem]"] = inv.Input.Filesystem
		}
		if inv.IsParameterSelected("Role") {
			ret["pool[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("RefquotaCheck") {
			ret["pool[refquota_check]"] = convertBoolToString(inv.Input.RefquotaCheck)
		}
		if inv.IsParameterSelected("Atime") {
			ret["pool[atime]"] = convertBoolToString(inv.Input.Atime)
		}
		if inv.IsParameterSelected("Compression") {
			ret["pool[compression]"] = convertBoolToString(inv.Input.Compression)
		}
		if inv.IsParameterSelected("Recordsize") {
			ret["pool[recordsize]"] = convertInt64ToString(inv.Input.Recordsize)
		}
		if inv.IsParameterSelected("Quota") {
			ret["pool[quota]"] = convertInt64ToString(inv.Input.Quota)
		}
		if inv.IsParameterSelected("Refquota") {
			ret["pool[refquota]"] = convertInt64ToString(inv.Input.Refquota)
		}
		if inv.IsParameterSelected("Relatime") {
			ret["pool[relatime]"] = convertBoolToString(inv.Input.Relatime)
		}
		if inv.IsParameterSelected("Sync") {
			ret["pool[sync]"] = inv.Input.Sync
		}
		if inv.IsParameterSelected("Sharenfs") {
			ret["pool[sharenfs]"] = inv.Input.Sharenfs
		}
	}
}

func (inv *ActionPoolIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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

