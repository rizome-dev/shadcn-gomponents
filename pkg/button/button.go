package button

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Button component
type Props struct {
	Variant  string // "default" | "destructive" | "outline" | "secondary" | "ghost" | "link"
	Size     string // "default" | "sm" | "lg" | "icon"
	Disabled bool
	Type     string // "button" | "submit" | "reset"
	Class    string // Additional custom classes
}

// buttonVariants defines the variant configuration for buttons
var buttonVariants = lib.VariantConfig{
	Base: "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
	Variants: map[string]map[string]string{
		"variant": {
			"default":     "bg-primary text-primary-foreground shadow-xs hover:bg-primary/90",
			"destructive": "bg-destructive text-white shadow-xs hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60",
			"outline":     "border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground dark:bg-input/30 dark:border-input dark:hover:bg-input/50",
			"secondary":   "bg-secondary text-secondary-foreground shadow-xs hover:bg-secondary/80",
			"ghost":       "hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50",
			"link":        "text-primary underline-offset-4 hover:underline",
		},
		"size": {
			"default": "h-9 px-4 py-2 has-[>svg]:px-3",
			"sm":      "h-8 rounded-md gap-1.5 px-3 has-[>svg]:px-2.5",
			"lg":      "h-10 rounded-md px-6 has-[>svg]:px-4",
			"icon":    "size-9",
		},
	},
	Defaults: map[string]string{
		"variant": "default",
		"size":    "default",
	},
}

// New creates a new Button component
func New(props Props, children ...g.Node) g.Node {
	// Set default type if not provided
	if props.Type == "" {
		props.Type = "button"
	}

	// Get variant classes
	classes := buttonVariants.GetClasses(lib.VariantProps{
		Variant: props.Variant,
		Size:    props.Size,
		Class:   props.Class,
	})

	// Build attributes
	attrs := []g.Node{
		html.Class(classes),
		html.Type(props.Type),
	}

	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}

	// Combine attributes and children
	return html.Button(append(attrs, children...)...)
}

// Default creates a button with default variant
func Default(children ...g.Node) g.Node {
	return New(Props{}, children...)
}

// Destructive creates a button with destructive variant
func Destructive(children ...g.Node) g.Node {
	return New(Props{Variant: "destructive"}, children...)
}

// Outline creates a button with outline variant
func Outline(children ...g.Node) g.Node {
	return New(Props{Variant: "outline"}, children...)
}

// Secondary creates a button with secondary variant
func Secondary(children ...g.Node) g.Node {
	return New(Props{Variant: "secondary"}, children...)
}

// Ghost creates a button with ghost variant
func Ghost(children ...g.Node) g.Node {
	return New(Props{Variant: "ghost"}, children...)
}

// LinkButton creates a button with link variant
func LinkButton(children ...g.Node) g.Node {
	return New(Props{Variant: "link"}, children...)
}

// Icon creates an icon-only button
func Icon(variant string, children ...g.Node) g.Node {
	return New(Props{
		Variant: variant,
		Size:    "icon",
	}, children...)
}