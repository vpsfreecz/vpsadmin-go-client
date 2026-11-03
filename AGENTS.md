# Repository Guidelines

## Project Structure & Module Organization

This repository is a small Go module for the vpsAdmin API client.

- `go.mod` declares module `github.com/vpsfreecz/vpsadmin-go-client` and Go 1.16.
- `client/` contains the generated package used by importers as `github.com/vpsfreecz/vpsadmin-go-client/client`.
- `README.md` documents public usage examples, authentication, timeouts, and client regeneration.
- `shell.nix` provides a Nix development shell with Go, Git, Go tools and Python.
- `schema/` retains raw and effective API inputs, source/generator provenance
  and resolved dependency locks. `tools/prepare-schema` makes the bounded OAuth
  endpoint projection documented in the README.

Files in `client/` are generated resource/action types named by API resource and action, for example `resource_vps_action_start.go` and `resource_dns_record_action_show.go`.

## Build, Test, and Development Commands

- `nix-shell`: enter the repository's Nix development environment.
- `go test ./...`: compile and test all packages. In minimal environments without `gcc`, use `CGO_ENABLED=0 go test ./...`.
- `go test ./client`: run checks for only the generated client package.
- `go fmt ./...`: format Go files before committing.
- `python3 -m unittest discover -s tools -p '*_test.py'`: test schema preparation.
- `tools/prepare-schema`: derive the generation input from the saved raw export.
- Follow README's pinned saved-schema command to generate into a fresh directory;
  compare the complete generated file set, including removals, while preserving
  handwritten tests and `go.mod`.

There is no runnable application in this repository; it is a library module.

## Coding Style & Naming Conventions

Use standard Go formatting (`gofmt`/`go fmt`) with tabs for indentation. Keep package code under package name `client`. Follow the existing generated naming patterns: exported structs and constructors use PascalCase (`ResourceVps`, `NewActionVpsStart`), while filenames use lowercase snake-style resource/action names.

Do not edit generated files manually. Changes to `client/` must be made by regenerating the client with `haveapi-go-client`.

## Testing Guidelines

`go test ./...` compiles the generated package and checks OAuth revocation with
an HTTP transport stub. Python tests beside the schema helper check its bounded
projection and failure behavior. Keep tests deterministic and offline. Preserve
handwritten `*_test.go` files when replacing generated output.

## Commit & Pull Request Guidelines

Recent commits use short, imperative summaries, for example
`Add generated HTTP timeout support` and `Update client for current API
version`. Keep commits focused: separate regeneration-only changes from
handwritten fixes when practical.

Commit messages must say what is changing and why. Describe the problem in the
body and summarize the solution. Wrap subject and body lines at 80 characters.
Write commit messages through a temporary file and commit with `git commit -F`.
Git hooks must run for every commit; do not use `--no-verify`, `SKIP=...`,
`LEFTHOOK=0`, or similar to bypass them.

Pull requests should include a concise description, the reason for the update,
regeneration details if applicable, and the test command run. Link related
issues when available. Screenshots are not relevant for this library.

## Agent-Specific Instructions

For vpsAdmin API and generator changes, follow the canonical
[generated-client synchronization procedure](https://github.com/vpsfreecz/vpsadmin/blob/master/docs/api-schema-export.md#generated-client-synchronization)
before declaring synchronization complete.

Do not patch files in `client/` by hand. Regenerate them with `haveapi-go-client`, then review the resulting diff for scope and accidental churn.
