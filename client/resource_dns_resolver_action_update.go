package client

import (
	"strings"
)

// ActionDnsResolverUpdate is a type for action Dns_resolver#Update
type ActionDnsResolverUpdate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsResolverUpdate(client *Client) *ActionDnsResolverUpdate {
	return &ActionDnsResolverUpdate{
		Client: client,
	}
}

// ActionDnsResolverUpdateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsResolverUpdateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsResolverUpdateMetaGlobalInput) SetIncludes(value string) *ActionDnsResolverUpdateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsResolverUpdateMetaGlobalInput) SetNo(value bool) *ActionDnsResolverUpdateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsResolverUpdateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsResolverUpdateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsResolverUpdateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsResolverUpdateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsResolverUpdateInput is a type for action input parameters
type ActionDnsResolverUpdateInput struct {
	IpAddr      string `json:"ip_addr"`
	IsUniversal bool   `json:"is_universal"`
	Label       string `json:"label"`
	Location    int64  `json:"location"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIpAddr sets parameter IpAddr to value and selects it for sending
func (in *ActionDnsResolverUpdateInput) SetIpAddr(value string) *ActionDnsResolverUpdateInput {
	in.IpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpAddr"] = nil
	return in
}

// SetIsUniversal sets parameter IsUniversal to value and selects it for sending
func (in *ActionDnsResolverUpdateInput) SetIsUniversal(value bool) *ActionDnsResolverUpdateInput {
	in.IsUniversal = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IsUniversal"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionDnsResolverUpdateInput) SetLabel(value string) *ActionDnsResolverUpdateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionDnsResolverUpdateInput) SetLocation(value int64) *ActionDnsResolverUpdateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionDnsResolverUpdateInput) SetLocationNil(set bool) *ActionDnsResolverUpdateInput {
	if in._nilParameters == nil {
		if !set {
			return in
		}
		in._nilParameters = make(map[string]interface{})
	}

	if set {
		in._nilParameters["Location"] = nil
		in.SelectParameters("Location")
	} else {
		delete(in._nilParameters, "Location")
	}
	return in
}

// SelectParameters sets parameters from ActionDnsResolverUpdateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsResolverUpdateInput) SelectParameters(params ...string) *ActionDnsResolverUpdateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsResolverUpdateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsResolverUpdateInput) UnselectParameters(params ...string) *ActionDnsResolverUpdateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsResolverUpdateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsResolverUpdateRequest is a type for the entire action request
type ActionDnsResolverUpdateRequest struct {
	DnsResolver map[string]interface{} `json:"dns_resolver"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionDnsResolverUpdateOutput is a type for action output parameters
type ActionDnsResolverUpdateOutput struct {
	Id          int64                     `json:"id"`
	IpAddr      string                    `json:"ip_addr"`
	IsUniversal bool                      `json:"is_universal"`
	Label       string                    `json:"label"`
	Location    *ActionLocationShowOutput `json:"location"`
}

// Type for action response, including envelope
type ActionDnsResolverUpdateResponse struct {
	Action *ActionDnsResolverUpdate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsResolver *ActionDnsResolverUpdateOutput `json:"dns_resolver"`
	}

	// Action output without the namespace
	Output *ActionDnsResolverUpdateOutput
}

// Prepare the action for invocation
func (action *ActionDnsResolverUpdate) Prepare() *ActionDnsResolverUpdateInvocation {
	return &ActionDnsResolverUpdateInvocation{
		Action: action,
		Path:   "/v6.0/dns_resolvers/{dns_resolver_id}",
	}
}

// ActionDnsResolverUpdateInvocation is used to configure action for invocation
type ActionDnsResolverUpdateInvocation struct {
	// Pointer to the action
	Action *ActionDnsResolverUpdate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsResolverUpdateInput
	// Global meta input parameters
	MetaInput *ActionDnsResolverUpdateMetaGlobalInput
}

// SetPathParamInt sets integer path parameter
func (inv *ActionDnsResolverUpdateInvocation) SetPathParamInt(param string, value int64) *ActionDnsResolverUpdateInvocation {
	return inv.SetPathParamString(param, convertInt64ToString(value))
}

// SetPathParamString sets string path parameter
func (inv *ActionDnsResolverUpdateInvocation) SetPathParamString(param string, value string) *ActionDnsResolverUpdateInvocation {
	inv.Path = strings.Replace(inv.Path, "{"+param+"}", value, 1)
	return inv
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsResolverUpdateInvocation) NewInput() *ActionDnsResolverUpdateInput {
	inv.Input = &ActionDnsResolverUpdateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsResolverUpdateInvocation) SetInput(input *ActionDnsResolverUpdateInput) *ActionDnsResolverUpdateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsResolverUpdateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsResolverUpdateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsResolverUpdateInvocation) NewMetaInput() *ActionDnsResolverUpdateMetaGlobalInput {
	inv.MetaInput = &ActionDnsResolverUpdateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsResolverUpdateInvocation) SetMetaInput(input *ActionDnsResolverUpdateMetaGlobalInput) *ActionDnsResolverUpdateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsResolverUpdateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsResolverUpdateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsResolverUpdateInvocation) Call() (*ActionDnsResolverUpdateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsResolverUpdateInvocation) callAsBody() (*ActionDnsResolverUpdateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsResolverUpdateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("PUT", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsResolver
	}
	return resp, err
}

func (inv *ActionDnsResolverUpdateInvocation) makeAllInputParams() *ActionDnsResolverUpdateRequest {
	return &ActionDnsResolverUpdateRequest{
		DnsResolver: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsResolverUpdateInvocation) makeInputParams() map[string]interface{} {
	ret := make(map[string]interface{})

	if inv.Input != nil {
		if inv.IsParameterSelected("IpAddr") {
			ret["ip_addr"] = inv.Input.IpAddr
		}
		if inv.IsParameterSelected("IsUniversal") {
			ret["is_universal"] = inv.Input.IsUniversal
		}
		if inv.IsParameterSelected("Label") {
			ret["label"] = inv.Input.Label
		}
		if inv.IsParameterSelected("Location") {
			if inv.IsParameterNil("Location") {
				ret["location"] = nil
			} else {
				ret["location"] = inv.Input.Location
			}
		}
	}

	return ret
}

func (inv *ActionDnsResolverUpdateInvocation) makeMetaInputParams() map[string]interface{} {
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
