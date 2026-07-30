package client

import (
	"net/url"
	"strings"
)

// ActionTransactionChainNotifyWhenDone is a type for action Transaction_chain#Notify_when_done
type ActionTransactionChainNotifyWhenDone struct {
	// Pointer to client
	Client *Client
}

func NewActionTransactionChainNotifyWhenDone(client *Client) *ActionTransactionChainNotifyWhenDone {
	return &ActionTransactionChainNotifyWhenDone{
		Client: client,
	}
}

// ActionTransactionChainNotifyWhenDoneMetaGlobalInput is a type for action global meta input parameters
type ActionTransactionChainNotifyWhenDoneMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionTransactionChainNotifyWhenDoneMetaGlobalInput) SetIncludes(value string) *ActionTransactionChainNotifyWhenDoneMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionTransactionChainNotifyWhenDoneMetaGlobalInput) SetNo(value bool) *ActionTransactionChainNotifyWhenDoneMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionTransactionChainNotifyWhenDoneMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionTransactionChainNotifyWhenDoneMetaGlobalInput) SelectParameters(params ...string) *ActionTransactionChainNotifyWhenDoneMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionTransactionChainNotifyWhenDoneMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionTransactionChainNotifyWhenDoneInput is a type for action input parameters
type ActionTransactionChainNotifyWhenDoneInput struct {
	ExpiresAt              string "json:\"expires_at\""
	NotificationReceiverId int64  "json:\"notification_receiver_id\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetExpiresAt sets parameter ExpiresAt to value and selects it for sending
func (in *ActionTransactionChainNotifyWhenDoneInput) SetExpiresAt(value string) *ActionTransactionChainNotifyWhenDoneInput {
	in.ExpiresAt = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetExpiresAtNil(false)
	in._selectedParameters["ExpiresAt"] = nil
	return in
}

// SetExpiresAtNil sets parameter ExpiresAt to nil and selects it for sending
func (in *ActionTransactionChainNotifyWhenDoneInput) SetExpiresAtNil(set bool) *ActionTransactionChainNotifyWhenDoneInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["ExpiresAt"] = nil
		in.SelectParameters("ExpiresAt")
	} else {
		delete(in._nilParameters, "ExpiresAt")
	}
	return in
}

// SetNotificationReceiverId sets parameter NotificationReceiverId to value and selects it for sending
func (in *ActionTransactionChainNotifyWhenDoneInput) SetNotificationReceiverId(value int64) *ActionTransactionChainNotifyWhenDoneInput {
	in.NotificationReceiverId = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetNotificationReceiverIdNil(false)
	in._selectedParameters["NotificationReceiverId"] = nil
	return in
}

// SetNotificationReceiverIdNil sets parameter NotificationReceiverId to nil and selects it for sending
func (in *ActionTransactionChainNotifyWhenDoneInput) SetNotificationReceiverIdNil(set bool) *ActionTransactionChainNotifyWhenDoneInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["NotificationReceiverId"] = nil
		in.SelectParameters("NotificationReceiverId")
	} else {
		delete(in._nilParameters, "NotificationReceiverId")
	}
	return in
}

// SelectParameters sets parameters from ActionTransactionChainNotifyWhenDoneInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionTransactionChainNotifyWhenDoneInput) SelectParameters(params ...string) *ActionTransactionChainNotifyWhenDoneInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionTransactionChainNotifyWhenDoneInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionTransactionChainNotifyWhenDoneInput) UnselectParameters(params ...string) *ActionTransactionChainNotifyWhenDoneInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionTransactionChainNotifyWhenDoneInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionTransactionChainNotifyWhenDoneRequest is a type for the entire action request
type ActionTransactionChainNotifyWhenDoneRequest struct {
	TransactionChain map[string]interface{} "json:\"transaction_chain\""
	Meta             map[string]interface{} "json:\"_meta\""
}

