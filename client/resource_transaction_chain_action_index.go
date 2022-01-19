package client

import ()

// ActionTransactionChainIndex is a type for action Transaction_chain#Index
type ActionTransactionChainIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionTransactionChainIndex(client *Client) *ActionTransactionChainIndex {
	return &ActionTransactionChainIndex{
		Client: client,
	}
}

// ActionTransactionChainIndexMetaGlobalInput is a type for action global meta input parameters
type ActionTransactionChainIndexMetaGlobalInput struct {
	Count    bool   `json:"count"`
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionTransactionChainIndexMetaGlobalInput) SetCount(value bool) *ActionTransactionChainIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionTransactionChainIndexMetaGlobalInput) SetIncludes(value string) *ActionTransactionChainIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionTransactionChainIndexMetaGlobalInput) SetNo(value bool) *ActionTransactionChainIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionTransactionChainIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionTransactionChainIndexMetaGlobalInput) SelectParameters(params ...string) *ActionTransactionChainIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionTransactionChainIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionTransactionChainIndexInput is a type for action input parameters
type ActionTransactionChainIndexInput struct {
	ClassName   string `json:"class_name"`
	Limit       int64  `json:"limit"`
	Name        string `json:"name"`
	Offset      int64  `json:"offset"`
	RowId       int64  `json:"row_id"`
	State       string `json:"state"`
	User        int64  `json:"user"`
	UserSession int64  `json:"user_session"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetClassName sets parameter ClassName to value and selects it for sending
func (in *ActionTransactionChainIndexInput) SetClassName(value string) *ActionTransactionChainIndexInput {
	in.ClassName = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["ClassName"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionTransactionChainIndexInput) SetLimit(value int64) *ActionTransactionChainIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetName sets parameter Name to value and selects it for sending
func (in *ActionTransactionChainIndexInput) SetName(value string) *ActionTransactionChainIndexInput {
	in.Name = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Name"] = nil
	return in
}

// SetOffset sets parameter Offset to value and selects it for sending
func (in *ActionTransactionChainIndexInput) SetOffset(value int64) *ActionTransactionChainIndexInput {
	in.Offset = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Offset"] = nil
	return in
}

// SetRowId sets parameter RowId to value and selects it for sending
func (in *ActionTransactionChainIndexInput) SetRowId(value int64) *ActionTransactionChainIndexInput {
	in.RowId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["RowId"] = nil
	return in
}

// SetState sets parameter State to value and selects it for sending
func (in *ActionTransactionChainIndexInput) SetState(value string) *ActionTransactionChainIndexInput {
	in.State = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["State"] = nil
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionTransactionChainIndexInput) SetUser(value int64) *ActionTransactionChainIndexInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["User"] = nil
	return in
}

// SetUserSession sets parameter UserSession to value and selects it for sending
func (in *ActionTransactionChainIndexInput) SetUserSession(value int64) *ActionTransactionChainIndexInput {
	in.UserSession = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["UserSession"] = nil
	return in
}

// SelectParameters sets parameters from ActionTransactionChainIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionTransactionChainIndexInput) SelectParameters(params ...string) *ActionTransactionChainIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionTransactionChainIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionTransactionChainIndexOutput is a type for action output parameters
type ActionTransactionChainIndexOutput struct {
	CreatedAt string                `json:"created_at"`
	Id        int64                 `json:"id"`
	Label     string                `json:"label"`
	Name      string                `json:"name"`
	Progress  int64                 `json:"progress"`
	Size      int64                 `json:"size"`
	State     string                `json:"state"`
	User      *ActionUserShowOutput `json:"user"`
}

// Type for action response, including envelope
type ActionTransactionChainIndexResponse struct {
	Action *ActionTransactionChainIndex `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		TransactionChains []*ActionTransactionChainIndexOutput `json:"transaction_chains"`
	}

	// Action output without the namespace
	Output []*ActionTransactionChainIndexOutput
}

// Prepare the action for invocation
func (action *ActionTransactionChainIndex) Prepare() *ActionTransactionChainIndexInvocation {
	return &ActionTransactionChainIndexInvocation{
		Action: action,
		Path:   "/v6.0/transaction_chains",
	}
}

// ActionTransactionChainIndexInvocation is used to configure action for invocation
type ActionTransactionChainIndexInvocation struct {
	// Pointer to the action
	Action *ActionTransactionChainIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionTransactionChainIndexInput
	// Global meta input parameters
	MetaInput *ActionTransactionChainIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionTransactionChainIndexInvocation) NewInput() *ActionTransactionChainIndexInput {
	inv.Input = &ActionTransactionChainIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionTransactionChainIndexInvocation) SetInput(input *ActionTransactionChainIndexInput) *ActionTransactionChainIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionTransactionChainIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionTransactionChainIndexInvocation) NewMetaInput() *ActionTransactionChainIndexMetaGlobalInput {
	inv.MetaInput = &ActionTransactionChainIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionTransactionChainIndexInvocation) SetMetaInput(input *ActionTransactionChainIndexMetaGlobalInput) *ActionTransactionChainIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionTransactionChainIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionTransactionChainIndexInvocation) Call() (*ActionTransactionChainIndexResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionTransactionChainIndexInvocation) callAsQuery() (*ActionTransactionChainIndexResponse, error) {
	queryParams := make(map[string]string)
	inv.convertInputToQueryParams(queryParams)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionTransactionChainIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.TransactionChains
	}
	return resp, err
}

func (inv *ActionTransactionChainIndexInvocation) convertInputToQueryParams(ret map[string]string) {
	if inv.Input != nil {
		if inv.IsParameterSelected("ClassName") {
			ret["transaction_chain[class_name]"] = inv.Input.ClassName
		}
		if inv.IsParameterSelected("Limit") {
			ret["transaction_chain[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Name") {
			ret["transaction_chain[name]"] = inv.Input.Name
		}
		if inv.IsParameterSelected("Offset") {
			ret["transaction_chain[offset]"] = convertInt64ToString(inv.Input.Offset)
		}
		if inv.IsParameterSelected("RowId") {
			ret["transaction_chain[row_id]"] = convertInt64ToString(inv.Input.RowId)
		}
		if inv.IsParameterSelected("State") {
			ret["transaction_chain[state]"] = inv.Input.State
		}
		if inv.IsParameterSelected("User") {
			ret["transaction_chain[user]"] = convertInt64ToString(inv.Input.User)
		}
		if inv.IsParameterSelected("UserSession") {
			ret["transaction_chain[user_session]"] = convertInt64ToString(inv.Input.UserSession)
		}
	}
}

func (inv *ActionTransactionChainIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) {
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
