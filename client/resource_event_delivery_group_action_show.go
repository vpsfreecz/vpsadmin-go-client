package client

import (
	"net/url"
	"strings"
)

// ActionEventDeliveryGroupShow is a type for action Event_delivery_group#Show
type ActionEventDeliveryGroupShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEventDeliveryGroupShow(client *Client) *ActionEventDeliveryGroupShow {
	return &ActionEventDeliveryGroupShow{
		Client: client,
	}
}

// ActionEventDeliveryGroupShowMetaGlobalInput is a type for action global meta input parameters
type ActionEventDeliveryGroupShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventDeliveryGroupShowMetaGlobalInput) SetIncludes(value string) *ActionEventDeliveryGroupShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventDeliveryGroupShowMetaGlobalInput) SetNo(value bool) *ActionEventDeliveryGroupShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventDeliveryGroupShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventDeliveryGroupShowMetaGlobalInput) SelectParameters(params ...string) *ActionEventDeliveryGroupShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventDeliveryGroupShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventDeliveryGroupShowOutput is a type for action output parameters
type ActionEventDeliveryGroupShowOutput struct {
	Actions                   interface{} "json:\"actions\""
	CreatedAt                 string      "json:\"created_at\""
	EventCount                int64       "json:\"event_count\""
	EventRouteId              int64       "json:\"event_route_id\""
	EventRouteLabel           string      "json:\"event_route_label\""
	GroupBy                   interface{} "json:\"group_by\""
	GroupIntervalSeconds      int64       "json:\"group_interval_seconds\""
	GroupKey                  string      "json:\"group_key\""
	GroupWaitSeconds          int64       "json:\"group_wait_seconds\""
	Id                        int64       "json:\"id\""
	Labels                    interface{} "json:\"labels\""
	LastSealedAt              string      "json:\"last_sealed_at\""
	NextFlushAt               string      "json:\"next_flush_at\""
	NotificationReceiverId    int64       "json:\"notification_receiver_id\""
	NotificationReceiverLabel string      "json:\"notification_receiver_label\""
	PendingEventCount         int64       "json:\"pending_event_count\""
	RouteOwnerId              int64       "json:\"route_owner_id\""
	RouteOwnerLogin           string      "json:\"route_owner_login\""
	State                     string      "json:\"state\""
	StreamCount               int64       "json:\"stream_count\""
	UpdatedAt                 string      "json:\"updated_at\""
}

// Type for action response, including envelope
type ActionEventDeliveryGroupShowResponse struct {
	Action *ActionEventDeliveryGroupShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		EventDeliveryGroup *ActionEventDeliveryGroupShowOutput "json:\"event_delivery_group\""
	}

	// Action output without the namespace
	Output *ActionEventDeliveryGroupShowOutput
}

// Prepare the action for invocation
func (action *ActionEventDeliveryGroupShow) Prepare() *ActionEventDeliveryGroupShowInvocation {
	return &ActionEventDeliveryGroupShowInvocation{
		Action: action,
		Path:   "/v7.0/event_delivery_groups/{event_delivery_group_id}",
	}
}

// ActionEventDeliveryGroupShowInvocation is used to configure action for invocation
type ActionEventDeliveryGroupShowInvocation struct {
	// Pointer to the action
	Action *ActionEventDeliveryGroupShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventDeliveryGroupShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventDeliveryGroupShowInvocation) SetPathParamInt(param string, value int64) *ActionEventDeliveryGroupShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventDeliveryGroupShowInvocation) SetPathParamString(param string, value string) *ActionEventDeliveryGroupShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventDeliveryGroupShowInvocation) NewMetaInput() *ActionEventDeliveryGroupShowMetaGlobalInput {
	inv.MetaInput = &ActionEventDeliveryGroupShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventDeliveryGroupShowInvocation) SetMetaInput(input *ActionEventDeliveryGroupShowMetaGlobalInput) *ActionEventDeliveryGroupShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventDeliveryGroupShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventDeliveryGroupShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventDeliveryGroupShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventDeliveryGroupShowInvocation) Call() (*ActionEventDeliveryGroupShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventDeliveryGroupShowInvocation) callAsQuery() (*ActionEventDeliveryGroupShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventDeliveryGroupShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.EventDeliveryGroup
	}
	return resp, err
}

func (inv *ActionEventDeliveryGroupShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
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
