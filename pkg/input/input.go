package input

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Input component
type Props struct {
	Type         string // "text" | "email" | "password" | "number" | etc.
	Placeholder  string
	Name         string
	ID           string
	Value        string
	Disabled     bool
	Required     bool
	AriaInvalid  bool
	AutoComplete string
	Class        string // Additional custom classes
}

// inputClasses defines the base classes for the input component
const inputClasses = "file:text-foreground placeholder:text-muted-foreground selection:bg-primary selection:text-primary-foreground dark:bg-input/30 border-input flex h-9 w-full min-w-0 rounded-md border bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none file:inline-flex file:h-7 file:border-0 file:bg-transparent file:text-sm file:font-medium md:text-sm focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"

// New creates a new Input component
func New(props Props) g.Node {
	// Set default type if not provided
	if props.Type == "" {
		props.Type = "text"
	}

	// Combine classes
	classes := lib.CN(inputClasses, props.Class)

	// Build attributes
	attrs := []g.Node{
		html.Class(classes),
		html.Type(props.Type),
	}

	// Add optional attributes
	if props.Placeholder != "" {
		attrs = append(attrs, html.Placeholder(props.Placeholder))
	}
	if props.Name != "" {
		attrs = append(attrs, html.Name(props.Name))
	}
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}
	if props.Value != "" {
		attrs = append(attrs, html.Value(props.Value))
	}
	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}
	if props.Required {
		attrs = append(attrs, html.Required())
	}
	if props.AriaInvalid {
		attrs = append(attrs, g.Attr("aria-invalid", "true"))
	}
	if props.AutoComplete != "" {
		attrs = append(attrs, html.AutoComplete(props.AutoComplete))
	}

	return html.Input(attrs...)
}

// Text creates a text input
func Text(placeholder string) g.Node {
	return New(Props{
		Type:        "text",
		Placeholder: placeholder,
	})
}

// Email creates an email input
func Email(placeholder string) g.Node {
	return New(Props{
		Type:        "email",
		Placeholder: placeholder,
		AutoComplete: "email",
	})
}

// Password creates a password input
func Password(placeholder string) g.Node {
	return New(Props{
		Type:        "password",
		Placeholder: placeholder,
		AutoComplete: "current-password",
	})
}

// Number creates a number input
func Number(placeholder string) g.Node {
	return New(Props{
		Type:        "number",
		Placeholder: placeholder,
	})
}

// Search creates a search input
func Search(placeholder string) g.Node {
	return New(Props{
		Type:        "search",
		Placeholder: placeholder,
	})
}

// Tel creates a telephone input
func Tel(placeholder string) g.Node {
	return New(Props{
		Type:        "tel",
		Placeholder: placeholder,
		AutoComplete: "tel",
	})
}

// URL creates a URL input
func URL(placeholder string) g.Node {
	return New(Props{
		Type:        "url",
		Placeholder: placeholder,
		AutoComplete: "url",
	})
}