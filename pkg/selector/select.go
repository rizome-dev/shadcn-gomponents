package selector

import (
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Select component
type Props struct {
	ID          string   // HTML id attribute
	Name        string   // Form field name
	Value       string   // Selected value
	Placeholder string   // Placeholder text
	Options     []OptionType // Available options
	Groups      []Group  // Grouped options
	Disabled    bool     // Whether the select is disabled
	Required    bool     // Whether the select is required
	Multiple    bool     // Whether multiple selection is allowed
	Size        string   // "sm" | "default" | "lg"
	Class       string   // Additional custom classes
	OnChange    string   // JavaScript onChange handler
}

// Option defines a select option
type OptionType struct {
	Value    string
	Label    string
	Disabled bool
	Selected bool
}

// Group defines a group of options
type Group struct {
	Label   string
	Options []OptionType
}

// New creates a new Select component (native HTML select)
func New(props Props) g.Node {
	// Determine size classes
	sizeClasses := "h-10 px-3 py-2"
	switch props.Size {
	case "sm":
		sizeClasses = "h-9 px-3 py-1.5 text-sm"
	case "lg":
		sizeClasses = "h-11 px-4 py-2.5"
	}

	// Base classes for the select
	classes := lib.CN(
		"flex w-full rounded-md border border-input bg-background",
		sizeClasses,
		"text-sm ring-offset-background",
		"focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
		"disabled:cursor-not-allowed disabled:opacity-50",
		"[&>option]:bg-background [&>option]:text-foreground",
		props.Class,
	)

	// Build attributes
	attrs := []g.Node{
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
	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}
	if props.Required {
		attrs = append(attrs, html.Required())
	}
	if props.Multiple {
		attrs = append(attrs, html.Multiple())
	}
	if props.OnChange != "" {
		attrs = append(attrs, g.Attr("onchange", props.OnChange))
	}

	// Build children (options)
	var children []g.Node

	// Add placeholder if provided
	if props.Placeholder != "" {
		children = append(children, 
			html.Option(
				html.Value(""),
				g.If(!props.Required, html.Selected()),
				html.Disabled(),
				g.Attr("hidden", ""),
				g.Text(props.Placeholder),
			),
		)
	}

	// Add regular options
	for _, opt := range props.Options {
		children = append(children, renderOption(opt, props.Value))
	}

	// Add grouped options
	for _, group := range props.Groups {
		groupChildren := []g.Node{
			g.Attr("label", group.Label),
		}
		for _, opt := range group.Options {
			groupChildren = append(groupChildren, renderOption(opt, props.Value))
		}
		children = append(children, html.OptGroup(groupChildren...))
	}

	return html.Select(append(attrs, children...)...)
}

// renderOption renders a single option
func renderOption(opt OptionType, selectedValue string) g.Node {
	attrs := []g.Node{
		html.Value(opt.Value),
	}
	
	// Check if this option should be selected
	if opt.Selected || (selectedValue != "" && opt.Value == selectedValue) {
		attrs = append(attrs, html.Selected())
	}
	
	if opt.Disabled {
		attrs = append(attrs, html.Disabled())
	}
	
	attrs = append(attrs, g.Text(opt.Label))
	
	return html.Option(attrs...)
}

// Simple creates a simple select with options
func Simple(name string, options []OptionType, defaultValue string) g.Node {
	return New(Props{
		Name:    name,
		Value:   defaultValue,
		Options: options,
	})
}

// WithPlaceholder creates a select with a placeholder
func WithPlaceholder(name, placeholder string, options []OptionType) g.Node {
	return New(Props{
		Name:        name,
		Placeholder: placeholder,
		Options:     options,
	})
}

// WithGroups creates a select with grouped options
func WithGroups(name string, groups []Group, defaultValue string) g.Node {
	return New(Props{
		Name:   name,
		Value:  defaultValue,
		Groups: groups,
	})
}

// FormField creates a select within a form field structure
func FormField(props Props, label, description string) g.Node {
	// Generate ID if not provided
	if props.ID == "" {
		props.ID = "select-" + props.Name
	}

	return html.Div(
		html.Class("space-y-2"),
		g.If(label != "",
			html.Label(
				html.For(props.ID),
				html.Class("text-sm font-medium"),
				g.Text(label),
			),
		),
		New(props),
		g.If(description != "",
			html.P(html.Class("text-sm text-muted-foreground"), g.Text(description)),
		),
	)
}

// Custom creates a custom styled select wrapper
// This mimics the Radix UI Select trigger appearance but uses native select
func Custom(props Props) g.Node {
	// Generate ID if not provided
	if props.ID == "" && props.Name != "" {
		props.ID = "select-custom-" + props.Name
	}

	// Create wrapper that looks like SelectTrigger
	disabledClass := ""
	if props.Disabled {
		disabledClass = "opacity-50"
	}
	wrapperClasses := lib.CN(
		"relative",
		disabledClass,
	)

	// Style the select to be invisible but functional
	cursorClass := ""
	if props.Disabled {
		cursorClass = "cursor-not-allowed"
	}
	selectProps := props
	selectProps.Class = lib.CN(
		"absolute inset-0 w-full h-full opacity-0 cursor-pointer",
		cursorClass,
	)

	// Determine display text
	displayText := props.Placeholder
	if props.Value != "" {
		// Find the label for the selected value
		for _, opt := range props.Options {
			if opt.Value == props.Value {
				displayText = opt.Label
				break
			}
		}
		// Check in groups if not found
		if displayText == props.Placeholder {
			for _, group := range props.Groups {
				for _, opt := range group.Options {
					if opt.Value == props.Value {
						displayText = opt.Label
						break
					}
				}
			}
		}
	}

	// Trigger button classes
	textColorClass := ""
	if props.Value == "" {
		textColorClass = "text-muted-foreground"
	}
	triggerClasses := lib.CN(
		"flex h-10 w-full items-center justify-between rounded-md border",
		"border-input bg-background px-3 py-2 text-sm ring-offset-background",
		"pointer-events-none", // Let select handle clicks
		textColorClass,
	)

	return html.Div(
		html.Class(wrapperClasses),
		// Visual trigger
		html.Div(
			html.Class(triggerClasses),
			html.Span(html.Class("line-clamp-1"), g.Text(displayText)),
			// Chevron down icon
			g.El("svg",
				g.Attr("xmlns", "http://www.w3.org/2000/svg"),
				g.Attr("viewBox", "0 0 24 24"),
				g.Attr("fill", "none"),
				g.Attr("stroke", "currentColor"),
				g.Attr("stroke-width", "2"),
				g.Attr("stroke-linecap", "round"),
				g.Attr("stroke-linejoin", "round"),
				html.Class("h-4 w-4 opacity-50"),
				g.El("polyline", g.Attr("points", "6 9 12 15 18 9")),
			),
		),
		// Actual select (invisible)
		New(selectProps),
	)
}