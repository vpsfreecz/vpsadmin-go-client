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

## VPS evidence

VPS evidence provides current rolling credit for one VPS, an overview of VPSes
with positive credit, and immutable intervals saved with an incident report.
Members can read their own VPSes; administrators can read all VPSes permitted by
the server. Check server capabilities before using these endpoints.

Use `DoQueryStringRequest` with an explicit path and a caller-owned response
struct for evidence reads. Generated scalar fields collapse null to zero,
false or an empty string, and generated responses omit global metadata and
Custom fields. These omissions lose the distinction between unavailable and
recorded credit, optional peaks and pagination facts.

The generated client intentionally has no `api.Vps.Evidence` accessor. The
pinned generator gives it the same Go name as the top-level `api.VpsEvidence`
and would route the child accessor to the fleet endpoint. Schema preparation
omits only that child from generation. Its API endpoint remains available through
the explicit request below. `api.VpsEvidence` selects the fleet endpoint, but
its generated response still has the scalar and metadata limitations above.

| Read | GET path | Query namespace | Response rows |
| --- | --- | --- | --- |
| Current credit for VPS 12 | `/v7.0/vpses/12/evidence` | No inputs | `evidence` |
| Positive-credit overview | `/v7.0/vps_evidence` | `vps_evidence` | `vps_evidence` |
| Intervals saved with report 12 | `/v7.0/incident_reports/12/evidence_intervals` | `evidence_interval` | `evidence_intervals` |

The following example uses the authenticated `api` from the usage example.
Its custom row selects credit, observation and peak fields; add other fields
from the API contract as your application needs them, preserving nullable scalars
as pointers.

```go
var current struct {
	client.Envelope
	Response struct {
		Rows []struct {
			Signal             string  `json:"signal"`
			CreditUs           *int64  `json:"credit_us"`
			ConfigurationValid bool    `json:"configuration_valid"`
			ObservationState   string  `json:"observation_state"`
			WindowComplete     *bool   `json:"window_complete"`
			ReportingEnabled   *bool   `json:"reporting_enabled"`
			ReportEligible     *bool   `json:"report_eligible"`
			ReportDue          *bool   `json:"report_due"`
			Measurement        *string `json:"measurement"`
			PeakUnit           *string `json:"peak_unit"`
			PeakValue          *int64  `json:"peak_value"`
			PeakIsLowerBound   *bool   `json:"peak_is_lower_bound"`
		} `json:"evidence"`
		Meta struct {
			ReadAt             string   `json:"read_at"`
			ConfigurationValid bool     `json:"configuration_valid"`
			UnavailableSignals []string `json:"unavailable_signals"`
		} `json:"_meta"`
	} `json:"response"`
}
err := api.DoQueryStringRequest("/v7.0/vpses/12/evidence", nil, &current)
if err != nil {
	log.Fatal(err)
}
if !current.Status {
	log.Fatalf("Read VPS evidence failed: %s (%v)", current.Message, current.Errors)
}
for _, row := range current.Response.Rows {
	if row.CreditUs == nil {
		fmt.Printf("%s: credit unavailable\n", row.Signal)
		continue
	}
	fmt.Printf("%s: %d microseconds of recorded credit\n", row.Signal, *row.CreditUs)
}
```

Read `Response`, not a generated convenience `Output` field. The method sends
an authenticated GET with query parameters and no body. Check its returned
error and then `Status`, `Message` and `Errors` before using the response.

The per-VPS response includes every supported signal, including no-data rows.
A nil `CreditUs` means unavailable, even when `ConfigurationValid` is true,
such as during an ownership change. A pointer to zero means no recorded credit;
it does not establish VPS health or collection coverage. Positive credit can
remain visible while observations are stale or reporting is disabled.
Freshness, source qualification, reporting eligibility and whether a report is
due are separate facts. Read `Measurement` and `PeakUnit` with `PeakValue`;
`PeakIsLowerBound` identifies an unfinished contributing observation.

