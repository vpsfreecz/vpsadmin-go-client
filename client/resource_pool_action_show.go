package client

import (
	"strings"
)

// ActionPoolShow is a type for action Pool#Show
type ActionPoolShow struct {
	// Pointer to client
	Client *Client
}

func NewActionPoolShow(client *Client) *ActionPoolShow {
	return &ActionPoolShow{
		Client: client,
	}
}

// ActionPoolShowMetaGlobalInput is a type for action global meta input parameters
type ActionPoolShowMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionPoolShowMetaGlobalInput) SetNo(value bool) *ActionPoolShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionPoolShowMetaGlobalInput) SetIncludes(value string) *ActionPoolShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionPoolShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionPoolShowMetaGlobalInput) SelectParameters(params ...string) *ActionPoolShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionPoolShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionPoolShowOutput is a type for action output parameters
type ActionPoolShowOutput struct {
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
type ActionPoolShowResponse struct {
	Action *ActionPoolShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Pool *ActionPoolShowOutput `json:"pool"`
	}

	// Action output without the namespace
	Output *ActionPoolShowOutput
}


// Prepare the action for invocation
func (action *ActionPoolShow) Prepare() *ActionPoolShowInvocation {
	return &ActionPoolShowInvocation{
		Action: action,
		Path: "/v5.0/pools/:pool_id",
	}
}

// ActionPoolShowInvocation is used to configure action for invocation
type ActionPoolShowInvocation struct {
	// Pointer to the action
	Action *ActionPoolShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionPoolShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionPoolShowInvocation) SetPathParamInt(param string, value int64) *ActionPoolShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionPoolShowInvocation) SetPathParamString(param string, value string) *ActionPoolShowInvocation {
	inv.Path = strings.Replace(inv.Path, ":"+param, value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionPoolShowInvocation) NewMetaInput() *ActionPoolShowMetaGlobalInput {
	inv.MetaInput = &ActionPoolShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionPoolShowInvocation) SetMetaInput(input *ActionPoolShowMetaGlobalInput) *ActionPoolShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionPoolShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionPoolShowInvocation) Call() (*ActionPoolShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionPoolShowInvocation) callAsQuery() (*ActionPoolShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionPoolShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Pool
	}
	return resp, err
}




func (inv *ActionPoolShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

