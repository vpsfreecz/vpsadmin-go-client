package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionMetricsAccessTokenShow is a type for action Metrics_access_token#Show
type ActionMetricsAccessTokenShow struct {
	// Pointer to client
	Client *Client
}

func NewActionMetricsAccessTokenShow(client *Client) *ActionMetricsAccessTokenShow {
	return &ActionMetricsAccessTokenShow{
		Client: client,
	}
}

// ActionMetricsAccessTokenShowMetaGlobalInput is a type for action global meta input parameters
type ActionMetricsAccessTokenShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionMetricsAccessTokenShowMetaGlobalInput) SetIncludes(value string) *ActionMetricsAccessTokenShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionMetricsAccessTokenShowMetaGlobalInput) SetNo(value bool) *ActionMetricsAccessTokenShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionMetricsAccessTokenShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionMetricsAccessTokenShowMetaGlobalInput) SelectParameters(params ...string) *ActionMetricsAccessTokenShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionMetricsAccessTokenShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionMetricsAccessTokenShowOutput is a type for action output parameters
type ActionMetricsAccessTokenShowOutput struct {
	AccessToken  string                "json:\"access_token\""
	CreatedAt    string                "json:\"created_at\""
	Id           int64                 "json:\"id\""
	LastUse      string                "json:\"last_use\""
	MetricPrefix string                "json:\"metric_prefix\""
	UpdatedAt    string                "json:\"updated_at\""
	UseCount     int64                 "json:\"use_count\""
	User         *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionMetricsAccessTokenShowResponse struct {
	Action *ActionMetricsAccessTokenShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		MetricsAccessToken *ActionMetricsAccessTokenShowOutput "json:\"metrics_access_token\""
	}

	// Action output without the namespace
	Output *ActionMetricsAccessTokenShowOutput
}

// Prepare the action for invocation
func (action *ActionMetricsAccessTokenShow) Prepare() *ActionMetricsAccessTokenShowInvocation {
	return &ActionMetricsAccessTokenShowInvocation{
		Action: action,
		Path:   "/v7.0/metrics_access_tokens/{metrics_access_token_id}",
	}
}

// ActionMetricsAccessTokenShowInvocation is used to configure action for invocation
type ActionMetricsAccessTokenShowInvocation struct {
	// Pointer to the action
	Action *ActionMetricsAccessTokenShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionMetricsAccessTokenShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionMetricsAccessTokenShowInvocation) SetPathParamInt(param string, value int64) *ActionMetricsAccessTokenShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionMetricsAccessTokenShowInvocation) SetPathParamString(param string, value string) *ActionMetricsAccessTokenShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionMetricsAccessTokenShowInvocation) NewMetaInput() *ActionMetricsAccessTokenShowMetaGlobalInput {
	inv.MetaInput = &ActionMetricsAccessTokenShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionMetricsAccessTokenShowInvocation) SetMetaInput(input *ActionMetricsAccessTokenShowMetaGlobalInput) *ActionMetricsAccessTokenShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionMetricsAccessTokenShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionMetricsAccessTokenShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionMetricsAccessTokenShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionMetricsAccessTokenShowInvocation) Call() (*ActionMetricsAccessTokenShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionMetricsAccessTokenShowInvocation) CallContext(ctx context.Context) (*ActionMetricsAccessTokenShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionMetricsAccessTokenShowResponse{Action: inv.Action}
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
		if resp.Response.MetricsAccessToken == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.MetricsAccessToken
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionMetricsAccessTokenShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionMetricsAccessTokenShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionMetricsAccessTokenShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
