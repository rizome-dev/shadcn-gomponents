# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go template library project (`github.com/rizome-dev/tmpl`) built by rizome labs. It serves as a production-ready template for starting new Go projects with proper structure, CI/CD, and build tooling.

## Build Commands

The project uses `just` for build tasks. Available commands:

```bash
# Build for macOS (local)
just build-darwin-local

# Build for Linux (local)
just build-linux-local

# Build for Linux using Docker (ensures compatibility)
just build-linux-docker

# Bootstrap the current directory as a new project
just bootstrap                                    # Uses defaults
just bootstrap github.com/user/myproject          # Custom module name
just bootstrap github.com/user/myproject sdk  # Custom module and type
```

### Bootstrap Command

The `bootstrap` command initializes the current directory as a new Go project. It:
- Creates a go.mod file with the specified module name
- Sets up the project structure based on type (cli, sdk, sharedlib, api)
- Updates all imports from the template to your module name
- Creates minimal starter files for your project type
- Does NOT manipulate go.mod beyond creation
- Does NOT install dependencies (use `make setup` after bootstrap)

Note: Run bootstrap from the root of your cloned template repository.

## Development Setup

1. This is a Go 1.23.4 project
2. After bootstrap, run `make setup` to install development tools
3. Build outputs go to `./build/` directory
4. Use `make build` for building, `make test` for testing

## Project Structure

Standard Go project layout created by bootstrap:
- `cmd/` - Main applications
- `internal/` - Private application and library code  
- `pkg/` - Library code for external use
- `docs/` - Documentation

## Testing

Use standard Go testing practices with `go test` or `make test`.

## Important Notes

- The template supports multiple project types: CLI, SDK, shared C library, and API server
- Shared library projects can build .dylib (macOS) and .so (Linux) files for C interoperability
- Contact: hi@rizome.dev
- Documentation: https://pkg.go.dev/github.com/rizome-dev/tmpl