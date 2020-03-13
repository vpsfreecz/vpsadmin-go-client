package client

import (
	"strings"
)

// ActionIntegrityObjectShow is a type for action Integrity_object#Show
type ActionIntegrityObjectShow struct {
	// Pointer to client
	Client *Client
}

func NewActionIntegrityObjectShow(client *Client) *ActionIntegrityObjectShow {
	return &ActionIntegrityObjectShow{
		Client: client,
	}
}

// ActionIntegrityObjectShowMetaGlobalInput is a type for action global meta input parameters
type ActionIntegrityObjectShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionIntegrityObjectShowMetaGlobalInput) SetNo(value bool) *ActionIntegrityObjectShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionIntegrityObjectShowMetaGlobalInput) SetIncludes(value string) *ActionIntegrityObjectShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionIntegrityObjectShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionIntegrityObjectShowMetaGlobalInput) SelectParameters(params ...string) *ActionIntegrityObjectShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionIntegrityObjectShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionIntegrityObjectShowOutput is a type for action output parameters
type ActionIntegrityObjectShowOutput struct {
	Id int64 `json:"id"`
	IntegrityCheck *ActionIntegrityCheckShowOutput `json:"integrity_check"`
	Node *ActionNodeShowOutput `json:"node"`
	ClassName string `json:"class_name"`
	RowId int64 `json:"row_id"`
	Parent *ActionIntegrityObjectShowOutput `json:"parent"`
	Status string `json:"status"`
	CheckedFacts int64 `json:"checked_facts"`
	TrueFacts int64 `json:"true_facts"`
	FalseFacts int64 `json:"false_facts"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Type for action response, including envelope
type ActionIntegrityObjectShowResponse struct {
	Action *ActionIntegrityObjectShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		IntegrityObject *ActionIntegrityObjectShowOutput `json:"integrity_object"`
	}

	// Action output without the namespace
	Output *ActionIntegrityObjectShowOutput
}


// Prepare the action for invocation
func (action *ActionIntegrityObjectShow) Prepare() *ActionIntegrityObjectShowInvocation {
	return &ActionIntegrityObjectShowInvocation{
		Action: action,
		Path: "/v6.0/integrity_objects/{integrity_object_id}",
	}
}

// ActionIntegrityObjectShowInvocation is used to configure action for invocation
type ActionIntegrityObjectShowInvocation struct {
	// Pointer to the action
	Action *ActionIntegrityObjectShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionIntegrityObjectShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionIntegrityObjectShowInvocation) SetPathParamInt(param string, value int64) *ActionIntegrityObjectShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionIntegrityObjectShowInvocation) SetPathParamString(param string, value string) *ActionIntegrityObjectShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionIntegrityObjectShowInvocation) NewMetaInput() *ActionIntegrityObjectShowMetaGlobalInput {
	inv.MetaInput = &ActionIntegrityObjectShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionIntegrityObjectShowInvocation) SetMetaInput(input *ActionIntegrityObjectShowMetaGlobalInput) *ActionIntegrityObjectShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionIntegrityObjectShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionIntegrityObjectShowInvocation) Call() (*ActionIntegrityObjectShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionIntegrityObjectShowInvocation) callAsQuery() (*ActionIntegrityObjectShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionIntegrityObjectShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.IntegrityObject
	}
	return resp, err
}




func (inv *ActionIntegrityObjectShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

