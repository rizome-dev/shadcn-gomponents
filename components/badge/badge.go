package badge

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Badge component
type Props struct {
	Variant string // "default" | "secondary" | "destructive" | "outline"
	Class   string // Additional custom classes
}

// badgeVariants defines the variant configuration for badges
var badgeVariants = lib.VariantConfig{
	Base: "inline-flex items-center rounded-md border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
	Variants: map[string]map[string]string{
		"variant": {
			"default":     "border-transparent bg-primary text-primary-foreground shadow hover:bg-primary/80",
			"secondary":   "border-transparent bg-secondary text-secondary-foreground hover:bg-secondary/80",
			"destructive": "border-transparent bg-destructive text-destructive-foreground shadow hover:bg-destructive/80",
			"outline":     "text-foreground",
		},
	},
	Defaults: map[string]string{
		"variant": "default",
	},
}

// New creates a new Badge component
func New(props Props, children ...g.Node) g.Node {
	// Get variant classes
	classes := badgeVariants.GetClasses(lib.VariantProps{
		Variant: props.Variant,
		Class:   props.Class,
	})

	// Build attributes
	attrs := []g.Node{
		html.Class(classes),
	}

	// Combine attributes and children
	return html.Div(append(attrs, children...)...)
}

// Default creates a badge with default variant
func Default(children ...g.Node) g.Node {
	return New(Props{}, children...)
}

// Secondary creates a badge with secondary variant
func Secondary(children ...g.Node) g.Node {
	return New(Props{Variant: "secondary"}, children...)
}

// Destructive creates a badge with destructive variant
func Destructive(children ...g.Node) g.Node {
	return New(Props{Variant: "destructive"}, children...)
}

// Outline creates a badge with outline variant
func Outline(children ...g.Node) g.Node {
	return New(Props{Variant: "outline"}, children...)
}

// WithIcon creates a badge with an icon
func WithIcon(icon g.Node, variant string, text string) g.Node {
	// Style the icon to match v4 behavior
	styledIcon := g.Group{
		icon,
		html.Class("h-3 w-3"),
	}
	
	return New(
		Props{
			Variant: variant,
			Class:   "gap-1",
		},
		styledIcon,
		g.Text(text),
	)
}

// Link creates a badge that functions as a link
func LinkComponent(href string, props Props, children ...g.Node) g.Node {
	// Get variant classes
	classes := badgeVariants.GetClasses(lib.VariantProps{
		Variant: props.Variant,
		Class:   props.Class,
	})
	
	// Add hover effect for links
	classes = lib.CN(classes, "hover:underline")

	return html.A(
		html.Href(href),
		html.Class(classes),
		g.Group(children),
	)
}