package client

import ()

// ActionTransactionIndex is a type for action Transaction#Index
type ActionTransactionIndex struct {
	// Pointer to client
	Client *Client
}

func NewActionTransactionIndex(client *Client) *ActionTransactionIndex {
	return &ActionTransactionIndex{
		Client: client,
	}
}

// ActionTransactionIndexMetaGlobalInput is a type for action global meta input parameters
type ActionTransactionIndexMetaGlobalInput struct {
	Count    bool   "json:\"count\""
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetCount sets parameter Count to value and selects it for sending
func (in *ActionTransactionIndexMetaGlobalInput) SetCount(value bool) *ActionTransactionIndexMetaGlobalInput {
	in.Count = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Count"] = nil
	return in
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionTransactionIndexMetaGlobalInput) SetIncludes(value string) *ActionTransactionIndexMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionTransactionIndexMetaGlobalInput) SetNo(value bool) *ActionTransactionIndexMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionTransactionIndexMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionTransactionIndexMetaGlobalInput) SelectParameters(params ...string) *ActionTransactionIndexMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionTransactionIndexMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionTransactionIndexInput is a type for action input parameters
type ActionTransactionIndexInput struct {
	Done             string "json:\"done\""
	FromId           int64  "json:\"from_id\""
	Limit            int64  "json:\"limit\""
	Node             int64  "json:\"node\""
	Success          int64  "json:\"success\""
	TransactionChain int64  "json:\"transaction_chain\""
	Type             int64  "json:\"type\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetDone sets parameter Done to value and selects it for sending
func (in *ActionTransactionIndexInput) SetDone(value string) *ActionTransactionIndexInput {
	in.Done = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Done"] = nil
	return in
}

// SetFromId sets parameter FromId to value and selects it for sending
func (in *ActionTransactionIndexInput) SetFromId(value int64) *ActionTransactionIndexInput {
	in.FromId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["FromId"] = nil
	return in
}

// SetLimit sets parameter Limit to value and selects it for sending
func (in *ActionTransactionIndexInput) SetLimit(value int64) *ActionTransactionIndexInput {
	in.Limit = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Limit"] = nil
	return in
}

// SetNode sets parameter Node to value and selects it for sending
func (in *ActionTransactionIndexInput) SetNode(value int64) *ActionTransactionIndexInput {
	in.Node = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Node"] = nil
	return in
}

// SetSuccess sets parameter Success to value and selects it for sending
func (in *ActionTransactionIndexInput) SetSuccess(value int64) *ActionTransactionIndexInput {
	in.Success = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Success"] = nil
	return in
}

// SetTransactionChain sets parameter TransactionChain to value and selects it for sending
func (in *ActionTransactionIndexInput) SetTransactionChain(value int64) *ActionTransactionIndexInput {
	in.TransactionChain = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["TransactionChain"] = nil
	return in
}

// SetType sets parameter Type to value and selects it for sending
func (in *ActionTransactionIndexInput) SetType(value int64) *ActionTransactionIndexInput {
	in.Type = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Type"] = nil
	return in
}

// SelectParameters sets parameters from ActionTransactionIndexInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionTransactionIndexInput) SelectParameters(params ...string) *ActionTransactionIndexInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionTransactionIndexInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionTransactionIndexInput) UnselectParameters(params ...string) *ActionTransactionIndexInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionTransactionIndexInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionTransactionIndexOutput is a type for action output parameters
type ActionTransactionIndexOutput struct {
	CreatedAt        string                            "json:\"created_at\""
	DependsOn        *ActionTransactionShowOutput      "json:\"depends_on\""
	Done             string                            "json:\"done\""
	FinishedAt       string                            "json:\"finished_at\""
	Id               int64                             "json:\"id\""
	Input            string                            "json:\"input\""
	Label            string                            "json:\"label\""
	Name             string                            "json:\"name\""
	Node             *ActionNodeShowOutput             "json:\"node\""
	Output           string                            "json:\"output\""
	Priority         int64                             "json:\"priority\""
	StartedAt        string                            "json:\"started_at\""
	Success          int64                             "json:\"success\""
	TransactionChain *ActionTransactionChainShowOutput "json:\"transaction_chain\""
	Type             int64                             "json:\"type\""
	Urgent           bool                              "json:\"urgent\""
	User             *ActionUserShowOutput             "json:\"user\""
	Vps              *ActionVpsShowOutput              "json:\"vps\""
}

// Type for action response, including envelope
type ActionTransactionIndexResponse struct {
	Action *ActionTransactionIndex "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Transactions []*ActionTransactionIndexOutput "json:\"transactions\""
	}

	// Action output without the namespace
	Output []*ActionTransactionIndexOutput
}

// Prepare the action for invocation
func (action *ActionTransactionIndex) Prepare() *ActionTransactionIndexInvocation {
	return &ActionTransactionIndexInvocation{
		Action: action,
		Path:   "/v7.0/transactions",
	}
}

// ActionTransactionIndexInvocation is used to configure action for invocation
type ActionTransactionIndexInvocation struct {
	// Pointer to the action
	Action *ActionTransactionIndex

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionTransactionIndexInput
	// Global meta input parameters
	MetaInput *ActionTransactionIndexMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionTransactionIndexInvocation) NewInput() *ActionTransactionIndexInput {
	inv.Input = &ActionTransactionIndexInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionTransactionIndexInvocation) SetInput(input *ActionTransactionIndexInput) *ActionTransactionIndexInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionTransactionIndexInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionTransactionIndexInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionTransactionIndexInvocation) NewMetaInput() *ActionTransactionIndexMetaGlobalInput {
	inv.MetaInput = &ActionTransactionIndexMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionTransactionIndexInvocation) SetMetaInput(input *ActionTransactionIndexMetaGlobalInput) *ActionTransactionIndexInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionTransactionIndexInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionTransactionIndexInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionTransactionIndexInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("Node") {
			if !inv.IsParameterNil("Node") {
				if inv.Input.Node < 0 {
					verr.Add("node", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("TransactionChain") {
			if !inv.IsParameterNil("TransactionChain") {
				if inv.Input.TransactionChain < 0 {
					verr.Add("transaction_chain", "not a valid resource id")
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
func (inv *ActionTransactionIndexInvocation) Call() (*ActionTransactionIndexResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionTransactionIndexInvocation) callAsQuery() (*ActionTransactionIndexResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionTransactionIndexResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Transactions
	}
	return resp, err
}

func (inv *ActionTransactionIndexInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("Done") {
			ret["transaction[done]"] = inv.Input.Done
		}
		if inv.IsParameterSelected("FromId") {
			ret["transaction[from_id]"] = convertInt64ToString(inv.Input.FromId)
		}
		if inv.IsParameterSelected("Limit") {
			ret["transaction[limit]"] = convertInt64ToString(inv.Input.Limit)
		}
		if inv.IsParameterSelected("Node") {
			ret["transaction[node]"] = convertInt64ToString(inv.Input.Node)
		}
		if inv.IsParameterSelected("Success") {
			ret["transaction[success]"] = convertInt64ToString(inv.Input.Success)
		}
		if inv.IsParameterSelected("TransactionChain") {
			ret["transaction[transaction_chain]"] = convertInt64ToString(inv.Input.TransactionChain)
		}
		if inv.IsParameterSelected("Type") {
			ret["transaction[type]"] = convertInt64ToString(inv.Input.Type)
		}
	}

	return nil
}

func (inv *ActionTransactionIndexInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
