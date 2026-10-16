vpsAdmin Client in Go
---------------------
The client is generated using `haveapi-go-client` from
[HaveAPI](https://github.com/vpsfreecz/haveapi).

## Usage
```go
package main

import (
	"fmt"
	"github.com/vpsfreecz/vpsadmin-go-client/client"
)

func main() {
	api := client.New("https://api.vpsfree.cz")

	response, err := api.Cluster.PublicStats.Call()

	if err != nil {
		fmt.Println(err)
		return
	} else if !response.Status {
		fmt.Printf("Action failed: %s\n", response.Message)
	}

	// Response envelope
	fmt.Printf("%+v\n", response)

	// Actual action output, namespaced
	fmt.Printf("%+v\n", response.Response)
	fmt.Printf("%+v\n", response.Response.Cluster)

	// Action output, directly accessed
	fmt.Printf("%+v\n", response.Output)
}
```

### Authentication
Both HTTP basic and token authentication is supported.

```go
// HTTP basic
api.SetBasicAuthentication("user", "password")

// New token
err := api.SetNewTokenAuth(&client.TokenAuthOptions{
	User: "user",
	Password: "secret",
})

if err == nil {
	fmt.Printf("Token is: %s\n", api.Authentication.(*client.TokenAuth).Token)
}

// Existing token
api.SetExistingTokenAuth("mytoken")
```

### HTTP timeouts
The client uses `http.DefaultClient` by default. Configure an overall request
timeout with `SetTimeout`, or provide a custom HTTP client with
`SetHTTPClient`.

```go
api.SetTimeout(60 * time.Second)
api.SetHTTPClient(&http.Client{Timeout: 60 * time.Second})
```

### Context and raw responses

Action invocations provide `CallContext(ctx)` alongside the existing `Call()`.
The context covers waiting for response headers and reading the response body.
For example, after importing `context` and `time`:

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
defer cancel()

action := api.Vps.Show.Prepare()
action.SetPathParamInt("vps_id", 123)
response, err := action.CallContext(ctx)
```

Use `CallRaw(ctx)` instead when the application needs the complete response
bytes or its own HTTP/API error handling. It returns `*client.RawResponse` with
`StatusCode int` and `Body json.RawMessage`. Check the HTTP status and decode the
body for the fields the application consumes. Raw calls preserve malformed JSON,
nulls and number representations; they do not classify the response or retry.
Validation, transport, cancellation and body-read failures remain Go errors.

Each call dispatches once through the configured HTTP client. The generator
adds no retry loop; redirects or transport behavior may cause further network
requests. Configure the HTTP client when strict redirect or replay control is
required. Direct actions that already have a parameter-free `Call()` also
provide the context and raw shortcuts. Existing `Call()` and
low-level helpers use a background context; higher-level authentication and
blocking-operation helpers retain their existing interfaces.

See the generator's [context and raw response contract](https://github.com/vpsfreecz/haveapi/blob/159631d3ce7f6804222754a3c27510d360c7bd18/clients/go/README.md#context-and-raw-responses)
for typed-response handling and body ownership.

### Custom JSON body inputs

Ordinary `Custom` body fields use `json.RawMessage`. For example, after importing
`encoding/json`, a generated value setter accepts an object directly:

```go
input.SetValue(json.RawMessage(`{"enabled":true}`))
```

This applies to IP-release address lists, WebUI settings, system configuration
values and WebAuthn finish credentials. The JSON encoder embeds these values
without converting numbers through `float64`; it may compact whitespace and
escape strings. Keep the supplied slice stable until the call returns.

Input selection remains unchanged. Setters select their fields; unselected
fields are omitted. An all-fields input includes every field, including null for
an unset raw value. A selected nil value or JSON `null` sends null. Selected
empty, malformed or trailing JSON fails before authentication and dispatch;
unselected malformed values are ignored. Required fields, nullability and
application validation remain the server's responsibility.

Metadata `includes` retains its string setter. Ordinary Custom GET inputs are
unsupported and fail generation. Custom output fields, including settings values
and WebAuthn options, remain absent from typed responses; use `CallRaw` to read
them. See the generator's [Custom input contract](https://github.com/vpsfreecz/haveapi/blob/159631d3ce7f6804222754a3c27510d360c7bd18/clients/go/README.md#custom-json-inputs)
for nullable helpers and selection details.

### Two-factor authentication
Two-factor authentication is supported only by token authentication. The API
server decides when the two-factor authentication is required, so the client
has to implement appropriate callbacks.

```go
err := api.SetNewTokenAuth(&client.TokenAuthOptions{
	User: "user",
	Password: "secret",
	TotpCallback: func (input *client.AuthTokenActionTokenTotpInput) error {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("TOTP code: ")
		code, _ := reader.ReadString('\n')
		input.SetCode(strings.TrimSpace(code))
		return nil
	},
})
```

### Path parameters
Path parameters are used to identify resources, e.g. by their ID.

```go
action := api.Vps.Show.Prepare()
action.SetPathParamInt("vps_id", 123)
response, err := action.Call()
```

### Input parameters
The client can be configured either to send just the configured parameters,
or to send all parameters, even unset ones.

```go
action := api.Vps.Update.Prepare()
action.SetPathParamInt("vps_id", 123)

input := action.NewInput()
input.SetHostname("new-hostname")

// Although Vps.Update has many more input parameters, only Hostname was set
// and so no other parameter will be sent. This is what we want in order to
// change just the hostname.
response, err := action.Call()
```

To send all input parameters, do not use the setter functions, but access
the input struct's fields directly. Or create the struct manually
(e.g. `ActionVpsUpdateInput{}`) and then use `action.SetInput()`.

## Updating the client

Files in `client/` are generated. Make shared behavior changes in the HaveAPI
Go generator and regenerate the client rather than editing generated files.

The current snapshot uses the public vpsAdmin v7.0 description (SHA256
`0eff1941670f0a7b344e6c9c8fb67c936c2bcf8f3ce19e1cfcfd1ce880d93fb8`)
and HaveAPI generator revision
`159631d3ce7f6804222754a3c27510d360c7bd18`.

### Compatibility when upgrading

Regeneration refreshes the API schema and shared generator code as well as
adding context and raw calls. It is not source compatible with every caller of
revision `1d9240f3d27b5831baf3694baab5fe06294875d8`.

An input struct that still compiles can change request behavior after a
schema refresh. Direct struct literals, direct field assignment and a nonnil
input with no selected-parameter map send every generated input field. New
fields therefore send their Go zero values instead of using server defaults.

In this update, `Vps.Replace` adds `PreserveBackups` and
`PreserveBackupHistory`. An older all-fields caller now sends both as `false`,
overriding the captured schema's `true` defaults and opting out of backup
preservation. Audit these callers before upgrading. Prefer `NewInput()` followed
by setters only for the intended fields; unselected preservation flags stay
omitted. Callers that deliberately send all fields must set both flags explicitly
according to their required backup policy. `Network.Create/Update` callers in
all-fields mode similarly send the new `Enabled` field as `false`; omit it with
selected setters or choose its value explicitly. Audit other added write flags
before upgrading an all-fields caller as well.

| API surface | Required caller change |
| --- | --- |
| `NewsLog.Create/Update` input | Replace `Message`/`SetMessage` with the appropriate `EnMessage` and `CsMessage` fields/setters. Output `Message` remains available. |
| `Node.Update` input | Remove use of `Cpus`, `TotalMemory`, `TotalSwap` and their setters; these inputs are absent from the current schema. |
| `SecurityAdvisory.NodeStatus.Create/Update` input | Replace `Note`/`SetNote` with localized notes. Review the added `ExpectedContentRevision` input against the API contract. |
| `SecurityAdvisory.NodeStatus.Create/Index/Update` output | Use the `Node` association and localized notes instead of `NodeId`, `NodeName` and `Note`. |

The generated tree also omits obsolete v5/v6 outputs for AuthToken and
SessionToken management, EnvironmentConfigChain, IntegrityCheck/Fact/Object,
IpTraffic/Monitor, VpsConfig, VpsOutageMount/Window, User TOTP management and
Cluster.AuthorizeMigrationKeys. These were already disconnected from the
previous client's resource wrappers, but their exported types and constructors
could still be used directly. Such callers must retain their existing module
pin until they have moved to a supported API. Token login helpers and
`UserSession` remain available.

Shared generator code adds language configuration and translated diagnostics;
English remains the default. `NewValidationError()` accepts an optional client:
ordinary calls still compile, but assignments to a `func() *ValidationError`
value need adjustment. Added generated struct fields can also affect callers
using unkeyed struct literals.

The OAuth revoke URL is now `https://auth.vpsfree.cz/_auth/oauth2/revoke`.
Callers must explicitly authorize that origin through `AllowOAuth2Origin`;
the existing origin checks remain in effect.

Review and compile each consumer before changing its module pin. Publishing a
client revision does not update consumers that retain an older pin. The module
keeps its existing Go 1.16 directive; that directive alone does not establish
that the generated snapshot has been tested with a Go 1.16 compiler.

### Regeneration

Run that generator into an empty output directory:

```sh
haveapi-go-client --version 7.0 \
  --module github.com/vpsfreecz/vpsadmin-go-client --package client \
  https://api.vpsfree.cz <empty-output-directory>
```

Review the emitted `go.mod` and complete `client/` tree, including additions and
obsolete generated files, before replacing the repository's generated tree.
The live API description can change; record its digest and review schema drift
when regenerating. Run `go test ./...` in the declared Nix shell before committing.
