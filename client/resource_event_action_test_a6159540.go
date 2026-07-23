package client

import ()

// ActionEventTest is a type for action Event#Test
type ActionEventTest struct {
	// Pointer to client
	Client *Client
}

func NewActionEventTest(client *Client) *ActionEventTest {
	return &ActionEventTest{
		Client: client,
	}
}

// ActionEventTestMetaGlobalInput is a type for action global meta input parameters
type ActionEventTestMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionEventTestMetaGlobalInput) SetIncludes(value string) *ActionEventTestMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionEventTestMetaGlobalInput) SetNo(value bool) *ActionEventTestMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionEventTestMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTestMetaGlobalInput) SelectParameters(params ...string) *ActionEventTestMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionEventTestMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTestInput is a type for action input parameters
type ActionEventTestInput struct {
	EventType    string "json:\"event_type\""
	PayloadJson  string "json:\"payload_json\""
	Subject      string "json:\"subject\""
	SubjectScope string "json:\"subject_scope\""
	Summary      string "json:\"summary\""
	User         int64  "json:\"user\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetEventType sets parameter EventType to value and selects it for sending
func (in *ActionEventTestInput) SetEventType(value string) *ActionEventTestInput {
	in.EventType = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetEventTypeNil(false)
	in._selectedParameters["EventType"] = nil
	return in
}

// SetEventTypeNil sets parameter EventType to nil and selects it for sending
func (in *ActionEventTestInput) SetEventTypeNil(set bool) *ActionEventTestInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["EventType"] = nil
		in.SelectParameters("EventType")
	} else {
		delete(in._nilParameters, "EventType")
	}
	return in
}

// SetPayloadJson sets parameter PayloadJson to value and selects it for sending
func (in *ActionEventTestInput) SetPayloadJson(value string) *ActionEventTestInput {
	in.PayloadJson = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetPayloadJsonNil(false)
	in._selectedParameters["PayloadJson"] = nil
	return in
}

// SetPayloadJsonNil sets parameter PayloadJson to nil and selects it for sending
func (in *ActionEventTestInput) SetPayloadJsonNil(set bool) *ActionEventTestInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["PayloadJson"] = nil
		in.SelectParameters("PayloadJson")
	} else {
		delete(in._nilParameters, "PayloadJson")
	}
	return in
}

// SetSubject sets parameter Subject to value and selects it for sending
func (in *ActionEventTestInput) SetSubject(value string) *ActionEventTestInput {
	in.Subject = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSubjectNil(false)
	in._selectedParameters["Subject"] = nil
	return in
}

// SetSubjectNil sets parameter Subject to nil and selects it for sending
func (in *ActionEventTestInput) SetSubjectNil(set bool) *ActionEventTestInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Subject"] = nil
		in.SelectParameters("Subject")
	} else {
		delete(in._nilParameters, "Subject")
	}
	return in
}

// SetSubjectScope sets parameter SubjectScope to value and selects it for sending
func (in *ActionEventTestInput) SetSubjectScope(value string) *ActionEventTestInput {
	in.SubjectScope = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSubjectScopeNil(false)
	in._selectedParameters["SubjectScope"] = nil
	return in
}

// SetSubjectScopeNil sets parameter SubjectScope to nil and selects it for sending
func (in *ActionEventTestInput) SetSubjectScopeNil(set bool) *ActionEventTestInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["SubjectScope"] = nil
		in.SelectParameters("SubjectScope")
	} else {
		delete(in._nilParameters, "SubjectScope")
	}
	return in
}

// SetSummary sets parameter Summary to value and selects it for sending
func (in *ActionEventTestInput) SetSummary(value string) *ActionEventTestInput {
	in.Summary = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetSummaryNil(false)
	in._selectedParameters["Summary"] = nil
	return in
}

// SetSummaryNil sets parameter Summary to nil and selects it for sending
func (in *ActionEventTestInput) SetSummaryNil(set bool) *ActionEventTestInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Summary"] = nil
		in.SelectParameters("Summary")
	} else {
		delete(in._nilParameters, "Summary")
	}
	return in
}

