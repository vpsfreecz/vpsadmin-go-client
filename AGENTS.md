# Repository Guidelines

## Project Structure & Module Organization

This repository is a small Go module for the vpsAdmin API client.

- `go.mod` declares module `github.com/vpsfreecz/vpsadmin-go-client` and Go 1.16.
- `client/` contains the generated package used by importers as `github.com/vpsfreecz/vpsadmin-go-client/client`.
- `README.md` documents public usage examples, authentication, timeouts, and client regeneration.
- `shell.nix` provides a Nix development shell with Go, Git, and Go tools.

Files in `client/` are generated resource/action types named by API resource and action, for example `resource_vps_action_start.go` and `resource_dns_record_action_show.go`.

## Build, Test, and Development Commands

- `nix-shell`: enter the repository's Nix development environment.
- `go test ./...`: compile and test all packages. In minimal environments without `gcc`, use `CGO_ENABLED=0 go test ./...`.
- `go test ./client`: run checks for only the generated client package.
- `go fmt ./...`: format Go files before committing.
- `haveapi-go-client https://api.vpsfree.cz ./client`: regenerate the client from the current vpsAdmin API.

There is no runnable application in this repository; it is a library module.

## Coding Style & Naming Conventions

Use standard Go formatting (`gofmt`/`go fmt`) with tabs for indentation. Keep package code under package name `client`. Follow the existing generated naming patterns: exported structs and constructors use PascalCase (`ResourceVps`, `NewActionVpsStart`), while filenames use lowercase snake-style resource/action names.

Do not edit generated files manually. Changes to `client/` must be made by regenerating the client with `haveapi-go-client`.

## Testing Guidelines

The repository currently has no `*_test.go` files, so `go test ./...` is primarily a compile-time smoke test. Add tests beside the package code as `*_test.go` when adding handwritten logic, bug fixes, or behavior not covered by generation. Keep tests deterministic and avoid live API calls unless explicitly gated.

## Commit & Pull Request Guidelines

Recent commits use short, imperative summaries, for example `Add generated HTTP timeout support` and `Update client for current API version`. Keep commits focused: separate regeneration-only changes from handwritten fixes when practical.

Pull requests should include a concise description, the reason for the update, regeneration details if applicable, and the test command run. Link related issues when available. Screenshots are not relevant for this library.

## Agent-Specific Instructions

Do not patch files in `client/` by hand. Regenerate them with `haveapi-go-client`, then review the resulting diff for scope and accidental churn.
