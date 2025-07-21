package switchcomp

import (
	"fmt"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Switch component
type Props struct {
	ID       string // HTML id attribute
	Name     string // Form field name
	Value    string // Form field value
	Checked  bool   // Whether the switch is on
	Disabled bool   // Whether the switch is disabled
	Required bool   // Whether the switch is required
	Size     string // "sm" | "default" | "lg"
	Class    string // Additional custom classes
	OnChange string // JavaScript onChange handler
}

// New creates a new Switch component
func New(props Props) g.Node {
	// Determine size classes
	var trackSize, thumbSize, thumbTranslate string
	switch props.Size {
	case "sm":
		trackSize = "h-5 w-9"
		thumbSize = "h-4 w-4"
		thumbTranslate = "translate-x-4"
	case "lg":
		trackSize = "h-7 w-14"
		thumbSize = "h-6 w-6"
		thumbTranslate = "translate-x-7"
	default: // default size
		trackSize = "h-6 w-11"
		thumbSize = "h-5 w-5"
		thumbTranslate = "translate-x-5"
	}

	// Track classes - includes state-based background colors
	trackClasses := lib.CN(
		"peer inline-flex shrink-0 cursor-pointer items-center rounded-full",
		"border-2 border-transparent transition-colors",
		"focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:ring-offset-background",
		"disabled:cursor-not-allowed disabled:opacity-50",
		trackSize,
		props.Class,
	)

	// Background color based on checked state
	var trackBgClass string
	if props.Checked {
		trackBgClass = "bg-primary"
	} else {
		trackBgClass = "bg-input"
	}

	// Thumb classes with positioning
	thumbClasses := lib.CN(
		"pointer-events-none block rounded-full bg-background shadow-lg ring-0 transition-transform",
		thumbSize,
	)

	// Thumb transform based on checked state
	var thumbTransform string
	if props.Checked {
		thumbTransform = thumbTranslate
	} else {
		thumbTransform = "translate-x-0"
	}

	// Build checkbox attributes
	checkboxAttrs := []g.Node{
		html.Type("checkbox"),
		html.Class("sr-only peer"), // Screen reader only, peer for CSS
	}

	// Add optional attributes
	if props.ID != "" {
		checkboxAttrs = append(checkboxAttrs, html.ID(props.ID))
	}
	if props.Name != "" {
		checkboxAttrs = append(checkboxAttrs, html.Name(props.Name))
	}
	if props.Value != "" {
		checkboxAttrs = append(checkboxAttrs, html.Value(props.Value))
	}
	if props.Checked {
		checkboxAttrs = append(checkboxAttrs, html.Checked())
	}
	if props.Disabled {
		checkboxAttrs = append(checkboxAttrs, html.Disabled())
	}
	if props.Required {
		checkboxAttrs = append(checkboxAttrs, html.Required())
	}
	if props.OnChange != "" {
		checkboxAttrs = append(checkboxAttrs, g.Attr("onchange", props.OnChange))
	}

	return html.Label(
		g.If(props.ID != "", html.For(props.ID)),
		html.Class("relative inline-block"),
		html.Input(checkboxAttrs...),
		html.Span(
			html.Class(lib.CN(trackClasses, trackBgClass)),
			html.Role("switch"),
			g.Attr("aria-checked", func() string {
				if props.Checked {
					return "true"
				}
				return "false"
			}()),
			g.If(props.Disabled, g.Attr("aria-disabled", "true")),
			html.Span(
				html.Class(thumbClasses),
				g.Attr("style", fmt.Sprintf("transform: %s;", thumbTransform)),
			),
		),
	)
}

// Default creates a switch with default settings
func Default() g.Node {
	return New(Props{})
}

// Checked creates a pre-checked switch
func Checked() g.Node {
	return New(Props{Checked: true})
}

// Disabled creates a disabled switch
func Disabled(checked bool) g.Node {
	return New(Props{
		Disabled: true,
		Checked:  checked,
	})
}

// Small creates a small switch
func SmallComponent() g.Node {
	return New(Props{Size: "sm"})
}

// Large creates a large switch
func Large() g.Node {
	return New(Props{Size: "lg"})
}

// WithLabel creates a switch with a label
func WithLabel(props Props, labelText string) g.Node {
	// Generate ID if not provided
	if props.ID == "" {
		props.ID = "switch-" + labelText
	}

	return html.Div(
		html.Class("flex items-center space-x-2"),
		New(props),
		html.Label(
			html.For(props.ID),
			html.Class("text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"),
			g.Text(labelText),
		),
	)
}

// FormField creates a switch within a form field structure
func FormField(props Props, labelText, description string) g.Node {
	// Generate ID if not provided
	if props.ID == "" {
		props.ID = "switch-field"
	}

	return html.Div(
		html.Class("flex flex-row items-center justify-between rounded-lg border p-4"),
		html.Div(html.Class("space-y-0.5"),
			html.Label(
				html.For(props.ID),
				html.Class("text-base font-medium"),
				g.Text(labelText),
			),
			g.If(description != "",
				html.P(html.Class("text-sm text-muted-foreground"), g.Text(description)),
			),
		),
		New(props),
	)
}

// Setting creates a switch in a settings-style layout
func Setting(name, label, description string, checked bool) g.Node {
	return FormField(
		Props{
			Name:    name,
			Checked: checked,
		},
		label,
		description,
	)
}