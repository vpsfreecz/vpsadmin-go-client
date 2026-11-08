# vpsAdmin Go client

A generated Go client for [vpsAdmin API 7.0](https://api.vpsfree.cz), with typed
resources, actions and inputs. It is generated with
[`haveapi-go-client`](https://github.com/vpsfreecz/haveapi/tree/master/clients/go).

## Installation

From your application's Go module:

```sh
go get github.com/vpsfreecz/vpsadmin-go-client/client
```

## Usage

Set `VPSADMIN_TOKEN` in your environment to an existing HaveAPI token. This
complete program reads the first page of up to 50 VPSes visible to that token
and prints their ID, hostname and running state:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vpsfreecz/vpsadmin-go-client/client"
)

func main() {
	token := os.Getenv("VPSADMIN_TOKEN")
	if token == "" {
		log.Fatal("VPSADMIN_TOKEN must be set")
	}

	api := client.New("https://api.vpsfree.cz")
	api.SetExistingTokenAuth(token)
	api.SetTimeout(30 * time.Second)

	action := api.Vps.Index.Prepare()
	action.NewInput().SetLimit(50)
	response, err := action.Call()
	if err != nil {
		log.Fatal(err)
	}
	if !response.Status {
		log.Fatalf("List VPSes failed: %s (%v)", response.Message, response.Errors)
	}

	for _, vps := range response.Output {
		fmt.Printf("%d\t%s\trunning=%t\n", vps.Id, vps.Hostname, vps.IsRunning)
	}
}
```

Use the unversioned API base URL. Generated action paths already contain the API
version. The server controls which resources and fields the token can access.

## Working with VPSes

The following snippets use the authenticated `api` from the usage example.
Choose a location and OS template, create a VPS, then use its returned ID to
inspect and manage it. Choose one creation example and only the power action
you need.

### Choose a location and OS template

List locations where you can create a VPS:

```go
action := api.Location.Index.Prepare()
action.NewInput().SetHasHypervisor(true)
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("List locations failed: %s (%v)", response.Message, response.Errors)
}
for _, location := range response.Output {
	fmt.Printf("%d\t%s\n", location.Id, location.Label)
}
```

Choose a location from the catalogue and store its ID in `locationId` (`int64`).
List templates available at that location:

```go
action := api.OsTemplate.Index.Prepare()
action.NewInput().SetLocation(locationId)
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("List OS templates failed: %s (%v)", response.Message, response.Errors)
}
for _, template := range response.Output {
	fmt.Printf("%d\t%s\n", template.Id, template.Label)
}
```

Choose a template from these results and store its ID in `templateId` (`int64`).

### Create a VPS

To create a VPS, choose a hostname, location and OS template. The minimal
request creates `my-vps`:

```go
action := api.Vps.Create.Prepare()
input := action.NewInput()
input.SetHostname("my-vps")
input.SetLocation(locationId)
input.SetOsTemplate(templateId)
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("Create VPS failed: %s (%v)", response.Message, response.Errors)
}
if response.Output == nil || response.Output.Id <= 0 {
	log.Fatal("Create VPS returned no usable VPS ID")
}
vpsId := response.Output.Id
fmt.Printf("VPS ID: %d\n", vpsId)
```

The VPS starts by default. Omitted resource sizes use the API's defaults. By
default it receives one IPv4 and one IPv6 allocation where IPv6 is available.

To choose resource sizes and public IP allocation counts, use this alternative
creation request:

```go
action := api.Vps.Create.Prepare()
input := action.NewInput()
input.SetHostname("my-vps")
input.SetLocation(locationId)
input.SetOsTemplate(templateId)
input.SetCpu(2)
input.SetMemory(2048)
input.SetDiskspace(10240)
input.SetIpv4(1)
input.SetIpv6(1)
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("Create VPS failed: %s (%v)", response.Message, response.Errors)
}
if response.Output == nil || response.Output.Id <= 0 {
	log.Fatal("Create VPS returned no usable VPS ID")
}
vpsId := response.Output.Id
fmt.Printf("VPS ID: %d\n", vpsId)
```

Memory and disk space use MiB: `2048` is 2 GiB and `10240` is 10 GiB. `cpu` sets
CPU capacity. Requested resources must fit your quota and be available at the
chosen location.

The `ipv4` and `ipv6` values are allocation counts, not addresses. An allocation
can be an individual address or a routed prefix. IPv6 requires support at the
chosen location.

Creating a VPS starts background work. Before continuing, check that creation
has completed successfully. A returned VPS ID or an unfinished operation does
not establish completion. Check completion again before each dependent action;
the same rule applies to power controls and deletion.

### Inspect and rename the VPS

Use `vpsId` from the selected create example to read the VPS:

```go
action := api.Vps.Show.Prepare()
action.SetPathParamInt("vps_id", vpsId)
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("Show VPS failed: %s (%v)", response.Message, response.Errors)
}
if response.Output == nil {
	log.Fatal("Show VPS returned no output")
}
fmt.Printf("%d\t%s\trunning=%t\n", response.Output.Id, response.Output.Hostname, response.Output.IsRunning)
```

To change only this VPS's hostname:

```go
action := api.Vps.Update.Prepare()
action.SetPathParamInt("vps_id", vpsId)
action.NewInput().SetHostname("new-hostname")
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("Rename VPS failed: %s (%v)", response.Message, response.Errors)
}
fmt.Println("Rename request accepted")
```

Input setters select which fields are sent. `SetHostname` sends only `hostname`,
even though the action accepts other inputs. Direct struct initialization or
field assignment without selecting any parameters sends all fields, including
zero values. Use setters for partial updates.

### Start, stop or restart

Each action uses the same `vpsId`. Run the operation you need as a separate
request, then check its completion before depending on the resulting state.

Start a stopped VPS. Creation already starts the VPS by default:

```go
action := api.Vps.Start.Prepare()
action.SetPathParamInt("vps_id", vpsId)
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("Start VPS failed: %s (%v)", response.Message, response.Errors)
}
fmt.Println("Start request accepted")
```

Stop the VPS gracefully:

```go
action := api.Vps.Stop.Prepare()
action.SetPathParamInt("vps_id", vpsId)
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("Stop VPS failed: %s (%v)", response.Message, response.Errors)
}
fmt.Println("Stop request accepted")
```

Restart the VPS gracefully:

```go
action := api.Vps.Restart.Prepare()
action.SetPathParamInt("vps_id", vpsId)
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("Restart VPS failed: %s (%v)", response.Message, response.Errors)
}
fmt.Println("Restart request accepted")
```

### Delete the VPS

Delete the VPS only when you no longer need it.

Deletion schedules the VPS to stop and its IP allocations to be released.
Disk data is removed later.

```go
action := api.Vps.Delete.Prepare()
action.SetPathParamInt("vps_id", vpsId)
response, err := action.Call()
if err != nil {
	log.Fatal(err)
}
if !response.Status {
	log.Fatalf("Delete VPS failed: %s (%v)", response.Message, response.Errors)
}
fmt.Println("Delete request accepted")
```

## Authentication

The usage example uses `SetExistingTokenAuth` with a token supplied through the
environment. HTTP Basic authentication and obtaining a new token with login
credentials are also supported. Keep credentials outside source code and never
print tokens.

For HTTP Basic authentication:

```go
user := os.Getenv("VPSADMIN_USER")
password := os.Getenv("VPSADMIN_PASSWORD")
if user == "" || password == "" {
	log.Fatal("VPSADMIN_USER and VPSADMIN_PASSWORD must be set")
}
api.SetBasicAuthentication(user, password)
```

To obtain a new token, check the authentication error before making API calls:

```go
user := os.Getenv("VPSADMIN_USER")
password := os.Getenv("VPSADMIN_PASSWORD")
if user == "" || password == "" {
	log.Fatal("VPSADMIN_USER and VPSADMIN_PASSWORD must be set")
}
err := api.SetNewTokenAuth(&client.TokenAuthOptions{
	User:     user,
	Password: password,
})
if err != nil {
	log.Fatal(err)
}
```

Two-factor authentication is supported through token authentication. When the
server requests TOTP, `SetNewTokenAuth` calls `TokenAuthOptions.TotpCallback`.
Pass the callback in `TokenAuthOptions`, importing `bufio` and `strings`:

```go
options := &client.TokenAuthOptions{
	User:     user,
	Password: password,
	TotpCallback: func(input *client.AuthTokenActionTokenTotpInput) error {
		fmt.Print("TOTP code: ")
		code, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		input.SetCode(strings.TrimSpace(code))
		return nil
	},
}
if err := api.SetNewTokenAuth(options); err != nil {
	log.Fatal(err)
}
```

The server decides which authentication steps are required. If it requires a
password reset, provide `TokenAuthOptions.ResetPasswordCallback` as well.

### OAuth token revocation

`SetExistingOAuth2Auth` accepts an existing OAuth access token. Token revocation
uses `https://auth.vpsfree.cz/_auth/oauth2/revoke`.
For a client at `https://api.vpsfree.cz`, explicitly trust the authentication
origin with `AllowOAuth2Origin("https://auth.vpsfree.cz")` before calling
`RevokeAccessToken`. Untrusted cross-origin requests fail before dispatch.
Successful revocation clears client authentication; a failed request preserves it.

