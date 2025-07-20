package toggle

import (
	"fmt"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Toggle component
type Props struct {
	ID          string   // HTML id attribute
	Pressed     bool     // Whether the toggle is pressed/on
	Disabled    bool     // Whether the toggle is disabled
	AriaLabel   string   // Accessibility label
	Variant     string   // "default" | "outline"
	Size        string   // "sm" | "default" | "lg"
	Class       string   // Additional custom classes
	OnClick     string   // JavaScript onClick handler
	DataState   string   // Override data-state attribute
	Attrs       []g.Node // Additional attributes to pass through
}

// toggleVariants defines the variant configuration for toggles
var toggleVariants = lib.VariantConfig{
	Base: "inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors hover:bg-muted hover:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0 gap-2 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground",
	Variants: map[string]map[string]string{
		"variant": {
			"default": "bg-transparent",
			"outline": "border border-input bg-transparent hover:bg-accent hover:text-accent-foreground",
		},
		"size": {
			"default": "h-10 px-3 min-w-10",
			"sm":      "h-9 px-2.5 min-w-9",
			"lg":      "h-11 px-5 min-w-11",
		},
	},
	Defaults: map[string]string{
		"variant": "default",
		"size":    "default",
	},
}

// New creates a new Toggle component
func New(props Props, children ...g.Node) g.Node {
	// Get variant classes
	classes := toggleVariants.GetClasses(lib.VariantProps{
		Variant: props.Variant,
		Size:    props.Size,
		Class:   props.Class,
	})

	// Determine data-state
	dataState := "off"
	if props.Pressed {
		dataState = "on"
	}
	if props.DataState != "" {
		dataState = props.DataState
	}

	// Build attributes
	attrs := []g.Node{
		html.Type("button"),
		html.Class(classes),
		g.Attr("data-state", dataState),
		html.Role("button"),
		g.Attr("aria-pressed", lib.CNIf(props.Pressed, "true", "false")),
	}

	// Add optional attributes
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}
	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}
	if props.AriaLabel != "" {
		attrs = append(attrs, g.Attr("aria-label", props.AriaLabel))
	}
	if props.OnClick != "" {
		attrs = append(attrs, g.Attr("onclick", props.OnClick))
	}
	
	// Add any additional attributes
	if len(props.Attrs) > 0 {
		attrs = append(attrs, props.Attrs...)
	}

	// Combine attributes and children
	return html.Button(append(attrs, children...)...)
}

// Default creates a toggle with default settings
func Default(children ...g.Node) g.Node {
	return New(Props{}, children...)
}

// Outline creates a toggle with outline variant
func Outline(children ...g.Node) g.Node {
	return New(Props{Variant: "outline"}, children...)
}

// Pressed creates a pre-pressed toggle
func Pressed(children ...g.Node) g.Node {
	return New(Props{Pressed: true}, children...)
}

// SmallToggle creates a small toggle
func SmallToggle(variant string, children ...g.Node) g.Node {
	return New(Props{
		Variant: variant,
		Size:    "sm",
	}, children...)
}

// Large creates a large toggle
func Large(variant string, children ...g.Node) g.Node {
	return New(Props{
		Variant: variant,
		Size:    "lg",
	}, children...)
}

// WithIcon creates a toggle with an icon and optional text
func WithIcon(icon g.Node, text string, pressed bool) g.Node {
	children := []g.Node{icon}
	if text != "" {
		children = append(children, g.Text(text))
	}
	
	return New(Props{
		Pressed:   pressed,
		AriaLabel: lib.CNIf(text == "", "Toggle", ""),
	}, children...)
}

// Icon creates an icon-only toggle with proper accessibility
func Icon(icon g.Node, ariaLabel string) g.Node {
	return New(Props{
		AriaLabel: ariaLabel,
	}, icon)
}

// Group creates a group of related toggles
func Group(toggles ...g.Node) g.Node {
	return html.Div(
		html.Class("flex items-center gap-1"),
		html.Role("group"),
		g.Group(toggles),
	)
}

// ToolbarItem creates a toggle suitable for toolbar usage
func ToolbarItem(icon g.Node, label string, pressed bool) g.Node {
	return New(Props{
		Variant:   "default",
		Size:      "sm",
		Pressed:   pressed,
		AriaLabel: label,
	}, icon)
}

// FormatButton creates a text formatting toggle (bold, italic, etc.)
func FormatButton(formatType string, icon g.Node, pressed bool) g.Node {
	return New(Props{
		Variant:   "default",
		Size:      "sm",
		Pressed:   pressed,
		AriaLabel: "Toggle " + formatType,
		OnClick:   fmt.Sprintf("document.execCommand('%s')", formatType),
	}, icon)
}

// ViewToggle creates a view mode toggle (grid/list view)
func ViewToggle(viewType string, icon g.Node, isActive bool) g.Node {
	return New(Props{
		Variant:   "outline",
		Size:      "sm",
		Pressed:   isActive,
		AriaLabel: viewType + " view",
	}, icon)
}

// GetToggleClasses returns the classes for a toggle button with the given properties
func GetToggleClasses(props Props) string {
	return toggleVariants.GetClasses(lib.VariantProps{
		Variant: props.Variant,
		Size:    props.Size,
		Class:   props.Class,
	})
}