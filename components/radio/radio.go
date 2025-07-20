package radio

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// GroupProps defines properties for the RadioGroup container
type GroupProps struct {
	Name         string // Form field name (required for radio groups)
	DefaultValue string // Default selected value
	Class        string // Additional custom classes
	Orientation  string // "vertical" | "horizontal" (default: vertical)
}

// ItemProps defines properties for individual RadioGroupItem
type ItemProps struct {
	ID       string // HTML id attribute
	Value    string // The value of this radio option
	Checked  bool   // Whether this radio is selected
	Disabled bool   // Whether this radio is disabled
	Class    string // Additional custom classes
}

// Group creates a new RadioGroup container
func Group(props GroupProps, children ...g.Node) g.Node {
	// Default orientation is vertical
	classes := "grid gap-2"
	
	if props.Orientation == "horizontal" {
		classes = "flex items-center gap-4"
	}
	
	classes = lib.CN(classes, props.Class)

	return html.Div(
		html.Class(classes),
		html.Role("radiogroup"),
		g.If(props.Name != "", g.Attr("data-name", props.Name)),
		g.If(props.DefaultValue != "", g.Attr("data-default-value", props.DefaultValue)),
		g.Group(children),
	)
}

// Item creates a new RadioGroupItem
func Item(props ItemProps, groupName string) g.Node {
	// Base classes for the radio item
	classes := lib.CN(
		"aspect-square h-4 w-4 rounded-full border border-primary text-primary",
		"ring-offset-background focus:outline-none focus-visible:ring-2",
		"focus-visible:ring-ring focus-visible:ring-offset-2",
		"disabled:cursor-not-allowed disabled:opacity-50",
		props.Class,
	)

	// Build attributes
	attrs := []g.Node{
		html.Type("radio"),
		html.Class(classes),
		html.Name(groupName),
		html.Value(props.Value),
	}

	// Add optional attributes
	if props.ID != "" {
		attrs = append(attrs, html.ID(props.ID))
	}
	if props.Checked {
		attrs = append(attrs, html.Checked())
		attrs = append(attrs, g.Attr("data-state", "checked"))
	} else {
		attrs = append(attrs, g.Attr("data-state", "unchecked"))
	}
	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}

	// Create wrapper that mimics Radix UI structure
	wrapperClasses := lib.CN(
		"relative inline-flex items-center justify-center",
		"h-4 w-4 rounded-full",
	)

	return html.Div(
		html.Class(wrapperClasses),
		html.Input(attrs...),
		// Indicator circle (shown when checked)
		g.If(props.Checked,
			renderIndicator(),
		),
	)
}

// WithLabel creates a radio item with a label
func WithLabel(props ItemProps, groupName, labelText string) g.Node {
	// Generate ID if not provided
	if props.ID == "" {
		props.ID = "radio-" + props.Value
	}

	return html.Div(
		html.Class("flex items-center space-x-2"),
		Item(props, groupName),
		html.Label(
			html.For(props.ID),
			html.Class("text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"),
			g.Text(labelText),
		),
	)
}

// renderIndicator renders the filled circle for selected state
func renderIndicator() g.Node {
	return html.Div(
		html.Class("absolute pointer-events-none inset-0 flex items-center justify-center"),
		g.El("svg",
			g.Attr("viewBox", "0 0 16 16"),
			g.Attr("fill", "currentColor"),
			html.Class("h-2.5 w-2.5 text-primary"),
			g.El("circle", g.Attr("cx", "8"), g.Attr("cy", "8"), g.Attr("r", "3")),
		),
	)
}

// Simple creates a simple radio group with options
func Simple(name string, options []Option, defaultValue string) g.Node {
	var items []g.Node
	
	for _, opt := range options {
		props := ItemProps{
			Value:    opt.Value,
			Checked:  opt.Value == defaultValue,
			Disabled: opt.Disabled,
		}
		items = append(items, WithLabel(props, name, opt.Label))
	}

	return Group(
		GroupProps{
			Name:         name,
			DefaultValue: defaultValue,
		},
		items...,
	)
}

// Option defines a radio option
type Option struct {
	Value    string
	Label    string
	Disabled bool
}

// Horizontal creates a horizontal radio group
func Horizontal(name string, options []Option, defaultValue string) g.Node {
	var items []g.Node
	
	for _, opt := range options {
		props := ItemProps{
			Value:    opt.Value,
			Checked:  opt.Value == defaultValue,
			Disabled: opt.Disabled,
		}
		items = append(items, WithLabel(props, name, opt.Label))
	}

	return Group(
		GroupProps{
			Name:         name,
			DefaultValue: defaultValue,
			Orientation:  "horizontal",
		},
		items...,
	)
}

// FormField creates a radio group within a form field structure
func FormField(name, label, description string, options []Option, defaultValue string) g.Node {
	return html.Div(
		html.Class("space-y-3"),
		html.Div(
			g.If(label != "",
				html.Label(html.Class("text-sm font-medium"), g.Text(label)),
			),
			g.If(description != "",
				html.P(html.Class("text-sm text-muted-foreground"), g.Text(description)),
			),
		),
		Simple(name, options, defaultValue),
	)
}

// Card creates a card-style radio option
func Card(props ItemProps, groupName string, title, description string) g.Node {
	// Generate ID if not provided
	if props.ID == "" {
		props.ID = "radio-card-" + props.Value
	}

	checkedClasses := ""
	if props.Checked {
		checkedClasses = "border-primary bg-accent"
	} else {
		checkedClasses = "border-muted"
	}
	
	disabledClasses := ""
	if props.Disabled {
		disabledClasses = "opacity-50 cursor-not-allowed"
	}
	
	cardClasses := lib.CN(
		"relative flex cursor-pointer rounded-lg border p-4",
		"hover:bg-accent",
		checkedClasses,
		disabledClasses,
	)

	return html.Label(
		html.For(props.ID),
		html.Class(cardClasses),
		html.Div(html.Class("flex items-start space-x-3"),
			Item(props, groupName),
			html.Div(html.Class("space-y-1"),
				html.P(html.Class("font-medium leading-none"), g.Text(title)),
				g.If(description != "",
					html.P(html.Class("text-sm text-muted-foreground"), g.Text(description)),
				),
			),
		),
	)
}