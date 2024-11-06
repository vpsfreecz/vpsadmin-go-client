package client

import (
	"strings"
)

// ActionDatasetExpansionShow is a type for action Dataset_expansion#Show
type ActionDatasetExpansionShow struct {
	// Pointer to client
	Client *Client
}

func NewActionDatasetExpansionShow(client *Client) *ActionDatasetExpansionShow {
	return &ActionDatasetExpansionShow{
		Client: client,
	}
}

// ActionDatasetExpansionShowMetaGlobalInput is a type for action global meta input parameters
type ActionDatasetExpansionShowMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDatasetExpansionShowMetaGlobalInput) SetIncludes(value string) *ActionDatasetExpansionShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDatasetExpansionShowMetaGlobalInput) SetNo(value bool) *ActionDatasetExpansionShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDatasetExpansionShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDatasetExpansionShowMetaGlobalInput) SelectParameters(params ...string) *ActionDatasetExpansionShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDatasetExpansionShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDatasetExpansionShowOutput is a type for action output parameters
type ActionDatasetExpansionShowOutput struct {
	AddedSpace             int64                    `json:"added_space"`
	CreatedAt              string                   `json:"created_at"`
	Dataset                *ActionDatasetShowOutput `json:"dataset"`
	EnableNotifications    bool                     `json:"enable_notifications"`
	EnableShrink           bool                     `json:"enable_shrink"`
	Id                     int64                    `json:"id"`
	MaxOverRefquotaSeconds int64                    `json:"max_over_refquota_seconds"`
	OriginalRefquota       int64                    `json:"original_refquota"`
	OverRefquotaSeconds    int64                    `json:"over_refquota_seconds"`
	State                  string                   `json:"state"`
	StopVps                bool                     `json:"stop_vps"`
	Vps                    *ActionVpsShowOutput     `json:"vps"`
}

// Type for action response, including envelope
type ActionDatasetExpansionShowResponse struct {
	Action *ActionDatasetExpansionShow `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DatasetExpansion *ActionDatasetExpansionShowOutput `json:"dataset_expansion"`
	}

	// Action output without the namespace
	Output *ActionDatasetExpansionShowOutput
}

// Prepare the action for invocation
func (action *ActionDatasetExpansionShow) Prepare() *ActionDatasetExpansionShowInvocation {
	return &ActionDatasetExpansionShowInvocation{
		Action: action,
		Path:   "/v7.0/dataset_expansions/{dataset_expansion_id}",
	}
}

// ActionDatasetExpansionShowInvocation is used to configure action for invocation
type ActionDatasetExpansionShowInvocation struct {
	// Pointer to the action
	Action *ActionDatasetExpansionShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionDatasetExpansionShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDatasetExpansionShowInvocation) SetPathParamInt(param string, value int64) *ActionDatasetExpansionShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDatasetExpansionShowInvocation) SetPathParamString(param string, value string) *ActionDatasetExpansionShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDatasetExpansionShowInvocation) NewMetaInput() *ActionDatasetExpansionShowMetaGlobalInput {
	inv.MetaInput = &ActionDatasetExpansionShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDatasetExpansionShowInvocation) SetMetaInput(input *ActionDatasetExpansionShowMetaGlobalInput) *ActionDatasetExpansionShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDatasetExpansionShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDatasetExpansionShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDatasetExpansionShowInvocation) Call() (*ActionDatasetExpansionShowResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionDatasetExpansionShowInvocation) callAsQuery() (*ActionDatasetExpansionShowResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionDatasetExpansionShowResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DatasetExpansion
	}
	return resp, err
}

func (inv *ActionDatasetExpansionShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
