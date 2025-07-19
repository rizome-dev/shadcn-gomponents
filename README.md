# tmpl

[![GoDoc](https://pkg.go.dev/badge/github.com/rizome-dev/tmpl)](https://pkg.go.dev/github.com/rizome-dev/tmpl)
[![Go Report Card](https://goreportcard.com/badge/github.com/rizome-dev/tmpl)](https://goreportcard.com/report/github.com/rizome-dev/tmpl)
[![CI](https://github.com/rizome-dev/tmpl/actions/workflows/ci.yml/badge.svg)](https://github.com/rizome-dev/tmpl/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A production-ready Go project template with CI/CD, testing, and build tooling.

built by [rizome labs](https://rizome.dev) | contact: [hi@rizome.dev](mailto:hi@rizome.dev)

## Quick Start

```bash
# Use this template on GitHub
# 1. Click "Use this template" button
# 2. Create your new repository
# 3. Clone your new repo

# Bootstrap your project
cd your-new-repo
just bootstrap github.com/yourusername/yourproject cli

# Set up development tools
make setup

# Build and test
make build
make test
```

## Bootstrap Options

```bash
just bootstrap                                    # Uses defaults (CLI project)
just bootstrap github.com/user/myproject          # Custom module name
just bootstrap github.com/user/myproject library  # Library project
just bootstrap github.com/user/myproject api      # API server project
```

Project types:
- `cli` - Command-line application
- `library` - Go library package
- `api` - HTTP API server
- `sharedlib` - Shared C library (.so/.dylib)

## Development

```bash
# Install tools
make setup

# Build
make build

# Test
make test

# Lint
make lint

# All checks
make ci
```

## Features

- **Go 1.23.4** with modern project structure
- **GitHub Actions** CI/CD pipeline
- **golangci-lint** with comprehensive rules
- **Multiple project types** supported
- **Cross-platform builds** (Linux, macOS, Windows)
- **Shared C library** support
- **Docker** build support

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see the [LICENSE](LICENSE) file for details.