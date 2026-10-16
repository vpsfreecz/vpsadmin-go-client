package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// ActionWebuiUserSettingShow is a type for action Webui_user_setting#Show
type ActionWebuiUserSettingShow struct {
	// Pointer to client
	Client *Client
}

func NewActionWebuiUserSettingShow(client *Client) *ActionWebuiUserSettingShow {
	return &ActionWebuiUserSettingShow{
		Client: client,
	}
}

// ActionWebuiUserSettingShowMetaGlobalInput is a type for action global meta input parameters
type ActionWebuiUserSettingShowMetaGlobalInput struct {
	Includes string "json:\"includes\""
	No       bool   "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionWebuiUserSettingShowMetaGlobalInput) SetIncludes(value string) *ActionWebuiUserSettingShowMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionWebuiUserSettingShowMetaGlobalInput) SetNo(value bool) *ActionWebuiUserSettingShowMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionWebuiUserSettingShowMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionWebuiUserSettingShowMetaGlobalInput) SelectParameters(params ...string) *ActionWebuiUserSettingShowMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionWebuiUserSettingShowMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionWebuiUserSettingShowOutput is a type for action output parameters
type ActionWebuiUserSettingShowOutput struct {
	CreatedAt string                "json:\"created_at\""
	Id        int64                 "json:\"id\""
	Key       string                "json:\"key\""
	Namespace string                "json:\"namespace\""
	UpdatedAt string                "json:\"updated_at\""
	User      *ActionUserShowOutput "json:\"user\""
}

// Type for action response, including envelope
type ActionWebuiUserSettingShowResponse struct {
	Action *ActionWebuiUserSettingShow "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		WebuiUserSetting *ActionWebuiUserSettingShowOutput "json:\"webui_user_setting\""
	}

	// Action output without the namespace
	Output *ActionWebuiUserSettingShowOutput
}

// Prepare the action for invocation
func (action *ActionWebuiUserSettingShow) Prepare() *ActionWebuiUserSettingShowInvocation {
	return &ActionWebuiUserSettingShowInvocation{
		Action: action,
		Path:   "/v7.0/webui_user_settings/{namespace}/{key}",
	}
}

// ActionWebuiUserSettingShowInvocation is used to configure action for invocation
type ActionWebuiUserSettingShowInvocation struct {
	// Pointer to the action
	Action *ActionWebuiUserSettingShow

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionWebuiUserSettingShowMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionWebuiUserSettingShowInvocation) SetPathParamInt(param string, value int64) *ActionWebuiUserSettingShowInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionWebuiUserSettingShowInvocation) SetPathParamString(param string, value string) *ActionWebuiUserSettingShowInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionWebuiUserSettingShowInvocation) NewMetaInput() *ActionWebuiUserSettingShowMetaGlobalInput {
	inv.MetaInput = &ActionWebuiUserSettingShowMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionWebuiUserSettingShowInvocation) SetMetaInput(input *ActionWebuiUserSettingShowMetaGlobalInput) *ActionWebuiUserSettingShowInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionWebuiUserSettingShowInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionWebuiUserSettingShowInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionWebuiUserSettingShowInvocation) validate() error {
	verr := NewValidationError(inv.Action.Client)
	if inv.MetaInput != nil {
	}

	if verr.Empty() {
		return nil
	}

	return verr
}

// Call invokes the action using a background context.
func (inv *ActionWebuiUserSettingShowInvocation) Call() (*ActionWebuiUserSettingShowResponse, error) {
	return inv.CallContext(context.Background())
}

// CallContext invokes the action and decodes its typed response.
func (inv *ActionWebuiUserSettingShowInvocation) CallContext(ctx context.Context) (*ActionWebuiUserSettingShowResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp := &ActionWebuiUserSettingShowResponse{Action: inv.Action}
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
		if resp.Response.WebuiUserSetting == nil {
			return resp, errMalformedResponse
		}
		resp.Output = resp.Response.WebuiUserSetting
	}
	return resp, nil
}

// CallRaw validates and sends the action without decoding its response body.
// The caller owns HTTP and API status classification.
func (inv *ActionWebuiUserSettingShowInvocation) CallRaw(ctx context.Context) (*RawResponse, error) {
	req, err := inv.newRequest(ctx)
	if err != nil {
		return nil, err
	}
	return inv.Action.Client.doRaw(req)
}

func (inv *ActionWebuiUserSettingShowInvocation) newRequest(ctx context.Context) (*http.Request, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	return inv.Action.Client.newQueryStringRequest(ctx, inv.Path, queryParams)
}

func (inv *ActionWebuiUserSettingShowInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}
}
