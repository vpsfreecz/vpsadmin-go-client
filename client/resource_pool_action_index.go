package client

import ()

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
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetNo sets parameter No to value and selects it for sending
func (in *ActionPoolIndexMetaGlobalInput) SetNo(value bool) *ActionPoolIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
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
	Atime         bool    `json:"atime"`
	CheckedAt     string  `json:"checked_at"`
	Compression   bool    `json:"compression"`
	Filesystem    string  `json:"filesystem"`
	Label         string  `json:"label"`
	Limit         int64   `json:"limit"`
	MaxDatasets   int64   `json:"max_datasets"`
	Name          string  `json:"name"`
	Node          int64   `json:"node"`
	Offset        int64   `json:"offset"`
	Quota         int64   `json:"quota"`
	Recordsize    int64   `json:"recordsize"`
	Refquota      int64   `json:"refquota"`
	RefquotaCheck bool    `json:"refquota_check"`
	Relatime      bool    `json:"relatime"`
	Role          string  `json:"role"`
	Scan          string  `json:"scan"`
	ScanPercent   float64 `json:"scan_percent"`
	Sharenfs      string  `json:"sharenfs"`
	State         string  `json:"state"`
	Sync          string  `json:"sync"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
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

// SetCheckedAt sets parameter CheckedAt to value and selects it for sending
func (in *ActionPoolIndexInput) SetCheckedAt(value string) *ActionPoolIndexInput {
	in.CheckedAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["CheckedAt"] = nil
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

// SetFilesystem sets parameter Filesystem to value and selects it for sending
func (in *ActionPoolIndexInput) SetFilesystem(value string) *ActionPoolIndexInput {
	in.Filesystem = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Filesystem"] = nil
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

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionPoolIndexInput) SetLimit(value int64) *ActionPoolIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetMaxDatasets sets parameter MaxDatasets to value and selects it for sending
func (in *ActionPoolIndexInput) SetMaxDatasets(value int64) *ActionPoolIndexInput {
	in.MaxDatasets = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["MaxDatasets"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionPoolIndexInput) SetName(value string) *ActionPoolIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionPoolIndexInput) SetNode(value int64) *ActionPoolIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNodeNil(false)
	in._selectedParameters["Node"] = nil
	return in
}

// SetNodeNil sets parameter Node to nil and selects it for sending
func (in *ActionPoolIndexInput) SetNodeNil(set bool) *ActionPoolIndexInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Node"] = nil
		in.SelectParameters("Node")
	} else {
		delete(in._nilParameters, "Node")
	}
	return in
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

// SetQuota sets parameter Quota to value and selects it for sending
func (in *ActionPoolIndexInput) SetQuota(value int64) *ActionPoolIndexInput {
	in.Quota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Quota"] = nil
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

// SetRefquota sets parameter Refquota to value and selects it for sending
func (in *ActionPoolIndexInput) SetRefquota(value int64) *ActionPoolIndexInput {
	in.Refquota = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Refquota"] = nil
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

// SetRelatime sets parameter Relatime to value and selects it for sending
func (in *ActionPoolIndexInput) SetRelatime(value bool) *ActionPoolIndexInput {
	in.Relatime = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Relatime"] = nil
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

// SetScan sets parameter Scan to value and selects it for sending
func (in *ActionPoolIndexInput) SetScan(value string) *ActionPoolIndexInput {
	in.Scan = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Scan"] = nil
	return in
}

// SetScanPercent sets parameter ScanPercent to value and selects it for sending
func (in *ActionPoolIndexInput) SetScanPercent(value float64) *ActionPoolIndexInput {
	in.ScanPercent = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ScanPercent"] = nil
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

// SetState sets parameter State to value and selects it for sending
func (in *ActionPoolIndexInput) SetState(value string) *ActionPoolIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
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

// UnselectParameters unsets parameters from ActionPoolIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionPoolIndexInput) UnselectParameters(params ...string) *ActionPoolIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
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
	Atime                 bool                  `json:"atime"`
	Avail                 int64                 `json:"avail"`
	CheckedAt             string                `json:"checked_at"`
	Compression           bool                  `json:"compression"`
	Filesystem            string                `json:"filesystem"`
	Id                    int64                 `json:"id"`
	Label                 string                `json:"label"`
	MaintenanceLock       string                `json:"maintenance_lock"`
	MaintenanceLockReason string                `json:"maintenance_lock_reason"`
	MaxDatasets           int64                 `json:"max_datasets"`
	Name                  string                `json:"name"`
	Node                  *ActionNodeShowOutput `json:"node"`
	Quota                 int64                 `json:"quota"`
	Recordsize            int64                 `json:"recordsize"`
	Referenced            int64                 `json:"referenced"`
	Refquota              int64                 `json:"refquota"`
	RefquotaCheck         bool                  `json:"refquota_check"`
	Relatime              bool                  `json:"relatime"`
	Role                  string                `json:"role"`
	Scan                  string                `json:"scan"`
	ScanPercent           float64               `json:"scan_percent"`
	Sharenfs              string                `json:"sharenfs"`
	State                 string                `json:"state"`
	Sync                  string                `json:"sync"`
	Used                  int64                 `json:"used"`
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
		Path:   "/v6.0/pools",
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

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionPoolIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
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

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionPoolIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
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
		if inv.IsParameterSelected("Atime") {
			ret["pool[atime]"] = convertBoolToString(inv.Input.Atime)
		}
		if inv.IsParameterSelected("CheckedAt") {
			ret["pool[checked_at]"] = inv.Input.CheckedAt
		}
		if inv.IsParameterSelected("Compression") {
			ret["pool[compression]"] = convertBoolToString(inv.Input.Compression)
		}
		if inv.IsParameterSelected("Filesystem") {
			ret["pool[filesystem]"] = inv.Input.Filesystem
		}
		if inv.IsParameterSelected("Label") {
			ret["pool[label]"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Limit") {
			ret["pool[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("MaxDatasets") {
			ret["pool[max_datasets]"] = convertInt64ToString(inv.Input.MaxDatasets)
		}
		if inv.IsParameterSelected("Name") {
			ret["pool[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Node") {
			ret["pool[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Offset") {
			ret["pool[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("Quota") {
			ret["pool[quota]"] = convertInt64ToString(inv.Input.Quota)
		}
		if inv.IsParameterSelected("Recordsize") {
			ret["pool[recordsize]"] = convertInt64ToString(inv.Input.Recordsize)
		}
		if inv.IsParameterSelected("Refquota") {
			ret["pool[refquota]"] = convertInt64ToString(inv.Input.Refquota)
		}
		if inv.IsParameterSelected("RefquotaCheck") {
			ret["pool[refquota_check]"] = convertBoolToString(inv.Input.RefquotaCheck)
		}
		if inv.IsParameterSelected("Relatime") {
			ret["pool[relatime]"] = convertBoolToString(inv.Input.Relatime)
		}
		if inv.IsParameterSelected("Role") {
			ret["pool[role]"] = inv.Input.Role
		}
		if inv.IsParameterSelected("Scan") {
			ret["pool[scan]"] = inv.Input.Scan
		}
		if inv.IsParameterSelected("ScanPercent") {
			ret["pool[scan_percent]"] = convertFloat64ToString(inv.Input.ScanPercent)
		}
		if inv.IsParameterSelected("Sharenfs") {
			ret["pool[sharenfs]"] = inv.Input.Sharenfs
		}
		if inv.IsParameterSelected("State") {
			ret["pool[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("Sync") {
			ret["pool[sync]"] = inv.Input.Sync
		}
	}
}

func (inv *ActionPoolIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
