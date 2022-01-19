package client

import (
	"strings"
)

// ActionExportShow is a type for action Export#Show
type ActionExportShow struct {
	// Pointer to client
	Client *Client
}

func NewActionExportShow(client *Client) *ActionExportShow {
	return &ActionExportShow{
		Client: client,
	}
}

// ActionExportShowMetaGlobalInput is a type for action global meta input parameters
type ActionExportShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No bool `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionExportShowMetaGlobalInput) SetIncludes(value string) *ActionExportShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}
// SetNo sets parameter No to value and selects it for sending
func (in *ActionExportShowMetaGlobalInput) SetNo(value bool) *ActionExportShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionExportShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionExportShowMetaGlobalInput) SelectParameters(params ...string) *ActionExportShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionExportShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionExportShowOutput is a type for action output parameters
type ActionExportShowOutput struct {
	AllVps bool `json:"all_vps"`
	CreatedAt string `json:"created_at"`
	Dataset *ActionDatasetShowOutput `json:"dataset"`
	Enabled bool `json:"enabled"`
	ExpirationDate string `json:"expiration_date"`
	HostIpAddress *ActionHostIpAddressShowOutput `json:"host_ip_address"`
	Id int64 `json:"id"`
	IpAddress *ActionIpAddressShowOutput `json:"ip_address"`
	Path string `json:"path"`
	RootSquash bool `json:"root_squash"`
	Rw bool `json:"rw"`
	Snapshot *ActionDatasetSnapshotShowOutput `json:"snapshot"`
	SubtreeCheck bool `json:"subtree_check"`
	Sync bool `json:"sync"`
	Threads int64 `json:"threads"`
	UpdatedAt string `json:"updated_at"`
	User *ActionUserShowOutput `json:"user"`
}


// Type for action response, including envelope
type ActionExportShowResponse struct {
	Action *ActionExportShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Export *ActionExportShowOutput `json:"export"`
	}

	// Action output without the namespace
	Output *ActionExportShowOutput
}


// Prepare the action for invocation
func (action *ActionExportShow) Prepare() *ActionExportShowInvocation {
	return &ActionExportShowInvocation{
		Action: action,
		Path: "/v6.0/exports/{export_id}",
	}
}

// ActionExportShowInvocation is used to configure action for invocation
type ActionExportShowInvocation struct {
	// Pointer to the action
	Action *ActionExportShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionExportShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionExportShowInvocation) SetPathParamInt(param string, value int64) *ActionExportShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionExportShowInvocation) SetPathParamString(param string, value string) *ActionExportShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionExportShowInvocation) NewMetaInput() *ActionExportShowMetaGlobalInput {
	inv.MetaInput = &ActionExportShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionExportShowInvocation) SetMetaInput(input *ActionExportShowMetaGlobalInput) *ActionExportShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionExportShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionExportShowInvocation) Call() (*ActionExportShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionExportShowInvocation) callAsQuery() (*ActionExportShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionExportShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Export
	}
	return resp, err
}




func (inv *ActionExportShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}

