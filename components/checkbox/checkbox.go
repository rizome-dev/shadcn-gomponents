package checkbox

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Checkbox component
type Props struct {
	ID             string // HTML id attribute
	Name           string // Form field name
	Value          string // Form field value
	Checked        bool   // Whether the checkbox is checked
	Indeterminate  bool   // Whether the checkbox is in indeterminate state
	Disabled       bool   // Whether the checkbox is disabled
	Required       bool   // Whether the checkbox is required
	Class          string // Additional custom classes
	OnChange       string // JavaScript onChange handler
}

// New creates a new Checkbox component
func New(props Props) g.Node {
	// Base classes for the checkbox
	classes := lib.CN(
		"peer h-4 w-4 shrink-0 rounded-sm border border-primary ring-offset-background",
		"focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
		"disabled:cursor-not-allowed disabled:opacity-50",
		"data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground",
		props.Class,
	)

	// Build attributes
	attrs := []g.Node{
		html.Type("checkbox"),
		html.Class(classes),
	}

	// Add optional attributes
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}
	if props.Name != "" {
		attrs = append(attrs, html.Name(props.Name))
	}
	if props.Value != "" {
		attrs = append(attrs, html.Value(props.Value))
	}
	if props.Checked {
		attrs = append(attrs, html.Checked())
		attrs = append(attrs, g.Attr("data-state", "checked"))
	} else {
		attrs = append(attrs, g.Attr("data-state", "unchecked"))
	}
	if props.Indeterminate {
		// Note: indeterminate is a JavaScript property, not an HTML attribute
		// This would need to be set via JavaScript
		attrs = append(attrs, g.Attr("data-indeterminate", "true"))
	}
	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}
	if props.Required {
		attrs = append(attrs, html.Required())
	}
	if props.OnChange != "" {
		attrs = append(attrs, g.Attr("onchange", props.OnChange))
	}

	// Create checkbox wrapper div that mimics Radix UI structure
	wrapperClasses := lib.CN(
		"relative inline-flex items-center justify-center",
		"h-4 w-4 shrink-0 rounded-sm border border-primary shadow-sm",
		"focus-within:outline-none focus-within:ring-2 focus-within:ring-ring focus-within:ring-offset-2",
		"has-[:disabled]:cursor-not-allowed has-[:disabled]:opacity-50",
		"has-[:checked]:bg-primary has-[:checked]:text-primary-foreground",
	)

	divAttrs := []g.Node{
		html.Class(wrapperClasses),
		g.Attr("role", "checkbox"),
	}
	
	if props.Checked {
		divAttrs = append(divAttrs, g.Attr("aria-checked", "true"))
	} else {
		divAttrs = append(divAttrs, g.Attr("aria-checked", "false"))
	}
	
	if props.Disabled {
		divAttrs = append(divAttrs, g.Attr("aria-disabled", "true"))
	}
	
	divAttrs = append(divAttrs, html.Input(attrs...))
	
	// Check icon (shown when checked)
	if props.Checked || props.Indeterminate {
		divAttrs = append(divAttrs, renderCheckIcon(props.Indeterminate))
	}
	
	return html.Div(divAttrs...)
}

// Default creates a checkbox with default settings
func Default() g.Node {
	return New(Props{})
}

// Checked creates a pre-checked checkbox
func Checked() g.Node {
	return New(Props{Checked: true})
}

// Disabled creates a disabled checkbox
func Disabled(checked bool) g.Node {
	return New(Props{
		Disabled: true,
		Checked:  checked,
	})
}

// WithLabel creates a checkbox with a label
func WithLabel(props Props, labelText string) g.Node {
	// Generate ID if not provided
	if props.ID == "" {
		props.ID = "checkbox-" + labelText // Simple ID generation
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

// FormField creates a checkbox within a form field structure
func FormField(props Props, labelText, description string) g.Node {
	// Generate ID if not provided
	if props.ID == "" {
		props.ID = "checkbox-field"
	}

	return html.Div(
		html.Class("space-y-2"),
		html.Div(
			html.Class("flex items-center space-x-2"),
			New(props),
			html.Div(html.Class("space-y-1 leading-none"),
				html.Label(
					html.For(props.ID),
					html.Class("text-sm font-medium"),
					g.Text(labelText),
				),
				g.If(description != "",
					html.P(html.Class("text-sm text-muted-foreground"), g.Text(description)),
				),
			),
		),
	)
}

// renderCheckIcon renders the appropriate icon for checked/indeterminate state
func renderCheckIcon(indeterminate bool) g.Node {
	classes := "absolute pointer-events-none h-4 w-4 text-current"
	
	if indeterminate {
		// Minus icon for indeterminate state
		return g.El("svg",
			g.Attr("xmlns", "http://www.w3.org/2000/svg"),
			g.Attr("viewBox", "0 0 24 24"),
			g.Attr("fill", "none"),
			g.Attr("stroke", "currentColor"),
			g.Attr("stroke-width", "3"),
			g.Attr("stroke-linecap", "round"),
			g.Attr("stroke-linejoin", "round"),
			html.Class(classes),
			g.El("line", g.Attr("x1", "5"), g.Attr("y1", "12"), g.Attr("x2", "19"), g.Attr("y2", "12")),
		)
	}
	
	// Check icon for checked state
	return g.El("svg",
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "3"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
		html.Class(classes),
		g.El("polyline", g.Attr("points", "20 6 9 17 4 12")),
	)
}

// Group creates a group of checkboxes
func Group(title string, checkboxes ...g.Node) g.Node {
	return html.Div(
		html.Class("space-y-3"),
		g.If(title != "",
			html.H3(html.Class("text-sm font-medium"), g.Text(title)),
		),
		html.Div(
			html.Class("space-y-2"),
			g.Group(checkboxes),
		),
	)
}