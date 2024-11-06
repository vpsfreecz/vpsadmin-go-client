package client

import ()

// ActionDnsResolverCreate is a type for action Dns_resolver#Create
type ActionDnsResolverCreate struct {
	// Pointer to client
	Client *Client
}

func NewActionDnsResolverCreate(client *Client) *ActionDnsResolverCreate {
	return &ActionDnsResolverCreate{
		Client: client,
	}
}

// ActionDnsResolverCreateMetaGlobalInput is a type for action global meta input parameters
type ActionDnsResolverCreateMetaGlobalInput struct {
	Includes string `json:"includes"`
	No       bool   `json:"no"`
	// Only selected parameters are sent to the API. Ignored if empty.
	_selectedParameters map[string]interface{}
	// Parameters that are set to nil instead of value
	_nilParameters map[string]interface{}
}

// SetIncludes sets parameter Includes to value and selects it for sending
func (in *ActionDnsResolverCreateMetaGlobalInput) SetIncludes(value string) *ActionDnsResolverCreateMetaGlobalInput {
	in.Includes = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Includes"] = nil
	return in
}

// SetNo sets parameter No to value and selects it for sending
func (in *ActionDnsResolverCreateMetaGlobalInput) SetNo(value bool) *ActionDnsResolverCreateMetaGlobalInput {
	in.No = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["No"] = nil
	return in
}

// SelectParameters sets parameters from ActionDnsResolverCreateMetaGlobalInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsResolverCreateMetaGlobalInput) SelectParameters(params ...string) *ActionDnsResolverCreateMetaGlobalInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

