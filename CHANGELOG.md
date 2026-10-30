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
