package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionOomReportUsageShow is a type for action Oom_report.Usage#Show
type ActionOomReportUsageShow struct {
	// Pointer to client
	Client *Client
}

func NewActionOomReportUsageShow(client *Client) *ActionOomReportUsageShow {
	return &ActionOomReportUsageShow{
		Client: client,
	}
}

// ActionOomReportUsageShowMetaGlobalInput is a type for action global meta input parameters
type ActionOomReportUsageShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionOomReportUsageShowMetaGlobalInput) SetIncludes(value string) *ActionOomReportUsageShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionOomReportUsageShowMetaGlobalInput) SetNo(value bool) *ActionOomReportUsageShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionOomReportUsageShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionOomReportUsageShowMetaGlobalInput) SelectParameters(params ...string) *ActionOomReportUsageShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionOomReportUsageShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionOomReportUsageShowOutput is a type for action output parameters
type ActionOomReportUsageShowOutput struct {
	Failcnt int64  "json:\"failcnt\""
	Id      int64  "json:\"id\""
	Limit   int64  "json:\"limit\""
	Memtype string "json:\"memtype\""
	Usage   int64  "json:\"usage\""
}

// Type for action response, including envelope
type ActionOomReportUsageShowResponse struct {
	Action *ActionOomReportUsageShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Usage *ActionOomReportUsageShowOutput "json:\"usage\""
	}

	// Action output without the namespace
	Output *ActionOomReportUsageShowOutput
}

// Prepare the action for invocation
func (action *ActionOomReportUsageShow) Prepare() *ActionOomReportUsageShowInvocation {
	return &ActionOomReportUsageShowInvocation{
		Action: action,
		Path:   "/v7.0/oom_reports/{oom_report_id}/usages/{usage_id}",
	}
}

// ActionOomReportUsageShowInvocation is used to configure action for invocation
type ActionOomReportUsageShowInvocation struct {
	// Pointer to the action
	Action *ActionOomReportUsageShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionOomReportUsageShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionOomReportUsageShowInvocation) SetPathParamInt(param string, value int64) *ActionOomReportUsageShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionOomReportUsageShowInvocation) SetPathParamString(param string, value string) *ActionOomReportUsageShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionOomReportUsageShowInvocation) NewMetaInput() *ActionOomReportUsageShowMetaGlobalInput {
	inv.MetaInput = &ActionOomReportUsageShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionOomReportUsageShowInvocation) SetMetaInput(input *ActionOomReportUsageShowMetaGlobalInput) *ActionOomReportUsageShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionOomReportUsageShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionOomReportUsageShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionOomReportUsageShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionOomReportUsageShowInvocation) Call() (*ActionOomReportUsageShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionOomReportUsageShowInvocation) CallContext(ctx context.Context) (*ActionOomReportUsageShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionOomReportUsageShowResponse{Action: inv.Action}
	raw, err := inv.Action.Client.doRaw(req)
	if err != nil {
		return resp, err
	}
	if err := decodeTypedResponse(raw.Body, resp); err != nil {
		return resp, err
	}
	if resp.Status {
		if resp.Response == nil {
			return resp, errMalformedResponse
		}
		if resp.Response.Usage == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.Usage
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionOomReportUsageShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionOomReportUsageShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionOomReportUsageShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
