# Unreleased

- Add DNS primary transfer checks, flags showing whether results are current,
  secondary serving status and historical transfer kinds. The new operations
  and fields require a compatible API. Missing evidence does not mean the
  primary is ready.
- Added exported fields can break unkeyed struct literals. Existing scalar
  outputs still cannot distinguish omitted values from zero values.

# 0.2.0

- Regenerate API 7.0 from vpsAdmin source commit
  `942c6eaed41331f823d8e0f670ed239e2b308440` with HaveAPI generator 0.32.2.
  Retain the full-fixture export, dependency locks and exact generation receipt.
- Add `SecurityAdvisory.Retract` and `SecurityAdvisory.Delete` (alias `Destroy`).
  Add `ExpectedState` and `SetExpectedState` to advisory Update and CVE/node-status
  Create, Update and Delete inputs. Draft synchronization consumers can send
  `expected_state: draft` alongside their content-revision precondition.
- No existing generated members are removed or change field types. Added fields
  can break unkeyed composite literals of exported structs. Keyed literals still
  compile, but input literals or direct field assignments without parameter
  selection can send an empty `expected_state`, which the server rejects. Use
  input setters or `SelectParameters` to select the fields sent to the server and
  omit an unused state guard. New calls and state guards require compatible API
  support.
- Preserve the shipped HTTPS OAuth revocation endpoint through a single-field
  input projection. Origin trust checks and clear-on-success behavior remain
  unchanged and have offline transport-stub coverage. Custom input limitations
  remain as documented in the README.

# 0.1.0

- Regenerate the client from the public vpsAdmin API 7.0 snapshot with HaveAPI
  0.32.0. Include IP release campaigns and requests, node kernel/eBPF/software
  state, password-change logs and WebUI settings.
- Keep the inherited limitation that generated inputs omit `Custom` parameters.
  Typed IP-release campaign create/exempt, IP-release request keep and WebUI
  setting calls cannot supply their required Custom inputs. The same limitation
  affects WebAuthn finish calls and SystemConfig's optional value. Use the
  low-level request method or another client when those fields are needed; see
  the README's Custom input parameters section.
- Add `SetLanguage`/`GetLanguage` and
  `SetLanguageHeader`/`GetLanguageHeader` for request language selection and
  Czech/English client-side validation messages. Use the production OAuth
  token-revocation endpoint from the captured API description.
- Replace news-log message and security-advisory note inputs with
  `CsMessage`/`EnMessage` and `CsNote`/`EnNote`. Advisory node-status outputs
  expose `Node` instead of `NodeId`/`NodeName`. Node updates no longer set
  `Cpus`, `TotalMemory` or `TotalSwap`. Include current advisory
  content-revision, user time-zone and VPS replacement backup-preservation
  inputs.
- Remove 50 obsolete generated files absent from the captured contract. These
  bindings were already disconnected from `Client` resource/action accessors.
  Code using their exported types or constructors directly must migrate to the
  current API contract. Affected groups: `AuthToken`, `SessionToken`,
  `EnvironmentConfigChain`, `IntegrityCheck`/`Fact`/`Object`,
  `IpTraffic`/`IpTrafficMonitor`, `VpsConfig`, `VpsOutageMount`/`Window`,
  `ClusterAuthorizeMigrationKeys` and `UserTotpConfirm`/`Disable`/`Enable`.
- Allow an optional `Client` argument to `NewValidationError` for language
  selection. Zero-argument calls remain valid. Code storing the constructor in
  a variable of the exact type `func() *ValidationError` must use a wrapper.