The overview includes only positive credit, including subsecond values. Use
query keys such as `vps_evidence[offset]`, `vps_evidence[limit]`,
`vps_evidence[vps]` and `vps_evidence[signal]`. Optional `vps_evidence[user]`,
`vps_evidence[node]`, `vps_evidence[location]` and `vps_evidence[environment]`
filter by resource ID; `user` is administrator-only. The server applies filters
within the caller's permissions before counting and paging. Decode rows under
`vps_evidence` into custom pointer fields, and retain `_meta.total_count`, `offset`, `limit`
and nullable `next_offset` alongside the current metadata shown above.
`unavailable_signals` is a sorted array of selected signals with invalid
configuration; an empty array means none. Missing or stale observations alone
do not make configuration invalid.

For saved intervals, send `evidence_interval[offset]` and
`evidence_interval[limit]`, then decode `evidence_intervals` into custom rows.
Pagination starts at offset 0, defaults to 25 rows and allows at most 100 rows.
Keep timestamps as strings to retain microsecond precision. Credit uses integer
microseconds; optional peak measurements have explicit units and nullable values.
Use pointers for nullable measurement, peak and lower-bound fields.

Saved `_meta` includes `available`, counts and pagination, plus nullable
`window_complete`, `evaluated_at`, `window_seconds`, `threshold_seconds`,
`repeat_seconds` and `credit_us`. Its nullable `condition_measurement`,
`condition_operator`, `condition_value` and `condition_unit` describe the signal
predicate saved with the incident. Use `*string`, `*int64` and `*bool` as
appropriate, check `available`, and follow `next_offset` for another page.
These facts describe the incident when it was created and remain unchanged
when current settings change. A non-nil false value differs from an unknown
value. Access to the parent report remains subject to server authorization.

The generic `evidence_interval` resource replaces the draft `rtnl_interval`
resource without a compatibility alias. Use the matching backend API before
calling the new endpoints.

### Memory pressure history

Use the root `/v7.0/vps_pressure_history` endpoint for daily recorded time.
The generated `api.VpsPressureHistory` accessor has the same scalar and metadata
limitations as evidence reads. Use a custom response to preserve nullable
values and all history metadata:

```go
var history struct {
	client.Envelope
	Response struct {
		Rows []struct {
			ID         *string `json:"id"`
			FromAt     *string `json:"from_at"`
			ToAt       *string `json:"to_at"`
			PartialDay *bool   `json:"partial_day"`
			RecordedUs *int64  `json:"recorded_us"`
		} `json:"vps_pressure_history"`
		Meta struct {
			ReadAt           *string `json:"read_at"`
			Timezone         *string `json:"timezone"`
			RangeFrom        *string `json:"range_from"`
			RangeTo          *string `json:"range_to"`
			TotalRecordedUs  *int64  `json:"total_recorded_us"`
			FirstRecordedAt  *string `json:"first_recorded_at"`
			AccountingSince  *string `json:"accounting_since"`
			OwnershipLimited *bool   `json:"ownership_limited"`
			Available        *bool   `json:"available"`
		} `json:"_meta"`
	} `json:"response"`
}
err := api.DoQueryStringRequest("/v7.0/vps_pressure_history", map[string]string{
	"vps_pressure_history[vps]":  "12",
	"vps_pressure_history[days]": "30",
}, &history)
if err != nil {
	log.Fatal(err)
}
if !history.Status {
	log.Fatalf("Read pressure history failed: %s (%v)", history.Message, history.Errors)
}
```

Read `history.Response.Rows` and `history.Response.Meta`, preserving nil
pointers. Keep timestamps as strings to retain microsecond precision.

The server requires a VPS ID and accepts `days` as 7 or 30, defaulting to 7.
The response contains that many ascending UTC calendar days, including today.
Each row spans `from_at` to `to_at`; today's end is next UTC midnight and
`partial_day` is true. `read_at` records when the response was read. Accepted
contributions within the server's clock tolerance are not clipped to that time.