## Responses, errors and timeouts

Check the `error` returned by `Call()` first; it covers request and decoding
failures. Then check `response.Status` before using `response.Output`. API
failures provide `Message` and per-parameter `Errors`.
`Output` exposes the action result directly.
`Response` contains the modeled response namespace and any metadata exposed by
that action. Scalar output fields use Go zero values, so an omitted field cannot
be distinguished from a returned `false`, `0` or empty string.

The list example reads one page. Request further pages with the index input's
`SetFromId` setter, passing the last returned VPS ID, and keep `SetLimit(50)`.
Check each response before reading its output and stop when a page is empty.

The client uses `http.DefaultClient` unless configured otherwise. The usage
example sets an overall 30-second request timeout. You can supply an
`*http.Client` through `SetHTTPClient` when you need custom transport settings.

Some actions start background work. A successful response does not mean that
work has finished. A timeout or lost connection during a write does not prove
the write failed. Inspect the server's state before retrying an uncertain write.

Use `api.SetLanguage("cs")` to select Czech messages and send `Accept-Language`;
check its returned error. `SetLanguageHeader` changes the request header used
for that value. Client-side validation messages support Czech and English and
fall back to English.

## Custom input parameters

The Go generator omits parameters whose schema type is `Custom`. Generated
input structs and request maps cannot send these fields. An action can have a
typed accessor even when its required Custom input is unsupported. In the
captured API 7.0 contract, this affects:

