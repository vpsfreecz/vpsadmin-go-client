package client

import (
	"net/url"
	"strings"
)

// ActionUserAvailableIps is a type for action User#Available_ips
type ActionUserAvailableIps struct {
	// Pointer to client
	Client *Client
}

func NewActionUserAvailableIps(client *Client) *ActionUserAvailableIps {
	return &ActionUserAvailableIps{
		Client: client,
	}
}

// ActionUserAvailableIpsMetaGlobalInput is a type for action global meta input parameters
type ActionUserAvailableIpsMetaGlobalInput struct {
	No bool "json:\"no\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionUserAvailableIpsMetaGlobalInput) SetNo(value bool) *ActionUserAvailableIpsMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserAvailableIpsMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserAvailableIpsMetaGlobalInput) SelectParameters(params ...string) *ActionUserAvailableIpsMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionUserAvailableIpsMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserAvailableIpsInput is a type for action input parameters
type ActionUserAvailableIpsInput struct {
	AddressLocation int64 "json:\"address_location\""
	Location        int64 "json:\"location\""
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetAddressLocation sets parameter AddressLocation to value and selects it for sending
func (in *ActionUserAvailableIpsInput) SetAddressLocation(value int64) *ActionUserAvailableIpsInput {
	in.AddressLocation = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["AddressLocation"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionUserAvailableIpsInput) SetLocation(value int64) *ActionUserAvailableIpsInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Location"] = nil
	return in
}

// SelectParameters sets parameters from ActionUserAvailableIpsInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionUserAvailableIpsInput) SelectParameters(params ...string) *ActionUserAvailableIpsInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionUserAvailableIpsInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionUserAvailableIpsInput) UnselectParameters(params ...string) *ActionUserAvailableIpsInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionUserAvailableIpsInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionUserAvailableIpsOutput is a type for action output parameters
type ActionUserAvailableIpsOutput struct {
	Ipv4        int64 "json:\"ipv4\""
	Ipv4Private int64 "json:\"ipv4_private\""
	Ipv6        int64 "json:\"ipv6\""
}

// Type for action response, including envelope
type ActionUserAvailableIpsResponse struct {
	Action *ActionUserAvailableIps "json:\"-\""
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		User *ActionUserAvailableIpsOutput "json:\"user\""
	}

	// Action output without the namespace
	Output *ActionUserAvailableIpsOutput
}

// Prepare the action for invocation
func (action *ActionUserAvailableIps) Prepare() *ActionUserAvailableIpsInvocation {
	return &ActionUserAvailableIpsInvocation{
		Action: action,
		Path:   "/v7.0/users/{user_id}/available_ips",
	}
}

// ActionUserAvailableIpsInvocation is used to configure action for invocation
type ActionUserAvailableIpsInvocation struct {
	// Pointer to the action
	Action *ActionUserAvailableIps

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionUserAvailableIpsInput
	// Global meta input parameters
	MetaInput *ActionUserAvailableIpsMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionUserAvailableIpsInvocation) SetPathParamInt(param string, value int64) *ActionUserAvailableIpsInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionUserAvailableIpsInvocation) SetPathParamString(param string, value string) *ActionUserAvailableIpsInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", url.PathEscape(value), 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionUserAvailableIpsInvocation) NewInput() *ActionUserAvailableIpsInput {
	inv.Input = &ActionUserAvailableIpsInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionUserAvailableIpsInvocation) SetInput(input *ActionUserAvailableIpsInput) *ActionUserAvailableIpsInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionUserAvailableIpsInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionUserAvailableIpsInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionUserAvailableIpsInvocation) NewMetaInput() *ActionUserAvailableIpsMetaGlobalInput {
	inv.MetaInput = &ActionUserAvailableIpsMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionUserAvailableIpsInvocation) SetMetaInput(input *ActionUserAvailableIpsMetaGlobalInput) *ActionUserAvailableIpsInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionUserAvailableIpsInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionUserAvailableIpsInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

func (inv *ActionUserAvailableIpsInvocation) validate() error {
	verr := NewValidationError()
	if inv.Input != nil {
		if inv.IsParameterSelected("AddressLocation") {
			if !inv.IsParameterNil("AddressLocation") {
				if inv.Input.AddressLocation < 0 {
					verr.Add("address_location", "not a valid resource id")
				}
			}
		}
		if inv.IsParameterSelected("Location") {
			if !inv.IsParameterNil("Location") {
				if inv.Input.Location < 0 {
					verr.Add("location", "not a valid resource id")
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
func (inv *ActionUserAvailableIpsInvocation) Call() (*ActionUserAvailableIpsResponse, error) {
	if err := inv.validate(); err != nil {
		return nil, err
	}
	return inv.callAsQuery()
}

func (inv *ActionUserAvailableIpsInvocation) callAsQuery() (*ActionUserAvailableIpsResponse, error) {
	queryParams := make(map[string]string)
	if err := inv.convertInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	if err := inv.convertMetaInputToQueryParams(queryParams); err != nil {
		return nil, err
	}
	resp := &ActionUserAvailableIpsResponse{Action: inv.Action}
	err := inv.Action.Client.DoQueryStringRequest(inv.Path, queryParams, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.User
	}
	return resp, err
}

func (inv *ActionUserAvailableIpsInvocation) convertInputToQueryParams(ret map[string]string) error {
	if inv.Input != nil {
		if inv.IsParameterSelected("AddressLocation") {
			ret["user[address_location]"] = convertInt64ToString(inv.Input.AddressLocation)
		}
		if inv.IsParameterSelected("Location") {
			ret["user[location]"] = convertInt64ToString(inv.Input.Location)
		}
	}

	return nil
}

func (inv *ActionUserAvailableIpsInvocation) convertMetaInputToQueryParams(ret map[string]string) error {
	if inv.MetaInput != nil {
		if inv.IsMetaParameterSelected("No") {
			ret["_meta[no]"] = convertBoolToString(inv.MetaInput.No)
		}
	}

	return nil
}
