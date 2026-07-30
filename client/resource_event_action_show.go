package client

import (
	"net/url"
	"strings"
)

// ActionEventShow is a type for action Event#Show
type ActionEventShow struct {
	// Pointer to client
	Client *Client
}

func NewActionEventShow(client *Client) *ActionEventShow {
	return &ActionEventShow{
		Client: client,
	}
}

// ActionEventShowMetaGlobalInput is a type for action global meta input parameters
type ActionEventShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventShowMetaGlobalInput) SetIncludes(value string) *ActionEventShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventShowMetaGlobalInput) SetNo(value bool) *ActionEventShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventShowMetaGlobalInput) SelectParameters(params ...string) *ActionEventShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventShowOutput is a type for action output parameters
type ActionEventShowOutput struct {
	Category        string                "json:\"category\""
	CreatedAt       string                "json:\"created_at\""
	EventType       string                "json:\"event_type\""
	Id              int64                 "json:\"id\""
	IpAddr          string                "json:\"ip_addr\""
	PayloadJson     string                "json:\"payload_json\""
	RoutingState    string                "json:\"routing_state\""
	Severity        string                "json:\"severity\""
	SourceClass     string                "json:\"source_class\""
	SourceId        int64                 "json:\"source_id\""
	Subject         string                "json:\"subject\""
	SubjectRelation string                "json:\"subject_relation\""
	Summary         string                "json:\"summary\""
	UpdatedAt       string                "json:\"updated_at\""
	User            *ActionUserShowOutput "json:\"user\""
	Vps             *ActionVpsShowOutput  "json:\"vps\""
}

// Type for action response, including envelope
type ActionEventShowResponse struct {
	Action *ActionEventShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Event *ActionEventShowOutput "json:\"event\""
	}

	// Action output without the namespace
	Output *ActionEventShowOutput
}

// Prepare the action for invocation
func (action *ActionEventShow) Prepare() *ActionEventShowInvocation {
	return &ActionEventShowInvocation{
		Action: action,
		Path:   "/v7.0/events/{event_id}",
	}
}

// ActionEventShowInvocation is used to configure action for invocation
type ActionEventShowInvocation struct {
	// Pointer to the action
	Action *ActionEventShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionEventShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionEventShowInvocation) SetPathParamInt(param string, value int64) *ActionEventShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionEventShowInvocation) SetPathParamString(param string, value string) *ActionEventShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventShowInvocation) NewMetaInput() *ActionEventShowMetaGlobalInput {
	inv.MetaInput = &ActionEventShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventShowInvocation) SetMetaInput(input *ActionEventShowMetaGlobalInput) *ActionEventShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventShowInvocation) validate() error {
	verr := NewValidationError()
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionEventShowInvocation) Call() (*ActionEventShowResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionEventShowInvocation) callAsQuery() (*ActionEventShowResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionEventShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Event
	}
	return resp, err
}

func (inv *ActionEventShowInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
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