// SetUser sets parameter User to value and selects it for sending
func (in *ActionEventTestInput) SetUser(value int64) *ActionEventTestInput {
	in.User = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetUserNil(false)
	in._selectedParameters["User"] = nil
	return in
}

// SetUserNil sets parameter User to nil and selects it for sending
func (in *ActionEventTestInput) SetUserNil(set bool) *ActionEventTestInput {
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

// SelectParameters sets parameters from ActionEventTestInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionEventTestInput) SelectParameters(params ...string) *ActionEventTestInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionEventTestInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionEventTestInput) UnselectParameters(params ...string) *ActionEventTestInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionEventTestInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionEventTestRequest is a type for the entire action request
type ActionEventTestRequest struct {
	Event map[string]interface{} "json:\"event\""
	Meta  map[string]interface{} "json:\"_meta\""
}

// ActionEventTestOutput is a type for action output parameters
type ActionEventTestOutput struct {
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
type ActionEventTestResponse struct {
	Action *ActionEventTest "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Event *ActionEventTestOutput "json:\"event\""
	}

	// Action output without the namespace
	Output *ActionEventTestOutput
}

// Prepare the action for invocation
func (action *ActionEventTest) Prepare() *ActionEventTestInvocation {
	return &ActionEventTestInvocation{
		Action: action,
		Path:   "/v7.0/events/test",
	}
}

// ActionEventTestInvocation is used to configure action for invocation
type ActionEventTestInvocation struct {
	// Pointer to the action
	Action *ActionEventTest

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionEventTestInput
	// Global meta input parameters
	MetaInput *ActionEventTestMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionEventTestInvocation) NewInput() *ActionEventTestInput {
	inv.Input = &ActionEventTestInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionEventTestInvocation) SetInput(input *ActionEventTestInput) *ActionEventTestInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionEventTestInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionEventTestInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionEventTestInvocation) NewMetaInput() *ActionEventTestMetaGlobalInput {
	inv.MetaInput = &ActionEventTestMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionEventTestInvocation) SetMetaInput(input *ActionEventTestMetaGlobalInput) *ActionEventTestInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionEventTestInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionEventTestInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionEventTestInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
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
func (inv *ActionEventTestInvocation) Call() (*ActionEventTestResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsBody()
}

func (inv *ActionEventTestInvocation) callAsBody() (*ActionEventTestResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionEventTestResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Event
	}
	return resp, err
}

func (inv *ActionEventTestInvocation) makeAllInputParams() *ActionEventTestRequest {
	return &ActionEventTestRequest{
		Event: inv.makeInputParams(),
		Meta:  inv.makeMetaInputParams(),
	}
}

func (inv *ActionEventTestInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("EventType") {
			if inv.IsParameterNil("EventType") {
				ret["event_type"] = nil
			} else {
				ret["event_type"] = inv.Input.EventType
			}
		}
		if inv.IsParameterSelected("PayloadJson") {
			if inv.IsParameterNil("PayloadJson") {
				ret["payload_json"] = nil
			} else {
				ret["payload_json"] = inv.Input.PayloadJson
			}
		}
		if inv.IsParameterSelected("Subject") {
			if inv.IsParameterNil("Subject") {
				ret["subject"] = nil
			} else {
				ret["subject"] = inv.Input.Subject
			}
		}
		if inv.IsParameterSelected("SubjectScope") {
			if inv.IsParameterNil("SubjectScope") {
				ret["subject_scope"] = nil
			} else {
				ret["subject_scope"] = inv.Input.SubjectScope
			}
		}
		if inv.IsParameterSelected("Summary") {
			if inv.IsParameterNil("Summary") {
				ret["summary"] = nil
			} else {
				ret["summary"] = inv.Input.Summary
			}
		}
		if inv.IsParameterSelected("User") {
			if inv.IsParameterNil("User") {
				ret["user"] = nil
			} else {
				ret["user"] = inv.Input.User
			}
		}
	}

	return ret
}

func (inv *ActionEventTestInvocation) makeMetaInputParams() map[string]interface{} {
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
