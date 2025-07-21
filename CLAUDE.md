# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

```bash
# Run tests with coverage
make test

# Run benchmarks
make benchmark

# Build CSS (required after style changes)
make build-css

# Watch CSS for development
make watch-css

# Run linter
make lint

# View test coverage report
make cover

# Run the demo application (builds CSS + starts server)
make start

# Build Docker images
make build-docker

# Run a single test
go test -v ./pkg/button -run TestButton

# Run tests for a specific package
go test -v ./pkg/accordion/...
```

## Architecture Overview

This is a Go implementation of shadcn/ui components using the gomponents library. Components are designed for server-side rendering with optional HTMX enhancement for interactivity.

### Component Structure

Each component lives in `pkg/<component-name>/` with this consistent structure:

1. **component.go** - Core component implementation
   - Props struct with variant configuration
   - Constructor functions (e.g., `New()`, `Root()`)
   - Helper functions for common patterns

2. **component_htmx.go** - HTMX-enhanced version (if interactive)
   - HTMXProps struct for HTMX configuration
   - Enhanced constructors with HTMX attributes
   - Server-side handlers for dynamic behavior

3. **component_test.go** - Table-driven tests
   - Tests HTML output contains expected elements
   - Verifies variant helpers work correctly

4. **example.go** - Comprehensive examples
   - Shows all variants and configurations
   - Used by the demo application

### Key Libraries

- **maragu.dev/gomponents** - HTML generation (aliased as `g`)
- **maragu.dev/gomponents/html** - HTML elements (dot-imported)
- **maragu.dev/gomponents-htmx** - HTMX attributes
- **lib/cn** - Class merging utilities (like clsx)
- **lib/variants** - Variant system for component styling

### Variant System

Components use a consistent variant system defined in `lib/variants.go`:

```go
variantConfig := lib.VariantConfig{
    Base: "base-classes",
    Variants: map[string]map[string]string{
        "variant": {
            "default": "default-classes",
            "secondary": "secondary-classes",
        },
        "size": {
            "sm": "small-classes",
            "md": "medium-classes",
        },
    },
    Defaults: map[string]string{
        "variant": "default",
        "size": "md",
    },
}
```

### Creating New Components

When creating a new component:

1. Create package in `pkg/<component-name>/`
2. Define Props struct with embedded `lib.VariantProps` if using variants
3. Create variant configuration
4. Implement constructor functions
5. Add HTMX support in separate file if needed
6. Write comprehensive examples in `example.go`
7. Add table-driven tests

### Testing Patterns

Tests verify HTML output using `strings.Contains`:

```go
tests := []struct {
    name     string
    props    Props
    children []g.Node
    contains []string // Expected HTML fragments
}{
    {
        name: "renders with default variant",
        props: Props{Variant: "default"},
        contains: []string{
            `class="`, 
            "default-variant-classes",
        },
    },
}
```

### HTMX Integration

Components with interactivity use HTMX. Server endpoints handle state:

```go
// In component_htmx.go
type HTMXProps struct {
    ID          string
    TriggerPath string
    UpdatePath  string
}

// Server handler updates component state
func HandleToggle(w http.ResponseWriter, r *http.Request) {
    // Return updated component HTML
}
```

### CSS and Styling

- Uses Tailwind CSS with custom theme configuration
- CSS variables for theming (light/dark mode)
- Classes merged using `lib.CN()` to handle duplicates
- Build CSS with `make build-css` after style changes

### Important Conventions

1. **Imports**: Always alias gomponents as `g` and dot-import html package
2. **Naming**: Use `New()` for main constructor, `Root()` for container components
3. **Props**: Embed `lib.VariantProps` for components with variants
4. **Testing**: Include examples of all variants in tests
5. **Documentation**: Examples in `example.go` serve as documentation