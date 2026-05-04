package client

import ()

// ActionExportOutageIndex is a type for action Export_outage#Index
type ActionExportOutageIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionExportOutageIndex(client *Client) *ActionExportOutageIndex {
	return &ActionExportOutageIndex{
		Client: client,
	}
}

// ActionExportOutageIndexMetaGlobalInput is a type for action global meta input parameters
type ActionExportOutageIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionExportOutageIndexMetaGlobalInput) SetCount(value bool) *ActionExportOutageIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportOutageIndexMetaGlobalInput) SetIncludes(value string) *ActionExportOutageIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportOutageIndexMetaGlobalInput) SetNo(value bool) *ActionExportOutageIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportOutageIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportOutageIndexMetaGlobalInput) SelectParameters(params ...string) *ActionExportOutageIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportOutageIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportOutageIndexInput is a type for action input parameters
type ActionExportOutageIndexInput struct {
	Environment int64 `json:"environment"`
	Export      int64 `json:"export"`
	FromId      int64 `json:"from_id"`
	Limit       int64 `json:"limit"`
	Location    int64 `json:"location"`
	Node        int64 `json:"node"`
	Outage      int64 `json:"outage"`
	User        int64 `json:"user"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEnvironment sets parameter Environment to value and selects it for sending
func (in *ActionExportOutageIndexInput) SetEnvironment(value int64) *ActionExportOutageIndexInput {
	in.Environment = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Environment"] = nil
	return in
}

// SetExport sets parameter Export to value and selects it for sending
func (in *ActionExportOutageIndexInput) SetExport(value int64) *ActionExportOutageIndexInput {
	in.Export = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Export"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionExportOutageIndexInput) SetFromId(value int64) *ActionExportOutageIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionExportOutageIndexInput) SetLimit(value int64) *ActionExportOutageIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionExportOutageIndexInput) SetLocation(value int64) *ActionExportOutageIndexInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionExportOutageIndexInput) SetNode(value int64) *ActionExportOutageIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetOutage sets parameter Outage to value and selects it for sending
func (in *ActionExportOutageIndexInput) SetOutage(value int64) *ActionExportOutageIndexInput {
	in.Outage = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Outage"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionExportOutageIndexInput) SetUser(value int64) *ActionExportOutageIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportOutageIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportOutageIndexInput) SelectParameters(params ...string) *ActionExportOutageIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionExportOutageIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionExportOutageIndexInput) UnselectParameters(params ...string) *ActionExportOutageIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionExportOutageIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionExportOutageIndexOutput is a type for action output parameters
type ActionExportOutageIndexOutput struct {
	Environment *ActionEnvironmentShowOutput `json:"environment"`
	Export      *ActionExportShowOutput      `json:"export"`
	Id          int64                        `json:"id"`
	Location    *ActionLocationShowOutput    `json:"location"`
	Node        *ActionNodeShowOutput        `json:"node"`
	Outage      *ActionOutageShowOutput      `json:"outage"`
	User        *ActionUserShowOutput        `json:"user"`
}

// Type for action response, including envelope
type ActionExportOutageIndexResponse struct {
	Action *ActionExportOutageIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		ExportOutages []*ActionExportOutageIndexOutput `json:"export_outages"`
	}

	// Action output without the namespace
	Output []*ActionExportOutageIndexOutput
}

// Prepare the action for invocation
func (action *ActionExportOutageIndex) Prepare() *ActionExportOutageIndexInvocation {
	return &ActionExportOutageIndexInvocation{
		Action: action,
		Path:   "/v7.0/export_outages",
	}
}

// ActionExportOutageIndexInvocation is used to configure action for invocation
type ActionExportOutageIndexInvocation struct {
	// Pointer to the action
	Action *ActionExportOutageIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionExportOutageIndexInput
	// Global meta input parameters
	MetaInput *ActionExportOutageIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionExportOutageIndexInvocation) NewInput() *ActionExportOutageIndexInput {
	inv.Input = &ActionExportOutageIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionExportOutageIndexInvocation) SetInput(input *ActionExportOutageIndexInput) *ActionExportOutageIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionExportOutageIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionExportOutageIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportOutageIndexInvocation) NewMetaInput() *ActionExportOutageIndexMetaGlobalInput {
	inv.MetaInput = &ActionExportOutageIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportOutageIndexInvocation) SetMetaInput(input *ActionExportOutageIndexMetaGlobalInput) *ActionExportOutageIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportOutageIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionExportOutageIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionExportOutageIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			if !inv.IsParameterNil("Environment") {
				if inv.Input.Environment < 0 {
					verr.Add("environment", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Export") {
			if !inv.IsParameterNil("Export") {
				if inv.Input.Export < 0 {
					verr.Add("export", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Location") {
			if !inv.IsParameterNil("Location") {
				if inv.Input.Location < 0 {
					verr.Add("location", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Outage") {
			if !inv.IsParameterNil("Outage") {
				if inv.Input.Outage < 0 {
					verr.Add("outage", "not a valid resource id")
				}
			}
		}
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
func (inv *ActionExportOutageIndexInvocation) Call() (*ActionExportOutageIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionExportOutageIndexInvocation) callAsQuery() (*ActionExportOutageIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionExportOutageIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.ExportOutages
	}
	return resp, err
}

func (inv *ActionExportOutageIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("Environment") {
			ret["export_outage[environment]"] = convertInt64ToString(inv.Input.Environment)
		}
		if inv.IsParameterSelected("Export") {
			ret["export_outage[export]"] = convertInt64ToString(inv.Input.Export)
		}
		if inv.IsParameterSelected("FromId") {
			ret["export_outage[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["export_outage[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Location") {
			ret["export_outage[location]"] = convertInt64ToString(inv.Input.Location)
		}
		if inv.IsParameterSelected("Node") {
			ret["export_outage[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Outage") {
			ret["export_outage[outage]"] = convertInt64ToString(inv.Input.Outage)
		}
		if inv.IsParameterSelected("User") {
			ret["export_outage[user]"] = convertInt64ToString(inv.Input.User)
		}
	}
}

func (inv *ActionExportOutageIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