`recorded_us` contains positive integer microseconds or null for a day without
recorded data. `total_recorded_us` sums the returned positive values and is null
when none are available. Preserve these unknown values instead of displaying
zero. History uses daily totals retained for 32 UTC calendar days; it is not
backfilled from incident reports or old observations and has no historical
reporting threshold.

History is scoped to the VPS's current owner and exact accounting epoch.
`accounting_since` is nullable; `ownership_limited` says whether that epoch
starts inside the requested range. `available` can be true without any recorded
days. A pending transfer or known ownership-state mismatch makes history
unavailable; a foreign VPS remains an authorization error. Current observation
freshness and reporting settings do not decide history availability.
`first_recorded_at` is the earliest retained acceptance timestamp for this owner
and epoch across the full 32-day retention period, independent of `days`.
It does not establish the first occurrence or the start of collection coverage.
Check the server's capabilities and authorization before offering this read.

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

This development snapshot uses a source-fixture API 7.0 export with all six
bundled plugins and English metadata. Its exact vpsAdmin revision is recorded
in [schema/provenance.json](schema/provenance.json).
The source fixture does not establish which API version or
permissions a deployed server provides. Read [CHANGELOG.md](CHANGELOG.md) before
upgrading.

Use `IncidentPolicy.Index` to discover policies and `Vps.IncidentPolicy.Index`
or `Vps.IncidentPolicy.Show` to read their current rows for a VPS. Policy keys
come from the registry at runtime; pass the returned string key in the path.
`Vps.IncidentPolicy.Update` preserves explicit null override resets and returns
the authoritative policy row. Change shared VPS shutdown permission through
`Vps.Update` separately, then use `Vps.Show` to confirm its current value.
Index filters are sent only when selected; initial reads need no policy selector.
Check server capabilities before enabling these calls.

Current rows include counts, effective limits, overrides and `ConfigurationValid`.
Limits reflect the declaration and current overrides. Invalid configurations
return null effective limits and status. Go scalar fields decode missing or null
values to zero, false or an empty string. Check `ConfigurationValid` before
interpreting limits or status. Use a reader that preserves null when you need to
distinguish missing values from scalar values.

Evaluation IDs, timestamps, states and incident-report policy evaluations describe
saved history. They do not supply current limits or predict whether a policy will
match a future event. The server validates policy keys and status strings.

Saved evaluations include `Matched`, `WarningPeriodDays` and `StopPeriodDays`.
The period fields describe the saved count windows. As a scalar boolean,
`Matched` cannot distinguish `false` from null or a missing field. A report's
notice policy key and label identify its saved originating policy; null or
missing values become empty strings. They do not select a current policy.
Active policy metadata has no `source_type` field.

The raw export is [schema/source-api-7.0.json](schema/source-api-7.0.json), with
its unmodified [provenance](schema/provenance.json) and exporter dependency lock.
`tools/prepare-schema` derives [schema/api-7.0.json](schema/api-7.0.json) with
two bounded changes: it replaces the fixture's OAuth revocation URL with
`https://auth.vpsfree.cz/_auth/oauth2/revoke` and omits the nested `vps.evidence`
resource to avoid the Go name collision described above. The raw export stays
unchanged. The effective schema contains 159 application resources and 533
actions; the raw export contains 160 resources, 534 actions and 9,176 parameters. Authentication remains one
resource with five actions in both inputs.

Preparation rejects an unexpected fixture profile, OAuth endpoint or evidence
resource signature before replacing the output. It retains the top-level
`vps_evidence` resource and checks both evidence routes and namespaces.
[schema/generator.json](schema/generator.json) records raw and effective hashes,
both transformations, helper identity, generator source, configuration and
toolchain. OAuth origin trust checks remain unchanged.

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