| Typed actions | Omitted input | Effect |
| --- | --- | --- |
| `IpReleaseCampaign.Create`, `IpReleaseCampaign.Exempt`, `IpReleaseRequest.Keep` | `addresses` | Required; these typed calls cannot supply it |
| `WebuiUserSetting.Set` | `value` | Required; the typed call cannot supply it |
| `Webauthn.Authentication.Finish`, `Webauthn.Registration.Finish` | `public_key_credential` | Required; these typed calls cannot supply it |
| `SystemConfig.Update` | `value` | Optional; the typed call cannot update it |

For these inputs, use the existing low-level
`Client.DoBodyRequest(method, path, params, output)` with a complete namespaced
request body and a suitable response type, or use another client. This method
JSON-encodes the supplied body and uses the client's authentication and
language settings. It bypasses generated input validation and does not populate
a generated response's convenience `Output` field. Check its returned error,
then the API response's `Status`, `Message` and `Errors`; read the returned
payload from its `Response` namespace. Encode variable URL path segments
appropriately.

## Development and regeneration

For vpsAdmin API and generator changes, follow the
[generated-client synchronization procedure](https://github.com/vpsfreecz/vpsadmin/blob/master/docs/api-schema-export.md#generated-client-synchronization)
for the required client checks and completion records. Synchronizing generated
source does not by itself require a release.

The saved API 7.0 schema is a source-fixture export from vpsAdmin commit
`68a36237b0bad9625fb25965bd415e801ddccf4c`, with all six bundled plugins and
English metadata. The source fixture does not establish which API version or
permissions a deployed server provides. Read [CHANGELOG.md](CHANGELOG.md) before
upgrading.

The raw export is [schema/source-api-7.0.json](schema/source-api-7.0.json), with
its unmodified [provenance](schema/provenance.json) and exporter dependency lock.
`tools/prepare-schema` derives [schema/api-7.0.json](schema/api-7.0.json) by
changing only the fixture's OAuth revocation URL to
`https://auth.vpsfree.cz/_auth/oauth2/revoke`. It rejects an unexpected fixture
profile or endpoint.
[schema/generator.json](schema/generator.json) records both hashes, the exact
change, generator source, configuration and toolchain. This preserves endpoint
selection without changing the client's origin trust checks.

Use `nix-shell` for Go and Python checks:

```sh
python3 -m unittest discover -s tools -p '*_test.py'
go test ./...
```

The Go tests use an HTTP transport stub and make no production requests. The
files in `client/`, except handwritten `*_test.go` files, are generated.
Prepare the input from this repository's root:

```sh
./tools/prepare-schema
```

Use a separate HaveAPI checkout at commit
`1ef9b204ba45bb261d3334be5c555cc1856b713b` (generator 0.32.2). Restore
`schema/generator.Gemfile.lock` as that checkout's `clients/go/Gemfile.lock` and
prepare its component bundle as described in the [Go generator
guide](https://github.com/vpsfreecz/haveapi/blob/1ef9b204ba45bb261d3334be5c555cc1856b713b/clients/go/README.md).
From the prepared HaveAPI environment, with `GO_CLIENT` pointing to this
repository and `FRESH_OUTPUT` to a fresh output directory, run:

```sh
BUNDLE_GEMFILE="$PWD/clients/go/Gemfile" bundle _4.0.23_ exec ruby \
  clients/go/bin/haveapi-go-client generate \
  --schema "$GO_CLIENT/schema/api-7.0.json" --api-version 7.0 \
  --output "$FRESH_OUTPUT" \
  --module github.com/vpsfreecz/vpsadmin-go-client --package client
```

Compare the fresh `client/*.go` output with this repository, including removed
files. Preserve handwritten tests and the existing `go.mod`. The generator does
not remove stale files when writing into an existing directory. Review the
complete diff before committing; do not edit generated files by hand.