// ActionTransactionChainNotifyWhenDoneOutput is a type for action output parameters
type ActionTransactionChainNotifyWhenDoneOutput struct {
	Continue               bool   "json:\"continue\""
	CreatedAt              string "json:\"created_at\""
	DisplayLabel           string "json:\"display_label\""
	Enabled                bool   "json:\"enabled\""
	EventType              string "json:\"event_type\""
	EventTypePattern       string "json:\"event_type_pattern\""
	ExpiresAt              string "json:\"expires_at\""
	HitCount               int64  "json:\"hit_count\""
	Id                     int64  "json:\"id\""
	Label                  string "json:\"label\""
	MatcherSummary         string "json:\"matcher_summary\""
	NotificationReceiverId int64  "json:\"notification_receiver_id\""
	ParentId               int64  "json:\"parent_id\""
	Position               int64  "json:\"position\""
	SingleUse              bool   "json:\"single_use\""
	SpentAt                string "json:\"spent_at\""
	UpdatedAt              string "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionTransactionChainNotifyWhenDoneResponse struct {
	Action *ActionTransactionChainNotifyWhenDone "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventRoute *ActionTransactionChainNotifyWhenDoneOutput "json:\"event_route\""
	}

	// Action output without the namespace
	Output *ActionTransactionChainNotifyWhenDoneOutput
}

// Prepare the action for invocation
func (action *ActionTransactionChainNotifyWhenDone) Prepare() *ActionTransactionChainNotifyWhenDoneInvocation {
	return &ActionTransactionChainNotifyWhenDoneInvocation{
		Action: action,
		Path:   "/v7.0/transaction_chains/{transaction_chain_id}/notify_when_done",
	}
}

// ActionTransactionChainNotifyWhenDoneInvocation is used to configure action for invocation
type ActionTransactionChainNotifyWhenDoneInvocation struct {
	// Pointer to the action
	Action *ActionTransactionChainNotifyWhenDone

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionTransactionChainNotifyWhenDoneInput
	// Global meta input parameters
	MetaInput *ActionTransactionChainNotifyWhenDoneMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) SetPathParamInt(param string, value int64) *ActionTransactionChainNotifyWhenDoneInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) SetPathParamString(param string, value string) *ActionTransactionChainNotifyWhenDoneInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) NewInput() *ActionTransactionChainNotifyWhenDoneInput {
	inv.Input = &ActionTransactionChainNotifyWhenDoneInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) SetInput(input *ActionTransactionChainNotifyWhenDoneInput) *ActionTransactionChainNotifyWhenDoneInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) NewMetaInput() *ActionTransactionChainNotifyWhenDoneMetaGlobalInput {
	inv.MetaInput = &ActionTransactionChainNotifyWhenDoneMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) SetMetaInput(input *ActionTransactionChainNotifyWhenDoneMetaGlobalInput) *ActionTransactionChainNotifyWhenDoneInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionTransactionChainNotifyWhenDoneInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("ExpiresAt") {
			if !inv.IsParameterNil("ExpiresAt") {
				normalized, ok := normalizeAndCheckDatetimeString(inv.Input.ExpiresAt)
				if !ok {
					verr.Add("expires_at", "not a valid datetime")
				} else {
					inv.Input.ExpiresAt = normalized
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
func (inv *ActionTransactionChainNotifyWhenDoneInvocation) Call() (*ActionTransactionChainNotifyWhenDoneResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionTransactionChainNotifyWhenDoneInvocation) callAsBody() (*ActionTransactionChainNotifyWhenDoneResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionTransactionChainNotifyWhenDoneResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventRoute
	}
	return resp, err
}

func (inv *ActionTransactionChainNotifyWhenDoneInvocation) makeAllInputParams() *ActionTransactionChainNotifyWhenDoneRequest {
	return &ActionTransactionChainNotifyWhenDoneRequest{
		TransactionChain: inv.makeInputParams(),
		Meta:             inv.makeMetaInputParams(),
	}
}

func (inv *ActionTransactionChainNotifyWhenDoneInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("ExpiresAt") {
			if inv.IsParameterNil("ExpiresAt") {
				ret["expires_at"] = nil
			} else {
				ret["expires_at"] = inv.Input.ExpiresAt
			}
		}
		if inv.IsParameterSelected("NotificationReceiverId") {
			if inv.IsParameterNil("NotificationReceiverId") {
				ret["notification_receiver_id"] = nil
			} else {
				ret["notification_receiver_id"] = inv.Input.NotificationReceiverId
			}
		}
	}

	return ret
}

func (inv *ActionTransactionChainNotifyWhenDoneInvocation) makeMetaInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["includes"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["no"] = inv.MetaInput.No
		}
	}

	return ret
}
