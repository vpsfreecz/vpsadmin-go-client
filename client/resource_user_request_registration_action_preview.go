package client

import (
	"strings"
)

// ActionUserRequestRegistrationPreview is a type for action User_request.Registration#Preview
type ActionUserRequestRegistrationPreview struct {
	// Pointer to client
	Client *Client
}

func NewActionUserRequestRegistrationPreview(client *Client) *ActionUserRequestRegistrationPreview {
	return &ActionUserRequestRegistrationPreview{
		Client: client,
	}
}

// ActionUserRequestRegistrationPreviewMetaGlobalInput is a type for action global meta input parameters
type ActionUserRequestRegistrationPreviewMetaGlobalInput struct {
	No bool `json:"no"`
	Includes string `json:"includes"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserRequestRegistrationPreviewMetaGlobalInput) SetNo(value bool) *ActionUserRequestRegistrationPreviewMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}
// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionUserRequestRegistrationPreviewMetaGlobalInput) SetIncludes(value string) *ActionUserRequestRegistrationPreviewMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserRequestRegistrationPreviewMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserRequestRegistrationPreviewMetaGlobalInput) SelectParameters(params ...string) *ActionUserRequestRegistrationPreviewMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserRequestRegistrationPreviewMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}



// ActionUserRequestRegistrationPreviewOutput is a type for action output parameters
type ActionUserRequestRegistrationPreviewOutput struct {
	Id int64 `json:"id"`
	AdminResponse string `json:"admin_response"`
	Login string `json:"login"`
	FullName string `json:"full_name"`
	OrgName string `json:"org_name"`
	OrgId string `json:"org_id"`
	Email string `json:"email"`
	Address string `json:"address"`
	YearOfBirth int64 `json:"year_of_birth"`
	How string `json:"how"`
	Note string `json:"note"`
	OsTemplate *ActionOsTemplateShowOutput `json:"os_template"`
	Location *ActionLocationShowOutput `json:"location"`
	Currency string `json:"currency"`
	Language *ActionLanguageShowOutput `json:"language"`
}


// Type for action response, including envelope
type ActionUserRequestRegistrationPreviewResponse struct {
	Action *ActionUserRequestRegistrationPreview `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		Registration *ActionUserRequestRegistrationPreviewOutput `json:"registration"`
	}

	// Action output without the namespace
	Output *ActionUserRequestRegistrationPreviewOutput
}


// Prepare the action for invocation
func (action *ActionUserRequestRegistrationPreview) Prepare() *ActionUserRequestRegistrationPreviewInvocation {
	return &ActionUserRequestRegistrationPreviewInvocation{
		Action: action,
		Path: "/v5.0/user_request/registrations/{registration_id}/{token}",
	}
}

// ActionUserRequestRegistrationPreviewInvocation is used to configure action for invocation
type ActionUserRequestRegistrationPreviewInvocation struct {
	// Pointer to the action
	Action *ActionUserRequestRegistrationPreview

	// Path which may contain parameters that need to be set
	Path string
	// Global meta input parameters
	MetaInput *ActionUserRequestRegistrationPreviewMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserRequestRegistrationPreviewInvocation) SetPathParamInt(param string, value int64) *ActionUserRequestRegistrationPreviewInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserRequestRegistrationPreviewInvocation) SetPathParamString(param string, value string) *ActionUserRequestRegistrationPreviewInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserRequestRegistrationPreviewInvocation) NewMetaInput() *ActionUserRequestRegistrationPreviewMetaGlobalInput {
	inv.MetaInput = &ActionUserRequestRegistrationPreviewMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserRequestRegistrationPreviewInvocation) SetMetaInput(input *ActionUserRequestRegistrationPreviewMetaGlobalInput) *ActionUserRequestRegistrationPreviewInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserRequestRegistrationPreviewInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionUserRequestRegistrationPreviewInvocation) Call() (*ActionUserRequestRegistrationPreviewResponse, error) {
	return inv.callAsQuery()
}

func (inv *ActionUserRequestRegistrationPreviewInvocation) callAsQuery() (*ActionUserRequestRegistrationPreviewResponse, error) {
	queryParams := make(map[string]string)
	inv.convertMetaInputToQueryParams(queryParams)
	resp := &ActionUserRequestRegistrationPreviewResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.Registration
	}
	return resp, err
}




func (inv *ActionUserRequestRegistrationPreviewInvocation) convertMetaInputToQueryParams(ret map[string]string) {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
		if inv.IsMetaParameterSelected("Includes") {
			ret["_meta[includes]"] = inv.MetaInput.Includes
		}
	}
}