func (in *ActionDnsResolverCreateMetaGlobalInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsResolverCreateInput is a type for action input parameters
type ActionDnsResolverCreateInput struct {
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
func (in *ActionDnsResolverCreateInput) SetIpAddr(value string) *ActionDnsResolverCreateInput {
	in.IpAddr = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IpAddr"] = nil
	return in
}

// SetIsUniversal sets parameter IsUniversal to value and selects it for sending
func (in *ActionDnsResolverCreateInput) SetIsUniversal(value bool) *ActionDnsResolverCreateInput {
	in.IsUniversal = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["IsUniversal"] = nil
	return in
}

// SetLabel sets parameter Label to value and selects it for sending
func (in *ActionDnsResolverCreateInput) SetLabel(value string) *ActionDnsResolverCreateInput {
	in.Label = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in._selectedParameters["Label"] = nil
	return in
}

// SetLocation sets parameter Location to value and selects it for sending
func (in *ActionDnsResolverCreateInput) SetLocation(value int64) *ActionDnsResolverCreateInput {
	in.Location = value

	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	in.SetLocationNil(false)
	in._selectedParameters["Location"] = nil
	return in
}

// SetLocationNil sets parameter Location to nil and selects it for sending
func (in *ActionDnsResolverCreateInput) SetLocationNil(set bool) *ActionDnsResolverCreateInput {
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

// SelectParameters sets parameters from ActionDnsResolverCreateInput
// that will be sent to the API.
// SelectParameters can be called multiple times.
func (in *ActionDnsResolverCreateInput) SelectParameters(params ...string) *ActionDnsResolverCreateInput {
	if in._selectedParameters == nil {
		in._selectedParameters = make(map[string]interface{})
	}

	for _, param := range params {
		in._selectedParameters[param] = nil
	}

	return in
}

// UnselectParameters unsets parameters from ActionDnsResolverCreateInput
// that will be sent to the API.
// UnsSelectParameters can be called multiple times.
func (in *ActionDnsResolverCreateInput) UnselectParameters(params ...string) *ActionDnsResolverCreateInput {
	if in._selectedParameters == nil {
		return in
	}

	for _, param := range params {
		delete(in._selectedParameters, param)
	}

	return in
}

func (in *ActionDnsResolverCreateInput) AnySelected() bool {
	if in._selectedParameters == nil {
		return false
	}

	return len(in._selectedParameters) > 0
}

// ActionDnsResolverCreateRequest is a type for the entire action request
type ActionDnsResolverCreateRequest struct {
	DnsResolver map[string]interface{} `json:"dns_resolver"`
	Meta        map[string]interface{} `json:"_meta"`
}

// ActionDnsResolverCreateOutput is a type for action output parameters
type ActionDnsResolverCreateOutput struct {
	Id          int64                     `json:"id"`
	IpAddr      string                    `json:"ip_addr"`
	IsUniversal bool                      `json:"is_universal"`
	Label       string                    `json:"label"`
	Location    *ActionLocationShowOutput `json:"location"`
}

// Type for action response, including envelope
type ActionDnsResolverCreateResponse struct {
	Action *ActionDnsResolverCreate `json:"-"`
	*Envelope
	// Action output encapsulated within a namespace
	Response *struct {
		DnsResolver *ActionDnsResolverCreateOutput `json:"dns_resolver"`
	}

	// Action output without the namespace
	Output *ActionDnsResolverCreateOutput
}

// Prepare the action for invocation
func (action *ActionDnsResolverCreate) Prepare() *ActionDnsResolverCreateInvocation {
	return &ActionDnsResolverCreateInvocation{
		Action: action,
		Path:   "/v7.0/dns_resolvers",
	}
}

// ActionDnsResolverCreateInvocation is used to configure action for invocation
type ActionDnsResolverCreateInvocation struct {
	// Pointer to the action
	Action *ActionDnsResolverCreate

	// Path which may contain parameters that need to be set
	Path string
	// Input parameters
	Input *ActionDnsResolverCreateInput
	// Global meta input parameters
	MetaInput *ActionDnsResolverCreateMetaGlobalInput
}

// NewInput returns a new struct for input parameters and sets it as with SetInput
func (inv *ActionDnsResolverCreateInvocation) NewInput() *ActionDnsResolverCreateInput {
	inv.Input = &ActionDnsResolverCreateInput{}
	return inv.Input
}

// SetInput provides input parameters to send to the API
func (inv *ActionDnsResolverCreateInvocation) SetInput(input *ActionDnsResolverCreateInput) *ActionDnsResolverCreateInvocation {
	inv.Input = input
	return inv
}

// IsParameterSelected returns true if param is to be sent to the API
func (inv *ActionDnsResolverCreateInvocation) IsParameterSelected(param string) bool {
	if inv.Input._selectedParameters == nil {
		return true
	}

	_, exists := inv.Input._selectedParameters[param]
	return exists
}

// IsParameterNil returns true if param is to be sent to the API as nil
func (inv *ActionDnsResolverCreateInvocation) IsParameterNil(param string) bool {
	if inv.Input._nilParameters == nil {
		return false
	}

	_, exists := inv.Input._nilParameters[param]
	return exists
}

// NewMetaInput returns a new struct for global meta input parameters and sets
// it as with SetMetaInput
func (inv *ActionDnsResolverCreateInvocation) NewMetaInput() *ActionDnsResolverCreateMetaGlobalInput {
	inv.MetaInput = &ActionDnsResolverCreateMetaGlobalInput{}
	return inv.MetaInput
}

// SetMetaInput provides global meta input parameters to send to the API
func (inv *ActionDnsResolverCreateInvocation) SetMetaInput(input *ActionDnsResolverCreateMetaGlobalInput) *ActionDnsResolverCreateInvocation {
	inv.MetaInput = input
	return inv
}

// IsMetaParameterSelected returns true if global meta param is to be sent to the API
func (inv *ActionDnsResolverCreateInvocation) IsMetaParameterSelected(param string) bool {
	if inv.MetaInput._selectedParameters == nil {
		return true
	}

	_, exists := inv.MetaInput._selectedParameters[param]
	return exists
}

// IsMetaParameterNil returns true if global meta param is to be sent to the API as nil
func (inv *ActionDnsResolverCreateInvocation) IsMetaParameterNil(param string) bool {
	if inv.MetaInput._nilParameters == nil {
		return false
	}

	_, exists := inv.MetaInput._nilParameters[param]
	return exists
}

// Call() invokes the action and returns a response from the API server
func (inv *ActionDnsResolverCreateInvocation) Call() (*ActionDnsResolverCreateResponse, error) {
	return inv.callAsBody()
}

func (inv *ActionDnsResolverCreateInvocation) callAsBody() (*ActionDnsResolverCreateResponse, error) {
	input := inv.makeAllInputParams()
	resp := &ActionDnsResolverCreateResponse{Action: inv.Action}
	err := inv.Action.Client.DoBodyRequest("POST", inv.Path, input, resp)
	if err == nil && resp.Status {
		resp.Output = resp.Response.DnsResolver
	}
	return resp, err
}

func (inv *ActionDnsResolverCreateInvocation) makeAllInputParams() *ActionDnsResolverCreateRequest {
	return &ActionDnsResolverCreateRequest{
		DnsResolver: inv.makeInputParams(),
		Meta:        inv.makeMetaInputParams(),
	}
}

func (inv *ActionDnsResolverCreateInvocation) makeInputParams() map[string]interface{} {
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

func (inv *ActionDnsResolverCreateInvocation) makeMetaInputParams() map[string]interface{} {
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
