package textarea

import (
	"fmt"
	
	g "maragu.dev/gomponents"
	html "maragu.dev/gomponents/html"
	"github.com/rizome-dev/shadcn-gomponents/lib"
)

// Props defines the properties for the Textarea component
type Props struct {
	ID          string // HTML id attribute
	Name        string // Form field name
	Value       string // Initial value
	Placeholder string // Placeholder text
	Rows        int    // Number of visible rows
	Cols        int    // Number of visible columns
	MaxLength   int    // Maximum character length
	MinLength   int    // Minimum character length
	Disabled    bool   // Whether the textarea is disabled
	Required    bool   // Whether the textarea is required
	ReadOnly    bool   // Whether the textarea is read-only
	AutoResize  bool   // Whether to enable auto-resize
	Resize      string // CSS resize property: "none" | "both" | "horizontal" | "vertical"
	Class       string // Additional custom classes
	OnChange    string // JavaScript onChange handler
	OnInput     string // JavaScript onInput handler
}

// New creates a new Textarea component
func New(props Props) g.Node {
	// Base classes
	baseClasses := "flex min-h-[80px] w-full rounded-md border border-input bg-background px-3 py-2"
	
	// Additional styling classes
	styleClasses := "text-sm ring-offset-background placeholder:text-muted-foreground"
	
	// Focus and state classes
	focusClasses := "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2"
	stateClasses := "disabled:cursor-not-allowed disabled:opacity-50"
	
	// Auto-resize class (CSS field-sizing-content)
	autoResizeClass := ""
	if props.AutoResize {
		autoResizeClass = "field-sizing-content"
	}
	
	// Resize class
	resizeClass := ""
	switch props.Resize {
	case "none":
		resizeClass = "resize-none"
	case "both":
		resizeClass = "resize"
	case "horizontal":
		resizeClass = "resize-x"
	case "vertical":
		resizeClass = "resize-y"
	}
	
	// Combine all classes
	classes := lib.CN(
		baseClasses,
		styleClasses,
		focusClasses,
		stateClasses,
		autoResizeClass,
		resizeClass,
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
	if props.Placeholder != "" {
		attrs = append(attrs, html.Placeholder(props.Placeholder))
	}
	if props.Rows > 0 {
		attrs = append(attrs, g.Attr("rows", fmt.Sprintf("%d", props.Rows)))
	}
	if props.Cols > 0 {
		attrs = append(attrs, g.Attr("cols", fmt.Sprintf("%d", props.Cols)))
	}
	if props.MaxLength > 0 {
		attrs = append(attrs, g.Attr("maxlength", fmt.Sprintf("%d", props.MaxLength)))
	}
	if props.MinLength > 0 {
		attrs = append(attrs, g.Attr("minlength", fmt.Sprintf("%d", props.MinLength)))
	}
	if props.Disabled {
		attrs = append(attrs, html.Disabled())
	}
	if props.Required {
		attrs = append(attrs, html.Required())
	}
	if props.ReadOnly {
		attrs = append(attrs, html.ReadOnly())
	}
	if props.OnChange != "" {
		attrs = append(attrs, g.Attr("onchange", props.OnChange))
	}
	if props.OnInput != "" {
		attrs = append(attrs, g.Attr("oninput", props.OnInput))
	}

	// Add the value as a child text node if provided
	if props.Value != "" {
		return html.Textarea(append(attrs, g.Text(props.Value))...)
	}

	return html.Textarea(attrs...)
}

// Default creates a textarea with default settings
func Default() g.Node {
	return New(Props{})
}

// WithPlaceholder creates a textarea with placeholder text
func WithPlaceholder(placeholder string) g.Node {
	return New(Props{
		Placeholder: placeholder,
	})
}

// AutoResize creates a textarea that automatically resizes based on content
func AutoResize() g.Node {
	return New(Props{
		AutoResize: true,
		Resize:     "none",
	})
}

// NoResize creates a textarea that cannot be manually resized
func NoResize() g.Node {
	return New(Props{
		Resize: "none",
	})
}

// FormField creates a textarea within a form field structure
func FormField(props Props, label, description string) g.Node {
	// Generate ID if not provided
	if props.ID == "" && props.Name != "" {
		props.ID = "textarea-" + props.Name
	}

	return html.Div(
		html.Class("space-y-2"),
		g.If(label != "",
			html.Label(
				g.If(props.ID != "", html.For(props.ID)),
				html.Class("text-sm font-medium"),
				g.Text(label),
				g.If(props.Required, 
					html.Span(html.Class("text-destructive ml-1"), g.Text("*")),
				),
			),
		),
		New(props),
		g.If(description != "",
			html.P(html.Class("text-sm text-muted-foreground"), g.Text(description)),
		),
	)
}

// WithCharacterCount creates a textarea with a character counter
func WithCharacterCount(props Props, label string) g.Node {
	// Generate ID if not provided
	if props.ID == "" && props.Name != "" {
		props.ID = "textarea-count-" + props.Name
	}
	
	countID := props.ID + "-count"
	
	// Add onInput handler to update character count
	props.OnInput = fmt.Sprintf("document.getElementById('%s').textContent = this.value.length", countID)

	return html.Div(
		html.Class("space-y-2"),
		g.If(label != "",
			html.Div(
				html.Class("flex items-center justify-between"),
				html.Label(
					g.If(props.ID != "", html.For(props.ID)),
					html.Class("text-sm font-medium"),
					g.Text(label),
				),
				g.If(props.MaxLength > 0,
					html.Span(
						html.Class("text-sm text-muted-foreground"),
						html.Span(html.ID(countID), g.Text("0")),
						g.Textf("/%d", props.MaxLength),
					),
				),
			),
		),
		New(props),
	)
}

// Message creates a simple message/comment textarea
func Message(name string, rows int) g.Node {
	return New(Props{
		Name:        name,
		Placeholder: "Type your message here...",
		Rows:        rows,
		Resize:      "none",
	})
}

// Bio creates a textarea suitable for biography/description fields
func Bio(name string, maxLength int) g.Node {
	return WithCharacterCount(
		Props{
			Name:        name,
			Placeholder: "Tell us a little bit about yourself",
			Rows:        4,
			MaxLength:   maxLength,
			AutoResize:  true,
		},
		"Bio",
	)
}