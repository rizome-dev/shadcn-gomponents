package label

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Label component
type Props struct {
	For      string // The ID of the input this label is for
	Required bool   // Whether to show a required indicator
	Class    string // Additional custom classes
}

// labelClasses defines the base classes for the label component
const labelClasses = "flex items-center gap-2 text-sm leading-none font-medium select-none group-data-[disabled=true]:pointer-events-none group-data-[disabled=true]:opacity-50 peer-disabled:cursor-not-allowed peer-disabled:opacity-50"

// New creates a new Label component
func New(props Props, children ...g.Node) g.Node {
	// Combine classes
	classes := lib.CN(labelClasses, props.Class)

	// Build attributes
	attrs := []g.Node{
		html.Class(classes),
	}

	// Add optional attributes
	if props.For != "" {
		attrs = append(attrs, html.For(props.For))
	}

	// Add required indicator if needed
	if props.Required {
		children = append(children, 
			html.Span(
				html.Class("text-destructive"),
				g.Text("*"),
			),
		)
	}

	// Combine attributes and children
	return html.Label(append(attrs, children...)...)
}

// Default creates a basic label
func Default(text string) g.Node {
	return New(Props{}, g.Text(text))
}

// WithRequired creates a label with required indicator
func WithRequired(text string, required bool) g.Node {
	return New(Props{Required: required}, g.Text(text))
}

// ForInput creates a label associated with an input
func ForInput(forID, text string) g.Node {
	return New(Props{For: forID}, g.Text(text))
}

// ForInputRequired creates a label associated with an input and shows required indicator
func ForInputRequired(forID, text string) g.Node {
	return New(Props{
		For:      forID,
		Required: true,
	}, g.Text(text))
}