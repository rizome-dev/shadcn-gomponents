package alert

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Alert component
type Props struct {
	Variant string // "default" | "destructive"
	Class   string // Additional custom classes
}

// alertVariants defines the variant configuration for alerts
var alertVariants = lib.VariantConfig{
	Base: "relative w-full rounded-lg border p-4 [&>svg]:absolute [&>svg]:left-4 [&>svg]:top-4 [&>svg]:text-foreground [&>svg+div]:pl-7",
	Variants: map[string]map[string]string{
		"variant": {
			"default":     "bg-card text-card-foreground",
			"destructive": "bg-card border-destructive/50 text-destructive dark:border-destructive [&>svg]:text-destructive",
		},
	},
	Defaults: map[string]string{
		"variant": "default",
	},
}

// New creates a new Alert component
func New(props Props, children ...g.Node) g.Node {
	// Get variant classes
	classes := alertVariants.GetClasses(lib.VariantProps{
		Variant: props.Variant,
		Class:   props.Class,
	})

	// Build attributes
	attrs := []g.Node{
		html.Class(classes),
		html.Role("alert"),
	}

	// Combine attributes and children
	return html.Div(append(attrs, children...)...)
}

// Default creates an alert with default variant
func Default(children ...g.Node) g.Node {
	return New(Props{}, children...)
}

// Destructive creates an alert with destructive variant
func Destructive(children ...g.Node) g.Node {
	return New(Props{Variant: "destructive"}, children...)
}

// Title creates an AlertTitle component
func Title(children ...g.Node) g.Node {
	return html.H5(
		html.Class("mb-1 font-medium leading-none tracking-tight"),
		g.Group(children),
	)
}

// Description creates an AlertDescription component
func Description(children ...g.Node) g.Node {
	return html.Div(
		html.Class("text-sm text-muted-foreground [&_p]:leading-relaxed"),
		g.Group(children),
	)
}

// WithIcon creates an alert with an icon
func WithIcon(icon g.Node, props Props, children ...g.Node) g.Node {
	return New(props, append([]g.Node{icon}, children...)...)
}