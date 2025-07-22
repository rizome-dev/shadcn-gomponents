package separator

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Separator component
type Props struct {
	Orientation string // "horizontal" | "vertical"
	Decorative  bool   // Whether the separator is decorative (true) or semantic (false)
	Class       string // Additional custom classes
}

// New creates a new Separator component
func New(props Props) g.Node {
	// Set defaults
	if props.Orientation == "" {
		props.Orientation = "horizontal"
	}
	// Default to decorative if not specified (matching shadcn-ui behavior)
	// Since bool fields default to false in Go, and false means non-decorative,
	// we don't need to set a default - the zero value (false) means semantic

	// Build base classes
	baseClasses := "shrink-0 bg-border"
	
	// Add orientation-specific classes
	var orientationClasses string
	if props.Orientation == "horizontal" {
		orientationClasses = "h-[1px] w-full"
	} else {
		orientationClasses = "h-full w-[1px]"
	}

	// Combine all classes
	classes := lib.CN(baseClasses, orientationClasses, props.Class)

	// Build attributes
	attrs := []g.Node{
		html.Class(classes),
		g.Attr("data-orientation", props.Orientation),
	}

	// Add appropriate role based on decorative prop
	if props.Decorative {
		attrs = append(attrs, html.Role("none"))
	} else {
		attrs = append(attrs, html.Role("separator"))
		// Add aria-orientation for semantic separators
		attrs = append(attrs, g.Attr("aria-orientation", props.Orientation))
	}

	return html.Div(attrs...)
}

// Horizontal creates a horizontal separator (default)
func Horizontal() g.Node {
	return New(Props{
		Orientation: "horizontal",
		Decorative:  true,
	})
}

// Vertical creates a vertical separator
func Vertical() g.Node {
	return New(Props{
		Orientation: "vertical",
		Decorative:  true,
	})
}

// Semantic creates a semantic (non-decorative) horizontal separator
func Semantic() g.Node {
	return New(Props{
		Orientation: "horizontal",
		Decorative:  false,
	})
}

// SemanticVertical creates a semantic (non-decorative) vertical separator
func SemanticVertical() g.Node {
	return New(Props{
		Orientation: "vertical",
		Decorative:  false,
	})
}

// WithClass creates a separator with custom classes
func WithClass(class string) g.Node {
	return New(Props{
		Orientation: "horizontal",
		Decorative:  true,
		Class:       class,
	})
}